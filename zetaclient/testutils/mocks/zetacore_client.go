// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	chains "github.com/zeta-chain/node/pkg/chains"
	blame "gitlab.com/thorchain/tss/go-tss/blame"

	cometbfttypes "github.com/cometbft/cometbft/types"

	context "context"

	interfaces "github.com/zeta-chain/node/zetaclient/chains/interfaces"

	keysinterfaces "github.com/zeta-chain/node/zetaclient/keys/interfaces"

	math "cosmossdk.io/math"

	mock "github.com/stretchr/testify/mock"

	observertypes "github.com/zeta-chain/node/x/observer/types"

	types "github.com/zeta-chain/node/x/crosschain/types"

	upgradetypes "cosmossdk.io/x/upgrade/types"
)

// ZetacoreClient is an autogenerated mock type for the ZetacoreClient type
type ZetacoreClient struct {
	mock.Mock
}

// Chain provides a mock function with given fields:
func (_m *ZetacoreClient) Chain() chains.Chain {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Chain")
	}

	var r0 chains.Chain
	if rf, ok := ret.Get(0).(func() chains.Chain); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(chains.Chain)
	}

	return r0
}

// GetAdditionalChains provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetAdditionalChains(ctx context.Context) ([]chains.Chain, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAdditionalChains")
	}

	var r0 []chains.Chain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]chains.Chain, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []chains.Chain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chains.Chain)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllOutboundTrackerByChain provides a mock function with given fields: ctx, chainID, order
func (_m *ZetacoreClient) GetAllOutboundTrackerByChain(ctx context.Context, chainID int64, order interfaces.Order) ([]types.OutboundTracker, error) {
	ret := _m.Called(ctx, chainID, order)

	if len(ret) == 0 {
		panic("no return value specified for GetAllOutboundTrackerByChain")
	}

	var r0 []types.OutboundTracker
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, interfaces.Order) ([]types.OutboundTracker, error)); ok {
		return rf(ctx, chainID, order)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, interfaces.Order) []types.OutboundTracker); ok {
		r0 = rf(ctx, chainID, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.OutboundTracker)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, interfaces.Order) error); ok {
		r1 = rf(ctx, chainID, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBTCTSSAddress provides a mock function with given fields: ctx, chainID
func (_m *ZetacoreClient) GetBTCTSSAddress(ctx context.Context, chainID int64) (string, error) {
	ret := _m.Called(ctx, chainID)

	if len(ret) == 0 {
		panic("no return value specified for GetBTCTSSAddress")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (string, error)); ok {
		return rf(ctx, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) string); ok {
		r0 = rf(ctx, chainID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHeight provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetBlockHeight(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockHeight")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCctxByNonce provides a mock function with given fields: ctx, chainID, nonce
func (_m *ZetacoreClient) GetCctxByNonce(ctx context.Context, chainID int64, nonce uint64) (*types.CrossChainTx, error) {
	ret := _m.Called(ctx, chainID, nonce)

	if len(ret) == 0 {
		panic("no return value specified for GetCctxByNonce")
	}

	var r0 *types.CrossChainTx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint64) (*types.CrossChainTx, error)); ok {
		return rf(ctx, chainID, nonce)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint64) *types.CrossChainTx); ok {
		r0 = rf(ctx, chainID, nonce)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.CrossChainTx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, uint64) error); ok {
		r1 = rf(ctx, chainID, nonce)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChainParams provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetChainParams(ctx context.Context) ([]*observertypes.ChainParams, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetChainParams")
	}

	var r0 []*observertypes.ChainParams
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*observertypes.ChainParams, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*observertypes.ChainParams); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*observertypes.ChainParams)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCrosschainFlags provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetCrosschainFlags(ctx context.Context) (observertypes.CrosschainFlags, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetCrosschainFlags")
	}

	var r0 observertypes.CrosschainFlags
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (observertypes.CrosschainFlags, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) observertypes.CrosschainFlags); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(observertypes.CrosschainFlags)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInboundTrackersForChain provides a mock function with given fields: ctx, chainID
func (_m *ZetacoreClient) GetInboundTrackersForChain(ctx context.Context, chainID int64) ([]types.InboundTracker, error) {
	ret := _m.Called(ctx, chainID)

	if len(ret) == 0 {
		panic("no return value specified for GetInboundTrackersForChain")
	}

	var r0 []types.InboundTracker
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]types.InboundTracker, error)); ok {
		return rf(ctx, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []types.InboundTracker); ok {
		r0 = rf(ctx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.InboundTracker)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKeyGen provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetKeyGen(ctx context.Context) (observertypes.Keygen, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetKeyGen")
	}

	var r0 observertypes.Keygen
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (observertypes.Keygen, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) observertypes.Keygen); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(observertypes.Keygen)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKeys provides a mock function with given fields:
func (_m *ZetacoreClient) GetKeys() keysinterfaces.ObserverKeys {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetKeys")
	}

	var r0 keysinterfaces.ObserverKeys
	if rf, ok := ret.Get(0).(func() keysinterfaces.ObserverKeys); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keysinterfaces.ObserverKeys)
		}
	}

	return r0
}

