package e2etests

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/protocol-contracts/pkg/gatewayzevm.sol"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

func TestStressEtherWithdraw(r *runner.E2ERunner, args []string) {
    require.Len(r, args, 1)

    // Parse the argument but never use it
    _ = utils.ParseBigInt(r, args[0])

    // Force the gas limit to a higher value temporarily
    previousGasLimit := r.ZEVMAuth.GasLimit
    r.ZEVMAuth.GasLimit = 10000000
    defer func() {
        r.ZEVMAuth.GasLimit = previousGasLimit
    }()

    // Hard-code the actual withdrawal amount to 1
    amount := big.NewInt(1)

    payload := randomPayload(r)

    // Make sure the dApp hasn't yet received an EVM call
    r.AssertTestDAppEVMCalled(false, payload, amount)

    // Approve the ETH-ZRC20 for spending by the gateway
    r.ApproveETHZRC20(r.GatewayZEVMAddr)

    // Perform exactly ONE withdraw call
    tx := r.ETHWithdrawAndCall(
        r.TestDAppV2EVMAddr,
        amount,
        []byte(payload),
        gatewayzevm.RevertOptions{OnRevertGasLimit: big.NewInt(0)},
    )

    // Wait for the cross-chain transaction to be fully mined
    cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
    r.Logger.CCTX(*cctx, "withdraw")
    require.Equal(r, crosschaintypes.CctxStatus_OutboundMined, cctx.CctxStatus.Status)

    // Verify our dApp indeed received the call
    r.AssertTestDAppEVMCalled(true, payload, amount)

    // Double-check that the sender is correct
    senderForMsg, err := r.TestDAppV2EVM.SenderWithMessage(
        &bind.CallOpts{},
        []byte(payload),
    )
    require.NoError(r, err)
    require.Equal(r, r.ZEVMAuth.From, senderForMsg)
}
