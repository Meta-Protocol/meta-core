// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/emissions/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	ValidatorEmissionPercentage string                                 `protobuf:"bytes,5,opt,name=validator_emission_percentage,json=validatorEmissionPercentage,proto3" json:"validator_emission_percentage,omitempty"`
	ObserverEmissionPercentage  string                                 `protobuf:"bytes,6,opt,name=observer_emission_percentage,json=observerEmissionPercentage,proto3" json:"observer_emission_percentage,omitempty"`
	TssSignerEmissionPercentage string                                 `protobuf:"bytes,7,opt,name=tss_signer_emission_percentage,json=tssSignerEmissionPercentage,proto3" json:"tss_signer_emission_percentage,omitempty"`
	ObserverSlashAmount         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=observer_slash_amount,json=observerSlashAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"observer_slash_amount"`
	BallotMaturityBlocks        int64                                  `protobuf:"varint,10,opt,name=ballot_maturity_blocks,json=ballotMaturityBlocks,proto3" json:"ballot_maturity_blocks,omitempty"`
	BlockRewardAmount           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=block_reward_amount,json=blockRewardAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"block_reward_amount"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_259272924aec0acf, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetValidatorEmissionPercentage() string {
	if m != nil {
		return m.ValidatorEmissionPercentage
	}
	return ""
}

func (m *Params) GetObserverEmissionPercentage() string {
	if m != nil {
		return m.ObserverEmissionPercentage
	}
	return ""
}

func (m *Params) GetTssSignerEmissionPercentage() string {
	if m != nil {
		return m.TssSignerEmissionPercentage
	}
	return ""
}

func (m *Params) GetBallotMaturityBlocks() int64 {
	if m != nil {
		return m.BallotMaturityBlocks
	}
	return 0
}

// Deprecated (v20): Do not use. Use Params Instead
type LegacyParams struct {
	MaxBondFactor               string                                 `protobuf:"bytes,1,opt,name=max_bond_factor,json=maxBondFactor,proto3" json:"max_bond_factor,omitempty"`
	MinBondFactor               string                                 `protobuf:"bytes,2,opt,name=min_bond_factor,json=minBondFactor,proto3" json:"min_bond_factor,omitempty"`
	AvgBlockTime                string                                 `protobuf:"bytes,3,opt,name=avg_block_time,json=avgBlockTime,proto3" json:"avg_block_time,omitempty"`
	TargetBondRatio             string                                 `protobuf:"bytes,4,opt,name=target_bond_ratio,json=targetBondRatio,proto3" json:"target_bond_ratio,omitempty"`
	ValidatorEmissionPercentage string                                 `protobuf:"bytes,5,opt,name=validator_emission_percentage,json=validatorEmissionPercentage,proto3" json:"validator_emission_percentage,omitempty"`
	ObserverEmissionPercentage  string                                 `protobuf:"bytes,6,opt,name=observer_emission_percentage,json=observerEmissionPercentage,proto3" json:"observer_emission_percentage,omitempty"`
	TssSignerEmissionPercentage string                                 `protobuf:"bytes,7,opt,name=tss_signer_emission_percentage,json=tssSignerEmissionPercentage,proto3" json:"tss_signer_emission_percentage,omitempty"`
	DurationFactorConstant      string                                 `protobuf:"bytes,8,opt,name=duration_factor_constant,json=durationFactorConstant,proto3" json:"duration_factor_constant,omitempty"`
	ObserverSlashAmount         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=observer_slash_amount,json=observerSlashAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"observer_slash_amount"`
	BallotMaturityBlocks        int64                                  `protobuf:"varint,10,opt,name=ballot_maturity_blocks,json=ballotMaturityBlocks,proto3" json:"ballot_maturity_blocks,omitempty"`
}

