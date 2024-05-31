// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/authority/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgUpdatePolicies defines the MsgUpdatePolicies service.
type MsgUpdatePolicies struct {
	Creator  string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Policies Policies `protobuf:"bytes,2,opt,name=policies,proto3" json:"policies"`
}

func (m *MsgUpdatePolicies) Reset()         { *m = MsgUpdatePolicies{} }
func (m *MsgUpdatePolicies) String() string { return proto.CompactTextString(m) }
func (*MsgUpdatePolicies) ProtoMessage()    {}
func (*MsgUpdatePolicies) Descriptor() ([]byte, []int) {
	return fileDescriptor_42e081863c477116, []int{0}
}
func (m *MsgUpdatePolicies) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdatePolicies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdatePolicies.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdatePolicies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdatePolicies.Merge(m, src)
}
func (m *MsgUpdatePolicies) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdatePolicies) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdatePolicies.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdatePolicies proto.InternalMessageInfo

func (m *MsgUpdatePolicies) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdatePolicies) GetPolicies() Policies {
	if m != nil {
		return m.Policies
	}
	return Policies{}
}

// MsgUpdatePoliciesResponse defines the MsgUpdatePoliciesResponse service.
type MsgUpdatePoliciesResponse struct {
}

func (m *MsgUpdatePoliciesResponse) Reset()         { *m = MsgUpdatePoliciesResponse{} }
func (m *MsgUpdatePoliciesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdatePoliciesResponse) ProtoMessage()    {}
func (*MsgUpdatePoliciesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42e081863c477116, []int{1}
}
func (m *MsgUpdatePoliciesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdatePoliciesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdatePoliciesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdatePoliciesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdatePoliciesResponse.Merge(m, src)
}
func (m *MsgUpdatePoliciesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdatePoliciesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdatePoliciesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdatePoliciesResponse proto.InternalMessageInfo

// MsgUpdateChainInfo defines the MsgUpdateChainInfo service.
type MsgUpdateChainInfo struct {
	Creator   string    `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ChainInfo ChainInfo `protobuf:"bytes,2,opt,name=chain_info,json=chainInfo,proto3" json:"chain_info"`
}

func (m *MsgUpdateChainInfo) Reset()         { *m = MsgUpdateChainInfo{} }
func (m *MsgUpdateChainInfo) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateChainInfo) ProtoMessage()    {}
func (*MsgUpdateChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_42e081863c477116, []int{2}
}
func (m *MsgUpdateChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateChainInfo.Merge(m, src)
}
func (m *MsgUpdateChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateChainInfo proto.InternalMessageInfo

func (m *MsgUpdateChainInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateChainInfo) GetChainInfo() ChainInfo {
	if m != nil {
		return m.ChainInfo
	}
	return ChainInfo{}
}

// MsgUpdateChainInfoResponse defines the MsgUpdateChainInfoResponse service.
type MsgUpdateChainInfoResponse struct {
}

func (m *MsgUpdateChainInfoResponse) Reset()         { *m = MsgUpdateChainInfoResponse{} }
func (m *MsgUpdateChainInfoResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateChainInfoResponse) ProtoMessage()    {}
func (*MsgUpdateChainInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42e081863c477116, []int{3}
}
func (m *MsgUpdateChainInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateChainInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateChainInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateChainInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateChainInfoResponse.Merge(m, src)
}
func (m *MsgUpdateChainInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateChainInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateChainInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateChainInfoResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdatePolicies)(nil), "zetachain.zetacore.authority.MsgUpdatePolicies")
	proto.RegisterType((*MsgUpdatePoliciesResponse)(nil), "zetachain.zetacore.authority.MsgUpdatePoliciesResponse")
	proto.RegisterType((*MsgUpdateChainInfo)(nil), "zetachain.zetacore.authority.MsgUpdateChainInfo")
	proto.RegisterType((*MsgUpdateChainInfoResponse)(nil), "zetachain.zetacore.authority.MsgUpdateChainInfoResponse")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/authority/tx.proto", fileDescriptor_42e081863c477116)
}

var fileDescriptor_42e081863c477116 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xad, 0x4a, 0x2d, 0x49,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0xb3, 0xf2, 0x8b, 0x52, 0xf5, 0x13, 0x4b, 0x4b, 0x32,
	0xf2, 0x8b, 0x32, 0x4b, 0x2a, 0xf5, 0x4b, 0x2a, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x64,
	0xe0, 0xca, 0xf4, 0x60, 0xca, 0xf4, 0xe0, 0xca, 0xa4, 0xb4, 0xf1, 0x1a, 0x52, 0x90, 0x9f, 0x93,
	0x99, 0x9c, 0x99, 0x5a, 0x0c, 0x31, 0x4a, 0x4a, 0x17, 0xaf, 0x62, 0xb0, 0x44, 0x7c, 0x66, 0x5e,
	0x5a, 0x3e, 0x54, 0xb9, 0x01, 0x5e, 0xe5, 0x50, 0x56, 0x55, 0x62, 0x49, 0x66, 0x7e, 0x1e, 0x54,
	0x87, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x98, 0xa9, 0x0f, 0x62, 0x41, 0x44, 0x95, 0xca, 0xb9, 0x04,
	0x7d, 0x8b, 0xd3, 0x43, 0x0b, 0x52, 0x12, 0x4b, 0x52, 0x03, 0xa0, 0x2e, 0x12, 0x92, 0xe0, 0x62,
	0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71,
	0x85, 0x3c, 0xb8, 0x38, 0x60, 0xee, 0x96, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x52, 0xd3, 0xc3,
	0x17, 0x06, 0x7a, 0x30, 0x33, 0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x82, 0xeb, 0x56, 0x92,
	0xe6, 0x92, 0xc4, 0xb0, 0x38, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xa9, 0x86, 0x4b,
	0x08, 0x2e, 0xe9, 0x0c, 0x32, 0xda, 0x33, 0x2f, 0x2d, 0x1f, 0x8f, 0xb3, 0x7c, 0xb8, 0xb8, 0x10,
	0x21, 0x04, 0x75, 0x98, 0x3a, 0x7e, 0x87, 0xc1, 0x8d, 0x85, 0xba, 0x8c, 0x33, 0x19, 0x26, 0xa0,
	0x24, 0xc3, 0x25, 0x85, 0x69, 0x3b, 0xcc, 0x6d, 0x46, 0x0d, 0x4c, 0x5c, 0xcc, 0xbe, 0xc5, 0xe9,
	0x42, 0x55, 0x5c, 0x7c, 0x68, 0xc1, 0xa6, 0x8f, 0xdf, 0x46, 0x0c, 0xef, 0x4a, 0x99, 0x93, 0xa8,
	0x01, 0xe6, 0x06, 0xa1, 0x5a, 0x2e, 0x7e, 0xf4, 0xc0, 0x31, 0x20, 0xd2, 0x2c, 0xb8, 0x0e, 0x29,
	0x0b, 0x52, 0x75, 0xc0, 0xac, 0x77, 0xf2, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0x83, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x70, 0xba, 0xd4,
	0x45, 0x4b, 0xa2, 0x15, 0xc8, 0xb9, 0xa8, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x9c, 0x0e, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xeb, 0xa6, 0x40, 0x72, 0x03, 0x00, 0x00,
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
	UpdatePolicies(ctx context.Context, in *MsgUpdatePolicies, opts ...grpc.CallOption) (*MsgUpdatePoliciesResponse, error)
	UpdateChainInfo(ctx context.Context, in *MsgUpdateChainInfo, opts ...grpc.CallOption) (*MsgUpdateChainInfoResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdatePolicies(ctx context.Context, in *MsgUpdatePolicies, opts ...grpc.CallOption) (*MsgUpdatePoliciesResponse, error) {
	out := new(MsgUpdatePoliciesResponse)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.authority.Msg/UpdatePolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateChainInfo(ctx context.Context, in *MsgUpdateChainInfo, opts ...grpc.CallOption) (*MsgUpdateChainInfoResponse, error) {
	out := new(MsgUpdateChainInfoResponse)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.authority.Msg/UpdateChainInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	UpdatePolicies(context.Context, *MsgUpdatePolicies) (*MsgUpdatePoliciesResponse, error)
	UpdateChainInfo(context.Context, *MsgUpdateChainInfo) (*MsgUpdateChainInfoResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdatePolicies(ctx context.Context, req *MsgUpdatePolicies) (*MsgUpdatePoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePolicies not implemented")
}
func (*UnimplementedMsgServer) UpdateChainInfo(ctx context.Context, req *MsgUpdateChainInfo) (*MsgUpdateChainInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChainInfo not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdatePolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePolicies)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.authority.Msg/UpdatePolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePolicies(ctx, req.(*MsgUpdatePolicies))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateChainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateChainInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateChainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.authority.Msg/UpdateChainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateChainInfo(ctx, req.(*MsgUpdateChainInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zetachain.zetacore.authority.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdatePolicies",
			Handler:    _Msg_UpdatePolicies_Handler,
		},
		{
			MethodName: "UpdateChainInfo",
			Handler:    _Msg_UpdateChainInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zetachain/zetacore/authority/tx.proto",
}

func (m *MsgUpdatePolicies) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdatePolicies) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdatePolicies) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Policies.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdatePoliciesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdatePoliciesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdatePoliciesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ChainInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateChainInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateChainInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateChainInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgUpdatePolicies) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Policies.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdatePoliciesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.ChainInfo.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateChainInfoResponse) Size() (n int) {
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
func (m *MsgUpdatePolicies) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdatePolicies: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdatePolicies: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Policies.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdatePoliciesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdatePoliciesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdatePoliciesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateChainInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChainInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdateChainInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateChainInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateChainInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
