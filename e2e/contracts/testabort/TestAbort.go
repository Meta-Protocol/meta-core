// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testabort

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

// AbortContext is an auto generated low-level Go binding around an user-defined struct.
type AbortContext struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}

// TestAbortMetaData contains all meta data concerning the TestAbort contract.
var TestAbortMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"abortedWithMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"outgoing\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"revertMessage\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"getAbortedWithMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"outgoing\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"revertMessage\",\"type\":\"bytes\"}],\"internalType\":\"structAbortContext\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"outgoing\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"revertMessage\",\"type\":\"bytes\"}],\"internalType\":\"structAbortContext\",\"name\":\"abortContext\",\"type\":\"tuple\"}],\"name\":\"onAbort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506110f78061001f6000396000f3fe6080604052600436106100385760003560e01c80632d4cfb7e1461004157806372748f7d1461006a5780639e59f463146100ac5761003f565b3661003f57005b005b34801561004d57600080fd5b5061006860048036038101906100639190610632565b6100e9565b005b34801561007657600080fd5b50610091600480360381019061008c91906106b1565b610151565b6040516100a3969594939291906107e3565b60405180910390f35b3480156100b857600080fd5b506100d360048036038101906100ce9190610987565b6102ca565b6040516100e09190610ad7565b60405180910390f35b61014e818060a001906100fc9190610b08565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508261014990610d72565b6104c5565b50565b600060205280600052604060002060009150905080600001805461017490610db4565b80601f01602080910402602001604051908101604052809291908181526020018280546101a090610db4565b80156101ed5780601f106101c2576101008083540402835291602001916101ed565b820191906000526020600020905b8154815290600101906020018083116101d057829003601f168201915b5050505050908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030160009054906101000a900460ff169080600401549080600501805461024790610db4565b80601f016020809104026020016040519081016040528092919081815260200182805461027390610db4565b80156102c05780601f10610295576101008083540402835291602001916102c0565b820191906000526020600020905b8154815290600101906020018083116102a357829003601f168201915b5050505050905086565b6102d26105ac565b600080836040516020016102e69190610e2c565b6040516020818303038152906040528051906020012081526020019081526020016000206040518060c001604052908160008201805461032590610db4565b80601f016020809104026020016040519081016040528092919081815260200182805461035190610db4565b801561039e5780601f106103735761010080835404028352916020019161039e565b820191906000526020600020905b81548152906001019060200180831161038157829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900460ff161515151581526020016004820154815260200160058201805461043c90610db4565b80601f016020809104026020016040519081016040528092919081815260200182805461046890610db4565b80156104b55780601f1061048a576101008083540402835291602001916104b5565b820191906000526020600020905b81548152906001019060200180831161049857829003601f168201915b5050505050815250509050919050565b80600080846040516020016104da9190610e2c565b60405160208183030381529060405280519060200120815260200190815260200160002060008201518160000190816105139190610fef565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030160006101000a81548160ff0219169083151502179055506080820151816004015560a08201518160050190816105a49190610fef565b509050505050565b6040518060c0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160001515815260200160008152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600060c082840312156106295761062861060e565b5b81905092915050565b60006020828403121561064857610647610604565b5b600082013567ffffffffffffffff81111561066657610665610609565b5b61067284828501610613565b91505092915050565b6000819050919050565b61068e8161067b565b811461069957600080fd5b50565b6000813590506106ab81610685565b92915050565b6000602082840312156106c7576106c6610604565b5b60006106d58482850161069c565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156107185780820151818401526020810190506106fd565b60008484015250505050565b6000601f19601f8301169050919050565b6000610740826106de565b61074a81856106e9565b935061075a8185602086016106fa565b61076381610724565b840191505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006107998261076e565b9050919050565b6107a98161078e565b82525050565b6000819050919050565b6107c2816107af565b82525050565b60008115159050919050565b6107dd816107c8565b82525050565b600060c08201905081810360008301526107fd8189610735565b905061080c60208301886107a0565b61081960408301876107b9565b61082660608301866107d4565b61083360808301856107b9565b81810360a08301526108458184610735565b9050979650505050505050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61089482610724565b810181811067ffffffffffffffff821117156108b3576108b261085c565b5b80604052505050565b60006108c66105fa565b90506108d2828261088b565b919050565b600067ffffffffffffffff8211156108f2576108f161085c565b5b6108fb82610724565b9050602081019050919050565b82818337600083830152505050565b600061092a610925846108d7565b6108bc565b90508281526020810184848401111561094657610945610857565b5b610951848285610908565b509392505050565b600082601f83011261096e5761096d610852565b5b813561097e848260208601610917565b91505092915050565b60006020828403121561099d5761099c610604565b5b600082013567ffffffffffffffff8111156109bb576109ba610609565b5b6109c784828501610959565b91505092915050565b600082825260208201905092915050565b60006109ec826106de565b6109f681856109d0565b9350610a068185602086016106fa565b610a0f81610724565b840191505092915050565b610a238161078e565b82525050565b610a32816107af565b82525050565b610a41816107c8565b82525050565b600060c0830160008301518482036000860152610a6482826109e1565b9150506020830151610a796020860182610a1a565b506040830151610a8c6040860182610a29565b506060830151610a9f6060860182610a38565b506080830151610ab26080860182610a29565b5060a083015184820360a0860152610aca82826109e1565b9150508091505092915050565b60006020820190508181036000830152610af18184610a47565b905092915050565b600080fd5b600080fd5b600080fd5b60008083356001602003843603038112610b2557610b24610af9565b5b80840192508235915067ffffffffffffffff821115610b4757610b46610afe565b5b602083019250600182023603831315610b6357610b62610b03565b5b509250929050565b600080fd5b600080fd5b600067ffffffffffffffff821115610b9057610b8f61085c565b5b610b9982610724565b9050602081019050919050565b6000610bb9610bb484610b75565b6108bc565b905082815260208101848484011115610bd557610bd4610857565b5b610be0848285610908565b509392505050565b600082601f830112610bfd57610bfc610852565b5b8135610c0d848260208601610ba6565b91505092915050565b610c1f8161078e565b8114610c2a57600080fd5b50565b600081359050610c3c81610c16565b92915050565b610c4b816107af565b8114610c5657600080fd5b50565b600081359050610c6881610c42565b92915050565b610c77816107c8565b8114610c8257600080fd5b50565b600081359050610c9481610c6e565b92915050565b600060c08284031215610cb057610caf610b6b565b5b610cba60c06108bc565b9050600082013567ffffffffffffffff811115610cda57610cd9610b70565b5b610ce684828501610be8565b6000830152506020610cfa84828501610c2d565b6020830152506040610d0e84828501610c59565b6040830152506060610d2284828501610c85565b6060830152506080610d3684828501610c59565b60808301525060a082013567ffffffffffffffff811115610d5a57610d59610b70565b5b610d6684828501610be8565b60a08301525092915050565b6000610d7e3683610c9a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610dcc57607f821691505b602082108103610ddf57610dde610d85565b5b50919050565b600081519050919050565b600081905092915050565b6000610e0682610de5565b610e108185610df0565b9350610e208185602086016106fa565b80840191505092915050565b6000610e388284610dfb565b915081905092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610ea57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610e68565b610eaf8683610e68565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610eec610ee7610ee2846107af565b610ec7565b6107af565b9050919050565b6000819050919050565b610f0683610ed1565b610f1a610f1282610ef3565b848454610e75565b825550505050565b600090565b610f2f610f22565b610f3a818484610efd565b505050565b5b81811015610f5e57610f53600082610f27565b600181019050610f40565b5050565b601f821115610fa357610f7481610e43565b610f7d84610e58565b81016020851015610f8c578190505b610fa0610f9885610e58565b830182610f3f565b50505b505050565b600082821c905092915050565b6000610fc660001984600802610fa8565b1980831691505092915050565b6000610fdf8383610fb5565b9150826002028217905092915050565b610ff8826106de565b67ffffffffffffffff8111156110115761101061085c565b5b61101b8254610db4565b611026828285610f62565b600060209050601f8311600181146110595760008415611047578287015190505b6110518582610fd3565b8655506110b9565b601f19841661106786610e43565b60005b8281101561108f5784890151825560018201915060208501945060208101905061106a565b868310156110ac57848901516110a8601f891682610fb5565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220ff750307c1447daa43b8c585df9b440046876431c31070cc22172c58e835efa064736f6c634300081a0033",
}

