package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	evmtypes "github.com/zeta-chain/ethermint/x/evm/types"

	zetachains "github.com/zeta-chain/node/pkg/chains"
	"github.com/zeta-chain/node/pkg/coin"
	"github.com/zeta-chain/node/pkg/crypto"
	"github.com/zeta-chain/node/testutil/sample"
	authoritytypes "github.com/zeta-chain/node/x/authority/types"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
	fungibletypes "github.com/zeta-chain/node/x/fungible/types"
	observertypes "github.com/zeta-chain/node/x/observer/types"
)

// simulation parameter constants
const (
	StakePerAccount           = "stake_per_account"
	InitiallyBondedValidators = "initially_bonded_validators"
)

// updateBankState updates the bank genesis state.
// It adds the following
// - The not bonded balance for the not bonded pool
func updateBankState(
	t *testing.T,
	rawState map[string]json.RawMessage,
	cdc codec.Codec,
	notBondedCoins sdk.Coin,
) *banktypes.GenesisState {
	bankStateBz, ok := rawState[banktypes.ModuleName]
	require.True(t, ok, "bank genesis state is missing")

	bankState := new(banktypes.GenesisState)
	err := cdc.UnmarshalJSON(bankStateBz, bankState)
	require.NoError(t, err)

	stakingAddr := authtypes.NewModuleAddress(stakingtypes.NotBondedPoolName).String()
	var found bool
	for _, balance := range bankState.Balances {
		if balance.Address == stakingAddr {
			found = true
			break
		}
	}
	if !found {
		bankState.Balances = append(bankState.Balances, banktypes.Balance{
			Address: stakingAddr,
			Coins:   sdk.NewCoins(notBondedCoins),
		})
	}

	return bankState
}

// updateEVMState updates the evm genesis state.
// It replaces the EvmDenom with BondDenom
func updateEVMState(
	t *testing.T,
	rawState map[string]json.RawMessage,
	cdc codec.Codec,
	bondDenom string,
) *evmtypes.GenesisState {
	evmStateBz, ok := rawState[evmtypes.ModuleName]
	require.True(t, ok, "evm genesis state is missing")

	evmState := new(evmtypes.GenesisState)
	cdc.MustUnmarshalJSON(evmStateBz, evmState)

	// replace the EvmDenom with BondDenom
	evmState.Params.EvmDenom = bondDenom

	return evmState
}

// updateStakingState updates the staking genesis state.
// It adds the following
// - The not bonded balance for the not bonded pool
// It additionally returns the non-bonded coins as well
func updateStakingState(
	t *testing.T,
	rawState map[string]json.RawMessage,
	cdc codec.Codec,
) (*stakingtypes.GenesisState, sdk.Coin) {
	stakingStateBz, ok := rawState[stakingtypes.ModuleName]
	require.True(t, ok, "staking genesis state is missing")

	stakingState := new(stakingtypes.GenesisState)
	err := cdc.UnmarshalJSON(stakingStateBz, stakingState)
	if err != nil {
		panic(err)
	}

	// compute not bonded balance
	notBondedTokens := math.ZeroInt()
	for _, val := range stakingState.Validators {
		if val.Status != stakingtypes.Unbonded {
			continue
		}
		notBondedTokens = notBondedTokens.Add(val.GetTokens())
	}
	notBondedCoins := sdk.NewCoin(stakingState.Params.BondDenom, notBondedTokens)

	return stakingState, notBondedCoins
}

