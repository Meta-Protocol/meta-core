// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testdappv2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TestDAppV2zContext is an auto generated low-level Go binding around an user-defined struct.
type TestDAppV2zContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// TestDAppV2MetaData contains all meta data concerning the TestDAppV2 contract.
var TestDAppV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"erc20Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"gasCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastContext\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMessage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastZRC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structTestDAppV2.zContext\",\"name\":\"context\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onCrossChainCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"simpleCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061126f806100206000396000f3fe60806040526004361061007b5760003560e01c8063b2f79b031161004e578063b2f79b031461011b578063b73f7eb114610146578063c7a339a914610173578063de43156e1461019c5761007b565b8063329707101461008057806336e980a0146100ab578063829a86d9146100d4578063a799911f146100ff575b600080fd5b34801561008c57600080fd5b506100956101c5565b6040516100a29190610ab2565b60405180910390f35b3480156100b757600080fd5b506100d260048036038101906100cd919061086d565b610253565b005b3480156100e057600080fd5b506100e9610280565b6040516100f69190610ad4565b60405180910390f35b6101196004803603810190610114919061086d565b610286565b005b34801561012757600080fd5b506101306102ba565b60405161013d9190610a22565b60405180910390f35b34801561015257600080fd5b5061015b6102e0565b60405161016a93929190610a74565b60405180910390f35b34801561017f57600080fd5b5061019a600480360381019061019591906107fe565b6103a0565b005b3480156101a857600080fd5b506101c360048036038101906101be91906108b6565b61046e565b005b600580546101d290610e9f565b80601f01602080910402602001604051908101604052809291908181526020018280546101fe90610e9f565b801561024b5780601f106102205761010080835404028352916020019161024b565b820191906000526020600020905b81548152906001019060200180831161022e57829003601f168201915b505050505081565b61025c81610538565b1561026657600080fd5b806005908051906020019061027c92919061056f565b5050565b60045481565b61028f81610538565b1561029957600080fd5b80600590805190602001906102af92919061056f565b503460048190555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060000180546102f190610e9f565b80601f016020809104026020016040519081016040528092919081815260200182805461031d90610e9f565b801561036a5780601f1061033f5761010080835404028352916020019161036a565b820191906000526020600020905b81548152906001019060200180831161034d57829003601f168201915b5050505050908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154905083565b6103a981610538565b156103b357600080fd5b8273ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b81526004016103f093929190610a3d565b602060405180830381600087803b15801561040a57600080fd5b505af115801561041e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061044291906107d1565b61044b57600080fd5b806005908051906020019061046192919061056f565b5081600481905550505050565b6104bb82828080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610538565b156104c557600080fd5b84600081816104d49190611182565b90505083600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826004819055508181600591906105309291906105f5565b505050505050565b600060405160200161054990610a0d565b604051602081830303815290604052805190602001208280519060200120149050919050565b82805461057b90610e9f565b90600052602060002090601f01602090048101928261059d57600085556105e4565b82601f106105b657805160ff19168380011785556105e4565b828001600101855582156105e4579182015b828111156105e35782518255916020019190600101906105c8565b5b5090506105f1919061067b565b5090565b82805461060190610e9f565b90600052602060002090601f016020900481019282610623576000855561066a565b82601f1061063c57803560ff191683800117855561066a565b8280016001018555821561066a579182015b8281111561066957823582559160200191906001019061064e565b5b509050610677919061067b565b5090565b5b8082111561069457600081600090555060010161067c565b5090565b60006106ab6106a684610b77565b610b52565b9050828152602081018484840111156106c7576106c6610ffd565b5b6106d2848285610de6565b509392505050565b6000813590506106e9816111d8565b92915050565b6000815190506106fe816111ef565b92915050565b60008083601f84011261071a57610719610fdf565b5b8235905067ffffffffffffffff81111561073757610736610fda565b5b60208301915083600182028301111561075357610752610ff3565b5b9250929050565b60008135905061076981611206565b92915050565b600082601f83011261078457610783610fdf565b5b8135610794848260208601610698565b91505092915050565b6000606082840312156107b3576107b2610fe9565b5b81905092915050565b6000813590506107cb8161121d565b92915050565b6000602082840312156107e7576107e6611007565b5b60006107f5848285016106ef565b91505092915050565b60008060006060848603121561081757610816611007565b5b60006108258682870161075a565b9350506020610836868287016107bc565b925050604084013567ffffffffffffffff81111561085757610856611002565b5b6108638682870161076f565b9150509250925092565b60006020828403121561088357610882611007565b5b600082013567ffffffffffffffff8111156108a1576108a0611002565b5b6108ad8482850161076f565b91505092915050565b6000806000806000608086880312156108d2576108d1611007565b5b600086013567ffffffffffffffff8111156108f0576108ef611002565b5b6108fc8882890161079d565b955050602061090d888289016106da565b945050604061091e888289016107bc565b935050606086013567ffffffffffffffff81111561093f5761093e611002565b5b61094b88828901610704565b92509250509295509295909350565b61096381610c51565b82525050565b600061097482610bc8565b61097e8185610bde565b935061098e818560208601610df5565b6109978161100c565b840191505092915050565b60006109ad82610bd3565b6109b78185610bef565b93506109c7818560208601610df5565b6109d08161100c565b840191505092915050565b60006109e8600683610c00565b91506109f38261105c565b600682019050919050565b610a0781610ca1565b82525050565b6000610a18826109db565b9150819050919050565b6000602082019050610a37600083018461095a565b92915050565b6000606082019050610a52600083018661095a565b610a5f602083018561095a565b610a6c60408301846109fe565b949350505050565b60006060820190508181036000830152610a8e8186610969565b9050610a9d602083018561095a565b610aaa60408301846109fe565b949350505050565b60006020820190508181036000830152610acc81846109a2565b905092915050565b6000602082019050610ae960008301846109fe565b92915050565b60008083356001602003843603038112610b0c57610b0b610fee565b5b80840192508235915067ffffffffffffffff821115610b2e57610b2d610fe4565b5b602083019250600182023603831315610b4a57610b49610ff8565b5b509250929050565b6000610b5c610b6d565b9050610b688282610eed565b919050565b6000604051905090565b600067ffffffffffffffff821115610b9257610b91610f6b565b5b610b9b8261100c565b9050602081019050919050565b60008190508160005260206000209050919050565b600082905092915050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b601f821115610c4c57610c1d81610ba8565b610c2684610e8f565b81016020851015610c35578190505b610c49610c4185610e8f565b830182610cab565b50505b505050565b6000610c5c82610c81565b9050919050565b60008115159050919050565b6000610c7a82610c51565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b5b81811015610cca57610cbf600082611044565b600181019050610cac565b5050565b6000610cd982610ce0565b9050919050565b6000610ceb82610cf2565b9050919050565b6000610cfd82610c81565b9050919050565b6000610d0f82610ca1565b9050919050565b610d208383610bbd565b67ffffffffffffffff811115610d3957610d38610f6b565b5b610d438254610e9f565b610d4e828285610c0b565b6000601f831160018114610d7d5760008415610d6b578287013590505b610d758582610ed1565b865550610ddd565b601f198416610d8b86610ba8565b60005b82811015610db357848901358255600182019150602085019450602081019050610d8e565b86831015610dd05784890135610dcc601f891682610f1e565b8355505b6001600288020188555050505b50505050505050565b82818337600083830152505050565b60005b83811015610e13578082015181840152602081019050610df8565b83811115610e22576000848401525b50505050565b6000810160008301610e3a8185610aef565b610e45818386611172565b50505050600181016020830180610e5b81610fae565b9050610e67818461114f565b505050600281016040830180610e7c81610fc4565b9050610e888184611190565b5050505050565b60006020601f8301049050919050565b60006002820490506001821680610eb757607f821691505b60208210811415610ecb57610eca610f3c565b5b50919050565b6000610edd8383610f1e565b9150826002028217905092915050565b610ef68261100c565b810181811067ffffffffffffffff82111715610f1557610f14610f6b565b5b80604052505050565b6000610f2f60001984600802611037565b1980831691505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000819050919050565b6000819050919050565b60008135610fbb816111d8565b80915050919050565b60008135610fd18161121d565b80915050919050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160001b9050919050565b600082821b905092915050565b600082821c905092915050565b61104c611234565b6110578184846111b3565b505050565b7f7265766572740000000000000000000000000000000000000000000000000000600082015250565b600073ffffffffffffffffffffffffffffffffffffffff6110a58461101d565b9350801983169250808416831791505092915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6110e78461101d565b9350801983169250808416831791505092915050565b60006008830261112d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261102a565b611137868361102a565b95508019841693508086168417925050509392505050565b61115882610cce565b61116b61116482610f9a565b8354611085565b8255505050565b61117d838383610d16565b505050565b61118c8282610e28565b5050565b61119982610d04565b6111ac6111a582610fa4565b83546110bb565b8255505050565b6111bc83610d04565b6111d06111c882610fa4565b8484546110fd565b825550505050565b6111e181610c51565b81146111ec57600080fd5b50565b6111f881610c63565b811461120357600080fd5b50565b61120f81610c6f565b811461121a57600080fd5b50565b61122681610ca1565b811461123157600080fd5b50565b60009056fea2646970667358221220cbcd3ee88401ed7a1eed4a04915d4a448b89997007ceb996d717aaece6d89a7364736f6c63430008070033",
}

