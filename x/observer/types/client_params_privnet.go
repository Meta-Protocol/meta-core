//go:build PRIVNET
// +build PRIVNET

package types

import (
	"github.com/zeta-chain/zetacore/common"
)

func GetClientParams() map[int64]ClientParams {
	return map[int64]ClientParams{
		common.GoerliLocalNetChain().ChainId: {
			ConfirmationCount:           14,
			GasPriceTicker:              5,
			ZetaTokenContractAddress:    "0xA8D5060feb6B456e886F023709A2795373691E63",
			ConnectorContractAddress:    "0x733aB8b06DDDEf27Eaa72294B0d7c9cEF7f12db9",
			Erc20CustodyContractAddress: "0xD28D6A0b8189305551a0A8bd247a6ECa9CE781Ca",
			BlockTimeExternal:           2,
			BlockTimeZeta:               6,
		},

		common.ZetaChain().ChainId: {
			ConfirmationCount:           0,
			GasPriceTicker:              5,
			ZetaTokenContractAddress:    "0x2DD9830f8Ac0E421aFF9B7c8f7E9DF6F65DBF6Ea",
			ConnectorContractAddress:    "",
			Erc20CustodyContractAddress: "",
			BlockTimeExternal:           6,
			BlockTimeZeta:               6,
		},
	}

}