func (m *LegacyParams) Reset()      { *m = LegacyParams{} }
func (*LegacyParams) ProtoMessage() {}
func (*LegacyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_259272924aec0acf, []int{1}
}
func (m *LegacyParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LegacyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LegacyParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LegacyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyParams.Merge(m, src)
}
func (m *LegacyParams) XXX_Size() int {
	return m.Size()
}
func (m *LegacyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyParams.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyParams proto.InternalMessageInfo

func (m *LegacyParams) GetMaxBondFactor() string {
	if m != nil {
		return m.MaxBondFactor
	}
	return ""
}

func (m *LegacyParams) GetMinBondFactor() string {
	if m != nil {
		return m.MinBondFactor
	}
	return ""
}

func (m *LegacyParams) GetAvgBlockTime() string {
	if m != nil {
		return m.AvgBlockTime
	}
	return ""
}

func (m *LegacyParams) GetTargetBondRatio() string {
	if m != nil {
		return m.TargetBondRatio
	}
	return ""
}

func (m *LegacyParams) GetValidatorEmissionPercentage() string {
	if m != nil {
		return m.ValidatorEmissionPercentage
	}
	return ""
}

func (m *LegacyParams) GetObserverEmissionPercentage() string {
	if m != nil {
		return m.ObserverEmissionPercentage
	}
	return ""
}

func (m *LegacyParams) GetTssSignerEmissionPercentage() string {
	if m != nil {
		return m.TssSignerEmissionPercentage
	}
	return ""
}

func (m *LegacyParams) GetDurationFactorConstant() string {
	if m != nil {
		return m.DurationFactorConstant
	}
	return ""
}

func (m *LegacyParams) GetBallotMaturityBlocks() int64 {
	if m != nil {
		return m.BallotMaturityBlocks
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "zetachain.zetacore.emissions.Params")
	proto.RegisterType((*LegacyParams)(nil), "zetachain.zetacore.emissions.LegacyParams")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/emissions/params.proto", fileDescriptor_259272924aec0acf)
}

var fileDescriptor_259272924aec0acf = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x1b, 0xda, 0x95, 0xd6, 0x0c, 0xb6, 0x65, 0x63, 0x8a, 0xc6, 0x48, 0xa7, 0x09, 0x4d,
	0x03, 0xb1, 0xe4, 0x00, 0x07, 0xc4, 0x09, 0x3a, 0x98, 0xc4, 0x04, 0xd2, 0x94, 0x71, 0xe2, 0x80,
	0xf5, 0x92, 0x98, 0xd4, 0x5a, 0x63, 0x57, 0xb6, 0x5b, 0x5a, 0xfe, 0x0a, 0x8e, 0x1c, 0xf9, 0x73,
	0x76, 0xdc, 0x05, 0x09, 0x71, 0x18, 0xa8, 0xfd, 0x47, 0x50, 0x5e, 0x7e, 0xa8, 0x48, 0xe5, 0xc0,
	0x0d, 0xed, 0x14, 0xcb, 0xef, 0xf3, 0x3e, 0xfa, 0xda, 0xb1, 0x1e, 0xb9, 0xff, 0x89, 0x19, 0x88,
	0x7a, 0xc0, 0x85, 0x8f, 0x2b, 0xa9, 0x98, 0xcf, 0x52, 0xae, 0x35, 0x97, 0x42, 0xfb, 0x03, 0x50,
	0x90, 0x6a, 0x6f, 0xa0, 0xa4, 0x91, 0xf6, 0x76, 0x85, 0x7a, 0x25, 0xea, 0x55, 0xe8, 0xd6, 0x46,
	0x22, 0x13, 0x89, 0xa0, 0x9f, 0xad, 0xf2, 0x9e, 0xdd, 0x9f, 0x75, 0xd2, 0x3c, 0x41, 0x89, 0xdd,
	0x25, 0x77, 0x47, 0xd0, 0xe7, 0x31, 0x18, 0xa9, 0x68, 0xd9, 0x47, 0x07, 0x4c, 0x45, 0x4c, 0x18,
	0x48, 0x98, 0xb3, 0xb4, 0x63, 0xed, 0xb7, 0x83, 0x3b, 0x15, 0xf4, 0xb2, 0x60, 0x4e, 0x2a, 0xc4,
	0x7e, 0x46, 0xb6, 0x65, 0xa8, 0x99, 0x1a, 0xb1, 0xc5, 0x8a, 0x26, 0x2a, 0xb6, 0x4a, 0x66, 0x81,
	0xe1, 0x90, 0xb8, 0x46, 0x6b, 0xaa, 0x79, 0x22, 0xfe, 0xe2, 0xb8, 0x9e, 0xc7, 0x30, 0x5a, 0x9f,
	0x22, 0xb4, 0x40, 0x12, 0x92, 0xdb, 0x55, 0x0c, 0xdd, 0x07, 0xdd, 0xa3, 0x90, 0xca, 0xa1, 0x30,
	0x4e, 0x3b, 0xeb, 0xed, 0x7a, 0xe7, 0x97, 0x9d, 0xda, 0x8f, 0xcb, 0xce, 0x5e, 0xc2, 0x4d, 0x6f,
	0x18, 0x7a, 0x91, 0x4c, 0xfd, 0x48, 0xea, 0x54, 0xea, 0xe2, 0x73, 0xa0, 0xe3, 0x33, 0xdf, 0x4c,
	0x06, 0x4c, 0x7b, 0xaf, 0x84, 0x09, 0xd6, 0x4b, 0xd9, 0x69, 0xe6, 0x7a, 0x8e, 0x2a, 0xfb, 0x31,
	0xd9, 0x0c, 0xa1, 0xdf, 0x97, 0x86, 0xa6, 0x60, 0x86, 0x8a, 0x9b, 0x09, 0x0d, 0xfb, 0x32, 0x3a,
	0xd3, 0x0e, 0xd9, 0xb1, 0xf6, 0xeb, 0xc1, 0x46, 0x5e, 0x7d, 0x53, 0x14, 0xbb, 0x58, 0xb3, 0xdf,
	0x93, 0x75, 0xa4, 0xa8, 0x62, 0x1f, 0x41, 0xc5, 0x65, 0xae, 0x1b, 0xff, 0x9c, 0xeb, 0x05, 0x8b,
	0x82, 0x35, 0x54, 0x05, 0x68, 0xca, 0x53, 0x3d, 0x6d, 0x7c, 0xf9, 0xda, 0xa9, 0x1d, 0x37, 0x5a,
	0xd6, 0xea, 0xd2, 0x71, 0xa3, 0xd5, 0x5a, 0x6d, 0xef, 0x7e, 0x6b, 0x90, 0xe5, 0xd7, 0x2c, 0x81,
	0x68, 0x52, 0xfc, 0xe7, 0x3d, 0xb2, 0x92, 0xc2, 0x98, 0x86, 0x52, 0xc4, 0xf4, 0x03, 0x44, 0x46,
	0x2a, 0xc7, 0xc2, 0x2b, 0xbd, 0x99, 0xc2, 0xb8, 0x2b, 0x45, 0x7c, 0x84, 0x9b, 0xc8, 0x71, 0xf1,
	0x07, 0x77, 0xad, 0xe0, 0xb8, 0x98, 0xe3, 0xee, 0x91, 0x5b, 0x30, 0x4a, 0xf2, 0xc3, 0x53, 0xc3,
	0x53, 0xe6, 0xd4, 0x11, 0x5b, 0x86, 0x51, 0x82, 0xa7, 0x7e, 0xcb, 0x53, 0x66, 0x3f, 0x20, 0x6b,
	0x06, 0x54, 0xc2, 0x4c, 0x2e, 0x54, 0x60, 0xb8, 0x74, 0x1a, 0x08, 0xae, 0xe4, 0x85, 0x4c, 0x19,
	0x64, 0xdb, 0x57, 0xe9, 0x25, 0x3e, 0x21, 0x4e, 0x3c, 0xc4, 0xc3, 0x8a, 0xe2, 0x12, 0x69, 0x24,
	0x85, 0x36, 0x20, 0x8c, 0xd3, 0xc2, 0xf6, 0xcd, 0xb2, 0x9e, 0x5f, 0xe7, 0x61, 0x51, 0xfd, 0x7f,
	0xdf, 0x70, 0xfe, 0xc6, 0xba, 0x47, 0xe7, 0x53, 0xd7, 0xba, 0x98, 0xba, 0xd6, 0xaf, 0xa9, 0x6b,
	0x7d, 0x9e, 0xb9, 0xb5, 0x8b, 0x99, 0x5b, 0xfb, 0x3e, 0x73, 0x6b, 0xef, 0x1e, 0xce, 0x45, 0xca,
	0x06, 0xd1, 0x41, 0x3e, 0xbe, 0x84, 0x8c, 0x99, 0x3f, 0x9e, 0x1b, 0x5e, 0x18, 0x2e, 0x6c, 0xe2,
	0x20, 0x7a, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xea, 0x36, 0x06, 0xa1, 0xe9, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BlockRewardAmount.Size()
		i -= size
		if _, err := m.BlockRewardAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if m.BallotMaturityBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BallotMaturityBlocks))
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.ObserverSlashAmount.Size()
		i -= size
		if _, err := m.ObserverSlashAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.TssSignerEmissionPercentage) > 0 {
		i -= len(m.TssSignerEmissionPercentage)
		copy(dAtA[i:], m.TssSignerEmissionPercentage)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TssSignerEmissionPercentage)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ObserverEmissionPercentage) > 0 {
		i -= len(m.ObserverEmissionPercentage)
		copy(dAtA[i:], m.ObserverEmissionPercentage)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ObserverEmissionPercentage)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ValidatorEmissionPercentage) > 0 {
		i -= len(m.ValidatorEmissionPercentage)
		copy(dAtA[i:], m.ValidatorEmissionPercentage)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ValidatorEmissionPercentage)))
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}

