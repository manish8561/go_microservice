// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

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

// TokensMetaData contains all meta data concerning the Tokens contract.
var TokensMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY_OWNER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY_STRATEGIES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMintedOwner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMintedStrategies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokensABI is the input ABI used to generate the binding from.
// Deprecated: Use TokensMetaData.ABI instead.
var TokensABI = TokensMetaData.ABI

// Tokens is an auto generated Go binding around an Ethereum contract.
type Tokens struct {
	TokensCaller     // Read-only binding to the contract
	TokensTransactor // Write-only binding to the contract
	TokensFilterer   // Log filterer for contract events
}

// TokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokensSession struct {
	Contract     *Tokens           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokensCallerSession struct {
	Contract *TokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokensTransactorSession struct {
	Contract     *TokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokensRaw struct {
	Contract *Tokens // Generic contract binding to access the raw methods on
}

// TokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokensCallerRaw struct {
	Contract *TokensCaller // Generic read-only contract binding to access the raw methods on
}

// TokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokensTransactorRaw struct {
	Contract *TokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokens creates a new instance of Tokens, bound to a specific deployed contract.
func NewTokens(address common.Address, backend bind.ContractBackend) (*Tokens, error) {
	contract, err := bindTokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokens{TokensCaller: TokensCaller{contract: contract}, TokensTransactor: TokensTransactor{contract: contract}, TokensFilterer: TokensFilterer{contract: contract}}, nil
}

// NewTokensCaller creates a new read-only instance of Tokens, bound to a specific deployed contract.
func NewTokensCaller(address common.Address, caller bind.ContractCaller) (*TokensCaller, error) {
	contract, err := bindTokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokensCaller{contract: contract}, nil
}

// NewTokensTransactor creates a new write-only instance of Tokens, bound to a specific deployed contract.
func NewTokensTransactor(address common.Address, transactor bind.ContractTransactor) (*TokensTransactor, error) {
	contract, err := bindTokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokensTransactor{contract: contract}, nil
}

// NewTokensFilterer creates a new log filterer instance of Tokens, bound to a specific deployed contract.
func NewTokensFilterer(address common.Address, filterer bind.ContractFilterer) (*TokensFilterer, error) {
	contract, err := bindTokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokensFilterer{contract: contract}, nil
}

// bindTokens binds a generic wrapper to an already deployed contract.
func bindTokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokens *TokensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokens.Contract.TokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokens *TokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokens.Contract.TokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokens *TokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokens.Contract.TokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokens *TokensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokens *TokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokens *TokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokens.Contract.contract.Transact(opts, method, params...)
}