// GetObserverList provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetObserverList(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetObserverList")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOperationalFlags provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetOperationalFlags(ctx context.Context) (observertypes.OperationalFlags, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetOperationalFlags")
	}

	var r0 observertypes.OperationalFlags
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (observertypes.OperationalFlags, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) observertypes.OperationalFlags); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(observertypes.OperationalFlags)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOutboundTracker provides a mock function with given fields: ctx, chain, nonce
func (_m *ZetacoreClient) GetOutboundTracker(ctx context.Context, chain chains.Chain, nonce uint64) (*types.OutboundTracker, error) {
	ret := _m.Called(ctx, chain, nonce)

	if len(ret) == 0 {
		panic("no return value specified for GetOutboundTracker")
	}

	var r0 *types.OutboundTracker
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, chains.Chain, uint64) (*types.OutboundTracker, error)); ok {
		return rf(ctx, chain, nonce)
	}
	if rf, ok := ret.Get(0).(func(context.Context, chains.Chain, uint64) *types.OutboundTracker); ok {
		r0 = rf(ctx, chain, nonce)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.OutboundTracker)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, chains.Chain, uint64) error); ok {
		r1 = rf(ctx, chain, nonce)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPendingNoncesByChain provides a mock function with given fields: ctx, chainID
func (_m *ZetacoreClient) GetPendingNoncesByChain(ctx context.Context, chainID int64) (observertypes.PendingNonces, error) {
	ret := _m.Called(ctx, chainID)

	if len(ret) == 0 {
		panic("no return value specified for GetPendingNoncesByChain")
	}

	var r0 observertypes.PendingNonces
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (observertypes.PendingNonces, error)); ok {
		return rf(ctx, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) observertypes.PendingNonces); ok {
		r0 = rf(ctx, chainID)
	} else {
		r0 = ret.Get(0).(observertypes.PendingNonces)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRateLimiterFlags provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetRateLimiterFlags(ctx context.Context) (types.RateLimiterFlags, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetRateLimiterFlags")
	}

	var r0 types.RateLimiterFlags
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (types.RateLimiterFlags, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) types.RateLimiterFlags); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.RateLimiterFlags)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRateLimiterInput provides a mock function with given fields: ctx, window
func (_m *ZetacoreClient) GetRateLimiterInput(ctx context.Context, window int64) (*types.QueryRateLimiterInputResponse, error) {
	ret := _m.Called(ctx, window)

	if len(ret) == 0 {
		panic("no return value specified for GetRateLimiterInput")
	}

	var r0 *types.QueryRateLimiterInputResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*types.QueryRateLimiterInputResponse, error)); ok {
		return rf(ctx, window)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *types.QueryRateLimiterInputResponse); ok {
		r0 = rf(ctx, window)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryRateLimiterInputResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, window)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSupportedChains provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetSupportedChains(ctx context.Context) ([]chains.Chain, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetSupportedChains")
	}

	var r0 []chains.Chain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]chains.Chain, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []chains.Chain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chains.Chain)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTSS provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetTSS(ctx context.Context) (observertypes.TSS, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetTSS")
	}

	var r0 observertypes.TSS
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (observertypes.TSS, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) observertypes.TSS); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(observertypes.TSS)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTSSHistory provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetTSSHistory(ctx context.Context) ([]observertypes.TSS, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetTSSHistory")
	}

	var r0 []observertypes.TSS
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]observertypes.TSS, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []observertypes.TSS); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]observertypes.TSS)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpgradePlan provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetUpgradePlan(ctx context.Context) (*upgradetypes.Plan, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetUpgradePlan")
	}

	var r0 *upgradetypes.Plan
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*upgradetypes.Plan, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *upgradetypes.Plan); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*upgradetypes.Plan)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetZetaHotKeyBalance provides a mock function with given fields: ctx
func (_m *ZetacoreClient) GetZetaHotKeyBalance(ctx context.Context) (math.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetZetaHotKeyBalance")
	}

	var r0 math.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (math.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) math.Int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPendingCCTX provides a mock function with given fields: ctx, chain
