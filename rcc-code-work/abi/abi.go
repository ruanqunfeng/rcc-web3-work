// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// AbiMetaData contains all meta data concerning the Abi contract.
var AbiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"computeSelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"bs\",\"type\":\"bytes\"}],\"name\":\"decodeData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"encodeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"encodeFunctionCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"hashFunctions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"mathTest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"recoverAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// AbiABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiMetaData.ABI instead.
var AbiABI = AbiMetaData.ABI

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// ComputeSelector is a free data retrieval call binding the contract method 0x68bc2ad1.
//
// Solidity: function computeSelector(string str) pure returns(bytes4)
func (_Abi *AbiCaller) ComputeSelector(opts *bind.CallOpts, str string) ([4]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "computeSelector", str)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ComputeSelector is a free data retrieval call binding the contract method 0x68bc2ad1.
//
// Solidity: function computeSelector(string str) pure returns(bytes4)
func (_Abi *AbiSession) ComputeSelector(str string) ([4]byte, error) {
	return _Abi.Contract.ComputeSelector(&_Abi.CallOpts, str)
}

// ComputeSelector is a free data retrieval call binding the contract method 0x68bc2ad1.
//
// Solidity: function computeSelector(string str) pure returns(bytes4)
func (_Abi *AbiCallerSession) ComputeSelector(str string) ([4]byte, error) {
	return _Abi.Contract.ComputeSelector(&_Abi.CallOpts, str)
}

