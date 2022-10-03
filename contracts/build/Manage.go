// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package build

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
)

// ManageMetaData contains all meta data concerning the Manage contract.
var ManageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"balanceToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405162000ee038038062000ee08339818101604052810190610034919061011f565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505061015f565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100ec826100c1565b9050919050565b6100fc816100e1565b811461010757600080fd5b50565b600081519050610119816100f3565b92915050565b60008060408385031215610136576101356100bc565b5b60006101448582860161010a565b92505060206101558582860161010a565b9150509250929050565b610d71806200016f6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063045990121461005157806347e7ef2414610081578063f3fef3a31461009d578063f5537ede146100b9575b600080fd5b61006b60048036038101906100669190610892565b6100d5565b60405161007891906108d8565b60405180910390f35b61009b6004803603810190610096919061091f565b61011e565b005b6100b760048036038101906100b2919061091f565b610385565b005b6100d360048036038101906100ce919061095f565b61066c565b005b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101a590610a0f565b60405180910390fd5b8060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231846040518263ffffffff1660e01b81526004016102089190610a3e565b602060405180830381865afa158015610225573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102499190610a6e565b101561028a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028190610ae7565b60405180910390fd5b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102d99190610b36565b9250508190555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd8330846040518463ffffffff1660e01b815260040161033d93929190610b6a565b6020604051808303816000875af115801561035c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103809190610bd9565b505050565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610415576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040c90610a0f565b60405180910390fd5b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610497576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048e90610c52565b60405180910390fd5b8060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016104f19190610a3e565b602060405180830381865afa15801561050e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105329190610a6e565b1015610573576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056a90610c52565b60405180910390fd5b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105c29190610c72565b9250508190555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401610624929190610ca6565b6020604051808303816000875af1158015610643573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106679190610bd9565b505050565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146106fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f390610a0f565b60405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561077e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077590610d1b565b60405180910390fd5b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107cd9190610b36565b9250508190555080600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108239190610c72565b92505081905550505050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061085f82610834565b9050919050565b61086f81610854565b811461087a57600080fd5b50565b60008135905061088c81610866565b92915050565b6000602082840312156108a8576108a761082f565b5b60006108b68482850161087d565b91505092915050565b6000819050919050565b6108d2816108bf565b82525050565b60006020820190506108ed60008301846108c9565b92915050565b6108fc816108bf565b811461090757600080fd5b50565b600081359050610919816108f3565b92915050565b600080604083850312156109365761093561082f565b5b60006109448582860161087d565b92505060206109558582860161090a565b9150509250929050565b6000806000606084860312156109785761097761082f565b5b60006109868682870161087d565b93505060206109978682870161087d565b92505060406109a88682870161090a565b9150509250925092565b600082825260208201905092915050565b7f4d616e6167653a20453031000000000000000000000000000000000000000000600082015250565b60006109f9600b836109b2565b9150610a04826109c3565b602082019050919050565b60006020820190508181036000830152610a28816109ec565b9050919050565b610a3881610854565b82525050565b6000602082019050610a536000830184610a2f565b92915050565b600081519050610a68816108f3565b92915050565b600060208284031215610a8457610a8361082f565b5b6000610a9284828501610a59565b91505092915050565b7f4d616e6167653a45303200000000000000000000000000000000000000000000600082015250565b6000610ad1600a836109b2565b9150610adc82610a9b565b602082019050919050565b60006020820190508181036000830152610b0081610ac4565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610b41826108bf565b9150610b4c836108bf565b9250828201905080821115610b6457610b63610b07565b5b92915050565b6000606082019050610b7f6000830186610a2f565b610b8c6020830185610a2f565b610b9960408301846108c9565b949350505050565b60008115159050919050565b610bb681610ba1565b8114610bc157600080fd5b50565b600081519050610bd381610bad565b92915050565b600060208284031215610bef57610bee61082f565b5b6000610bfd84828501610bc4565b91505092915050565b7f4d616e6167653a45303300000000000000000000000000000000000000000000600082015250565b6000610c3c600a836109b2565b9150610c4782610c06565b602082019050919050565b60006020820190508181036000830152610c6b81610c2f565b9050919050565b6000610c7d826108bf565b9150610c88836108bf565b9250828203905081811115610ca057610c9f610b07565b5b92915050565b6000604082019050610cbb6000830185610a2f565b610cc860208301846108c9565b9392505050565b7f4d616e6167653a45303100000000000000000000000000000000000000000000600082015250565b6000610d05600a836109b2565b9150610d1082610ccf565b602082019050919050565b60006020820190508181036000830152610d3481610cf8565b905091905056fea2646970667358221220e80ed2e5d9854d6593659251525c9553b5efd40fa3fb2b68105bb454b018df2764736f6c63430008100033",
}

// ManageABI is the input ABI used to generate the binding from.
// Deprecated: Use ManageMetaData.ABI instead.
var ManageABI = ManageMetaData.ABI

// ManageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ManageMetaData.Bin instead.
var ManageBin = ManageMetaData.Bin

