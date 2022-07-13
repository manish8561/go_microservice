// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transactions

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

// TransactionsMetaData contains all meta data concerning the Transactions contract.
var TransactionsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collectible\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dragonsliar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quick\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingDualRewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_autoCompoundTokenPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_ops\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalDeposits\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalSupply\",\"type\":\"uint256\"}],\"name\":\"Reinvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"UpdateAdminFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"UpdateReinvestReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newValue\",\"type\":\"bool\"}],\"name\":\"UpdateRequireReinvestBeforeDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_FEE_BIPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BONUS_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_AUTOCOMPOUND_TOKEN\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REINVEST_REWARD_BIPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REQUIRE_REINVEST_BEFORE_DEPOSIT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acStakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"autoCompoundToken\",\"outputs\":[{\"internalType\":\"contractIERC20AutoCompound\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"autoCompoundTokenPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"autoCompoundTokenPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectible\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositToken\",\"outputs\":[{\"internalType\":\"contractIPair\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dragonsliar\",\"outputs\":[{\"internalType\":\"contractIDragonLiar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"}],\"name\":\"dualTokenDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"dualWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimateReinvestReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getDepositTokensForShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getSharesForDepositTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintAutocompoundTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ops\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quick\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverNativeAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinvestOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"revokeAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setAllowances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"}],\"name\":\"singleTokenDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"singleWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingDualRewards\",\"outputs\":[{\"internalType\":\"contractIStakingDualRewards\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"updateAdminFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_autoCompoundTokenPerBlock\",\"type\":\"uint256\"}],\"name\":\"updateAutoCompoundTokenPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"multiplierNumber\",\"type\":\"uint256\"}],\"name\":\"updateMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"updateReinvestReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRequireReinvestBeforeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rewardDebt\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TransactionsABI is the input ABI used to generate the binding from.
// Deprecated: Use TransactionsMetaData.ABI instead.
var TransactionsABI = TransactionsMetaData.ABI

// Transactions is an auto generated Go binding around an Ethereum contract.
type Transactions struct {
	TransactionsCaller     // Read-only binding to the contract
	TransactionsTransactor // Write-only binding to the contract
	TransactionsFilterer   // Log filterer for contract events
}

// TransactionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransactionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransactionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransactionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransactionsSession struct {
	Contract     *Transactions     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransactionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransactionsCallerSession struct {
	Contract *TransactionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TransactionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransactionsTransactorSession struct {
	Contract     *TransactionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TransactionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransactionsRaw struct {
	Contract *Transactions // Generic contract binding to access the raw methods on
}

// TransactionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransactionsCallerRaw struct {
	Contract *TransactionsCaller // Generic read-only contract binding to access the raw methods on
}

// TransactionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactionsTransactorRaw struct {
	Contract *TransactionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransactions creates a new instance of Transactions, bound to a specific deployed contract.
func NewTransactions(address common.Address, backend bind.ContractBackend) (*Transactions, error) {
	contract, err := bindTransactions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transactions{TransactionsCaller: TransactionsCaller{contract: contract}, TransactionsTransactor: TransactionsTransactor{contract: contract}, TransactionsFilterer: TransactionsFilterer{contract: contract}}, nil
}

// NewTransactionsCaller creates a new read-only instance of Transactions, bound to a specific deployed contract.
func NewTransactionsCaller(address common.Address, caller bind.ContractCaller) (*TransactionsCaller, error) {
	contract, err := bindTransactions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionsCaller{contract: contract}, nil
}

// NewTransactionsTransactor creates a new write-only instance of Transactions, bound to a specific deployed contract.
func NewTransactionsTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactionsTransactor, error) {
	contract, err := bindTransactions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionsTransactor{contract: contract}, nil
}

// NewTransactionsFilterer creates a new log filterer instance of Transactions, bound to a specific deployed contract.
func NewTransactionsFilterer(address common.Address, filterer bind.ContractFilterer) (*TransactionsFilterer, error) {
	contract, err := bindTransactions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransactionsFilterer{contract: contract}, nil
}

// bindTransactions binds a generic wrapper to an already deployed contract.
func bindTransactions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transactions *TransactionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transactions.Contract.TransactionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transactions *TransactionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactions.Contract.TransactionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transactions *TransactionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transactions.Contract.TransactionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transactions *TransactionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transactions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transactions *TransactionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transactions *TransactionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transactions.Contract.contract.Transact(opts, method, params...)
}

// ADMINFEEBIPS is a free data retrieval call binding the contract method 0x07677111.
//
// Solidity: function ADMIN_FEE_BIPS() view returns(uint256)
func (_Transactions *TransactionsCaller) ADMINFEEBIPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "ADMIN_FEE_BIPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ADMINFEEBIPS is a free data retrieval call binding the contract method 0x07677111.
//
// Solidity: function ADMIN_FEE_BIPS() view returns(uint256)
func (_Transactions *TransactionsSession) ADMINFEEBIPS() (*big.Int, error) {
	return _Transactions.Contract.ADMINFEEBIPS(&_Transactions.CallOpts)
}

// ADMINFEEBIPS is a free data retrieval call binding the contract method 0x07677111.
//
// Solidity: function ADMIN_FEE_BIPS() view returns(uint256)
func (_Transactions *TransactionsCallerSession) ADMINFEEBIPS() (*big.Int, error) {
	return _Transactions.Contract.ADMINFEEBIPS(&_Transactions.CallOpts)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_Transactions *TransactionsCaller) BONUSMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "BONUS_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_Transactions *TransactionsSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _Transactions.Contract.BONUSMULTIPLIER(&_Transactions.CallOpts)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_Transactions *TransactionsCallerSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _Transactions.Contract.BONUSMULTIPLIER(&_Transactions.CallOpts)
}

// MINTAUTOCOMPOUNDTOKEN is a free data retrieval call binding the contract method 0xab4787a1.
//
// Solidity: function MINT_AUTOCOMPOUND_TOKEN() view returns(bool)
func (_Transactions *TransactionsCaller) MINTAUTOCOMPOUNDTOKEN(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "MINT_AUTOCOMPOUND_TOKEN")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MINTAUTOCOMPOUNDTOKEN is a free data retrieval call binding the contract method 0xab4787a1.
//
// Solidity: function MINT_AUTOCOMPOUND_TOKEN() view returns(bool)
func (_Transactions *TransactionsSession) MINTAUTOCOMPOUNDTOKEN() (bool, error) {
	return _Transactions.Contract.MINTAUTOCOMPOUNDTOKEN(&_Transactions.CallOpts)
}

// MINTAUTOCOMPOUNDTOKEN is a free data retrieval call binding the contract method 0xab4787a1.
//
// Solidity: function MINT_AUTOCOMPOUND_TOKEN() view returns(bool)
func (_Transactions *TransactionsCallerSession) MINTAUTOCOMPOUNDTOKEN() (bool, error) {
	return _Transactions.Contract.MINTAUTOCOMPOUNDTOKEN(&_Transactions.CallOpts)
}

// REINVESTREWARDBIPS is a free data retrieval call binding the contract method 0x8aff733d.
//
// Solidity: function REINVEST_REWARD_BIPS() view returns(uint256)
func (_Transactions *TransactionsCaller) REINVESTREWARDBIPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "REINVEST_REWARD_BIPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REINVESTREWARDBIPS is a free data retrieval call binding the contract method 0x8aff733d.
//
// Solidity: function REINVEST_REWARD_BIPS() view returns(uint256)
func (_Transactions *TransactionsSession) REINVESTREWARDBIPS() (*big.Int, error) {
	return _Transactions.Contract.REINVESTREWARDBIPS(&_Transactions.CallOpts)
}

