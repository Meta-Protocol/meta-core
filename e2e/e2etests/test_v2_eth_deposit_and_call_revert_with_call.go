package e2etests

import (
	"math/big"

	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/protocol-contracts/v2/pkg/gatewayevm.sol"

	"github.com/zeta-chain/zetacore/e2e/runner"
	"github.com/zeta-chain/zetacore/e2e/utils"
	crosschaintypes "github.com/zeta-chain/zetacore/x/crosschain/types"
)

func TestV2ETHDepositAndCallRevertWithCall(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 1)

	amount, ok := big.NewInt(0).SetString(args[0], 10)
	require.True(r, ok, "Invalid amount specified for TestV2ETHDepositAndCallRevertWithCall")

	r.ApproveERC20OnEVM(r.GatewayEVMAddr)

	payload := randomPayload(r)

	r.AssertTestDAppEVMCalled(false, payload, amount)

	// perform the deposit
	tx := r.V2ETHDepositAndCall(r.TestDAppV2ZEVMAddr, amount, []byte("revert"), gatewayevm.RevertOptions{
		RevertAddress:    r.TestDAppV2EVMAddr,
		CallOnRevert:     true,
		RevertMessage:    []byte(payload),
		OnRevertGasLimit: big.NewInt(200000),
	})

	// wait for the cctx to be mined
	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	r.Logger.CCTX(*cctx, "deposit")
	require.Equal(r, crosschaintypes.CctxStatus_Reverted, cctx.CctxStatus.Status)

	// check the payload was received on the contract
	r.AssertTestDAppEVMCalled(true, payload, big.NewInt(0))
}
