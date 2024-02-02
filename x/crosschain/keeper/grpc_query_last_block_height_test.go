package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/assert"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLastBlockHeightQuerySingle(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLastBlockHeight(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLastBlockHeightRequest
		response *types.QueryGetLastBlockHeightResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLastBlockHeightRequest{Index: msgs[0].Index},
			response: &types.QueryGetLastBlockHeightResponse{LastBlockHeight: &msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetLastBlockHeightRequest{Index: msgs[1].Index},
			response: &types.QueryGetLastBlockHeightResponse{LastBlockHeight: &msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetLastBlockHeightRequest{Index: "missing"},
			err:     status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LastBlockHeight(wctx, tc.request)
			if tc.err != nil {
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.Equal(t, tc.response, response)
			}
		})
	}
}

func TestLastBlockHeightQueryPaginated(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLastBlockHeight(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLastBlockHeightRequest {
		return &types.QueryAllLastBlockHeightRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LastBlockHeightAll(wctx, request(nil, uint64(i), uint64(step), false))
			assert.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, &msgs[j], resp.LastBlockHeight[j-i])
			}
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LastBlockHeightAll(wctx, request(next, 0, uint64(step), false))
			assert.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, &msgs[j], resp.LastBlockHeight[j-i])
			}
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.LastBlockHeightAll(wctx, request(nil, 0, 0, true))
		assert.NoError(t, err)
		assert.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.LastBlockHeightAll(wctx, nil)
		assert.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
