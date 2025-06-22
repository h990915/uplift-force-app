// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// BoostChainMainContractOrder is an auto generated low-level Go binding around an user-defined struct.
type BoostChainMainContractOrder struct {
	OrderId         *big.Int
	Player          common.Address
	Booster         common.Address
	TotalAmount     *big.Int
	PlayerDeposit   *big.Int
	BoosterDeposit  *big.Int
	RemainingAmount *big.Int
	Status          uint8
	Deadline        *big.Int
	ChildContract   common.Address
	CreatedAt       *big.Int
	AcceptedAt      *big.Int
	ConfirmedAt     *big.Int
	GameType        string
	GameMode        string
	Requirements    string
}

// BoostChainMainContractMetaData contains all meta data concerning the BoostChainMainContract contract.
var BoostChainMainContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_platformTreasury\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"booster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boosterDeposit\",\"type\":\"uint256\"}],\"name\":\"OrderAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cancelledBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penaltyAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"penaltyReceiver\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boosterReward\",\"type\":\"uint256\"}],\"name\":\"OrderCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmountLocked\",\"type\":\"uint256\"}],\"name\":\"OrderConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"playerDeposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gameType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gameMode\",\"type\":\"string\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"playerRefund\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penaltyToPlayer\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penaltyToPlatform\",\"type\":\"uint256\"}],\"name\":\"OrderFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"acceptOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boosterOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"calculateDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"completeOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"confirmOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_gameType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_gameMode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_requirements\",\"type\":\"string\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"failOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_booster\",\"type\":\"address\"}],\"name\":\"getBoosterOrders\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"booster\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"playerDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boosterDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBoostChainMainContract.OrderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"childContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"confirmedAt\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gameType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gameMode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"requirements\",\"type\":\"string\"}],\"internalType\":\"structBoostChainMainContract.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_player\",\"type\":\"address\"}],\"name\":\"getPlayerOrders\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"booster\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"playerDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boosterDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBoostChainMainContract.OrderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"childContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"confirmedAt\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gameType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gameMode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"requirements\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyToPlatformRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyToVictimRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"playerOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"}],\"name\":\"setDepositRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"}],\"name\":\"setPlatformFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTreasury\",\"type\":\"address\"}],\"name\":\"setPlatformTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BoostChainMainContractABI is the input ABI used to generate the binding from.
// Deprecated: Use BoostChainMainContractMetaData.ABI instead.
var BoostChainMainContractABI = BoostChainMainContractMetaData.ABI

// BoostChainMainContract is an auto generated Go binding around an Ethereum contract.
type BoostChainMainContract struct {
	BoostChainMainContractCaller     // Read-only binding to the contract
	BoostChainMainContractTransactor // Write-only binding to the contract
	BoostChainMainContractFilterer   // Log filterer for contract events
}

// BoostChainMainContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoostChainMainContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoostChainMainContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoostChainMainContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoostChainMainContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoostChainMainContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoostChainMainContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoostChainMainContractSession struct {
	Contract     *BoostChainMainContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BoostChainMainContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoostChainMainContractCallerSession struct {
	Contract *BoostChainMainContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// BoostChainMainContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoostChainMainContractTransactorSession struct {
	Contract     *BoostChainMainContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// BoostChainMainContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoostChainMainContractRaw struct {
	Contract *BoostChainMainContract // Generic contract binding to access the raw methods on
}

// BoostChainMainContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoostChainMainContractCallerRaw struct {
	Contract *BoostChainMainContractCaller // Generic read-only contract binding to access the raw methods on
}

// BoostChainMainContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoostChainMainContractTransactorRaw struct {
	Contract *BoostChainMainContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBoostChainMainContract creates a new instance of BoostChainMainContract, bound to a specific deployed contract.
func NewBoostChainMainContract(address common.Address, backend bind.ContractBackend) (*BoostChainMainContract, error) {
	contract, err := bindBoostChainMainContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContract{BoostChainMainContractCaller: BoostChainMainContractCaller{contract: contract}, BoostChainMainContractTransactor: BoostChainMainContractTransactor{contract: contract}, BoostChainMainContractFilterer: BoostChainMainContractFilterer{contract: contract}}, nil
}

// NewBoostChainMainContractCaller creates a new read-only instance of BoostChainMainContract, bound to a specific deployed contract.
func NewBoostChainMainContractCaller(address common.Address, caller bind.ContractCaller) (*BoostChainMainContractCaller, error) {
	contract, err := bindBoostChainMainContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContractCaller{contract: contract}, nil
}

// NewBoostChainMainContractTransactor creates a new write-only instance of BoostChainMainContract, bound to a specific deployed contract.
func NewBoostChainMainContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BoostChainMainContractTransactor, error) {
	contract, err := bindBoostChainMainContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContractTransactor{contract: contract}, nil
}

// NewBoostChainMainContractFilterer creates a new log filterer instance of BoostChainMainContract, bound to a specific deployed contract.
func NewBoostChainMainContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BoostChainMainContractFilterer, error) {
	contract, err := bindBoostChainMainContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContractFilterer{contract: contract}, nil
}

// bindBoostChainMainContract binds a generic wrapper to an already deployed contract.
func bindBoostChainMainContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BoostChainMainContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoostChainMainContract *BoostChainMainContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BoostChainMainContract.Contract.BoostChainMainContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoostChainMainContract *BoostChainMainContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.BoostChainMainContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoostChainMainContract *BoostChainMainContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.BoostChainMainContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoostChainMainContract *BoostChainMainContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BoostChainMainContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoostChainMainContract *BoostChainMainContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoostChainMainContract *BoostChainMainContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCaller) BASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractSession) BASISPOINTS() (*big.Int, error) {
	return _BoostChainMainContract.Contract.BASISPOINTS(&_BoostChainMainContract.CallOpts)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) BASISPOINTS() (*big.Int, error) {
	return _BoostChainMainContract.Contract.BASISPOINTS(&_BoostChainMainContract.CallOpts)
}

// BoosterOrders is a free data retrieval call binding the contract method 0x22961a1a.
//
// Solidity: function boosterOrders(address , uint256 ) view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCaller) BoosterOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "boosterOrders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoosterOrders is a free data retrieval call binding the contract method 0x22961a1a.
//
// Solidity: function boosterOrders(address , uint256 ) view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractSession) BoosterOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BoostChainMainContract.Contract.BoosterOrders(&_BoostChainMainContract.CallOpts, arg0, arg1)
}

