package observer

import (
	"context"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
	"github.com/tonkeeper/tongo/liteclient"
	"github.com/tonkeeper/tongo/tlb"
	"github.com/tonkeeper/tongo/ton"

	toncontracts "github.com/zeta-chain/node/pkg/contracts/ton"
	"github.com/zeta-chain/node/zetaclient/chains/base"
	"github.com/zeta-chain/node/zetaclient/chains/ton/config"
)

// Observer is a TON observer.
type Observer struct {
	*base.Observer

	client  LiteClient
	gateway *toncontracts.Gateway

	outbounds *lru.Cache
}

const outboundsCacheSize = 1024

// LiteClient represents a TON client
// see https://github.com/ton-blockchain/ton/blob/master/tl/generate/scheme/tonlib_api.tl
//
//go:generate mockery --name LiteClient --filename ton_liteclient.go --case underscore --output ../../../testutils/mocks
type LiteClient interface {
	config.Getter
	GetMasterchainInfo(ctx context.Context) (liteclient.LiteServerMasterchainInfoC, error)
	GetBlockHeader(ctx context.Context, blockID ton.BlockIDExt, mode uint32) (tlb.BlockInfo, error)
	GetTransactionsSince(ctx context.Context, acc ton.AccountID, lt uint64, hash ton.Bits256) ([]ton.Transaction, error)
	GetFirstTransaction(ctx context.Context, acc ton.AccountID) (*ton.Transaction, int, error)
	GetTransaction(ctx context.Context, acc ton.AccountID, lt uint64, hash ton.Bits256) (ton.Transaction, error)
}

// New constructor for TON Observer.
func New(bo *base.Observer, client LiteClient, gateway *toncontracts.Gateway) (*Observer, error) {
	switch {
	case !bo.Chain().IsTONChain():
		return nil, errors.New("base observer chain is not TON")
	case client == nil:
		return nil, errors.New("liteapi client is nil")
	case gateway == nil:
		return nil, errors.New("gateway is nil")
	}

	outbounds, err := lru.New(outboundsCacheSize)
	if err != nil {
		return nil, err
	}

	bo.LoadLastTxScanned()

	return &Observer{
		Observer:  bo,
		client:    client,
		gateway:   gateway,
		outbounds: outbounds,
	}, nil
}

// PostGasPrice fetches on-chain gas config and reports it to Zetacore.
func (ob *Observer) PostGasPrice(ctx context.Context) error {
	cfg, err := config.FetchGasConfig(ctx, ob.client)
	if err != nil {
		return errors.Wrap(err, "failed to fetch gas config")
	}

	gasPrice, err := config.ParseGasPrice(cfg)
	if err != nil {
		return errors.Wrap(err, "failed to parse gas price")
	}

	blockID, err := ob.getLatestMasterchainBlock(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get latest masterchain block")
	}

	// There's no concept of priority fee in TON
	const priorityFee = 0

	_, errVote := ob.
		ZetacoreClient().
		PostVoteGasPrice(ctx, ob.Chain(), gasPrice, priorityFee, uint64(blockID.Seqno))

	return errVote
}

// CheckRPCStatus checks TON RPC status and alerts if necessary.
func (ob *Observer) CheckRPCStatus(ctx context.Context) error {
	blockID, err := ob.getLatestMasterchainBlock(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get latest masterchain block")
	}

	block, err := ob.client.GetBlockHeader(ctx, blockID, 0)
	switch {
	case err != nil:
		return errors.Wrap(err, "failed to get masterchain block header")
	case block.NotMaster:
		return errors.Errorf("block %q is not a master block", blockID.BlockID.String())
	}

	blockTime := time.Unix(int64(block.GenUtime), 0).UTC()

	ob.ReportBlockLatency(blockTime)

	return nil
}

func (ob *Observer) getLatestMasterchainBlock(ctx context.Context) (ton.BlockIDExt, error) {
	mc, err := ob.client.GetMasterchainInfo(ctx)
	if err != nil {
		return ton.BlockIDExt{}, errors.Wrap(err, "failed to get masterchain info")
	}

	return mc.Last.ToBlockIdExt(), nil
}
