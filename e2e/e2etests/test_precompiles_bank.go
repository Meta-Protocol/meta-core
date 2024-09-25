package e2etests

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	"github.com/zeta-chain/node/precompiles/bank"
)

func TestPrecompilesBank(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 0, "No arguments expected")

	// Increase the gasLimit. It's required because of the gas consumed by precompiled functions.
	previousGasLimit := r.ZEVMAuth.GasLimit
	r.ZEVMAuth.GasLimit = 10_000_000
	defer func() {
		r.ZEVMAuth.GasLimit = previousGasLimit

		// Reset the allowance to 0; this is needed when running upgrade tests where
		// this test runs twice.
		tx, err := r.ERC20ZRC20.Approve(r.ZEVMAuth, bank.ContractAddress, big.NewInt(0))
		require.NoError(r, err)
		receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
		utils.RequireTxSuccessful(r, receipt, "Resetting allowance failed")
	}()

	totalAmount := big.NewInt(1e3)
	depositAmount := big.NewInt(500)
	higherBalanceAmount := big.NewInt(1001)
	higherAllowanceAmount := big.NewInt(501)
	spender := r.EVMAddress()

	// Get ERC20ZRC20.
	txHash := r.DepositERC20WithAmountAndMessage(r.EVMAddress(), totalAmount, []byte{})
	utils.WaitCctxMinedByInboundHash(r.Ctx, txHash.Hex(), r.CctxClient, r.Logger, r.CctxTimeout)

	// Create a bank contract caller.
	bankContract, err := bank.NewIBank(bank.ContractAddress, r.ZEVMClient)
	require.NoError(r, err, "Failed to create bank contract caller")

	// Cosmos coin balance should be 0 at this point.
	cosmosBalance, err := bankContract.BalanceOf(&bind.CallOpts{Context: r.Ctx}, r.ERC20ZRC20Addr, spender)
	require.NoError(r, err, "Call bank.BalanceOf()")
	require.Equal(r, uint64(0), cosmosBalance.Uint64(), "spender cosmos coin balance should be 0")

	// Approve allowance of 500 ERC20ZRC20 tokens for the bank contract. Should pass.
	tx, err := r.ERC20ZRC20.Approve(r.ZEVMAuth, bank.ContractAddress, depositAmount)
	require.NoError(r, err)
	receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequireTxSuccessful(r, receipt, "Approve ETHZRC20 bank allowance tx failed")

	// Deposit 501 ERC20ZRC20 tokens to the bank contract.
	// It's higher than allowance but lower than balance, should fail.
	tx, err = bankContract.Deposit(r.ZEVMAuth, r.ERC20ZRC20Addr, higherAllowanceAmount)
	require.NoError(r, err, "Call bank.Deposit() with amout higher than allowance")
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequiredTxFailed(r, receipt, "Depositting an amount higher than allowed should fail")

	// Approve allowance of 1000 ERC20ZRC20 tokens.
	tx, err = r.ERC20ZRC20.Approve(r.ZEVMAuth, bank.ContractAddress, big.NewInt(1e3))
	require.NoError(r, err)
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequireTxSuccessful(r, receipt, "Approve ETHZRC20 bank allowance tx failed")

	// Deposit 1001 ERC20ZRC20 tokens to the bank contract.
	// It's higher than spender balance but within approved allowance, should fail.
	tx, err = bankContract.Deposit(r.ZEVMAuth, r.ERC20ZRC20Addr, higherBalanceAmount)
	require.NoError(r, err, "Call bank.Deposit() with amout higher than balance")
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequiredTxFailed(r, receipt, "Depositting an amount higher than balance should fail")

	// Deposit 500 ERC20ZRC20 tokens to the bank contract. Should pass.
	tx, err = bankContract.Deposit(r.ZEVMAuth, r.ERC20ZRC20Addr, depositAmount)
	require.NoError(r, err, "Call bank.Deposit() with correct amount")
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequireTxSuccessful(r, receipt, "Depositting a correct amount should pass")

	// Check the deposit event.
	eventDeposit, err := bankContract.ParseDeposit(*receipt.Logs[0])
	require.NoError(r, err, "Parse Deposit event")
	require.Equal(r, r.EVMAddress(), eventDeposit.Zrc20Depositor, "Deposit event token should be r.EVMAddress()")
	require.Equal(r, r.ERC20ZRC20Addr, eventDeposit.Zrc20Token, "Deposit event token should be ERC20ZRC20Addr")
	require.Equal(r, depositAmount, eventDeposit.Amount, "Deposit event amount should be 500")

	// Spender: cosmos coin balance should be 500 at this point.
	cosmosBalance, err = bankContract.BalanceOf(&bind.CallOpts{Context: r.Ctx}, r.ERC20ZRC20Addr, spender)
	require.NoError(r, err, "Call bank.BalanceOf()")
	require.Equal(r, uint64(500), cosmosBalance.Uint64(), "spender cosmos coin balance should be 500")

	// Bank: ERC20ZRC20 balance should be 500 tokens locked.
	bankZRC20Balance, err := r.ERC20ZRC20.BalanceOf(&bind.CallOpts{Context: r.Ctx}, bank.ContractAddress)
	require.NoError(r, err, "Call ERC20ZRC20.BalanceOf")
	require.Equal(r, uint64(500), bankZRC20Balance.Uint64(), "bank ERC20ZRC20 balance should be 500")

	// Try to withdraw 501 ERC20ZRC20 tokens. Should fail.
	tx, err = bankContract.Withdraw(r.ZEVMAuth, r.ERC20ZRC20Addr, big.NewInt(501))
	require.NoError(r, err, "Error calling bank.withdraw()")
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequiredTxFailed(r, receipt, "Withdrawing more than cosmos coin balance amount should fail")

	// Bank: ERC20ZRC20 balance should be 500 tokens locked after a failed withdraw.
	// No tokens should be unlocked with a failed withdraw.
	bankZRC20Balance, err = r.ERC20ZRC20.BalanceOf(&bind.CallOpts{Context: r.Ctx}, bank.ContractAddress)
	require.NoError(r, err, "Call ERC20ZRC20.BalanceOf")
	require.Equal(r, uint64(500), bankZRC20Balance.Uint64(), "bank ERC20ZRC20 balance should be 500")

	// Try to withdraw 500 ERC20ZRC20 tokens. Should pass.
	tx, err = bankContract.Withdraw(r.ZEVMAuth, r.ERC20ZRC20Addr, depositAmount)
	require.NoError(r, err, "Error calling bank.withdraw()")
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequireTxSuccessful(r, receipt, "Withdraw correct amount should pass")

	// Check the withdraw event.
	eventWithdraw, err := bankContract.ParseWithdraw(*receipt.Logs[0])
	require.NoError(r, err, "Parse Withdraw event")
	require.Equal(r, r.EVMAddress(), eventWithdraw.Zrc20Withdrawer, "Withdrawer should be r.EVMAddress()")
	require.Equal(r, r.ERC20ZRC20Addr, eventWithdraw.Zrc20Token, "Withdraw event token should be ERC20ZRC20Addr")
	require.Equal(r, depositAmount, eventWithdraw.Amount, "Withdraw event amount should be 500")

	// Spender: cosmos coin balance should be 0 at this point.
	cosmosBalance, err = bankContract.BalanceOf(&bind.CallOpts{Context: r.Ctx}, r.ERC20ZRC20Addr, spender)
	require.NoError(r, err, "Call bank.BalanceOf()")
	require.Equal(r, uint64(0), cosmosBalance.Uint64(), "spender cosmos coin balance should be 0")

	// Spender: ERC20ZRC20 balance should be 1000 at this point.
	zrc20Balance, err := r.ERC20ZRC20.BalanceOf(&bind.CallOpts{Context: r.Ctx}, spender)
	require.NoError(r, err, "Call bank.BalanceOf()")
	require.Equal(r, uint64(1000), zrc20Balance.Uint64(), "spender ERC20ZRC20 balance should be 1000")

	// Bank: ERC20ZRC20 balance should be 0 tokens locked.
	bankZRC20Balance, err = r.ERC20ZRC20.BalanceOf(&bind.CallOpts{Context: r.Ctx}, bank.ContractAddress)
	require.NoError(r, err, "Call ERC20ZRC20.BalanceOf")
	require.Equal(r, uint64(0), bankZRC20Balance.Uint64(), "bank ERC20ZRC20 balance should be 0")
}

