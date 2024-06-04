// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/authority/policies.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// PolicyType defines the type of policy
type PolicyType int32

const (
	PolicyType_groupEmergency   PolicyType = 0
	PolicyType_groupOperational PolicyType = 1
	// non-sensitive protocol parameters
	PolicyType_groupAdmin PolicyType = 2
	// protocol parameters or moving funds
	PolicyType_groupEmpty PolicyType = 3
)

var PolicyType_name = map[int32]string{
	0: "groupEmergency",
	1: "groupOperational",
	2: "groupAdmin",
	3: "groupEmpty",
}

var PolicyType_value = map[string]int32{
	"groupEmergency":   0,
	"groupOperational": 1,
	"groupAdmin":       2,
	"groupEmpty":       3,
}

func (x PolicyType) String() string {
	return proto.EnumName(PolicyType_name, int32(x))
}

func (PolicyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_afa9e3e7b996ef74, []int{0}
}

type Policy struct {
	PolicyType PolicyType `protobuf:"varint,1,opt,name=policy_type,json=policyType,proto3,enum=zetachain.zetacore.authority.PolicyType" json:"policy_type,omitempty"`
	Address    string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa9e3e7b996ef74, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return m.Size()
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPolicyType() PolicyType {
	if m != nil {
		return m.PolicyType
	}
	return PolicyType_groupEmergency
}

func (m *Policy) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// Policy contains info about authority policies
type Policies struct {
	Items []*Policy `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (m *Policies) Reset()         { *m = Policies{} }
func (m *Policies) String() string { return proto.CompactTextString(m) }
func (*Policies) ProtoMessage()    {}
func (*Policies) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa9e3e7b996ef74, []int{1}
}
func (m *Policies) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Policies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Policies.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Policies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policies.Merge(m, src)
}
func (m *Policies) XXX_Size() int {
	return m.Size()
}
func (m *Policies) XXX_DiscardUnknown() {
	xxx_messageInfo_Policies.DiscardUnknown(m)
}

var xxx_messageInfo_Policies proto.InternalMessageInfo

func (m *Policies) GetItems() []*Policy {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterEnum("zetachain.zetacore.authority.PolicyType", PolicyType_name, PolicyType_value)
	proto.RegisterType((*Policy)(nil), "zetachain.zetacore.authority.Policy")
	proto.RegisterType((*Policies)(nil), "zetachain.zetacore.authority.Policies")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/authority/policies.proto", fileDescriptor_afa9e3e7b996ef74)
}

var fileDescriptor_afa9e3e7b996ef74 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0xdc, 0xb4, 0x5a, 0xf5, 0x15, 0xca, 0x12, 0x7a, 0x58, 0x8a, 0x84, 0x52, 0x3c, 0x2c, 0x8a,
	0x59, 0xa9, 0x37, 0x6f, 0x0a, 0x15, 0xf4, 0x62, 0x29, 0x9e, 0x44, 0x90, 0x74, 0x1b, 0xb6, 0x81,
	0x66, 0x13, 0xb2, 0x29, 0x18, 0xbf, 0xc2, 0x8f, 0xf0, 0xe0, 0xa7, 0x78, 0xec, 0xd1, 0xa3, 0xb4,
	0x3f, 0x22, 0xcd, 0xda, 0x55, 0x3c, 0x88, 0xb7, 0x79, 0x8f, 0x99, 0x79, 0x8f, 0x19, 0x38, 0x7a,
	0xe2, 0x96, 0xa5, 0x53, 0x26, 0xf2, 0xc4, 0x23, 0x65, 0x78, 0xc2, 0xe6, 0x76, 0xaa, 0x8c, 0xb0,
	0x2e, 0xd1, 0x6a, 0x26, 0x52, 0xc1, 0x0b, 0xaa, 0x8d, 0xb2, 0x0a, 0xef, 0x57, 0x64, 0xba, 0x21,
	0xd3, 0x8a, 0xdc, 0x69, 0x67, 0x2a, 0x53, 0x9e, 0x98, 0xac, 0x51, 0xa9, 0xe9, 0x49, 0x68, 0x0c,
	0xd7, 0x2e, 0x0e, 0x5f, 0x41, 0xd3, 0xfb, 0xb9, 0x07, 0xeb, 0x34, 0x8f, 0x50, 0x17, 0xc5, 0xad,
	0x7e, 0x4c, 0xff, 0xf2, 0xa4, 0xa5, 0xf4, 0xd6, 0x69, 0x3e, 0x02, 0x5d, 0x61, 0x1c, 0xc1, 0x0e,
	0x9b, 0x4c, 0x0c, 0x2f, 0x8a, 0xa8, 0xd6, 0x45, 0xf1, 0xde, 0x68, 0x33, 0xf6, 0x2e, 0x61, 0x77,
	0xf8, 0xf5, 0x34, 0x3e, 0x83, 0x6d, 0x61, 0xb9, 0x2c, 0x22, 0xd4, 0xad, 0xc7, 0xcd, 0xfe, 0xc1,
	0x7f, 0x4e, 0x8d, 0x4a, 0xc9, 0xe1, 0x3d, 0xc0, 0xf7, 0x6d, 0x8c, 0xa1, 0x95, 0x19, 0x35, 0xd7,
	0x03, 0xc9, 0x4d, 0xc6, 0xf3, 0xd4, 0x85, 0x01, 0x6e, 0x43, 0xe8, 0x77, 0x37, 0x9a, 0x1b, 0x66,
	0x85, 0xca, 0xd9, 0x2c, 0x44, 0xb8, 0x05, 0xe0, 0xb7, 0xe7, 0x13, 0x29, 0xf2, 0xb0, 0x56, 0xcd,
	0x03, 0xa9, 0xad, 0x0b, 0xeb, 0x9d, 0xad, 0xd7, 0x17, 0x82, 0x2e, 0xae, 0xdf, 0x96, 0x04, 0x2d,
	0x96, 0x04, 0x7d, 0x2c, 0x09, 0x7a, 0x5e, 0x91, 0x60, 0xb1, 0x22, 0xc1, 0xfb, 0x8a, 0x04, 0x77,
	0x27, 0x99, 0xb0, 0xd3, 0xf9, 0x98, 0xa6, 0x4a, 0xfa, 0x42, 0x8e, 0x7f, 0x75, 0xf3, 0xf8, 0xa3,
	0x9d, 0x75, 0x8c, 0xc5, 0xb8, 0xe1, 0x73, 0x3e, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x37, 0x5d,
	0x49, 0x63, 0xca, 0x01, 0x00, 0x00,
}

func (m *Policy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Policy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Policy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPolicies(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.PolicyType != 0 {
		i = encodeVarintPolicies(dAtA, i, uint64(m.PolicyType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Policies) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Policies) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Policies) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPolicies(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPolicies(dAtA []byte, offset int, v uint64) int {
	offset -= sovPolicies(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Policy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PolicyType != 0 {
		n += 1 + sovPolicies(uint64(m.PolicyType))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPolicies(uint64(l))
	}
	return n
}

func (m *Policies) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovPolicies(uint64(l))
		}
	}
	return n
}

func sovPolicies(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPolicies(x uint64) (n int) {
	return sovPolicies(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Policy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicies
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Policy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Policy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PolicyType", wireType)
			}
			m.PolicyType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicies
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PolicyType |= PolicyType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicies
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPolicies
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicies
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicies(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicies
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Policies) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicies
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Policies: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Policies: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicies
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPolicies
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolicies
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Policy{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicies(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicies
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPolicies(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPolicies
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPolicies
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPolicies
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPolicies
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPolicies
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPolicies
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPolicies        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPolicies          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPolicies = fmt.Errorf("proto: unexpected end of group")
)
