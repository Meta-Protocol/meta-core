// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/observer/confirmation.proto

package types

import (
	fmt "fmt"
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

type Confirmation struct {
	SafeInboundCount  uint64 `protobuf:"varint,1,opt,name=safe_inbound_count,json=safeInboundCount,proto3" json:"safe_inbound_count,omitempty"`
	FastInboundCount  uint64 `protobuf:"varint,2,opt,name=fast_inbound_count,json=fastInboundCount,proto3" json:"fast_inbound_count,omitempty"`
	SafeOutboundCount uint64 `protobuf:"varint,3,opt,name=safe_outbound_count,json=safeOutboundCount,proto3" json:"safe_outbound_count,omitempty"`
	FastOutboundCount uint64 `protobuf:"varint,4,opt,name=fast_outbound_count,json=fastOutboundCount,proto3" json:"fast_outbound_count,omitempty"`
}

func (m *Confirmation) Reset()         { *m = Confirmation{} }
func (m *Confirmation) String() string { return proto.CompactTextString(m) }
func (*Confirmation) ProtoMessage()    {}
func (*Confirmation) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b45655bc4002d5, []int{0}
}
func (m *Confirmation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Confirmation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Confirmation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Confirmation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Confirmation.Merge(m, src)
}
func (m *Confirmation) XXX_Size() int {
	return m.Size()
}
func (m *Confirmation) XXX_DiscardUnknown() {
	xxx_messageInfo_Confirmation.DiscardUnknown(m)
}

var xxx_messageInfo_Confirmation proto.InternalMessageInfo

func (m *Confirmation) GetSafeInboundCount() uint64 {
	if m != nil {
		return m.SafeInboundCount
	}
	return 0
}

func (m *Confirmation) GetFastInboundCount() uint64 {
	if m != nil {
		return m.FastInboundCount
	}
	return 0
}

func (m *Confirmation) GetSafeOutboundCount() uint64 {
	if m != nil {
		return m.SafeOutboundCount
	}
	return 0
}

func (m *Confirmation) GetFastOutboundCount() uint64 {
	if m != nil {
		return m.FastOutboundCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Confirmation)(nil), "zetachain.zetacore.observer.Confirmation")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/observer/confirmation.proto", fileDescriptor_26b45655bc4002d5)
}

var fileDescriptor_26b45655bc4002d5 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xab, 0x4a, 0x2d, 0x49,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0xb3, 0xf2, 0x8b, 0x52, 0xf5, 0xf3, 0x93, 0x8a, 0x53,
	0x8b, 0xca, 0x52, 0x8b, 0xf4, 0x93, 0xf3, 0xf3, 0xd2, 0x32, 0x8b, 0x72, 0x13, 0x4b, 0x32, 0xf3,
	0xf3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xa4, 0xe1, 0xea, 0xf5, 0x60, 0xea, 0xf5, 0x60,
	0xea, 0x95, 0x4e, 0x31, 0x72, 0xf1, 0x38, 0x23, 0xe9, 0x11, 0xd2, 0xe1, 0x12, 0x2a, 0x4e, 0x4c,
	0x4b, 0x8d, 0xcf, 0xcc, 0x4b, 0xca, 0x2f, 0xcd, 0x4b, 0x89, 0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x12, 0x00, 0xc9, 0x78, 0x42, 0x24, 0x9c, 0x41, 0xe2, 0x20,
	0xd5, 0x69, 0x89, 0xc5, 0x25, 0x68, 0xaa, 0x99, 0x20, 0xaa, 0x41, 0x32, 0x28, 0xaa, 0xf5, 0xb8,
	0x84, 0xc1, 0x66, 0xe7, 0x97, 0x96, 0x20, 0x2b, 0x67, 0x06, 0x2b, 0x17, 0x04, 0x49, 0xf9, 0x43,
	0x65, 0xe0, 0xea, 0xc1, 0xa6, 0xa3, 0xa9, 0x67, 0x81, 0xa8, 0x07, 0x49, 0xa1, 0xa8, 0x77, 0x72,
	0x3d, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96,
	0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xed, 0xf4, 0xcc, 0x92, 0x8c,
	0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x70, 0xa0, 0xe9, 0x42, 0xc2, 0x2f, 0x2f, 0x3f, 0x25, 0x55,
	0xbf, 0x02, 0x11, 0x7a, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0x70, 0x33, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x34, 0xb5, 0xf0, 0xb5, 0x69, 0x01, 0x00, 0x00,
}

func (m *Confirmation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Confirmation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Confirmation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FastOutboundCount != 0 {
		i = encodeVarintConfirmation(dAtA, i, uint64(m.FastOutboundCount))
		i--
		dAtA[i] = 0x20
	}
	if m.SafeOutboundCount != 0 {
		i = encodeVarintConfirmation(dAtA, i, uint64(m.SafeOutboundCount))
		i--
		dAtA[i] = 0x18
	}
	if m.FastInboundCount != 0 {
		i = encodeVarintConfirmation(dAtA, i, uint64(m.FastInboundCount))
		i--
		dAtA[i] = 0x10
	}
	if m.SafeInboundCount != 0 {
		i = encodeVarintConfirmation(dAtA, i, uint64(m.SafeInboundCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfirmation(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfirmation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Confirmation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SafeInboundCount != 0 {
		n += 1 + sovConfirmation(uint64(m.SafeInboundCount))
	}
	if m.FastInboundCount != 0 {
		n += 1 + sovConfirmation(uint64(m.FastInboundCount))
	}
	if m.SafeOutboundCount != 0 {
		n += 1 + sovConfirmation(uint64(m.SafeOutboundCount))
	}
	if m.FastOutboundCount != 0 {
		n += 1 + sovConfirmation(uint64(m.FastOutboundCount))
	}
	return n
}

func sovConfirmation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfirmation(x uint64) (n int) {
	return sovConfirmation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Confirmation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfirmation
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
			return fmt.Errorf("proto: Confirmation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Confirmation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafeInboundCount", wireType)
			}
			m.SafeInboundCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirmation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SafeInboundCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FastInboundCount", wireType)
			}
			m.FastInboundCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirmation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FastInboundCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafeOutboundCount", wireType)
			}
			m.SafeOutboundCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirmation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SafeOutboundCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FastOutboundCount", wireType)
			}
			m.FastOutboundCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirmation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FastOutboundCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfirmation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfirmation
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
func skipConfirmation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfirmation
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
					return 0, ErrIntOverflowConfirmation
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
					return 0, ErrIntOverflowConfirmation
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
				return 0, ErrInvalidLengthConfirmation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfirmation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfirmation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfirmation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfirmation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfirmation = fmt.Errorf("proto: unexpected end of group")
)
