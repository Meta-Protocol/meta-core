package observer

import (
	"context"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/pkg/errors"

	"github.com/zeta-chain/node/pkg/coin"
	"github.com/zeta-chain/node/zetaclient/chains/bitcoin/common"
	"github.com/zeta-chain/node/zetaclient/zetacore"
)

// ProcessInboundTrackers processes inbound trackers
func (ob *Observer) ProcessInboundTrackers(ctx context.Context) error {
	trackers, err := ob.ZetacoreClient().GetInboundTrackersForChain(ctx, ob.Chain().ChainId)
	if err != nil {
		return err
	}

	for _, tracker := range trackers {
		ob.logger.Inbound.Info().
			Str("tracker.hash", tracker.TxHash).
			Str("tracker.coin-type", tracker.CoinType.String()).
			Msgf("checking tracker")
		ballotIdentifier, err := ob.CheckReceiptForBtcTxHash(ctx, tracker.TxHash, true)
		if err != nil {
			return err
		}
		ob.logger.Inbound.Info().
			Str("inbound.chain", ob.Chain().Name).
			Str("inbound.ballot", ballotIdentifier).
			Str("inbound.coin-type", coin.CoinType_Gas.String()).
			Msgf("Vote submitted for inbound Tracker")
	}

	return nil
}

// CheckReceiptForBtcTxHash checks the receipt for a btc tx hash
func (ob *Observer) CheckReceiptForBtcTxHash(ctx context.Context, txHash string, vote bool) (string, error) {
	hash, err := chainhash.NewHashFromStr(txHash)
	if err != nil {
		return "", err
	}

	tx, err := ob.rpc.GetRawTransactionVerbose(ctx, hash)
	if err != nil {
		return "", err
	}

	blockHash, err := chainhash.NewHashFromStr(tx.BlockHash)
	if err != nil {
		return "", err
	}

	blockVb, err := ob.rpc.GetBlockVerbose(ctx, blockHash)
	if err != nil {
		return "", err
	}

	if len(blockVb.Tx) <= 1 {
		return "", fmt.Errorf("block %d has no transactions", blockVb.Height)
	}

	tss, err := ob.ZetacoreClient().GetBTCTSSAddress(ctx, ob.Chain().ChainId)
	if err != nil {
		return "", err
	}

	// check confirmation
	// #nosec G115 block height always positive
	if !ob.IsBlockConfirmedForInboundSafe(uint64(blockVb.Height)) {
		return "", fmt.Errorf("block %d is not confirmed yet", blockVb.Height)
	}

	// #nosec G115 always positive
	event, err := GetBtcEvent(
		ctx,
		ob.rpc,
		*tx,
		tss,
		uint64(blockVb.Height),
		ob.logger.Inbound,
		ob.netParams,
		common.CalcDepositorFee,
	)
	if err != nil {
		return "", err
	}

	if event == nil {
		return "", errors.New("no btc deposit event found")
	}

	msg := ob.GetInboundVoteFromBtcEvent(event)
	if msg == nil {
		return "", errors.New("no message built for btc sent to TSS")
	}

	if !vote {
		return msg.Digest(), nil
	}

	return ob.PostVoteInbound(ctx, msg, zetacore.PostVoteInboundExecutionGasLimit)
}
