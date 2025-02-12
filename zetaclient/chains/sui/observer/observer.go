package observer

import (
	"context"
	"strconv"
	"time"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/pkg/errors"

	"github.com/zeta-chain/node/zetaclient/chains/base"
)

// Observer SUI observer
type Observer struct {
	*base.Observer
	client RPC
}

// RPC represents subset of SUI RPC methods.
type RPC interface {
	HealthCheck(ctx context.Context) (time.Time, error)
	GetLatestCheckpoint(ctx context.Context) (models.CheckpointResponse, error)

	SuiXGetReferenceGasPrice(ctx context.Context) (uint64, error)
}

// New Observer constructor.
func New(baseObserver *base.Observer, client RPC) *Observer {
	return &Observer{
		Observer: baseObserver,
		client:   client,
	}
}

// CheckRPCStatus checks the RPC status of the chain.
func (ob *Observer) CheckRPCStatus(ctx context.Context) error {
	blockTime, err := ob.client.HealthCheck(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to check rpc health")
	}

	// It's not a "real" block latency as SUI uses concept of "checkpoints"
	ob.ReportBlockLatency(blockTime)

	return nil
}

// PostGasPrice posts SUI gas price to zetacore.
// Note (1) that SUI changes gas per EPOCH (not block)
// Note (2) that SuiXGetCurrentEpoch() is deprecated.
//
// See https://docs.sui.io/concepts/tokenomics/gas-pricing
// See https://docs.sui.io/concepts/sui-architecture/transaction-lifecycle#epoch-change
//
// TLDR:
// - GasFees = CompUnits * (ReferencePrice + Tip) + StorageUnits * StoragePrice
// - "During regular network usage, users are NOT expected to pay tips"
// - "Validators update the ReferencePrice every epoch (~24h)"
// - "Storage price is updated infrequently through gov proposals"
func (ob *Observer) PostGasPrice(ctx context.Context) error {
	checkpoint, err := ob.client.GetLatestCheckpoint(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to get latest checkpoint")
	}

	epochNum, err := uint64FromStr(checkpoint.Epoch)
	if err != nil {
		return errors.Wrap(err, "unable to parse epoch number")
	}

	// gas price in MIST. 1 SUI = 10^9 MIST (a billion)
	// e.g. { "jsonrpc": "2.0", "id": 1, "result": "750" }
	gasPrice, err := ob.client.SuiXGetReferenceGasPrice(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to get ref gas price")
	}

	// no priority fee for SUI
	const priorityFee = 0

	_, err = ob.ZetacoreClient().PostVoteGasPrice(ctx, ob.Chain(), gasPrice, priorityFee, epochNum)
	if err != nil {
		return errors.Wrap(err, "unable to post vote for gas price")
	}

	return nil
}

func uint64FromStr(raw string) (uint64, error) {
	v, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to parse uint64 from %s", raw)
	}

	return v, nil
}
