// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/observer/ballot.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
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

type VoteType int32

const (
	VoteType_SuccessObservation VoteType = 0
	VoteType_FailureObservation VoteType = 1
	// this voter is observing failed / reverted . It does
	// not mean it was unable to observe.
	VoteType_NotYetVoted VoteType = 2
)

var VoteType_name = map[int32]string{
	0: "SuccessObservation",
	1: "FailureObservation",
	2: "NotYetVoted",
}

var VoteType_value = map[string]int32{
	"SuccessObservation": 0,
	"FailureObservation": 1,
	"NotYetVoted":        2,
}

func (x VoteType) String() string {
	return proto.EnumName(VoteType_name, int32(x))
}

func (VoteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_18c7141b763f2e87, []int{0}
}

type BallotStatus int32

const (
	BallotStatus_BallotFinalized_SuccessObservation BallotStatus = 0
	BallotStatus_BallotFinalized_FailureObservation BallotStatus = 1
	BallotStatus_BallotInProgress                   BallotStatus = 2
)

var BallotStatus_name = map[int32]string{
	0: "BallotFinalized_SuccessObservation",
	1: "BallotFinalized_FailureObservation",
	2: "BallotInProgress",
}

var BallotStatus_value = map[string]int32{
	"BallotFinalized_SuccessObservation": 0,
	"BallotFinalized_FailureObservation": 1,
	"BallotInProgress":                   2,
}

func (x BallotStatus) String() string {
	return proto.EnumName(BallotStatus_name, int32(x))
}

func (BallotStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_18c7141b763f2e87, []int{1}
}

// https://github.com/zeta-chain/node/issues/939
type Ballot struct {
	Index                string                      `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	BallotIdentifier     string                      `protobuf:"bytes,2,opt,name=ballot_identifier,json=ballotIdentifier,proto3" json:"ballot_identifier,omitempty"`
	VoterList            []string                    `protobuf:"bytes,3,rep,name=voter_list,json=voterList,proto3" json:"voter_list,omitempty"`
	Votes                []VoteType                  `protobuf:"varint,4,rep,packed,name=votes,proto3,enum=zetachain.zetacore.observer.VoteType" json:"votes,omitempty"`
	ObservationType      ObservationType             `protobuf:"varint,5,opt,name=observation_type,json=observationType,proto3,enum=zetachain.zetacore.observer.ObservationType" json:"observation_type,omitempty"`
	BallotThreshold      cosmossdk_io_math.LegacyDec `protobuf:"bytes,6,opt,name=ballot_threshold,json=ballotThreshold,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"ballot_threshold"`
	BallotStatus         BallotStatus                `protobuf:"varint,7,opt,name=ballot_status,json=ballotStatus,proto3,enum=zetachain.zetacore.observer.BallotStatus" json:"ballot_status,omitempty"`
	BallotCreationHeight int64                       `protobuf:"varint,8,opt,name=ballot_creation_height,json=ballotCreationHeight,proto3" json:"ballot_creation_height,omitempty"`
}

func (m *Ballot) Reset()         { *m = Ballot{} }
func (m *Ballot) String() string { return proto.CompactTextString(m) }
func (*Ballot) ProtoMessage()    {}
func (*Ballot) Descriptor() ([]byte, []int) {
	return fileDescriptor_18c7141b763f2e87, []int{0}
}
func (m *Ballot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ballot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ballot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ballot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ballot.Merge(m, src)
}
func (m *Ballot) XXX_Size() int {
	return m.Size()
}
func (m *Ballot) XXX_DiscardUnknown() {
	xxx_messageInfo_Ballot.DiscardUnknown(m)
}

var xxx_messageInfo_Ballot proto.InternalMessageInfo

func (m *Ballot) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Ballot) GetBallotIdentifier() string {
	if m != nil {
		return m.BallotIdentifier
	}
	return ""
}

func (m *Ballot) GetVoterList() []string {
	if m != nil {
		return m.VoterList
	}
	return nil
}

func (m *Ballot) GetVotes() []VoteType {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *Ballot) GetObservationType() ObservationType {
	if m != nil {
		return m.ObservationType
	}
	return ObservationType_EmptyObserverType
}

func (m *Ballot) GetBallotStatus() BallotStatus {
	if m != nil {
		return m.BallotStatus
	}
	return BallotStatus_BallotFinalized_SuccessObservation
}

func (m *Ballot) GetBallotCreationHeight() int64 {
	if m != nil {
		return m.BallotCreationHeight
	}
	return 0
}

