package chains

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/require"
)

func TestBitcoinNetParamsFromChainID(t *testing.T) {
	tests := []struct {
		name     string
		chainID  int64
		expected *chaincfg.Params
		wantErr  bool
	}{
		{"Regnet", BtcRegtestChain.ChainId, BitcoinRegnetParams, false},
		{"Mainnet", BtcMainnetChain.ChainId, BitcoinMainnetParams, false},
		{"Testnet", BtcTestNetChain.ChainId, BitcoinTestnetParams, false},
		{"Unknown", -1, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, err := BitcoinNetParamsFromChainID(tt.chainID)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, params)
			}
		})
	}
}

func TestBitcoinChainIDFromNetParams(t *testing.T) {
	tests := []struct {
		name            string
		networkName     string
		expectedChainID int64
		wantErr         bool
	}{
		{"Regnet", BitcoinRegnetParams.Name, BtcRegtestChain.ChainId, false},
		{"Mainnet", BitcoinMainnetParams.Name, BtcMainnetChain.ChainId, false},
		{"Testnet", BitcoinTestnetParams.Name, BtcTestNetChain.ChainId, false},
		{"Unknown", "Unknown", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chainID, err := BitcoinChainIDFromNetworkName(tt.networkName)
			if tt.wantErr {
				require.Error(t, err)
				require.Zero(t, chainID)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedChainID, chainID)
			}
		})
	}
}

func TestIsBitcoinRegnet(t *testing.T) {
	require.True(t, IsBitcoinRegnet(BtcRegtestChain.ChainId))
	require.False(t, IsBitcoinRegnet(BtcMainnetChain.ChainId))
	require.False(t, IsBitcoinRegnet(BtcTestNetChain.ChainId))
}

func TestIsBitcoinMainnet(t *testing.T) {
	require.True(t, IsBitcoinMainnet(BtcMainnetChain.ChainId))
	require.False(t, IsBitcoinMainnet(BtcRegtestChain.ChainId))
	require.False(t, IsBitcoinMainnet(BtcTestNetChain.ChainId))
}
