#!/bin/bash

# WARNING: DO NOT EDIT, THIS FILE IS PROBABLY A COPY
#
# The original version of this file is located in the https://github.com/istio/common-files repo.
# If you're looking at this file in a different repo and want to make a change, please go to the
# common-files repo, make the change there and check it in. Then come back to this repo and run
# "make update-common".

# Copyright Istio Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e
set -x

# The purpose of this file is to unify ocp setup in both istio/istio and istio-ecosystem/sail-operator.
# repos to avoid code duplication. This is needed to setup the OCP environment for the tests.

WD=$(dirname "$0")
WD=$(cd "$WD"; pwd)
TIMEOUT=300
export NAMESPACE="${NAMESPACE:-"istio-system"}"
SAIL_REPO_URL="https://github.com/istio-ecosystem/sail-operator.git"
SAIL_OPERATOR_BRANCH="${SAIL_OPERATOR_BRANCH:-}"  # Will be auto-detected if not set
IBM="${IBM:-"false"}"

function setup_internal_registry() {
  # Validate that the internal registry is running in the OCP Cluster, configure the variable to be used in the make target.
  # If there is no internal registry, the test can't be executed targeting to the internal registry

  # For multicluster mode, setup registry in all clusters
  if [[ "${TOPOLOGY:-SINGLE_CLUSTER}" != "SINGLE_CLUSTER" ]]; then
    echo "Setting up internal registry for multicluster mode..."

    # Get available contexts for each cluster
    local -a cluster_contexts
    for cluster_name in "${CLUSTER_NAMES[@]}"; do
      local cluster_context=""
      mapfile -t available_contexts < <(kubectl config get-contexts -o name 2>/dev/null || true)
      for context in "${available_contexts[@]}"; do
        if [[ "${context}" == "${cluster_name}" ]] || [[ "${context}" == *"/${cluster_name}/"* ]] || [[ "${context}" == *"-${cluster_name}-"* ]]; then
          cluster_context="${context}"
          break
        fi
      done
      cluster_contexts+=("${cluster_context}")
    done

    # Setup registry in each cluster
    for i in "${!CLUSTER_NAMES[@]}"; do
      local cluster_name="${CLUSTER_NAMES[i]}"
      local context="${cluster_contexts[i]}"

      echo "Setting up registry in cluster ${cluster_name}..."

      # Check if the registry pods are running
      if oc --context="${context}" get pods -n openshift-image-registry --no-headers 2>/dev/null | grep -v "Running\|Completed"; then
        echo "Warning: Image registry in cluster ${cluster_name} is not fully running"
      fi

      # Check if default route already exist
      if [ -z "$(oc --context="${context}" get route default-route -n openshift-image-registry -o name 2>/dev/null)" ]; then
        echo "Route default-route does not exist in cluster ${cluster_name}, patching DefaultRoute to true on Image Registry."
        oc --context="${context}" patch configs.imageregistry.operator.openshift.io/cluster --patch '{"spec":{"defaultRoute":true}}' --type=merge

        timeout --foreground -v -s SIGHUP -k ${TIMEOUT} ${TIMEOUT} bash --verbose -c \
          "until oc --context=${context} get route default-route -n openshift-image-registry &> /dev/null; do sleep 5; done && echo 'The default-route has been created in cluster ${cluster_name}.'"
      fi

      # Create namespace in each cluster
      oc --context="${context}" create namespace "${NAMESPACE}" 2>/dev/null || true

      # Deploy rolebinding for cross-cluster image pulls
      echo "Deploying rolebindings in cluster ${cluster_name}..."
      echo '
kind: List
apiVersion: v1
items:
- apiVersion: rbac.authorization.k8s.io/v1
  kind: RoleBinding
  metadata:
    name: image-puller
    namespace: '"${NAMESPACE}"'
  roleRef:
    apiGroup: rbac.authorization.k8s.io
    kind: ClusterRole
    name: system:image-puller
  subjects:
  - kind: Group
    apiGroup: rbac.authorization.k8s.io
    name: system:unauthenticated
  - kind: Group
    name: system:serviceaccounts
    apiGroup: rbac.authorization.k8s.io
- apiVersion: rbac.authorization.k8s.io/v1
  kind: RoleBinding
  metadata:
    name: image-pusher
    namespace: '"${NAMESPACE}"'
  roleRef:
    apiGroup: rbac.authorization.k8s.io
    kind: ClusterRole
    name: system:image-builder
  subjects:
  - kind: Group
    apiGroup: rbac.authorization.k8s.io
    name: system:unauthenticated
' | oc --context="${context}" apply -f -
    done

    # Use the primary cluster's registry as the main hub
    local primary_context="${cluster_contexts[0]}"
    URL=$(oc --context="${primary_context}" get route default-route -n openshift-image-registry --template='{{ .spec.host }}')
    export HUB="${URL}/${NAMESPACE}"
    echo "Primary cluster internal registry URL: ${HUB}"

    return
  fi

  # Single-cluster mode: original logic
  # Check if the registry pods are running
  oc get pods -n openshift-image-registry --no-headers | grep -v "Running\|Completed" && echo "It looks like the OCP image registry is not deployed or Running. This tests scenario requires it. Aborting." && exit 1

  # Check if default route already exist
  if [ -z "$(oc get route default-route -n openshift-image-registry -o name)" ]; then
    echo "Route default-route does not exist, patching DefaultRoute to true on Image Registry."
    oc patch configs.imageregistry.operator.openshift.io/cluster --patch '{"spec":{"defaultRoute":true}}' --type=merge

    timeout --foreground -v -s SIGHUP -k ${TIMEOUT} ${TIMEOUT} bash --verbose -c \
      "until oc get route default-route -n openshift-image-registry &> /dev/null; do sleep 5; done && echo 'The 'default-route' has been created.'"
  fi

  # Get the registry route
  URL=$(oc get route default-route -n openshift-image-registry --template='{{ .spec.host }}')
  # Hub will be equal to the route url/project-name(NameSpace)
  export HUB="${URL}/${NAMESPACE}"
  echo "Internal registry URL: ${HUB}"

  # Create namespace from where the image are going to be pushed
  # This is needed because in the internal registry the images are stored in the namespace.
  # If the namespace already exist, it will not fail
  oc create namespace "${NAMESPACE}" || true

  deploy_rolebinding

  # Login to the internal registry when running on CRC (Only for local development)
  # Take into count that you will need to add before the registry URL as Insecure registry in "/etc/docker/daemon.json"
  if [[ ${URL} == *".apps-crc.testing"* ]]; then
    echo "Executing Docker login to the internal registry"
    if ! oc whoami -t | docker login -u "$(oc whoami)" --password-stdin "${URL}"; then
      echo "***** Error: Failed to log in to Docker registry."
      echo "***** Check the error and if is related to 'tls: failed to verify certificate' please add the registry URL as Insecure registry in '/etc/docker/daemon.json'"
      exit 1
    fi
  fi
}

