package runner

import (
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/zeta-chain/node/cmd/zetacored/config"
	"github.com/zeta-chain/node/e2e/txserver"
	e2eutils "github.com/zeta-chain/node/e2e/utils"
	emissionstypes "github.com/zeta-chain/node/x/emissions/types"
	observertypes "github.com/zeta-chain/node/x/observer/types"
)

// FundEmissionsPool funds the emissions pool on ZetaChain with the same value as used originally on mainnet (20M ZETA)
func (r *E2ERunner) FundEmissionsPool() error {
	r.Logger.Print("⚙️ funding the emissions pool on ZetaChain with 20M ZETA (%s)", txserver.EmissionsPoolAddress)

	return r.ZetaTxServer.FundEmissionsPool(e2eutils.OperationalPolicyName, EmissionsPoolFunding)
}

// WithdrawEmissions withdraws emissions from the emission pool on ZetaChain for all observers
// This functions uses the UserEmissionsWithdrawName to create the withdraw tx.
// UserEmissionsWithdraw can sign the authz transactions because the necessary permissions are granted in the genesis file
func (r *E2ERunner) WithdrawEmissions() error {
	observerSet, err := r.ObserverClient.ObserverSet(r.Ctx, &observertypes.QueryObserverSet{})
	if err != nil {
		return err
	}

	for _, observer := range observerSet.Observers {
		r.Logger.Print("🏃 Withdrawing emissions for observer %s", observer)
		var (
			baseDenom            = config.BaseDenom
			queryObserverBalance = &banktypes.QueryBalanceRequest{
				Address: observer,
				Denom:   baseDenom,
			}
		)

		balanceBefore, err := r.BankClient.Balance(r.Ctx, queryObserverBalance)
		if err != nil {
			return errors.Wrapf(err, "failed to get balance for observer before withdrawing emissions %s", observer)
		}

		availableAmount, err := r.EmissionsClient.ShowAvailableEmissions(
			r.Ctx,
			&emissionstypes.QueryShowAvailableEmissionsRequest{
				Address: observer,
			},
		)
		if err != nil {
			return fmt.Errorf("failed to get available emissions for observer %s: %w", observer, err)
		}

		availableCoin, err := sdk.ParseCoinNormalized(availableAmount.Amount)
		if err != nil {
			return fmt.Errorf("failed to parse coin amount: %w", err)
		}

		if availableCoin.Amount.IsZero() {
			r.Logger.Print("no emissions to withdraw for observer %s", observer)
			continue
		}

		if err := r.ZetaTxServer.WithdrawAllEmissions(availableCoin.Amount, e2eutils.UserEmissionsWithdrawName, observer); err != nil {
			return err
		}

		balanceAfter, err := r.BankClient.Balance(r.Ctx, queryObserverBalance)
		if err != nil {
			return errors.Wrapf(err, "failed to get balance for observer after withdrawing emissions %s", observer)
		}

		changeInBalance := balanceAfter.Balance.Sub(*balanceBefore.Balance).Amount
		if !changeInBalance.Equal(availableCoin.Amount) {
			return fmt.Errorf(
				"invalid balance change for observer %s, expected %s, got %s",
				observer,
				availableCoin.Amount,
				changeInBalance,
			)
		}
	}

	return nil
}