// BoosterOrders is a free data retrieval call binding the contract method 0x22961a1a.
//
// Solidity: function boosterOrders(address , uint256 ) view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) BoosterOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BoostChainMainContract.Contract.BoosterOrders(&_BoostChainMainContract.CallOpts, arg0, arg1)
}

// CalculateDeposit is a free data retrieval call binding the contract method 0x4ec508f9.
//
// Solidity: function calculateDeposit(uint256 _totalAmount) view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCaller) CalculateDeposit(opts *bind.CallOpts, _totalAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "calculateDeposit", _totalAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateDeposit is a free data retrieval call binding the contract method 0x4ec508f9.
//
// Solidity: function calculateDeposit(uint256 _totalAmount) view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractSession) CalculateDeposit(_totalAmount *big.Int) (*big.Int, error) {
	return _BoostChainMainContract.Contract.CalculateDeposit(&_BoostChainMainContract.CallOpts, _totalAmount)
}

// CalculateDeposit is a free data retrieval call binding the contract method 0x4ec508f9.
//
// Solidity: function calculateDeposit(uint256 _totalAmount) view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) CalculateDeposit(_totalAmount *big.Int) (*big.Int, error) {
	return _BoostChainMainContract.Contract.CalculateDeposit(&_BoostChainMainContract.CallOpts, _totalAmount)
}

// DepositRate is a free data retrieval call binding the contract method 0x7e74a1ed.
//
// Solidity: function depositRate() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCaller) DepositRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "depositRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositRate is a free data retrieval call binding the contract method 0x7e74a1ed.
//
// Solidity: function depositRate() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractSession) DepositRate() (*big.Int, error) {
	return _BoostChainMainContract.Contract.DepositRate(&_BoostChainMainContract.CallOpts)
}

// DepositRate is a free data retrieval call binding the contract method 0x7e74a1ed.
//
// Solidity: function depositRate() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) DepositRate() (*big.Int, error) {
	return _BoostChainMainContract.Contract.DepositRate(&_BoostChainMainContract.CallOpts)
}

// GetBoosterOrders is a free data retrieval call binding the contract method 0xef27b867.
//
// Solidity: function getBoosterOrders(address _booster) view returns(uint256[])
func (_BoostChainMainContract *BoostChainMainContractCaller) GetBoosterOrders(opts *bind.CallOpts, _booster common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "getBoosterOrders", _booster)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBoosterOrders is a free data retrieval call binding the contract method 0xef27b867.
//
// Solidity: function getBoosterOrders(address _booster) view returns(uint256[])
func (_BoostChainMainContract *BoostChainMainContractSession) GetBoosterOrders(_booster common.Address) ([]*big.Int, error) {
	return _BoostChainMainContract.Contract.GetBoosterOrders(&_BoostChainMainContract.CallOpts, _booster)
}

// GetBoosterOrders is a free data retrieval call binding the contract method 0xef27b867.
//
// Solidity: function getBoosterOrders(address _booster) view returns(uint256[])
func (_BoostChainMainContract *BoostChainMainContractCallerSession) GetBoosterOrders(_booster common.Address) ([]*big.Int, error) {
	return _BoostChainMainContract.Contract.GetBoosterOrders(&_BoostChainMainContract.CallOpts, _booster)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 _orderId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint8,uint256,address,uint256,uint256,uint256,string,string,string))
func (_BoostChainMainContract *BoostChainMainContractCaller) GetOrder(opts *bind.CallOpts, _orderId *big.Int) (BoostChainMainContractOrder, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "getOrder", _orderId)

	if err != nil {
		return *new(BoostChainMainContractOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(BoostChainMainContractOrder)).(*BoostChainMainContractOrder)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 _orderId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint8,uint256,address,uint256,uint256,uint256,string,string,string))
func (_BoostChainMainContract *BoostChainMainContractSession) GetOrder(_orderId *big.Int) (BoostChainMainContractOrder, error) {
	return _BoostChainMainContract.Contract.GetOrder(&_BoostChainMainContract.CallOpts, _orderId)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 _orderId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint8,uint256,address,uint256,uint256,uint256,string,string,string))
func (_BoostChainMainContract *BoostChainMainContractCallerSession) GetOrder(_orderId *big.Int) (BoostChainMainContractOrder, error) {
	return _BoostChainMainContract.Contract.GetOrder(&_BoostChainMainContract.CallOpts, _orderId)
}

// GetPlayerOrders is a free data retrieval call binding the contract method 0x1b1d30c6.
//
// Solidity: function getPlayerOrders(address _player) view returns(uint256[])
func (_BoostChainMainContract *BoostChainMainContractCaller) GetPlayerOrders(opts *bind.CallOpts, _player common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "getPlayerOrders", _player)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPlayerOrders is a free data retrieval call binding the contract method 0x1b1d30c6.
//
// Solidity: function getPlayerOrders(address _player) view returns(uint256[])
func (_BoostChainMainContract *BoostChainMainContractSession) GetPlayerOrders(_player common.Address) ([]*big.Int, error) {
	return _BoostChainMainContract.Contract.GetPlayerOrders(&_BoostChainMainContract.CallOpts, _player)
}

// GetPlayerOrders is a free data retrieval call binding the contract method 0x1b1d30c6.
//
// Solidity: function getPlayerOrders(address _player) view returns(uint256[])
func (_BoostChainMainContract *BoostChainMainContractCallerSession) GetPlayerOrders(_player common.Address) ([]*big.Int, error) {
	return _BoostChainMainContract.Contract.GetPlayerOrders(&_BoostChainMainContract.CallOpts, _player)
}

// OrderCounter is a free data retrieval call binding the contract method 0xb789bf52.
//
// Solidity: function orderCounter() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCaller) OrderCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "orderCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderCounter is a free data retrieval call binding the contract method 0xb789bf52.
//
// Solidity: function orderCounter() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractSession) OrderCounter() (*big.Int, error) {
	return _BoostChainMainContract.Contract.OrderCounter(&_BoostChainMainContract.CallOpts)
}

// OrderCounter is a free data retrieval call binding the contract method 0xb789bf52.
//
// Solidity: function orderCounter() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) OrderCounter() (*big.Int, error) {
	return _BoostChainMainContract.Contract.OrderCounter(&_BoostChainMainContract.CallOpts)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 orderId, address player, address booster, uint256 totalAmount, uint256 playerDeposit, uint256 boosterDeposit, uint256 remainingAmount, uint8 status, uint256 deadline, address childContract, uint256 createdAt, uint256 acceptedAt, uint256 confirmedAt, string gameType, string gameMode, string requirements)
