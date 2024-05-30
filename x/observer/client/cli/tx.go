package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/zeta-chain/zetacore/x/observer/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdAddObserver(),
		CmdUpdateChainParams(),
		CmdRemoveChainParams(),
		CmdUpdateKeygen(),
		CmdVoteBlame(),
		CmdUpdateObserver(),
		CmdEncode(),
		CmdResetChainNonces(),
		CmdVoteTSS(),
		CmdEnableCCTX(),
		CmdDisableCCTX(),
		CmdUpdateGasPriceIncreaseFlags(),
	)

	return cmd
}