function deploy_rolebinding() {
    # Adding roles to avoid the need to be authenticated to push images to the internal registry 
    # and pull them later in the any namespace
      echo '
kind: List
apiVersion: v1
items:
- apiVersion: rbac.authorization.k8s.io/v1
  kind: RoleBinding
  metadata:
    name: image-puller
    namespace: '"$NAMESPACE"'
  roleRef:
    apiGroup: rbac.authorization.k8s.io
    kind: ClusterRole
    name: system:image-puller
  subjects:
  - kind: Group
    apiGroup: rbac.authorization.k8s.io
    name: system:unauthenticated
  - kind: Group
    name: system:serviceaccounts
    apiGroup: rbac.authorization.k8s.io
- apiVersion: rbac.authorization.k8s.io/v1
  kind: RoleBinding
  metadata:
    name: image-pusher
    namespace: '"$NAMESPACE"'
  roleRef:
    apiGroup: rbac.authorization.k8s.io
    kind: ClusterRole
    name: system:image-builder
  subjects:
  - kind: Group
    apiGroup: rbac.authorization.k8s.io
    name: system:unauthenticated
' | oc apply -f -
}

# Set gcr.io as mirror to docker.io/istio to be able to get images in downstream tests.
function addGcrMirror(){
  oc apply -f - <<__EOF__
apiVersion: config.openshift.io/v1
kind: ImageDigestMirrorSet
metadata:
  name: docker-images-from-gcr
spec:
  imageDigestMirrors:
  - mirrors:
    - mirror.gcr.io
    source: docker.io
    mirrorSourcePolicy: NeverContactSource
---
apiVersion: config.openshift.io/v1
kind: ImageTagMirrorSet
metadata:
  name: docker-images-from-gcr
spec:
  imageTagMirrors:
  - mirrors:
    - mirror.gcr.io
    source: docker.io
    mirrorSourcePolicy: NeverContactSource
__EOF__
}

