package observer

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/zeta-chain/node/pkg/chains"
	"github.com/zeta-chain/node/zetaclient/chains/bitcoin/client"
	"github.com/zeta-chain/node/zetaclient/chains/bitcoin/common"
)

// PostGasPrice posts gas price to zetacore
func (ob *Observer) PostGasPrice(ctx context.Context) error {
	var (
		err              error
		feeRateEstimated int64
	)

	// estimate fee rate according to network type
	switch ob.Chain().NetworkType {
	case chains.NetworkType_privnet:
		// regnet RPC 'EstimateSmartFee' is not available
		feeRateEstimated = client.FeeRateRegnet
	case chains.NetworkType_testnet:
		// testnet RPC 'EstimateSmartFee' can return unreasonable high fee rate
		feeRateEstimated, err = common.GetRecentFeeRate(ctx, ob.rpc, ob.netParams)
		if err != nil {
			return errors.Wrapf(err, "unable to get recent fee rate")
		}
	case chains.NetworkType_mainnet:
		feeRateEstimated, err = ob.rpc.GetEstimatedFeeRate(ctx, 1, false)
		if err != nil {
			return errors.Wrap(err, "unable to get estimated fee rate")
		}
	default:
		return fmt.Errorf("unsupported bitcoin network type %d", ob.Chain().NetworkType)
	}

	// query the current block number
	blockNumber, err := ob.rpc.GetBlockCount(ctx)
	if err != nil {
		return errors.Wrap(err, "GetBlockCount error")
	}

	// Bitcoin has no concept of priority fee (like eth)
	const priorityFee = 0

	// #nosec G115 always positive
	_, err = ob.ZetacoreClient().
		PostVoteGasPrice(ctx, ob.Chain(), uint64(feeRateEstimated), priorityFee, uint64(blockNumber))
	if err != nil {
		return errors.Wrap(err, "PostVoteGasPrice error")
	}

	return nil
}