func (_BoostChainMainContract *BoostChainMainContractCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OrderId         *big.Int
	Player          common.Address
	Booster         common.Address
	TotalAmount     *big.Int
	PlayerDeposit   *big.Int
	BoosterDeposit  *big.Int
	RemainingAmount *big.Int
	Status          uint8
	Deadline        *big.Int
	ChildContract   common.Address
	CreatedAt       *big.Int
	AcceptedAt      *big.Int
	ConfirmedAt     *big.Int
	GameType        string
	GameMode        string
	Requirements    string
}, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		OrderId         *big.Int
		Player          common.Address
		Booster         common.Address
		TotalAmount     *big.Int
		PlayerDeposit   *big.Int
		BoosterDeposit  *big.Int
		RemainingAmount *big.Int
		Status          uint8
		Deadline        *big.Int
		ChildContract   common.Address
		CreatedAt       *big.Int
		AcceptedAt      *big.Int
		ConfirmedAt     *big.Int
		GameType        string
		GameMode        string
		Requirements    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Player = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Booster = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TotalAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PlayerDeposit = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BoosterDeposit = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RemainingAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.Deadline = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.ChildContract = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.AcceptedAt = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.ConfirmedAt = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.GameType = *abi.ConvertType(out[13], new(string)).(*string)
	outstruct.GameMode = *abi.ConvertType(out[14], new(string)).(*string)
	outstruct.Requirements = *abi.ConvertType(out[15], new(string)).(*string)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 orderId, address player, address booster, uint256 totalAmount, uint256 playerDeposit, uint256 boosterDeposit, uint256 remainingAmount, uint8 status, uint256 deadline, address childContract, uint256 createdAt, uint256 acceptedAt, uint256 confirmedAt, string gameType, string gameMode, string requirements)
func (_BoostChainMainContract *BoostChainMainContractSession) Orders(arg0 *big.Int) (struct {
	OrderId         *big.Int
	Player          common.Address
	Booster         common.Address
	TotalAmount     *big.Int
	PlayerDeposit   *big.Int
	BoosterDeposit  *big.Int
	RemainingAmount *big.Int
	Status          uint8
	Deadline        *big.Int
	ChildContract   common.Address
	CreatedAt       *big.Int
	AcceptedAt      *big.Int
	ConfirmedAt     *big.Int
	GameType        string
	GameMode        string
	Requirements    string
}, error) {
	return _BoostChainMainContract.Contract.Orders(&_BoostChainMainContract.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 orderId, address player, address booster, uint256 totalAmount, uint256 playerDeposit, uint256 boosterDeposit, uint256 remainingAmount, uint8 status, uint256 deadline, address childContract, uint256 createdAt, uint256 acceptedAt, uint256 confirmedAt, string gameType, string gameMode, string requirements)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) Orders(arg0 *big.Int) (struct {
	OrderId         *big.Int
	Player          common.Address
	Booster         common.Address
	TotalAmount     *big.Int
	PlayerDeposit   *big.Int
	BoosterDeposit  *big.Int
	RemainingAmount *big.Int
	Status          uint8
	Deadline        *big.Int
	ChildContract   common.Address
	CreatedAt       *big.Int
	AcceptedAt      *big.Int
	ConfirmedAt     *big.Int
	GameType        string
	GameMode        string
	Requirements    string
}, error) {
	return _BoostChainMainContract.Contract.Orders(&_BoostChainMainContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BoostChainMainContract *BoostChainMainContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BoostChainMainContract *BoostChainMainContractSession) Owner() (common.Address, error) {
	return _BoostChainMainContract.Contract.Owner(&_BoostChainMainContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) Owner() (common.Address, error) {
	return _BoostChainMainContract.Contract.Owner(&_BoostChainMainContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BoostChainMainContract *BoostChainMainContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BoostChainMainContract *BoostChainMainContractSession) Paused() (bool, error) {
	return _BoostChainMainContract.Contract.Paused(&_BoostChainMainContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) Paused() (bool, error) {
	return _BoostChainMainContract.Contract.Paused(&_BoostChainMainContract.CallOpts)
}

// PenaltyToPlatformRate is a free data retrieval call binding the contract method 0xecf4f557.
//
// Solidity: function penaltyToPlatformRate() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCaller) PenaltyToPlatformRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "penaltyToPlatformRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyToPlatformRate is a free data retrieval call binding the contract method 0xecf4f557.
//
// Solidity: function penaltyToPlatformRate() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractSession) PenaltyToPlatformRate() (*big.Int, error) {
	return _BoostChainMainContract.Contract.PenaltyToPlatformRate(&_BoostChainMainContract.CallOpts)
}

// PenaltyToPlatformRate is a free data retrieval call binding the contract method 0xecf4f557.
//
// Solidity: function penaltyToPlatformRate() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) PenaltyToPlatformRate() (*big.Int, error) {
	return _BoostChainMainContract.Contract.PenaltyToPlatformRate(&_BoostChainMainContract.CallOpts)
}

// PenaltyToVictimRate is a free data retrieval call binding the contract method 0x8cbe8ae8.
//
// Solidity: function penaltyToVictimRate() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCaller) PenaltyToVictimRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "penaltyToVictimRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyToVictimRate is a free data retrieval call binding the contract method 0x8cbe8ae8.
//
// Solidity: function penaltyToVictimRate() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractSession) PenaltyToVictimRate() (*big.Int, error) {
	return _BoostChainMainContract.Contract.PenaltyToVictimRate(&_BoostChainMainContract.CallOpts)
}

// PenaltyToVictimRate is a free data retrieval call binding the contract method 0x8cbe8ae8.
//
// Solidity: function penaltyToVictimRate() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) PenaltyToVictimRate() (*big.Int, error) {
	return _BoostChainMainContract.Contract.PenaltyToVictimRate(&_BoostChainMainContract.CallOpts)
}

// PlatformFeeRate is a free data retrieval call binding the contract method 0xeeca08f0.
//
// Solidity: function platformFeeRate() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCaller) PlatformFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "platformFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFeeRate is a free data retrieval call binding the contract method 0xeeca08f0.
//
// Solidity: function platformFeeRate() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractSession) PlatformFeeRate() (*big.Int, error) {
	return _BoostChainMainContract.Contract.PlatformFeeRate(&_BoostChainMainContract.CallOpts)
}