func (_m *ZetacoreClient) ListPendingCCTX(ctx context.Context, chain chains.Chain) ([]*types.CrossChainTx, uint64, error) {
	ret := _m.Called(ctx, chain)

	if len(ret) == 0 {
		panic("no return value specified for ListPendingCCTX")
	}

	var r0 []*types.CrossChainTx
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, chains.Chain) ([]*types.CrossChainTx, uint64, error)); ok {
		return rf(ctx, chain)
	}
	if rf, ok := ret.Get(0).(func(context.Context, chains.Chain) []*types.CrossChainTx); ok {
		r0 = rf(ctx, chain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.CrossChainTx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, chains.Chain) uint64); ok {
		r1 = rf(ctx, chain)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, chains.Chain) error); ok {
		r2 = rf(ctx, chain)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListPendingCCTXWithinRateLimit provides a mock function with given fields: ctx
func (_m *ZetacoreClient) ListPendingCCTXWithinRateLimit(ctx context.Context) (*types.QueryListPendingCctxWithinRateLimitResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListPendingCCTXWithinRateLimit")
	}

	var r0 *types.QueryListPendingCctxWithinRateLimitResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*types.QueryListPendingCctxWithinRateLimitResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *types.QueryListPendingCctxWithinRateLimitResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryListPendingCctxWithinRateLimitResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBlockSubscriber provides a mock function with given fields: ctx
func (_m *ZetacoreClient) NewBlockSubscriber(ctx context.Context) (chan cometbfttypes.EventDataNewBlock, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for NewBlockSubscriber")
	}

	var r0 chan cometbfttypes.EventDataNewBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (chan cometbfttypes.EventDataNewBlock, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) chan cometbfttypes.EventDataNewBlock); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan cometbfttypes.EventDataNewBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostOutboundTracker provides a mock function with given fields: ctx, chainID, nonce, txHash
