//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.1-0.20241121165744-79df5c4772f2
// source: workloadapi/security/authorization.proto

package security

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	emptypb1 "github.com/planetscale/vtprotobuf/types/known/emptypb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (this *Authorization) EqualVT(that *Authorization) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if this.Namespace != that.Namespace {
		return false
	}
	if this.Scope != that.Scope {
		return false
	}
	if this.Action != that.Action {
		return false
	}
	if len(this.Groups) != len(that.Groups) {
		return false
	}
	for i, vx := range this.Groups {
		vy := that.Groups[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Group{}
			}
			if q == nil {
				q = &Group{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Authorization) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Authorization)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Group) EqualVT(that *Group) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Rules) != len(that.Rules) {
		return false
	}
	for i, vx := range this.Rules {
		vy := that.Rules[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Rules{}
			}
			if q == nil {
				q = &Rules{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Group) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Group)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Rules) EqualVT(that *Rules) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Matches) != len(that.Matches) {
		return false
	}
	for i, vx := range this.Matches {
		vy := that.Matches[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Match{}
			}
			if q == nil {
				q = &Match{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Rules) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Rules)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Match) EqualVT(that *Match) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Namespaces) != len(that.Namespaces) {
		return false
	}
	for i, vx := range this.Namespaces {
		vy := that.Namespaces[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &StringMatch{}
			}
			if q == nil {
				q = &StringMatch{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.NotNamespaces) != len(that.NotNamespaces) {
		return false
	}
	for i, vx := range this.NotNamespaces {
		vy := that.NotNamespaces[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &StringMatch{}
			}
			if q == nil {
				q = &StringMatch{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.Principals) != len(that.Principals) {
		return false
	}
	for i, vx := range this.Principals {
		vy := that.Principals[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &StringMatch{}
			}
			if q == nil {
				q = &StringMatch{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.NotPrincipals) != len(that.NotPrincipals) {
		return false
	}
	for i, vx := range this.NotPrincipals {
		vy := that.NotPrincipals[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &StringMatch{}
			}
			if q == nil {
				q = &StringMatch{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.SourceIps) != len(that.SourceIps) {
		return false
	}
	for i, vx := range this.SourceIps {
		vy := that.SourceIps[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Address{}
			}
			if q == nil {
				q = &Address{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.NotSourceIps) != len(that.NotSourceIps) {
		return false
	}
	for i, vx := range this.NotSourceIps {
		vy := that.NotSourceIps[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Address{}
			}
			if q == nil {
				q = &Address{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.DestinationIps) != len(that.DestinationIps) {
		return false
	}
	for i, vx := range this.DestinationIps {
		vy := that.DestinationIps[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Address{}
			}
			if q == nil {
				q = &Address{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.NotDestinationIps) != len(that.NotDestinationIps) {
		return false
	}
	for i, vx := range this.NotDestinationIps {
		vy := that.NotDestinationIps[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Address{}
			}
			if q == nil {
				q = &Address{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.DestinationPorts) != len(that.DestinationPorts) {
		return false
	}
	for i, vx := range this.DestinationPorts {
		vy := that.DestinationPorts[i]
		if vx != vy {
			return false
		}
	}
	if len(this.NotDestinationPorts) != len(that.NotDestinationPorts) {
		return false
	}
	for i, vx := range this.NotDestinationPorts {
		vy := that.NotDestinationPorts[i]
		if vx != vy {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Match) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Match)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Address) EqualVT(that *Address) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if string(this.Address) != string(that.Address) {
		return false
	}
	if this.Length != that.Length {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Address) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Address)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *StringMatch) EqualVT(that *StringMatch) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.MatchType == nil && that.MatchType != nil {
		return false
	} else if this.MatchType != nil {
		if that.MatchType == nil {
			return false
		}
		if !this.MatchType.(interface {
			EqualVT(isStringMatch_MatchType) bool
		}).EqualVT(that.MatchType) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *StringMatch) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*StringMatch)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *StringMatch_Exact) EqualVT(thatIface isStringMatch_MatchType) bool {
	that, ok := thatIface.(*StringMatch_Exact)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Exact != that.Exact {
		return false
	}
	return true
}

func (this *StringMatch_Prefix) EqualVT(thatIface isStringMatch_MatchType) bool {
	that, ok := thatIface.(*StringMatch_Prefix)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Prefix != that.Prefix {
		return false
	}
	return true
}

func (this *StringMatch_Suffix) EqualVT(thatIface isStringMatch_MatchType) bool {
	that, ok := thatIface.(*StringMatch_Suffix)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Suffix != that.Suffix {
		return false
	}
	return true
}

func (this *StringMatch_Presence) EqualVT(thatIface isStringMatch_MatchType) bool {
	that, ok := thatIface.(*StringMatch_Presence)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Presence, that.Presence; p != q {
		if p == nil {
			p = &emptypb.Empty{}
		}
		if q == nil {
			q = &emptypb.Empty{}
		}
		if !(*emptypb1.Empty)(p).EqualVT((*emptypb1.Empty)(q)) {
			return false
		}
	}
	return true
}

func (m *Authorization) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Authorization) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Authorization) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Groups) > 0 {
		for iNdEx := len(m.Groups) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Groups[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Action != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x20
	}
	if m.Scope != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Scope))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Group) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Group) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Group) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Rules) > 0 {
		for iNdEx := len(m.Rules) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Rules[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Rules) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rules) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Rules) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Matches) > 0 {
		for iNdEx := len(m.Matches) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Matches[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *Match) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Match) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Match) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.NotDestinationPorts) > 0 {
		var pksize2 int
		for _, num := range m.NotDestinationPorts {
			pksize2 += protohelpers.SizeOfVarint(uint64(num))
		}
		i -= pksize2
		j1 := i
		for _, num := range m.NotDestinationPorts {
			for num >= 1<<7 {
				dAtA[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA[j1] = uint8(num)
			j1++
		}
		i = protohelpers.EncodeVarint(dAtA, i, uint64(pksize2))
		i--
		dAtA[i] = 0x52
	}
	if len(m.DestinationPorts) > 0 {
		var pksize4 int
		for _, num := range m.DestinationPorts {
			pksize4 += protohelpers.SizeOfVarint(uint64(num))
		}
		i -= pksize4
		j3 := i
		for _, num := range m.DestinationPorts {
			for num >= 1<<7 {
				dAtA[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA[j3] = uint8(num)
			j3++
		}
		i = protohelpers.EncodeVarint(dAtA, i, uint64(pksize4))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.NotDestinationIps) > 0 {
		for iNdEx := len(m.NotDestinationIps) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.NotDestinationIps[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.DestinationIps) > 0 {
		for iNdEx := len(m.DestinationIps) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DestinationIps[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.NotSourceIps) > 0 {
		for iNdEx := len(m.NotSourceIps) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.NotSourceIps[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.SourceIps) > 0 {
		for iNdEx := len(m.SourceIps) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.SourceIps[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.NotPrincipals) > 0 {
		for iNdEx := len(m.NotPrincipals) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.NotPrincipals[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Principals) > 0 {
		for iNdEx := len(m.Principals) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Principals[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.NotNamespaces) > 0 {
		for iNdEx := len(m.NotNamespaces) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.NotNamespaces[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Namespaces) > 0 {
		for iNdEx := len(m.Namespaces) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Namespaces[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Address) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Address) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Address) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Length != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Length))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StringMatch) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringMatch) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *StringMatch) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if msg, ok := m.MatchType.(*StringMatch_Presence); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.MatchType.(*StringMatch_Suffix); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.MatchType.(*StringMatch_Prefix); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.MatchType.(*StringMatch_Exact); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *StringMatch_Exact) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *StringMatch_Exact) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Exact)
	copy(dAtA[i:], m.Exact)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Exact)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *StringMatch_Prefix) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *StringMatch_Prefix) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Prefix)
	copy(dAtA[i:], m.Prefix)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Prefix)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *StringMatch_Suffix) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *StringMatch_Suffix) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Suffix)
	copy(dAtA[i:], m.Suffix)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Suffix)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *StringMatch_Presence) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *StringMatch_Presence) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Presence != nil {
		size, err := (*emptypb1.Empty)(m.Presence).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Authorization) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Scope != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Scope))
	}
	if m.Action != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Action))
	}
	if len(m.Groups) > 0 {
		for _, e := range m.Groups {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *Group) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *Rules) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Matches) > 0 {
		for _, e := range m.Matches {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *Match) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Namespaces) > 0 {
		for _, e := range m.Namespaces {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.NotNamespaces) > 0 {
		for _, e := range m.NotNamespaces {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.Principals) > 0 {
		for _, e := range m.Principals {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.NotPrincipals) > 0 {
		for _, e := range m.NotPrincipals {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.SourceIps) > 0 {
		for _, e := range m.SourceIps {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.NotSourceIps) > 0 {
		for _, e := range m.NotSourceIps {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.DestinationIps) > 0 {
		for _, e := range m.DestinationIps {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.NotDestinationIps) > 0 {
		for _, e := range m.NotDestinationIps {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.DestinationPorts) > 0 {
		l = 0
		for _, e := range m.DestinationPorts {
			l += protohelpers.SizeOfVarint(uint64(e))
		}
		n += 1 + protohelpers.SizeOfVarint(uint64(l)) + l
	}
	if len(m.NotDestinationPorts) > 0 {
		l = 0
		for _, e := range m.NotDestinationPorts {
			l += protohelpers.SizeOfVarint(uint64(e))
		}
		n += 1 + protohelpers.SizeOfVarint(uint64(l)) + l
	}
	n += len(m.unknownFields)
	return n
}

func (m *Address) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Length != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Length))
	}
	n += len(m.unknownFields)
	return n
}

func (m *StringMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.MatchType.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *StringMatch_Exact) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Exact)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *StringMatch_Prefix) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Prefix)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *StringMatch_Suffix) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Suffix)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *StringMatch_Presence) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Presence != nil {
		l = (*emptypb1.Empty)(m.Presence).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 2
	}
	return n
}
