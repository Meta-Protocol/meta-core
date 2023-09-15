//go:build !PRIVNET && !TESTNET
// +build !PRIVNET,!TESTNET

package common

func EthChain() Chain {
	return Chain{
		ChainName: ChainName_eth_mainnet,
		ChainId:   1,
	}
}

func BscMainnetChain() Chain {
	return Chain{
		ChainName: ChainName_bsc_mainnet,
		ChainId:   56,
	}
}

func ZetaChain() Chain {
	return Chain{
		ChainName: ChainName_zeta_mainnet,
		ChainId:   70000,
	}
}

func BtcMainnetChain() Chain {
	return Chain{
		ChainName: ChainName_btc_mainnet,
		ChainId:   8332,
	}
}

func PolygonChain() Chain {
	return Chain{
		ChainName: ChainName_polygon_mainnet,
		ChainId:   137,
	}
}

func DefaultChainsList() []*Chain {
	chains := []Chain{
		BtcMainnetChain(),
		PolygonChain(),
		BscMainnetChain(),
		EthChain(),
		ZetaChain(),
	}
	var c []*Chain
	for i := 0; i < len(chains); i++ {
		c = append(c, &chains[i])
	}
	return c
}
