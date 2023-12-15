// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/zeta-chain/zetacore/common"

	mock "github.com/stretchr/testify/mock"

	observertypes "github.com/zeta-chain/zetacore/x/observer/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// CrosschainObserverKeeper is an autogenerated mock type for the CrosschainObserverKeeper type
type CrosschainObserverKeeper struct {
	mock.Mock
}

// AddBallotToList provides a mock function with given fields: ctx, ballot
func (_m *CrosschainObserverKeeper) AddBallotToList(ctx types.Context, ballot observertypes.Ballot) {
	_m.Called(ctx, ballot)
}

// AddVoteToBallot provides a mock function with given fields: ctx, ballot, address, observationType
func (_m *CrosschainObserverKeeper) AddVoteToBallot(ctx types.Context, ballot observertypes.Ballot, address string, observationType observertypes.VoteType) (observertypes.Ballot, error) {
	ret := _m.Called(ctx, ballot, address, observationType)

	if len(ret) == 0 {
		panic("no return value specified for AddVoteToBallot")
	}

	var r0 observertypes.Ballot
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, observertypes.Ballot, string, observertypes.VoteType) (observertypes.Ballot, error)); ok {
		return rf(ctx, ballot, address, observationType)
	}
	if rf, ok := ret.Get(0).(func(types.Context, observertypes.Ballot, string, observertypes.VoteType) observertypes.Ballot); ok {
		r0 = rf(ctx, ballot, address, observationType)
	} else {
		r0 = ret.Get(0).(observertypes.Ballot)
	}

	if rf, ok := ret.Get(1).(func(types.Context, observertypes.Ballot, string, observertypes.VoteType) error); ok {
		r1 = rf(ctx, ballot, address, observationType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckIfFinalizingVote provides a mock function with given fields: ctx, ballot
func (_m *CrosschainObserverKeeper) CheckIfFinalizingVote(ctx types.Context, ballot observertypes.Ballot) (observertypes.Ballot, bool) {
	ret := _m.Called(ctx, ballot)

	if len(ret) == 0 {
		panic("no return value specified for CheckIfFinalizingVote")
	}

	var r0 observertypes.Ballot
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, observertypes.Ballot) (observertypes.Ballot, bool)); ok {
		return rf(ctx, ballot)
	}
	if rf, ok := ret.Get(0).(func(types.Context, observertypes.Ballot) observertypes.Ballot); ok {
		r0 = rf(ctx, ballot)
	} else {
		r0 = ret.Get(0).(observertypes.Ballot)
	}

	if rf, ok := ret.Get(1).(func(types.Context, observertypes.Ballot) bool); ok {
		r1 = rf(ctx, ballot)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// CheckIfTssPubkeyHasBeenGenerated provides a mock function with given fields: ctx, tssPubkey
func (_m *CrosschainObserverKeeper) CheckIfTssPubkeyHasBeenGenerated(ctx types.Context, tssPubkey string) (observertypes.TSS, bool) {
	ret := _m.Called(ctx, tssPubkey)

	if len(ret) == 0 {
		panic("no return value specified for CheckIfTssPubkeyHasBeenGenerated")
	}

	var r0 observertypes.TSS
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, string) (observertypes.TSS, bool)); ok {
		return rf(ctx, tssPubkey)
	}
	if rf, ok := ret.Get(0).(func(types.Context, string) observertypes.TSS); ok {
		r0 = rf(ctx, tssPubkey)
	} else {
		r0 = ret.Get(0).(observertypes.TSS)
	}

	if rf, ok := ret.Get(1).(func(types.Context, string) bool); ok {
		r1 = rf(ctx, tssPubkey)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// FindBallot provides a mock function with given fields: ctx, index, chain, observationType
func (_m *CrosschainObserverKeeper) FindBallot(ctx types.Context, index string, chain *common.Chain, observationType observertypes.ObservationType) (observertypes.Ballot, bool, error) {
	ret := _m.Called(ctx, index, chain, observationType)

	if len(ret) == 0 {
		panic("no return value specified for FindBallot")
	}

	var r0 observertypes.Ballot
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Context, string, *common.Chain, observertypes.ObservationType) (observertypes.Ballot, bool, error)); ok {
		return rf(ctx, index, chain, observationType)
	}
	if rf, ok := ret.Get(0).(func(types.Context, string, *common.Chain, observertypes.ObservationType) observertypes.Ballot); ok {
		r0 = rf(ctx, index, chain, observationType)
	} else {
		r0 = ret.Get(0).(observertypes.Ballot)
	}

	if rf, ok := ret.Get(1).(func(types.Context, string, *common.Chain, observertypes.ObservationType) bool); ok {
		r1 = rf(ctx, index, chain, observationType)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(types.Context, string, *common.Chain, observertypes.ObservationType) error); ok {
		r2 = rf(ctx, index, chain, observationType)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllBallots provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetAllBallots(ctx types.Context) []*observertypes.Ballot {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllBallots")
	}

	var r0 []*observertypes.Ballot
	if rf, ok := ret.Get(0).(func(types.Context) []*observertypes.Ballot); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*observertypes.Ballot)
		}
	}

	return r0
}

