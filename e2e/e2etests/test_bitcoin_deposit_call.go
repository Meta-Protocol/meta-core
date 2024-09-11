package e2etests

import (
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/zetacore/e2e/runner"
	"github.com/zeta-chain/zetacore/e2e/utils"
	testcontract "github.com/zeta-chain/zetacore/testutil/contracts"
	crosschaintypes "github.com/zeta-chain/zetacore/x/crosschain/types"
	zetabitcoin "github.com/zeta-chain/zetacore/zetaclient/chains/bitcoin"
)

func TestBitcoinDepositAndCall(r *runner.E2ERunner, args []string) {
	// ARRANGE
	// Given BTC address
	r.SetBtcAddress(r.Name, false)

	// Given "Live" BTC network
	stop := r.MineBlocksIfLocalBitcoin()
	defer stop()

	// Given amount to send
	require.Len(r, args, 1)
	amount := parseFloat(r, args[0])
	amountTotal := amount + zetabitcoin.DefaultDepositorFee

	// Given a list of UTXOs
	utxos, err := r.ListDeployerUTXOs()
	require.NoError(r, err)
	require.NotEmpty(r, utxos)

	// deploy an example contract in ZEVM
	contractAddr, _, _, err := testcontract.DeployExample(r.ZEVMAuth, r.ZEVMClient)
	require.NoError(r, err)
	r.Logger.Print("Bitcoin: Example contract deployed at: %s", contractAddr.String())

	// ACT
	// Send BTC to TSS address with a dummy memo
	data := []byte("hello shatoshi")
	memo := append(contractAddr.Bytes(), data...)
	txHash, err := r.SendToTSSFromDeployerWithMemo(amountTotal, utxos, memo)
	require.NoError(r, err)
	require.NotEmpty(r, txHash)
	r.Logger.Print("Bitcoin: Sent %f BTC to TSS address", amountTotal)

	// wait for the cctx to be mined
	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, txHash.String(), r.CctxClient, r.Logger, r.CctxTimeout)
	r.Logger.CCTX(*cctx, "bitcoin_deposit_and_call")
	utils.RequireCCTXStatus(r, cctx, crosschaintypes.CctxStatus_OutboundMined)
	r.Logger.Print("Bitcoin: CCTX mined")

	// check if example contract has been called, 'bar' value should be set to amount
	// amoutSats, err := zetabitcoin.GetSatoshis(amount)
	// require.NoError(r, err)
	// utils.MustHaveCalledExampleContract(r, contract, big.NewInt(amoutSats))
}
