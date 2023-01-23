// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: emissions/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgAddTokenEmission struct {
	Creator  string                                 `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Category EmissionCategory                       `protobuf:"varint,2,opt,name=category,proto3,enum=zetachain.zetacore.emissions.EmissionCategory" json:"category,omitempty"`
	Amount   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=Amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"Amount" yaml:"amount"`
}

func (m *MsgAddTokenEmission) Reset()         { *m = MsgAddTokenEmission{} }
func (m *MsgAddTokenEmission) String() string { return proto.CompactTextString(m) }
func (*MsgAddTokenEmission) ProtoMessage()    {}
func (*MsgAddTokenEmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_618f91fd090d1520, []int{0}
}
func (m *MsgAddTokenEmission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddTokenEmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddTokenEmission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddTokenEmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddTokenEmission.Merge(m, src)
}
func (m *MsgAddTokenEmission) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddTokenEmission) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddTokenEmission.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddTokenEmission proto.InternalMessageInfo

func (m *MsgAddTokenEmission) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAddTokenEmission) GetCategory() EmissionCategory {
	if m != nil {
		return m.Category
	}
	return EmissionCategory_ObserverEmission
}

type MsgAddTokenEmissionResponse struct {
}

func (m *MsgAddTokenEmissionResponse) Reset()         { *m = MsgAddTokenEmissionResponse{} }
func (m *MsgAddTokenEmissionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddTokenEmissionResponse) ProtoMessage()    {}
func (*MsgAddTokenEmissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_618f91fd090d1520, []int{1}
}
func (m *MsgAddTokenEmissionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddTokenEmissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddTokenEmissionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddTokenEmissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddTokenEmissionResponse.Merge(m, src)
}
func (m *MsgAddTokenEmissionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddTokenEmissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddTokenEmissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddTokenEmissionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddTokenEmission)(nil), "zetachain.zetacore.emissions.MsgAddTokenEmission")
	proto.RegisterType((*MsgAddTokenEmissionResponse)(nil), "zetachain.zetacore.emissions.MsgAddTokenEmissionResponse")
}

func init() { proto.RegisterFile("emissions/tx.proto", fileDescriptor_618f91fd090d1520) }

var fileDescriptor_618f91fd090d1520 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xcd, 0xcd, 0x2c,
	0x2e, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0xa9, 0x4a, 0x2d, 0x49, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x03, 0xb3, 0xf2, 0x8b, 0x52, 0xf5,
	0xe0, 0xca, 0xa4, 0x14, 0x10, 0x3a, 0x60, 0xac, 0xf8, 0x92, 0xa2, 0xc4, 0xe4, 0xec, 0xd4, 0x22,
	0x88, 0x7e, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x30, 0x53, 0x1f, 0xc4, 0x82, 0x88, 0x2a, 0x5d,
	0x61, 0xe4, 0x12, 0xf6, 0x2d, 0x4e, 0x77, 0x4c, 0x49, 0x09, 0xc9, 0xcf, 0x4e, 0xcd, 0x73, 0x85,
	0xea, 0x15, 0x92, 0xe0, 0x62, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xbc, 0xb8, 0x38, 0x92, 0x13, 0x4b, 0x52, 0xd3, 0xf3, 0x8b,
	0x2a, 0x25, 0x98, 0x14, 0x18, 0x35, 0xf8, 0x8c, 0xf4, 0xf4, 0xf0, 0x39, 0x4d, 0x0f, 0x66, 0xa6,
	0x33, 0x54, 0x57, 0x10, 0x5c, 0xbf, 0x50, 0x38, 0x17, 0x9b, 0x63, 0x6e, 0x7e, 0x69, 0x5e, 0x89,
	0x04, 0x33, 0xc8, 0x12, 0x27, 0xfb, 0x13, 0xf7, 0xe4, 0x19, 0x6e, 0xdd, 0x93, 0x57, 0x4b, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x86,
	0x52, 0xba, 0xc5, 0x29, 0xd9, 0xfa, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x7a, 0x9e, 0x79, 0x25, 0x9f,
	0xee, 0xc9, 0xf3, 0x56, 0x26, 0xe6, 0xe6, 0x58, 0x29, 0x25, 0x82, 0x4d, 0x51, 0x0a, 0x82, 0x1a,
	0xa7, 0x24, 0xcb, 0x25, 0x8d, 0xc5, 0x57, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x46,
	0x1d, 0x8c, 0x5c, 0xcc, 0xbe, 0xc5, 0xe9, 0x42, 0x0d, 0x8c, 0x5c, 0x02, 0x18, 0x5e, 0x37, 0xc4,
	0xef, 0x1d, 0x2c, 0xe6, 0x4a, 0x59, 0x92, 0xac, 0x05, 0xe6, 0x14, 0x27, 0xaf, 0x13, 0x8f, 0xe4,
	0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f,
	0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x40, 0x0a, 0x04, 0x90, 0xa1, 0xba, 0x60, 0xf3,
	0xf5, 0x61, 0xe6, 0xeb, 0x57, 0xe8, 0x23, 0xa5, 0x12, 0x50, 0x90, 0x24, 0xb1, 0x81, 0xe3, 0xd4,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x42, 0x89, 0xb9, 0x04, 0x3f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	AddTokenEmission(ctx context.Context, in *MsgAddTokenEmission, opts ...grpc.CallOption) (*MsgAddTokenEmissionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddTokenEmission(ctx context.Context, in *MsgAddTokenEmission, opts ...grpc.CallOption) (*MsgAddTokenEmissionResponse, error) {
	out := new(MsgAddTokenEmissionResponse)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.emissions.Msg/AddTokenEmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AddTokenEmission(context.Context, *MsgAddTokenEmission) (*MsgAddTokenEmissionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddTokenEmission(ctx context.Context, req *MsgAddTokenEmission) (*MsgAddTokenEmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTokenEmission not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddTokenEmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddTokenEmission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddTokenEmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.emissions.Msg/AddTokenEmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddTokenEmission(ctx, req.(*MsgAddTokenEmission))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zetachain.zetacore.emissions.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTokenEmission",
			Handler:    _Msg_AddTokenEmission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emissions/tx.proto",
}

func (m *MsgAddTokenEmission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddTokenEmission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddTokenEmission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Category != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Category))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddTokenEmissionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddTokenEmissionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddTokenEmissionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddTokenEmission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Category != 0 {
		n += 1 + sovTx(uint64(m.Category))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgAddTokenEmissionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddTokenEmission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddTokenEmission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddTokenEmission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			m.Category = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Category |= EmissionCategory(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddTokenEmissionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddTokenEmissionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddTokenEmissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
