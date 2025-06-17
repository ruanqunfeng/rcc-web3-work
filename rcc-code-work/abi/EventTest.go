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

// EventTestMetaData contains all meta data concerning the EventTest contract.
var EventTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"foo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"Myevent\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// EventTestABI is the input ABI used to generate the binding from.
// Deprecated: Use EventTestMetaData.ABI instead.
var EventTestABI = EventTestMetaData.ABI

// EventTest is an auto generated Go binding around an Ethereum contract.
type EventTest struct {
	EventTestCaller     // Read-only binding to the contract
	EventTestTransactor // Write-only binding to the contract
	EventTestFilterer   // Log filterer for contract events
}

// EventTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventTestSession struct {
	Contract     *EventTest        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventTestCallerSession struct {
	Contract *EventTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EventTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventTestTransactorSession struct {
	Contract     *EventTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EventTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventTestRaw struct {
	Contract *EventTest // Generic contract binding to access the raw methods on
}

// EventTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventTestCallerRaw struct {
	Contract *EventTestCaller // Generic read-only contract binding to access the raw methods on
}

// EventTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventTestTransactorRaw struct {
	Contract *EventTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventTest creates a new instance of EventTest, bound to a specific deployed contract.
func NewEventTest(address common.Address, backend bind.ContractBackend) (*EventTest, error) {
	contract, err := bindEventTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventTest{EventTestCaller: EventTestCaller{contract: contract}, EventTestTransactor: EventTestTransactor{contract: contract}, EventTestFilterer: EventTestFilterer{contract: contract}}, nil
}

// NewEventTestCaller creates a new read-only instance of EventTest, bound to a specific deployed contract.
func NewEventTestCaller(address common.Address, caller bind.ContractCaller) (*EventTestCaller, error) {
	contract, err := bindEventTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventTestCaller{contract: contract}, nil
}

// NewEventTestTransactor creates a new write-only instance of EventTest, bound to a specific deployed contract.
func NewEventTestTransactor(address common.Address, transactor bind.ContractTransactor) (*EventTestTransactor, error) {
	contract, err := bindEventTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventTestTransactor{contract: contract}, nil
}

// NewEventTestFilterer creates a new log filterer instance of EventTest, bound to a specific deployed contract.
func NewEventTestFilterer(address common.Address, filterer bind.ContractFilterer) (*EventTestFilterer, error) {
	contract, err := bindEventTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventTestFilterer{contract: contract}, nil
}

// bindEventTest binds a generic wrapper to an already deployed contract.
func bindEventTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventTest *EventTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventTest.Contract.EventTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventTest *EventTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventTest.Contract.EventTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventTest *EventTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventTest.Contract.EventTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventTest *EventTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventTest *EventTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventTest *EventTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventTest.Contract.contract.Transact(opts, method, params...)
}

// Foo is a paid mutator transaction binding the contract method 0xc2985578.
//
// Solidity: function foo() returns()
func (_EventTest *EventTestTransactor) Foo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventTest.contract.Transact(opts, "foo")
}

// Foo is a paid mutator transaction binding the contract method 0xc2985578.
//
// Solidity: function foo() returns()
func (_EventTest *EventTestSession) Foo() (*types.Transaction, error) {
	return _EventTest.Contract.Foo(&_EventTest.TransactOpts)
}

// Foo is a paid mutator transaction binding the contract method 0xc2985578.
//
// Solidity: function foo() returns()
func (_EventTest *EventTestTransactorSession) Foo() (*types.Transaction, error) {
	return _EventTest.Contract.Foo(&_EventTest.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EventTest *EventTestTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventTest.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EventTest *EventTestSession) Receive() (*types.Transaction, error) {
	return _EventTest.Contract.Receive(&_EventTest.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EventTest *EventTestTransactorSession) Receive() (*types.Transaction, error) {
	return _EventTest.Contract.Receive(&_EventTest.TransactOpts)
}

// EventTestMyeventIterator is returned from FilterMyevent and is used to iterate over the raw logs and unpacked data for Myevent events raised by the EventTest contract.
type EventTestMyeventIterator struct {
	Event *EventTestMyevent // Event containing the contract specifics and raw log

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
func (it *EventTestMyeventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventTestMyevent)
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
		it.Event = new(EventTestMyevent)
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
func (it *EventTestMyeventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventTestMyeventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventTestMyevent represents a Myevent event raised by the EventTest contract.
type EventTestMyevent struct {
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMyevent is a free log retrieval operation binding the contract event 0xb58728ea1fc0389ea1078b5d1833ae97cedee74c5b53b4f6d4281af24be5867c.
//
// Solidity: event Myevent(string description)
func (_EventTest *EventTestFilterer) FilterMyevent(opts *bind.FilterOpts) (*EventTestMyeventIterator, error) {

	logs, sub, err := _EventTest.contract.FilterLogs(opts, "Myevent")
	if err != nil {
		return nil, err
	}
	return &EventTestMyeventIterator{contract: _EventTest.contract, event: "Myevent", logs: logs, sub: sub}, nil
}

// WatchMyevent is a free log subscription operation binding the contract event 0xb58728ea1fc0389ea1078b5d1833ae97cedee74c5b53b4f6d4281af24be5867c.
//
// Solidity: event Myevent(string description)
func (_EventTest *EventTestFilterer) WatchMyevent(opts *bind.WatchOpts, sink chan<- *EventTestMyevent) (event.Subscription, error) {

	logs, sub, err := _EventTest.contract.WatchLogs(opts, "Myevent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventTestMyevent)
				if err := _EventTest.contract.UnpackLog(event, "Myevent", log); err != nil {
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

// ParseMyevent is a log parse operation binding the contract event 0xb58728ea1fc0389ea1078b5d1833ae97cedee74c5b53b4f6d4281af24be5867c.
//
// Solidity: event Myevent(string description)
func (_EventTest *EventTestFilterer) ParseMyevent(log types.Log) (*EventTestMyevent, error) {
	event := new(EventTestMyevent)
	if err := _EventTest.contract.UnpackLog(event, "Myevent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
