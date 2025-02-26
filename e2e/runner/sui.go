package runner

import (
	"fmt"
	"strings"

	"cosmossdk.io/math"
	"github.com/block-vision/sui-go-sdk/models"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/pkg/contracts/sui"
)

// SuiDepositSUI calls Deposit on Sui
func (r *E2ERunner) SuiDepositSUI(
	receiver ethcommon.Address,
	amount math.Uint,
) models.SuiTransactionBlockResponse {
	signer, err := r.Account.SuiSigner()
	require.NoError(r, err, "get deployer signer")

	// retrieve SUI coin object to deposit
	coinObjectID := r.suiSplitSUI(signer, amount)

	// create the tx
	return r.suiExecuteDeposit(signer, string(sui.SUI), coinObjectID, receiver)
}

// SuiDepositAndCallSUI calls DepositAndCall on Sui
func (r *E2ERunner) SuiDepositAndCallSUI(
	receiver ethcommon.Address,
	amount math.Uint,
	payload []byte,
) models.SuiTransactionBlockResponse {
	signer, err := r.Account.SuiSigner()
	require.NoError(r, err, "get deployer signer")

	// retrieve SUI coin object to deposit
	coinObjectID := r.suiSplitSUI(signer, amount)

	// create the tx
	return r.suiExecuteDepositAndCall(signer, string(sui.SUI), coinObjectID, receiver, payload)
}

// SuiDepositFungibleToken calls Deposit with fungible token on Sui
func (r *E2ERunner) SuiDepositFungibleToken(
	receiver ethcommon.Address,
	amount math.Uint,
) models.SuiTransactionBlockResponse {
	signer, err := r.Account.SuiSigner()
	require.NoError(r, err, "get deployer signer")

	// retrieve SUI coin object to deposit
	coinObjectID := r.suiSplitUSDC(signer, amount)

	// create the tx
	return r.suiExecuteDeposit(signer, "0x"+r.SuiTokenCoinType, coinObjectID, receiver)
}

// SuiFungibleTokenDepositAndCall calls DepositAndCall with fungible token on Sui
func (r *E2ERunner) SuiFungibleTokenDepositAndCall(
	receiver ethcommon.Address,
	amount math.Uint,
	payload []byte,
) models.SuiTransactionBlockResponse {
	signer, err := r.Account.SuiSigner()
	require.NoError(r, err, "get deployer signer")

	// retrieve SUI coin object to deposit
	coinObjectID := r.suiSplitUSDC(signer, amount)

	// create the tx
	return r.suiExecuteDepositAndCall(signer, "0x"+r.SuiTokenCoinType, coinObjectID, receiver, payload)
}

// SuiMintUSDC mints FakeUSDC on Sui to a receiver
// this function requires the signer to be the owner of the trasuryCap
func (r *E2ERunner) SuiMintUSDC(
	amount,
	receiver string,
) models.SuiTransactionBlockResponse {
	signer, err := r.Account.SuiSigner()
	require.NoError(r, err, "get deployer signer")

	// extract the package ID from the coin type
	splitted := strings.Split(r.SuiTokenCoinType, "::")
	require.Len(r, splitted, 3, "coinType should be in format <packageID>::<module>::<name>")
	packageID := "0x" + splitted[0]

	// create the tx
	tx, err := r.Clients.Sui.MoveCall(r.Ctx, models.MoveCallRequest{
		Signer:          signer.Address(),
		PackageObjectId: packageID,
		Module:          "fake_usdc",
		Function:        "mint",
		TypeArguments:   []any{},
		Arguments:       []any{r.SuiTokenTreasuryCap, amount, receiver},
		GasBudget:       "5000000000",
	})
	require.NoError(r, err)

	return r.suiExecuteTx(signer, tx)
}

// suiExecuteDeposit executes a deposit on the SUI contract
func (r *E2ERunner) suiExecuteDeposit(
	signer *sui.SignerSecp256k1,
	coinType string,
	coinObjectID string,
	receiver ethcommon.Address,
) models.SuiTransactionBlockResponse {
	// create the tx
	tx, err := r.Clients.Sui.MoveCall(r.Ctx, models.MoveCallRequest{
		Signer:          signer.Address(),
		PackageObjectId: r.GatewayPackageID,
		Module:          "gateway",
		Function:        "deposit",
		TypeArguments:   []any{coinType},
		Arguments:       []any{r.GatewayObjectID, coinObjectID, receiver.Hex()},
		GasBudget:       "5000000000",
	})
	require.NoError(r, err)

	return r.suiExecuteTx(signer, tx)
}