// GetAllChainNonces provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetAllChainNonces(ctx types.Context) []observertypes.ChainNonces {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllChainNonces")
	}

	var r0 []observertypes.ChainNonces
	if rf, ok := ret.Get(0).(func(types.Context) []observertypes.ChainNonces); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]observertypes.ChainNonces)
		}
	}

	return r0
}

// GetAllNodeAccount provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetAllNodeAccount(ctx types.Context) []observertypes.NodeAccount {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllNodeAccount")
	}

	var r0 []observertypes.NodeAccount
	if rf, ok := ret.Get(0).(func(types.Context) []observertypes.NodeAccount); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]observertypes.NodeAccount)
		}
	}

	return r0
}

// GetAllNonceToCctx provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetAllNonceToCctx(ctx types.Context) []observertypes.NonceToCctx {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllNonceToCctx")
	}

	var r0 []observertypes.NonceToCctx
	if rf, ok := ret.Get(0).(func(types.Context) []observertypes.NonceToCctx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]observertypes.NonceToCctx)
		}
	}

	return r0
}

// GetAllPendingNonces provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetAllPendingNonces(ctx types.Context) ([]observertypes.PendingNonces, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllPendingNonces")
	}

	var r0 []observertypes.PendingNonces
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context) ([]observertypes.PendingNonces, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) []observertypes.PendingNonces); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]observertypes.PendingNonces)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTSS provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetAllTSS(ctx types.Context) []observertypes.TSS {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTSS")
	}

	var r0 []observertypes.TSS
	if rf, ok := ret.Get(0).(func(types.Context) []observertypes.TSS); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]observertypes.TSS)
		}
	}

	return r0
}

// GetAllTssFundMigrators provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetAllTssFundMigrators(ctx types.Context) []observertypes.TssFundMigratorInfo {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTssFundMigrators")
	}

	var r0 []observertypes.TssFundMigratorInfo
	if rf, ok := ret.Get(0).(func(types.Context) []observertypes.TssFundMigratorInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]observertypes.TssFundMigratorInfo)
		}
	}

	return r0
}

// GetBallot provides a mock function with given fields: ctx, index
func (_m *CrosschainObserverKeeper) GetBallot(ctx types.Context, index string) (observertypes.Ballot, bool) {
	ret := _m.Called(ctx, index)

	if len(ret) == 0 {
		panic("no return value specified for GetBallot")
	}

	var r0 observertypes.Ballot
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, string) (observertypes.Ballot, bool)); ok {
		return rf(ctx, index)
	}
	if rf, ok := ret.Get(0).(func(types.Context, string) observertypes.Ballot); ok {
		r0 = rf(ctx, index)
	} else {
		r0 = ret.Get(0).(observertypes.Ballot)
	}

	if rf, ok := ret.Get(1).(func(types.Context, string) bool); ok {
		r1 = rf(ctx, index)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetBlockHeader provides a mock function with given fields: ctx, hash
func (_m *CrosschainObserverKeeper) GetBlockHeader(ctx types.Context, hash []byte) (common.BlockHeader, bool) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockHeader")
	}

	var r0 common.BlockHeader
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, []byte) (common.BlockHeader, bool)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(types.Context, []byte) common.BlockHeader); ok {
		r0 = rf(ctx, hash)
	} else {
		r0 = ret.Get(0).(common.BlockHeader)
	}

	if rf, ok := ret.Get(1).(func(types.Context, []byte) bool); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetChainNonces provides a mock function with given fields: ctx, index
