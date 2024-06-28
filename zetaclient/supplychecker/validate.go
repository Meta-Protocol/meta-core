package supplychecker

import (
	sdkmath "cosmossdk.io/math"
	"github.com/rs/zerolog"
)

// ValidateZetaSupply validates the zeta supply from the checked values
func ValidateZetaSupply(
	logger zerolog.Logger,
	abortedTxAmounts, zetaInTransit, genesisAmounts, externalChainTotalSupply, zetaTokenSupplyOnNode, ethLockedAmount sdkmath.Int,
) bool {
	lhs := ethLockedAmount.Sub(abortedTxAmounts)
	rhs := zetaTokenSupplyOnNode.Add(zetaInTransit).Add(externalChainTotalSupply).Sub(genesisAmounts)

	copyZetaTokenSupplyOnNode := zetaTokenSupplyOnNode
	copyGenesisAmounts := genesisAmounts
	nodeAmounts := copyZetaTokenSupplyOnNode.Sub(copyGenesisAmounts)
	logs := ZetaSupplyCheckLogs{
		Logger:                   logger,
		AbortedTxAmounts:         abortedTxAmounts,
		ZetaInTransit:            zetaInTransit,
		ExternalChainTotalSupply: externalChainTotalSupply,
		NodeAmounts:              nodeAmounts,
		ZetaTokenSupplyOnNode:    zetaTokenSupplyOnNode,
		EthLockedAmount:          ethLockedAmount,
		LHS:                      lhs,
		RHS:                      rhs,
	}
	defer logs.LogOutput()

	logs.SupplyCheckSuccess = lhs.Equal(rhs)
	return logs.SupplyCheckSuccess
}
