// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: observer/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	common "github.com/zeta-chain/zetacore/common"
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

type MsgSetSupportedChains struct {
	Creator   string           `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ChainId   int64            `protobuf:"varint,2,opt,name=ChainId,proto3" json:"ChainId,omitempty"`
	ChainName common.ChainName `protobuf:"varint,3,opt,name=ChainName,proto3,enum=common.ChainName" json:"ChainName,omitempty"`
}

func (m *MsgSetSupportedChains) Reset()         { *m = MsgSetSupportedChains{} }
func (m *MsgSetSupportedChains) String() string { return proto.CompactTextString(m) }
func (*MsgSetSupportedChains) ProtoMessage()    {}
func (*MsgSetSupportedChains) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bcd40fa296a2b1d, []int{0}
}
func (m *MsgSetSupportedChains) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetSupportedChains) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetSupportedChains.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetSupportedChains) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetSupportedChains.Merge(m, src)
}
func (m *MsgSetSupportedChains) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetSupportedChains) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetSupportedChains.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetSupportedChains proto.InternalMessageInfo

func (m *MsgSetSupportedChains) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSetSupportedChains) GetChainId() int64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *MsgSetSupportedChains) GetChainName() common.ChainName {
	if m != nil {
		return m.ChainName
	}
	return common.ChainName_Empty
}

type MsgSetSupportedChainsResponse struct {
}

func (m *MsgSetSupportedChainsResponse) Reset()         { *m = MsgSetSupportedChainsResponse{} }
func (m *MsgSetSupportedChainsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetSupportedChainsResponse) ProtoMessage()    {}
func (*MsgSetSupportedChainsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bcd40fa296a2b1d, []int{1}
}
func (m *MsgSetSupportedChainsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetSupportedChainsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetSupportedChainsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetSupportedChainsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetSupportedChainsResponse.Merge(m, src)
}
func (m *MsgSetSupportedChainsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetSupportedChainsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetSupportedChainsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetSupportedChainsResponse proto.InternalMessageInfo

type MsgAddObserver struct {
	Creator         string          `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ChainId         int64           `protobuf:"varint,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	ObservationType ObservationType `protobuf:"varint,3,opt,name=observationType,proto3,enum=zetachain.zetacore.observer.ObservationType" json:"observationType,omitempty"`
}

func (m *MsgAddObserver) Reset()         { *m = MsgAddObserver{} }
func (m *MsgAddObserver) String() string { return proto.CompactTextString(m) }
func (*MsgAddObserver) ProtoMessage()    {}
func (*MsgAddObserver) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bcd40fa296a2b1d, []int{2}
}
func (m *MsgAddObserver) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddObserver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddObserver.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddObserver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddObserver.Merge(m, src)
}
func (m *MsgAddObserver) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddObserver) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddObserver.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddObserver proto.InternalMessageInfo

func (m *MsgAddObserver) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAddObserver) GetChainId() int64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *MsgAddObserver) GetObservationType() ObservationType {
	if m != nil {
		return m.ObservationType
	}
	return ObservationType_EmptyObserverType
}

type MsgAddObserverResponse struct {
}

func (m *MsgAddObserverResponse) Reset()         { *m = MsgAddObserverResponse{} }
func (m *MsgAddObserverResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddObserverResponse) ProtoMessage()    {}
func (*MsgAddObserverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bcd40fa296a2b1d, []int{3}
}
func (m *MsgAddObserverResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddObserverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddObserverResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddObserverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddObserverResponse.Merge(m, src)
}
func (m *MsgAddObserverResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddObserverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddObserverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddObserverResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetSupportedChains)(nil), "zetachain.zetacore.observer.MsgSetSupportedChains")
	proto.RegisterType((*MsgSetSupportedChainsResponse)(nil), "zetachain.zetacore.observer.MsgSetSupportedChainsResponse")
	proto.RegisterType((*MsgAddObserver)(nil), "zetachain.zetacore.observer.MsgAddObserver")
	proto.RegisterType((*MsgAddObserverResponse)(nil), "zetachain.zetacore.observer.MsgAddObserverResponse")
}

func init() { proto.RegisterFile("observer/tx.proto", fileDescriptor_1bcd40fa296a2b1d) }

var fileDescriptor_1bcd40fa296a2b1d = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x4f, 0x2a, 0x4e,
	0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xae,
	0x4a, 0x2d, 0x49, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x03, 0xb3, 0xf2, 0x8b, 0x52, 0xf5, 0x60,
	0xaa, 0xa4, 0x84, 0x93, 0xf3, 0x73, 0x73, 0xf3, 0xf3, 0xf4, 0x21, 0x14, 0x44, 0x87, 0x94, 0x48,
	0x7a, 0x7e, 0x7a, 0x3e, 0x98, 0xa9, 0x0f, 0x62, 0x41, 0x45, 0xc5, 0xe1, 0x46, 0xc3, 0x18, 0x10,
	0x09, 0xa5, 0x1a, 0x2e, 0x51, 0xdf, 0xe2, 0xf4, 0xe0, 0xd4, 0x92, 0xe0, 0xd2, 0x82, 0x82, 0xfc,
	0xa2, 0x92, 0xd4, 0x14, 0x67, 0x90, 0x6d, 0xc5, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0x45, 0xa9, 0x89,
	0x25, 0xf9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0x2e, 0x48, 0x06, 0xac, 0xc6,
	0x33, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x08, 0xc6, 0x15, 0xd2, 0xe7, 0xe2, 0x04, 0x33,
	0xfd, 0x12, 0x73, 0x53, 0x25, 0x98, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x04, 0xf5, 0xa0, 0xae, 0x83,
	0x4b, 0x04, 0x21, 0xd4, 0x28, 0xc9, 0x73, 0xc9, 0x62, 0xb5, 0x3d, 0x28, 0xb5, 0xb8, 0x20, 0x3f,
	0xaf, 0x38, 0x55, 0x69, 0x0e, 0x23, 0x17, 0x9f, 0x6f, 0x71, 0xba, 0x63, 0x4a, 0x8a, 0x3f, 0xd4,
	0xdd, 0xf8, 0x1d, 0x96, 0x8c, 0xea, 0x30, 0x28, 0x57, 0x28, 0x8c, 0x8b, 0x1f, 0xe2, 0xef, 0xc4,
	0x92, 0xcc, 0xfc, 0xbc, 0x90, 0xca, 0x02, 0x98, 0xf3, 0x74, 0xf4, 0xf0, 0x04, 0xb0, 0x9e, 0x3f,
	0xaa, 0x9e, 0x20, 0x74, 0x43, 0x94, 0x24, 0xb8, 0xc4, 0x50, 0x5d, 0x07, 0x73, 0xb8, 0xd1, 0x5f,
	0x46, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x16, 0x46, 0x2e, 0x21, 0x2c, 0xa1, 0x6b, 0x84, 0xd7,
	0x5e, 0xac, 0x61, 0x22, 0x65, 0x45, 0xba, 0x1e, 0x98, 0x73, 0x84, 0xf2, 0xb9, 0xb8, 0x91, 0xc3,
	0x50, 0x9b, 0x90, 0x51, 0x48, 0x8a, 0xa5, 0x8c, 0x49, 0x50, 0x0c, 0xb3, 0xd0, 0xc9, 0xf3, 0xc4,
	0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1,
	0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xf4, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93,
	0x40, 0xe9, 0x42, 0x1f, 0x64, 0x9c, 0x2e, 0xd8, 0x64, 0x7d, 0x98, 0xc9, 0xfa, 0x15, 0xfa, 0x88,
	0x6c, 0x50, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x4e, 0xa9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1a, 0xf0, 0xdc, 0x88, 0x1f, 0x03, 0x00, 0x00,
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
	SetSupportedChains(ctx context.Context, in *MsgSetSupportedChains, opts ...grpc.CallOption) (*MsgSetSupportedChainsResponse, error)
	AddObserver(ctx context.Context, in *MsgAddObserver, opts ...grpc.CallOption) (*MsgAddObserverResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetSupportedChains(ctx context.Context, in *MsgSetSupportedChains, opts ...grpc.CallOption) (*MsgSetSupportedChainsResponse, error) {
	out := new(MsgSetSupportedChainsResponse)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.observer.Msg/SetSupportedChains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddObserver(ctx context.Context, in *MsgAddObserver, opts ...grpc.CallOption) (*MsgAddObserverResponse, error) {
	out := new(MsgAddObserverResponse)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.observer.Msg/AddObserver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SetSupportedChains(context.Context, *MsgSetSupportedChains) (*MsgSetSupportedChainsResponse, error)
	AddObserver(context.Context, *MsgAddObserver) (*MsgAddObserverResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetSupportedChains(ctx context.Context, req *MsgSetSupportedChains) (*MsgSetSupportedChainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSupportedChains not implemented")
}
func (*UnimplementedMsgServer) AddObserver(ctx context.Context, req *MsgAddObserver) (*MsgAddObserverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddObserver not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetSupportedChains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetSupportedChains)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetSupportedChains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.observer.Msg/SetSupportedChains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetSupportedChains(ctx, req.(*MsgSetSupportedChains))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddObserver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddObserver)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddObserver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.observer.Msg/AddObserver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddObserver(ctx, req.(*MsgAddObserver))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zetachain.zetacore.observer.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetSupportedChains",
			Handler:    _Msg_SetSupportedChains_Handler,
		},
		{
			MethodName: "AddObserver",
			Handler:    _Msg_AddObserver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "observer/tx.proto",
}

func (m *MsgSetSupportedChains) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetSupportedChains) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetSupportedChains) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ChainName != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ChainName))
		i--
		dAtA[i] = 0x18
	}
	if m.ChainId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ChainId))
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

func (m *MsgSetSupportedChainsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetSupportedChainsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetSupportedChainsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgAddObserver) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddObserver) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddObserver) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ObservationType != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ObservationType))
		i--
		dAtA[i] = 0x18
	}
	if m.ChainId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ChainId))
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

func (m *MsgAddObserverResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddObserverResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddObserverResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSetSupportedChains) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ChainId != 0 {
		n += 1 + sovTx(uint64(m.ChainId))
	}
	if m.ChainName != 0 {
		n += 1 + sovTx(uint64(m.ChainName))
	}
	return n
}

func (m *MsgSetSupportedChainsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgAddObserver) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ChainId != 0 {
		n += 1 + sovTx(uint64(m.ChainId))
	}
	if m.ObservationType != 0 {
		n += 1 + sovTx(uint64(m.ObservationType))
	}
	return n
}

func (m *MsgAddObserverResponse) Size() (n int) {
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
func (m *MsgSetSupportedChains) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetSupportedChains: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetSupportedChains: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			m.ChainName = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainName |= common.ChainName(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgSetSupportedChainsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetSupportedChainsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetSupportedChainsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgAddObserver) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddObserver: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddObserver: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return fmt.Errorf("proto: wrong wireType = %d for field ObservationType", wireType)
			}
			m.ObservationType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ObservationType |= ObservationType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgAddObserverResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddObserverResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddObserverResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
