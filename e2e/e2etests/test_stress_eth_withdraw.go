package e2etests

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/protocol-contracts/pkg/gatewayzevm.sol"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

// TestStressEtherWithdraw uses the Gateway contract's withdraw function to
// perform multiple withdrawals in parallel.
func TestStressEtherWithdraw(r *runner.E2ERunner, args []string) {
    require.Len(r, args, 2)

    // Log the args at the beginning
    r.Logger.Info(fmt.Sprintf("Starting TestStressEtherWithdraw, args=%v", args))

    // Increase gas limit for this test
    previousGasLimit := r.ZEVMAuth.GasLimit
    r.ZEVMAuth.GasLimit = 10000000
    defer func() {
        // Restore the original gas limit
        r.Logger.Info(fmt.Sprintf("Restoring original gas limit: %d", previousGasLimit))
        r.ZEVMAuth.GasLimit = previousGasLimit
    }()

    // Parse the arguments
    amount := utils.ParseBigInt(r, args[0])
    payload := randomPayload(r) // randomPayload returns a string

    // Log parsed inputs (convert payload to []byte for hex encoding)
    r.Logger.Info(fmt.Sprintf("Parsed test inputs: amount=%s, payloadHex=%s",
        amount.String(),
        hex.EncodeToString([]byte(payload)),
    ))

    // Initially verify that the dApp has not been called
    r.AssertTestDAppEVMCalled(false, payload, amount)

    // Approve ETHZRC20 for the Gateway contract
    r.Logger.Info(fmt.Sprintf("Approving ETHZRC20 for Gateway, gatewayAddress=%s", r.GatewayZEVMAddr.String()))
    r.ApproveETHZRC20(r.GatewayZEVMAddr)

    // Perform the withdrawal with ETHWithdrawAndCall
    r.Logger.Info(fmt.Sprintf("Calling ETHWithdrawAndCall, target=%s, amount=%s, payloadHex=%s",
        r.TestDAppV2EVMAddr.String(),
        amount.String(),
        hex.EncodeToString([]byte(payload)),
    ))
    tx := r.ETHWithdrawAndCall(
        r.TestDAppV2EVMAddr,
        amount,
        []byte(payload),
        gatewayzevm.RevertOptions{OnRevertGasLimit: big.NewInt(0)},
    )

    r.Logger.Info(fmt.Sprintf("Withdraw transaction sent, txHash=%s", tx.Hash().Hex()))

    // Wait for the CCTX to be mined
    r.Logger.Info(fmt.Sprintf("Waiting for CCTX to be mined, txHash=%s", tx.Hash().Hex()))
    cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)

    // Use the correct field from CrossChainTx (replace cctx.Index with whichever field is correct)
    r.Logger.Info(fmt.Sprintf("CCTX mined, cctxIndex=%s, status=%v",
        cctx.Index,
        cctx.CctxStatus.Status,
    ))

    // Log with the dedicated CCTX logger (optional/custom logic)
    r.Logger.CCTX(*cctx, "withdraw")

    // Ensure the transaction is in the expected mined status
    require.Equal(r, crosschaintypes.CctxStatus_OutboundMined, cctx.CctxStatus.Status)

    // Assert the DApp was called after withdrawal
    r.Logger.Info("Verifying TestDAppEVM was called after withdrawal")
    r.AssertTestDAppEVMCalled(true, payload, amount)

    // Check that the correct sender was recorded in the DApp
    r.Logger.Info("Checking sender for payload on TestDAppV2EVM")
    senderForMsg, err := r.TestDAppV2EVM.SenderWithMessage(&bind.CallOpts{}, []byte(payload))
    require.NoError(r, err)

    r.Logger.Info(fmt.Sprintf("Comparing expected sender with the contract's recognized sender: expected=%s, actual=%s",
        r.ZEVMAuth.From.String(),
        senderForMsg.String(),
    ))
    require.Equal(r, r.ZEVMAuth.From, senderForMsg)

    r.Logger.Info("TestStressEtherWithdraw completed successfully")
}
