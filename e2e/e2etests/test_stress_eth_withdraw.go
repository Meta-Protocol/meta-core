package e2etests

import (
	"fmt"
	"math/big"

	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/protocol-contracts/pkg/gatewayzevm.sol"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

func TestStressEtherWithdraw(r *runner.E2ERunner, args []string) {
		require.Len(r, args, 2)

		// Start of test logging
		fmt.Println("Starting TestETHWithdraw, args=", args)

		// Parse the withdraw amount
		amount := utils.ParseBigInt(r, args[0])
		fmt.Println("Parsed withdraw amount =", amount.String())

		// Retrieve old balance for comparison
		oldBalance, err := r.EVMClient.BalanceAt(r.Ctx, r.EVMAddress(), nil)
		require.NoError(r, err)
		fmt.Println("Old balance for address", r.EVMAddress().Hex(), "=", oldBalance.String())

		// Approve ETHZRC20 for the Gateway contract
		fmt.Println("Approving ETHZRC20 for Gateway, gatewayAddress=", r.GatewayZEVMAddr.Hex())
		r.ApproveETHZRC20(r.GatewayZEVMAddr)

		// Perform the withdraw
		fmt.Println("Performing ETHWithdraw, amount=", amount.String())
		tx := r.ETHWithdraw(r.EVMAddress(), amount, gatewayzevm.RevertOptions{OnRevertGasLimit: big.NewInt(0)})

		fmt.Println("Withdraw transaction broadcasted, txHash=", tx.Hash().Hex())

		// Wait for the cctx to be mined
		fmt.Println("Waiting for CCTX to be mined, inboundHash=", tx.Hash().Hex())
		cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)

		// Log with your existing CCTX logger (optional/custom)
		r.Logger.CCTX(*cctx, "withdraw")

		// Log the CCTX status
		fmt.Println("CCTX mined with status=", cctx.CctxStatus.Status)
		require.Equal(r, crosschaintypes.CctxStatus_OutboundMined, cctx.CctxStatus.Status)

		// Retrieve the new balance
		newBalance, err := r.EVMClient.BalanceAt(r.Ctx, r.EVMAddress(), nil)
		require.NoError(r, err)
		fmt.Println("New balance for address", r.EVMAddress().Hex(), "=", newBalance.String())

		// Ensure the new balance is greater (minus gas fees)
		require.Greater(r, newBalance.Uint64(), oldBalance.Uint64())

		// Test end
		fmt.Println("TestETHWithdraw completed successfully")
}
