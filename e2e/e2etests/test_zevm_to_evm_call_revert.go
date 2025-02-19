package e2etests

import (
	"github.com/zeta-chain/node/testutil/sample"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/protocol-contracts/pkg/gatewayzevm.sol"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

func TestZEVMToEVMCallRevert(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 0)

	payload := randomPayload(r)

	r.ApproveETHZRC20(r.GatewayZEVMAddr)

	// perform the withdraw
	tx := r.ZEVMToEMVCall(
		sample.EthAddress(), // non-existing address
		[]byte("revert"),
		gatewayzevm.RevertOptions{
			RevertAddress:    r.TestDAppV2ZEVMAddr,
			CallOnRevert:     true,
			RevertMessage:    []byte(payload),
			OnRevertGasLimit: big.NewInt(200000),
		},
	)

	// wait for the cctx to be mined
	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	r.Logger.CCTX(*cctx, "call")
	require.Equal(r, crosschaintypes.CctxStatus_Reverted, cctx.CctxStatus.Status)

	r.AssertTestDAppZEVMCalled(true, payload, big.NewInt(0))

	// check expected sender was used
	senderForMsg, err := r.TestDAppV2ZEVM.SenderWithMessage(
		&bind.CallOpts{},
		[]byte(payload),
	)
	require.NoError(r, err)
	require.Equal(r, r.ZEVMAuth.From, senderForMsg)
}