// updateObserverState updates the observer genesis state.
// It adds the following
// - A random observer set which is a subset of the current validator set
// - A randomised node account for each observer
// - A random TSS
// - A TSS history for the TSS created
// - Chain nonces for each chain
// - Pending nonces for each chain
// - Crosschain flags, inbound and outbound enabled
func updateObserverState(
	t *testing.T,
	rawState map[string]json.RawMessage,
	cdc codec.Codec,
	r *rand.Rand,
	validators stakingtypes.Validators,
) *observertypes.GenesisState {
	observerStateBz, ok := rawState[observertypes.ModuleName]
	require.True(t, ok, "observer genesis state is missing")

	observerState := new(observertypes.GenesisState)
	cdc.MustUnmarshalJSON(observerStateBz, observerState)

	// Create an observer set as a subset of the current validator set
	observers := make([]string, 0)
	for _, validator := range validators {
		accAddress, err := observertypes.GetAccAddressFromOperatorAddress(validator.OperatorAddress)
		if err != nil {
			continue
		}
		observers = append(observers, accAddress.String())
	}

	r.Shuffle(len(observers), func(i, j int) {
		observers[i], observers[j] = observers[j], observers[i]
	})

	numObservers := r.Intn(21) + 5
	if numObservers > len(observers) {
		numObservers = len(observers)
	}
	observers = observers[:numObservers]

	// Create node account list for the observers set
	nodeAccounts := make([]*observertypes.NodeAccount, len(observers))
	for i, observer := range observers {
		nodeAccounts[i] = &observertypes.NodeAccount{
			Operator:       observer,
			GranteeAddress: observer,
			GranteePubkey:  &crypto.PubKeySet{},
			NodeStatus:     observertypes.NodeStatus_Active,
		}
	}
	// Create a random tss
	tss, err := sample.TSSFromRand(r)
	require.NoError(t, err)
	tss.OperatorAddressList = observers

	// Create a tss history
	tssHistory := make([]observertypes.TSS, 0)
	tssHistory = append(tssHistory, tss)

	// Create chainnonces and pendingnonces
	chains := zetachains.DefaultChainsList()
	chainsNonces := make([]observertypes.ChainNonces, 0)
	pendingNonces := make([]observertypes.PendingNonces, 0)
	for _, chain := range chains {
		chainNonce := observertypes.ChainNonces{
			ChainId: chain.ChainId,
			Nonce:   0,
		}
		chainsNonces = append(chainsNonces, chainNonce)
		pendingNonce := observertypes.PendingNonces{
			NonceLow:  0,
			NonceHigh: 0,
			ChainId:   chain.ChainId,
			Tss:       tss.TssPubkey,
		}
		pendingNonces = append(pendingNonces, pendingNonce)
	}

	observerState.Tss = &tss
	observerState.Observers.ObserverList = observers
	observerState.NodeAccountList = nodeAccounts
	observerState.CrosschainFlags.IsInboundEnabled = true
	observerState.CrosschainFlags.IsOutboundEnabled = true
	observerState.ChainNonces = chainsNonces
	observerState.PendingNonces = pendingNonces
	observerState.TssHistory = tssHistory

	return observerState
}

// updateAuthorityState updates the authority genesis state.
// It adds the following
// - A policy for each policy type, the address is a random account address selected from the simulation accounts list
func updateAuthorityState(
	t *testing.T,
	rawState map[string]json.RawMessage,
	cdc codec.Codec,
	r *rand.Rand,
	accs []simtypes.Account,
) *authoritytypes.GenesisState {
	authorityStateBz, ok := rawState[authoritytypes.ModuleName]
	require.True(t, ok, "authority genesis state is missing")

	authorityState := new(authoritytypes.GenesisState)
	cdc.MustUnmarshalJSON(authorityStateBz, authorityState)

	randomAccount := accs[r.Intn(len(accs))]
	policies := authoritytypes.Policies{
		Items: []*authoritytypes.Policy{
			{
				Address:    randomAccount.Address.String(),
				PolicyType: authoritytypes.PolicyType_groupEmergency,
			},
			{
				Address:    randomAccount.Address.String(),
				PolicyType: authoritytypes.PolicyType_groupAdmin,
			},
			{
				Address:    randomAccount.Address.String(),
				PolicyType: authoritytypes.PolicyType_groupOperational,
			},
		},
	}
	authorityState.Policies = policies

	return authorityState
}