// TestAbortABI is the input ABI used to generate the binding from.
// Deprecated: Use TestAbortMetaData.ABI instead.
var TestAbortABI = TestAbortMetaData.ABI

// TestAbortBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestAbortMetaData.Bin instead.
var TestAbortBin = TestAbortMetaData.Bin

// DeployTestAbort deploys a new Ethereum contract, binding an instance of TestAbort to it.
func DeployTestAbort(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestAbort, error) {
	parsed, err := TestAbortMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestAbortBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestAbort{TestAbortCaller: TestAbortCaller{contract: contract}, TestAbortTransactor: TestAbortTransactor{contract: contract}, TestAbortFilterer: TestAbortFilterer{contract: contract}}, nil
}

// TestAbort is an auto generated Go binding around an Ethereum contract.
type TestAbort struct {
	TestAbortCaller     // Read-only binding to the contract
	TestAbortTransactor // Write-only binding to the contract
	TestAbortFilterer   // Log filterer for contract events
}

// TestAbortCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestAbortCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAbortTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestAbortTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAbortFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestAbortFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAbortSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestAbortSession struct {
	Contract     *TestAbort        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestAbortCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestAbortCallerSession struct {
	Contract *TestAbortCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TestAbortTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestAbortTransactorSession struct {
	Contract     *TestAbortTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestAbortRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestAbortRaw struct {
	Contract *TestAbort // Generic contract binding to access the raw methods on
}

// TestAbortCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestAbortCallerRaw struct {
	Contract *TestAbortCaller // Generic read-only contract binding to access the raw methods on
}

// TestAbortTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestAbortTransactorRaw struct {
	Contract *TestAbortTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestAbort creates a new instance of TestAbort, bound to a specific deployed contract.
func NewTestAbort(address common.Address, backend bind.ContractBackend) (*TestAbort, error) {
	contract, err := bindTestAbort(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestAbort{TestAbortCaller: TestAbortCaller{contract: contract}, TestAbortTransactor: TestAbortTransactor{contract: contract}, TestAbortFilterer: TestAbortFilterer{contract: contract}}, nil
}

// NewTestAbortCaller creates a new read-only instance of TestAbort, bound to a specific deployed contract.
func NewTestAbortCaller(address common.Address, caller bind.ContractCaller) (*TestAbortCaller, error) {
	contract, err := bindTestAbort(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestAbortCaller{contract: contract}, nil
}

// NewTestAbortTransactor creates a new write-only instance of TestAbort, bound to a specific deployed contract.
func NewTestAbortTransactor(address common.Address, transactor bind.ContractTransactor) (*TestAbortTransactor, error) {
	contract, err := bindTestAbort(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestAbortTransactor{contract: contract}, nil
}

// NewTestAbortFilterer creates a new log filterer instance of TestAbort, bound to a specific deployed contract.
func NewTestAbortFilterer(address common.Address, filterer bind.ContractFilterer) (*TestAbortFilterer, error) {
	contract, err := bindTestAbort(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestAbortFilterer{contract: contract}, nil
}

// bindTestAbort binds a generic wrapper to an already deployed contract.
func bindTestAbort(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestAbortMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAbort *TestAbortRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAbort.Contract.TestAbortCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAbort *TestAbortRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAbort.Contract.TestAbortTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAbort *TestAbortRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAbort.Contract.TestAbortTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAbort *TestAbortCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAbort.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAbort *TestAbortTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAbort.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAbort *TestAbortTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAbort.Contract.contract.Transact(opts, method, params...)
}

// AbortedWithMessage is a free data retrieval call binding the contract method 0x72748f7d.
//
// Solidity: function abortedWithMessage(bytes32 ) view returns(bytes sender, address asset, uint256 amount, bool outgoing, uint256 chainID, bytes revertMessage)
func (_TestAbort *TestAbortCaller) AbortedWithMessage(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}, error) {
	var out []interface{}
	err := _TestAbort.contract.Call(opts, &out, "abortedWithMessage", arg0)

	outstruct := new(struct {
		Sender        []byte
		Asset         common.Address
		Amount        *big.Int
		Outgoing      bool
		ChainID       *big.Int
		RevertMessage []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Asset = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Outgoing = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.ChainID = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RevertMessage = *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return *outstruct, err

}

// AbortedWithMessage is a free data retrieval call binding the contract method 0x72748f7d.
//
// Solidity: function abortedWithMessage(bytes32 ) view returns(bytes sender, address asset, uint256 amount, bool outgoing, uint256 chainID, bytes revertMessage)
func (_TestAbort *TestAbortSession) AbortedWithMessage(arg0 [32]byte) (struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}, error) {
	return _TestAbort.Contract.AbortedWithMessage(&_TestAbort.CallOpts, arg0)
}

// AbortedWithMessage is a free data retrieval call binding the contract method 0x72748f7d.
//
// Solidity: function abortedWithMessage(bytes32 ) view returns(bytes sender, address asset, uint256 amount, bool outgoing, uint256 chainID, bytes revertMessage)
func (_TestAbort *TestAbortCallerSession) AbortedWithMessage(arg0 [32]byte) (struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}, error) {
	return _TestAbort.Contract.AbortedWithMessage(&_TestAbort.CallOpts, arg0)
}

// GetAbortedWithMessage is a free data retrieval call binding the contract method 0x9e59f463.
//
// Solidity: function getAbortedWithMessage(string message) view returns((bytes,address,uint256,bool,uint256,bytes))
func (_TestAbort *TestAbortCaller) GetAbortedWithMessage(opts *bind.CallOpts, message string) (AbortContext, error) {
	var out []interface{}
	err := _TestAbort.contract.Call(opts, &out, "getAbortedWithMessage", message)

	if err != nil {
		return *new(AbortContext), err
	}

	out0 := *abi.ConvertType(out[0], new(AbortContext)).(*AbortContext)

	return out0, err

}

// GetAbortedWithMessage is a free data retrieval call binding the contract method 0x9e59f463.
//
// Solidity: function getAbortedWithMessage(string message) view returns((bytes,address,uint256,bool,uint256,bytes))
func (_TestAbort *TestAbortSession) GetAbortedWithMessage(message string) (AbortContext, error) {
	return _TestAbort.Contract.GetAbortedWithMessage(&_TestAbort.CallOpts, message)
}

// GetAbortedWithMessage is a free data retrieval call binding the contract method 0x9e59f463.
//
// Solidity: function getAbortedWithMessage(string message) view returns((bytes,address,uint256,bool,uint256,bytes))
func (_TestAbort *TestAbortCallerSession) GetAbortedWithMessage(message string) (AbortContext, error) {
	return _TestAbort.Contract.GetAbortedWithMessage(&_TestAbort.CallOpts, message)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_TestAbort *TestAbortTransactor) OnAbort(opts *bind.TransactOpts, abortContext AbortContext) (*types.Transaction, error) {
	return _TestAbort.contract.Transact(opts, "onAbort", abortContext)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_TestAbort *TestAbortSession) OnAbort(abortContext AbortContext) (*types.Transaction, error) {
	return _TestAbort.Contract.OnAbort(&_TestAbort.TransactOpts, abortContext)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_TestAbort *TestAbortTransactorSession) OnAbort(abortContext AbortContext) (*types.Transaction, error) {
	return _TestAbort.Contract.OnAbort(&_TestAbort.TransactOpts, abortContext)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestAbort *TestAbortTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TestAbort.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestAbort *TestAbortSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestAbort.Contract.Fallback(&_TestAbort.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestAbort *TestAbortTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestAbort.Contract.Fallback(&_TestAbort.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestAbort *TestAbortTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAbort.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestAbort *TestAbortSession) Receive() (*types.Transaction, error) {
	return _TestAbort.Contract.Receive(&_TestAbort.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestAbort *TestAbortTransactorSession) Receive() (*types.Transaction, error) {
	return _TestAbort.Contract.Receive(&_TestAbort.TransactOpts)
}