# Deploy MetalLB in the OCP cluster and configure IP address pool
function deployMetalLB() {
  # Check if MetalLB is already deployed
  echo "Checking if MetalLB is already deployed..."
  if oc get metallb metallb -n metallb-system && oc get ipaddresspool default -n metallb-system &> /dev/null; then
    echo "MetalLB is already deployed (MetalLB CR and IPAddressPool CR exist), skipping..."
    return 0
  else
    echo "MetalLB CR or IPAddressPool CR is not deployed, deploying..."
  fi

  # Create the metallb-system namespace
  echo '
apiVersion: v1
kind: Namespace
metadata:
  name: metallb-system' | oc apply -f -

  # Create Subscription for MetalLB
  echo '
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: metallb-operator-sub
  namespace: metallb-system
spec:
  channel: stable
  name: metallb-operator
  source: redhat-operators
  sourceNamespace: openshift-marketplace
  installPlanApproval: Automatic' | oc apply -f -

  # Check operator Phase is Succeeded
# shellcheck disable=SC2016
timeout --foreground -v -s SIGHUP -k ${TIMEOUT} ${TIMEOUT} bash -c 'until [ "$(oc get csv -n metallb-system | awk "/metallb-operator/ {print \$NF}")" == "Succeeded" ]; do sleep 5; done && echo "The MetalLB operator has been installed."'

  # Create MetalLB CR
  echo '
apiVersion: metallb.io/v1beta1
kind: MetalLB
metadata:
  name: metallb
  namespace: metallb-system' | oc apply -f -

  # Check MetalLB controller is running
timeout --foreground -v -s SIGHUP -k ${TIMEOUT} ${TIMEOUT} bash -c 'until oc get pods -n metallb-system --no-headers | grep controller | grep "Running"; do sleep 5; done && echo "The MetalLB controller is running."'

  # Get Nodes Internal IP by using: kubectl get nodes -l node-role.kubernetes.io/worker -o jsonpath='{.items[*].status.addresses[?(@.type=="InternalIP")].address}'
  NODE_IPS=$(oc get nodes -l node-role.kubernetes.io/worker -o jsonpath='{.items[*].status.addresses[?(@.type=="InternalIP")].address}' | tr ' ' ',')

  # Split The IPs by , to create the IP address pool
  IFS=',' read -r -a array <<< "${NODE_IPS}"

  # Iterate over the Create IPS to create address pool
  IP_POOL_YAML='
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: default
  namespace: metallb-system
spec:
  addresses:'

  # Iterate over the IPs to create address pool entries
  for ip in "${array[@]}"; do
    IP_POOL_YAML+=$'\n  - '"${ip}-${ip}"
  done
  echo "IP Pool YAML: ${IP_POOL_YAML}"

  # Apply the IP address pool
  echo "${IP_POOL_YAML}" | oc apply -f -

  # Check the IP address pool is created
  timeout --foreground -v -s SIGHUP -k ${TIMEOUT} ${TIMEOUT} bash -c 'until oc get IPAddressPool default -n metallb-system; do sleep 5; done && echo "The IP address pool has been created."'
  
  # IBM specific modifications
  if [ "${IBM}" == "true" ]; then
    # Create L2Advertisement CR
    echo '
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  name: default
  namespace: metallb-system' | oc apply -f -
  fi

  echo "MetalLB has been deployed and configured with the IP address pool."
}

#need to change env variables since make deploy of sail-operator uses them
function env_save(){
  INICIAL_NAMESPACE="$NAMESPACE"
  INICIAL_HUB="$HUB"
  INITIAL_TAG="$TAG"
}
function cleanup_sail_repo() {
    echo "Cleaning up..."
    cd .. 2>/dev/null || true
    rm -rf sail-operator
    export NAMESPACE="$INICIAL_NAMESPACE"
    export HUB="$INICIAL_HUB"
    export TAG="$INITIAL_TAG"
}