// suiExecuteDepositAndCall executes a depositAndCall on the SUI contract
func (r *E2ERunner) suiExecuteDepositAndCall(
	signer *sui.SignerSecp256k1,
	coinType string,
	coinObjectID string,
	receiver ethcommon.Address,
	payload []byte,
) models.SuiTransactionBlockResponse {
	// create the tx
	tx, err := r.Clients.Sui.MoveCall(r.Ctx, models.MoveCallRequest{
		Signer:          signer.Address(),
		PackageObjectId: r.GatewayPackageID,
		Module:          "gateway",
		Function:        "deposit_and_call",
		TypeArguments:   []any{coinType},
		Arguments:       []any{r.GatewayObjectID, coinObjectID, receiver.Hex(), payload},
		GasBudget:       "5000000000",
	})
	require.NoError(r, err)

	return r.suiExecuteTx(signer, tx)
}

// suiSplitUSDC splits USDC coin and obtain a USDC coin object with the wanted balance
func (r *E2ERunner) suiSplitUSDC(signer *sui.SignerSecp256k1, balance math.Uint) (objID string) {
	// find the coin to split
	originalCoin := r.suiFindCoinWithBalanceAbove(signer.Address(), balance, "0x"+r.SuiTokenCoinType)

	tx, err := r.Clients.Sui.SplitCoin(r.Ctx, models.SplitCoinRequest{
		Signer:       signer.Address(),
		CoinObjectId: originalCoin,
		SplitAmounts: []string{balance.String()},
		GasBudget:    "5000000000",
	})

	require.NoError(r, err)
	r.suiExecuteTx(signer, tx)

	// find the split coin
	return r.suiFindCoinWithBalance(signer.Address(), balance, "0x"+r.SuiTokenCoinType)
}

// suiSplitSUI splits SUI coin and obtain a SUI coin object with the wanted balance
func (r *E2ERunner) suiSplitSUI(signer *sui.SignerSecp256k1, balance math.Uint) (objID string) {
	// find the coin to split
	originalCoin := r.suiFindCoinWithBalanceAbove(signer.Address(), balance, string(sui.SUI))

	// split the coin using the PaySui API
	tx, err := r.Clients.Sui.PaySui(r.Ctx, models.PaySuiRequest{
		Signer:      signer.Address(),
		SuiObjectId: []string{originalCoin},
		Recipient:   []string{signer.Address()},
		Amount:      []string{balance.String()},
		GasBudget:   "5000000000",
	})
	require.NoError(r, err)

	r.suiExecuteTx(signer, tx)

	// find the split coin
	return r.suiFindCoinWithBalance(signer.Address(), balance, string(sui.SUI))
}

func (r *E2ERunner) suiFindCoinWithBalance(
	address string,
	balance math.Uint,
	coinType string,
) (coinID string) {
	return r.suiFindCoin(address, balance, coinType, func(a, b math.Uint) bool {
		return a.Equal(b)
	})
}

func (r *E2ERunner) suiFindCoinWithBalanceAbove(
	address string,
	balanceAbove math.Uint,
	coinType string,
) (coinID string) {
	return r.suiFindCoin(address, balanceAbove, coinType, func(a, b math.Uint) bool {
		return a.GTE(b)
	})
}

type compFunc func(a, b math.Uint) bool

func (r *E2ERunner) suiFindCoin(
	address string,
	balance math.Uint,
	coinType string,
	comp compFunc,
) (coinID string) {
	res, err := r.Clients.Sui.SuiXGetCoins(r.Ctx, models.SuiXGetCoinsRequest{
		Owner:    address,
		CoinType: coinType,
	})
	require.NoError(r, err)

	for _, data := range res.Data {
		coinBalance, err := math.ParseUint(data.Balance)
		require.NoError(r, err)

		if comp(coinBalance, balance) {
			return data.CoinObjectId
		}
	}

	require.FailNow(r, fmt.Sprintf("coin %s not found for address %s", coinType, address))
	return ""
}

func (r *E2ERunner) suiExecuteTx(
	signer *sui.SignerSecp256k1,
	tx models.TxnMetaData,
) models.SuiTransactionBlockResponse {
	// sign the tx
	signature, err := signer.SignTxBlock(tx)
	require.NoError(r, err, "sign transaction")

	resp, err := r.Clients.Sui.SuiExecuteTransactionBlock(r.Ctx, models.SuiExecuteTransactionBlockRequest{
		TxBytes:     tx.TxBytes,
		Signature:   []string{signature},
		RequestType: "WaitForLocalExecution",
	})
	require.NoError(r, err)

	return resp
}
