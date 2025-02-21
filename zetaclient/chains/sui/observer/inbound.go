package observer

import (
	"context"
	"encoding/hex"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/pkg/errors"

	"github.com/zeta-chain/node/pkg/coin"
	"github.com/zeta-chain/node/pkg/contracts/sui"
	cctypes "github.com/zeta-chain/node/x/crosschain/types"
	"github.com/zeta-chain/node/zetaclient/chains/sui/client"
	"github.com/zeta-chain/node/zetaclient/logs"
	"github.com/zeta-chain/node/zetaclient/zetacore"
)

var errTxNotFound = errors.New("no tx found")

// ObserveInbound processes inbound deposit cross-chain transactions.
func (ob *Observer) ObserveInbound(ctx context.Context) error {
	if err := ob.ensureCursor(ctx); err != nil {
		return errors.Wrap(err, "unable to ensure inbound cursor")
	}

	ob.Logger().Inbound.Info().Msg("Cursor: " + ob.getCursor())

	query := client.EventQuery{
		PackageID: ob.gateway.PackageID(),
		Module:    ob.gateway.Module(),
		// TODO: fix issue with cursor
		Cursor: "", //ob.getCursor(),
		Limit:  client.DefaultEventsLimit,
	}

	// Sui has a nice access-pattern of scrolling through contract events
	events, _, err := ob.client.QueryModuleEvents(ctx, query)
	if err != nil {
		return errors.Wrap(err, "unable to query module events")
	}

	ob.Logger().Inbound.Info().Int("events", len(events)).Msg("Processing sui inbound events")

	for _, event := range events {
		// Note: we can make this concurrent if needed.
		// Let's revisit later
		err := ob.processInboundEvent(ctx, event, nil)

		switch {
		case errors.Is(err, errTxNotFound):
			// try again later
			ob.Logger().Inbound.Warn().Err(err).
				Str(logs.FieldTx, event.Id.TxDigest).
				Msg("TX not found or unfinalized. Pausing")
			return nil
		case err != nil:
			// failed processing also updates the cursor
			ob.Logger().Inbound.Err(err).
				Str(logs.FieldTx, event.Id.TxDigest).
				Msg("Unable to process inbound event")
		}

		// update the cursor
		if err := ob.setCursor(client.EncodeCursor(event.Id)); err != nil {
			return errors.Wrapf(err, "unable to set cursor %+v", event.Id)
		}
	}

	return nil
}

// ProcessInboundTrackers processes trackers for inbound transactions.
func (ob *Observer) ProcessInboundTrackers(ctx context.Context) error {
	chainID := ob.Chain().ChainId

	trackers, err := ob.ZetacoreClient().GetInboundTrackersForChain(ctx, chainID)
	if err != nil {
		return errors.Wrap(err, "unable to get inbound trackers")
	}

	for _, tracker := range trackers {
		if err := ob.processInboundTracker(ctx, tracker); err != nil {
			ob.Logger().Inbound.Err(err).
				Str(logs.FieldTx, tracker.TxHash).
				Msg("Unable to process inbound tracker")
		}
	}

	return nil
}

// processInboundEvent parses raw event into Inbound, augments it with origin tx and votes on the inbound.
// - Invalid/Non-inbound txs are skipped.
// - Unconfirmed txs pause the whole tail sequence.
// - If tx is empty, it fetches the tx from RPC.
// - Sui tx is finalized if it's returned from RPC
//
// See https://docs.sui.io/concepts/sui-architecture/transaction-lifecycle#verifying-finality
func (ob *Observer) processInboundEvent(
	ctx context.Context,
	raw models.SuiEventResponse,
	tx *models.SuiTransactionBlockResponse,
) error {
	event, err := ob.gateway.ParseEvent(raw)
	switch {
	case errors.Is(err, sui.ErrParseEvent):
		ob.Logger().Inbound.Err(err).Msg("Unable to parse event. Skipping")
		return nil
	case err != nil:
		return errors.Wrap(err, "unable to parse event")
	case !event.IsInbound():
		ob.Logger().Inbound.Info().Msg("Not an inbound event. Skipping")
	case event.EventIndex != 0:
		// Is it possible to have multiple events per tx?
		// e.g. contract "A" calls Gateway multiple times in a single tx (deposit to multiple accounts)
		// most likely not, so let's explicitly fail to prevent undefined behavior.
		return errors.Errorf("unexpected event index %d for tx %s", event.EventIndex, event.TxHash)
	}

	if tx == nil {
		txReq := models.SuiGetTransactionBlockRequest{Digest: event.TxHash}
		txFresh, err := ob.client.SuiGetTransactionBlock(ctx, txReq)
		if err != nil {
			return errors.Wrap(errTxNotFound, err.Error())
		}

		tx = &txFresh
	}

	msg, err := ob.constructInboundVote(event, *tx)
	if err != nil {
		return errors.Wrap(err, "unable to construct inbound vote")
	}

	_, err = ob.PostVoteInbound(ctx, msg, zetacore.PostVoteInboundExecutionGasLimit)
	if err != nil {
		return errors.Wrap(err, "unable to post vote inbound")
	}

	return nil
}

// processInboundTracker queries tx with its events by tracker and then votes.
func (ob *Observer) processInboundTracker(ctx context.Context, tracker cctypes.InboundTracker) error {
	req := models.SuiGetTransactionBlockRequest{
		Digest:  tracker.TxHash,
		Options: models.SuiTransactionBlockOptions{ShowEvents: true},
	}

	tx, err := ob.client.SuiGetTransactionBlock(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "unable to get transaction block")
	}

	for _, event := range tx.Events {
		if err := ob.processInboundEvent(ctx, event, &tx); err != nil {
			return errors.Wrapf(err, "unable to process inbound event %s", event.Id.EventSeq)
		}
	}

	return nil
}

// constructInboundVote creates a vote message for inbound deposit
func (ob *Observer) constructInboundVote(
	event sui.Event,
	tx models.SuiTransactionBlockResponse,
) (*cctypes.MsgVoteInbound, error) {
	inbound, err := event.Inbound()
	if err != nil {
		return nil, errors.Wrap(err, "unable to extract inbound")
	}

	coinType := coin.CoinType_Gas
	if !inbound.IsGasDeposit() {
		coinType = coin.CoinType_ERC20
	}

	// Sui uses checkpoint seq num instead of block height
	checkpointSeqNum, err := uint64FromStr(tx.Checkpoint)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse checkpoint")
	}

	// Empty or full SUI coin name
	var asset string
	if !inbound.IsGasDeposit() {
		asset = string(inbound.CoinType)
	}

	return cctypes.NewMsgVoteInbound(
		ob.ZetacoreClient().GetKeys().GetOperatorAddress().String(),
		inbound.Sender,
		ob.Chain().ChainId,
		inbound.Sender,
		inbound.Receiver.String(),
		ob.ZetacoreClient().Chain().ChainId,
		inbound.Amount,
		hex.EncodeToString(inbound.Payload),
		event.TxHash,
		checkpointSeqNum,
		zetacore.PostVoteInboundCallOptionsGasLimit,
		coinType,
		asset,
		event.EventIndex,
		cctypes.ProtocolContractVersion_V2,
		false,
		cctypes.InboundStatus_SUCCESS,
		cctypes.ConfirmationMode_SAFE,
		cctypes.WithCrossChainCall(inbound.IsCrossChainCall),
	), nil
}
