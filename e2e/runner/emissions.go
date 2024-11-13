package runner

import (
	"fmt"
	"strings"

	"cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
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
		r.Logger.Print("🏃 withdrawing emissions from the emissions pool on ZetaChain for observer %s", observer)
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

		amount, found := strings.CutSuffix(availableAmount.Amount, baseDenom)
		if !found {
			return fmt.Errorf("invalid amount %s", availableAmount.Amount)
		}

		amountInt, ok := sdkmath.NewIntFromString(amount)
		if !ok {
			return fmt.Errorf("failed to convert string to int")
		}

		if amountInt.IsZero() {
			r.Logger.Print("no emissions to withdraw for observer %s", observer)
			continue
		}

		err = r.ZetaTxServer.WithdrawAllEmissions(amountInt, e2eutils.UserEmissionsWithdrawName, observer)
		if err != nil {
			return err
		}

		balanceAfter, err := r.BankClient.Balance(r.Ctx, queryObserverBalance)
		if err != nil {
			return errors.Wrapf(err, "failed to get balance for observer after withdrawing emissions %s", observer)
		}

		changeInBalance := balanceAfter.Balance.Sub(*balanceBefore.Balance).Amount.String()
		if changeInBalance != amount {
			return fmt.Errorf(
				"invalid balance change for observer %s, expected %s, got %s",
				observer,
				amount,
				changeInBalance,
			)
		}
	}

	return nil
}