func (_m *CrosschainObserverKeeper) GetChainNonces(ctx types.Context, index string) (observertypes.ChainNonces, bool) {
	ret := _m.Called(ctx, index)

	if len(ret) == 0 {
		panic("no return value specified for GetChainNonces")
	}

	var r0 observertypes.ChainNonces
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, string) (observertypes.ChainNonces, bool)); ok {
		return rf(ctx, index)
	}
	if rf, ok := ret.Get(0).(func(types.Context, string) observertypes.ChainNonces); ok {
		r0 = rf(ctx, index)
	} else {
		r0 = ret.Get(0).(observertypes.ChainNonces)
	}

	if rf, ok := ret.Get(1).(func(types.Context, string) bool); ok {
		r1 = rf(ctx, index)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetCoreParamsByChainID provides a mock function with given fields: ctx, chainID
func (_m *CrosschainObserverKeeper) GetCoreParamsByChainID(ctx types.Context, chainID int64) (*observertypes.CoreParams, bool) {
	ret := _m.Called(ctx, chainID)

	if len(ret) == 0 {
		panic("no return value specified for GetCoreParamsByChainID")
	}

	var r0 *observertypes.CoreParams
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, int64) (*observertypes.CoreParams, bool)); ok {
		return rf(ctx, chainID)
	}
	if rf, ok := ret.Get(0).(func(types.Context, int64) *observertypes.CoreParams); ok {
		r0 = rf(ctx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*observertypes.CoreParams)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, int64) bool); ok {
		r1 = rf(ctx, chainID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetCrosschainFlags provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetCrosschainFlags(ctx types.Context) (observertypes.CrosschainFlags, bool) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetCrosschainFlags")
	}

	var r0 observertypes.CrosschainFlags
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context) (observertypes.CrosschainFlags, bool)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) observertypes.CrosschainFlags); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(observertypes.CrosschainFlags)
	}

	if rf, ok := ret.Get(1).(func(types.Context) bool); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetFundMigrator provides a mock function with given fields: ctx, chainID
func (_m *CrosschainObserverKeeper) GetFundMigrator(ctx types.Context, chainID int64) (observertypes.TssFundMigratorInfo, bool) {
	ret := _m.Called(ctx, chainID)

	if len(ret) == 0 {
		panic("no return value specified for GetFundMigrator")
	}

	var r0 observertypes.TssFundMigratorInfo
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, int64) (observertypes.TssFundMigratorInfo, bool)); ok {
		return rf(ctx, chainID)
	}
	if rf, ok := ret.Get(0).(func(types.Context, int64) observertypes.TssFundMigratorInfo); ok {
		r0 = rf(ctx, chainID)
	} else {
		r0 = ret.Get(0).(observertypes.TssFundMigratorInfo)
	}

	if rf, ok := ret.Get(1).(func(types.Context, int64) bool); ok {
		r1 = rf(ctx, chainID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetKeygen provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetKeygen(ctx types.Context) (observertypes.Keygen, bool) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetKeygen")
	}

	var r0 observertypes.Keygen
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context) (observertypes.Keygen, bool)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) observertypes.Keygen); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(observertypes.Keygen)
	}

	if rf, ok := ret.Get(1).(func(types.Context) bool); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetNodeAccount provides a mock function with given fields: ctx, address
