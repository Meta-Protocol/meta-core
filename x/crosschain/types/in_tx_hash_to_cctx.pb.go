// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/crosschain/in_tx_hash_to_cctx.proto

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

type InTxHashToCctx struct {
	InTxHash  string   `protobuf:"bytes,1,opt,name=in_tx_hash,json=inTxHash,proto3" json:"in_tx_hash,omitempty"`
	CctxIndex []string `protobuf:"bytes,2,rep,name=cctx_index,json=cctxIndex,proto3" json:"cctx_index,omitempty"`
}

func (m *InTxHashToCctx) Reset()         { *m = InTxHashToCctx{} }
func (m *InTxHashToCctx) String() string { return proto.CompactTextString(m) }
func (*InTxHashToCctx) ProtoMessage()    {}
func (*InTxHashToCctx) Descriptor() ([]byte, []int) {
	return fileDescriptor_fba0f1091d4145fb, []int{0}
}
func (m *InTxHashToCctx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InTxHashToCctx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InTxHashToCctx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InTxHashToCctx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InTxHashToCctx.Merge(m, src)
}
func (m *InTxHashToCctx) XXX_Size() int {
	return m.Size()
}
func (m *InTxHashToCctx) XXX_DiscardUnknown() {
	xxx_messageInfo_InTxHashToCctx.DiscardUnknown(m)
}

var xxx_messageInfo_InTxHashToCctx proto.InternalMessageInfo

func (m *InTxHashToCctx) GetInTxHash() string {
	if m != nil {
		return m.InTxHash
	}
	return ""
}

func (m *InTxHashToCctx) GetCctxIndex() []string {
	if m != nil {
		return m.CctxIndex
	}
	return nil
}

func init() {
	proto.RegisterType((*InTxHashToCctx)(nil), "zetachain.zetacore.crosschain.InTxHashToCctx")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/crosschain/in_tx_hash_to_cctx.proto", fileDescriptor_fba0f1091d4145fb)
}

var fileDescriptor_fba0f1091d4145fb = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xab, 0x4a, 0x2d, 0x49,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0xb3, 0xf2, 0x8b, 0x52, 0xf5, 0x93, 0x8b, 0xf2, 0x8b,
	0x8b, 0x21, 0x62, 0x99, 0x79, 0xf1, 0x25, 0x15, 0xf1, 0x19, 0x89, 0xc5, 0x19, 0xf1, 0x25, 0xf9,
	0xf1, 0xc9, 0xc9, 0x25, 0x15, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xb2, 0x70, 0x7d, 0x7a,
	0x30, 0x7d, 0x7a, 0x08, 0x7d, 0x4a, 0xbe, 0x5c, 0x7c, 0x9e, 0x79, 0x21, 0x15, 0x1e, 0x89, 0xc5,
	0x19, 0x21, 0xf9, 0xce, 0xc9, 0x25, 0x15, 0x42, 0x32, 0x5c, 0x5c, 0x08, 0xc3, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x38, 0x32, 0xa1, 0x6a, 0x84, 0x64, 0xb9, 0xb8, 0x40, 0x86, 0xc7, 0x67,
	0xe6, 0xa5, 0xa4, 0x56, 0x48, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x06, 0x71, 0x82, 0x44, 0x3c, 0x41,
	0x02, 0x4e, 0xde, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x98, 0x9e,
	0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0x0b, 0xf6, 0x80, 0x2e, 0x9a, 0x5f, 0x2a, 0x90,
	0x7d, 0x53, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0x81, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x70, 0x97, 0x6c, 0x57, 0xfb, 0x00, 0x00, 0x00,
}

func (m *InTxHashToCctx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InTxHashToCctx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InTxHashToCctx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CctxIndex) > 0 {
		for iNdEx := len(m.CctxIndex) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CctxIndex[iNdEx])
			copy(dAtA[i:], m.CctxIndex[iNdEx])
			i = encodeVarintInTxHashToCctx(dAtA, i, uint64(len(m.CctxIndex[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.InTxHash) > 0 {
		i -= len(m.InTxHash)
		copy(dAtA[i:], m.InTxHash)
		i = encodeVarintInTxHashToCctx(dAtA, i, uint64(len(m.InTxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInTxHashToCctx(dAtA []byte, offset int, v uint64) int {
	offset -= sovInTxHashToCctx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InTxHashToCctx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InTxHash)
	if l > 0 {
		n += 1 + l + sovInTxHashToCctx(uint64(l))
	}
	if len(m.CctxIndex) > 0 {
		for _, s := range m.CctxIndex {
			l = len(s)
			n += 1 + l + sovInTxHashToCctx(uint64(l))
		}
	}
	return n
}

func sovInTxHashToCctx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInTxHashToCctx(x uint64) (n int) {
	return sovInTxHashToCctx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InTxHashToCctx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInTxHashToCctx
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
			return fmt.Errorf("proto: InTxHashToCctx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InTxHashToCctx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInTxHashToCctx
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
				return ErrInvalidLengthInTxHashToCctx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInTxHashToCctx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CctxIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInTxHashToCctx
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
				return ErrInvalidLengthInTxHashToCctx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInTxHashToCctx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CctxIndex = append(m.CctxIndex, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInTxHashToCctx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInTxHashToCctx
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
func skipInTxHashToCctx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInTxHashToCctx
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
					return 0, ErrIntOverflowInTxHashToCctx
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
					return 0, ErrIntOverflowInTxHashToCctx
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
				return 0, ErrInvalidLengthInTxHashToCctx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInTxHashToCctx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInTxHashToCctx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInTxHashToCctx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInTxHashToCctx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInTxHashToCctx = fmt.Errorf("proto: unexpected end of group")
)