// updateCrosschainState updates the crosschain genesis state.
// It adds the following
// - A gas price list for each chain
func updateCrosschainState(
	t *testing.T,
	rawState map[string]json.RawMessage,
	cdc codec.Codec,
	r *rand.Rand,
) *crosschaintypes.GenesisState {
	crossChainStateBz, ok := rawState[crosschaintypes.ModuleName]
	require.True(t, ok, "crosschain genesis state is missing")

	crossChainState := new(crosschaintypes.GenesisState)
	cdc.MustUnmarshalJSON(crossChainStateBz, crossChainState)

	// Add a gasprice for each chain
	chains := zetachains.DefaultChainsList()
	gasPriceList := make([]*crosschaintypes.GasPrice, len(chains))
	for i, chain := range chains {
		gasPriceList[i] = sample.GasPriceFromRand(r, chain.ChainId)
	}

	crossChainState.GasPriceList = gasPriceList

	return crossChainState
}

// updateFungibleState updates the fungible genesis state.
// It adds the following
// - A random system contract address
// - A random connector zevm address
// - A random gateway address
// - A foreign coin for each chain under the default chain list.
func updateFungibleState(
	t *testing.T,
	rawState map[string]json.RawMessage,
	cdc codec.Codec,
	r *rand.Rand,
) *fungibletypes.GenesisState {
	fungibleStateBz, ok := rawState[fungibletypes.ModuleName]
	require.True(t, ok, "fungible genesis state is missing")

	fungibleState := new(fungibletypes.GenesisState)
	cdc.MustUnmarshalJSON(fungibleStateBz, fungibleState)
	fungibleState.SystemContract = &fungibletypes.SystemContract{
		SystemContract: sample.EthAddressFromRand(r).String(),
		ConnectorZevm:  sample.EthAddressFromRand(r).String(),
		Gateway:        sample.EthAddressFromRand(r).String(),
	}

	foreignCoins := make([]fungibletypes.ForeignCoins, 0)
	chains := zetachains.DefaultChainsList()

	for _, chain := range chains {
		foreignCoin := fungibletypes.ForeignCoins{
			ForeignChainId:       chain.ChainId,
			Asset:                sample.EthAddressFromRand(r).String(),
			Zrc20ContractAddress: sample.EthAddressFromRand(r).String(),
			Decimals:             18,
			Paused:               false,
			CoinType:             coin.CoinType_Gas,
			LiquidityCap:         math.ZeroUint(),
		}
		foreignCoins = append(foreignCoins, foreignCoin)
	}
	fungibleState.ForeignCoinsList = foreignCoins

	return fungibleState
}

// updateRawState updates the raw genesis state for the application.
// This is used to inject values needed to run the simulation tests.
func updateRawState(
	t *testing.T,
	rawState map[string]json.RawMessage,
	cdc codec.Codec,
	r *rand.Rand,
	accs []simtypes.Account,
) {
	stakingState, notBondedCoins := updateStakingState(t, rawState, cdc)
	bankState := updateBankState(t, rawState, cdc, notBondedCoins)
	evmState := updateEVMState(t, rawState, cdc, stakingState.Params.BondDenom)
	observerState := updateObserverState(t, rawState, cdc, r, stakingState.Validators)
	authorityState := updateAuthorityState(t, rawState, cdc, r, accs)
	fungibleState := updateFungibleState(t, rawState, cdc, r)

	rawState[stakingtypes.ModuleName] = cdc.MustMarshalJSON(stakingState)
	rawState[banktypes.ModuleName] = cdc.MustMarshalJSON(bankState)
	rawState[evmtypes.ModuleName] = cdc.MustMarshalJSON(evmState)
	rawState[observertypes.ModuleName] = cdc.MustMarshalJSON(observerState)
	rawState[authoritytypes.ModuleName] = cdc.MustMarshalJSON(authorityState)
	rawState[fungibletypes.ModuleName] = cdc.MustMarshalJSON(fungibleState)
	rawState[crosschaintypes.ModuleName] = cdc.MustMarshalJSON(updateCrosschainState(t, rawState, cdc, r))
}