// TestDAppV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use TestDAppV2MetaData.ABI instead.
var TestDAppV2ABI = TestDAppV2MetaData.ABI

// TestDAppV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestDAppV2MetaData.Bin instead.
var TestDAppV2Bin = TestDAppV2MetaData.Bin

// DeployTestDAppV2 deploys a new Ethereum contract, binding an instance of TestDAppV2 to it.
func DeployTestDAppV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestDAppV2, error) {
	parsed, err := TestDAppV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestDAppV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestDAppV2{TestDAppV2Caller: TestDAppV2Caller{contract: contract}, TestDAppV2Transactor: TestDAppV2Transactor{contract: contract}, TestDAppV2Filterer: TestDAppV2Filterer{contract: contract}}, nil
}

// TestDAppV2 is an auto generated Go binding around an Ethereum contract.
type TestDAppV2 struct {
	TestDAppV2Caller     // Read-only binding to the contract
	TestDAppV2Transactor // Write-only binding to the contract
	TestDAppV2Filterer   // Log filterer for contract events
}

// TestDAppV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestDAppV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDAppV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestDAppV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDAppV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestDAppV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDAppV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestDAppV2Session struct {
	Contract     *TestDAppV2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestDAppV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestDAppV2CallerSession struct {
	Contract *TestDAppV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestDAppV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestDAppV2TransactorSession struct {
	Contract     *TestDAppV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestDAppV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestDAppV2Raw struct {
	Contract *TestDAppV2 // Generic contract binding to access the raw methods on
}

// TestDAppV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestDAppV2CallerRaw struct {
	Contract *TestDAppV2Caller // Generic read-only contract binding to access the raw methods on
}

// TestDAppV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestDAppV2TransactorRaw struct {
	Contract *TestDAppV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestDAppV2 creates a new instance of TestDAppV2, bound to a specific deployed contract.
func NewTestDAppV2(address common.Address, backend bind.ContractBackend) (*TestDAppV2, error) {
	contract, err := bindTestDAppV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestDAppV2{TestDAppV2Caller: TestDAppV2Caller{contract: contract}, TestDAppV2Transactor: TestDAppV2Transactor{contract: contract}, TestDAppV2Filterer: TestDAppV2Filterer{contract: contract}}, nil
}

// NewTestDAppV2Caller creates a new read-only instance of TestDAppV2, bound to a specific deployed contract.
func NewTestDAppV2Caller(address common.Address, caller bind.ContractCaller) (*TestDAppV2Caller, error) {
	contract, err := bindTestDAppV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestDAppV2Caller{contract: contract}, nil
}

// NewTestDAppV2Transactor creates a new write-only instance of TestDAppV2, bound to a specific deployed contract.
func NewTestDAppV2Transactor(address common.Address, transactor bind.ContractTransactor) (*TestDAppV2Transactor, error) {
	contract, err := bindTestDAppV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestDAppV2Transactor{contract: contract}, nil
}

// NewTestDAppV2Filterer creates a new log filterer instance of TestDAppV2, bound to a specific deployed contract.
func NewTestDAppV2Filterer(address common.Address, filterer bind.ContractFilterer) (*TestDAppV2Filterer, error) {
	contract, err := bindTestDAppV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestDAppV2Filterer{contract: contract}, nil
}

// bindTestDAppV2 binds a generic wrapper to an already deployed contract.
func bindTestDAppV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestDAppV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestDAppV2 *TestDAppV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestDAppV2.Contract.TestDAppV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestDAppV2 *TestDAppV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestDAppV2.Contract.TestDAppV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestDAppV2 *TestDAppV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestDAppV2.Contract.TestDAppV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestDAppV2 *TestDAppV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestDAppV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestDAppV2 *TestDAppV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestDAppV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestDAppV2 *TestDAppV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestDAppV2.Contract.contract.Transact(opts, method, params...)
}