// MAXSUPPLYOWNER is a free data retrieval call binding the contract method 0x51761747.
//
// Solidity: function MAX_SUPPLY_OWNER() view returns(uint256)
func (_Tokens *TokensCaller) MAXSUPPLYOWNER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokens.contract.Call(opts, &out, "MAX_SUPPLY_OWNER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLYOWNER is a free data retrieval call binding the contract method 0x51761747.
//
// Solidity: function MAX_SUPPLY_OWNER() view returns(uint256)
func (_Tokens *TokensSession) MAXSUPPLYOWNER() (*big.Int, error) {
	return _Tokens.Contract.MAXSUPPLYOWNER(&_Tokens.CallOpts)
}

// MAXSUPPLYOWNER is a free data retrieval call binding the contract method 0x51761747.
//
// Solidity: function MAX_SUPPLY_OWNER() view returns(uint256)
func (_Tokens *TokensCallerSession) MAXSUPPLYOWNER() (*big.Int, error) {
	return _Tokens.Contract.MAXSUPPLYOWNER(&_Tokens.CallOpts)
}

// MAXSUPPLYSTRATEGIES is a free data retrieval call binding the contract method 0x45a2d8fe.
//
// Solidity: function MAX_SUPPLY_STRATEGIES() view returns(uint256)
func (_Tokens *TokensCaller) MAXSUPPLYSTRATEGIES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokens.contract.Call(opts, &out, "MAX_SUPPLY_STRATEGIES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLYSTRATEGIES is a free data retrieval call binding the contract method 0x45a2d8fe.
//
// Solidity: function MAX_SUPPLY_STRATEGIES() view returns(uint256)
func (_Tokens *TokensSession) MAXSUPPLYSTRATEGIES() (*big.Int, error) {
	return _Tokens.Contract.MAXSUPPLYSTRATEGIES(&_Tokens.CallOpts)
}

// MAXSUPPLYSTRATEGIES is a free data retrieval call binding the contract method 0x45a2d8fe.
//
// Solidity: function MAX_SUPPLY_STRATEGIES() view returns(uint256)
func (_Tokens *TokensCallerSession) MAXSUPPLYSTRATEGIES() (*big.Int, error) {
	return _Tokens.Contract.MAXSUPPLYSTRATEGIES(&_Tokens.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokens *TokensCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokens.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokens *TokensSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Tokens.Contract.Allowance(&_Tokens.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokens *TokensCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Tokens.Contract.Allowance(&_Tokens.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokens *TokensCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokens.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokens *TokensSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Tokens.Contract.BalanceOf(&_Tokens.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokens *TokensCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Tokens.Contract.BalanceOf(&_Tokens.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tokens *TokensCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokens.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tokens *TokensSession) Decimals() (uint8, error) {
	return _Tokens.Contract.Decimals(&_Tokens.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tokens *TokensCallerSession) Decimals() (uint8, error) {
	return _Tokens.Contract.Decimals(&_Tokens.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokens *TokensCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokens.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokens *TokensSession) Name() (string, error) {
	return _Tokens.Contract.Name(&_Tokens.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokens *TokensCallerSession) Name() (string, error) {
	return _Tokens.Contract.Name(&_Tokens.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Tokens *TokensCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Tokens.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Tokens *TokensSession) Operators(arg0 common.Address) (bool, error) {
	return _Tokens.Contract.Operators(&_Tokens.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Tokens *TokensCallerSession) Operators(arg0 common.Address) (bool, error) {
	return _Tokens.Contract.Operators(&_Tokens.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokens *TokensCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokens.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokens *TokensSession) Owner() (common.Address, error) {
	return _Tokens.Contract.Owner(&_Tokens.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokens *TokensCallerSession) Owner() (common.Address, error) {
	return _Tokens.Contract.Owner(&_Tokens.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokens *TokensCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokens.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokens *TokensSession) Symbol() (string, error) {
	return _Tokens.Contract.Symbol(&_Tokens.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokens *TokensCallerSession) Symbol() (string, error) {
	return _Tokens.Contract.Symbol(&_Tokens.CallOpts)
}

// TokenMintedOwner is a free data retrieval call binding the contract method 0xe928d7c0.
//
// Solidity: function tokenMintedOwner() view returns(uint256)
func (_Tokens *TokensCaller) TokenMintedOwner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokens.contract.Call(opts, &out, "tokenMintedOwner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMintedOwner is a free data retrieval call binding the contract method 0xe928d7c0.
//
// Solidity: function tokenMintedOwner() view returns(uint256)
func (_Tokens *TokensSession) TokenMintedOwner() (*big.Int, error) {
	return _Tokens.Contract.TokenMintedOwner(&_Tokens.CallOpts)
}

// TokenMintedOwner is a free data retrieval call binding the contract method 0xe928d7c0.
//
// Solidity: function tokenMintedOwner() view returns(uint256)
func (_Tokens *TokensCallerSession) TokenMintedOwner() (*big.Int, error) {
	return _Tokens.Contract.TokenMintedOwner(&_Tokens.CallOpts)
}

// TokenMintedStrategies is a free data retrieval call binding the contract method 0x44703032.
//
// Solidity: function tokenMintedStrategies() view returns(uint256)
func (_Tokens *TokensCaller) TokenMintedStrategies(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokens.contract.Call(opts, &out, "tokenMintedStrategies")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMintedStrategies is a free data retrieval call binding the contract method 0x44703032.
//
// Solidity: function tokenMintedStrategies() view returns(uint256)
func (_Tokens *TokensSession) TokenMintedStrategies() (*big.Int, error) {
	return _Tokens.Contract.TokenMintedStrategies(&_Tokens.CallOpts)
}

// TokenMintedStrategies is a free data retrieval call binding the contract method 0x44703032.
//
// Solidity: function tokenMintedStrategies() view returns(uint256)
func (_Tokens *TokensCallerSession) TokenMintedStrategies() (*big.Int, error) {
	return _Tokens.Contract.TokenMintedStrategies(&_Tokens.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokens *TokensCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokens.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokens *TokensSession) TotalSupply() (*big.Int, error) {
	return _Tokens.Contract.TotalSupply(&_Tokens.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokens *TokensCallerSession) TotalSupply() (*big.Int, error) {
	return _Tokens.Contract.TotalSupply(&_Tokens.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tokens *TokensTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokens.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tokens *TokensSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokens.Contract.Approve(&_Tokens.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tokens *TokensTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokens.Contract.Approve(&_Tokens.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tokens *TokensTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tokens.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tokens *TokensSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tokens.Contract.DecreaseAllowance(&_Tokens.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tokens *TokensTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tokens.Contract.DecreaseAllowance(&_Tokens.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tokens *TokensTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tokens.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tokens *TokensSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tokens.Contract.IncreaseAllowance(&_Tokens.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tokens *TokensTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tokens.Contract.IncreaseAllowance(&_Tokens.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Tokens *TokensTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokens.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Tokens *TokensSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokens.Contract.Mint(&_Tokens.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Tokens *TokensTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokens.Contract.Mint(&_Tokens.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokens *TokensTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokens.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokens *TokensSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokens.Contract.RenounceOwnership(&_Tokens.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokens *TokensTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokens.Contract.RenounceOwnership(&_Tokens.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool value) returns()
func (_Tokens *TokensTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address, value bool) (*types.Transaction, error) {
	return _Tokens.contract.Transact(opts, "setOperator", _operator, value)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool value) returns()
func (_Tokens *TokensSession) SetOperator(_operator common.Address, value bool) (*types.Transaction, error) {
	return _Tokens.Contract.SetOperator(&_Tokens.TransactOpts, _operator, value)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool value) returns()
func (_Tokens *TokensTransactorSession) SetOperator(_operator common.Address, value bool) (*types.Transaction, error) {
	return _Tokens.Contract.SetOperator(&_Tokens.TransactOpts, _operator, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Tokens *TokensTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokens.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Tokens *TokensSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokens.Contract.Transfer(&_Tokens.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Tokens *TokensTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokens.Contract.Transfer(&_Tokens.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Tokens *TokensTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokens.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Tokens *TokensSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokens.Contract.TransferFrom(&_Tokens.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Tokens *TokensTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokens.Contract.TransferFrom(&_Tokens.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokens *TokensTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tokens.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokens *TokensSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokens.Contract.TransferOwnership(&_Tokens.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokens *TokensTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokens.Contract.TransferOwnership(&_Tokens.TransactOpts, newOwner)
}

// TokensApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tokens contract.
type TokensApprovalIterator struct {
	Event *TokensApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokensApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokensApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokensApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensApproval represents a Approval event raised by the Tokens contract.
type TokensApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokens *TokensFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokensApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tokens.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokensApprovalIterator{contract: _Tokens.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokens *TokensFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokensApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tokens.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensApproval)
				if err := _Tokens.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokens *TokensFilterer) ParseApproval(log types.Log) (*TokensApproval, error) {
	event := new(TokensApproval)
	if err := _Tokens.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokensOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tokens contract.
type TokensOwnershipTransferredIterator struct {
	Event *TokensOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokensOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokensOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokensOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensOwnershipTransferred represents a OwnershipTransferred event raised by the Tokens contract.
type TokensOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokens *TokensFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokensOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokens.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokensOwnershipTransferredIterator{contract: _Tokens.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokens *TokensFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokensOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokens.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensOwnershipTransferred)
				if err := _Tokens.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokens *TokensFilterer) ParseOwnershipTransferred(log types.Log) (*TokensOwnershipTransferred, error) {
	event := new(TokensOwnershipTransferred)
	if err := _Tokens.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokensTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Tokens contract.
type TokensTransferIterator struct {
	Event *TokensTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokensTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokensTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokensTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensTransfer represents a Transfer event raised by the Tokens contract.
type TokensTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokens *TokensFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokensTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tokens.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokensTransferIterator{contract: _Tokens.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokens *TokensFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokensTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tokens.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensTransfer)
				if err := _Tokens.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokens *TokensFilterer) ParseTransfer(log types.Log) (*TokensTransfer, error) {
	event := new(TokensTransfer)
	if err := _Tokens.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
