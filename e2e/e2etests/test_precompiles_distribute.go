package e2etests

import (
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	"github.com/zeta-chain/node/precompiles/bank"
	"github.com/zeta-chain/node/precompiles/staking"
	ptypes "github.com/zeta-chain/node/precompiles/types"
)

func TestPrecompilesDistribute(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 0, "No arguments expected")

	var (
		spenderAddress            = r.EVMAddress()
		distributeContractAddress = staking.ContractAddress
		feeCollectorAddress       = authtypes.NewModuleAddress("fee_collector")
		lockerAddress             = bank.ContractAddress

		zrc20Address = r.ERC20ZRC20Addr
		zrc20Denom   = ptypes.ZRC20ToCosmosDenom(zrc20Address)

		oneThousand    = big.NewInt(1e3)
		oneThousandOne = big.NewInt(1001)
		fiveHundred    = big.NewInt(500)
		fiveHundredOne = big.NewInt(501)

		previousGasLimit = r.ZEVMAuth.GasLimit
	)

	// Set new gas limit to avoid out of gas errors.
	r.ZEVMAuth.GasLimit = 10_000_000

	// Set the test to reset the state after it finishes.
	defer resetTest(r, lockerAddress, previousGasLimit)

	// Get ERC20ZRC20.
	txHash := r.DepositERC20WithAmountAndMessage(spenderAddress, oneThousand, []byte{})
	utils.WaitCctxMinedByInboundHash(r.Ctx, txHash.Hex(), r.CctxClient, r.Logger, r.CctxTimeout)

	dstrContract, err := staking.NewIStaking(distributeContractAddress, r.ZEVMClient)
	require.NoError(r, err, "failed to create distribute contract caller")

	validators, err := dstrContract.GetAllValidators(&bind.CallOpts{})
	require.NoError(r, err)
	fmt.Println(validators)

	// Check initial balances.
	balanceShouldBe(r, 1000, checkZRC20Balance(r, spenderAddress))
	balanceShouldBe(r, 0, checkZRC20Balance(r, lockerAddress))
	balanceShouldBe(r, 0, checkCosmosBalance(r, feeCollectorAddress, zrc20Denom))

	tx, err := dstrContract.Distribute(r.ZEVMAuth, zrc20Address, oneThousand)
	require.NoError(r, err)
	receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequiredTxFailed(r, receipt, "distribute should fail when there's no allowance")

	// Balances shouldn't change after a failed attempt.
	balanceShouldBe(r, 1000, checkZRC20Balance(r, spenderAddress))
	balanceShouldBe(r, 0, checkZRC20Balance(r, lockerAddress))
	balanceShouldBe(r, 0, checkCosmosBalance(r, feeCollectorAddress, zrc20Denom))

	// Allow 500.
	approveAllowance(r, distributeContractAddress, fiveHundred)

	// Shouldn't be able to distribute more than allowed.
	tx, err = dstrContract.Distribute(r.ZEVMAuth, zrc20Address, fiveHundredOne)
	require.NoError(r, err)
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequiredTxFailed(r, receipt, "distribute should fail trying to distribute more than allowed")

	// Balances shouldn't change after a failed attempt.
	balanceShouldBe(r, 1000, checkZRC20Balance(r, spenderAddress))
	balanceShouldBe(r, 0, checkZRC20Balance(r, lockerAddress))
	balanceShouldBe(r, 0, checkCosmosBalance(r, feeCollectorAddress, zrc20Denom))

	// Raise the allowance to 1000.
	approveAllowance(r, distributeContractAddress, oneThousand)

	// Shouldn't be able to distribute more than owned balance.
	tx, err = dstrContract.Distribute(r.ZEVMAuth, zrc20Address, oneThousandOne)
	require.NoError(r, err)
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequiredTxFailed(r, receipt, "distribute should fail trying to distribute more than owned balance")

	// Balances shouldn't change after a failed attempt.
	balanceShouldBe(r, 1000, checkZRC20Balance(r, spenderAddress))
	balanceShouldBe(r, 0, checkZRC20Balance(r, lockerAddress))
	balanceShouldBe(r, 0, checkCosmosBalance(r, feeCollectorAddress, zrc20Denom))

	// Should be able to distribute 500, which is within balance and allowance.
	tx, err = dstrContract.Distribute(r.ZEVMAuth, zrc20Address, fiveHundred)
	require.NoError(r, err)
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequireTxSuccessful(r, receipt, "distribute should succeed when distributing within balance and allowance")

	balanceShouldBe(r, 500, checkZRC20Balance(r, spenderAddress))
	balanceShouldBe(r, 500, checkZRC20Balance(r, lockerAddress))
	balanceShouldBe(r, 500, checkCosmosBalance(r, feeCollectorAddress, zrc20Denom))

	// After one block the rewards should have been distributed and fee collector should have 0 ZRC20 balance.
	r.WaitForBlocks(1)
	balanceShouldBe(r, 0, checkCosmosBalance(r, feeCollectorAddress, zrc20Denom))

	accAddress := types.AccAddress("zetavaloper16vuh496n7wahw6m8dmzc0p5meymmctleuaz562")
	valBal := checkCosmosBalance(r, accAddress, zrc20Denom)
	fmt.Println(valBal)

	accAddress = types.AccAddress(validators[0].OperatorAddress)
	valBal = checkCosmosBalance(r, accAddress, zrc20Denom)
	fmt.Println(valBal)

	coins := checkAllCosmosBalance(r, accAddress)
	fmt.Printf("Coins: %+v\n", coins)
}

func checkCosmosBalance(r *runner.E2ERunner, address types.AccAddress, denom string) *big.Int {
	bal, err := r.BankClient.Balance(
		r.Ctx,
		&banktypes.QueryBalanceRequest{Address: address.String(), Denom: denom},
	)
	require.NoError(r, err)

	return bal.Balance.Amount.BigInt()
}

func checkAllCosmosBalance(r *runner.E2ERunner, address types.AccAddress) types.Coins {
	bal, err := r.BankClient.AllBalances(
		r.Ctx,
		&banktypes.QueryAllBalancesRequest{Address: address.String()},
	)
	require.NoError(r, err)

	return bal.Balances
}

func resetTest(r *runner.E2ERunner, lockerAddress common.Address, previousGasLimit uint64) {
	r.ZEVMAuth.GasLimit = previousGasLimit

	// Reset the allowance to 0; this is needed when running upgrade tests where this test runs twice.
	tx, err := r.ERC20ZRC20.Approve(r.ZEVMAuth, lockerAddress, big.NewInt(0))
	require.NoError(r, err)
	receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequireTxSuccessful(r, receipt, "Resetting allowance failed")
}
