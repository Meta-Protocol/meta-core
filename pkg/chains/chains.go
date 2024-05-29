package chains

import "fmt"

var (
	/**
	* Mainnet chains
	 */

	// ZetaChainMainnet is the mainnet chain for Zeta
	ZetaChainMainnet = Chain{
		ChainName:   ChainName_zeta_mainnet,
		ChainId:     7000,
		Network:     Network_zeta,
		NetworkType: NetworkType_mainnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_tendermint,
		IsExternal:  false,
		CctxGateway: CCTXGateway_zevm,
	}

	// Ethereum is Ethereum mainnet
	Ethereum = Chain{
		ChainName:   ChainName_eth_mainnet,
		ChainId:     1,
		Network:     Network_eth,
		NetworkType: NetworkType_mainnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_ethereum,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	// BscMainnet is Binance Smart Chain mainnet
	BscMainnet = Chain{
		ChainName:   ChainName_bsc_mainnet,
		ChainId:     56,
		Network:     Network_bsc,
		NetworkType: NetworkType_mainnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_ethereum,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	// BitcoinMainnet is Bitcoin mainnet
	BitcoinMainnet = Chain{
		ChainName:   ChainName_btc_mainnet,
		ChainId:     8332,
		Network:     Network_btc,
		NetworkType: NetworkType_mainnet,
		Vm:          Vm_no_vm,
		Consensus:   Consensus_bitcoin,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	// Polygon is Polygon mainnet
	Polygon = Chain{
		ChainName:   ChainName_polygon_mainnet,
		ChainId:     137,
		Network:     Network_polygon,
		NetworkType: NetworkType_mainnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_ethereum,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	// OptimismMainnet is Optimism mainnet
	OptimismMainnet = Chain{
		ChainName:   ChainName_optimism_mainnet,
		ChainId:     10,
		Network:     Network_optimism,
		NetworkType: NetworkType_mainnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_op_stack,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	// BaseMainnet is Base mainnet
	BaseMainnet = Chain{
		ChainName:   ChainName_base_mainnet,
		ChainId:     8453,
		Network:     Network_base,
		NetworkType: NetworkType_mainnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_op_stack,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	/**
	* Testnet chains
	 */

	// ZetaChainTestnet is the testnet chain for Zeta
	ZetaChainTestnet = Chain{
		ChainName:   ChainName_zeta_testnet,
		ChainId:     7001,
		Network:     Network_zeta,
		NetworkType: NetworkType_testnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_tendermint,
		IsExternal:  false,
		CctxGateway: CCTXGateway_zevm,
	}

	// Sepolia is Ethereum sepolia testnet
	Sepolia = Chain{
		ChainName:   ChainName_sepolia_testnet,
		ChainId:     11155111,
		Network:     Network_eth,
		NetworkType: NetworkType_testnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_ethereum,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	// BscTestnet is Binance Smart Chain testnet
	BscTestnet = Chain{
		ChainName:   ChainName_bsc_testnet,
		ChainId:     97,
		Network:     Network_bsc,
		NetworkType: NetworkType_testnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_ethereum,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	// BitcoinTestnet is Bitcoin testnet3
	BitcoinTestnet = Chain{
		ChainName:   ChainName_btc_testnet,
		ChainId:     18332,
		Network:     Network_btc,
		NetworkType: NetworkType_testnet,
		Vm:          Vm_no_vm,
		Consensus:   Consensus_bitcoin,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	// Amoy is Polygon amoy testnet
	Amoy = Chain{
		ChainName:   ChainName_amoy_testnet,
		ChainId:     80002,
		Network:     Network_polygon,
		NetworkType: NetworkType_testnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_ethereum,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	// OptimismSepolia is Optimism sepolia testnet
	OptimismSepolia = Chain{
		ChainName:   ChainName_optimism_sepolia,
		ChainId:     11155420,
		Network:     Network_optimism,
		NetworkType: NetworkType_testnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_op_stack,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	// BaseSepolia is Base sepolia testnet
	BaseSepolia = Chain{
		ChainName:   ChainName_base_sepolia,
		ChainId:     84532,
		Network:     Network_base,
		NetworkType: NetworkType_testnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_op_stack,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	/**
	* Devnet chains
	 */

	// ZetaChainDevnet is the devnet chain for Zeta
	// used as live testing environment
	ZetaChainDevnet = Chain{
		ChainName:   ChainName_zeta_mainnet,
		ChainId:     70000,
		Network:     Network_zeta,
		NetworkType: NetworkType_devnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_tendermint,
		IsExternal:  false,
		CctxGateway: CCTXGateway_zevm,
	}

	/**
	* Privnet chains
	 */

	// ZetaChainPrivnet is the privnet chain for Zeta (localnet)
	ZetaChainPrivnet = Chain{
		ChainName:   ChainName_zeta_mainnet,
		ChainId:     101,
		Network:     Network_zeta,
		NetworkType: NetworkType_privnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_tendermint,
		IsExternal:  false,
		CctxGateway: CCTXGateway_zevm,
	}

	// BitcoinRegtest is Bitcoin regtest (localnet)
	BitcoinRegtest = Chain{
		ChainName:   ChainName_btc_regtest,
		ChainId:     18444,
		Network:     Network_btc,
		NetworkType: NetworkType_privnet,
		Vm:          Vm_no_vm,
		Consensus:   Consensus_bitcoin,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	// GoerliLocalnet is Ethereum local goerli (localnet)
	GoerliLocalnet = Chain{
		ChainName:   ChainName_goerli_localnet,
		ChainId:     1337,
		Network:     Network_eth,
		NetworkType: NetworkType_privnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_ethereum,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	/**
	* Deprecated chains
	 */

	// Goerli is Ethereum goerli testnet (deprecated for sepolia)
	Goerli = Chain{
		ChainName:   ChainName_goerli_testnet,
		ChainId:     5,
		Network:     Network_eth,
		NetworkType: NetworkType_testnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_ethereum,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}

	// Mumbai is Polygon mumbai testnet (deprecated for amoy)
	Mumbai = Chain{
		ChainName:   ChainName_mumbai_testnet,
		ChainId:     80001,
		Network:     Network_polygon,
		NetworkType: NetworkType_testnet,
		Vm:          Vm_evm,
		Consensus:   Consensus_ethereum,
		IsExternal:  true,
		CctxGateway: CCTXGateway_observers,
	}
)

func BtcDustOffset() int64 {
	return 2000
}

// DefaultChainsList returns a list of default chains
func DefaultChainsList() []*Chain {
	return chainListPointers([]Chain{
		BitcoinMainnet,
		BscMainnet,
		Ethereum,
		BitcoinTestnet,
		Mumbai,
		Amoy,
		BscTestnet,
		Goerli,
		Sepolia,
		BitcoinRegtest,
		GoerliLocalnet,
		ZetaChainMainnet,
		ZetaChainTestnet,
		ZetaChainDevnet,
		ZetaChainPrivnet,
		Polygon,
		OptimismMainnet,
		OptimismSepolia,
		BaseMainnet,
		BaseSepolia,
	})
}

// ChainListByNetworkType returns a list of chains by network type
func ChainListByNetworkType(networkType NetworkType) []*Chain {
	var chainList []*Chain
	for _, chain := range DefaultChainsList() {
		if chain.NetworkType == networkType {
			chainList = append(chainList, chain)
		}
	}
	return chainList
}

// ChainListByNetwork returns a list of chains by network
func ChainListByNetwork(network Network) []*Chain {
	var chainList []*Chain
	for _, chain := range DefaultChainsList() {
		if chain.Network == network {
			chainList = append(chainList, chain)
		}
	}
	return chainList
}

// ExternalChainList returns a list chains that are not Zeta
func ExternalChainList() []*Chain {
	var chainList []*Chain
	for _, chain := range DefaultChainsList() {
		if chain.IsExternal {
			chainList = append(chainList, chain)
		}
	}
	return chainList
}

// ChainListByConsensus returns a list of chains by consensus
func ChainListByConsensus(consensus Consensus) []*Chain {
	var chainList []*Chain
	for _, chain := range DefaultChainsList() {
		if chain.Consensus == consensus {
			chainList = append(chainList, chain)
		}
	}
	return chainList
}

// ChainListForHeaderSupport returns a list of chains that support headers
func ChainListForHeaderSupport() []*Chain {
	var chainList []*Chain
	for _, chain := range DefaultChainsList() {
		if chain.Consensus == Consensus_ethereum || chain.Consensus == Consensus_bitcoin {
			chainList = append(chainList, chain)
		}
	}
	return chainList
}

// ZetaChainFromChainID returns a ZetaChain chain object  from a Cosmos chain ID
func ZetaChainFromChainID(chainID string) (Chain, error) {
	ethChainID, err := CosmosToEthChainID(chainID)
	if err != nil {
		return Chain{}, err
	}

	switch ethChainID {
	case ZetaChainPrivnet.ChainId:
		return ZetaChainPrivnet, nil
	case ZetaChainMainnet.ChainId:
		return ZetaChainMainnet, nil
	case ZetaChainTestnet.ChainId:
		return ZetaChainTestnet, nil
	case ZetaChainDevnet.ChainId:
		return ZetaChainDevnet, nil
	default:
		return Chain{}, fmt.Errorf("chain %d not found", ethChainID)
	}
}

// TODO : https://github.com/zeta-chain/node/issues/2080
// remove the usage of this function
// chainListPointers returns a list of chain pointers
func chainListPointers(chains []Chain) []*Chain {
	var c []*Chain
	for i := 0; i < len(chains); i++ {
		c = append(c, &chains[i])
	}
	return c
}