func (_m *CrosschainObserverKeeper) GetNodeAccount(ctx types.Context, address string) (observertypes.NodeAccount, bool) {
	ret := _m.Called(ctx, address)

	if len(ret) == 0 {
		panic("no return value specified for GetNodeAccount")
	}

	var r0 observertypes.NodeAccount
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, string) (observertypes.NodeAccount, bool)); ok {
		return rf(ctx, address)
	}
	if rf, ok := ret.Get(0).(func(types.Context, string) observertypes.NodeAccount); ok {
		r0 = rf(ctx, address)
	} else {
		r0 = ret.Get(0).(observertypes.NodeAccount)
	}

	if rf, ok := ret.Get(1).(func(types.Context, string) bool); ok {
		r1 = rf(ctx, address)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetNonceToCctx provides a mock function with given fields: ctx, tss, chainID, nonce
func (_m *CrosschainObserverKeeper) GetNonceToCctx(ctx types.Context, tss string, chainID int64, nonce int64) (observertypes.NonceToCctx, bool) {
	ret := _m.Called(ctx, tss, chainID, nonce)

	if len(ret) == 0 {
		panic("no return value specified for GetNonceToCctx")
	}

	var r0 observertypes.NonceToCctx
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, string, int64, int64) (observertypes.NonceToCctx, bool)); ok {
		return rf(ctx, tss, chainID, nonce)
	}
	if rf, ok := ret.Get(0).(func(types.Context, string, int64, int64) observertypes.NonceToCctx); ok {
		r0 = rf(ctx, tss, chainID, nonce)
	} else {
		r0 = ret.Get(0).(observertypes.NonceToCctx)
	}

	if rf, ok := ret.Get(1).(func(types.Context, string, int64, int64) bool); ok {
		r1 = rf(ctx, tss, chainID, nonce)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetObserverSet provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetObserverSet(ctx types.Context) (observertypes.ObserverSet, bool) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetObserverSet")
	}

	var r0 observertypes.ObserverSet
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context) (observertypes.ObserverSet, bool)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) observertypes.ObserverSet); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(observertypes.ObserverSet)
	}

	if rf, ok := ret.Get(1).(func(types.Context) bool); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetParams provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetParams(ctx types.Context) observertypes.Params {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetParams")
	}

	var r0 observertypes.Params
	if rf, ok := ret.Get(0).(func(types.Context) observertypes.Params); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(observertypes.Params)
	}

	return r0
}

// GetPendingNonces provides a mock function with given fields: ctx, tss, chainID
func (_m *CrosschainObserverKeeper) GetPendingNonces(ctx types.Context, tss string, chainID int64) (observertypes.PendingNonces, bool) {
	ret := _m.Called(ctx, tss, chainID)

	if len(ret) == 0 {
		panic("no return value specified for GetPendingNonces")
	}

	var r0 observertypes.PendingNonces
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, string, int64) (observertypes.PendingNonces, bool)); ok {
		return rf(ctx, tss, chainID)
	}
	if rf, ok := ret.Get(0).(func(types.Context, string, int64) observertypes.PendingNonces); ok {
		r0 = rf(ctx, tss, chainID)
	} else {
		r0 = ret.Get(0).(observertypes.PendingNonces)
	}

	if rf, ok := ret.Get(1).(func(types.Context, string, int64) bool); ok {
		r1 = rf(ctx, tss, chainID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetPreviousTSS provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetPreviousTSS(ctx types.Context) (observertypes.TSS, bool) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetPreviousTSS")
	}

	var r0 observertypes.TSS
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context) (observertypes.TSS, bool)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) observertypes.TSS); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(observertypes.TSS)
	}

	if rf, ok := ret.Get(1).(func(types.Context) bool); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetTSS provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) GetTSS(ctx types.Context) (observertypes.TSS, bool) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetTSS")
	}

	var r0 observertypes.TSS
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context) (observertypes.TSS, bool)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) observertypes.TSS); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(observertypes.TSS)
	}

	if rf, ok := ret.Get(1).(func(types.Context) bool); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetTssAddress provides a mock function with given fields: goCtx, req
