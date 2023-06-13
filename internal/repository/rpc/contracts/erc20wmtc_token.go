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
)

// ErcWrappedMtcMetaData contains all meta data concerning the ErcWrappedMtc contract.
var ErcWrappedMtcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_INVALID_ZERO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_ERROR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ErcWrappedMtcABI is the input ABI used to generate the binding from.
// Deprecated: Use ErcWrappedMtcMetaData.ABI instead.
var ErcWrappedMtcABI = ErcWrappedMtcMetaData.ABI

// ErcWrappedMtc is an auto generated Go binding around an Ethereum contract.
type ErcWrappedMtc struct {
	ErcWrappedMtcCaller     // Read-only binding to the contract
	ErcWrappedMtcTransactor // Write-only binding to the contract
	ErcWrappedMtcFilterer   // Log filterer for contract events
}

// ErcWrappedMtcCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErcWrappedMtcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcWrappedMtcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErcWrappedMtcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcWrappedMtcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErcWrappedMtcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcWrappedMtcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErcWrappedMtcSession struct {
	Contract     *ErcWrappedMtc    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErcWrappedMtcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErcWrappedMtcCallerSession struct {
	Contract *ErcWrappedMtcCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ErcWrappedMtcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErcWrappedMtcTransactorSession struct {
	Contract     *ErcWrappedMtcTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ErcWrappedMtcRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErcWrappedMtcRaw struct {
	Contract *ErcWrappedMtc // Generic contract binding to access the raw methods on
}

// ErcWrappedMtcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErcWrappedMtcCallerRaw struct {
	Contract *ErcWrappedMtcCaller // Generic read-only contract binding to access the raw methods on
}

// ErcWrappedMtcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErcWrappedMtcTransactorRaw struct {
	Contract *ErcWrappedMtcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErcWrappedMtc creates a new instance of ErcWrappedMtc, bound to a specific deployed contract.
func NewErcWrappedMtc(address common.Address, backend bind.ContractBackend) (*ErcWrappedMtc, error) {
	contract, err := bindErcWrappedMtc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedMtc{ErcWrappedMtcCaller: ErcWrappedMtcCaller{contract: contract}, ErcWrappedMtcTransactor: ErcWrappedMtcTransactor{contract: contract}, ErcWrappedMtcFilterer: ErcWrappedMtcFilterer{contract: contract}}, nil
}

// NewErcWrappedMtcCaller creates a new read-only instance of ErcWrappedMtc, bound to a specific deployed contract.
func NewErcWrappedMtcCaller(address common.Address, caller bind.ContractCaller) (*ErcWrappedMtcCaller, error) {
	contract, err := bindErcWrappedMtc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedMtcCaller{contract: contract}, nil
}

// NewErcWrappedMtcTransactor creates a new write-only instance of ErcWrappedMtc, bound to a specific deployed contract.
func NewErcWrappedMtcTransactor(address common.Address, transactor bind.ContractTransactor) (*ErcWrappedMtcTransactor, error) {
	contract, err := bindErcWrappedMtc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedMtcTransactor{contract: contract}, nil
}

// NewErcWrappedMtcFilterer creates a new log filterer instance of ErcWrappedMtc, bound to a specific deployed contract.
func NewErcWrappedMtcFilterer(address common.Address, filterer bind.ContractFilterer) (*ErcWrappedMtcFilterer, error) {
	contract, err := bindErcWrappedMtc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedMtcFilterer{contract: contract}, nil
}

// bindErcWrappedMtc binds a generic wrapper to an already deployed contract.
func bindErcWrappedMtc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ErcWrappedMtcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ErcWrappedMtc *ErcWrappedMtcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ErcWrappedMtc.Contract.ErcWrappedMtcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ErcWrappedMtc *ErcWrappedMtcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.ErcWrappedMtcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ErcWrappedMtc *ErcWrappedMtcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.ErcWrappedMtcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ErcWrappedMtc *ErcWrappedMtcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ErcWrappedMtc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ErcWrappedMtc *ErcWrappedMtcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ErcWrappedMtc *ErcWrappedMtcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.contract.Transact(opts, method, params...)
}