// PlatformFeeRate is a free data retrieval call binding the contract method 0xeeca08f0.
//
// Solidity: function platformFeeRate() view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) PlatformFeeRate() (*big.Int, error) {
	return _BoostChainMainContract.Contract.PlatformFeeRate(&_BoostChainMainContract.CallOpts)
}

// PlatformTreasury is a free data retrieval call binding the contract method 0xe138818c.
//
// Solidity: function platformTreasury() view returns(address)
func (_BoostChainMainContract *BoostChainMainContractCaller) PlatformTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "platformTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatformTreasury is a free data retrieval call binding the contract method 0xe138818c.
//
// Solidity: function platformTreasury() view returns(address)
func (_BoostChainMainContract *BoostChainMainContractSession) PlatformTreasury() (common.Address, error) {
	return _BoostChainMainContract.Contract.PlatformTreasury(&_BoostChainMainContract.CallOpts)
}

// PlatformTreasury is a free data retrieval call binding the contract method 0xe138818c.
//
// Solidity: function platformTreasury() view returns(address)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) PlatformTreasury() (common.Address, error) {
	return _BoostChainMainContract.Contract.PlatformTreasury(&_BoostChainMainContract.CallOpts)
}

// PlayerOrders is a free data retrieval call binding the contract method 0x402c76fd.
//
// Solidity: function playerOrders(address , uint256 ) view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCaller) PlayerOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BoostChainMainContract.contract.Call(opts, &out, "playerOrders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlayerOrders is a free data retrieval call binding the contract method 0x402c76fd.
//
// Solidity: function playerOrders(address , uint256 ) view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractSession) PlayerOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BoostChainMainContract.Contract.PlayerOrders(&_BoostChainMainContract.CallOpts, arg0, arg1)
}

// PlayerOrders is a free data retrieval call binding the contract method 0x402c76fd.
//
// Solidity: function playerOrders(address , uint256 ) view returns(uint256)
func (_BoostChainMainContract *BoostChainMainContractCallerSession) PlayerOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BoostChainMainContract.Contract.PlayerOrders(&_BoostChainMainContract.CallOpts, arg0, arg1)
}

// AcceptOrder is a paid mutator transaction binding the contract method 0xef18e9ed.
//
// Solidity: function acceptOrder(uint256 _orderId) payable returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) AcceptOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "acceptOrder", _orderId)
}

// AcceptOrder is a paid mutator transaction binding the contract method 0xef18e9ed.
//
// Solidity: function acceptOrder(uint256 _orderId) payable returns()
func (_BoostChainMainContract *BoostChainMainContractSession) AcceptOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.AcceptOrder(&_BoostChainMainContract.TransactOpts, _orderId)
}

// AcceptOrder is a paid mutator transaction binding the contract method 0xef18e9ed.
//
// Solidity: function acceptOrder(uint256 _orderId) payable returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) AcceptOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.AcceptOrder(&_BoostChainMainContract.TransactOpts, _orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _orderId) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) CancelOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "cancelOrder", _orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _orderId) returns()
func (_BoostChainMainContract *BoostChainMainContractSession) CancelOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.CancelOrder(&_BoostChainMainContract.TransactOpts, _orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _orderId) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) CancelOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.CancelOrder(&_BoostChainMainContract.TransactOpts, _orderId)
}

// CompleteOrder is a paid mutator transaction binding the contract method 0xb6adaaff.
//
// Solidity: function completeOrder(uint256 _orderId) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) CompleteOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "completeOrder", _orderId)
}

// CompleteOrder is a paid mutator transaction binding the contract method 0xb6adaaff.
//
// Solidity: function completeOrder(uint256 _orderId) returns()
func (_BoostChainMainContract *BoostChainMainContractSession) CompleteOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.CompleteOrder(&_BoostChainMainContract.TransactOpts, _orderId)
}

// CompleteOrder is a paid mutator transaction binding the contract method 0xb6adaaff.
//
// Solidity: function completeOrder(uint256 _orderId) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) CompleteOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.CompleteOrder(&_BoostChainMainContract.TransactOpts, _orderId)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0x8ac7d79c.
//
// Solidity: function confirmOrder(uint256 _orderId) payable returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) ConfirmOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "confirmOrder", _orderId)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0x8ac7d79c.
//
// Solidity: function confirmOrder(uint256 _orderId) payable returns()
func (_BoostChainMainContract *BoostChainMainContractSession) ConfirmOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.ConfirmOrder(&_BoostChainMainContract.TransactOpts, _orderId)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0x8ac7d79c.
//
// Solidity: function confirmOrder(uint256 _orderId) payable returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) ConfirmOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.ConfirmOrder(&_BoostChainMainContract.TransactOpts, _orderId)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xf62154ad.
//
// Solidity: function createOrder(uint256 _totalAmount, uint256 _deadline, string _gameType, string _gameMode, string _requirements) payable returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) CreateOrder(opts *bind.TransactOpts, _totalAmount *big.Int, _deadline *big.Int, _gameType string, _gameMode string, _requirements string) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "createOrder", _totalAmount, _deadline, _gameType, _gameMode, _requirements)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xf62154ad.
//
// Solidity: function createOrder(uint256 _totalAmount, uint256 _deadline, string _gameType, string _gameMode, string _requirements) payable returns()
func (_BoostChainMainContract *BoostChainMainContractSession) CreateOrder(_totalAmount *big.Int, _deadline *big.Int, _gameType string, _gameMode string, _requirements string) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.CreateOrder(&_BoostChainMainContract.TransactOpts, _totalAmount, _deadline, _gameType, _gameMode, _requirements)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xf62154ad.
//
// Solidity: function createOrder(uint256 _totalAmount, uint256 _deadline, string _gameType, string _gameMode, string _requirements) payable returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) CreateOrder(_totalAmount *big.Int, _deadline *big.Int, _gameType string, _gameMode string, _requirements string) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.CreateOrder(&_BoostChainMainContract.TransactOpts, _totalAmount, _deadline, _gameType, _gameMode, _requirements)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_BoostChainMainContract *BoostChainMainContractSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.EmergencyWithdraw(&_BoostChainMainContract.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.EmergencyWithdraw(&_BoostChainMainContract.TransactOpts)
}

// FailOrder is a paid mutator transaction binding the contract method 0xfcc2b553.
//
// Solidity: function failOrder(uint256 _orderId) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) FailOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "failOrder", _orderId)
}

