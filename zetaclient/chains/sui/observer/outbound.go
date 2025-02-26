package observer

import (
	"context"
	"strconv"

	"cosmossdk.io/math"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/pkg/errors"

	"github.com/zeta-chain/node/pkg/chains"
	"github.com/zeta-chain/node/pkg/coin"
	"github.com/zeta-chain/node/pkg/contracts/sui"
	cctypes "github.com/zeta-chain/node/x/crosschain/types"
	"github.com/zeta-chain/node/zetaclient/chains/interfaces"
	"github.com/zeta-chain/node/zetaclient/logs"
	"github.com/zeta-chain/node/zetaclient/zetacore"
)

// https://github.com/zeta-chain/protocol-contracts-sui/blob/9d08a70817d8cc7cf799b9ae12c59b6e0b8aaab9/sources/gateway.move#L125
// (excluding last arg of `ctx`)
const expectedWithdrawArgs = 5

// 50 SUI
// https://docs.sui.io/concepts/tokenomics/gas-in-sui#gas-budgets
const maxGasLimit = 50_000_000_000

// OutboundCreated checks if the outbound tx exists in the memory
// and has valid nonce & signature
func (ob *Observer) OutboundCreated(nonce uint64) bool {
	_, ok := ob.getTx(nonce)
	return ok
}

// ProcessOutboundTrackers loads all freshly-included Sui transactions in-memory
// for further voting by Observer-Signer.
func (ob *Observer) ProcessOutboundTrackers(ctx context.Context) error {
	chainID := ob.Chain().ChainId

	trackers, err := ob.ZetacoreClient().GetAllOutboundTrackerByChain(ctx, chainID, interfaces.Ascending)
	if err != nil {
		return errors.Wrap(err, "unable to get outbound trackers")
	}

	for _, tracker := range trackers {
		nonce := tracker.Nonce

		// should not happen
		if len(tracker.HashList) == 0 {
			return errors.Errorf("empty outbound tracker hash for nonce %d", nonce)
		}

		// already loaded
		if _, ok := ob.getTx(nonce); ok {
			continue
		}

		digest := tracker.HashList[0].TxHash

		cctx, err := ob.ZetacoreClient().GetCctxByNonce(ctx, chainID, tracker.Nonce)
		if err != nil {
			return errors.Wrapf(err, "unable to get cctx by nonce %d (sui digest %q)", tracker.Nonce, digest)
		}

		if err := ob.loadOutboundTx(ctx, cctx, digest); err != nil {
			// we don't want to block other cctxs, so let's error and continue
			ob.Logger().Outbound.
				Error().Err(err).
				Str(logs.FieldMethod, "ProcessOutboundTrackers").
				Uint64(logs.FieldNonce, nonce).
				Str(logs.FieldTx, digest).
				Msg("Unable to load outbound transaction")
		}
	}

	return nil
}

// VoteOutbound calculates outbound result based on cctx and in-mem Sui tx
// and votes the ballot to zetacore.
func (ob *Observer) VoteOutbound(ctx context.Context, cctx *cctypes.CrossChainTx) error {
	chainID := ob.Chain().ChainId
	nonce := cctx.GetCurrentOutboundParam().TssNonce

	// should be fetched by ProcessOutboundTrackers routine
	// if exists, we can safely assume it's authentic and nonce is valid
	tx, ok := ob.getTx(nonce)
	if !ok {
		return errors.Errorf("missing tx for nonce %d", nonce)
	}

	// used instead of block height
	checkpoint, err := strconv.ParseUint(tx.Checkpoint, 10, 64)
	if err != nil {
		return errors.Wrap(err, "unable to parse checkpoint")
	}

	// parse status, coinType, and amount
	var (
		status    = chains.ReceiveStatus_failed
		coinType  = coin.CoinType_Gas
		amount    = math.NewUint(0)
		isSuccess = tx.Effects.Status.Status == "success"
	)

	if isSuccess {
		status = chains.ReceiveStatus_success

		_, w, err := ob.gateway.ParseTxWithdrawal(tx)
		if err != nil {
			return errors.Wrap(err, "unable to parse tx withdrawal")
		}

		if !w.IsGas() {
			coinType = coin.CoinType_ERC20
		}

		amount = w.Amount
	}

	// Gas parameters
	// Gas price *might* change once per epoch (~24h), so using the latest value is fine.
	// #nosec G115 - always in range
	outboundGasPrice := math.NewInt(int64(ob.getLatestGasPrice()))

	// This might happen after zetacore restart when PostGasPrice has not been called yet. retry later.
	if outboundGasPrice.IsZero() {
		return errors.New("latest gas price is zero")
	}

	outboundGasUsed, err := parseGasUsed(tx)
	if err != nil {
		return errors.Wrap(err, "unable to parse gas used")
	}

	// Create message
	msg := cctypes.NewMsgVoteOutbound(
		ob.ZetacoreClient().GetKeys().GetOperatorAddress().String(),
		cctx.Index,
		tx.Digest,
		checkpoint,
		outboundGasUsed,
		outboundGasPrice,
		maxGasLimit,
		amount,
		status,
		chainID,
		nonce,
		coinType,
		cctypes.ConfirmationMode_SAFE,
	)

	// TODO compliance checks
	// https://github.com/zeta-chain/node/issues/3584

	if err := ob.postVoteOutbound(ctx, msg); err != nil {
		return errors.Wrap(err, "unable to post vote outbound")
	}

	ob.unsetTx(nonce)

	return nil
}

