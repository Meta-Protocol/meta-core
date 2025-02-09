// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/observer/confirmation_params.proto

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

type ConfirmationParams struct {
	// This is the safe number of confirmations to wait before an inbound is
	// considered finalized.
	SafeInboundCount uint64 `protobuf:"varint,1,opt,name=safe_inbound_count,json=safeInboundCount,proto3" json:"safe_inbound_count,omitempty"`
	// This is the number of confirmations for fast inbound observation, which is
	// shorter than safe_inbound_count.
	FastInboundCount uint64 `protobuf:"varint,2,opt,name=fast_inbound_count,json=fastInboundCount,proto3" json:"fast_inbound_count,omitempty"`
	// This is the safe number of confirmations to wait before an outbound is
	// considered finalized.
	SafeOutboundCount uint64 `protobuf:"varint,3,opt,name=safe_outbound_count,json=safeOutboundCount,proto3" json:"safe_outbound_count,omitempty"`
	// This is the number of confirmations for fast outbound observation, which is
	// shorter than safe_outbound_count.
	FastOutboundCount uint64 `protobuf:"varint,4,opt,name=fast_outbound_count,json=fastOutboundCount,proto3" json:"fast_outbound_count,omitempty"`
}

func (m *ConfirmationParams) Reset()         { *m = ConfirmationParams{} }
func (m *ConfirmationParams) String() string { return proto.CompactTextString(m) }
func (*ConfirmationParams) ProtoMessage()    {}
func (*ConfirmationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d50fce477b3e8db, []int{0}
}
func (m *ConfirmationParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfirmationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfirmationParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfirmationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmationParams.Merge(m, src)
}
func (m *ConfirmationParams) XXX_Size() int {
	return m.Size()
}
func (m *ConfirmationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmationParams.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmationParams proto.InternalMessageInfo

func (m *ConfirmationParams) GetSafeInboundCount() uint64 {
	if m != nil {
		return m.SafeInboundCount
	}
	return 0
}

func (m *ConfirmationParams) GetFastInboundCount() uint64 {
	if m != nil {
		return m.FastInboundCount
	}
	return 0
}

func (m *ConfirmationParams) GetSafeOutboundCount() uint64 {
	if m != nil {
		return m.SafeOutboundCount
	}
	return 0
}

func (m *ConfirmationParams) GetFastOutboundCount() uint64 {
	if m != nil {
		return m.FastOutboundCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ConfirmationParams)(nil), "zetachain.zetacore.observer.ConfirmationParams")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/observer/confirmation_params.proto", fileDescriptor_3d50fce477b3e8db)
}

