package runner

import (
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	sdkmath "cosmossdk.io/math"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/zeta-chain/zetacore/cmd/zetacored/config"
	"github.com/zeta-chain/zetacore/e2e/txserver"
	crosschaintypes "github.com/zeta-chain/zetacore/x/crosschain/types"
)

// TestReport is a struct that contains the report for a specific e2e test
// It can be generated with the RunE2ETestsIntoReport method
type TestReport struct {
	Name     string
	Success  bool
	Time     time.Duration
	GasSpent AccountBalancesDiff
}

// TestReports is a slice of TestReport
type TestReports []TestReport

// String returns the string representation of the test report as a table
// it uses text/tabwriter to format the table
func (tr TestReports) String(prefix string) (string, error) {
	var b strings.Builder
	writer := tabwriter.NewWriter(&b, 0, 4, 4, ' ', 0)
	if _, err := fmt.Fprintln(writer, "Name\tSuccess\tTime\tSpent"); err != nil {
		return "", err
	}

	for _, report := range tr {
		spent := formatBalances(report.GasSpent)
		success := "✅"
		if !report.Success {
			success = "❌"
		}
		if _, err := fmt.Fprintf(writer, "%s%s\t%s\t%s\t%s\n", prefix, report.Name, success, report.Time, spent); err != nil {
			return "", err
		}
	}

	if err := writer.Flush(); err != nil {
		return "", err
	}
	return b.String(), nil
}

// PrintTestReports prints the test reports
func (runner *E2ERunner) PrintTestReports(tr TestReports) {
	runner.Logger.Print(" ---📈 E2E Test Report ---")
	table, err := tr.String("")
	if err != nil {
		runner.Logger.Print("Error rendering test report: %s", err)
	}
	runner.Logger.PrintNoPrefix(table)
}

// NetworkReport is a struct that contains the report for the network used after running e2e tests
// This report has been initialized to check the emissions pool balance and if the pool is decreasing
// TODO: add more complete data and validation to the network
// https://github.com/zeta-chain/node/issues/1873
type NetworkReport struct {
	EmissionsPoolBalance sdkmath.Int
	Height               uint64
	CctxCount            int
}

// Validate validates the network report
// This method is used to validate the network after running e2e tests
// It checks the emissions pool balance and if the pool is decreasing
func (nr NetworkReport) Validate() error {
	if nr.EmissionsPoolBalance.GTE(sdkmath.NewIntFromBigInt(EmissionsPoolFunding)) {
		return fmt.Errorf(
			"emissions pool balance is not decreasing, expected less than %s, got %s",
			EmissionsPoolFunding,
			nr.EmissionsPoolBalance,
		)
	}
	return nil
}

// GenerateNetworkReport generates a report for the network used after running e2e tests
func (runner *E2ERunner) GenerateNetworkReport() (NetworkReport, error) {
	// get the emissions pool balance
	balanceRes, err := runner.BankClient.Balance(runner.Ctx, &banktypes.QueryBalanceRequest{
		Address: txserver.EmissionsPoolAddress,
		Denom:   config.BaseDenom,
	})
	if err != nil {
		return NetworkReport{}, err
	}
	emissionsPoolBalance := balanceRes.Balance

	// fetch the height and number of cctxs, this gives a better idea on the activity of the network

	// get the block height
	blockRes, err := runner.ZevmClient.BlockNumber(runner.Ctx)
	if err != nil {
		return NetworkReport{}, err
	}

	// get the number of cctxs
	cctxsRes, err := runner.CctxClient.CctxAll(runner.Ctx, &crosschaintypes.QueryAllCctxRequest{})
	if err != nil {
		return NetworkReport{}, err
	}
	cctxCount := len(cctxsRes.CrossChainTx)

	return NetworkReport{
		EmissionsPoolBalance: emissionsPoolBalance.Amount,
		Height:               blockRes,
		CctxCount:            cctxCount,
	}, nil
}

// PrintNetworkReport prints the network report
func (runner *E2ERunner) PrintNetworkReport(nr NetworkReport) {
	runner.Logger.Print(" ---📈 Network Report ---")
	runner.Logger.Print("Block Height:           %d", nr.Height)
	runner.Logger.Print("CCTX Processed:         %d", nr.CctxCount)
	runner.Logger.Print("Emissions Pool Balance: %sZETA", nr.EmissionsPoolBalance.Quo(sdkmath.NewIntFromUint64(1e18)))

}