// LastAmount is a free data retrieval call binding the contract method 0x829a86d9.
//
// Solidity: function lastAmount() view returns(uint256)
func (_TestDAppV2 *TestDAppV2Caller) LastAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "lastAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAmount is a free data retrieval call binding the contract method 0x829a86d9.
//
// Solidity: function lastAmount() view returns(uint256)
func (_TestDAppV2 *TestDAppV2Session) LastAmount() (*big.Int, error) {
	return _TestDAppV2.Contract.LastAmount(&_TestDAppV2.CallOpts)
}

// LastAmount is a free data retrieval call binding the contract method 0x829a86d9.
//
// Solidity: function lastAmount() view returns(uint256)
func (_TestDAppV2 *TestDAppV2CallerSession) LastAmount() (*big.Int, error) {
	return _TestDAppV2.Contract.LastAmount(&_TestDAppV2.CallOpts)
}

// LastContext is a free data retrieval call binding the contract method 0xb73f7eb1.
//
// Solidity: function lastContext() view returns(bytes origin, address sender, uint256 chainID)
func (_TestDAppV2 *TestDAppV2Caller) LastContext(opts *bind.CallOpts) (struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "lastContext")

	outstruct := new(struct {
		Origin  []byte
		Sender  common.Address
		ChainID *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Origin = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Sender = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ChainID = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LastContext is a free data retrieval call binding the contract method 0xb73f7eb1.
//
// Solidity: function lastContext() view returns(bytes origin, address sender, uint256 chainID)
func (_TestDAppV2 *TestDAppV2Session) LastContext() (struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}, error) {
	return _TestDAppV2.Contract.LastContext(&_TestDAppV2.CallOpts)
}

// LastContext is a free data retrieval call binding the contract method 0xb73f7eb1.
//
// Solidity: function lastContext() view returns(bytes origin, address sender, uint256 chainID)
func (_TestDAppV2 *TestDAppV2CallerSession) LastContext() (struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}, error) {
	return _TestDAppV2.Contract.LastContext(&_TestDAppV2.CallOpts)
}

