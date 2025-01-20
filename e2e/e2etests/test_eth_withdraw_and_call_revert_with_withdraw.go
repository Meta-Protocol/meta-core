package e2etests

import (
	"math/big"

	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/protocol-contracts/v2/pkg/gatewayzevm.sol"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

func TestETHWithdrawAndCallRevertWithWithdraw(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 1)

	amount := utils.ParseBigInt(r, args[0])

	r.ApproveETHZRC20(r.GatewayZEVMAddr)

	// perform the withdraw
	tx := r.ETHWithdrawAndArbitraryCall(
		r.TestDAppV2EVMAddr,
		amount,
		r.EncodeGasCall("revert"),
		gatewayzevm.RevertOptions{
			RevertAddress:    r.TestDAppV2ZEVMAddr,
			CallOnRevert:     true,
			RevertMessage:    []byte("withdraw"), // call withdraw in the onRevert hook
			OnRevertGasLimit: big.NewInt(0),
		},
	)

	// wait for the cctx to be mined
	cctxRevert := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	r.Logger.CCTX(*cctxRevert, "withdraw")
	require.Equal(r, crosschaintypes.CctxStatus_Reverted, cctxRevert.CctxStatus.Status)

	r.Logger.Print("cctxRevert")
	r.Logger.Print(cctxRevert.String())
	cctxWithdrawFromRevert := utils.WaitCctxMinedByInboundHash(r.Ctx, cctxRevert.Index, r.CctxClient, r.Logger, r.CctxTimeout)

	//check the cctx status
	utils.RequireCCTXStatus(r, cctxWithdrawFromRevert, crosschaintypes.CctxStatus_OutboundMined)
}
