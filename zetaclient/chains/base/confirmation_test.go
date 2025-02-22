package base_test

import (
	"context"
	"errors"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"
	"github.com/test-go/testify/mock"
	"github.com/zeta-chain/node/pkg/chains"
	"github.com/zeta-chain/node/pkg/coin"
	"github.com/zeta-chain/node/pkg/constant"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
	fungibletypes "github.com/zeta-chain/node/x/fungible/types"
	observertypes "github.com/zeta-chain/node/x/observer/types"
)

func Test_GetScanRangeInboundSafe(t *testing.T) {
	chain := chains.BitcoinMainnet

	tests := []struct {
		name               string
		lastBlock          uint64
		lastScanned        uint64
		blockLimit         uint64
		confParams         observertypes.ConfirmationParams
		expectedBlockRange [2]uint64
	}{
		{
			name:        "no unscanned blocks",
			lastBlock:   99,
			lastScanned: 90,
			blockLimit:  10,
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 10,
			},
			expectedBlockRange: [2]uint64{91, 91}, // [91, 91), nothing to scan
		},
		{
			name:        "1 unscanned blocks",
			lastBlock:   100,
			lastScanned: 90,
			blockLimit:  10,
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 10,
			},
			expectedBlockRange: [2]uint64{91, 92}, // [91, 92)
		},
		{
			name:        "10 unscanned blocks",
			lastBlock:   109,
			lastScanned: 90,
			blockLimit:  10,
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 10,
			},
			expectedBlockRange: [2]uint64{91, 101}, // [91, 101)
		},
		{
			name:        "block limit applied",
			lastBlock:   110,
			lastScanned: 90,
			blockLimit:  10,
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 10,
			},
			expectedBlockRange: [2]uint64{91, 101}, // [91, 101), 11 unscanned blocks, but capped to 10
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ob := newTestSuite(t, chain, withConfirmationParams(tt.confParams))
			ob.Observer.WithLastBlock(tt.lastBlock)
			ob.Observer.WithLastBlockScanned(tt.lastScanned)

			start, end := ob.GetScanRangeInboundSafe(tt.blockLimit)
			require.Equal(t, tt.expectedBlockRange, [2]uint64{start, end})
		})
	}
}

func Test_GetScanRangeInboundFast(t *testing.T) {
	chain := chains.BitcoinMainnet

	tests := []struct {
		name               string
		lastBlock          uint64
		lastScanned        uint64
		blockLimit         uint64
		confParams         observertypes.ConfirmationParams
		expectedBlockRange [2]uint64
	}{
		{
			name:        "no unscanned blocks",
			lastBlock:   99,
			lastScanned: 90,
			blockLimit:  10,
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 10,
				FastInboundCount: 10,
			},
			expectedBlockRange: [2]uint64{91, 91}, // [91, 91), nothing to scan
		},
		{
			name:        "1 unscanned blocks",
			lastBlock:   100,
			lastScanned: 90,
			blockLimit:  10,
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 10,
				FastInboundCount: 0, // should fall back to safe confirmation
			},
			expectedBlockRange: [2]uint64{91, 92}, // [91, 92)
		},
		{
			name:        "10 unscanned blocks",
			lastBlock:   109,
			lastScanned: 90,
			blockLimit:  10,
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 10,
				FastInboundCount: 10,
			},
			expectedBlockRange: [2]uint64{91, 101}, // [91, 101)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ob := newTestSuite(t, chain, withConfirmationParams(tt.confParams))
			ob.Observer.WithLastBlock(tt.lastBlock)
			ob.Observer.WithLastBlockScanned(tt.lastScanned)

			start, end := ob.GetScanRangeInboundFast(tt.blockLimit)
			require.Equal(t, tt.expectedBlockRange, [2]uint64{start, end})
		})
	}
}