// REINVESTREWARDBIPS is a free data retrieval call binding the contract method 0x8aff733d.
//
// Solidity: function REINVEST_REWARD_BIPS() view returns(uint256)
func (_Transactions *TransactionsCallerSession) REINVESTREWARDBIPS() (*big.Int, error) {
	return _Transactions.Contract.REINVESTREWARDBIPS(&_Transactions.CallOpts)
}

// REQUIREREINVESTBEFOREDEPOSIT is a free data retrieval call binding the contract method 0x13317314.
//
// Solidity: function REQUIRE_REINVEST_BEFORE_DEPOSIT() view returns(bool)
func (_Transactions *TransactionsCaller) REQUIREREINVESTBEFOREDEPOSIT(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "REQUIRE_REINVEST_BEFORE_DEPOSIT")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// REQUIREREINVESTBEFOREDEPOSIT is a free data retrieval call binding the contract method 0x13317314.
//
// Solidity: function REQUIRE_REINVEST_BEFORE_DEPOSIT() view returns(bool)
func (_Transactions *TransactionsSession) REQUIREREINVESTBEFOREDEPOSIT() (bool, error) {
	return _Transactions.Contract.REQUIREREINVESTBEFOREDEPOSIT(&_Transactions.CallOpts)
}

// REQUIREREINVESTBEFOREDEPOSIT is a free data retrieval call binding the contract method 0x13317314.
//
// Solidity: function REQUIRE_REINVEST_BEFORE_DEPOSIT() view returns(bool)
func (_Transactions *TransactionsCallerSession) REQUIREREINVESTBEFOREDEPOSIT() (bool, error) {
	return _Transactions.Contract.REQUIREREINVESTBEFOREDEPOSIT(&_Transactions.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Transactions *TransactionsCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Transactions *TransactionsSession) WETH() (common.Address, error) {
	return _Transactions.Contract.WETH(&_Transactions.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Transactions *TransactionsCallerSession) WETH() (common.Address, error) {
	return _Transactions.Contract.WETH(&_Transactions.CallOpts)
}

// AcStakingContract is a free data retrieval call binding the contract method 0x7c54b0bb.
//
// Solidity: function acStakingContract() view returns(address)
func (_Transactions *TransactionsCaller) AcStakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "acStakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AcStakingContract is a free data retrieval call binding the contract method 0x7c54b0bb.
//
// Solidity: function acStakingContract() view returns(address)
func (_Transactions *TransactionsSession) AcStakingContract() (common.Address, error) {
	return _Transactions.Contract.AcStakingContract(&_Transactions.CallOpts)
}

// AcStakingContract is a free data retrieval call binding the contract method 0x7c54b0bb.
//
// Solidity: function acStakingContract() view returns(address)
func (_Transactions *TransactionsCallerSession) AcStakingContract() (common.Address, error) {
	return _Transactions.Contract.AcStakingContract(&_Transactions.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address account, address spender) view returns(uint256)
func (_Transactions *TransactionsCaller) Allowance(opts *bind.CallOpts, account common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "allowance", account, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address account, address spender) view returns(uint256)
func (_Transactions *TransactionsSession) Allowance(account common.Address, spender common.Address) (*big.Int, error) {
	return _Transactions.Contract.Allowance(&_Transactions.CallOpts, account, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address account, address spender) view returns(uint256)
func (_Transactions *TransactionsCallerSession) Allowance(account common.Address, spender common.Address) (*big.Int, error) {
	return _Transactions.Contract.Allowance(&_Transactions.CallOpts, account, spender)
}

// AutoCompoundToken is a free data retrieval call binding the contract method 0xe51d3693.
//
// Solidity: function autoCompoundToken() view returns(address)
func (_Transactions *TransactionsCaller) AutoCompoundToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "autoCompoundToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AutoCompoundToken is a free data retrieval call binding the contract method 0xe51d3693.
//
// Solidity: function autoCompoundToken() view returns(address)
func (_Transactions *TransactionsSession) AutoCompoundToken() (common.Address, error) {
	return _Transactions.Contract.AutoCompoundToken(&_Transactions.CallOpts)
}

// AutoCompoundToken is a free data retrieval call binding the contract method 0xe51d3693.
//
// Solidity: function autoCompoundToken() view returns(address)
func (_Transactions *TransactionsCallerSession) AutoCompoundToken() (common.Address, error) {
	return _Transactions.Contract.AutoCompoundToken(&_Transactions.CallOpts)
}

// AutoCompoundTokenPerBlock is a free data retrieval call binding the contract method 0xb5e1e2fa.
//
// Solidity: function autoCompoundTokenPerBlock() view returns(uint256)
func (_Transactions *TransactionsCaller) AutoCompoundTokenPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "autoCompoundTokenPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AutoCompoundTokenPerBlock is a free data retrieval call binding the contract method 0xb5e1e2fa.
//
// Solidity: function autoCompoundTokenPerBlock() view returns(uint256)
func (_Transactions *TransactionsSession) AutoCompoundTokenPerBlock() (*big.Int, error) {
	return _Transactions.Contract.AutoCompoundTokenPerBlock(&_Transactions.CallOpts)
}

// AutoCompoundTokenPerBlock is a free data retrieval call binding the contract method 0xb5e1e2fa.
//
// Solidity: function autoCompoundTokenPerBlock() view returns(uint256)
func (_Transactions *TransactionsCallerSession) AutoCompoundTokenPerBlock() (*big.Int, error) {
	return _Transactions.Contract.AutoCompoundTokenPerBlock(&_Transactions.CallOpts)
}

// AutoCompoundTokenPerShare is a free data retrieval call binding the contract method 0xc5ebc26e.
//
// Solidity: function autoCompoundTokenPerShare() view returns(uint256)
func (_Transactions *TransactionsCaller) AutoCompoundTokenPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "autoCompoundTokenPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AutoCompoundTokenPerShare is a free data retrieval call binding the contract method 0xc5ebc26e.
//
// Solidity: function autoCompoundTokenPerShare() view returns(uint256)
func (_Transactions *TransactionsSession) AutoCompoundTokenPerShare() (*big.Int, error) {
	return _Transactions.Contract.AutoCompoundTokenPerShare(&_Transactions.CallOpts)
}

// AutoCompoundTokenPerShare is a free data retrieval call binding the contract method 0xc5ebc26e.
//
// Solidity: function autoCompoundTokenPerShare() view returns(uint256)
func (_Transactions *TransactionsCallerSession) AutoCompoundTokenPerShare() (*big.Int, error) {
	return _Transactions.Contract.AutoCompoundTokenPerShare(&_Transactions.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Transactions *TransactionsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Transactions *TransactionsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Transactions.Contract.BalanceOf(&_Transactions.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Transactions *TransactionsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Transactions.Contract.BalanceOf(&_Transactions.CallOpts, account)
}

// Collectible is a free data retrieval call binding the contract method 0xea05a7d0.
//
// Solidity: function collectible() view returns(address)
func (_Transactions *TransactionsCaller) Collectible(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "collectible")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collectible is a free data retrieval call binding the contract method 0xea05a7d0.
//
// Solidity: function collectible() view returns(address)
func (_Transactions *TransactionsSession) Collectible() (common.Address, error) {
	return _Transactions.Contract.Collectible(&_Transactions.CallOpts)
}

// Collectible is a free data retrieval call binding the contract method 0xea05a7d0.
//
// Solidity: function collectible() view returns(address)
func (_Transactions *TransactionsCallerSession) Collectible() (common.Address, error) {
	return _Transactions.Contract.Collectible(&_Transactions.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Transactions *TransactionsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Transactions *TransactionsSession) Decimals() (uint8, error) {
	return _Transactions.Contract.Decimals(&_Transactions.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Transactions *TransactionsCallerSession) Decimals() (uint8, error) {
	return _Transactions.Contract.Decimals(&_Transactions.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_Transactions *TransactionsCaller) DepositToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "depositToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_Transactions *TransactionsSession) DepositToken() (common.Address, error) {
	return _Transactions.Contract.DepositToken(&_Transactions.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_Transactions *TransactionsCallerSession) DepositToken() (common.Address, error) {
	return _Transactions.Contract.DepositToken(&_Transactions.CallOpts)
}

// Dragonsliar is a free data retrieval call binding the contract method 0x29453a48.
//
// Solidity: function dragonsliar() view returns(address)
func (_Transactions *TransactionsCaller) Dragonsliar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "dragonsliar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dragonsliar is a free data retrieval call binding the contract method 0x29453a48.
//
// Solidity: function dragonsliar() view returns(address)
func (_Transactions *TransactionsSession) Dragonsliar() (common.Address, error) {
	return _Transactions.Contract.Dragonsliar(&_Transactions.CallOpts)
}

// Dragonsliar is a free data retrieval call binding the contract method 0x29453a48.
//
// Solidity: function dragonsliar() view returns(address)
func (_Transactions *TransactionsCallerSession) Dragonsliar() (common.Address, error) {
	return _Transactions.Contract.Dragonsliar(&_Transactions.CallOpts)
}

// EstimateReinvestReward is a free data retrieval call binding the contract method 0xb9e57b80.
//
// Solidity: function estimateReinvestReward() view returns(uint256)
func (_Transactions *TransactionsCaller) EstimateReinvestReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "estimateReinvestReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateReinvestReward is a free data retrieval call binding the contract method 0xb9e57b80.
//
// Solidity: function estimateReinvestReward() view returns(uint256)
func (_Transactions *TransactionsSession) EstimateReinvestReward() (*big.Int, error) {
	return _Transactions.Contract.EstimateReinvestReward(&_Transactions.CallOpts)
}

// EstimateReinvestReward is a free data retrieval call binding the contract method 0xb9e57b80.
//
// Solidity: function estimateReinvestReward() view returns(uint256)
func (_Transactions *TransactionsCallerSession) EstimateReinvestReward() (*big.Int, error) {
	return _Transactions.Contract.EstimateReinvestReward(&_Transactions.CallOpts)
}

// GetDepositTokensForShares is a free data retrieval call binding the contract method 0xeab89a5a.
//
// Solidity: function getDepositTokensForShares(uint256 amount) view returns(uint256)
func (_Transactions *TransactionsCaller) GetDepositTokensForShares(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "getDepositTokensForShares", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositTokensForShares is a free data retrieval call binding the contract method 0xeab89a5a.
//
// Solidity: function getDepositTokensForShares(uint256 amount) view returns(uint256)
func (_Transactions *TransactionsSession) GetDepositTokensForShares(amount *big.Int) (*big.Int, error) {
	return _Transactions.Contract.GetDepositTokensForShares(&_Transactions.CallOpts, amount)
}

// GetDepositTokensForShares is a free data retrieval call binding the contract method 0xeab89a5a.
//
// Solidity: function getDepositTokensForShares(uint256 amount) view returns(uint256)
func (_Transactions *TransactionsCallerSession) GetDepositTokensForShares(amount *big.Int) (*big.Int, error) {
	return _Transactions.Contract.GetDepositTokensForShares(&_Transactions.CallOpts, amount)
}

// GetSharesForDepositTokens is a free data retrieval call binding the contract method 0xdd8ce4d6.
//
// Solidity: function getSharesForDepositTokens(uint256 amount) view returns(uint256)
func (_Transactions *TransactionsCaller) GetSharesForDepositTokens(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "getSharesForDepositTokens", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesForDepositTokens is a free data retrieval call binding the contract method 0xdd8ce4d6.
//
// Solidity: function getSharesForDepositTokens(uint256 amount) view returns(uint256)
func (_Transactions *TransactionsSession) GetSharesForDepositTokens(amount *big.Int) (*big.Int, error) {
	return _Transactions.Contract.GetSharesForDepositTokens(&_Transactions.CallOpts, amount)
}

// GetSharesForDepositTokens is a free data retrieval call binding the contract method 0xdd8ce4d6.
//
// Solidity: function getSharesForDepositTokens(uint256 amount) view returns(uint256)
func (_Transactions *TransactionsCallerSession) GetSharesForDepositTokens(amount *big.Int) (*big.Int, error) {
	return _Transactions.Contract.GetSharesForDepositTokens(&_Transactions.CallOpts, amount)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Transactions *TransactionsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Transactions *TransactionsSession) Name() (string, error) {
	return _Transactions.Contract.Name(&_Transactions.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Transactions *TransactionsCallerSession) Name() (string, error) {
	return _Transactions.Contract.Name(&_Transactions.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Transactions *TransactionsCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Transactions *TransactionsSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Transactions.Contract.Nonces(&_Transactions.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Transactions *TransactionsCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Transactions.Contract.Nonces(&_Transactions.CallOpts, arg0)
}

// Ops is a free data retrieval call binding the contract method 0xe70abe92.
//
// Solidity: function ops() view returns(address)
func (_Transactions *TransactionsCaller) Ops(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "ops")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ops is a free data retrieval call binding the contract method 0xe70abe92.
//
// Solidity: function ops() view returns(address)
func (_Transactions *TransactionsSession) Ops() (common.Address, error) {
	return _Transactions.Contract.Ops(&_Transactions.CallOpts)
}

// Ops is a free data retrieval call binding the contract method 0xe70abe92.
//
// Solidity: function ops() view returns(address)
func (_Transactions *TransactionsCallerSession) Ops() (common.Address, error) {
	return _Transactions.Contract.Ops(&_Transactions.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Transactions *TransactionsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Transactions *TransactionsSession) Owner() (common.Address, error) {
	return _Transactions.Contract.Owner(&_Transactions.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Transactions *TransactionsCallerSession) Owner() (common.Address, error) {
	return _Transactions.Contract.Owner(&_Transactions.CallOpts)
}

// Quick is a free data retrieval call binding the contract method 0xfdd3a879.
//
// Solidity: function quick() view returns(address)
func (_Transactions *TransactionsCaller) Quick(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "quick")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Quick is a free data retrieval call binding the contract method 0xfdd3a879.
//
// Solidity: function quick() view returns(address)
func (_Transactions *TransactionsSession) Quick() (common.Address, error) {
	return _Transactions.Contract.Quick(&_Transactions.CallOpts)
}

// Quick is a free data retrieval call binding the contract method 0xfdd3a879.
//
// Solidity: function quick() view returns(address)
func (_Transactions *TransactionsCallerSession) Quick() (common.Address, error) {
	return _Transactions.Contract.Quick(&_Transactions.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Transactions *TransactionsCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Transactions *TransactionsSession) Router() (common.Address, error) {
	return _Transactions.Contract.Router(&_Transactions.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Transactions *TransactionsCallerSession) Router() (common.Address, error) {
	return _Transactions.Contract.Router(&_Transactions.CallOpts)
}

// StakingDualRewards is a free data retrieval call binding the contract method 0xa135edec.
//
// Solidity: function stakingDualRewards() view returns(address)
func (_Transactions *TransactionsCaller) StakingDualRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "stakingDualRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingDualRewards is a free data retrieval call binding the contract method 0xa135edec.
//
// Solidity: function stakingDualRewards() view returns(address)
func (_Transactions *TransactionsSession) StakingDualRewards() (common.Address, error) {
	return _Transactions.Contract.StakingDualRewards(&_Transactions.CallOpts)
}

// StakingDualRewards is a free data retrieval call binding the contract method 0xa135edec.
//
// Solidity: function stakingDualRewards() view returns(address)
func (_Transactions *TransactionsCallerSession) StakingDualRewards() (common.Address, error) {
	return _Transactions.Contract.StakingDualRewards(&_Transactions.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Transactions *TransactionsCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Transactions *TransactionsSession) StartBlock() (*big.Int, error) {
	return _Transactions.Contract.StartBlock(&_Transactions.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Transactions *TransactionsCallerSession) StartBlock() (*big.Int, error) {
	return _Transactions.Contract.StartBlock(&_Transactions.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Transactions *TransactionsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Transactions *TransactionsSession) Symbol() (string, error) {
	return _Transactions.Contract.Symbol(&_Transactions.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Transactions *TransactionsCallerSession) Symbol() (string, error) {
	return _Transactions.Contract.Symbol(&_Transactions.CallOpts)
}

// TotalDeposits is a free data retrieval call binding the contract method 0x7d882097.
//
// Solidity: function totalDeposits() view returns(uint256)
func (_Transactions *TransactionsCaller) TotalDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "totalDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeposits is a free data retrieval call binding the contract method 0x7d882097.
//
// Solidity: function totalDeposits() view returns(uint256)
func (_Transactions *TransactionsSession) TotalDeposits() (*big.Int, error) {
	return _Transactions.Contract.TotalDeposits(&_Transactions.CallOpts)
}

// TotalDeposits is a free data retrieval call binding the contract method 0x7d882097.
//
// Solidity: function totalDeposits() view returns(uint256)
func (_Transactions *TransactionsCallerSession) TotalDeposits() (*big.Int, error) {
	return _Transactions.Contract.TotalDeposits(&_Transactions.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Transactions *TransactionsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Transactions *TransactionsSession) TotalSupply() (*big.Int, error) {
	return _Transactions.Contract.TotalSupply(&_Transactions.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Transactions *TransactionsCallerSession) TotalSupply() (*big.Int, error) {
	return _Transactions.Contract.TotalSupply(&_Transactions.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transactions *TransactionsCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transactions *TransactionsSession) Treasury() (common.Address, error) {
	return _Transactions.Contract.Treasury(&_Transactions.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transactions *TransactionsCallerSession) Treasury() (common.Address, error) {
	return _Transactions.Contract.Treasury(&_Transactions.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint128 amount, uint128 rewardDebt)
func (_Transactions *TransactionsCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Transactions.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint128 amount, uint128 rewardDebt)
func (_Transactions *TransactionsSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Transactions.Contract.UserInfo(&_Transactions.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint128 amount, uint128 rewardDebt)
func (_Transactions *TransactionsCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Transactions.Contract.UserInfo(&_Transactions.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Transactions *TransactionsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Transactions *TransactionsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.Approve(&_Transactions.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Transactions *TransactionsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.Approve(&_Transactions.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Transactions *TransactionsTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Transactions *TransactionsSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.Deposit(&_Transactions.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Transactions *TransactionsTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.Deposit(&_Transactions.TransactOpts, amount)
}

// DualTokenDeposit is a paid mutator transaction binding the contract method 0x7161bff1.
//
// Solidity: function dualTokenDeposit(uint256 amount0, address _token0, uint256 amount1, address _token1, uint256 slippage) returns()
func (_Transactions *TransactionsTransactor) DualTokenDeposit(opts *bind.TransactOpts, amount0 *big.Int, _token0 common.Address, amount1 *big.Int, _token1 common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "dualTokenDeposit", amount0, _token0, amount1, _token1, slippage)
}

// DualTokenDeposit is a paid mutator transaction binding the contract method 0x7161bff1.
//
// Solidity: function dualTokenDeposit(uint256 amount0, address _token0, uint256 amount1, address _token1, uint256 slippage) returns()
func (_Transactions *TransactionsSession) DualTokenDeposit(amount0 *big.Int, _token0 common.Address, amount1 *big.Int, _token1 common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.DualTokenDeposit(&_Transactions.TransactOpts, amount0, _token0, amount1, _token1, slippage)
}

// DualTokenDeposit is a paid mutator transaction binding the contract method 0x7161bff1.
//
// Solidity: function dualTokenDeposit(uint256 amount0, address _token0, uint256 amount1, address _token1, uint256 slippage) returns()
func (_Transactions *TransactionsTransactorSession) DualTokenDeposit(amount0 *big.Int, _token0 common.Address, amount1 *big.Int, _token1 common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.DualTokenDeposit(&_Transactions.TransactOpts, amount0, _token0, amount1, _token1, slippage)
}

// DualWithdraw is a paid mutator transaction binding the contract method 0x1bf5e6c9.
//
// Solidity: function dualWithdraw(uint256 amount) returns()
func (_Transactions *TransactionsTransactor) DualWithdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "dualWithdraw", amount)
}

// DualWithdraw is a paid mutator transaction binding the contract method 0x1bf5e6c9.
//
// Solidity: function dualWithdraw(uint256 amount) returns()
func (_Transactions *TransactionsSession) DualWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.DualWithdraw(&_Transactions.TransactOpts, amount)
}

// DualWithdraw is a paid mutator transaction binding the contract method 0x1bf5e6c9.
//
// Solidity: function dualWithdraw(uint256 amount) returns()
func (_Transactions *TransactionsTransactorSession) DualWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.DualWithdraw(&_Transactions.TransactOpts, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Transactions *TransactionsTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Transactions *TransactionsSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Transactions.Contract.EmergencyWithdraw(&_Transactions.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Transactions *TransactionsTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Transactions.Contract.EmergencyWithdraw(&_Transactions.TransactOpts)
}

// MintAutocompoundTokens is a paid mutator transaction binding the contract method 0xdc65b760.
//
// Solidity: function mintAutocompoundTokens() returns()
func (_Transactions *TransactionsTransactor) MintAutocompoundTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "mintAutocompoundTokens")
}

// MintAutocompoundTokens is a paid mutator transaction binding the contract method 0xdc65b760.
//
// Solidity: function mintAutocompoundTokens() returns()
func (_Transactions *TransactionsSession) MintAutocompoundTokens() (*types.Transaction, error) {
	return _Transactions.Contract.MintAutocompoundTokens(&_Transactions.TransactOpts)
}

// MintAutocompoundTokens is a paid mutator transaction binding the contract method 0xdc65b760.
//
// Solidity: function mintAutocompoundTokens() returns()
func (_Transactions *TransactionsTransactorSession) MintAutocompoundTokens() (*types.Transaction, error) {
	return _Transactions.Contract.MintAutocompoundTokens(&_Transactions.TransactOpts)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_Transactions *TransactionsTransactor) RecoverERC20(opts *bind.TransactOpts, tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "recoverERC20", tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_Transactions *TransactionsSession) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.RecoverERC20(&_Transactions.TransactOpts, tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_Transactions *TransactionsTransactorSession) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.RecoverERC20(&_Transactions.TransactOpts, tokenAddress, tokenAmount)
}

// RecoverNativeAsset is a paid mutator transaction binding the contract method 0x4d00c033.
//
// Solidity: function recoverNativeAsset(uint256 amount) returns()
func (_Transactions *TransactionsTransactor) RecoverNativeAsset(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "recoverNativeAsset", amount)
}

// RecoverNativeAsset is a paid mutator transaction binding the contract method 0x4d00c033.
//
// Solidity: function recoverNativeAsset(uint256 amount) returns()
func (_Transactions *TransactionsSession) RecoverNativeAsset(amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.RecoverNativeAsset(&_Transactions.TransactOpts, amount)
}

// RecoverNativeAsset is a paid mutator transaction binding the contract method 0x4d00c033.
//
// Solidity: function recoverNativeAsset(uint256 amount) returns()
func (_Transactions *TransactionsTransactorSession) RecoverNativeAsset(amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.RecoverNativeAsset(&_Transactions.TransactOpts, amount)
}

// Reinvest is a paid mutator transaction binding the contract method 0xfdb5a03e.
//
// Solidity: function reinvest() returns()
func (_Transactions *TransactionsTransactor) Reinvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "reinvest")
}

// Reinvest is a paid mutator transaction binding the contract method 0xfdb5a03e.
//
// Solidity: function reinvest() returns()
func (_Transactions *TransactionsSession) Reinvest() (*types.Transaction, error) {
	return _Transactions.Contract.Reinvest(&_Transactions.TransactOpts)
}

// Reinvest is a paid mutator transaction binding the contract method 0xfdb5a03e.
//
// Solidity: function reinvest() returns()
func (_Transactions *TransactionsTransactorSession) Reinvest() (*types.Transaction, error) {
	return _Transactions.Contract.Reinvest(&_Transactions.TransactOpts)
}

// ReinvestOps is a paid mutator transaction binding the contract method 0x5bd93e99.
//
// Solidity: function reinvestOps() returns()
func (_Transactions *TransactionsTransactor) ReinvestOps(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "reinvestOps")
}

// ReinvestOps is a paid mutator transaction binding the contract method 0x5bd93e99.
//
// Solidity: function reinvestOps() returns()
func (_Transactions *TransactionsSession) ReinvestOps() (*types.Transaction, error) {
	return _Transactions.Contract.ReinvestOps(&_Transactions.TransactOpts)
}

// ReinvestOps is a paid mutator transaction binding the contract method 0x5bd93e99.
//
// Solidity: function reinvestOps() returns()
func (_Transactions *TransactionsTransactorSession) ReinvestOps() (*types.Transaction, error) {
	return _Transactions.Contract.ReinvestOps(&_Transactions.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Transactions *TransactionsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Transactions *TransactionsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Transactions.Contract.RenounceOwnership(&_Transactions.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Transactions *TransactionsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Transactions.Contract.RenounceOwnership(&_Transactions.TransactOpts)
}

// RevokeAllowance is a paid mutator transaction binding the contract method 0x7ae26773.
//
// Solidity: function revokeAllowance(address token, address spender) returns()
func (_Transactions *TransactionsTransactor) RevokeAllowance(opts *bind.TransactOpts, token common.Address, spender common.Address) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "revokeAllowance", token, spender)
}

// RevokeAllowance is a paid mutator transaction binding the contract method 0x7ae26773.
//
// Solidity: function revokeAllowance(address token, address spender) returns()
func (_Transactions *TransactionsSession) RevokeAllowance(token common.Address, spender common.Address) (*types.Transaction, error) {
	return _Transactions.Contract.RevokeAllowance(&_Transactions.TransactOpts, token, spender)
}

// RevokeAllowance is a paid mutator transaction binding the contract method 0x7ae26773.
//
// Solidity: function revokeAllowance(address token, address spender) returns()
func (_Transactions *TransactionsTransactorSession) RevokeAllowance(token common.Address, spender common.Address) (*types.Transaction, error) {
	return _Transactions.Contract.RevokeAllowance(&_Transactions.TransactOpts, token, spender)
}

// SetAllowances is a paid mutator transaction binding the contract method 0xdbd9a4d4.
//
// Solidity: function setAllowances() returns()
func (_Transactions *TransactionsTransactor) SetAllowances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "setAllowances")
}

// SetAllowances is a paid mutator transaction binding the contract method 0xdbd9a4d4.
//
// Solidity: function setAllowances() returns()
func (_Transactions *TransactionsSession) SetAllowances() (*types.Transaction, error) {
	return _Transactions.Contract.SetAllowances(&_Transactions.TransactOpts)
}

// SetAllowances is a paid mutator transaction binding the contract method 0xdbd9a4d4.
//
// Solidity: function setAllowances() returns()
func (_Transactions *TransactionsTransactorSession) SetAllowances() (*types.Transaction, error) {
	return _Transactions.Contract.SetAllowances(&_Transactions.TransactOpts)
}

// SingleTokenDeposit is a paid mutator transaction binding the contract method 0xa733e527.
//
// Solidity: function singleTokenDeposit(uint256 amount, address _token, uint256 slippage) returns()
func (_Transactions *TransactionsTransactor) SingleTokenDeposit(opts *bind.TransactOpts, amount *big.Int, _token common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "singleTokenDeposit", amount, _token, slippage)
}

// SingleTokenDeposit is a paid mutator transaction binding the contract method 0xa733e527.
//
// Solidity: function singleTokenDeposit(uint256 amount, address _token, uint256 slippage) returns()
func (_Transactions *TransactionsSession) SingleTokenDeposit(amount *big.Int, _token common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.SingleTokenDeposit(&_Transactions.TransactOpts, amount, _token, slippage)
}

// SingleTokenDeposit is a paid mutator transaction binding the contract method 0xa733e527.
//
// Solidity: function singleTokenDeposit(uint256 amount, address _token, uint256 slippage) returns()
func (_Transactions *TransactionsTransactorSession) SingleTokenDeposit(amount *big.Int, _token common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.SingleTokenDeposit(&_Transactions.TransactOpts, amount, _token, slippage)
}

// SingleWithdraw is a paid mutator transaction binding the contract method 0x743178d4.
//
// Solidity: function singleWithdraw(uint256 amount, address _token) returns()
func (_Transactions *TransactionsTransactor) SingleWithdraw(opts *bind.TransactOpts, amount *big.Int, _token common.Address) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "singleWithdraw", amount, _token)
}

// SingleWithdraw is a paid mutator transaction binding the contract method 0x743178d4.
//
// Solidity: function singleWithdraw(uint256 amount, address _token) returns()
func (_Transactions *TransactionsSession) SingleWithdraw(amount *big.Int, _token common.Address) (*types.Transaction, error) {
	return _Transactions.Contract.SingleWithdraw(&_Transactions.TransactOpts, amount, _token)
}

// SingleWithdraw is a paid mutator transaction binding the contract method 0x743178d4.
//
// Solidity: function singleWithdraw(uint256 amount, address _token) returns()
func (_Transactions *TransactionsTransactorSession) SingleWithdraw(amount *big.Int, _token common.Address) (*types.Transaction, error) {
	return _Transactions.Contract.SingleWithdraw(&_Transactions.TransactOpts, amount, _token)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Transactions *TransactionsTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Transactions *TransactionsSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.Transfer(&_Transactions.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Transactions *TransactionsTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.Transfer(&_Transactions.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Transactions *TransactionsTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Transactions *TransactionsSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.TransferFrom(&_Transactions.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Transactions *TransactionsTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.TransferFrom(&_Transactions.TransactOpts, src, dst, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Transactions *TransactionsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Transactions *TransactionsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Transactions.Contract.TransferOwnership(&_Transactions.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Transactions *TransactionsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Transactions.Contract.TransferOwnership(&_Transactions.TransactOpts, newOwner)
}

// UpdateAdminFee is a paid mutator transaction binding the contract method 0xcff1b6ef.
//
// Solidity: function updateAdminFee(uint256 newValue) returns()
func (_Transactions *TransactionsTransactor) UpdateAdminFee(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "updateAdminFee", newValue)
}

// UpdateAdminFee is a paid mutator transaction binding the contract method 0xcff1b6ef.
//
// Solidity: function updateAdminFee(uint256 newValue) returns()
func (_Transactions *TransactionsSession) UpdateAdminFee(newValue *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.UpdateAdminFee(&_Transactions.TransactOpts, newValue)
}

// UpdateAdminFee is a paid mutator transaction binding the contract method 0xcff1b6ef.
//
// Solidity: function updateAdminFee(uint256 newValue) returns()
func (_Transactions *TransactionsTransactorSession) UpdateAdminFee(newValue *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.UpdateAdminFee(&_Transactions.TransactOpts, newValue)
}

// UpdateAutoCompoundTokenPerBlock is a paid mutator transaction binding the contract method 0x092b176e.
//
// Solidity: function updateAutoCompoundTokenPerBlock(uint256 _autoCompoundTokenPerBlock) returns()
func (_Transactions *TransactionsTransactor) UpdateAutoCompoundTokenPerBlock(opts *bind.TransactOpts, _autoCompoundTokenPerBlock *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "updateAutoCompoundTokenPerBlock", _autoCompoundTokenPerBlock)
}

// UpdateAutoCompoundTokenPerBlock is a paid mutator transaction binding the contract method 0x092b176e.
//
// Solidity: function updateAutoCompoundTokenPerBlock(uint256 _autoCompoundTokenPerBlock) returns()
func (_Transactions *TransactionsSession) UpdateAutoCompoundTokenPerBlock(_autoCompoundTokenPerBlock *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.UpdateAutoCompoundTokenPerBlock(&_Transactions.TransactOpts, _autoCompoundTokenPerBlock)
}

// UpdateAutoCompoundTokenPerBlock is a paid mutator transaction binding the contract method 0x092b176e.
//
// Solidity: function updateAutoCompoundTokenPerBlock(uint256 _autoCompoundTokenPerBlock) returns()
func (_Transactions *TransactionsTransactorSession) UpdateAutoCompoundTokenPerBlock(_autoCompoundTokenPerBlock *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.UpdateAutoCompoundTokenPerBlock(&_Transactions.TransactOpts, _autoCompoundTokenPerBlock)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_Transactions *TransactionsTransactor) UpdateMultiplier(opts *bind.TransactOpts, multiplierNumber *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "updateMultiplier", multiplierNumber)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_Transactions *TransactionsSession) UpdateMultiplier(multiplierNumber *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.UpdateMultiplier(&_Transactions.TransactOpts, multiplierNumber)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_Transactions *TransactionsTransactorSession) UpdateMultiplier(multiplierNumber *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.UpdateMultiplier(&_Transactions.TransactOpts, multiplierNumber)
}

// UpdateReinvestReward is a paid mutator transaction binding the contract method 0xa8ae2b7c.
//
// Solidity: function updateReinvestReward(uint256 newValue) returns()
func (_Transactions *TransactionsTransactor) UpdateReinvestReward(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "updateReinvestReward", newValue)
}

// UpdateReinvestReward is a paid mutator transaction binding the contract method 0xa8ae2b7c.
//
// Solidity: function updateReinvestReward(uint256 newValue) returns()
func (_Transactions *TransactionsSession) UpdateReinvestReward(newValue *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.UpdateReinvestReward(&_Transactions.TransactOpts, newValue)
}

// UpdateReinvestReward is a paid mutator transaction binding the contract method 0xa8ae2b7c.
//
// Solidity: function updateReinvestReward(uint256 newValue) returns()
func (_Transactions *TransactionsTransactorSession) UpdateReinvestReward(newValue *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.UpdateReinvestReward(&_Transactions.TransactOpts, newValue)
}

// UpdateRequireReinvestBeforeDeposit is a paid mutator transaction binding the contract method 0x236aecd5.
//
// Solidity: function updateRequireReinvestBeforeDeposit() returns()
func (_Transactions *TransactionsTransactor) UpdateRequireReinvestBeforeDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "updateRequireReinvestBeforeDeposit")
}

// UpdateRequireReinvestBeforeDeposit is a paid mutator transaction binding the contract method 0x236aecd5.
//
// Solidity: function updateRequireReinvestBeforeDeposit() returns()
func (_Transactions *TransactionsSession) UpdateRequireReinvestBeforeDeposit() (*types.Transaction, error) {
	return _Transactions.Contract.UpdateRequireReinvestBeforeDeposit(&_Transactions.TransactOpts)
}

// UpdateRequireReinvestBeforeDeposit is a paid mutator transaction binding the contract method 0x236aecd5.
//
// Solidity: function updateRequireReinvestBeforeDeposit() returns()
func (_Transactions *TransactionsTransactorSession) UpdateRequireReinvestBeforeDeposit() (*types.Transaction, error) {
	return _Transactions.Contract.UpdateRequireReinvestBeforeDeposit(&_Transactions.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Transactions *TransactionsTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Transactions.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Transactions *TransactionsSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.Withdraw(&_Transactions.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Transactions *TransactionsTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Transactions.Contract.Withdraw(&_Transactions.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Transactions *TransactionsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactions.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Transactions *TransactionsSession) Receive() (*types.Transaction, error) {
	return _Transactions.Contract.Receive(&_Transactions.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Transactions *TransactionsTransactorSession) Receive() (*types.Transaction, error) {
	return _Transactions.Contract.Receive(&_Transactions.TransactOpts)
}

// TransactionsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Transactions contract.
type TransactionsApprovalIterator struct {
	Event *TransactionsApproval // Event containing the contract specifics and raw log

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
func (it *TransactionsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionsApproval)
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
		it.Event = new(TransactionsApproval)
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
func (it *TransactionsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionsApproval represents a Approval event raised by the Transactions contract.
type TransactionsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Transactions *TransactionsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TransactionsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Transactions.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TransactionsApprovalIterator{contract: _Transactions.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Transactions *TransactionsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TransactionsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Transactions.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionsApproval)
				if err := _Transactions.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Transactions *TransactionsFilterer) ParseApproval(log types.Log) (*TransactionsApproval, error) {
	event := new(TransactionsApproval)
	if err := _Transactions.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransactionsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Transactions contract.
type TransactionsDepositIterator struct {
	Event *TransactionsDeposit // Event containing the contract specifics and raw log

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
func (it *TransactionsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionsDeposit)
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
		it.Event = new(TransactionsDeposit)
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
func (it *TransactionsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionsDeposit represents a Deposit event raised by the Transactions contract.
type TransactionsDeposit struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_Transactions *TransactionsFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*TransactionsDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Transactions.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &TransactionsDepositIterator{contract: _Transactions.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_Transactions *TransactionsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TransactionsDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Transactions.contract.WatchLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionsDeposit)
				if err := _Transactions.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_Transactions *TransactionsFilterer) ParseDeposit(log types.Log) (*TransactionsDeposit, error) {
	event := new(TransactionsDeposit)
	if err := _Transactions.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransactionsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Transactions contract.
type TransactionsOwnershipTransferredIterator struct {
	Event *TransactionsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TransactionsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionsOwnershipTransferred)
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
		it.Event = new(TransactionsOwnershipTransferred)
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
func (it *TransactionsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionsOwnershipTransferred represents a OwnershipTransferred event raised by the Transactions contract.
type TransactionsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Transactions *TransactionsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TransactionsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Transactions.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TransactionsOwnershipTransferredIterator{contract: _Transactions.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Transactions *TransactionsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TransactionsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Transactions.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionsOwnershipTransferred)
				if err := _Transactions.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Transactions *TransactionsFilterer) ParseOwnershipTransferred(log types.Log) (*TransactionsOwnershipTransferred, error) {
	event := new(TransactionsOwnershipTransferred)
	if err := _Transactions.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransactionsRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the Transactions contract.
type TransactionsRecoveredIterator struct {
	Event *TransactionsRecovered // Event containing the contract specifics and raw log

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
func (it *TransactionsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionsRecovered)
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
		it.Event = new(TransactionsRecovered)
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
func (it *TransactionsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionsRecovered represents a Recovered event raised by the Transactions contract.
type TransactionsRecovered struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_Transactions *TransactionsFilterer) FilterRecovered(opts *bind.FilterOpts) (*TransactionsRecoveredIterator, error) {

	logs, sub, err := _Transactions.contract.FilterLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return &TransactionsRecoveredIterator{contract: _Transactions.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_Transactions *TransactionsFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *TransactionsRecovered) (event.Subscription, error) {

	logs, sub, err := _Transactions.contract.WatchLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionsRecovered)
				if err := _Transactions.contract.UnpackLog(event, "Recovered", log); err != nil {
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

// ParseRecovered is a log parse operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_Transactions *TransactionsFilterer) ParseRecovered(log types.Log) (*TransactionsRecovered, error) {
	event := new(TransactionsRecovered)
	if err := _Transactions.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransactionsReinvestIterator is returned from FilterReinvest and is used to iterate over the raw logs and unpacked data for Reinvest events raised by the Transactions contract.
type TransactionsReinvestIterator struct {
	Event *TransactionsReinvest // Event containing the contract specifics and raw log

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
func (it *TransactionsReinvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionsReinvest)
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
		it.Event = new(TransactionsReinvest)
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
func (it *TransactionsReinvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionsReinvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionsReinvest represents a Reinvest event raised by the Transactions contract.
type TransactionsReinvest struct {
	NewTotalDeposits *big.Int
	NewTotalSupply   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReinvest is a free log retrieval operation binding the contract event 0xc7606d21ac05cd309191543e409f0845c016120563783d70e4f41419dc0ef234.
//
// Solidity: event Reinvest(uint256 newTotalDeposits, uint256 newTotalSupply)
func (_Transactions *TransactionsFilterer) FilterReinvest(opts *bind.FilterOpts) (*TransactionsReinvestIterator, error) {

	logs, sub, err := _Transactions.contract.FilterLogs(opts, "Reinvest")
	if err != nil {
		return nil, err
	}
	return &TransactionsReinvestIterator{contract: _Transactions.contract, event: "Reinvest", logs: logs, sub: sub}, nil
}

// WatchReinvest is a free log subscription operation binding the contract event 0xc7606d21ac05cd309191543e409f0845c016120563783d70e4f41419dc0ef234.
//
// Solidity: event Reinvest(uint256 newTotalDeposits, uint256 newTotalSupply)
func (_Transactions *TransactionsFilterer) WatchReinvest(opts *bind.WatchOpts, sink chan<- *TransactionsReinvest) (event.Subscription, error) {

	logs, sub, err := _Transactions.contract.WatchLogs(opts, "Reinvest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionsReinvest)
				if err := _Transactions.contract.UnpackLog(event, "Reinvest", log); err != nil {
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

// ParseReinvest is a log parse operation binding the contract event 0xc7606d21ac05cd309191543e409f0845c016120563783d70e4f41419dc0ef234.
//
// Solidity: event Reinvest(uint256 newTotalDeposits, uint256 newTotalSupply)
func (_Transactions *TransactionsFilterer) ParseReinvest(log types.Log) (*TransactionsReinvest, error) {
	event := new(TransactionsReinvest)
	if err := _Transactions.contract.UnpackLog(event, "Reinvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransactionsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Transactions contract.
type TransactionsTransferIterator struct {
	Event *TransactionsTransfer // Event containing the contract specifics and raw log

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
func (it *TransactionsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionsTransfer)
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
		it.Event = new(TransactionsTransfer)
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
func (it *TransactionsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionsTransfer represents a Transfer event raised by the Transactions contract.
type TransactionsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Transactions *TransactionsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TransactionsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Transactions.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TransactionsTransferIterator{contract: _Transactions.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Transactions *TransactionsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TransactionsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Transactions.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionsTransfer)
				if err := _Transactions.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Transactions *TransactionsFilterer) ParseTransfer(log types.Log) (*TransactionsTransfer, error) {
	event := new(TransactionsTransfer)
	if err := _Transactions.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransactionsUpdateAdminFeeIterator is returned from FilterUpdateAdminFee and is used to iterate over the raw logs and unpacked data for UpdateAdminFee events raised by the Transactions contract.
type TransactionsUpdateAdminFeeIterator struct {
	Event *TransactionsUpdateAdminFee // Event containing the contract specifics and raw log

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
func (it *TransactionsUpdateAdminFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionsUpdateAdminFee)
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
		it.Event = new(TransactionsUpdateAdminFee)
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
func (it *TransactionsUpdateAdminFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionsUpdateAdminFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionsUpdateAdminFee represents a UpdateAdminFee event raised by the Transactions contract.
type TransactionsUpdateAdminFee struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateAdminFee is a free log retrieval operation binding the contract event 0x3cc372f330f95ac9540626dc8a25f5bf21ba607215a5d58304cb804d446f104a.
//
// Solidity: event UpdateAdminFee(uint256 oldValue, uint256 newValue)
func (_Transactions *TransactionsFilterer) FilterUpdateAdminFee(opts *bind.FilterOpts) (*TransactionsUpdateAdminFeeIterator, error) {

	logs, sub, err := _Transactions.contract.FilterLogs(opts, "UpdateAdminFee")
	if err != nil {
		return nil, err
	}
	return &TransactionsUpdateAdminFeeIterator{contract: _Transactions.contract, event: "UpdateAdminFee", logs: logs, sub: sub}, nil
}

// WatchUpdateAdminFee is a free log subscription operation binding the contract event 0x3cc372f330f95ac9540626dc8a25f5bf21ba607215a5d58304cb804d446f104a.
//
// Solidity: event UpdateAdminFee(uint256 oldValue, uint256 newValue)
func (_Transactions *TransactionsFilterer) WatchUpdateAdminFee(opts *bind.WatchOpts, sink chan<- *TransactionsUpdateAdminFee) (event.Subscription, error) {

	logs, sub, err := _Transactions.contract.WatchLogs(opts, "UpdateAdminFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionsUpdateAdminFee)
				if err := _Transactions.contract.UnpackLog(event, "UpdateAdminFee", log); err != nil {
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

// ParseUpdateAdminFee is a log parse operation binding the contract event 0x3cc372f330f95ac9540626dc8a25f5bf21ba607215a5d58304cb804d446f104a.
//
// Solidity: event UpdateAdminFee(uint256 oldValue, uint256 newValue)
func (_Transactions *TransactionsFilterer) ParseUpdateAdminFee(log types.Log) (*TransactionsUpdateAdminFee, error) {
	event := new(TransactionsUpdateAdminFee)
	if err := _Transactions.contract.UnpackLog(event, "UpdateAdminFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransactionsUpdateReinvestRewardIterator is returned from FilterUpdateReinvestReward and is used to iterate over the raw logs and unpacked data for UpdateReinvestReward events raised by the Transactions contract.
type TransactionsUpdateReinvestRewardIterator struct {
	Event *TransactionsUpdateReinvestReward // Event containing the contract specifics and raw log

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
func (it *TransactionsUpdateReinvestRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionsUpdateReinvestReward)
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
		it.Event = new(TransactionsUpdateReinvestReward)
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
func (it *TransactionsUpdateReinvestRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionsUpdateReinvestRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionsUpdateReinvestReward represents a UpdateReinvestReward event raised by the Transactions contract.
type TransactionsUpdateReinvestReward struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateReinvestReward is a free log retrieval operation binding the contract event 0xe7f97d51d307dc44045597c9978bec0f842e6bb40d19b9444084cfa30d9ed4f2.
//
// Solidity: event UpdateReinvestReward(uint256 oldValue, uint256 newValue)
func (_Transactions *TransactionsFilterer) FilterUpdateReinvestReward(opts *bind.FilterOpts) (*TransactionsUpdateReinvestRewardIterator, error) {

	logs, sub, err := _Transactions.contract.FilterLogs(opts, "UpdateReinvestReward")
	if err != nil {
		return nil, err
	}
	return &TransactionsUpdateReinvestRewardIterator{contract: _Transactions.contract, event: "UpdateReinvestReward", logs: logs, sub: sub}, nil
}

// WatchUpdateReinvestReward is a free log subscription operation binding the contract event 0xe7f97d51d307dc44045597c9978bec0f842e6bb40d19b9444084cfa30d9ed4f2.
//
// Solidity: event UpdateReinvestReward(uint256 oldValue, uint256 newValue)
func (_Transactions *TransactionsFilterer) WatchUpdateReinvestReward(opts *bind.WatchOpts, sink chan<- *TransactionsUpdateReinvestReward) (event.Subscription, error) {

	logs, sub, err := _Transactions.contract.WatchLogs(opts, "UpdateReinvestReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionsUpdateReinvestReward)
				if err := _Transactions.contract.UnpackLog(event, "UpdateReinvestReward", log); err != nil {
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

// ParseUpdateReinvestReward is a log parse operation binding the contract event 0xe7f97d51d307dc44045597c9978bec0f842e6bb40d19b9444084cfa30d9ed4f2.
//
// Solidity: event UpdateReinvestReward(uint256 oldValue, uint256 newValue)
func (_Transactions *TransactionsFilterer) ParseUpdateReinvestReward(log types.Log) (*TransactionsUpdateReinvestReward, error) {
	event := new(TransactionsUpdateReinvestReward)
	if err := _Transactions.contract.UnpackLog(event, "UpdateReinvestReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransactionsUpdateRequireReinvestBeforeDepositIterator is returned from FilterUpdateRequireReinvestBeforeDeposit and is used to iterate over the raw logs and unpacked data for UpdateRequireReinvestBeforeDeposit events raised by the Transactions contract.
type TransactionsUpdateRequireReinvestBeforeDepositIterator struct {
	Event *TransactionsUpdateRequireReinvestBeforeDeposit // Event containing the contract specifics and raw log

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
func (it *TransactionsUpdateRequireReinvestBeforeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionsUpdateRequireReinvestBeforeDeposit)
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
		it.Event = new(TransactionsUpdateRequireReinvestBeforeDeposit)
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
func (it *TransactionsUpdateRequireReinvestBeforeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionsUpdateRequireReinvestBeforeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionsUpdateRequireReinvestBeforeDeposit represents a UpdateRequireReinvestBeforeDeposit event raised by the Transactions contract.
type TransactionsUpdateRequireReinvestBeforeDeposit struct {
	NewValue bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateRequireReinvestBeforeDeposit is a free log retrieval operation binding the contract event 0xd46852adf64681b12b81a308b01efd0a546667f68ab41ae5855c2bae7756380f.
//
// Solidity: event UpdateRequireReinvestBeforeDeposit(bool newValue)
func (_Transactions *TransactionsFilterer) FilterUpdateRequireReinvestBeforeDeposit(opts *bind.FilterOpts) (*TransactionsUpdateRequireReinvestBeforeDepositIterator, error) {

	logs, sub, err := _Transactions.contract.FilterLogs(opts, "UpdateRequireReinvestBeforeDeposit")
	if err != nil {
		return nil, err
	}
	return &TransactionsUpdateRequireReinvestBeforeDepositIterator{contract: _Transactions.contract, event: "UpdateRequireReinvestBeforeDeposit", logs: logs, sub: sub}, nil
}

// WatchUpdateRequireReinvestBeforeDeposit is a free log subscription operation binding the contract event 0xd46852adf64681b12b81a308b01efd0a546667f68ab41ae5855c2bae7756380f.
//
// Solidity: event UpdateRequireReinvestBeforeDeposit(bool newValue)
func (_Transactions *TransactionsFilterer) WatchUpdateRequireReinvestBeforeDeposit(opts *bind.WatchOpts, sink chan<- *TransactionsUpdateRequireReinvestBeforeDeposit) (event.Subscription, error) {

	logs, sub, err := _Transactions.contract.WatchLogs(opts, "UpdateRequireReinvestBeforeDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionsUpdateRequireReinvestBeforeDeposit)
				if err := _Transactions.contract.UnpackLog(event, "UpdateRequireReinvestBeforeDeposit", log); err != nil {
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

// ParseUpdateRequireReinvestBeforeDeposit is a log parse operation binding the contract event 0xd46852adf64681b12b81a308b01efd0a546667f68ab41ae5855c2bae7756380f.
//
// Solidity: event UpdateRequireReinvestBeforeDeposit(bool newValue)
func (_Transactions *TransactionsFilterer) ParseUpdateRequireReinvestBeforeDeposit(log types.Log) (*TransactionsUpdateRequireReinvestBeforeDeposit, error) {
	event := new(TransactionsUpdateRequireReinvestBeforeDeposit)
	if err := _Transactions.contract.UnpackLog(event, "UpdateRequireReinvestBeforeDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransactionsWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Transactions contract.
type TransactionsWithdrawIterator struct {
	Event *TransactionsWithdraw // Event containing the contract specifics and raw log

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
func (it *TransactionsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionsWithdraw)
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
		it.Event = new(TransactionsWithdraw)
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
func (it *TransactionsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionsWithdraw represents a Withdraw event raised by the Transactions contract.
type TransactionsWithdraw struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_Transactions *TransactionsFilterer) FilterWithdraw(opts *bind.FilterOpts, account []common.Address) (*TransactionsWithdrawIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Transactions.contract.FilterLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return &TransactionsWithdrawIterator{contract: _Transactions.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_Transactions *TransactionsFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TransactionsWithdraw, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Transactions.contract.WatchLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionsWithdraw)
				if err := _Transactions.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_Transactions *TransactionsFilterer) ParseWithdraw(log types.Log) (*TransactionsWithdraw, error) {
	event := new(TransactionsWithdraw)
	if err := _Transactions.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