type BallotListForHeight struct {
	Height           int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	BallotsIndexList []string `protobuf:"bytes,2,rep,name=ballots_index_list,json=ballotsIndexList,proto3" json:"ballots_index_list,omitempty"`
}

func (m *BallotListForHeight) Reset()         { *m = BallotListForHeight{} }
func (m *BallotListForHeight) String() string { return proto.CompactTextString(m) }
func (*BallotListForHeight) ProtoMessage()    {}
func (*BallotListForHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_18c7141b763f2e87, []int{1}
}
func (m *BallotListForHeight) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BallotListForHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BallotListForHeight.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BallotListForHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BallotListForHeight.Merge(m, src)
}
func (m *BallotListForHeight) XXX_Size() int {
	return m.Size()
}
func (m *BallotListForHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_BallotListForHeight.DiscardUnknown(m)
}

var xxx_messageInfo_BallotListForHeight proto.InternalMessageInfo

func (m *BallotListForHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BallotListForHeight) GetBallotsIndexList() []string {
	if m != nil {
		return m.BallotsIndexList
	}
	return nil
}

func init() {
	proto.RegisterEnum("zetachain.zetacore.observer.VoteType", VoteType_name, VoteType_value)
	proto.RegisterEnum("zetachain.zetacore.observer.BallotStatus", BallotStatus_name, BallotStatus_value)
	proto.RegisterType((*Ballot)(nil), "zetachain.zetacore.observer.Ballot")
	proto.RegisterType((*BallotListForHeight)(nil), "zetachain.zetacore.observer.BallotListForHeight")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/observer/ballot.proto", fileDescriptor_18c7141b763f2e87)
}