// FailOrder is a paid mutator transaction binding the contract method 0xfcc2b553.
//
// Solidity: function failOrder(uint256 _orderId) returns()
func (_BoostChainMainContract *BoostChainMainContractSession) FailOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.FailOrder(&_BoostChainMainContract.TransactOpts, _orderId)
}

// FailOrder is a paid mutator transaction binding the contract method 0xfcc2b553.
//
// Solidity: function failOrder(uint256 _orderId) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) FailOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.FailOrder(&_BoostChainMainContract.TransactOpts, _orderId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BoostChainMainContract *BoostChainMainContractSession) Pause() (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.Pause(&_BoostChainMainContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) Pause() (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.Pause(&_BoostChainMainContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BoostChainMainContract *BoostChainMainContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.RenounceOwnership(&_BoostChainMainContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.RenounceOwnership(&_BoostChainMainContract.TransactOpts)
}

// SetDepositRate is a paid mutator transaction binding the contract method 0xbf35588b.
//
// Solidity: function setDepositRate(uint256 _newRate) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) SetDepositRate(opts *bind.TransactOpts, _newRate *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "setDepositRate", _newRate)
}

// SetDepositRate is a paid mutator transaction binding the contract method 0xbf35588b.
//
// Solidity: function setDepositRate(uint256 _newRate) returns()
func (_BoostChainMainContract *BoostChainMainContractSession) SetDepositRate(_newRate *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.SetDepositRate(&_BoostChainMainContract.TransactOpts, _newRate)
}

// SetDepositRate is a paid mutator transaction binding the contract method 0xbf35588b.
//
// Solidity: function setDepositRate(uint256 _newRate) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) SetDepositRate(_newRate *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.SetDepositRate(&_BoostChainMainContract.TransactOpts, _newRate)
}

// SetPlatformFeeRate is a paid mutator transaction binding the contract method 0x927fef2e.
//
// Solidity: function setPlatformFeeRate(uint256 _newRate) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) SetPlatformFeeRate(opts *bind.TransactOpts, _newRate *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "setPlatformFeeRate", _newRate)
}

// SetPlatformFeeRate is a paid mutator transaction binding the contract method 0x927fef2e.
//
// Solidity: function setPlatformFeeRate(uint256 _newRate) returns()
func (_BoostChainMainContract *BoostChainMainContractSession) SetPlatformFeeRate(_newRate *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.SetPlatformFeeRate(&_BoostChainMainContract.TransactOpts, _newRate)
}

// SetPlatformFeeRate is a paid mutator transaction binding the contract method 0x927fef2e.
//
// Solidity: function setPlatformFeeRate(uint256 _newRate) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) SetPlatformFeeRate(_newRate *big.Int) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.SetPlatformFeeRate(&_BoostChainMainContract.TransactOpts, _newRate)
}

// SetPlatformTreasury is a paid mutator transaction binding the contract method 0x7cd86d60.
//
// Solidity: function setPlatformTreasury(address _newTreasury) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) SetPlatformTreasury(opts *bind.TransactOpts, _newTreasury common.Address) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "setPlatformTreasury", _newTreasury)
}

// SetPlatformTreasury is a paid mutator transaction binding the contract method 0x7cd86d60.
//
// Solidity: function setPlatformTreasury(address _newTreasury) returns()
func (_BoostChainMainContract *BoostChainMainContractSession) SetPlatformTreasury(_newTreasury common.Address) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.SetPlatformTreasury(&_BoostChainMainContract.TransactOpts, _newTreasury)
}

// SetPlatformTreasury is a paid mutator transaction binding the contract method 0x7cd86d60.
//
// Solidity: function setPlatformTreasury(address _newTreasury) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) SetPlatformTreasury(_newTreasury common.Address) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.SetPlatformTreasury(&_BoostChainMainContract.TransactOpts, _newTreasury)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BoostChainMainContract *BoostChainMainContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.TransferOwnership(&_BoostChainMainContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.TransferOwnership(&_BoostChainMainContract.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BoostChainMainContract *BoostChainMainContractSession) Unpause() (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.Unpause(&_BoostChainMainContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.Unpause(&_BoostChainMainContract.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BoostChainMainContract *BoostChainMainContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.Fallback(&_BoostChainMainContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.Fallback(&_BoostChainMainContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BoostChainMainContract *BoostChainMainContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoostChainMainContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BoostChainMainContract *BoostChainMainContractSession) Receive() (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.Receive(&_BoostChainMainContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BoostChainMainContract *BoostChainMainContractTransactorSession) Receive() (*types.Transaction, error) {
	return _BoostChainMainContract.Contract.Receive(&_BoostChainMainContract.TransactOpts)
}