# Detect and set the sail-operator branch based on current Istio branch
function detect_sail_operator_branch() {
  # Allow explicit override via environment variable
  if [ -n "${SAIL_OPERATOR_BRANCH:-}" ]; then
    echo "Using explicitly set SAIL_OPERATOR_BRANCH: ${SAIL_OPERATOR_BRANCH}"
    return 0
  fi

  # Detect current Istio branch
  local current_branch
  current_branch="$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "main")"

  # Map Istio branch to sail-operator branch
  if [[ "${current_branch}" == "main" ]] || [[ "${current_branch}" == "master" ]]; then
    SAIL_OPERATOR_BRANCH="main"
  elif [[ "${current_branch}" =~ ^release-[0-9]+\.[0-9]+$ ]]; then
    # Direct mapping: release-X.Y -> release-X.Y
    SAIL_OPERATOR_BRANCH="${current_branch}"
  else
    # Fallback to main for unrecognized patterns (e.g., feature branches)
    echo "Warning: Unrecognized branch pattern '${current_branch}', defaulting to sail-operator main branch"
    SAIL_OPERATOR_BRANCH="main"
  fi

  export SAIL_OPERATOR_BRANCH
  echo "Detected sail-operator branch: ${SAIL_OPERATOR_BRANCH} (from Istio branch: ${current_branch})"
}

function deploy_operator(){
  # Detect appropriate sail-operator branch before cloning
  detect_sail_operator_branch

  # Save and unset env vars so sail-operator's make deploy uses its own defaults
  env_save
  unset HUB
  unset TAG
  unset NAMESPACE

  git clone --depth 1 --branch "${SAIL_OPERATOR_BRANCH}" $SAIL_REPO_URL || { echo "Failed to clone sail-operator repo on branch ${SAIL_OPERATOR_BRANCH}"; exit 1; }
  cd sail-operator
  make deploy || { echo "sail-operator make deploy failed"; cleanup_sail_repo ; exit 1; }
  oc -n sail-operator wait --for=condition=Available deployment/sail-operator --timeout=240s || { echo "Failed to start sail-operator"; exit 1; }

  # Restore original env vars for subsequent Istio operations
  cleanup_sail_repo
  echo "Sail operator deployed from branch: ${SAIL_OPERATOR_BRANCH}"

}

# Multicluster setup functions

function load_cluster_topology() {
  local topology_file="$1"

  if [ ! -f "${topology_file}" ]; then
    echo "Error: Topology file ${topology_file} not found"
    exit 1
  fi

  echo "Loading cluster topology from ${topology_file}"

  # Check if jq is installed
  if ! command -v jq &> /dev/null; then
    echo "Error: jq is required for topology parsing. Please install jq."
    exit 1
  fi

  # Extract cluster names and networks from topology JSON
  mapfile -t CLUSTER_NAMES < <(jq -r '.[] | .clusterName' "${topology_file}")
  mapfile -t CLUSTER_NETWORKS < <(jq -r '.[] | .network' "${topology_file}")

  export CLUSTER_NAMES
  export CLUSTER_NETWORKS
  export NUM_CLUSTERS="${#CLUSTER_NAMES[@]}"

  echo "Loaded ${NUM_CLUSTERS} clusters from topology:"
  for i in "${!CLUSTER_NAMES[@]}"; do
    echo "  - ${CLUSTER_NAMES[i]} (network: ${CLUSTER_NETWORKS[i]})"
  done
}

