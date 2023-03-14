// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crosschain/keygen.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Keygen struct {
	Creator     string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Status      uint64   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Pubkeys     []string `protobuf:"bytes,3,rep,name=pubkeys,proto3" json:"pubkeys,omitempty"`
	BlockNumber int64    `protobuf:"varint,4,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
}

func (m *Keygen) Reset()         { *m = Keygen{} }
func (m *Keygen) String() string { return proto.CompactTextString(m) }
func (*Keygen) ProtoMessage()    {}
func (*Keygen) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2a31a17f5bb2128, []int{0}
}
func (m *Keygen) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Keygen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Keygen.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Keygen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keygen.Merge(m, src)
}
func (m *Keygen) XXX_Size() int {
	return m.Size()
}
func (m *Keygen) XXX_DiscardUnknown() {
	xxx_messageInfo_Keygen.DiscardUnknown(m)
}

var xxx_messageInfo_Keygen proto.InternalMessageInfo

func (m *Keygen) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Keygen) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Keygen) GetPubkeys() []string {
	if m != nil {
		return m.Pubkeys
	}
	return nil
}

func (m *Keygen) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*Keygen)(nil), "zetachain.zetacore.crosschain.Keygen")
}

func init() { proto.RegisterFile("crosschain/keygen.proto", fileDescriptor_b2a31a17f5bb2128) }

var fileDescriptor_b2a31a17f5bb2128 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2e, 0xca, 0x2f,
	0x2e, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0x4e, 0xad, 0x4c, 0x4f, 0xcd, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0xad, 0x4a, 0x2d, 0x49, 0x04, 0x8b, 0xeb, 0x81, 0x59, 0xf9, 0x45,
	0xa9, 0x7a, 0x08, 0xb5, 0x4a, 0x65, 0x5c, 0x6c, 0xde, 0x60, 0xe5, 0x42, 0x12, 0x5c, 0xec, 0xc9,
	0x45, 0xa9, 0x89, 0x25, 0xf9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90,
	0x18, 0x17, 0x5b, 0x71, 0x49, 0x62, 0x49, 0x69, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4b, 0x10,
	0x94, 0x07, 0xd2, 0x51, 0x50, 0x9a, 0x94, 0x9d, 0x5a, 0x59, 0x2c, 0xc1, 0xac, 0xc0, 0x0c, 0xd2,
	0x01, 0xe5, 0x0a, 0x29, 0x70, 0x71, 0x27, 0xe5, 0xe4, 0x27, 0x67, 0xfb, 0x95, 0xe6, 0x26, 0xa5,
	0x16, 0x49, 0xb0, 0x28, 0x30, 0x6a, 0x30, 0x07, 0x21, 0x0b, 0x39, 0x79, 0x9f, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c,
	0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x61, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x3e, 0xc8, 0xc5, 0xba, 0x10, 0x4f, 0xc1, 0x1c, 0xaf, 0x5f, 0xa1, 0x8f, 0xe4, 0xd5,
	0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x57, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x7f, 0x19, 0x28, 0x7f, 0x05, 0x01, 0x00, 0x00,
}

func (m *Keygen) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Keygen) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Keygen) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockNumber != 0 {
		i = encodeVarintKeygen(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Pubkeys) > 0 {
		for iNdEx := len(m.Pubkeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Pubkeys[iNdEx])
			copy(dAtA[i:], m.Pubkeys[iNdEx])
			i = encodeVarintKeygen(dAtA, i, uint64(len(m.Pubkeys[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Status != 0 {
		i = encodeVarintKeygen(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintKeygen(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeygen(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeygen(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Keygen) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovKeygen(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovKeygen(uint64(m.Status))
	}
	if len(m.Pubkeys) > 0 {
		for _, s := range m.Pubkeys {
			l = len(s)
			n += 1 + l + sovKeygen(uint64(l))
		}
	}
	if m.BlockNumber != 0 {
		n += 1 + sovKeygen(uint64(m.BlockNumber))
	}
	return n
}

func sovKeygen(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeygen(x uint64) (n int) {
	return sovKeygen(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Keygen) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeygen
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
			return fmt.Errorf("proto: Keygen: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Keygen: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
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
				return ErrInvalidLengthKeygen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeygen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
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
				return ErrInvalidLengthKeygen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeygen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkeys = append(m.Pubkeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKeygen(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeygen
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
func skipKeygen(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeygen
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
					return 0, ErrIntOverflowKeygen
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
					return 0, ErrIntOverflowKeygen
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
				return 0, ErrInvalidLengthKeygen
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeygen
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeygen
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeygen        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeygen          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeygen = fmt.Errorf("proto: unexpected end of group")
)