// DeployManage deploys a new Ethereum contract, binding an instance of Manage to it.
func DeployManage(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _token common.Address) (common.Address, *types.Transaction, *Manage, error) {
	parsed, err := ManageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ManageBin), backend, _admin, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Manage{ManageCaller: ManageCaller{contract: contract}, ManageTransactor: ManageTransactor{contract: contract}, ManageFilterer: ManageFilterer{contract: contract}}, nil
}

// Manage is an auto generated Go binding around an Ethereum contract.
type Manage struct {
	ManageCaller     // Read-only binding to the contract
	ManageTransactor // Write-only binding to the contract
	ManageFilterer   // Log filterer for contract events
}

// ManageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManageSession struct {
	Contract     *Manage           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManageCallerSession struct {
	Contract *ManageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ManageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManageTransactorSession struct {
	Contract     *ManageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManageRaw struct {
	Contract *Manage // Generic contract binding to access the raw methods on
}

// ManageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManageCallerRaw struct {
	Contract *ManageCaller // Generic read-only contract binding to access the raw methods on
}

// ManageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManageTransactorRaw struct {
	Contract *ManageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManage creates a new instance of Manage, bound to a specific deployed contract.
func NewManage(address common.Address, backend bind.ContractBackend) (*Manage, error) {
	contract, err := bindManage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Manage{ManageCaller: ManageCaller{contract: contract}, ManageTransactor: ManageTransactor{contract: contract}, ManageFilterer: ManageFilterer{contract: contract}}, nil
}

// NewManageCaller creates a new read-only instance of Manage, bound to a specific deployed contract.
func NewManageCaller(address common.Address, caller bind.ContractCaller) (*ManageCaller, error) {
	contract, err := bindManage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManageCaller{contract: contract}, nil
}

// NewManageTransactor creates a new write-only instance of Manage, bound to a specific deployed contract.
func NewManageTransactor(address common.Address, transactor bind.ContractTransactor) (*ManageTransactor, error) {
	contract, err := bindManage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManageTransactor{contract: contract}, nil
}

// NewManageFilterer creates a new log filterer instance of Manage, bound to a specific deployed contract.
func NewManageFilterer(address common.Address, filterer bind.ContractFilterer) (*ManageFilterer, error) {
	contract, err := bindManage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManageFilterer{contract: contract}, nil
}

// bindManage binds a generic wrapper to an already deployed contract.
func bindManage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manage *ManageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manage.Contract.ManageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manage *ManageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manage.Contract.ManageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manage *ManageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manage.Contract.ManageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manage *ManageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manage *ManageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manage *ManageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manage.Contract.contract.Transact(opts, method, params...)
}

// BalanceToken is a free data retrieval call binding the contract method 0x04599012.
//
// Solidity: function balanceToken(address from) view returns(uint256 value)
func (_Manage *ManageCaller) BalanceToken(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Manage.contract.Call(opts, &out, "balanceToken", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceToken is a free data retrieval call binding the contract method 0x04599012.
//
// Solidity: function balanceToken(address from) view returns(uint256 value)
func (_Manage *ManageSession) BalanceToken(from common.Address) (*big.Int, error) {
	return _Manage.Contract.BalanceToken(&_Manage.CallOpts, from)
}

// BalanceToken is a free data retrieval call binding the contract method 0x04599012.
//
// Solidity: function balanceToken(address from) view returns(uint256 value)
func (_Manage *ManageCallerSession) BalanceToken(from common.Address) (*big.Int, error) {
	return _Manage.Contract.BalanceToken(&_Manage.CallOpts, from)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address from, uint256 value) returns()
func (_Manage *ManageTransactor) Deposit(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Manage.contract.Transact(opts, "deposit", from, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address from, uint256 value) returns()
func (_Manage *ManageSession) Deposit(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Manage.Contract.Deposit(&_Manage.TransactOpts, from, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address from, uint256 value) returns()
func (_Manage *ManageTransactorSession) Deposit(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Manage.Contract.Deposit(&_Manage.TransactOpts, from, value)
}

// TransferToken is a paid mutator transaction binding the contract method 0xf5537ede.
//
// Solidity: function transferToken(address from, address to, uint256 value) returns()
func (_Manage *ManageTransactor) TransferToken(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Manage.contract.Transact(opts, "transferToken", from, to, value)
}

// TransferToken is a paid mutator transaction binding the contract method 0xf5537ede.
//
// Solidity: function transferToken(address from, address to, uint256 value) returns()
func (_Manage *ManageSession) TransferToken(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Manage.Contract.TransferToken(&_Manage.TransactOpts, from, to, value)
}

// TransferToken is a paid mutator transaction binding the contract method 0xf5537ede.
//
// Solidity: function transferToken(address from, address to, uint256 value) returns()
func (_Manage *ManageTransactorSession) TransferToken(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Manage.Contract.TransferToken(&_Manage.TransactOpts, from, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address from, uint256 value) returns()
func (_Manage *ManageTransactor) Withdraw(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Manage.contract.Transact(opts, "withdraw", from, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address from, uint256 value) returns()
func (_Manage *ManageSession) Withdraw(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Manage.Contract.Withdraw(&_Manage.TransactOpts, from, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address from, uint256 value) returns()
func (_Manage *ManageTransactorSession) Withdraw(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Manage.Contract.Withdraw(&_Manage.TransactOpts, from, value)
}
