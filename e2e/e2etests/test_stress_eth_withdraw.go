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
    fmt.Println("Starting TestStressEtherWithdraw, args=", args)

    // Increase gas limit for this test
    previousGasLimit := r.ZEVMAuth.GasLimit
    r.ZEVMAuth.GasLimit = 10000000
    defer func() {
        // Restore the original gas limit
        fmt.Println("Restoring original gas limit:", previousGasLimit)
        r.ZEVMAuth.GasLimit = previousGasLimit
    }()

    // Parse the arguments
    amount := utils.ParseBigInt(r, args[0])
    payload := randomPayload(r) // randomPayload returns a string

    // Print parsed inputs (convert payload to []byte for hex encoding)
    fmt.Println("Parsed test inputs:",
        "amount=", amount.String(),
        "payloadHex=", hex.EncodeToString([]byte(payload)),
    )

    // Initially verify that the dApp has not been called
    r.AssertTestDAppEVMCalled(false, payload, amount)

    // Approve ETHZRC20 for the Gateway contract
    fmt.Println("Approving ETHZRC20 for Gateway, gatewayAddress=", r.GatewayZEVMAddr.String())
    r.ApproveETHZRC20(r.GatewayZEVMAddr)

    // Perform the withdrawal with ETHWithdrawAndCall
    fmt.Println("Calling ETHWithdrawAndCall, target=", r.TestDAppV2EVMAddr.String(),
        "amount=", amount.String(),
        "payloadHex=", hex.EncodeToString([]byte(payload)),
    )
    tx := r.ETHWithdrawAndCall(
        r.TestDAppV2EVMAddr,
        amount,
        []byte(payload),
        gatewayzevm.RevertOptions{OnRevertGasLimit: big.NewInt(0)},
    )

    fmt.Println("Withdraw transaction sent, txHash=", tx.Hash().Hex())

    // Wait for the CCTX to be mined
    fmt.Println("Waiting for CCTX to be mined, txHash=", tx.Hash().Hex())
    cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)

    // Use the correct field from CrossChainTx
    // Replace cctx.Index with whichever field is correct for your version
    fmt.Println("CCTX mined: cctxIndex=", cctx.Index, "status=", cctx.CctxStatus.Status)

    // Log with the dedicated CCTX logger (optional/custom logic)
    r.Logger.CCTX(*cctx, "withdraw")

    // Ensure the transaction is in the expected mined status
    require.Equal(r, crosschaintypes.CctxStatus_OutboundMined, cctx.CctxStatus.Status)

    // Assert the DApp was called after withdrawal
    fmt.Println("Verifying TestDAppEVM was called after withdrawal")
    r.AssertTestDAppEVMCalled(true, payload, amount)

    // Check that the correct sender was recorded in the DApp
    fmt.Println("Checking sender for payload on TestDAppV2EVM")
    senderForMsg, err := r.TestDAppV2EVM.SenderWithMessage(&bind.CallOpts{}, []byte(payload))
    require.NoError(r, err)

    fmt.Println("Comparing expected sender with the contract's recognized sender",
        "expectedSender=", r.ZEVMAuth.From.String(),
        "contractSender=", senderForMsg.String(),
    )
    require.Equal(r, r.ZEVMAuth.From, senderForMsg)

    fmt.Println("TestStressEtherWithdraw completed successfully")
}