func (_m *CrosschainObserverKeeper) GetTssAddress(goCtx context.Context, req *observertypes.QueryGetTssAddressRequest) (*observertypes.QueryGetTssAddressResponse, error) {
	ret := _m.Called(goCtx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetTssAddress")
	}

	var r0 *observertypes.QueryGetTssAddressResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *observertypes.QueryGetTssAddressRequest) (*observertypes.QueryGetTssAddressResponse, error)); ok {
		return rf(goCtx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *observertypes.QueryGetTssAddressRequest) *observertypes.QueryGetTssAddressResponse); ok {
		r0 = rf(goCtx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*observertypes.QueryGetTssAddressResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *observertypes.QueryGetTssAddressRequest) error); ok {
		r1 = rf(goCtx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAuthorized provides a mock function with given fields: ctx, address
func (_m *CrosschainObserverKeeper) IsAuthorized(ctx types.Context, address string) bool {
	ret := _m.Called(ctx, address)

	if len(ret) == 0 {
		panic("no return value specified for IsAuthorized")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.Context, string) bool); ok {
		r0 = rf(ctx, address)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsInboundEnabled provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) IsInboundEnabled(ctx types.Context) bool {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for IsInboundEnabled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoveAllExistingMigrators provides a mock function with given fields: ctx
func (_m *CrosschainObserverKeeper) RemoveAllExistingMigrators(ctx types.Context) {
	_m.Called(ctx)
}

// RemoveChainNonces provides a mock function with given fields: ctx, index
func (_m *CrosschainObserverKeeper) RemoveChainNonces(ctx types.Context, index string) {
	_m.Called(ctx, index)
}

// RemoveFromPendingNonces provides a mock function with given fields: ctx, tss, chainID, nonce
func (_m *CrosschainObserverKeeper) RemoveFromPendingNonces(ctx types.Context, tss string, chainID int64, nonce int64) {
	_m.Called(ctx, tss, chainID, nonce)
}

// RemoveNonceToCctx provides a mock function with given fields: ctx, cctx
func (_m *CrosschainObserverKeeper) RemoveNonceToCctx(ctx types.Context, cctx observertypes.NonceToCctx) {
	_m.Called(ctx, cctx)
}

// SetBallot provides a mock function with given fields: ctx, ballot
func (_m *CrosschainObserverKeeper) SetBallot(ctx types.Context, ballot *observertypes.Ballot) {
	_m.Called(ctx, ballot)
}

// SetChainNonces provides a mock function with given fields: ctx, chainNonces
func (_m *CrosschainObserverKeeper) SetChainNonces(ctx types.Context, chainNonces observertypes.ChainNonces) {
	_m.Called(ctx, chainNonces)
}

// SetCrosschainFlags provides a mock function with given fields: ctx, crosschainFlags
func (_m *CrosschainObserverKeeper) SetCrosschainFlags(ctx types.Context, crosschainFlags observertypes.CrosschainFlags) {
	_m.Called(ctx, crosschainFlags)
}

// SetFundMigrator provides a mock function with given fields: ctx, fm
func (_m *CrosschainObserverKeeper) SetFundMigrator(ctx types.Context, fm observertypes.TssFundMigratorInfo) {
	_m.Called(ctx, fm)
}

// SetKeygen provides a mock function with given fields: ctx, keygen
func (_m *CrosschainObserverKeeper) SetKeygen(ctx types.Context, keygen observertypes.Keygen) {
	_m.Called(ctx, keygen)
}

// SetLastObserverCount provides a mock function with given fields: ctx, lbc
func (_m *CrosschainObserverKeeper) SetLastObserverCount(ctx types.Context, lbc *observertypes.LastObserverCount) {
	_m.Called(ctx, lbc)
}

// SetNodeAccount provides a mock function with given fields: ctx, nodeAccount
func (_m *CrosschainObserverKeeper) SetNodeAccount(ctx types.Context, nodeAccount observertypes.NodeAccount) {
	_m.Called(ctx, nodeAccount)
}

// SetNonceToCctx provides a mock function with given fields: ctx, nonceToCctx
func (_m *CrosschainObserverKeeper) SetNonceToCctx(ctx types.Context, nonceToCctx observertypes.NonceToCctx) {
	_m.Called(ctx, nonceToCctx)
}

// SetObservers provides a mock function with given fields: ctx, om
func (_m *CrosschainObserverKeeper) SetObservers(ctx types.Context, om observertypes.ObserverSet) {
	_m.Called(ctx, om)
}

// SetPendingNonces provides a mock function with given fields: ctx, pendingNonces
func (_m *CrosschainObserverKeeper) SetPendingNonces(ctx types.Context, pendingNonces observertypes.PendingNonces) {
	_m.Called(ctx, pendingNonces)
}

// SetTSS provides a mock function with given fields: ctx, tss
func (_m *CrosschainObserverKeeper) SetTSS(ctx types.Context, tss observertypes.TSS) {
	_m.Called(ctx, tss)
}

// SetTSSHistory provides a mock function with given fields: ctx, tss
func (_m *CrosschainObserverKeeper) SetTSSHistory(ctx types.Context, tss observertypes.TSS) {
	_m.Called(ctx, tss)
}

// SetTssAndUpdateNonce provides a mock function with given fields: ctx, tss
func (_m *CrosschainObserverKeeper) SetTssAndUpdateNonce(ctx types.Context, tss observertypes.TSS) {
	_m.Called(ctx, tss)
}

// NewCrosschainObserverKeeper creates a new instance of CrosschainObserverKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCrosschainObserverKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *CrosschainObserverKeeper {
	mock := &CrosschainObserverKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