// LastMessage is a free data retrieval call binding the contract method 0x32970710.
//
// Solidity: function lastMessage() view returns(string)
func (_TestDAppV2 *TestDAppV2Caller) LastMessage(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "lastMessage")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LastMessage is a free data retrieval call binding the contract method 0x32970710.
//
// Solidity: function lastMessage() view returns(string)
func (_TestDAppV2 *TestDAppV2Session) LastMessage() (string, error) {
	return _TestDAppV2.Contract.LastMessage(&_TestDAppV2.CallOpts)
}

// LastMessage is a free data retrieval call binding the contract method 0x32970710.
//
// Solidity: function lastMessage() view returns(string)
func (_TestDAppV2 *TestDAppV2CallerSession) LastMessage() (string, error) {
	return _TestDAppV2.Contract.LastMessage(&_TestDAppV2.CallOpts)
}

// LastZRC20 is a free data retrieval call binding the contract method 0xb2f79b03.
//
// Solidity: function lastZRC20() view returns(address)
func (_TestDAppV2 *TestDAppV2Caller) LastZRC20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "lastZRC20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastZRC20 is a free data retrieval call binding the contract method 0xb2f79b03.
//
// Solidity: function lastZRC20() view returns(address)
func (_TestDAppV2 *TestDAppV2Session) LastZRC20() (common.Address, error) {
	return _TestDAppV2.Contract.LastZRC20(&_TestDAppV2.CallOpts)
}