func Test_IsBlockConfirmedForInboundSafe(t *testing.T) {
	chain := chains.BitcoinMainnet

	tests := []struct {
		name        string
		blockNumber uint64
		lastBlock   uint64
		confParams  observertypes.ConfirmationParams
		expected    bool
	}{
		{
			name:        "should confirm block 100 when confirmation arrives 2",
			blockNumber: 100,
			lastBlock:   101, // got 2 confirmations
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 2,
			},
			expected: true,
		},
		{
			name:        "should not confirm block 100 when confirmation < 2",
			blockNumber: 100,
			lastBlock:   100, // got 1 confirmation, need one more
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 2,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ob := newTestSuite(t, chain, withConfirmationParams(tt.confParams))
			ob.Observer.WithLastBlock(tt.lastBlock)

			isConfirmed := ob.IsBlockConfirmedForInboundSafe(tt.blockNumber)
			require.Equal(t, tt.expected, isConfirmed)
		})
	}
}

func Test_IsBlockConfirmedForInboundFast(t *testing.T) {
	chain := chains.BitcoinMainnet

	tests := []struct {
		name        string
		blockNumber uint64
		lastBlock   uint64
		confParams  observertypes.ConfirmationParams
		expected    bool
	}{
		{
			name:        "should confirm block 100 when confirmation arrives 2",
			blockNumber: 100,
			lastBlock:   101, // got 2 confirmations
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 2,
				FastInboundCount: 0, // should fall back to safe confirmation
			},
			expected: true,
		},
		{
			name:        "should not confirm block 100 when confirmation < 2",
			blockNumber: 100,
			lastBlock:   100, // got 1 confirmation, need one more
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 2,
				FastInboundCount: 2,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ob := newTestSuite(t, chain, withConfirmationParams(tt.confParams))
			ob.Observer.WithLastBlock(tt.lastBlock)

			isConfirmed := ob.IsBlockConfirmedForInboundFast(tt.blockNumber)
			require.Equal(t, tt.expected, isConfirmed)
		})
	}
}

func Test_GetInboundConfirmationMode(t *testing.T) {
	chain := chains.BitcoinMainnet

	tests := []struct {
		name         string
		scannedBlock uint64
		lastBlock    uint64
		confParams   observertypes.ConfirmationParams
		expected     crosschaintypes.ConfirmationMode
	}{
		{
			name:         "should return SAFE confirmation mode",
			scannedBlock: 100,
			lastBlock:    101, // got 2 confirmations
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 2,
			},
			expected: crosschaintypes.ConfirmationMode_SAFE,
		},
		{
			name:         "should return FAST confirmation mode",
			scannedBlock: 100,
			lastBlock:    100, // got 1 confirmation, need one more
			confParams: observertypes.ConfirmationParams{
				SafeInboundCount: 2,
			},
			expected: crosschaintypes.ConfirmationMode_FAST,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ob := newTestSuite(t, chain, withConfirmationParams(tt.confParams))
			ob.Observer.WithLastBlock(tt.lastBlock)

			mode := ob.GetInboundConfirmationMode(tt.scannedBlock)
			require.Equal(t, tt.expected, mode)
		})
	}
}

func Test_IsBlockConfirmedForOutboundSafe(t *testing.T) {
	chain := chains.BitcoinMainnet

	tests := []struct {
		name        string
		blockNumber uint64
		lastBlock   uint64
		confParams  observertypes.ConfirmationParams
		expected    bool
	}{
		{
			name:        "should confirm block 100 when confirmation arrives 2",
			blockNumber: 100,
			lastBlock:   101, // got 2 confirmations
			confParams: observertypes.ConfirmationParams{
				SafeOutboundCount: 2,
			},
			expected: true,
		},
		{
			name:        "should not confirm block 100 when confirmation < 2",
			blockNumber: 100,
			lastBlock:   100, // got 1 confirmation, need one more
			confParams: observertypes.ConfirmationParams{
				SafeOutboundCount: 2,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ob := newTestSuite(t, chain, withConfirmationParams(tt.confParams))
			ob.Observer.WithLastBlock(tt.lastBlock)

			isConfirmed := ob.IsBlockConfirmedForOutboundSafe(tt.blockNumber)
			require.Equal(t, tt.expected, isConfirmed)
		})
	}
}