var fileDescriptor_18c7141b763f2e87 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x93, 0x26, 0x34, 0x4b, 0x69, 0xc2, 0x12, 0x45, 0x56, 0x2a, 0xdc, 0x28, 0x08, 0x64,
	0xd2, 0x62, 0x4b, 0x85, 0x1b, 0xb7, 0x00, 0x11, 0x91, 0xaa, 0x00, 0x6e, 0x05, 0x02, 0x0e, 0x96,
	0x63, 0x0f, 0xf6, 0x0a, 0xc7, 0x1b, 0xed, 0x6e, 0xaa, 0x26, 0x4f, 0xc1, 0x43, 0x70, 0xe0, 0x51,
	0x7a, 0xec, 0x09, 0x21, 0x0e, 0x15, 0x4a, 0x5e, 0x04, 0x79, 0xd7, 0x6e, 0x53, 0x29, 0xca, 0x6d,
	0x67, 0xe6, 0xfb, 0xbe, 0xf9, 0xdb, 0x41, 0xe6, 0x1c, 0x84, 0xe7, 0x47, 0x1e, 0x49, 0x6c, 0xf9,
	0xa2, 0x0c, 0x6c, 0x3a, 0xe2, 0xc0, 0xce, 0x80, 0xd9, 0x23, 0x2f, 0x8e, 0xa9, 0xb0, 0x26, 0x8c,
	0x0a, 0x8a, 0xf7, 0xae, 0x91, 0x56, 0x8e, 0xb4, 0x72, 0x64, 0xab, 0x11, 0xd2, 0x90, 0x4a, 0x9c,
	0x9d, 0xbe, 0x14, 0xa5, 0xd5, 0xdd, 0x24, 0x9e, 0x3f, 0x14, 0xb6, 0xf3, 0xbb, 0x84, 0x2a, 0x3d,
	0x99, 0x0f, 0x37, 0x50, 0x99, 0x24, 0x01, 0x9c, 0xeb, 0x5a, 0x5b, 0x33, 0xab, 0x8e, 0x32, 0xf0,
	0x01, 0xba, 0xaf, 0xea, 0x71, 0x49, 0x00, 0x89, 0x20, 0xdf, 0x08, 0x30, 0xbd, 0x28, 0x11, 0x75,
	0x15, 0x18, 0x5c, 0xfb, 0xf1, 0x43, 0x84, 0xce, 0xa8, 0x00, 0xe6, 0xc6, 0x84, 0x0b, 0xbd, 0xd4,
	0x2e, 0x99, 0x55, 0xa7, 0x2a, 0x3d, 0xc7, 0x84, 0x0b, 0xfc, 0x12, 0x95, 0x53, 0x83, 0xeb, 0x5b,
	0xed, 0x92, 0xb9, 0x7b, 0xf4, 0xd8, 0xda, 0xd0, 0x9b, 0xf5, 0x91, 0x0a, 0x38, 0x9d, 0x4d, 0xc0,
	0x51, 0x1c, 0xfc, 0x09, 0xd5, 0x55, 0xcc, 0x13, 0x84, 0x26, 0xae, 0x98, 0x4d, 0x40, 0x2f, 0xb7,
	0x35, 0x73, 0xf7, 0xe8, 0x70, 0xa3, 0xce, 0xbb, 0x1b, 0x92, 0x94, 0xab, 0xd1, 0xdb, 0x0e, 0x3c,
	0x44, 0x59, 0x23, 0xae, 0x88, 0x18, 0xf0, 0x88, 0xc6, 0x81, 0x5e, 0x49, 0x1b, 0xec, 0x3d, 0xba,
	0xb8, 0xda, 0x2f, 0xfc, 0xbd, 0xda, 0xdf, 0xf3, 0x29, 0x1f, 0x53, 0xce, 0x83, 0xef, 0x16, 0xa1,
	0xf6, 0xd8, 0x13, 0x91, 0x75, 0x0c, 0xa1, 0xe7, 0xcf, 0x5e, 0x83, 0xef, 0xd4, 0x14, 0xf9, 0x34,
	0xe7, 0xe2, 0x21, 0xba, 0x97, 0xe9, 0x71, 0xe1, 0x89, 0x29, 0xd7, 0xef, 0xc8, 0x2a, 0x9f, 0x6e,
	0xac, 0x52, 0xed, 0xe0, 0x44, 0x12, 0x9c, 0x9d, 0xd1, 0x8a, 0x85, 0x5f, 0xa0, 0x66, 0xa6, 0xe7,
	0x33, 0x50, 0xcd, 0x47, 0x40, 0xc2, 0x48, 0xe8, 0xdb, 0x6d, 0xcd, 0x2c, 0x39, 0x0d, 0x15, 0x7d,
	0x95, 0x05, 0xdf, 0xca, 0x58, 0xe7, 0x2b, 0x7a, 0xa0, 0x34, 0xd3, 0xc9, 0xf7, 0x29, 0x53, 0x6e,
	0xdc, 0x44, 0x95, 0x8c, 0xac, 0x49, 0x72, 0x66, 0xe1, 0x43, 0x84, 0x95, 0x0c, 0x77, 0xe5, 0xde,
	0xd5, 0x06, 0x8b, 0x72, 0x83, 0xd9, 0x78, 0xf8, 0x20, 0x0d, 0xa4, 0x72, 0xdd, 0x0f, 0x68, 0x3b,
	0x5f, 0x0f, 0x6e, 0x22, 0x7c, 0x32, 0xf5, 0x7d, 0xe0, 0x7c, 0x65, 0xd2, 0xf5, 0x42, 0xea, 0xef,
	0x7b, 0x24, 0x9e, 0x32, 0x58, 0xf5, 0x6b, 0xb8, 0x86, 0xee, 0x0e, 0xa9, 0xf8, 0x0c, 0x22, 0x55,
	0x08, 0xea, 0xc5, 0xd6, 0xd6, 0xaf, 0x9f, 0x86, 0xd6, 0x9d, 0xa3, 0x9d, 0xd5, 0x19, 0xe0, 0x27,
	0xa8, 0xa3, 0xec, 0x3e, 0x49, 0xbc, 0x98, 0xcc, 0x21, 0x70, 0xd7, 0xa6, 0x59, 0x83, 0x5b, 0x9b,
	0xb6, 0x81, 0xea, 0x0a, 0x37, 0x48, 0xde, 0x33, 0x1a, 0x32, 0xe0, 0x3c, 0xcf, 0xdd, 0x7b, 0x73,
	0xb1, 0x30, 0xb4, 0xcb, 0x85, 0xa1, 0xfd, 0x5b, 0x18, 0xda, 0x8f, 0xa5, 0x51, 0xb8, 0x5c, 0x1a,
	0x85, 0x3f, 0x4b, 0xa3, 0xf0, 0xe5, 0x20, 0x24, 0x22, 0x9a, 0x8e, 0x2c, 0x9f, 0x8e, 0xe5, 0x2d,
	0x3d, 0x53, 0x67, 0x95, 0xd0, 0x00, 0xec, 0xf3, 0x9b, 0xa3, 0x4a, 0x3f, 0x23, 0x1f, 0x55, 0xe4,
	0x49, 0x3d, 0xff, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x59, 0x7b, 0x9c, 0x5c, 0xdd, 0x03, 0x00, 0x00,
}