// ERRINVALIDZEROVALUE is a free data retrieval call binding the contract method 0x6d7497b3.
//
// Solidity: function ERR_INVALID_ZERO_VALUE() view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcCaller) ERRINVALIDZEROVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedMtc.contract.Call(opts, &out, "ERR_INVALID_ZERO_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRINVALIDZEROVALUE is a free data retrieval call binding the contract method 0x6d7497b3.
//
// Solidity: function ERR_INVALID_ZERO_VALUE() view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcSession) ERRINVALIDZEROVALUE() (*big.Int, error) {
	return _ErcWrappedMtc.Contract.ERRINVALIDZEROVALUE(&_ErcWrappedMtc.CallOpts)
}

// ERRINVALIDZEROVALUE is a free data retrieval call binding the contract method 0x6d7497b3.
//
// Solidity: function ERR_INVALID_ZERO_VALUE() view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcCallerSession) ERRINVALIDZEROVALUE() (*big.Int, error) {
	return _ErcWrappedMtc.Contract.ERRINVALIDZEROVALUE(&_ErcWrappedMtc.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcCaller) ERRNOERROR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedMtc.contract.Call(opts, &out, "ERR_NO_ERROR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcSession) ERRNOERROR() (*big.Int, error) {
	return _ErcWrappedMtc.Contract.ERRNOERROR(&_ErcWrappedMtc.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcCallerSession) ERRNOERROR() (*big.Int, error) {
	return _ErcWrappedMtc.Contract.ERRNOERROR(&_ErcWrappedMtc.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedMtc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ErcWrappedMtc.Contract.Allowance(&_ErcWrappedMtc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ErcWrappedMtc.Contract.Allowance(&_ErcWrappedMtc.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedMtc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ErcWrappedMtc.Contract.BalanceOf(&_ErcWrappedMtc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ErcWrappedMtc.Contract.BalanceOf(&_ErcWrappedMtc.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ErcWrappedMtc *ErcWrappedMtcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ErcWrappedMtc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ErcWrappedMtc *ErcWrappedMtcSession) Decimals() (uint8, error) {
	return _ErcWrappedMtc.Contract.Decimals(&_ErcWrappedMtc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ErcWrappedMtc *ErcWrappedMtcCallerSession) Decimals() (uint8, error) {
	return _ErcWrappedMtc.Contract.Decimals(&_ErcWrappedMtc.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ErcWrappedMtc.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcSession) IsPauser(account common.Address) (bool, error) {
	return _ErcWrappedMtc.Contract.IsPauser(&_ErcWrappedMtc.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcCallerSession) IsPauser(account common.Address) (bool, error) {
	return _ErcWrappedMtc.Contract.IsPauser(&_ErcWrappedMtc.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ErcWrappedMtc *ErcWrappedMtcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ErcWrappedMtc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ErcWrappedMtc *ErcWrappedMtcSession) Name() (string, error) {
	return _ErcWrappedMtc.Contract.Name(&_ErcWrappedMtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ErcWrappedMtc *ErcWrappedMtcCallerSession) Name() (string, error) {
	return _ErcWrappedMtc.Contract.Name(&_ErcWrappedMtc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ErcWrappedMtc.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcSession) Paused() (bool, error) {
	return _ErcWrappedMtc.Contract.Paused(&_ErcWrappedMtc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcCallerSession) Paused() (bool, error) {
	return _ErcWrappedMtc.Contract.Paused(&_ErcWrappedMtc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ErcWrappedMtc *ErcWrappedMtcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ErcWrappedMtc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ErcWrappedMtc *ErcWrappedMtcSession) Symbol() (string, error) {
	return _ErcWrappedMtc.Contract.Symbol(&_ErcWrappedMtc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ErcWrappedMtc *ErcWrappedMtcCallerSession) Symbol() (string, error) {
	return _ErcWrappedMtc.Contract.Symbol(&_ErcWrappedMtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedMtc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcSession) TotalSupply() (*big.Int, error) {
	return _ErcWrappedMtc.Contract.TotalSupply(&_ErcWrappedMtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcCallerSession) TotalSupply() (*big.Int, error) {
	return _ErcWrappedMtc.Contract.TotalSupply(&_ErcWrappedMtc.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ErcWrappedMtc *ErcWrappedMtcTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ErcWrappedMtc.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ErcWrappedMtc *ErcWrappedMtcSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.AddPauser(&_ErcWrappedMtc.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ErcWrappedMtc *ErcWrappedMtcTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.AddPauser(&_ErcWrappedMtc.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.Approve(&_ErcWrappedMtc.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.Approve(&_ErcWrappedMtc.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.DecreaseAllowance(&_ErcWrappedMtc.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.DecreaseAllowance(&_ErcWrappedMtc.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedMtc.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcSession) Deposit() (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.Deposit(&_ErcWrappedMtc.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcTransactorSession) Deposit() (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.Deposit(&_ErcWrappedMtc.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.IncreaseAllowance(&_ErcWrappedMtc.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.IncreaseAllowance(&_ErcWrappedMtc.TransactOpts, spender, addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ErcWrappedMtc *ErcWrappedMtcTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedMtc.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ErcWrappedMtc *ErcWrappedMtcSession) Pause() (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.Pause(&_ErcWrappedMtc.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ErcWrappedMtc *ErcWrappedMtcTransactorSession) Pause() (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.Pause(&_ErcWrappedMtc.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ErcWrappedMtc *ErcWrappedMtcTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedMtc.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ErcWrappedMtc *ErcWrappedMtcSession) RenouncePauser() (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.RenouncePauser(&_ErcWrappedMtc.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ErcWrappedMtc *ErcWrappedMtcTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.RenouncePauser(&_ErcWrappedMtc.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.Transfer(&_ErcWrappedMtc.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.Transfer(&_ErcWrappedMtc.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.TransferFrom(&_ErcWrappedMtc.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ErcWrappedMtc *ErcWrappedMtcTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.TransferFrom(&_ErcWrappedMtc.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ErcWrappedMtc *ErcWrappedMtcTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedMtc.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ErcWrappedMtc *ErcWrappedMtcSession) Unpause() (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.Unpause(&_ErcWrappedMtc.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ErcWrappedMtc *ErcWrappedMtcTransactorSession) Unpause() (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.Unpause(&_ErcWrappedMtc.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.Withdraw(&_ErcWrappedMtc.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_ErcWrappedMtc *ErcWrappedMtcTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ErcWrappedMtc.Contract.Withdraw(&_ErcWrappedMtc.TransactOpts, amount)
}

// ErcWrappedMtcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ErcWrappedMtc contract.
type ErcWrappedMtcApprovalIterator struct {
	Event *ErcWrappedMtcApproval // Event containing the contract specifics and raw log

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
func (it *ErcWrappedMtcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedMtcApproval)
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
		it.Event = new(ErcWrappedMtcApproval)
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
func (it *ErcWrappedMtcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedMtcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedMtcApproval represents a Approval event raised by the ErcWrappedMtc contract.
type ErcWrappedMtcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ErcWrappedMtcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ErcWrappedMtc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedMtcApprovalIterator{contract: _ErcWrappedMtc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ErcWrappedMtcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ErcWrappedMtc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedMtcApproval)
				if err := _ErcWrappedMtc.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) ParseApproval(log types.Log) (*ErcWrappedMtcApproval, error) {
	event := new(ErcWrappedMtcApproval)
	if err := _ErcWrappedMtc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedMtcPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ErcWrappedMtc contract.
type ErcWrappedMtcPausedIterator struct {
	Event *ErcWrappedMtcPaused // Event containing the contract specifics and raw log

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
func (it *ErcWrappedMtcPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedMtcPaused)
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
		it.Event = new(ErcWrappedMtcPaused)
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
func (it *ErcWrappedMtcPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedMtcPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedMtcPaused represents a Paused event raised by the ErcWrappedMtc contract.
type ErcWrappedMtcPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) FilterPaused(opts *bind.FilterOpts) (*ErcWrappedMtcPausedIterator, error) {

	logs, sub, err := _ErcWrappedMtc.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ErcWrappedMtcPausedIterator{contract: _ErcWrappedMtc.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ErcWrappedMtcPaused) (event.Subscription, error) {

	logs, sub, err := _ErcWrappedMtc.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedMtcPaused)
				if err := _ErcWrappedMtc.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) ParsePaused(log types.Log) (*ErcWrappedMtcPaused, error) {
	event := new(ErcWrappedMtcPaused)
	if err := _ErcWrappedMtc.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedMtcPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the ErcWrappedMtc contract.
type ErcWrappedMtcPauserAddedIterator struct {
	Event *ErcWrappedMtcPauserAdded // Event containing the contract specifics and raw log

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
func (it *ErcWrappedMtcPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedMtcPauserAdded)
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
		it.Event = new(ErcWrappedMtcPauserAdded)
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
func (it *ErcWrappedMtcPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedMtcPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedMtcPauserAdded represents a PauserAdded event raised by the ErcWrappedMtc contract.
type ErcWrappedMtcPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*ErcWrappedMtcPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedMtc.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedMtcPauserAddedIterator{contract: _ErcWrappedMtc.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *ErcWrappedMtcPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedMtc.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedMtcPauserAdded)
				if err := _ErcWrappedMtc.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) ParsePauserAdded(log types.Log) (*ErcWrappedMtcPauserAdded, error) {
	event := new(ErcWrappedMtcPauserAdded)
	if err := _ErcWrappedMtc.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedMtcPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the ErcWrappedMtc contract.
type ErcWrappedMtcPauserRemovedIterator struct {
	Event *ErcWrappedMtcPauserRemoved // Event containing the contract specifics and raw log

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
func (it *ErcWrappedMtcPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedMtcPauserRemoved)
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
		it.Event = new(ErcWrappedMtcPauserRemoved)
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
func (it *ErcWrappedMtcPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedMtcPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedMtcPauserRemoved represents a PauserRemoved event raised by the ErcWrappedMtc contract.
type ErcWrappedMtcPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*ErcWrappedMtcPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedMtc.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedMtcPauserRemovedIterator{contract: _ErcWrappedMtc.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *ErcWrappedMtcPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedMtc.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedMtcPauserRemoved)
				if err := _ErcWrappedMtc.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) ParsePauserRemoved(log types.Log) (*ErcWrappedMtcPauserRemoved, error) {
	event := new(ErcWrappedMtcPauserRemoved)
	if err := _ErcWrappedMtc.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedMtcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ErcWrappedMtc contract.
type ErcWrappedMtcTransferIterator struct {
	Event *ErcWrappedMtcTransfer // Event containing the contract specifics and raw log

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
func (it *ErcWrappedMtcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedMtcTransfer)
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
		it.Event = new(ErcWrappedMtcTransfer)
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
func (it *ErcWrappedMtcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedMtcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedMtcTransfer represents a Transfer event raised by the ErcWrappedMtc contract.
type ErcWrappedMtcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ErcWrappedMtcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ErcWrappedMtc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedMtcTransferIterator{contract: _ErcWrappedMtc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ErcWrappedMtcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ErcWrappedMtc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedMtcTransfer)
				if err := _ErcWrappedMtc.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) ParseTransfer(log types.Log) (*ErcWrappedMtcTransfer, error) {
	event := new(ErcWrappedMtcTransfer)
	if err := _ErcWrappedMtc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedMtcUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ErcWrappedMtc contract.
type ErcWrappedMtcUnpausedIterator struct {
	Event *ErcWrappedMtcUnpaused // Event containing the contract specifics and raw log

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
func (it *ErcWrappedMtcUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedMtcUnpaused)
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
		it.Event = new(ErcWrappedMtcUnpaused)
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
func (it *ErcWrappedMtcUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedMtcUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedMtcUnpaused represents a Unpaused event raised by the ErcWrappedMtc contract.
type ErcWrappedMtcUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ErcWrappedMtcUnpausedIterator, error) {

	logs, sub, err := _ErcWrappedMtc.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ErcWrappedMtcUnpausedIterator{contract: _ErcWrappedMtc.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ErcWrappedMtcUnpaused) (event.Subscription, error) {

	logs, sub, err := _ErcWrappedMtc.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedMtcUnpaused)
				if err := _ErcWrappedMtc.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ErcWrappedMtc *ErcWrappedMtcFilterer) ParseUnpaused(log types.Log) (*ErcWrappedMtcUnpaused, error) {
	event := new(ErcWrappedMtcUnpaused)
	if err := _ErcWrappedMtc.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