var fileDescriptor_3d50fce477b3e8db = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbf, 0x4e, 0xc5, 0x20,
	0x14, 0x87, 0x8b, 0xde, 0x38, 0x74, 0xd2, 0xba, 0x98, 0x98, 0x10, 0xe3, 0x64, 0xa2, 0xc2, 0x60,
	0x7c, 0x01, 0x6f, 0x1c, 0x9c, 0x34, 0x8e, 0x2e, 0x0d, 0x70, 0xa9, 0x97, 0xa1, 0x9c, 0x06, 0x0e,
	0x46, 0x7d, 0x0a, 0x1f, 0xcb, 0xb1, 0xa3, 0xa3, 0x69, 0x5f, 0xc4, 0x00, 0xde, 0x7f, 0xdd, 0x08,
	0xdf, 0xf7, 0x9d, 0xe1, 0x57, 0xde, 0x7e, 0x6a, 0x14, 0x6a, 0x29, 0x8c, 0xe5, 0xe9, 0x05, 0x4e,
	0x73, 0x90, 0x5e, 0xbb, 0x37, 0xed, 0xb8, 0x02, 0xdb, 0x18, 0xd7, 0x0a, 0x34, 0x60, 0xeb, 0x4e,
	0x38, 0xd1, 0x7a, 0xd6, 0x39, 0x40, 0xa8, 0x4e, 0xd7, 0x19, 0x5b, 0x65, 0x6c, 0x95, 0x9d, 0xf7,
	0xa4, 0xac, 0xe6, 0x5b, 0xe9, 0x53, 0x2a, 0xab, 0xab, 0xb2, 0xf2, 0xa2, 0xd1, 0xb5, 0xb1, 0x12,
	0x82, 0x5d, 0xd4, 0x0a, 0x82, 0xc5, 0x13, 0x72, 0x46, 0x2e, 0x66, 0xcf, 0x87, 0x91, 0x3c, 0x64,
	0x30, 0x8f, 0xff, 0xd1, 0x6e, 0x84, 0xc7, 0x89, 0xbd, 0x97, 0xed, 0x48, 0x76, 0x6c, 0x56, 0x1e,
	0xa7, 0xdb, 0x10, 0x70, 0x5b, 0xdf, 0x4f, 0xfa, 0x51, 0x44, 0x8f, 0xff, 0x64, 0xed, 0xa7, 0xeb,
	0x13, 0x7f, 0x96, 0xfd, 0x88, 0x76, 0xfc, 0xbb, 0xfb, 0xef, 0x81, 0x92, 0x7e, 0xa0, 0xe4, 0x77,
	0xa0, 0xe4, 0x6b, 0xa4, 0x45, 0x3f, 0xd2, 0xe2, 0x67, 0xa4, 0xc5, 0xcb, 0xe5, 0xab, 0xc1, 0x65,
	0x90, 0x4c, 0x41, 0x9b, 0x16, 0xbc, 0xce, 0x63, 0x5a, 0x58, 0x68, 0xfe, 0xbe, 0x99, 0x12, 0x3f,
	0x3a, 0xed, 0xe5, 0x41, 0x5a, 0xef, 0xe6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x5a, 0xf5, 0x7c,
	0x76, 0x01, 0x00, 0x00,
}

func (m *ConfirmationParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfirmationParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfirmationParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FastOutboundCount != 0 {
		i = encodeVarintConfirmationParams(dAtA, i, uint64(m.FastOutboundCount))
		i--
		dAtA[i] = 0x20
	}
	if m.SafeOutboundCount != 0 {
		i = encodeVarintConfirmationParams(dAtA, i, uint64(m.SafeOutboundCount))
		i--
		dAtA[i] = 0x18
	}
	if m.FastInboundCount != 0 {
		i = encodeVarintConfirmationParams(dAtA, i, uint64(m.FastInboundCount))
		i--
		dAtA[i] = 0x10
	}
	if m.SafeInboundCount != 0 {
		i = encodeVarintConfirmationParams(dAtA, i, uint64(m.SafeInboundCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfirmationParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfirmationParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConfirmationParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SafeInboundCount != 0 {
		n += 1 + sovConfirmationParams(uint64(m.SafeInboundCount))
	}
	if m.FastInboundCount != 0 {
		n += 1 + sovConfirmationParams(uint64(m.FastInboundCount))
	}
	if m.SafeOutboundCount != 0 {
		n += 1 + sovConfirmationParams(uint64(m.SafeOutboundCount))
	}
	if m.FastOutboundCount != 0 {
		n += 1 + sovConfirmationParams(uint64(m.FastOutboundCount))
	}
	return n
}

func sovConfirmationParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfirmationParams(x uint64) (n int) {
	return sovConfirmationParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConfirmationParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfirmationParams
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
			return fmt.Errorf("proto: ConfirmationParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfirmationParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafeInboundCount", wireType)
			}
			m.SafeInboundCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirmationParams
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
					return ErrIntOverflowConfirmationParams
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
					return ErrIntOverflowConfirmationParams
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
					return ErrIntOverflowConfirmationParams
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
			skippy, err := skipConfirmationParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfirmationParams
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
func skipConfirmationParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfirmationParams
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
					return 0, ErrIntOverflowConfirmationParams
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
					return 0, ErrIntOverflowConfirmationParams
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
				return 0, ErrInvalidLengthConfirmationParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfirmationParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfirmationParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfirmationParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfirmationParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfirmationParams = fmt.Errorf("proto: unexpected end of group")
)