// AppStateFn returns the initial application state using a genesis or the simulation parameters.
// It panics if the user provides files for both of them.
// If a file is not given for the genesis or the sim params, it creates a randomized one.
// All modifications to the genesis state should be done in this function.
func AppStateFn(
	t *testing.T,
	cdc codec.Codec,
	simManager *module.SimulationManager,
	genesisState map[string]json.RawMessage,
	exportedState json.RawMessage,
) simtypes.AppStateFn {
	return func(r *rand.Rand, accs []simtypes.Account, config simtypes.Config,
	) (appState json.RawMessage, simAccs []simtypes.Account, chainID string, genesisTimestamp time.Time) {
		if FlagGenesisTimeValue == 0 {
			genesisTimestamp = simtypes.RandTimestamp(r)
		} else {
			genesisTimestamp = time.Unix(FlagGenesisTimeValue, 0)
		}

		chainID = config.ChainID

		// if exported state is provided, then use it
		if exportedState != nil {
			return exportedState, accs, chainID, genesisTimestamp
		}

		appParams := make(simtypes.AppParams)
		appState, simAccs = AppStateRandomizedFn(
			simManager,
			r,
			cdc,
			accs,
			genesisTimestamp,
			appParams,
			genesisState,
		)

		rawState := make(map[string]json.RawMessage)
		err := json.Unmarshal(appState, &rawState)
		if err != nil {
			panic(err)
		}

		updateRawState(t, rawState, cdc, r, simAccs)

		// replace appstate
		appState, err = json.Marshal(rawState)
		require.NoError(t, err)

		return appState, simAccs, chainID, genesisTimestamp
	}
}

// AppStateRandomizedFn creates calls each module's GenesisState generator function
// and creates the simulation params
func AppStateRandomizedFn(
	simManager *module.SimulationManager, r *rand.Rand, cdc codec.Codec,
	accs []simtypes.Account, genesisTimestamp time.Time, appParams simtypes.AppParams,
	genesisState map[string]json.RawMessage,
) (json.RawMessage, []simtypes.Account) {
	numAccs := int64(len(accs))
	// generate a random amount of initial stake coins and a random initial
	// number of bonded accounts
	var (
		numInitiallyBonded int64
		initialStake       math.Int
	)

	appParams.GetOrGenerate(cdc,
		StakePerAccount, &initialStake, r,
		func(r *rand.Rand) { initialStake = math.NewInt(r.Int63n(1e12)) },
	)
	appParams.GetOrGenerate(cdc,
		InitiallyBondedValidators, &numInitiallyBonded, r,
		func(r *rand.Rand) { numInitiallyBonded = int64(r.Intn(300)) },
	)

	if numInitiallyBonded > numAccs {
		numInitiallyBonded = numAccs
	}

	// set the default power reduction to be one less than the initial stake so that all randomised validators are part of the validator set
	sdk.DefaultPowerReduction = initialStake.Sub(sdk.OneInt())

	fmt.Printf(
		`Selected randomly generated parameters for simulated genesis:
{
  stake_per_account: "%d",
  initially_bonded_validators: "%d"
}
`, initialStake, numInitiallyBonded,
	)

	simState := &module.SimulationState{
		AppParams:    appParams,
		Cdc:          cdc,
		Rand:         r,
		GenState:     genesisState,
		Accounts:     accs,
		InitialStake: initialStake,
		NumBonded:    numInitiallyBonded,
		GenTimestamp: genesisTimestamp,
	}

	simManager.GenerateGenesisStates(simState)

	appState, err := json.Marshal(genesisState)
	if err != nil {
		panic(err)
	}
	return appState, accs
}