// loadOutboundTx loads cross-chain outbound tx by digest and ensures its authenticity.
func (ob *Observer) loadOutboundTx(ctx context.Context, cctx *cctypes.CrossChainTx, digest string) error {
	res, err := ob.client.SuiGetTransactionBlock(ctx, models.SuiGetTransactionBlockRequest{
		Digest: digest,
		Options: models.SuiTransactionBlockOptions{
			ShowEvents:  true,
			ShowInput:   true,
			ShowEffects: true,
		},
	})

	if err != nil {
		return errors.Wrap(err, "unable to get tx")
	}

	if err := ob.validateOutbound(cctx, res); err != nil {
		return errors.Wrap(err, "tx validation failed")
	}

	ob.setTx(res, cctx.GetCurrentOutboundParam().TssNonce)

	return nil
}

// validateOutbound validates the authenticity of the outbound transaction.
// Note that it doesn't care about successful execution (e.g. something failed).
func (ob *Observer) validateOutbound(cctx *cctypes.CrossChainTx, tx models.SuiTransactionBlockResponse) error {
	nonce := cctx.GetCurrentOutboundParam().TssNonce

	inputs := tx.Transaction.Data.Transaction.Inputs

	// Check args length
	if len(inputs) != expectedWithdrawArgs {
		return errors.Errorf("invalid number of input arguments (got %d, want %d)", len(inputs), expectedWithdrawArgs)
	}

	txNonce, err := parseNonceFromWithdrawInputs(inputs)
	if err != nil {
		return errors.Wrap(err, "unable to parse nonce from inputs")
	}

	if txNonce != nonce {
		return errors.Errorf("nonce mismatch (tx nonce %d, cctx nonce %d)", txNonce, nonce)
	}

	if len(tx.Transaction.TxSignatures) == 0 {
		return errors.New("missing tx signature")
	}

	pubKey, _, err := sui.DeserializeSignatureECDSA(tx.Transaction.TxSignatures[0])
	if err != nil {
		return errors.Wrap(err, "unable to deserialize tx signature")
	}

	if !ob.TSS().PubKey().AsECDSA().Equal(pubKey) {
		return errors.New("pubKey mismatch")
	}

	return nil
}

func (ob *Observer) postVoteOutbound(ctx context.Context, msg *cctypes.MsgVoteOutbound) error {
	const gasLimit = zetacore.PostVoteOutboundGasLimit

	retryGasLimit := uint64(0)
	if msg.Status == chains.ReceiveStatus_failed {
		retryGasLimit = zetacore.PostVoteOutboundRevertGasLimit
	}

	zetaTxHash, ballot, err := ob.ZetacoreClient().PostVoteOutbound(ctx, gasLimit, retryGasLimit, msg)
	switch {
	case err != nil:
		return errors.Wrap(err, "unable to post vote outbound")
	case zetaTxHash != "":
		ob.Logger().Outbound.Info().
			Str(logs.FieldZetaTx, zetaTxHash).
			Str(logs.FieldBallot, ballot).
			Msg("PostVoteOutbound: posted outbound vote successfully")
	}

	return nil
}

func (ob *Observer) getTx(nonce uint64) (models.SuiTransactionBlockResponse, bool) {
	ob.txMu.RLock()
	defer ob.txMu.RUnlock()

	tx, ok := ob.txMap[nonce]

	return tx, ok
}

func (ob *Observer) setTx(tx models.SuiTransactionBlockResponse, nonce uint64) {
	ob.txMu.Lock()
	defer ob.txMu.Unlock()

	ob.txMap[nonce] = tx
}

func (ob *Observer) unsetTx(nonce uint64) {
	ob.txMu.Lock()
	defer ob.txMu.Unlock()

	delete(ob.txMap, nonce)
}

func parseNonceFromWithdrawInputs(inputs []models.SuiCallArg) (uint64, error) {
	if len(inputs) != expectedWithdrawArgs {
		return 0, errors.New("invalid number of input arguments")
	}

	const nonceIdx = 2

	// {
	//   "type": "pure",
	//   "valueType": "u64",
	//   "value": "12345"
	// }
	raw := inputs[nonceIdx]

	if raw["type"] != "pure" || raw["valueType"] != "u64" {
		return 0, errors.Errorf("invalid nonce object %+v", raw)
	}

	return strconv.ParseUint(raw["value"].(string), 10, 64)
}

func parseGasUsed(tx models.SuiTransactionBlockResponse) (uint64, error) {
	gas := tx.Effects.GasUsed

	compCost, err := parseUint64(gas.ComputationCost)
	if err != nil {
		return 0, errors.Wrap(err, "comp cost")
	}

	storageCost, err := parseUint64(gas.StorageCost)
	if err != nil {
		return 0, errors.Wrap(err, "storage cost")
	}

	storageRebate, err := parseUint64(gas.StorageRebate)
	if err != nil {
		return 0, errors.Wrap(err, "storage rebate")
	}

	// should not happen
	if (compCost + storageCost) < storageRebate {
		return 0, errors.New("storage rebate exceeds total costs")
	}

	return compCost + storageCost - storageRebate, nil
}

func parseUint64(v string) (uint64, error) {
	return strconv.ParseUint(v, 10, 64)
}