func (m *Ballot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ballot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ballot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BallotCreationHeight != 0 {
		i = encodeVarintBallot(dAtA, i, uint64(m.BallotCreationHeight))
		i--
		dAtA[i] = 0x40
	}
	if m.BallotStatus != 0 {
		i = encodeVarintBallot(dAtA, i, uint64(m.BallotStatus))
		i--
		dAtA[i] = 0x38
	}
	{
		size := m.BallotThreshold.Size()
		i -= size
		if _, err := m.BallotThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBallot(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.ObservationType != 0 {
		i = encodeVarintBallot(dAtA, i, uint64(m.ObservationType))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Votes) > 0 {
		dAtA2 := make([]byte, len(m.Votes)*10)
		var j1 int
		for _, num := range m.Votes {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintBallot(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.VoterList) > 0 {
		for iNdEx := len(m.VoterList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.VoterList[iNdEx])
			copy(dAtA[i:], m.VoterList[iNdEx])
			i = encodeVarintBallot(dAtA, i, uint64(len(m.VoterList[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.BallotIdentifier) > 0 {
		i -= len(m.BallotIdentifier)
		copy(dAtA[i:], m.BallotIdentifier)
		i = encodeVarintBallot(dAtA, i, uint64(len(m.BallotIdentifier)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintBallot(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BallotListForHeight) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BallotListForHeight) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BallotListForHeight) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BallotsIndexList) > 0 {
		for iNdEx := len(m.BallotsIndexList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BallotsIndexList[iNdEx])
			copy(dAtA[i:], m.BallotsIndexList[iNdEx])
			i = encodeVarintBallot(dAtA, i, uint64(len(m.BallotsIndexList[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Height != 0 {
		i = encodeVarintBallot(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBallot(dAtA []byte, offset int, v uint64) int {
	offset -= sovBallot(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Ballot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovBallot(uint64(l))
	}
	l = len(m.BallotIdentifier)
	if l > 0 {
		n += 1 + l + sovBallot(uint64(l))
	}
	if len(m.VoterList) > 0 {
		for _, s := range m.VoterList {
			l = len(s)
			n += 1 + l + sovBallot(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		l = 0
		for _, e := range m.Votes {
			l += sovBallot(uint64(e))
		}
		n += 1 + sovBallot(uint64(l)) + l
	}
	if m.ObservationType != 0 {
		n += 1 + sovBallot(uint64(m.ObservationType))
	}
	l = m.BallotThreshold.Size()
	n += 1 + l + sovBallot(uint64(l))
	if m.BallotStatus != 0 {
		n += 1 + sovBallot(uint64(m.BallotStatus))
	}
	if m.BallotCreationHeight != 0 {
		n += 1 + sovBallot(uint64(m.BallotCreationHeight))
	}
	return n
}

func (m *BallotListForHeight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBallot(uint64(m.Height))
	}
	if len(m.BallotsIndexList) > 0 {
		for _, s := range m.BallotsIndexList {
			l = len(s)
			n += 1 + l + sovBallot(uint64(l))
		}
	}
	return n
}

func sovBallot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBallot(x uint64) (n int) {
	return sovBallot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ballot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBallot
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
			return fmt.Errorf("proto: Ballot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ballot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
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
				return ErrInvalidLengthBallot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBallot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
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
				return ErrInvalidLengthBallot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBallot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BallotIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoterList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
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
				return ErrInvalidLengthBallot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBallot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoterList = append(m.VoterList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v VoteType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBallot
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= VoteType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Votes = append(m.Votes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBallot
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBallot
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthBallot
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Votes) == 0 {
					m.Votes = make([]VoteType, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v VoteType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBallot
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= VoteType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Votes = append(m.Votes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObservationType", wireType)
			}
			m.ObservationType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
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
				return ErrInvalidLengthBallot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBallot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BallotThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotStatus", wireType)
			}
			m.BallotStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BallotStatus |= BallotStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotCreationHeight", wireType)
			}
			m.BallotCreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BallotCreationHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBallot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBallot
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
func (m *BallotListForHeight) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBallot
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
			return fmt.Errorf("proto: BallotListForHeight: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BallotListForHeight: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotsIndexList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
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
				return ErrInvalidLengthBallot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBallot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BallotsIndexList = append(m.BallotsIndexList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBallot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBallot
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
func skipBallot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBallot
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
					return 0, ErrIntOverflowBallot
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
					return 0, ErrIntOverflowBallot
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
				return 0, ErrInvalidLengthBallot
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBallot
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBallot
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBallot        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBallot          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBallot = fmt.Errorf("proto: unexpected end of group")
)
