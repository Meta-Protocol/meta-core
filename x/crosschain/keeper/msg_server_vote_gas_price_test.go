package keeper_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/zetacore/pkg/chains"
	keepertest "github.com/zeta-chain/zetacore/testutil/keeper"
	"github.com/zeta-chain/zetacore/testutil/sample"
	"github.com/zeta-chain/zetacore/x/crosschain/keeper"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

func TestMsgServer_VoteGasPrice(t *testing.T) {
	t.Run("should error if unsupported chain", func(t *testing.T) {
		k, ctx, _, _ := keepertest.CrosschainKeeperWithMocks(t, keepertest.CrosschainMockOptions{
			UseObserverMock: true,
		})

		observerMock := keepertest.GetCrosschainObserverMock(t, k)
		keepertest.MockFailedGetSupportedChainFromChainID(observerMock, sample.Chain(5))
		msgServer := keeper.NewMsgServerImpl(*k)

		res, err := msgServer.VoteGasPrice(ctx, &types.MsgVoteGasPrice{
			ChainId: 5,
		})
		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("should error if not non tombstoned observer", func(t *testing.T) {
		k, ctx, _, _ := keepertest.CrosschainKeeperWithMocks(t, keepertest.CrosschainMockOptions{
			UseObserverMock: true,
		})

		observerMock := keepertest.GetCrosschainObserverMock(t, k)
		keepertest.MockGetSupportedChainFromChainID(observerMock, sample.Chain(5))
		observerMock.On("IsNonTombstonedObserver", mock.Anything, mock.Anything).Return(false)

		msgServer := keeper.NewMsgServerImpl(*k)

		res, err := msgServer.VoteGasPrice(ctx, &types.MsgVoteGasPrice{
			ChainId: 5,
		})
		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("should error if gas price not found and set gas price in fungible keeper fails", func(t *testing.T) {
		k, ctx, _, _ := keepertest.CrosschainKeeperWithMocks(t, keepertest.CrosschainMockOptions{
			UseObserverMock: true,
			UseFungibleMock: true,
		})

		observerMock := keepertest.GetCrosschainObserverMock(t, k)
		keepertest.MockGetSupportedChainFromChainID(observerMock, chains.Chain{
			ChainId: 5,
		})
		observerMock.On("IsNonTombstonedObserver", mock.Anything, mock.Anything).Return(true)

		fungibleMock := keepertest.GetCrosschainFungibleMock(t, k)
		fungibleMock.On("SetGasPrice", mock.Anything, mock.Anything, mock.Anything).Return(uint64(0), errors.New("err"))
		msgServer := keeper.NewMsgServerImpl(*k)

		res, err := msgServer.VoteGasPrice(ctx, &types.MsgVoteGasPrice{
			ChainId: 5,
		})
		require.Error(t, err)
		require.Nil(t, res)
		_, found := k.GetGasPrice(ctx, 5)
		require.True(t, found)
	})

	t.Run("should not error if gas price not found and set gas price in fungible keeper succeeds", func(t *testing.T) {
		k, ctx, _, _ := keepertest.CrosschainKeeperWithMocks(t, keepertest.CrosschainMockOptions{
			UseObserverMock: true,
			UseFungibleMock: true,
		})

		observerMock := keepertest.GetCrosschainObserverMock(t, k)
		keepertest.MockGetSupportedChainFromChainID(observerMock, chains.Chain{
			ChainId: 5,
		})
		observerMock.On("IsNonTombstonedObserver", mock.Anything, mock.Anything).Return(true)

		fungibleMock := keepertest.GetCrosschainFungibleMock(t, k)
		fungibleMock.On("SetGasPrice", mock.Anything, mock.Anything, mock.Anything).Return(uint64(1), nil)
		msgServer := keeper.NewMsgServerImpl(*k)
		creator := sample.AccAddress()
		res, err := msgServer.VoteGasPrice(ctx, &types.MsgVoteGasPrice{
			Creator:     creator,
			ChainId:     5,
			Price:       1,
			BlockNumber: 1,
		})
		require.NoError(t, err)
		require.Empty(t, res)
		gp, found := k.GetGasPrice(ctx, 5)
		require.True(t, found)
		require.Equal(t, types.GasPrice{
			Creator:     creator,
			Index:       "5",
			ChainId:     5,
			Signers:     []string{creator},
			BlockNums:   []uint64{1},
			Prices:      []uint64{1},
			MedianIndex: 0,
		}, gp)
	})

	t.Run("should not error if gas price found and msg.creator in signers", func(t *testing.T) {
		k, ctx, _, _ := keepertest.CrosschainKeeperWithMocks(t, keepertest.CrosschainMockOptions{
			UseObserverMock: true,
			UseFungibleMock: true,
		})

		observerMock := keepertest.GetCrosschainObserverMock(t, k)
		keepertest.MockGetSupportedChainFromChainID(observerMock, chains.Chain{
			ChainId: 5,
		})
		observerMock.On("IsNonTombstonedObserver", mock.Anything, mock.Anything).Return(true)

		fungibleMock := keepertest.GetCrosschainFungibleMock(t, k)
		fungibleMock.On("SetGasPrice", mock.Anything, mock.Anything, mock.Anything).Return(uint64(1), nil)
		msgServer := keeper.NewMsgServerImpl(*k)

		creator := sample.AccAddress()
		k.SetGasPrice(ctx, types.GasPrice{
			Creator:   creator,
			ChainId:   5,
			Signers:   []string{creator},
			BlockNums: []uint64{1},
			Prices:    []uint64{1},
		})

		res, err := msgServer.VoteGasPrice(ctx, &types.MsgVoteGasPrice{
			Creator:     creator,
			ChainId:     5,
			BlockNumber: 2,
			Price:       2,
		})
		require.NoError(t, err)
		require.Empty(t, res)
		gp, found := k.GetGasPrice(ctx, 5)
		require.True(t, found)
		require.Equal(t, types.GasPrice{
			Creator:     creator,
			Index:       "",
			ChainId:     5,
			Signers:     []string{creator},
			BlockNums:   []uint64{2},
			Prices:      []uint64{2},
			MedianIndex: 0,
		}, gp)
	})

	t.Run("should not error if gas price found and msg.creator not in signers", func(t *testing.T) {
		k, ctx, _, _ := keepertest.CrosschainKeeperWithMocks(t, keepertest.CrosschainMockOptions{
			UseObserverMock: true,
			UseFungibleMock: true,
		})

		observerMock := keepertest.GetCrosschainObserverMock(t, k)
		keepertest.MockGetSupportedChainFromChainID(observerMock, chains.Chain{
			ChainId: 5,
		})
		observerMock.On("IsNonTombstonedObserver", mock.Anything, mock.Anything).Return(true)

		fungibleMock := keepertest.GetCrosschainFungibleMock(t, k)
		fungibleMock.On("SetGasPrice", mock.Anything, mock.Anything, mock.Anything).Return(uint64(1), nil)
		msgServer := keeper.NewMsgServerImpl(*k)

		creator := sample.AccAddress()
		k.SetGasPrice(ctx, types.GasPrice{
			Creator:   creator,
			ChainId:   5,
			BlockNums: []uint64{1},
			Prices:    []uint64{1},
		})

		res, err := msgServer.VoteGasPrice(ctx, &types.MsgVoteGasPrice{
			Creator:     creator,
			ChainId:     5,
			BlockNumber: 2,
			Price:       2,
		})
		require.NoError(t, err)
		require.Empty(t, res)
		gp, found := k.GetGasPrice(ctx, 5)
		require.True(t, found)
		require.Equal(t, types.GasPrice{
			Creator:     creator,
			Index:       "",
			ChainId:     5,
			Signers:     []string{creator},
			BlockNums:   []uint64{1, 2},
			Prices:      []uint64{1, 2},
			MedianIndex: 1,
		}, gp)
	})
}
