package e2etests

import (
	"strconv"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

// TestBitcoinWithdrawRBF tests the RBF (Replace-By-Fee) feature in Zetaclient.
// It needs block mining to be stopped and runs as the last test in the suite.
//
// IMPORTANT: the test requires to simulate a stuck tx in the Bitcoin regnet.
// Changing the 'minTxConfirmations' to 1 to not include Bitcoin pending txs.
// https://github.com/zeta-chain/node/blob/feat-bitcoin-Replace-By-Fee/zetaclient/chains/bitcoin/observer/outbound.go#L30
func TestBitcoinWithdrawRBF(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 2)

	// parse arguments
	defaultReceiver := r.BTCDeployerAddress.EncodeAddress()
	to, amount := utils.ParseBitcoinWithdrawArgs(r, args, defaultReceiver, r.GetBitcoinChainID())

	// initiate a withdraw CCTX
	receipt := approveAndWithdrawBTCZRC20(r, to, amount)
	cctx := utils.GetCCTXByInboundHash(r.Ctx, r.CctxClient, receipt.TxHash.Hex())

	// wait for the 1st outbound tracker hash to come in
	nonce := cctx.GetCurrentOutboundParam().TssNonce
	hashes := utils.WaitOutboundTracker(r.Ctx, r.CctxClient, r.GetBitcoinChainID(), nonce, 1, r.Logger, 3*time.Minute)
	txHash, err := chainhash.NewHashFromStr(hashes[0])
	r.Logger.Info("got 1st tracker hash: %s", txHash)

	// get original tx
	require.NoError(r, err)
	txResult, err := r.BtcRPCClient.GetTransaction(r.Ctx, txHash)
	require.NoError(r, err)
	require.Zero(r, txResult.Confirmations)

	// wait for RBF tx to kick in
	hashes = utils.WaitOutboundTracker(r.Ctx, r.CctxClient, r.GetBitcoinChainID(), nonce, 2, r.Logger, 3*time.Minute)
	txHashRBF, err := chainhash.NewHashFromStr(hashes[1])
	require.NoError(r, err)
	r.Logger.Info("got 2nd tracker hash: %s", txHashRBF)

	// resume block mining
	stop := r.MineBlocksIfLocalBitcoin()
	defer stop()

	// waiting for CCTX to be mined
	cctx = utils.WaitCctxMinedByInboundHash(r.Ctx, receipt.TxHash.Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	utils.RequireCCTXStatus(r, cctx, crosschaintypes.CctxStatus_OutboundMined)

	// ensure the original tx is dropped
	utils.MustHaveDroppedTx(r.Ctx, r.BtcRPCClient, txHash)

	// ensure the RBF tx is mined
	rawResult := utils.MustHaveMinedTx(r.Ctx, r.BtcRPCClient, txHashRBF)

	// ensure RBF fee rate > old rate
	params := cctx.GetCurrentOutboundParam()
	oldRate, err := strconv.ParseInt(params.GasPrice, 10, 64)
	require.NoError(r, err)

	_, newRate, err := r.BtcRPCClient.GetTransactionFeeAndRate(r.Ctx, rawResult)
	require.NoError(r, err)
	require.Greater(r, newRate, oldRate, "RBF fee rate should be higher than the original tx")
}