func (m *LegacyParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LegacyParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LegacyParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BallotMaturityBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BallotMaturityBlocks))
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.ObserverSlashAmount.Size()
		i -= size
		if _, err := m.ObserverSlashAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.DurationFactorConstant) > 0 {
		i -= len(m.DurationFactorConstant)
		copy(dAtA[i:], m.DurationFactorConstant)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DurationFactorConstant)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.TssSignerEmissionPercentage) > 0 {
		i -= len(m.TssSignerEmissionPercentage)
		copy(dAtA[i:], m.TssSignerEmissionPercentage)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TssSignerEmissionPercentage)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ObserverEmissionPercentage) > 0 {
		i -= len(m.ObserverEmissionPercentage)
		copy(dAtA[i:], m.ObserverEmissionPercentage)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ObserverEmissionPercentage)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ValidatorEmissionPercentage) > 0 {
		i -= len(m.ValidatorEmissionPercentage)
		copy(dAtA[i:], m.ValidatorEmissionPercentage)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ValidatorEmissionPercentage)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TargetBondRatio) > 0 {
		i -= len(m.TargetBondRatio)
		copy(dAtA[i:], m.TargetBondRatio)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TargetBondRatio)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AvgBlockTime) > 0 {
		i -= len(m.AvgBlockTime)
		copy(dAtA[i:], m.AvgBlockTime)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AvgBlockTime)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MinBondFactor) > 0 {
		i -= len(m.MinBondFactor)
		copy(dAtA[i:], m.MinBondFactor)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MinBondFactor)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MaxBondFactor) > 0 {
		i -= len(m.MaxBondFactor)
		copy(dAtA[i:], m.MaxBondFactor)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MaxBondFactor)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorEmissionPercentage)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ObserverEmissionPercentage)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TssSignerEmissionPercentage)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.ObserverSlashAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.BallotMaturityBlocks != 0 {
		n += 1 + sovParams(uint64(m.BallotMaturityBlocks))
	}
	l = m.BlockRewardAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *LegacyParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MaxBondFactor)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.MinBondFactor)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.AvgBlockTime)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TargetBondRatio)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ValidatorEmissionPercentage)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ObserverEmissionPercentage)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TssSignerEmissionPercentage)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DurationFactorConstant)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.ObserverSlashAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.BallotMaturityBlocks != 0 {
		n += 1 + sovParams(uint64(m.BallotMaturityBlocks))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorEmissionPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorEmissionPercentage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverEmissionPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObserverEmissionPercentage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TssSignerEmissionPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TssSignerEmissionPercentage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverSlashAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObserverSlashAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotMaturityBlocks", wireType)
			}
			m.BallotMaturityBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BallotMaturityBlocks |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockRewardAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockRewardAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *LegacyParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: LegacyParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LegacyParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBondFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxBondFactor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBondFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinBondFactor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvgBlockTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AvgBlockTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetBondRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetBondRatio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorEmissionPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorEmissionPercentage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverEmissionPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObserverEmissionPercentage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TssSignerEmissionPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TssSignerEmissionPercentage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationFactorConstant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DurationFactorConstant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverSlashAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObserverSlashAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotMaturityBlocks", wireType)
			}
			m.BallotMaturityBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BallotMaturityBlocks |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