func Test_IsBlockConfirmedForOutboundFast(t *testing.T) {
	chain := chains.BitcoinMainnet

	tests := []struct {
		name        string
		blockNumber uint64
		lastBlock   uint64
		confParams  observertypes.ConfirmationParams
		expected    bool
	}{
		{
			name:        "should confirm block 100 when confirmation arrives 2",
			blockNumber: 100,
			lastBlock:   101, // got 2 confirmations
			confParams: observertypes.ConfirmationParams{
				SafeOutboundCount: 2,
				FastOutboundCount: 0, // should fall back to safe confirmation
			},
			expected: true,
		},
		{
			name:        "should not confirm block 100 when confirmation < 2",
			blockNumber: 100,
			lastBlock:   100, // got 1 confirmation, need one more
			confParams: observertypes.ConfirmationParams{
				SafeOutboundCount: 2,
				FastOutboundCount: 2,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ob := newTestSuite(t, chain, withConfirmationParams(tt.confParams))
			ob.Observer.WithLastBlock(tt.lastBlock)

			isConfirmed := ob.IsBlockConfirmedForOutboundFast(tt.blockNumber)
			require.Equal(t, tt.expected, isConfirmed)
		})
	}
}

func Test_IsInboundEligibleForFastConfirmation(t *testing.T) {
	chain := chains.Ethereum
	liquidityCap := sdkmath.NewUint(100_000)
	multiplier := constant.DefaultInboundFastConfirmationLiquidityMultiplier
	fastAmountCap := constant.CalcInboundFastAmountCap(liquidityCap, multiplier)

	tests := []struct {
		name                string
		msg                 *crosschaintypes.MsgVoteInbound
		failForeignCoinsRPC bool
		eligible            bool
		errMsg              string
	}{
		{
			name: "eligible for fast confirmation",
			msg: &crosschaintypes.MsgVoteInbound{
				SenderChainId:           chain.ChainId,
				Amount:                  sdkmath.NewUint(fastAmountCap.Uint64()),
				CoinType:                coin.CoinType_Gas,
				Asset:                   "",
				ProtocolContractVersion: crosschaintypes.ProtocolContractVersion_V2,
			},
			eligible: true,
		},
		{
			name: "not eligible if multiplier not set for chain id",
			msg: &crosschaintypes.MsgVoteInbound{
				SenderChainId: chains.SolanaMainnet.ChainId, // not set for Solana
			},
			eligible: false,
		},
		{
			name: "not eligible if not fungible",
			msg: &crosschaintypes.MsgVoteInbound{
				SenderChainId:           chain.ChainId,
				ProtocolContractVersion: crosschaintypes.ProtocolContractVersion_V1, // not eligible for V1
			},
			eligible: false,
		},
		{
			name: "return error if foreign coins query RPC fails",
			msg: &crosschaintypes.MsgVoteInbound{
				SenderChainId:           chain.ChainId,
				CoinType:                coin.CoinType_Gas,
				Asset:                   "",
				ProtocolContractVersion: crosschaintypes.ProtocolContractVersion_V2,
			},
			failForeignCoinsRPC: true,
			eligible:            false,
			errMsg:              "unable to get foreign coins",
		},
		{
			name: "not eligible if amount exceeds fast amount cap",
			msg: &crosschaintypes.MsgVoteInbound{
				SenderChainId:           chain.ChainId,
				Amount:                  sdkmath.NewUint(fastAmountCap.Uint64() + 1), // +1 to exceed
				CoinType:                coin.CoinType_Gas,
				Asset:                   "",
				ProtocolContractVersion: crosschaintypes.ProtocolContractVersion_V2,
			},
			eligible: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ARRANGE
			ob := newTestSuite(t, chain)

			// mock up the foreign coins RPC
			if tt.failForeignCoinsRPC {
				ob.zetacore.On("GetForeignCoinsFromAsset", mock.Anything, chain.ChainId, tt.msg.Asset).
					Maybe().
					Return(fungibletypes.ForeignCoins{}, errors.New("rpc failed"))
			} else {
				ob.zetacore.On("GetForeignCoinsFromAsset", mock.Anything, chain.ChainId, tt.msg.Asset).Maybe().Return(fungibletypes.ForeignCoins{LiquidityCap: liquidityCap}, nil)
			}

			// ACT
			ctx := context.Background()
			eligible, err := ob.IsInboundEligibleForFastConfirmation(ctx, tt.msg)

			// ASSERT
			require.Equal(t, tt.eligible, eligible)
			if tt.errMsg != "" {
				require.Contains(t, err.Error(), tt.errMsg)
				return
			}
			require.NoError(t, err)
		})
	}
}