// DecodeData is a free data retrieval call binding the contract method 0x447a06f5.
//
// Solidity: function decodeData(bytes bs) pure returns(string, uint256)
func (_Abi *AbiCaller) DecodeData(opts *bind.CallOpts, bs []byte) (string, *big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "decodeData", bs)

	if err != nil {
		return *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// DecodeData is a free data retrieval call binding the contract method 0x447a06f5.
//
// Solidity: function decodeData(bytes bs) pure returns(string, uint256)
func (_Abi *AbiSession) DecodeData(bs []byte) (string, *big.Int, error) {
	return _Abi.Contract.DecodeData(&_Abi.CallOpts, bs)
}

// DecodeData is a free data retrieval call binding the contract method 0x447a06f5.
//
// Solidity: function decodeData(bytes bs) pure returns(string, uint256)
func (_Abi *AbiCallerSession) DecodeData(bs []byte) (string, *big.Int, error) {
	return _Abi.Contract.DecodeData(&_Abi.CallOpts, bs)
}

// EncodeData is a free data retrieval call binding the contract method 0x47da7f8e.
//
// Solidity: function encodeData(string text, uint256 number) pure returns(bytes, bytes)
func (_Abi *AbiCaller) EncodeData(opts *bind.CallOpts, text string, number *big.Int) ([]byte, []byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "encodeData", text, number)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// EncodeData is a free data retrieval call binding the contract method 0x47da7f8e.
//
// Solidity: function encodeData(string text, uint256 number) pure returns(bytes, bytes)
func (_Abi *AbiSession) EncodeData(text string, number *big.Int) ([]byte, []byte, error) {
	return _Abi.Contract.EncodeData(&_Abi.CallOpts, text, number)
}

// EncodeData is a free data retrieval call binding the contract method 0x47da7f8e.
//
// Solidity: function encodeData(string text, uint256 number) pure returns(bytes, bytes)
func (_Abi *AbiCallerSession) EncodeData(text string, number *big.Int) ([]byte, []byte, error) {
	return _Abi.Contract.EncodeData(&_Abi.CallOpts, text, number)
}

// EncodeFunctionCall is a free data retrieval call binding the contract method 0xf435898b.
//
// Solidity: function encodeFunctionCall() pure returns(bytes)
func (_Abi *AbiCaller) EncodeFunctionCall(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "encodeFunctionCall")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeFunctionCall is a free data retrieval call binding the contract method 0xf435898b.
//
// Solidity: function encodeFunctionCall() pure returns(bytes)
func (_Abi *AbiSession) EncodeFunctionCall() ([]byte, error) {
	return _Abi.Contract.EncodeFunctionCall(&_Abi.CallOpts)
}

// EncodeFunctionCall is a free data retrieval call binding the contract method 0xf435898b.
//
// Solidity: function encodeFunctionCall() pure returns(bytes)
func (_Abi *AbiCallerSession) EncodeFunctionCall() ([]byte, error) {
	return _Abi.Contract.EncodeFunctionCall(&_Abi.CallOpts)
}

// GetSelector is a free data retrieval call binding the contract method 0x034899bc.
//
// Solidity: function getSelector() pure returns(bytes4)
func (_Abi *AbiCaller) GetSelector(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getSelector")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// GetSelector is a free data retrieval call binding the contract method 0x034899bc.
//
// Solidity: function getSelector() pure returns(bytes4)
func (_Abi *AbiSession) GetSelector() ([4]byte, error) {
	return _Abi.Contract.GetSelector(&_Abi.CallOpts)
}

// GetSelector is a free data retrieval call binding the contract method 0x034899bc.
//
// Solidity: function getSelector() pure returns(bytes4)
func (_Abi *AbiCallerSession) GetSelector() ([4]byte, error) {
	return _Abi.Contract.GetSelector(&_Abi.CallOpts)
}

// HashFunctions is a free data retrieval call binding the contract method 0xb2b34f77.
//
// Solidity: function hashFunctions(string str) pure returns(bytes32, bytes32, bytes32)
func (_Abi *AbiCaller) HashFunctions(opts *bind.CallOpts, str string) ([32]byte, [32]byte, [32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "hashFunctions", str)

	if err != nil {
		return *new([32]byte), *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	out2 := *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return out0, out1, out2, err

}

// HashFunctions is a free data retrieval call binding the contract method 0xb2b34f77.
//
// Solidity: function hashFunctions(string str) pure returns(bytes32, bytes32, bytes32)
func (_Abi *AbiSession) HashFunctions(str string) ([32]byte, [32]byte, [32]byte, error) {
	return _Abi.Contract.HashFunctions(&_Abi.CallOpts, str)
}

// HashFunctions is a free data retrieval call binding the contract method 0xb2b34f77.
//
// Solidity: function hashFunctions(string str) pure returns(bytes32, bytes32, bytes32)
func (_Abi *AbiCallerSession) HashFunctions(str string) ([32]byte, [32]byte, [32]byte, error) {
	return _Abi.Contract.HashFunctions(&_Abi.CallOpts, str)
}

// MathTest is a free data retrieval call binding the contract method 0xec5dfcd8.
//
// Solidity: function mathTest(uint256 x, uint256 y, uint256 m) pure returns(uint256, uint256)
func (_Abi *AbiCaller) MathTest(opts *bind.CallOpts, x *big.Int, y *big.Int, m *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "mathTest", x, y, m)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// MathTest is a free data retrieval call binding the contract method 0xec5dfcd8.
//
// Solidity: function mathTest(uint256 x, uint256 y, uint256 m) pure returns(uint256, uint256)
func (_Abi *AbiSession) MathTest(x *big.Int, y *big.Int, m *big.Int) (*big.Int, *big.Int, error) {
	return _Abi.Contract.MathTest(&_Abi.CallOpts, x, y, m)
}

// MathTest is a free data retrieval call binding the contract method 0xec5dfcd8.
//
// Solidity: function mathTest(uint256 x, uint256 y, uint256 m) pure returns(uint256, uint256)
func (_Abi *AbiCallerSession) MathTest(x *big.Int, y *big.Int, m *big.Int) (*big.Int, *big.Int, error) {
	return _Abi.Contract.MathTest(&_Abi.CallOpts, x, y, m)
}

// RecoverAddress is a free data retrieval call binding the contract method 0x8428cf83.
//
// Solidity: function recoverAddress(bytes32 hash, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Abi *AbiCaller) RecoverAddress(opts *bind.CallOpts, hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "recoverAddress", hash, v, r, s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverAddress is a free data retrieval call binding the contract method 0x8428cf83.
//
// Solidity: function recoverAddress(bytes32 hash, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Abi *AbiSession) RecoverAddress(hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Abi.Contract.RecoverAddress(&_Abi.CallOpts, hash, v, r, s)
}

// RecoverAddress is a free data retrieval call binding the contract method 0x8428cf83.
//
// Solidity: function recoverAddress(bytes32 hash, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Abi *AbiCallerSession) RecoverAddress(hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Abi.Contract.RecoverAddress(&_Abi.CallOpts, hash, v, r, s)
}

// Transfer is a free data retrieval call binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address addr, uint256 amount) pure returns(bytes)
func (_Abi *AbiCaller) Transfer(opts *bind.CallOpts, addr common.Address, amount *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "transfer", addr, amount)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Transfer is a free data retrieval call binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address addr, uint256 amount) pure returns(bytes)
func (_Abi *AbiSession) Transfer(addr common.Address, amount *big.Int) ([]byte, error) {
	return _Abi.Contract.Transfer(&_Abi.CallOpts, addr, amount)
}

// Transfer is a free data retrieval call binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address addr, uint256 amount) pure returns(bytes)
func (_Abi *AbiCallerSession) Transfer(addr common.Address, amount *big.Int) ([]byte, error) {
	return _Abi.Contract.Transfer(&_Abi.CallOpts, addr, amount)
}

// CallContract is a paid mutator transaction binding the contract method 0x03710668.
//
// Solidity: function callContract(address addr, bytes data) returns(bool, bytes)
func (_Abi *AbiTransactor) CallContract(opts *bind.TransactOpts, addr common.Address, data []byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "callContract", addr, data)
}

// CallContract is a paid mutator transaction binding the contract method 0x03710668.
//
// Solidity: function callContract(address addr, bytes data) returns(bool, bytes)
func (_Abi *AbiSession) CallContract(addr common.Address, data []byte) (*types.Transaction, error) {
	return _Abi.Contract.CallContract(&_Abi.TransactOpts, addr, data)
}

// CallContract is a paid mutator transaction binding the contract method 0x03710668.
//
// Solidity: function callContract(address addr, bytes data) returns(bool, bytes)
func (_Abi *AbiTransactorSession) CallContract(addr common.Address, data []byte) (*types.Transaction, error) {
	return _Abi.Contract.CallContract(&_Abi.TransactOpts, addr, data)
}
