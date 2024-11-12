// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/crosschain/outbound_tracker.proto

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

type TxHash struct {
	TxHash   string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	TxSigner string `protobuf:"bytes,2,opt,name=tx_signer,json=txSigner,proto3" json:"tx_signer,omitempty"`
}

func (m *TxHash) Reset()         { *m = TxHash{} }
func (m *TxHash) String() string { return proto.CompactTextString(m) }
func (*TxHash) ProtoMessage()    {}
func (*TxHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_77cb2cfe04eb42d9, []int{0}
}
func (m *TxHash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxHash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxHash.Merge(m, src)
}
func (m *TxHash) XXX_Size() int {
	return m.Size()
}
func (m *TxHash) XXX_DiscardUnknown() {
	xxx_messageInfo_TxHash.DiscardUnknown(m)
}

var xxx_messageInfo_TxHash proto.InternalMessageInfo

func (m *TxHash) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *TxHash) GetTxSigner() string {
	if m != nil {
		return m.TxSigner
	}
	return ""
}

type OutboundTracker struct {
	Index    string    `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	ChainId  int64     `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Nonce    uint64    `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	HashList []*TxHash `protobuf:"bytes,4,rep,name=hash_list,json=hashList,proto3" json:"hash_list,omitempty"`
}

func (m *OutboundTracker) Reset()         { *m = OutboundTracker{} }
func (m *OutboundTracker) String() string { return proto.CompactTextString(m) }
func (*OutboundTracker) ProtoMessage()    {}
func (*OutboundTracker) Descriptor() ([]byte, []int) {
	return fileDescriptor_77cb2cfe04eb42d9, []int{1}
}
func (m *OutboundTracker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutboundTracker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutboundTracker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutboundTracker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutboundTracker.Merge(m, src)
}
func (m *OutboundTracker) XXX_Size() int {
	return m.Size()
}
func (m *OutboundTracker) XXX_DiscardUnknown() {
	xxx_messageInfo_OutboundTracker.DiscardUnknown(m)
}

var xxx_messageInfo_OutboundTracker proto.InternalMessageInfo

func (m *OutboundTracker) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *OutboundTracker) GetChainId() int64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *OutboundTracker) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *OutboundTracker) GetHashList() []*TxHash {
	if m != nil {
		return m.HashList
	}
	return nil
}

func init() {
	proto.RegisterType((*TxHash)(nil), "zetachain.zetacore.crosschain.TxHash")
	proto.RegisterType((*OutboundTracker)(nil), "zetachain.zetacore.crosschain.OutboundTracker")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/crosschain/outbound_tracker.proto", fileDescriptor_77cb2cfe04eb42d9)
}

var fileDescriptor_77cb2cfe04eb42d9 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xa9, 0x4a, 0x2d, 0x49,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0xb3, 0xf2, 0x8b, 0x52, 0xf5, 0x93, 0x8b, 0xf2, 0x8b,
	0x8b, 0x21, 0x62, 0xf9, 0xa5, 0x25, 0x49, 0xf9, 0xa5, 0x79, 0x29, 0xf1, 0x25, 0x45, 0x89, 0xc9,
	0xd9, 0xa9, 0x45, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xb2, 0x70, 0x5d, 0x7a, 0x30, 0x5d,
	0x7a, 0x08, 0x5d, 0x4a, 0x2e, 0x5c, 0x6c, 0x21, 0x15, 0x1e, 0x89, 0xc5, 0x19, 0x42, 0xe2, 0x5c,
	0xec, 0x25, 0x15, 0xf1, 0x19, 0x89, 0xc5, 0x19, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x6c,
	0x25, 0x10, 0x09, 0x69, 0x2e, 0xce, 0x92, 0x8a, 0xf8, 0xe2, 0xcc, 0xf4, 0xbc, 0xd4, 0x22, 0x09,
	0x26, 0xb0, 0x14, 0x47, 0x49, 0x45, 0x30, 0x98, 0xef, 0xc5, 0xc2, 0xc1, 0x2c, 0xc0, 0xa2, 0x34,
	0x87, 0x91, 0x8b, 0xdf, 0x1f, 0x6a, 0x7f, 0x08, 0xc4, 0x7a, 0x21, 0x11, 0x2e, 0xd6, 0xcc, 0xbc,
	0x94, 0xd4, 0x0a, 0xa8, 0x69, 0x10, 0x8e, 0x90, 0x24, 0x17, 0x07, 0xd8, 0xe2, 0xf8, 0xcc, 0x14,
	0xb0, 0x59, 0xcc, 0x41, 0xec, 0x60, 0xbe, 0x67, 0x0a, 0x48, 0x43, 0x5e, 0x7e, 0x5e, 0x72, 0xaa,
	0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x84, 0x23, 0xe4, 0xc4, 0xc5, 0x09, 0x72, 0x53, 0x7c,
	0x4e, 0x66, 0x71, 0x89, 0x04, 0x8b, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0xaa, 0x1e, 0x5e, 0x3f, 0xe9,
	0x41, 0x3c, 0x14, 0xc4, 0x01, 0xd2, 0xe7, 0x93, 0x59, 0x5c, 0xe2, 0xe4, 0x7e, 0xe2, 0x91, 0x1c,
	0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1,
	0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xe0, 0x40, 0xd5, 0x85, 0x84, 0x65, 0x5e, 0x7e, 0x4a, 0xaa, 0x7e, 0x05, 0x72, 0xe8,
	0x96, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xc3, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xd8, 0x18, 0x1c, 0x85, 0x8b, 0x01, 0x00, 0x00,
}

func (m *TxHash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxHash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxHash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxSigner) > 0 {
		i -= len(m.TxSigner)
		copy(dAtA[i:], m.TxSigner)
		i = encodeVarintOutboundTracker(dAtA, i, uint64(len(m.TxSigner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintOutboundTracker(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OutboundTracker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutboundTracker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutboundTracker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HashList) > 0 {
		for iNdEx := len(m.HashList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HashList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOutboundTracker(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Nonce != 0 {
		i = encodeVarintOutboundTracker(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x18
	}
	if m.ChainId != 0 {
		i = encodeVarintOutboundTracker(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintOutboundTracker(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOutboundTracker(dAtA []byte, offset int, v uint64) int {
	offset -= sovOutboundTracker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxHash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovOutboundTracker(uint64(l))
	}
	l = len(m.TxSigner)
	if l > 0 {
		n += 1 + l + sovOutboundTracker(uint64(l))
	}
	return n
}

func (m *OutboundTracker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovOutboundTracker(uint64(l))
	}
	if m.ChainId != 0 {
		n += 1 + sovOutboundTracker(uint64(m.ChainId))
	}
	if m.Nonce != 0 {
		n += 1 + sovOutboundTracker(uint64(m.Nonce))
	}
	if len(m.HashList) > 0 {
		for _, e := range m.HashList {
			l = e.Size()
			n += 1 + l + sovOutboundTracker(uint64(l))
		}
	}
	return n
}

func sovOutboundTracker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOutboundTracker(x uint64) (n int) {
	return sovOutboundTracker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxHash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOutboundTracker
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
			return fmt.Errorf("proto: TxHash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxHash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTracker
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
				return ErrInvalidLengthOutboundTracker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTracker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxSigner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTracker
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
				return ErrInvalidLengthOutboundTracker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTracker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxSigner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOutboundTracker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOutboundTracker
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
func (m *OutboundTracker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOutboundTracker
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
			return fmt.Errorf("proto: OutboundTracker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutboundTracker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTracker
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
				return ErrInvalidLengthOutboundTracker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTracker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTracker
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
				return ErrInvalidLengthOutboundTracker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTracker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashList = append(m.HashList, &TxHash{})
			if err := m.HashList[len(m.HashList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOutboundTracker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOutboundTracker
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
func skipOutboundTracker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOutboundTracker
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
					return 0, ErrIntOverflowOutboundTracker
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
					return 0, ErrIntOverflowOutboundTracker
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
				return 0, ErrInvalidLengthOutboundTracker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOutboundTracker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOutboundTracker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOutboundTracker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOutboundTracker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOutboundTracker = fmt.Errorf("proto: unexpected end of group")
)