// LastZRC20 is a free data retrieval call binding the contract method 0xb2f79b03.
//
// Solidity: function lastZRC20() view returns(address)
func (_TestDAppV2 *TestDAppV2CallerSession) LastZRC20() (common.Address, error) {
	return _TestDAppV2.Contract.LastZRC20(&_TestDAppV2.CallOpts)
}

// Erc20Call is a paid mutator transaction binding the contract method 0xc7a339a9.
//
// Solidity: function erc20Call(address erc20, uint256 amount, string message) returns()
func (_TestDAppV2 *TestDAppV2Transactor) Erc20Call(opts *bind.TransactOpts, erc20 common.Address, amount *big.Int, message string) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "erc20Call", erc20, amount, message)
}

// Erc20Call is a paid mutator transaction binding the contract method 0xc7a339a9.
//
// Solidity: function erc20Call(address erc20, uint256 amount, string message) returns()
func (_TestDAppV2 *TestDAppV2Session) Erc20Call(erc20 common.Address, amount *big.Int, message string) (*types.Transaction, error) {
	return _TestDAppV2.Contract.Erc20Call(&_TestDAppV2.TransactOpts, erc20, amount, message)
}

// Erc20Call is a paid mutator transaction binding the contract method 0xc7a339a9.
//
// Solidity: function erc20Call(address erc20, uint256 amount, string message) returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) Erc20Call(erc20 common.Address, amount *big.Int, message string) (*types.Transaction, error) {
	return _TestDAppV2.Contract.Erc20Call(&_TestDAppV2.TransactOpts, erc20, amount, message)
}

// GasCall is a paid mutator transaction binding the contract method 0xa799911f.
//
// Solidity: function gasCall(string message) payable returns()
func (_TestDAppV2 *TestDAppV2Transactor) GasCall(opts *bind.TransactOpts, message string) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "gasCall", message)
}

// GasCall is a paid mutator transaction binding the contract method 0xa799911f.
//
// Solidity: function gasCall(string message) payable returns()
func (_TestDAppV2 *TestDAppV2Session) GasCall(message string) (*types.Transaction, error) {
	return _TestDAppV2.Contract.GasCall(&_TestDAppV2.TransactOpts, message)
}

// GasCall is a paid mutator transaction binding the contract method 0xa799911f.
//
// Solidity: function gasCall(string message) payable returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) GasCall(message string) (*types.Transaction, error) {
	return _TestDAppV2.Contract.GasCall(&_TestDAppV2.TransactOpts, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestDAppV2 *TestDAppV2Transactor) OnCrossChainCall(opts *bind.TransactOpts, context TestDAppV2zContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "onCrossChainCall", context, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestDAppV2 *TestDAppV2Session) OnCrossChainCall(context TestDAppV2zContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestDAppV2.Contract.OnCrossChainCall(&_TestDAppV2.TransactOpts, context, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) OnCrossChainCall(context TestDAppV2zContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestDAppV2.Contract.OnCrossChainCall(&_TestDAppV2.TransactOpts, context, zrc20, amount, message)
}

// SimpleCall is a paid mutator transaction binding the contract method 0x36e980a0.
//
// Solidity: function simpleCall(string message) returns()
func (_TestDAppV2 *TestDAppV2Transactor) SimpleCall(opts *bind.TransactOpts, message string) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "simpleCall", message)
}

// SimpleCall is a paid mutator transaction binding the contract method 0x36e980a0.
//
// Solidity: function simpleCall(string message) returns()
func (_TestDAppV2 *TestDAppV2Session) SimpleCall(message string) (*types.Transaction, error) {
	return _TestDAppV2.Contract.SimpleCall(&_TestDAppV2.TransactOpts, message)
}

// SimpleCall is a paid mutator transaction binding the contract method 0x36e980a0.
//
// Solidity: function simpleCall(string message) returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) SimpleCall(message string) (*types.Transaction, error) {
	return _TestDAppV2.Contract.SimpleCall(&_TestDAppV2.TransactOpts, message)
}