func TestPrecompilesBankNonZRC20(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 0, "No arguments expected")

	// Increase the gasLimit. It's required because of the gas consumed by precompiled functions.
	previousGasLimit := r.ZEVMAuth.GasLimit
	r.ZEVMAuth.GasLimit = 10_000_000
	defer func() {
		r.ZEVMAuth.GasLimit = previousGasLimit
	}()

	spender, bankAddr := r.EVMAddress(), bank.ContractAddress

	// Create a bank contract caller.
	bankContract, err := bank.NewIBank(bank.ContractAddress, r.ZEVMClient)
	require.NoError(r, err, "Failed to create bank contract caller")

	// Deposit and approve 50 WZETA for the test.
	approveAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(50))
	r.DepositAndApproveWZeta(approveAmount)

	// Non ZRC20 balanceOf check should fail.
	_, err = bankContract.BalanceOf(&bind.CallOpts{Context: r.Ctx}, r.WZetaAddr, spender)
	require.Error(r, err, "bank.balanceOf() should error out when checking for non ZRC20 balance")
	require.Contains(
		r,
		err.Error(),
		"invalid token 0x5F0b1a82749cb4E2278EC87F8BF6B618dC71a8bf: token is not a whitelisted ZRC20",
		"Error should be 'token is not a whitelisted ZRC20'",
	)

	// Allow the bank contract to spend 25 WZeta tokens.
	tx, err := r.WZeta.Approve(r.ZEVMAuth, bankAddr, big.NewInt(25))
	require.NoError(r, err, "Error approving allowance for bank contract")
	receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	require.EqualValues(r, uint64(1), receipt.Status, "approve allowance tx failed")

	// Check the allowance of the bank in WZeta tokens. Should be 25.
	allowance, err := r.WZeta.Allowance(&bind.CallOpts{Context: r.Ctx}, spender, bankAddr)
	require.NoError(r, err, "Error retrieving bank allowance")
	require.EqualValues(r, uint64(25), allowance.Uint64(), "Error allowance for bank contract")

	// Call Deposit with 25 Non ZRC20 tokens. Should fail.
	tx, err = bankContract.Deposit(r.ZEVMAuth, r.WZetaAddr, big.NewInt(25))
	require.NoError(r, err, "Error calling bank.deposit()")
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	require.Equal(r, uint64(0), receipt.Status, "Non ZRC20 deposit should fail")

	// Call Withdraw with 25 on ZRC20 tokens. Should fail.
	tx, err = bankContract.Withdraw(r.ZEVMAuth, r.WZetaAddr, big.NewInt(25))
	require.NoError(r, err, "Error calling bank.withdraw()")
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	require.Equal(r, uint64(0), receipt.Status, "Non ZRC20 withdraw should fail")
}
