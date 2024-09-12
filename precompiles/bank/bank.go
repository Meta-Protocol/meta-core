package bank

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	ptypes "github.com/zeta-chain/node/precompiles/types"
	fungiblekeeper "github.com/zeta-chain/node/x/fungible/keeper"
)

var (
	ABI                 abi.ABI
	ContractAddress     = common.HexToAddress("0x0000000000000000000000000000000000000067")
	GasRequiredByMethod = map[[4]byte]uint64{}
	ViewMethod          = map[[4]byte]bool{}
)

func init() {
	initABI()
}

func initABI() {
	if err := ABI.UnmarshalJSON([]byte(IBankMetaData.ABI)); err != nil {
		panic(err)
	}

	GasRequiredByMethod = map[[4]byte]uint64{}
	for methodName := range ABI.Methods {
		var methodID [4]byte
		copy(methodID[:], ABI.Methods[methodName].ID[:4])
		switch methodName {
		case DepositMethodName:
			GasRequiredByMethod[methodID] = DepositMethodGas
		case WithdrawMethodName:
			GasRequiredByMethod[methodID] = WithdrawMethodGas
		case BalanceOfMethodName:
			GasRequiredByMethod[methodID] = BalanceOfGas
		default:
			GasRequiredByMethod[methodID] = DefaultGas
		}
	}
}

type Contract struct {
	ptypes.BaseContract

	bankKeeper     bank.Keeper
	fungibleKeeper fungiblekeeper.Keeper
	cdc            codec.Codec
	kvGasConfig    storetypes.GasConfig
}

func NewIBankContract(
	bankKeeper bank.Keeper,
	fungibleKeeper fungiblekeeper.Keeper,
	cdc codec.Codec,
	kvGasConfig storetypes.GasConfig,
) *Contract {
	return &Contract{
		BaseContract:   ptypes.NewBaseContract(ContractAddress),
		bankKeeper:     bankKeeper,
		fungibleKeeper: fungibleKeeper,
		cdc:            cdc,
		kvGasConfig:    kvGasConfig,
	}
}

// Address() is required to implement the PrecompiledContract interface.
func (c *Contract) Address() common.Address {
	return ContractAddress
}

// Abi() is required to implement the PrecompiledContract interface.
func (c *Contract) Abi() abi.ABI {
	return ABI
}

// RequiredGas is required to implement the PrecompiledContract interface.
// The gas has to be calculated deterministically based on the input.
func (c *Contract) RequiredGas(input []byte) uint64 {
	// get methodID (first 4 bytes)
	var methodID [4]byte
	copy(methodID[:], input[:4])
	// base cost to prevent large input size
	baseCost := uint64(len(input)) * c.kvGasConfig.WriteCostPerByte
	if ViewMethod[methodID] {
		baseCost = uint64(len(input)) * c.kvGasConfig.ReadCostPerByte
	}

	if requiredGas, ok := GasRequiredByMethod[methodID]; ok {
		return requiredGas + baseCost
	}

	// Can not happen, but return 0 if the method is not found.
	return 0
}

// Run is the entrypoint of the precompiled contract, it switches over the input method,
// and execute them accordingly.
func (c *Contract) Run(evm *vm.EVM, contract *vm.Contract, readOnly bool) ([]byte, error) {
	fmt.Println("DEBUG: bank.Run()")
	method, err := ABI.MethodById(contract.Input[:4])
	if err != nil {
		return nil, err
	}

	args, err := method.Inputs.Unpack(contract.Input[4:])
	if err != nil {
		return nil, err
	}

	stateDB := evm.StateDB.(ptypes.ExtStateDB)

	switch method.Name {
	case DepositMethodName:
		fmt.Println("DEBUG: bank.Run(): DepositMethodName")
		if readOnly {
			return nil, ptypes.ErrUnexpected{
				Got: "method not allowed in read-only mode " + method.Name,
			}
		}

		var res []byte
		execErr := stateDB.ExecuteNativeAction(contract.Address(), nil, func(ctx sdk.Context) error {
			fmt.Println("DEBUG: bank.Run(): DepositMethodName: ExecuteNativeAction c.deposit()")
			res, err = c.deposit(ctx, method, contract.CallerAddress, args)
			return err
		})
		if execErr != nil {
			fmt.Printf("DEBUG: bank.Run(): execErr %s", execErr.Error())
			return nil, err
		}
		return res, nil

	case WithdrawMethodName:
		if readOnly {
			return nil, ptypes.ErrUnexpected{
				Got: "method not allowed in read-only mode " + method.Name,
			}
		}

		return nil, nil
		// TODO

	case BalanceOfMethodName:
		fmt.Println("DEBUG: bank.Run(): BalanceOfMethodName")
		var res []byte
		execErr := stateDB.ExecuteNativeAction(contract.Address(), nil, func(ctx sdk.Context) error {
			fmt.Println("DEBUG: bank.Run(): DepositMethodName: ExecuteNativeAction c.balanceOf()")
			res, err = c.balanceOf(ctx, method, args)
			return err
		})
		if execErr != nil {
			return nil, err
		}
		return res, nil

	default:
		return nil, ptypes.ErrInvalidMethod{
			Method: method.Name,
		}
	}
}