function validate_ocp_multicluster_kubeconfigs() {
  echo "Validating OCP multicluster kubeconfigs..."

  # Check if KUBECONFIG environment variable is set
  if [ -z "${KUBECONFIG:-}" ]; then
    echo "Error: KUBECONFIG environment variable is not set"
    echo "Expected format: KUBECONFIG=/path/to/primary.kubeconfig:/path/to/remote.kubeconfig"
    exit 1
  fi

  echo "KUBECONFIG: ${KUBECONFIG}"

  # Get all available contexts from merged kubeconfig
  mapfile -t AVAILABLE_CONTEXTS < <(kubectl config get-contexts -o name 2>/dev/null || true)

  if [ ${#AVAILABLE_CONTEXTS[@]} -eq 0 ]; then
    echo "Error: No contexts found in KUBECONFIG"
    exit 1
  fi

  echo "Available contexts:"
  printf '  - %s\n' "${AVAILABLE_CONTEXTS[@]}"

  # Validate each cluster in topology has a matching context
  local missing_contexts=()
  for cluster_name in "${CLUSTER_NAMES[@]}"; do
    local found=false
    for context in "${AVAILABLE_CONTEXTS[@]}"; do
      # Match exact cluster name or context containing cluster name
      if [[ "${context}" == "${cluster_name}" ]] || [[ "${context}" == *"/${cluster_name}/"* ]] || [[ "${context}" == *"-${cluster_name}-"* ]]; then
        found=true
        echo "Found context for cluster ${cluster_name}: ${context}"
        break
      fi
    done

    if [ "${found}" = false ]; then
      missing_contexts+=("${cluster_name}")
    fi
  done

  if [ ${#missing_contexts[@]} -gt 0 ]; then
    echo "Error: Missing contexts for the following clusters:"
    printf '  - %s\n' "${missing_contexts[@]}"
    echo ""
    echo "Available contexts:"
    printf '  - %s\n' "${AVAILABLE_CONTEXTS[@]}"
    exit 1
  fi

  # Validate cluster access by running oc cluster-info for each cluster
  echo "Validating cluster access..."
  for cluster_name in "${CLUSTER_NAMES[@]}"; do
    # Find the matching context
    local cluster_context=""
    for context in "${AVAILABLE_CONTEXTS[@]}"; do
      if [[ "${context}" == "${cluster_name}" ]] || [[ "${context}" == *"/${cluster_name}/"* ]] || [[ "${context}" == *"-${cluster_name}-"* ]]; then
        cluster_context="${context}"
        break
      fi
    done

    if ! oc --context="${cluster_context}" cluster-info &> /dev/null; then
      echo "Error: Cannot access cluster ${cluster_name} using context ${cluster_context}"
      exit 1
    fi

    echo "Successfully validated access to cluster ${cluster_name}"
  done

  echo "All multicluster kubeconfigs validated successfully"
}

function validate_ocp_cluster_connectivity() {
  echo "Validating cross-cluster network connectivity..."

  # Create a test namespace for connectivity checks
  local test_namespace="istio-mc-connectivity-test"

  # Get available contexts for each cluster
  local -a cluster_contexts
  for cluster_name in "${CLUSTER_NAMES[@]}"; do
    local cluster_context=""
    mapfile -t available_contexts < <(kubectl config get-contexts -o name 2>/dev/null || true)
    for context in "${available_contexts[@]}"; do
      if [[ "${context}" == "${cluster_name}" ]] || [[ "${context}" == *"/${cluster_name}/"* ]] || [[ "${context}" == *"-${cluster_name}-"* ]]; then
        cluster_context="${context}"
        break
      fi
    done
    cluster_contexts+=("${cluster_context}")
  done

  # Deploy test pods in each cluster
  echo "Deploying test pods in each cluster..."
  local -a pod_ips
  for i in "${!CLUSTER_NAMES[@]}"; do
    local cluster_name="${CLUSTER_NAMES[i]}"
    local context="${cluster_contexts[i]}"

    # Create test namespace
    oc --context="${context}" create namespace "${test_namespace}" 2>/dev/null || true

    # Deploy a simple test pod
    oc --context="${context}" run -n "${test_namespace}" connectivity-test-pod-"${cluster_name}" \
      --image=registry.access.redhat.com/ubi9/ubi-minimal:latest \
      --command -- sleep 3600 2>/dev/null || true

    # Wait for pod to be running
    if ! oc --context="${context}" wait --for=condition=Ready -n "${test_namespace}" \
      pod/connectivity-test-pod-"${cluster_name}" --timeout=60s 2>/dev/null; then
      echo "Warning: Test pod for cluster ${cluster_name} not ready within timeout, skipping detailed connectivity check"
      continue
    fi

    # Get pod IP
    local pod_ip
    pod_ip=$(oc --context="${context}" get pod -n "${test_namespace}" connectivity-test-pod-"${cluster_name}" \
      -o jsonpath='{.status.podIP}' 2>/dev/null || echo "")
    pod_ips+=("${pod_ip}")
    echo "Test pod in cluster ${cluster_name}: IP=${pod_ip}"
  done

  # Test connectivity between cluster pairs on the same network
  echo "Testing connectivity between clusters on the same network..."
  local connectivity_success=true
  for i in "${!CLUSTER_NAMES[@]}"; do
    for j in "${!CLUSTER_NAMES[@]}"; do
      if [ "$i" -lt "$j" ]; then
        # Check if clusters are on the same network
        if [ "${CLUSTER_NETWORKS[i]}" == "${CLUSTER_NETWORKS[j]}" ]; then
          local from_cluster="${CLUSTER_NAMES[i]}"
          local to_cluster="${CLUSTER_NAMES[j]}"
          local to_ip="${pod_ips[j]}"
          local from_context="${cluster_contexts[i]}"

          if [ -n "${to_ip}" ]; then
            echo "Testing connectivity from ${from_cluster} to ${to_cluster} (${to_ip})..."
            if oc --context="${from_context}" exec -n "${test_namespace}" \
              connectivity-test-pod-"${from_cluster}" -- \
              timeout 5 sh -c "ping -c 1 -W 1 ${to_ip}" &>/dev/null; then
              echo "  ✓ Connectivity OK"
            else
              echo "  ✗ Connectivity FAILED"
              echo "Warning: Clusters ${from_cluster} and ${to_cluster} are on same network (${CLUSTER_NETWORKS[i]}) but cannot communicate"
              connectivity_success=false
            fi
          fi
        fi
      fi
    done
  done

  # Cleanup test resources
  echo "Cleaning up connectivity test resources..."
  for context in "${cluster_contexts[@]}"; do
    oc --context="${context}" delete namespace "${test_namespace}" --ignore-not-found=true &>/dev/null || true
  done

  if [ "${connectivity_success}" = false ]; then
    echo "Error: Connectivity validation failed. Please check network policies and cluster network configuration."
    echo "For same-network clusters, ensure:"
    echo "  - Network policies allow pod-to-pod communication"
    echo "  - Cluster CNI configurations are compatible"
    echo "  - Firewall rules allow cross-cluster traffic"
    exit 1
  fi

  echo "Cross-cluster connectivity validation completed successfully"
}

function deploy_east_west_gateways() {
  echo "Deploying East-West gateways for cross-network scenarios..."

  # Get available contexts for each cluster
  local -a cluster_contexts
  for cluster_name in "${CLUSTER_NAMES[@]}"; do
    local cluster_context=""
    mapfile -t available_contexts < <(kubectl config get-contexts -o name 2>/dev/null || true)
    for context in "${available_contexts[@]}"; do
      if [[ "${context}" == "${cluster_name}" ]] || [[ "${context}" == *"/${cluster_name}/"* ]] || [[ "${context}" == *"-${cluster_name}-"* ]]; then
        cluster_context="${context}"
        break
      fi
    done
    cluster_contexts+=("${cluster_context}")
  done

  # Deploy gateway in each cluster
  for i in "${!CLUSTER_NAMES[@]}"; do
    local cluster_name="${CLUSTER_NAMES[i]}"
    local network="${CLUSTER_NETWORKS[i]}"
    local context="${cluster_contexts[i]}"

    echo "Deploying East-West gateway in cluster ${cluster_name} (network: ${network})..."

    # Create eastwest gateway using Istio samples configuration
    # This is a simplified version - the actual implementation will be handled by the test framework
    # when --istio.test.ambient.multinetwork flag is passed

    # Check if istio-system namespace exists
    if ! oc --context="${context}" get namespace "${NAMESPACE}" &>/dev/null; then
      echo "Note: Namespace ${NAMESPACE} does not exist in cluster ${cluster_name}"
      echo "Gateway will be deployed by test framework during test execution"
    else
      echo "Gateway deployment will be handled by test framework with --istio.test.ambient.multinetwork flag"
    fi
  done

  echo "East-West gateway deployment configuration completed"
  echo "Note: Actual gateway deployment is handled by the test framework"
}

function setup_ocp_multicluster_topology() {
  echo "Setting up OCP multicluster topology configuration..."

  local topology_file="$1"
  local runtime_topology_file="${ARTIFACTS_DIR}/topology-config.json"

  # Read the original topology JSON
  local topology_json
  topology_json=$(cat "${topology_file}")

  # Get available contexts
  mapfile -t available_contexts < <(kubectl config get-contexts -o name 2>/dev/null || true)

  # Extract kubeconfig paths from KUBECONFIG environment variable
  IFS=':' read -r -a kubeconfig_paths <<< "${KUBECONFIG}"

  # For each cluster, inject the kubeconfig path
  for i in "${!CLUSTER_NAMES[@]}"; do
    local cluster_name="${CLUSTER_NAMES[i]}"

    # Find matching context
    local cluster_context=""
    for context in "${available_contexts[@]}"; do
      if [[ "${context}" == "${cluster_name}" ]] || [[ "${context}" == *"/${cluster_name}/"* ]] || [[ "${context}" == *"-${cluster_name}-"* ]]; then
        cluster_context="${context}"
        break
      fi
    done

    # Find the kubeconfig file that contains this context
    local kubeconfig_path=""
    for kconfig in "${kubeconfig_paths[@]}"; do
      if kubectl --kubeconfig="${kconfig}" config get-contexts "${cluster_context}" &>/dev/null; then
        kubeconfig_path="${kconfig}"
        break
      fi
    done

    if [ -z "${kubeconfig_path}" ]; then
      echo "Warning: Could not find kubeconfig file for cluster ${cluster_name}, using merged KUBECONFIG"
      # Use the first kubeconfig in the list as fallback
      kubeconfig_path="${kubeconfig_paths[0]}"
    fi

    echo "Cluster ${cluster_name}: kubeconfig=${kubeconfig_path}"

    # Source the set_topology_value function from lib.sh
    # shellcheck source=prow/lib.sh
    if [ -f "${ROOT}/prow/lib.sh" ]; then
      source "${ROOT}/prow/lib.sh"
    fi

    # Inject kubeconfig path into topology JSON
    topology_json=$(set_topology_value "${topology_json}" "${cluster_name}" "meta.kubeconfig" "${kubeconfig_path}")
  done

  # Write the runtime topology configuration
  echo "${topology_json}" > "${runtime_topology_file}"

  echo "Runtime topology configuration written to ${runtime_topology_file}"
  echo "Topology contents:"
  jq '.' "${runtime_topology_file}"

  export INTEGRATION_TEST_TOPOLOGY_FILE="${runtime_topology_file}"
}

function generate_dynamic_topology() {
  local num_clusters="$1"
  local topology_type="${2:-MULTICLUSTER}"
  local output_file="${ARTIFACTS_DIR}/dynamic-topology.json"

  echo "Generating dynamic topology for ${num_clusters} clusters (type: ${topology_type})"

  # Get available contexts from KUBECONFIG
  mapfile -t available_contexts < <(kubectl config get-contexts -o name 2>/dev/null || true)

  if [[ ${#available_contexts[@]} -lt ${num_clusters} ]]; then
    echo "Error: Requested ${num_clusters} clusters, but only ${#available_contexts[@]} contexts available in KUBECONFIG"
    exit 1
  fi

  # Well-known cluster names for Istio test framework
  local -a cluster_names=("primary" "remote" "cross-network-primary")
  local -a networks=("network-1" "network-1" "network-2")

  # For ambient multicluster, use cross-network configuration
  if [[ "${topology_type}" == "AMBIENT_MULTICLUSTER" ]]; then
    networks=("network-1" "network-2" "network-3")
  fi

  # Build topology JSON
  local topology_json="["

  for i in $(seq 0 $((num_clusters - 1))); do
    local cluster_name="${cluster_names[i]}"
    local network="${networks[i]}"

    # Add cluster entry
    if [[ $i -gt 0 ]]; then
      topology_json+=","
    fi

    topology_json+='
  {
    "kind": "Kubernetes",
    "clusterName": "'${cluster_name}'",
    "network": "'${network}'"'

    # Add primaryClusterName for remote clusters
    if [[ "${cluster_name}" == "remote" ]]; then
      topology_json+=',
    "primaryClusterName": "primary"'
    fi

    topology_json+='
  }'
  done

  topology_json+='
]'

  # Write topology file
  echo "${topology_json}" > "${output_file}"

  echo "Dynamic topology generated at ${output_file}:"
  jq '.' "${output_file}"

  echo "${output_file}"
}