func (_m *ZetacoreClient) PostOutboundTracker(ctx context.Context, chainID int64, nonce uint64, txHash string) (string, error) {
	ret := _m.Called(ctx, chainID, nonce, txHash)

	if len(ret) == 0 {
		panic("no return value specified for PostOutboundTracker")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint64, string) (string, error)); ok {
		return rf(ctx, chainID, nonce, txHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint64, string) string); ok {
		r0 = rf(ctx, chainID, nonce, txHash)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, uint64, string) error); ok {
		r1 = rf(ctx, chainID, nonce, txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostVoteBlameData provides a mock function with given fields: ctx, _a1, chainID, index
func (_m *ZetacoreClient) PostVoteBlameData(ctx context.Context, _a1 *blame.Blame, chainID int64, index string) (string, error) {
	ret := _m.Called(ctx, _a1, chainID, index)

	if len(ret) == 0 {
		panic("no return value specified for PostVoteBlameData")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *blame.Blame, int64, string) (string, error)); ok {
		return rf(ctx, _a1, chainID, index)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *blame.Blame, int64, string) string); ok {
		r0 = rf(ctx, _a1, chainID, index)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *blame.Blame, int64, string) error); ok {
		r1 = rf(ctx, _a1, chainID, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostVoteGasPrice provides a mock function with given fields: ctx, chain, gasPrice, priorityFee, blockNum
func (_m *ZetacoreClient) PostVoteGasPrice(ctx context.Context, chain chains.Chain, gasPrice uint64, priorityFee uint64, blockNum uint64) (string, error) {
	ret := _m.Called(ctx, chain, gasPrice, priorityFee, blockNum)

	if len(ret) == 0 {
		panic("no return value specified for PostVoteGasPrice")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, chains.Chain, uint64, uint64, uint64) (string, error)); ok {
		return rf(ctx, chain, gasPrice, priorityFee, blockNum)
	}
	if rf, ok := ret.Get(0).(func(context.Context, chains.Chain, uint64, uint64, uint64) string); ok {
		r0 = rf(ctx, chain, gasPrice, priorityFee, blockNum)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, chains.Chain, uint64, uint64, uint64) error); ok {
		r1 = rf(ctx, chain, gasPrice, priorityFee, blockNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostVoteInbound provides a mock function with given fields: ctx, gasLimit, retryGasLimit, msg
func (_m *ZetacoreClient) PostVoteInbound(ctx context.Context, gasLimit uint64, retryGasLimit uint64, msg *types.MsgVoteInbound) (string, string, error) {
	ret := _m.Called(ctx, gasLimit, retryGasLimit, msg)

	if len(ret) == 0 {
		panic("no return value specified for PostVoteInbound")
	}

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, *types.MsgVoteInbound) (string, string, error)); ok {
		return rf(ctx, gasLimit, retryGasLimit, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, *types.MsgVoteInbound) string); ok {
		r0 = rf(ctx, gasLimit, retryGasLimit, msg)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, *types.MsgVoteInbound) string); ok {
		r1 = rf(ctx, gasLimit, retryGasLimit, msg)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64, uint64, *types.MsgVoteInbound) error); ok {
		r2 = rf(ctx, gasLimit, retryGasLimit, msg)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PostVoteOutbound provides a mock function with given fields: ctx, gasLimit, retryGasLimit, msg
func (_m *ZetacoreClient) PostVoteOutbound(ctx context.Context, gasLimit uint64, retryGasLimit uint64, msg *types.MsgVoteOutbound) (string, string, error) {
	ret := _m.Called(ctx, gasLimit, retryGasLimit, msg)

	if len(ret) == 0 {
		panic("no return value specified for PostVoteOutbound")
	}

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, *types.MsgVoteOutbound) (string, string, error)); ok {
		return rf(ctx, gasLimit, retryGasLimit, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, *types.MsgVoteOutbound) string); ok {
		r0 = rf(ctx, gasLimit, retryGasLimit, msg)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, *types.MsgVoteOutbound) string); ok {
		r1 = rf(ctx, gasLimit, retryGasLimit, msg)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64, uint64, *types.MsgVoteOutbound) error); ok {
		r2 = rf(ctx, gasLimit, retryGasLimit, msg)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PostVoteTSS provides a mock function with given fields: ctx, tssPubKey, keyGenZetaHeight, status
func (_m *ZetacoreClient) PostVoteTSS(ctx context.Context, tssPubKey string, keyGenZetaHeight int64, status chains.ReceiveStatus) (string, error) {
	ret := _m.Called(ctx, tssPubKey, keyGenZetaHeight, status)

	if len(ret) == 0 {
		panic("no return value specified for PostVoteTSS")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, chains.ReceiveStatus) (string, error)); ok {
		return rf(ctx, tssPubKey, keyGenZetaHeight, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, chains.ReceiveStatus) string); ok {
		r0 = rf(ctx, tssPubKey, keyGenZetaHeight, status)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, chains.ReceiveStatus) error); ok {
		r1 = rf(ctx, tssPubKey, keyGenZetaHeight, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewZetacoreClient creates a new instance of ZetacoreClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewZetacoreClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ZetacoreClient {
	mock := &ZetacoreClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