// BoostChainMainContractOrderAcceptedIterator is returned from FilterOrderAccepted and is used to iterate over the raw logs and unpacked data for OrderAccepted events raised by the BoostChainMainContract contract.
type BoostChainMainContractOrderAcceptedIterator struct {
	Event *BoostChainMainContractOrderAccepted // Event containing the contract specifics and raw log

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
func (it *BoostChainMainContractOrderAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoostChainMainContractOrderAccepted)
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
		it.Event = new(BoostChainMainContractOrderAccepted)
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
func (it *BoostChainMainContractOrderAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoostChainMainContractOrderAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoostChainMainContractOrderAccepted represents a OrderAccepted event raised by the BoostChainMainContract contract.
type BoostChainMainContractOrderAccepted struct {
	OrderId        *big.Int
	Booster        common.Address
	BoosterDeposit *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOrderAccepted is a free log retrieval operation binding the contract event 0x5e7ed488618ff5431a6372dfbaf675f24c13c93842f4e2d54d85e4dc8ffb8ef5.
//
// Solidity: event OrderAccepted(uint256 indexed orderId, address indexed booster, uint256 boosterDeposit)
func (_BoostChainMainContract *BoostChainMainContractFilterer) FilterOrderAccepted(opts *bind.FilterOpts, orderId []*big.Int, booster []common.Address) (*BoostChainMainContractOrderAcceptedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var boosterRule []interface{}
	for _, boosterItem := range booster {
		boosterRule = append(boosterRule, boosterItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.FilterLogs(opts, "OrderAccepted", orderIdRule, boosterRule)
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContractOrderAcceptedIterator{contract: _BoostChainMainContract.contract, event: "OrderAccepted", logs: logs, sub: sub}, nil
}

// WatchOrderAccepted is a free log subscription operation binding the contract event 0x5e7ed488618ff5431a6372dfbaf675f24c13c93842f4e2d54d85e4dc8ffb8ef5.
//
// Solidity: event OrderAccepted(uint256 indexed orderId, address indexed booster, uint256 boosterDeposit)
func (_BoostChainMainContract *BoostChainMainContractFilterer) WatchOrderAccepted(opts *bind.WatchOpts, sink chan<- *BoostChainMainContractOrderAccepted, orderId []*big.Int, booster []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var boosterRule []interface{}
	for _, boosterItem := range booster {
		boosterRule = append(boosterRule, boosterItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.WatchLogs(opts, "OrderAccepted", orderIdRule, boosterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoostChainMainContractOrderAccepted)
				if err := _BoostChainMainContract.contract.UnpackLog(event, "OrderAccepted", log); err != nil {
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

// ParseOrderAccepted is a log parse operation binding the contract event 0x5e7ed488618ff5431a6372dfbaf675f24c13c93842f4e2d54d85e4dc8ffb8ef5.
//
// Solidity: event OrderAccepted(uint256 indexed orderId, address indexed booster, uint256 boosterDeposit)
func (_BoostChainMainContract *BoostChainMainContractFilterer) ParseOrderAccepted(log types.Log) (*BoostChainMainContractOrderAccepted, error) {
	event := new(BoostChainMainContractOrderAccepted)
	if err := _BoostChainMainContract.contract.UnpackLog(event, "OrderAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoostChainMainContractOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the BoostChainMainContract contract.
type BoostChainMainContractOrderCancelledIterator struct {
	Event *BoostChainMainContractOrderCancelled // Event containing the contract specifics and raw log

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
func (it *BoostChainMainContractOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoostChainMainContractOrderCancelled)
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
		it.Event = new(BoostChainMainContractOrderCancelled)
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
func (it *BoostChainMainContractOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoostChainMainContractOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoostChainMainContractOrderCancelled represents a OrderCancelled event raised by the BoostChainMainContract contract.
type BoostChainMainContractOrderCancelled struct {
	OrderId         *big.Int
	CancelledBy     common.Address
	PenaltyAmount   *big.Int
	PenaltyReceiver common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xcf137cc8ba486b28681a39807956960712339eb42ea3229202cfd9d5087b4e46.
//
// Solidity: event OrderCancelled(uint256 indexed orderId, address indexed cancelledBy, uint256 penaltyAmount, address penaltyReceiver)
func (_BoostChainMainContract *BoostChainMainContractFilterer) FilterOrderCancelled(opts *bind.FilterOpts, orderId []*big.Int, cancelledBy []common.Address) (*BoostChainMainContractOrderCancelledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var cancelledByRule []interface{}
	for _, cancelledByItem := range cancelledBy {
		cancelledByRule = append(cancelledByRule, cancelledByItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.FilterLogs(opts, "OrderCancelled", orderIdRule, cancelledByRule)
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContractOrderCancelledIterator{contract: _BoostChainMainContract.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xcf137cc8ba486b28681a39807956960712339eb42ea3229202cfd9d5087b4e46.
//
// Solidity: event OrderCancelled(uint256 indexed orderId, address indexed cancelledBy, uint256 penaltyAmount, address penaltyReceiver)
func (_BoostChainMainContract *BoostChainMainContractFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *BoostChainMainContractOrderCancelled, orderId []*big.Int, cancelledBy []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var cancelledByRule []interface{}
	for _, cancelledByItem := range cancelledBy {
		cancelledByRule = append(cancelledByRule, cancelledByItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.WatchLogs(opts, "OrderCancelled", orderIdRule, cancelledByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoostChainMainContractOrderCancelled)
				if err := _BoostChainMainContract.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0xcf137cc8ba486b28681a39807956960712339eb42ea3229202cfd9d5087b4e46.
//
// Solidity: event OrderCancelled(uint256 indexed orderId, address indexed cancelledBy, uint256 penaltyAmount, address penaltyReceiver)
func (_BoostChainMainContract *BoostChainMainContractFilterer) ParseOrderCancelled(log types.Log) (*BoostChainMainContractOrderCancelled, error) {
	event := new(BoostChainMainContractOrderCancelled)
	if err := _BoostChainMainContract.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoostChainMainContractOrderCompletedIterator is returned from FilterOrderCompleted and is used to iterate over the raw logs and unpacked data for OrderCompleted events raised by the BoostChainMainContract contract.
type BoostChainMainContractOrderCompletedIterator struct {
	Event *BoostChainMainContractOrderCompleted // Event containing the contract specifics and raw log

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
func (it *BoostChainMainContractOrderCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoostChainMainContractOrderCompleted)
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
		it.Event = new(BoostChainMainContractOrderCompleted)
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
func (it *BoostChainMainContractOrderCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoostChainMainContractOrderCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoostChainMainContractOrderCompleted represents a OrderCompleted event raised by the BoostChainMainContract contract.
type BoostChainMainContractOrderCompleted struct {
	OrderId       *big.Int
	PlatformFee   *big.Int
	BoosterReward *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderCompleted is a free log retrieval operation binding the contract event 0xc4bfc30eeb4d37af0bfe8327e0acfdc5e10da19745a4677dd5f4244993d7ca98.
//
// Solidity: event OrderCompleted(uint256 indexed orderId, uint256 platformFee, uint256 boosterReward)
func (_BoostChainMainContract *BoostChainMainContractFilterer) FilterOrderCompleted(opts *bind.FilterOpts, orderId []*big.Int) (*BoostChainMainContractOrderCompletedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.FilterLogs(opts, "OrderCompleted", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContractOrderCompletedIterator{contract: _BoostChainMainContract.contract, event: "OrderCompleted", logs: logs, sub: sub}, nil
}

// WatchOrderCompleted is a free log subscription operation binding the contract event 0xc4bfc30eeb4d37af0bfe8327e0acfdc5e10da19745a4677dd5f4244993d7ca98.
//
// Solidity: event OrderCompleted(uint256 indexed orderId, uint256 platformFee, uint256 boosterReward)
func (_BoostChainMainContract *BoostChainMainContractFilterer) WatchOrderCompleted(opts *bind.WatchOpts, sink chan<- *BoostChainMainContractOrderCompleted, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.WatchLogs(opts, "OrderCompleted", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoostChainMainContractOrderCompleted)
				if err := _BoostChainMainContract.contract.UnpackLog(event, "OrderCompleted", log); err != nil {
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

// ParseOrderCompleted is a log parse operation binding the contract event 0xc4bfc30eeb4d37af0bfe8327e0acfdc5e10da19745a4677dd5f4244993d7ca98.
//
// Solidity: event OrderCompleted(uint256 indexed orderId, uint256 platformFee, uint256 boosterReward)
func (_BoostChainMainContract *BoostChainMainContractFilterer) ParseOrderCompleted(log types.Log) (*BoostChainMainContractOrderCompleted, error) {
	event := new(BoostChainMainContractOrderCompleted)
	if err := _BoostChainMainContract.contract.UnpackLog(event, "OrderCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoostChainMainContractOrderConfirmedIterator is returned from FilterOrderConfirmed and is used to iterate over the raw logs and unpacked data for OrderConfirmed events raised by the BoostChainMainContract contract.
type BoostChainMainContractOrderConfirmedIterator struct {
	Event *BoostChainMainContractOrderConfirmed // Event containing the contract specifics and raw log

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
func (it *BoostChainMainContractOrderConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoostChainMainContractOrderConfirmed)
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
		it.Event = new(BoostChainMainContractOrderConfirmed)
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
func (it *BoostChainMainContractOrderConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoostChainMainContractOrderConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoostChainMainContractOrderConfirmed represents a OrderConfirmed event raised by the BoostChainMainContract contract.
type BoostChainMainContractOrderConfirmed struct {
	OrderId           *big.Int
	ChildContract     common.Address
	TotalAmountLocked *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrderConfirmed is a free log retrieval operation binding the contract event 0xd9153bed9da362a771596f693014f5226ba4b890b57155e405c1d3b58b180aaf.
//
// Solidity: event OrderConfirmed(uint256 indexed orderId, address indexed childContract, uint256 totalAmountLocked)
func (_BoostChainMainContract *BoostChainMainContractFilterer) FilterOrderConfirmed(opts *bind.FilterOpts, orderId []*big.Int, childContract []common.Address) (*BoostChainMainContractOrderConfirmedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var childContractRule []interface{}
	for _, childContractItem := range childContract {
		childContractRule = append(childContractRule, childContractItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.FilterLogs(opts, "OrderConfirmed", orderIdRule, childContractRule)
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContractOrderConfirmedIterator{contract: _BoostChainMainContract.contract, event: "OrderConfirmed", logs: logs, sub: sub}, nil
}

// WatchOrderConfirmed is a free log subscription operation binding the contract event 0xd9153bed9da362a771596f693014f5226ba4b890b57155e405c1d3b58b180aaf.
//
// Solidity: event OrderConfirmed(uint256 indexed orderId, address indexed childContract, uint256 totalAmountLocked)
func (_BoostChainMainContract *BoostChainMainContractFilterer) WatchOrderConfirmed(opts *bind.WatchOpts, sink chan<- *BoostChainMainContractOrderConfirmed, orderId []*big.Int, childContract []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var childContractRule []interface{}
	for _, childContractItem := range childContract {
		childContractRule = append(childContractRule, childContractItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.WatchLogs(opts, "OrderConfirmed", orderIdRule, childContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoostChainMainContractOrderConfirmed)
				if err := _BoostChainMainContract.contract.UnpackLog(event, "OrderConfirmed", log); err != nil {
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

// ParseOrderConfirmed is a log parse operation binding the contract event 0xd9153bed9da362a771596f693014f5226ba4b890b57155e405c1d3b58b180aaf.
//
// Solidity: event OrderConfirmed(uint256 indexed orderId, address indexed childContract, uint256 totalAmountLocked)
func (_BoostChainMainContract *BoostChainMainContractFilterer) ParseOrderConfirmed(log types.Log) (*BoostChainMainContractOrderConfirmed, error) {
	event := new(BoostChainMainContractOrderConfirmed)
	if err := _BoostChainMainContract.contract.UnpackLog(event, "OrderConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoostChainMainContractOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the BoostChainMainContract contract.
type BoostChainMainContractOrderCreatedIterator struct {
	Event *BoostChainMainContractOrderCreated // Event containing the contract specifics and raw log

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
func (it *BoostChainMainContractOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoostChainMainContractOrderCreated)
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
		it.Event = new(BoostChainMainContractOrderCreated)
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
func (it *BoostChainMainContractOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoostChainMainContractOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoostChainMainContractOrderCreated represents a OrderCreated event raised by the BoostChainMainContract contract.
type BoostChainMainContractOrderCreated struct {
	OrderId       *big.Int
	Player        common.Address
	TotalAmount   *big.Int
	PlayerDeposit *big.Int
	GameType      string
	GameMode      string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x6ffbef5b13a3765327808a6fec6c0e80175b6ea16168958142444975916fe865.
//
// Solidity: event OrderCreated(uint256 indexed orderId, address indexed player, uint256 totalAmount, uint256 playerDeposit, string gameType, string gameMode)
func (_BoostChainMainContract *BoostChainMainContractFilterer) FilterOrderCreated(opts *bind.FilterOpts, orderId []*big.Int, player []common.Address) (*BoostChainMainContractOrderCreatedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.FilterLogs(opts, "OrderCreated", orderIdRule, playerRule)
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContractOrderCreatedIterator{contract: _BoostChainMainContract.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x6ffbef5b13a3765327808a6fec6c0e80175b6ea16168958142444975916fe865.
//
// Solidity: event OrderCreated(uint256 indexed orderId, address indexed player, uint256 totalAmount, uint256 playerDeposit, string gameType, string gameMode)
func (_BoostChainMainContract *BoostChainMainContractFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *BoostChainMainContractOrderCreated, orderId []*big.Int, player []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.WatchLogs(opts, "OrderCreated", orderIdRule, playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoostChainMainContractOrderCreated)
				if err := _BoostChainMainContract.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x6ffbef5b13a3765327808a6fec6c0e80175b6ea16168958142444975916fe865.
//
// Solidity: event OrderCreated(uint256 indexed orderId, address indexed player, uint256 totalAmount, uint256 playerDeposit, string gameType, string gameMode)
func (_BoostChainMainContract *BoostChainMainContractFilterer) ParseOrderCreated(log types.Log) (*BoostChainMainContractOrderCreated, error) {
	event := new(BoostChainMainContractOrderCreated)
	if err := _BoostChainMainContract.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoostChainMainContractOrderFailedIterator is returned from FilterOrderFailed and is used to iterate over the raw logs and unpacked data for OrderFailed events raised by the BoostChainMainContract contract.
type BoostChainMainContractOrderFailedIterator struct {
	Event *BoostChainMainContractOrderFailed // Event containing the contract specifics and raw log

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
func (it *BoostChainMainContractOrderFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoostChainMainContractOrderFailed)
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
		it.Event = new(BoostChainMainContractOrderFailed)
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
func (it *BoostChainMainContractOrderFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoostChainMainContractOrderFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoostChainMainContractOrderFailed represents a OrderFailed event raised by the BoostChainMainContract contract.
type BoostChainMainContractOrderFailed struct {
	OrderId           *big.Int
	PlayerRefund      *big.Int
	PenaltyToPlayer   *big.Int
	PenaltyToPlatform *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrderFailed is a free log retrieval operation binding the contract event 0xd523100b3ba1ea2dbee908e8a7ece3c17ca1048deb23df6dcdf3618499ddcf1a.
//
// Solidity: event OrderFailed(uint256 indexed orderId, uint256 playerRefund, uint256 penaltyToPlayer, uint256 penaltyToPlatform)
func (_BoostChainMainContract *BoostChainMainContractFilterer) FilterOrderFailed(opts *bind.FilterOpts, orderId []*big.Int) (*BoostChainMainContractOrderFailedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.FilterLogs(opts, "OrderFailed", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContractOrderFailedIterator{contract: _BoostChainMainContract.contract, event: "OrderFailed", logs: logs, sub: sub}, nil
}

// WatchOrderFailed is a free log subscription operation binding the contract event 0xd523100b3ba1ea2dbee908e8a7ece3c17ca1048deb23df6dcdf3618499ddcf1a.
//
// Solidity: event OrderFailed(uint256 indexed orderId, uint256 playerRefund, uint256 penaltyToPlayer, uint256 penaltyToPlatform)
func (_BoostChainMainContract *BoostChainMainContractFilterer) WatchOrderFailed(opts *bind.WatchOpts, sink chan<- *BoostChainMainContractOrderFailed, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.WatchLogs(opts, "OrderFailed", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoostChainMainContractOrderFailed)
				if err := _BoostChainMainContract.contract.UnpackLog(event, "OrderFailed", log); err != nil {
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

// ParseOrderFailed is a log parse operation binding the contract event 0xd523100b3ba1ea2dbee908e8a7ece3c17ca1048deb23df6dcdf3618499ddcf1a.
//
// Solidity: event OrderFailed(uint256 indexed orderId, uint256 playerRefund, uint256 penaltyToPlayer, uint256 penaltyToPlatform)
func (_BoostChainMainContract *BoostChainMainContractFilterer) ParseOrderFailed(log types.Log) (*BoostChainMainContractOrderFailed, error) {
	event := new(BoostChainMainContractOrderFailed)
	if err := _BoostChainMainContract.contract.UnpackLog(event, "OrderFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoostChainMainContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BoostChainMainContract contract.
type BoostChainMainContractOwnershipTransferredIterator struct {
	Event *BoostChainMainContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BoostChainMainContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoostChainMainContractOwnershipTransferred)
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
		it.Event = new(BoostChainMainContractOwnershipTransferred)
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
func (it *BoostChainMainContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoostChainMainContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoostChainMainContractOwnershipTransferred represents a OwnershipTransferred event raised by the BoostChainMainContract contract.
type BoostChainMainContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BoostChainMainContract *BoostChainMainContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BoostChainMainContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContractOwnershipTransferredIterator{contract: _BoostChainMainContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BoostChainMainContract *BoostChainMainContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BoostChainMainContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BoostChainMainContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoostChainMainContractOwnershipTransferred)
				if err := _BoostChainMainContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BoostChainMainContract *BoostChainMainContractFilterer) ParseOwnershipTransferred(log types.Log) (*BoostChainMainContractOwnershipTransferred, error) {
	event := new(BoostChainMainContractOwnershipTransferred)
	if err := _BoostChainMainContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoostChainMainContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BoostChainMainContract contract.
type BoostChainMainContractPausedIterator struct {
	Event *BoostChainMainContractPaused // Event containing the contract specifics and raw log

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
func (it *BoostChainMainContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoostChainMainContractPaused)
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
		it.Event = new(BoostChainMainContractPaused)
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
func (it *BoostChainMainContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoostChainMainContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoostChainMainContractPaused represents a Paused event raised by the BoostChainMainContract contract.
type BoostChainMainContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BoostChainMainContract *BoostChainMainContractFilterer) FilterPaused(opts *bind.FilterOpts) (*BoostChainMainContractPausedIterator, error) {

	logs, sub, err := _BoostChainMainContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContractPausedIterator{contract: _BoostChainMainContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BoostChainMainContract *BoostChainMainContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BoostChainMainContractPaused) (event.Subscription, error) {

	logs, sub, err := _BoostChainMainContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoostChainMainContractPaused)
				if err := _BoostChainMainContract.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BoostChainMainContract *BoostChainMainContractFilterer) ParsePaused(log types.Log) (*BoostChainMainContractPaused, error) {
	event := new(BoostChainMainContractPaused)
	if err := _BoostChainMainContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoostChainMainContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BoostChainMainContract contract.
type BoostChainMainContractUnpausedIterator struct {
	Event *BoostChainMainContractUnpaused // Event containing the contract specifics and raw log

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
func (it *BoostChainMainContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoostChainMainContractUnpaused)
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
		it.Event = new(BoostChainMainContractUnpaused)
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
func (it *BoostChainMainContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoostChainMainContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoostChainMainContractUnpaused represents a Unpaused event raised by the BoostChainMainContract contract.
type BoostChainMainContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BoostChainMainContract *BoostChainMainContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BoostChainMainContractUnpausedIterator, error) {

	logs, sub, err := _BoostChainMainContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BoostChainMainContractUnpausedIterator{contract: _BoostChainMainContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BoostChainMainContract *BoostChainMainContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BoostChainMainContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _BoostChainMainContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoostChainMainContractUnpaused)
				if err := _BoostChainMainContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BoostChainMainContract *BoostChainMainContractFilterer) ParseUnpaused(log types.Log) (*BoostChainMainContractUnpaused, error) {
	event := new(BoostChainMainContractUnpaused)
	if err := _BoostChainMainContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
