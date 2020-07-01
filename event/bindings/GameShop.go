// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ActivatableABI is the input ABI used to generate the binding from.
const ActivatableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Activate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Deactivate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ActivatableFuncSigs maps the 4-byte function signature to its string representation.
var ActivatableFuncSigs = map[string]string{
	"0f15f4c0": "activate()",
	"02fb0c5e": "active()",
	"51b42b00": "deactivate()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// ActivatableBin is the compiled bytecode used for deploying new contracts.
var ActivatableBin = "0x6080604052600080546001600160a81b031916331790556102fb806100256000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806302fb0c5e146100675780630f15f4c01461008357806351b42b001461008d578063715018a6146100955780638da5cb5b1461009d578063f2fde38b146100c1575b600080fd5b61006f6100e7565b604080519115158252519081900360200190f35b61008b6100f7565b005b61008b610162565b61008b6101c6565b6100a5610225565b604080516001600160a01b039092168252519081900360200190f35b61008b600480360360208110156100d757600080fd5b50356001600160a01b0316610234565b600054600160a01b900460ff1681565b600054600160a01b900460ff161561010e57600080fd5b6000546001600160a01b0316331461012557600080fd5b6000805460ff60a01b1916600160a01b17815560405133917ff7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba1591a2565b600054600160a01b900460ff1661017857600080fd5b6000546001600160a01b0316331461018f57600080fd5b6000805460ff60a01b1916815560405133917f238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c91a2565b6000546001600160a01b031633146101dd57600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6000546001600160a01b031681565b6000546001600160a01b0316331461024b57600080fd5b61025481610257565b50565b6001600160a01b03811661026a57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea26469706673582212209a44b143237860ba4f1c57338359675eb2fd33c19b7d5a5081cfb10f8db6b77c64736f6c63430006070033"

// DeployActivatable deploys a new Ethereum contract, binding an instance of Activatable to it.
func DeployActivatable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Activatable, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivatableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ActivatableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Activatable{ActivatableCaller: ActivatableCaller{contract: contract}, ActivatableTransactor: ActivatableTransactor{contract: contract}, ActivatableFilterer: ActivatableFilterer{contract: contract}}, nil
}

// Activatable is an auto generated Go binding around an Ethereum contract.
type Activatable struct {
	ActivatableCaller     // Read-only binding to the contract
	ActivatableTransactor // Write-only binding to the contract
	ActivatableFilterer   // Log filterer for contract events
}

// ActivatableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivatableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivatableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivatableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivatableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivatableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivatableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivatableSession struct {
	Contract     *Activatable      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActivatableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivatableCallerSession struct {
	Contract *ActivatableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ActivatableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivatableTransactorSession struct {
	Contract     *ActivatableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ActivatableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivatableRaw struct {
	Contract *Activatable // Generic contract binding to access the raw methods on
}

// ActivatableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivatableCallerRaw struct {
	Contract *ActivatableCaller // Generic read-only contract binding to access the raw methods on
}

// ActivatableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivatableTransactorRaw struct {
	Contract *ActivatableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivatable creates a new instance of Activatable, bound to a specific deployed contract.
func NewActivatable(address common.Address, backend bind.ContractBackend) (*Activatable, error) {
	contract, err := bindActivatable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Activatable{ActivatableCaller: ActivatableCaller{contract: contract}, ActivatableTransactor: ActivatableTransactor{contract: contract}, ActivatableFilterer: ActivatableFilterer{contract: contract}}, nil
}

// NewActivatableCaller creates a new read-only instance of Activatable, bound to a specific deployed contract.
func NewActivatableCaller(address common.Address, caller bind.ContractCaller) (*ActivatableCaller, error) {
	contract, err := bindActivatable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivatableCaller{contract: contract}, nil
}

// NewActivatableTransactor creates a new write-only instance of Activatable, bound to a specific deployed contract.
func NewActivatableTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivatableTransactor, error) {
	contract, err := bindActivatable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivatableTransactor{contract: contract}, nil
}

// NewActivatableFilterer creates a new log filterer instance of Activatable, bound to a specific deployed contract.
func NewActivatableFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivatableFilterer, error) {
	contract, err := bindActivatable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivatableFilterer{contract: contract}, nil
}

// bindActivatable binds a generic wrapper to an already deployed contract.
func bindActivatable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivatableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Activatable *ActivatableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Activatable.Contract.ActivatableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Activatable *ActivatableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.Contract.ActivatableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Activatable *ActivatableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Activatable.Contract.ActivatableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Activatable *ActivatableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Activatable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Activatable *ActivatableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Activatable *ActivatableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Activatable.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Activatable *ActivatableCaller) Active(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Activatable.contract.Call(opts, out, "active")
	return *ret0, err
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Activatable *ActivatableSession) Active() (bool, error) {
	return _Activatable.Contract.Active(&_Activatable.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Activatable *ActivatableCallerSession) Active() (bool, error) {
	return _Activatable.Contract.Active(&_Activatable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Activatable *ActivatableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Activatable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Activatable *ActivatableSession) Owner() (common.Address, error) {
	return _Activatable.Contract.Owner(&_Activatable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Activatable *ActivatableCallerSession) Owner() (common.Address, error) {
	return _Activatable.Contract.Owner(&_Activatable.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Activatable *ActivatableTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Activatable *ActivatableSession) Activate() (*types.Transaction, error) {
	return _Activatable.Contract.Activate(&_Activatable.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Activatable *ActivatableTransactorSession) Activate() (*types.Transaction, error) {
	return _Activatable.Contract.Activate(&_Activatable.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Activatable *ActivatableTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Activatable *ActivatableSession) Deactivate() (*types.Transaction, error) {
	return _Activatable.Contract.Deactivate(&_Activatable.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Activatable *ActivatableTransactorSession) Deactivate() (*types.Transaction, error) {
	return _Activatable.Contract.Deactivate(&_Activatable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Activatable *ActivatableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Activatable *ActivatableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Activatable.Contract.RenounceOwnership(&_Activatable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Activatable *ActivatableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Activatable.Contract.RenounceOwnership(&_Activatable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Activatable *ActivatableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Activatable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Activatable *ActivatableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Activatable.Contract.TransferOwnership(&_Activatable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Activatable *ActivatableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Activatable.Contract.TransferOwnership(&_Activatable.TransactOpts, _newOwner)
}

// ActivatableActivateIterator is returned from FilterActivate and is used to iterate over the raw logs and unpacked data for Activate events raised by the Activatable contract.
type ActivatableActivateIterator struct {
	Event *ActivatableActivate // Event containing the contract specifics and raw log

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
func (it *ActivatableActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivatableActivate)
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
		it.Event = new(ActivatableActivate)
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
func (it *ActivatableActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivatableActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivatableActivate represents a Activate event raised by the Activatable contract.
type ActivatableActivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterActivate is a free log retrieval operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Activatable *ActivatableFilterer) FilterActivate(opts *bind.FilterOpts, _sender []common.Address) (*ActivatableActivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Activatable.contract.FilterLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ActivatableActivateIterator{contract: _Activatable.contract, event: "Activate", logs: logs, sub: sub}, nil
}

// WatchActivate is a free log subscription operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Activatable *ActivatableFilterer) WatchActivate(opts *bind.WatchOpts, sink chan<- *ActivatableActivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Activatable.contract.WatchLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivatableActivate)
				if err := _Activatable.contract.UnpackLog(event, "Activate", log); err != nil {
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

// ParseActivate is a log parse operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Activatable *ActivatableFilterer) ParseActivate(log types.Log) (*ActivatableActivate, error) {
	event := new(ActivatableActivate)
	if err := _Activatable.contract.UnpackLog(event, "Activate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ActivatableDeactivateIterator is returned from FilterDeactivate and is used to iterate over the raw logs and unpacked data for Deactivate events raised by the Activatable contract.
type ActivatableDeactivateIterator struct {
	Event *ActivatableDeactivate // Event containing the contract specifics and raw log

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
func (it *ActivatableDeactivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivatableDeactivate)
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
		it.Event = new(ActivatableDeactivate)
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
func (it *ActivatableDeactivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivatableDeactivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivatableDeactivate represents a Deactivate event raised by the Activatable contract.
type ActivatableDeactivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeactivate is a free log retrieval operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Activatable *ActivatableFilterer) FilterDeactivate(opts *bind.FilterOpts, _sender []common.Address) (*ActivatableDeactivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Activatable.contract.FilterLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ActivatableDeactivateIterator{contract: _Activatable.contract, event: "Deactivate", logs: logs, sub: sub}, nil
}

// WatchDeactivate is a free log subscription operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Activatable *ActivatableFilterer) WatchDeactivate(opts *bind.WatchOpts, sink chan<- *ActivatableDeactivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Activatable.contract.WatchLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivatableDeactivate)
				if err := _Activatable.contract.UnpackLog(event, "Deactivate", log); err != nil {
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

// ParseDeactivate is a log parse operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Activatable *ActivatableFilterer) ParseDeactivate(log types.Log) (*ActivatableDeactivate, error) {
	event := new(ActivatableDeactivate)
	if err := _Activatable.contract.UnpackLog(event, "Deactivate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ActivatableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Activatable contract.
type ActivatableOwnershipRenouncedIterator struct {
	Event *ActivatableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ActivatableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivatableOwnershipRenounced)
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
		it.Event = new(ActivatableOwnershipRenounced)
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
func (it *ActivatableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivatableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivatableOwnershipRenounced represents a OwnershipRenounced event raised by the Activatable contract.
type ActivatableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Activatable *ActivatableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ActivatableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Activatable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ActivatableOwnershipRenouncedIterator{contract: _Activatable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Activatable *ActivatableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ActivatableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Activatable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivatableOwnershipRenounced)
				if err := _Activatable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Activatable *ActivatableFilterer) ParseOwnershipRenounced(log types.Log) (*ActivatableOwnershipRenounced, error) {
	event := new(ActivatableOwnershipRenounced)
	if err := _Activatable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ActivatableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Activatable contract.
type ActivatableOwnershipTransferredIterator struct {
	Event *ActivatableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ActivatableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivatableOwnershipTransferred)
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
		it.Event = new(ActivatableOwnershipTransferred)
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
func (it *ActivatableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivatableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivatableOwnershipTransferred represents a OwnershipTransferred event raised by the Activatable contract.
type ActivatableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Activatable *ActivatableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ActivatableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Activatable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ActivatableOwnershipTransferredIterator{contract: _Activatable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Activatable *ActivatableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ActivatableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Activatable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivatableOwnershipTransferred)
				if err := _Activatable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Activatable *ActivatableFilterer) ParseOwnershipTransferred(log types.Log) (*ActivatableOwnershipTransferred, error) {
	event := new(ActivatableOwnershipTransferred)
	if err := _Activatable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DestructibleABI is the input ABI used to generate the binding from.
const DestructibleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DestructibleFuncSigs maps the 4-byte function signature to its string representation.
var DestructibleFuncSigs = map[string]string{
	"83197ef0": "destroy()",
	"f5074f41": "destroyAndSend(address)",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// DestructibleBin is the compiled bytecode used for deploying new contracts.
var DestructibleBin = "0x6080604052600080546001600160a01b0319163317905561025b806100256000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063715018a61461005c57806383197ef0146100665780638da5cb5b1461006e578063f2fde38b14610092578063f5074f41146100b8575b600080fd5b6100646100de565b005b61006461013d565b610076610162565b604080516001600160a01b039092168252519081900360200190f35b610064600480360360208110156100a857600080fd5b50356001600160a01b0316610171565b610064600480360360208110156100ce57600080fd5b50356001600160a01b0316610194565b6000546001600160a01b031633146100f557600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6000546001600160a01b0316331461015457600080fd5b6000546001600160a01b0316ff5b6000546001600160a01b031681565b6000546001600160a01b0316331461018857600080fd5b610191816101b7565b50565b6000546001600160a01b031633146101ab57600080fd5b806001600160a01b0316ff5b6001600160a01b0381166101ca57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea26469706673582212204d278c566aac78281cd8d70e0fa53eff6e1ac785bef4366f886d6ae1e15ce16c64736f6c63430006070033"

// DeployDestructible deploys a new Ethereum contract, binding an instance of Destructible to it.
func DeployDestructible(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Destructible, error) {
	parsed, err := abi.JSON(strings.NewReader(DestructibleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DestructibleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Destructible{DestructibleCaller: DestructibleCaller{contract: contract}, DestructibleTransactor: DestructibleTransactor{contract: contract}, DestructibleFilterer: DestructibleFilterer{contract: contract}}, nil
}

// Destructible is an auto generated Go binding around an Ethereum contract.
type Destructible struct {
	DestructibleCaller     // Read-only binding to the contract
	DestructibleTransactor // Write-only binding to the contract
	DestructibleFilterer   // Log filterer for contract events
}

// DestructibleCaller is an auto generated read-only Go binding around an Ethereum contract.
type DestructibleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestructibleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DestructibleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestructibleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DestructibleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestructibleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DestructibleSession struct {
	Contract     *Destructible     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DestructibleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DestructibleCallerSession struct {
	Contract *DestructibleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DestructibleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DestructibleTransactorSession struct {
	Contract     *DestructibleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DestructibleRaw is an auto generated low-level Go binding around an Ethereum contract.
type DestructibleRaw struct {
	Contract *Destructible // Generic contract binding to access the raw methods on
}

// DestructibleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DestructibleCallerRaw struct {
	Contract *DestructibleCaller // Generic read-only contract binding to access the raw methods on
}

// DestructibleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DestructibleTransactorRaw struct {
	Contract *DestructibleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDestructible creates a new instance of Destructible, bound to a specific deployed contract.
func NewDestructible(address common.Address, backend bind.ContractBackend) (*Destructible, error) {
	contract, err := bindDestructible(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Destructible{DestructibleCaller: DestructibleCaller{contract: contract}, DestructibleTransactor: DestructibleTransactor{contract: contract}, DestructibleFilterer: DestructibleFilterer{contract: contract}}, nil
}

// NewDestructibleCaller creates a new read-only instance of Destructible, bound to a specific deployed contract.
func NewDestructibleCaller(address common.Address, caller bind.ContractCaller) (*DestructibleCaller, error) {
	contract, err := bindDestructible(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DestructibleCaller{contract: contract}, nil
}

// NewDestructibleTransactor creates a new write-only instance of Destructible, bound to a specific deployed contract.
func NewDestructibleTransactor(address common.Address, transactor bind.ContractTransactor) (*DestructibleTransactor, error) {
	contract, err := bindDestructible(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DestructibleTransactor{contract: contract}, nil
}

// NewDestructibleFilterer creates a new log filterer instance of Destructible, bound to a specific deployed contract.
func NewDestructibleFilterer(address common.Address, filterer bind.ContractFilterer) (*DestructibleFilterer, error) {
	contract, err := bindDestructible(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DestructibleFilterer{contract: contract}, nil
}

// bindDestructible binds a generic wrapper to an already deployed contract.
func bindDestructible(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DestructibleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destructible *DestructibleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Destructible.Contract.DestructibleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destructible *DestructibleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destructible.Contract.DestructibleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destructible *DestructibleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destructible.Contract.DestructibleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destructible *DestructibleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Destructible.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destructible *DestructibleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destructible.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destructible *DestructibleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destructible.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destructible *DestructibleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Destructible.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destructible *DestructibleSession) Owner() (common.Address, error) {
	return _Destructible.Contract.Owner(&_Destructible.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destructible *DestructibleCallerSession) Owner() (common.Address, error) {
	return _Destructible.Contract.Owner(&_Destructible.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Destructible *DestructibleTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destructible.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Destructible *DestructibleSession) Destroy() (*types.Transaction, error) {
	return _Destructible.Contract.Destroy(&_Destructible.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Destructible *DestructibleTransactorSession) Destroy() (*types.Transaction, error) {
	return _Destructible.Contract.Destroy(&_Destructible.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_Destructible *DestructibleTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _Destructible.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_Destructible *DestructibleSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Destructible.Contract.DestroyAndSend(&_Destructible.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_Destructible *DestructibleTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Destructible.Contract.DestroyAndSend(&_Destructible.TransactOpts, _recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destructible *DestructibleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destructible.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destructible *DestructibleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Destructible.Contract.RenounceOwnership(&_Destructible.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destructible *DestructibleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Destructible.Contract.RenounceOwnership(&_Destructible.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Destructible *DestructibleTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Destructible.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Destructible *DestructibleSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Destructible.Contract.TransferOwnership(&_Destructible.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Destructible *DestructibleTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Destructible.Contract.TransferOwnership(&_Destructible.TransactOpts, _newOwner)
}

// DestructibleOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Destructible contract.
type DestructibleOwnershipRenouncedIterator struct {
	Event *DestructibleOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *DestructibleOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestructibleOwnershipRenounced)
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
		it.Event = new(DestructibleOwnershipRenounced)
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
func (it *DestructibleOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestructibleOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestructibleOwnershipRenounced represents a OwnershipRenounced event raised by the Destructible contract.
type DestructibleOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Destructible *DestructibleFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DestructibleOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Destructible.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DestructibleOwnershipRenouncedIterator{contract: _Destructible.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Destructible *DestructibleFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DestructibleOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Destructible.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestructibleOwnershipRenounced)
				if err := _Destructible.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Destructible *DestructibleFilterer) ParseOwnershipRenounced(log types.Log) (*DestructibleOwnershipRenounced, error) {
	event := new(DestructibleOwnershipRenounced)
	if err := _Destructible.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DestructibleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Destructible contract.
type DestructibleOwnershipTransferredIterator struct {
	Event *DestructibleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DestructibleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestructibleOwnershipTransferred)
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
		it.Event = new(DestructibleOwnershipTransferred)
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
func (it *DestructibleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestructibleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestructibleOwnershipTransferred represents a OwnershipTransferred event raised by the Destructible contract.
type DestructibleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Destructible *DestructibleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DestructibleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destructible.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DestructibleOwnershipTransferredIterator{contract: _Destructible.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Destructible *DestructibleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DestructibleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destructible.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestructibleOwnershipTransferred)
				if err := _Destructible.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Destructible *DestructibleFilterer) ParseOwnershipTransferred(log types.Log) (*DestructibleOwnershipTransferred, error) {
	event := new(DestructibleOwnershipTransferred)
	if err := _Destructible.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameShopABI is the input ABI used to generate the binding from.
const GameShopABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Activate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"BuyItem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Deactivate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_exhibitor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"}],\"name\":\"PayToExhibitor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_refundBalance\",\"type\":\"uint256\"}],\"name\":\"RefuncdToOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_exhibitor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"RegisterItem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"exhibitor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"payToExhibitor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundToOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"registerItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GameShopFuncSigs maps the 4-byte function signature to its string representation.
var GameShopFuncSigs = map[string]string{
	"0f15f4c0": "activate()",
	"02fb0c5e": "active()",
	"e7fb74c7": "buyItem(uint256)",
	"51b42b00": "deactivate()",
	"83197ef0": "destroy()",
	"f5074f41": "destroyAndSend(address)",
	"ddca3f43": "fee()",
	"bfb231d2": "items(uint256)",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"dc803012": "payToExhibitor(uint256)",
	"3e4e10d6": "refundToOwner()",
	"62e7d264": "registerItem(uint256)",
	"715018a6": "renounceOwnership()",
	"69fe0e2d": "setFee(uint256)",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// GameShopBin is the compiled bytecode used for deploying new contracts.
var GameShopBin = "0x6080604052600080546001600160b01b03191633179055610a30806100256000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806383197ef0116100a2578063dc80301211610071578063dc80301214610217578063ddca3f4314610234578063e7fb74c71461024e578063f2fde38b1461026b578063f5074f411461029157610116565b806383197ef0146101a35780638456cb59146101ab5780638da5cb5b146101b3578063bfb231d2146101d757610116565b806351b42b00116100e957806351b42b00146101515780635c975abb1461015957806362e7d2641461016157806369fe0e2d1461017e578063715018a61461019b57610116565b806302fb0c5e1461011b5780630f15f4c0146101375780633e4e10d6146101415780633f4ba83a14610149575b600080fd5b6101236102b7565b604080519115158252519081900360200190f35b61013f6102c7565b005b61013f610332565b61013f6103e7565b61013f61044a565b6101236104ae565b61013f6004803603602081101561017757600080fd5b50356104be565b61013f6004803603602081101561019457600080fd5b50356105c9565b61013f6105f2565b61013f610651565b61013f610676565b6101bb6106e0565b604080516001600160a01b039092168252519081900360200190f35b6101f4600480360360208110156101ed57600080fd5b50356106ef565b604080516001600160a01b03909316835260208301919091528051918290030190f35b61013f6004803603602081101561022d57600080fd5b5035610724565b61023c610893565b60408051918252519081900360200190f35b61013f6004803603602081101561026457600080fd5b5035610899565b61013f6004803603602081101561028157600080fd5b50356001600160a01b031661092f565b61013f600480360360208110156102a757600080fd5b50356001600160a01b0316610952565b600054600160a81b900460ff1681565b600054600160a81b900460ff16156102de57600080fd5b6000546001600160a01b031633146102f557600080fd5b6000805460ff60a81b1916600160a81b17815560405133917ff7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba1591a2565b600054600160a81b900460ff161561034957600080fd5b6000546001600160a01b0316331461036057600080fd5b6000471161036d57600080fd5b6000805460405147926001600160a01b039092169183156108fc02918491818181858888f193505050501580156103a8573d6000803e3d6000fd5b50604080513381526020810183905281517f2b2022abde867d8a7a7f0428c7ea254e589d9084d8db8dfafd6663991867f158929181900390910190a150565b6000546001600160a01b031633146103fe57600080fd5b600054600160a01b900460ff1661041457600080fd5b6000805460ff60a01b191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b600054600160a81b900460ff1661046057600080fd5b6000546001600160a01b0316331461047757600080fd5b6000805460ff60a81b1916815560405133917f238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c91a2565b600054600160a01b900460ff1681565b600054600160a01b900460ff16156104d557600080fd5b600081116104e257600080fd5b604080518082018252338082526020808301858152600480546001810182556000829052945160029095027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b810180546001600160a01b03979097166001600160a01b03199097169690961790955590517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c90940193909355915483516000199190910180825292810191909152808301849052915190917fba5b480d77b78f6111acbcab260e6ddf3f148ac5a0eef77d0bf4106498ebedbf919081900360600190a15050565b6000546001600160a01b031633146105e057600080fd5b600081116105ed57600080fd5b600155565b6000546001600160a01b0316331461060957600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6000546001600160a01b0316331461066857600080fd5b6000546001600160a01b0316ff5b6000546001600160a01b0316331461068d57600080fd5b600054600160a01b900460ff16156106a457600080fd5b6000805460ff60a01b1916600160a01b1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b6000546001600160a01b031681565b600481815481106106fc57fe5b6000918252602090912060029091020180546001909101546001600160a01b03909116915082565b6000546001600160a01b0316331461073b57600080fd5b6107436109e3565b6004828154811061075057fe5b60009182526020808320604080518082018252600290940290910180546001600160a01b03168452600101548383015285845260039091529091205490915060ff161561079c57600080fd5b6001548160200151116107ae57600080fd5b60015481602001510347116107c257600080fd5b60005481516001600160a01b03908116911614156107df57600080fd5b6000828152600360209081526040808320805460ff19166001908117909155845190549285015191516001600160a01b03909116939290910380156108fc0292909190818181858888f1935050505015801561083f573d6000803e3d6000fd5b508051600154602080840151604080516001600160a01b0390951685529290039083015280517f3ead597c83fac751fe46dd2d957eee90483295ccb7af5cc389e317e7ceee18529281900390910190a15050565b60015481565b600054600160a01b900460ff16156108b057600080fd5b6000818152600260205260409020546001600160a01b03166108d157600080fd5b60008181526002602090815260409182902080546001600160a01b0319163390811790915582518481529182015281517f3b389fd404da5e441ca41a7645775a68f254e9153177fa10d255a7315391aa52929181900390910190a150565b6000546001600160a01b0316331461094657600080fd5b61094f81610975565b50565b6000546001600160a01b0316331461096957600080fd5b806001600160a01b0316ff5b6001600160a01b03811661098857600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201529056fea2646970667358221220dc1bc7e1034600b6e44530d5b1dc6c99ae95c98c9e51212da0449d126439e62f64736f6c63430006070033"

// DeployGameShop deploys a new Ethereum contract, binding an instance of GameShop to it.
func DeployGameShop(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GameShop, error) {
	parsed, err := abi.JSON(strings.NewReader(GameShopABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GameShopBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GameShop{GameShopCaller: GameShopCaller{contract: contract}, GameShopTransactor: GameShopTransactor{contract: contract}, GameShopFilterer: GameShopFilterer{contract: contract}}, nil
}

// GameShop is an auto generated Go binding around an Ethereum contract.
type GameShop struct {
	GameShopCaller     // Read-only binding to the contract
	GameShopTransactor // Write-only binding to the contract
	GameShopFilterer   // Log filterer for contract events
}

// GameShopCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameShopCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameShopTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameShopTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameShopFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameShopFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameShopSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameShopSession struct {
	Contract     *GameShop         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameShopCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameShopCallerSession struct {
	Contract *GameShopCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GameShopTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameShopTransactorSession struct {
	Contract     *GameShopTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GameShopRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameShopRaw struct {
	Contract *GameShop // Generic contract binding to access the raw methods on
}

// GameShopCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameShopCallerRaw struct {
	Contract *GameShopCaller // Generic read-only contract binding to access the raw methods on
}

// GameShopTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameShopTransactorRaw struct {
	Contract *GameShopTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameShop creates a new instance of GameShop, bound to a specific deployed contract.
func NewGameShop(address common.Address, backend bind.ContractBackend) (*GameShop, error) {
	contract, err := bindGameShop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameShop{GameShopCaller: GameShopCaller{contract: contract}, GameShopTransactor: GameShopTransactor{contract: contract}, GameShopFilterer: GameShopFilterer{contract: contract}}, nil
}

// NewGameShopCaller creates a new read-only instance of GameShop, bound to a specific deployed contract.
func NewGameShopCaller(address common.Address, caller bind.ContractCaller) (*GameShopCaller, error) {
	contract, err := bindGameShop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameShopCaller{contract: contract}, nil
}

// NewGameShopTransactor creates a new write-only instance of GameShop, bound to a specific deployed contract.
func NewGameShopTransactor(address common.Address, transactor bind.ContractTransactor) (*GameShopTransactor, error) {
	contract, err := bindGameShop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameShopTransactor{contract: contract}, nil
}

// NewGameShopFilterer creates a new log filterer instance of GameShop, bound to a specific deployed contract.
func NewGameShopFilterer(address common.Address, filterer bind.ContractFilterer) (*GameShopFilterer, error) {
	contract, err := bindGameShop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameShopFilterer{contract: contract}, nil
}

// bindGameShop binds a generic wrapper to an already deployed contract.
func bindGameShop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameShopABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameShop *GameShopRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GameShop.Contract.GameShopCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameShop *GameShopRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameShop.Contract.GameShopTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameShop *GameShopRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameShop.Contract.GameShopTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameShop *GameShopCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GameShop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameShop *GameShopTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameShop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameShop *GameShopTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameShop.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_GameShop *GameShopCaller) Active(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GameShop.contract.Call(opts, out, "active")
	return *ret0, err
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_GameShop *GameShopSession) Active() (bool, error) {
	return _GameShop.Contract.Active(&_GameShop.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_GameShop *GameShopCallerSession) Active() (bool, error) {
	return _GameShop.Contract.Active(&_GameShop.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_GameShop *GameShopCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GameShop.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_GameShop *GameShopSession) Fee() (*big.Int, error) {
	return _GameShop.Contract.Fee(&_GameShop.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_GameShop *GameShopCallerSession) Fee() (*big.Int, error) {
	return _GameShop.Contract.Fee(&_GameShop.CallOpts)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(address exhibitor, uint256 price)
func (_GameShop *GameShopCaller) Items(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Exhibitor common.Address
	Price     *big.Int
}, error) {
	ret := new(struct {
		Exhibitor common.Address
		Price     *big.Int
	})
	out := ret
	err := _GameShop.contract.Call(opts, out, "items", arg0)
	return *ret, err
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(address exhibitor, uint256 price)
func (_GameShop *GameShopSession) Items(arg0 *big.Int) (struct {
	Exhibitor common.Address
	Price     *big.Int
}, error) {
	return _GameShop.Contract.Items(&_GameShop.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(address exhibitor, uint256 price)
func (_GameShop *GameShopCallerSession) Items(arg0 *big.Int) (struct {
	Exhibitor common.Address
	Price     *big.Int
}, error) {
	return _GameShop.Contract.Items(&_GameShop.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameShop *GameShopCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GameShop.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameShop *GameShopSession) Owner() (common.Address, error) {
	return _GameShop.Contract.Owner(&_GameShop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameShop *GameShopCallerSession) Owner() (common.Address, error) {
	return _GameShop.Contract.Owner(&_GameShop.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GameShop *GameShopCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GameShop.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GameShop *GameShopSession) Paused() (bool, error) {
	return _GameShop.Contract.Paused(&_GameShop.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GameShop *GameShopCallerSession) Paused() (bool, error) {
	return _GameShop.Contract.Paused(&_GameShop.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_GameShop *GameShopTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_GameShop *GameShopSession) Activate() (*types.Transaction, error) {
	return _GameShop.Contract.Activate(&_GameShop.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_GameShop *GameShopTransactorSession) Activate() (*types.Transaction, error) {
	return _GameShop.Contract.Activate(&_GameShop.TransactOpts)
}

// BuyItem is a paid mutator transaction binding the contract method 0xe7fb74c7.
//
// Solidity: function buyItem(uint256 _id) returns()
func (_GameShop *GameShopTransactor) BuyItem(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "buyItem", _id)
}

// BuyItem is a paid mutator transaction binding the contract method 0xe7fb74c7.
//
// Solidity: function buyItem(uint256 _id) returns()
func (_GameShop *GameShopSession) BuyItem(_id *big.Int) (*types.Transaction, error) {
	return _GameShop.Contract.BuyItem(&_GameShop.TransactOpts, _id)
}

// BuyItem is a paid mutator transaction binding the contract method 0xe7fb74c7.
//
// Solidity: function buyItem(uint256 _id) returns()
func (_GameShop *GameShopTransactorSession) BuyItem(_id *big.Int) (*types.Transaction, error) {
	return _GameShop.Contract.BuyItem(&_GameShop.TransactOpts, _id)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_GameShop *GameShopTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_GameShop *GameShopSession) Deactivate() (*types.Transaction, error) {
	return _GameShop.Contract.Deactivate(&_GameShop.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_GameShop *GameShopTransactorSession) Deactivate() (*types.Transaction, error) {
	return _GameShop.Contract.Deactivate(&_GameShop.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_GameShop *GameShopTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_GameShop *GameShopSession) Destroy() (*types.Transaction, error) {
	return _GameShop.Contract.Destroy(&_GameShop.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_GameShop *GameShopTransactorSession) Destroy() (*types.Transaction, error) {
	return _GameShop.Contract.Destroy(&_GameShop.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_GameShop *GameShopTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_GameShop *GameShopSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _GameShop.Contract.DestroyAndSend(&_GameShop.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_GameShop *GameShopTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _GameShop.Contract.DestroyAndSend(&_GameShop.TransactOpts, _recipient)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GameShop *GameShopTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GameShop *GameShopSession) Pause() (*types.Transaction, error) {
	return _GameShop.Contract.Pause(&_GameShop.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GameShop *GameShopTransactorSession) Pause() (*types.Transaction, error) {
	return _GameShop.Contract.Pause(&_GameShop.TransactOpts)
}

// PayToExhibitor is a paid mutator transaction binding the contract method 0xdc803012.
//
// Solidity: function payToExhibitor(uint256 _id) returns()
func (_GameShop *GameShopTransactor) PayToExhibitor(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "payToExhibitor", _id)
}

// PayToExhibitor is a paid mutator transaction binding the contract method 0xdc803012.
//
// Solidity: function payToExhibitor(uint256 _id) returns()
func (_GameShop *GameShopSession) PayToExhibitor(_id *big.Int) (*types.Transaction, error) {
	return _GameShop.Contract.PayToExhibitor(&_GameShop.TransactOpts, _id)
}

// PayToExhibitor is a paid mutator transaction binding the contract method 0xdc803012.
//
// Solidity: function payToExhibitor(uint256 _id) returns()
func (_GameShop *GameShopTransactorSession) PayToExhibitor(_id *big.Int) (*types.Transaction, error) {
	return _GameShop.Contract.PayToExhibitor(&_GameShop.TransactOpts, _id)
}

// RefundToOwner is a paid mutator transaction binding the contract method 0x3e4e10d6.
//
// Solidity: function refundToOwner() returns()
func (_GameShop *GameShopTransactor) RefundToOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "refundToOwner")
}

// RefundToOwner is a paid mutator transaction binding the contract method 0x3e4e10d6.
//
// Solidity: function refundToOwner() returns()
func (_GameShop *GameShopSession) RefundToOwner() (*types.Transaction, error) {
	return _GameShop.Contract.RefundToOwner(&_GameShop.TransactOpts)
}

// RefundToOwner is a paid mutator transaction binding the contract method 0x3e4e10d6.
//
// Solidity: function refundToOwner() returns()
func (_GameShop *GameShopTransactorSession) RefundToOwner() (*types.Transaction, error) {
	return _GameShop.Contract.RefundToOwner(&_GameShop.TransactOpts)
}

// RegisterItem is a paid mutator transaction binding the contract method 0x62e7d264.
//
// Solidity: function registerItem(uint256 _price) returns()
func (_GameShop *GameShopTransactor) RegisterItem(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "registerItem", _price)
}

// RegisterItem is a paid mutator transaction binding the contract method 0x62e7d264.
//
// Solidity: function registerItem(uint256 _price) returns()
func (_GameShop *GameShopSession) RegisterItem(_price *big.Int) (*types.Transaction, error) {
	return _GameShop.Contract.RegisterItem(&_GameShop.TransactOpts, _price)
}

// RegisterItem is a paid mutator transaction binding the contract method 0x62e7d264.
//
// Solidity: function registerItem(uint256 _price) returns()
func (_GameShop *GameShopTransactorSession) RegisterItem(_price *big.Int) (*types.Transaction, error) {
	return _GameShop.Contract.RegisterItem(&_GameShop.TransactOpts, _price)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameShop *GameShopTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameShop *GameShopSession) RenounceOwnership() (*types.Transaction, error) {
	return _GameShop.Contract.RenounceOwnership(&_GameShop.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameShop *GameShopTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GameShop.Contract.RenounceOwnership(&_GameShop.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _newFee) returns()
func (_GameShop *GameShopTransactor) SetFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "setFee", _newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _newFee) returns()
func (_GameShop *GameShopSession) SetFee(_newFee *big.Int) (*types.Transaction, error) {
	return _GameShop.Contract.SetFee(&_GameShop.TransactOpts, _newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _newFee) returns()
func (_GameShop *GameShopTransactorSession) SetFee(_newFee *big.Int) (*types.Transaction, error) {
	return _GameShop.Contract.SetFee(&_GameShop.TransactOpts, _newFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_GameShop *GameShopTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_GameShop *GameShopSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _GameShop.Contract.TransferOwnership(&_GameShop.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_GameShop *GameShopTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _GameShop.Contract.TransferOwnership(&_GameShop.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GameShop *GameShopTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameShop.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GameShop *GameShopSession) Unpause() (*types.Transaction, error) {
	return _GameShop.Contract.Unpause(&_GameShop.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GameShop *GameShopTransactorSession) Unpause() (*types.Transaction, error) {
	return _GameShop.Contract.Unpause(&_GameShop.TransactOpts)
}

// GameShopActivateIterator is returned from FilterActivate and is used to iterate over the raw logs and unpacked data for Activate events raised by the GameShop contract.
type GameShopActivateIterator struct {
	Event *GameShopActivate // Event containing the contract specifics and raw log

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
func (it *GameShopActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameShopActivate)
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
		it.Event = new(GameShopActivate)
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
func (it *GameShopActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameShopActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameShopActivate represents a Activate event raised by the GameShop contract.
type GameShopActivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterActivate is a free log retrieval operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_GameShop *GameShopFilterer) FilterActivate(opts *bind.FilterOpts, _sender []common.Address) (*GameShopActivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _GameShop.contract.FilterLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &GameShopActivateIterator{contract: _GameShop.contract, event: "Activate", logs: logs, sub: sub}, nil
}

// WatchActivate is a free log subscription operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_GameShop *GameShopFilterer) WatchActivate(opts *bind.WatchOpts, sink chan<- *GameShopActivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _GameShop.contract.WatchLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameShopActivate)
				if err := _GameShop.contract.UnpackLog(event, "Activate", log); err != nil {
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

// ParseActivate is a log parse operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_GameShop *GameShopFilterer) ParseActivate(log types.Log) (*GameShopActivate, error) {
	event := new(GameShopActivate)
	if err := _GameShop.contract.UnpackLog(event, "Activate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameShopBuyItemIterator is returned from FilterBuyItem and is used to iterate over the raw logs and unpacked data for BuyItem events raised by the GameShop contract.
type GameShopBuyItemIterator struct {
	Event *GameShopBuyItem // Event containing the contract specifics and raw log

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
func (it *GameShopBuyItemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameShopBuyItem)
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
		it.Event = new(GameShopBuyItem)
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
func (it *GameShopBuyItemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameShopBuyItemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameShopBuyItem represents a BuyItem event raised by the GameShop contract.
type GameShopBuyItem struct {
	Id    *big.Int
	Buyer common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBuyItem is a free log retrieval operation binding the contract event 0x3b389fd404da5e441ca41a7645775a68f254e9153177fa10d255a7315391aa52.
//
// Solidity: event BuyItem(uint256 _id, address _buyer)
func (_GameShop *GameShopFilterer) FilterBuyItem(opts *bind.FilterOpts) (*GameShopBuyItemIterator, error) {

	logs, sub, err := _GameShop.contract.FilterLogs(opts, "BuyItem")
	if err != nil {
		return nil, err
	}
	return &GameShopBuyItemIterator{contract: _GameShop.contract, event: "BuyItem", logs: logs, sub: sub}, nil
}

// WatchBuyItem is a free log subscription operation binding the contract event 0x3b389fd404da5e441ca41a7645775a68f254e9153177fa10d255a7315391aa52.
//
// Solidity: event BuyItem(uint256 _id, address _buyer)
func (_GameShop *GameShopFilterer) WatchBuyItem(opts *bind.WatchOpts, sink chan<- *GameShopBuyItem) (event.Subscription, error) {

	logs, sub, err := _GameShop.contract.WatchLogs(opts, "BuyItem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameShopBuyItem)
				if err := _GameShop.contract.UnpackLog(event, "BuyItem", log); err != nil {
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

// ParseBuyItem is a log parse operation binding the contract event 0x3b389fd404da5e441ca41a7645775a68f254e9153177fa10d255a7315391aa52.
//
// Solidity: event BuyItem(uint256 _id, address _buyer)
func (_GameShop *GameShopFilterer) ParseBuyItem(log types.Log) (*GameShopBuyItem, error) {
	event := new(GameShopBuyItem)
	if err := _GameShop.contract.UnpackLog(event, "BuyItem", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameShopDeactivateIterator is returned from FilterDeactivate and is used to iterate over the raw logs and unpacked data for Deactivate events raised by the GameShop contract.
type GameShopDeactivateIterator struct {
	Event *GameShopDeactivate // Event containing the contract specifics and raw log

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
func (it *GameShopDeactivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameShopDeactivate)
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
		it.Event = new(GameShopDeactivate)
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
func (it *GameShopDeactivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameShopDeactivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameShopDeactivate represents a Deactivate event raised by the GameShop contract.
type GameShopDeactivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeactivate is a free log retrieval operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_GameShop *GameShopFilterer) FilterDeactivate(opts *bind.FilterOpts, _sender []common.Address) (*GameShopDeactivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _GameShop.contract.FilterLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &GameShopDeactivateIterator{contract: _GameShop.contract, event: "Deactivate", logs: logs, sub: sub}, nil
}

// WatchDeactivate is a free log subscription operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_GameShop *GameShopFilterer) WatchDeactivate(opts *bind.WatchOpts, sink chan<- *GameShopDeactivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _GameShop.contract.WatchLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameShopDeactivate)
				if err := _GameShop.contract.UnpackLog(event, "Deactivate", log); err != nil {
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

// ParseDeactivate is a log parse operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_GameShop *GameShopFilterer) ParseDeactivate(log types.Log) (*GameShopDeactivate, error) {
	event := new(GameShopDeactivate)
	if err := _GameShop.contract.UnpackLog(event, "Deactivate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameShopOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the GameShop contract.
type GameShopOwnershipRenouncedIterator struct {
	Event *GameShopOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *GameShopOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameShopOwnershipRenounced)
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
		it.Event = new(GameShopOwnershipRenounced)
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
func (it *GameShopOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameShopOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameShopOwnershipRenounced represents a OwnershipRenounced event raised by the GameShop contract.
type GameShopOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_GameShop *GameShopFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*GameShopOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _GameShop.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GameShopOwnershipRenouncedIterator{contract: _GameShop.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_GameShop *GameShopFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *GameShopOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _GameShop.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameShopOwnershipRenounced)
				if err := _GameShop.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_GameShop *GameShopFilterer) ParseOwnershipRenounced(log types.Log) (*GameShopOwnershipRenounced, error) {
	event := new(GameShopOwnershipRenounced)
	if err := _GameShop.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameShopOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GameShop contract.
type GameShopOwnershipTransferredIterator struct {
	Event *GameShopOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GameShopOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameShopOwnershipTransferred)
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
		it.Event = new(GameShopOwnershipTransferred)
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
func (it *GameShopOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameShopOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameShopOwnershipTransferred represents a OwnershipTransferred event raised by the GameShop contract.
type GameShopOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameShop *GameShopFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GameShopOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameShop.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GameShopOwnershipTransferredIterator{contract: _GameShop.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameShop *GameShopFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GameShopOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameShop.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameShopOwnershipTransferred)
				if err := _GameShop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GameShop *GameShopFilterer) ParseOwnershipTransferred(log types.Log) (*GameShopOwnershipTransferred, error) {
	event := new(GameShopOwnershipTransferred)
	if err := _GameShop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameShopPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the GameShop contract.
type GameShopPauseIterator struct {
	Event *GameShopPause // Event containing the contract specifics and raw log

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
func (it *GameShopPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameShopPause)
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
		it.Event = new(GameShopPause)
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
func (it *GameShopPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameShopPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameShopPause represents a Pause event raised by the GameShop contract.
type GameShopPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_GameShop *GameShopFilterer) FilterPause(opts *bind.FilterOpts) (*GameShopPauseIterator, error) {

	logs, sub, err := _GameShop.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &GameShopPauseIterator{contract: _GameShop.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_GameShop *GameShopFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *GameShopPause) (event.Subscription, error) {

	logs, sub, err := _GameShop.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameShopPause)
				if err := _GameShop.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_GameShop *GameShopFilterer) ParsePause(log types.Log) (*GameShopPause, error) {
	event := new(GameShopPause)
	if err := _GameShop.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameShopPayToExhibitorIterator is returned from FilterPayToExhibitor and is used to iterate over the raw logs and unpacked data for PayToExhibitor events raised by the GameShop contract.
type GameShopPayToExhibitorIterator struct {
	Event *GameShopPayToExhibitor // Event containing the contract specifics and raw log

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
func (it *GameShopPayToExhibitorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameShopPayToExhibitor)
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
		it.Event = new(GameShopPayToExhibitor)
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
func (it *GameShopPayToExhibitorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameShopPayToExhibitorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameShopPayToExhibitor represents a PayToExhibitor event raised by the GameShop contract.
type GameShopPayToExhibitor struct {
	Exhibitor  common.Address
	AmountPaid *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPayToExhibitor is a free log retrieval operation binding the contract event 0x3ead597c83fac751fe46dd2d957eee90483295ccb7af5cc389e317e7ceee1852.
//
// Solidity: event PayToExhibitor(address _exhibitor, uint256 amountPaid)
func (_GameShop *GameShopFilterer) FilterPayToExhibitor(opts *bind.FilterOpts) (*GameShopPayToExhibitorIterator, error) {

	logs, sub, err := _GameShop.contract.FilterLogs(opts, "PayToExhibitor")
	if err != nil {
		return nil, err
	}
	return &GameShopPayToExhibitorIterator{contract: _GameShop.contract, event: "PayToExhibitor", logs: logs, sub: sub}, nil
}

// WatchPayToExhibitor is a free log subscription operation binding the contract event 0x3ead597c83fac751fe46dd2d957eee90483295ccb7af5cc389e317e7ceee1852.
//
// Solidity: event PayToExhibitor(address _exhibitor, uint256 amountPaid)
func (_GameShop *GameShopFilterer) WatchPayToExhibitor(opts *bind.WatchOpts, sink chan<- *GameShopPayToExhibitor) (event.Subscription, error) {

	logs, sub, err := _GameShop.contract.WatchLogs(opts, "PayToExhibitor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameShopPayToExhibitor)
				if err := _GameShop.contract.UnpackLog(event, "PayToExhibitor", log); err != nil {
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

// ParsePayToExhibitor is a log parse operation binding the contract event 0x3ead597c83fac751fe46dd2d957eee90483295ccb7af5cc389e317e7ceee1852.
//
// Solidity: event PayToExhibitor(address _exhibitor, uint256 amountPaid)
func (_GameShop *GameShopFilterer) ParsePayToExhibitor(log types.Log) (*GameShopPayToExhibitor, error) {
	event := new(GameShopPayToExhibitor)
	if err := _GameShop.contract.UnpackLog(event, "PayToExhibitor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameShopRefuncdToOwnerIterator is returned from FilterRefuncdToOwner and is used to iterate over the raw logs and unpacked data for RefuncdToOwner events raised by the GameShop contract.
type GameShopRefuncdToOwnerIterator struct {
	Event *GameShopRefuncdToOwner // Event containing the contract specifics and raw log

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
func (it *GameShopRefuncdToOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameShopRefuncdToOwner)
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
		it.Event = new(GameShopRefuncdToOwner)
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
func (it *GameShopRefuncdToOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameShopRefuncdToOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameShopRefuncdToOwner represents a RefuncdToOwner event raised by the GameShop contract.
type GameShopRefuncdToOwner struct {
	Owner         common.Address
	RefundBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRefuncdToOwner is a free log retrieval operation binding the contract event 0x2b2022abde867d8a7a7f0428c7ea254e589d9084d8db8dfafd6663991867f158.
//
// Solidity: event RefuncdToOwner(address _owner, uint256 _refundBalance)
func (_GameShop *GameShopFilterer) FilterRefuncdToOwner(opts *bind.FilterOpts) (*GameShopRefuncdToOwnerIterator, error) {

	logs, sub, err := _GameShop.contract.FilterLogs(opts, "RefuncdToOwner")
	if err != nil {
		return nil, err
	}
	return &GameShopRefuncdToOwnerIterator{contract: _GameShop.contract, event: "RefuncdToOwner", logs: logs, sub: sub}, nil
}

// WatchRefuncdToOwner is a free log subscription operation binding the contract event 0x2b2022abde867d8a7a7f0428c7ea254e589d9084d8db8dfafd6663991867f158.
//
// Solidity: event RefuncdToOwner(address _owner, uint256 _refundBalance)
func (_GameShop *GameShopFilterer) WatchRefuncdToOwner(opts *bind.WatchOpts, sink chan<- *GameShopRefuncdToOwner) (event.Subscription, error) {

	logs, sub, err := _GameShop.contract.WatchLogs(opts, "RefuncdToOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameShopRefuncdToOwner)
				if err := _GameShop.contract.UnpackLog(event, "RefuncdToOwner", log); err != nil {
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

// ParseRefuncdToOwner is a log parse operation binding the contract event 0x2b2022abde867d8a7a7f0428c7ea254e589d9084d8db8dfafd6663991867f158.
//
// Solidity: event RefuncdToOwner(address _owner, uint256 _refundBalance)
func (_GameShop *GameShopFilterer) ParseRefuncdToOwner(log types.Log) (*GameShopRefuncdToOwner, error) {
	event := new(GameShopRefuncdToOwner)
	if err := _GameShop.contract.UnpackLog(event, "RefuncdToOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameShopRegisterItemIterator is returned from FilterRegisterItem and is used to iterate over the raw logs and unpacked data for RegisterItem events raised by the GameShop contract.
type GameShopRegisterItemIterator struct {
	Event *GameShopRegisterItem // Event containing the contract specifics and raw log

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
func (it *GameShopRegisterItemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameShopRegisterItem)
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
		it.Event = new(GameShopRegisterItem)
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
func (it *GameShopRegisterItemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameShopRegisterItemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameShopRegisterItem represents a RegisterItem event raised by the GameShop contract.
type GameShopRegisterItem struct {
	Id        *big.Int
	Exhibitor common.Address
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegisterItem is a free log retrieval operation binding the contract event 0xba5b480d77b78f6111acbcab260e6ddf3f148ac5a0eef77d0bf4106498ebedbf.
//
// Solidity: event RegisterItem(uint256 _id, address _exhibitor, uint256 _price)
func (_GameShop *GameShopFilterer) FilterRegisterItem(opts *bind.FilterOpts) (*GameShopRegisterItemIterator, error) {

	logs, sub, err := _GameShop.contract.FilterLogs(opts, "RegisterItem")
	if err != nil {
		return nil, err
	}
	return &GameShopRegisterItemIterator{contract: _GameShop.contract, event: "RegisterItem", logs: logs, sub: sub}, nil
}

// WatchRegisterItem is a free log subscription operation binding the contract event 0xba5b480d77b78f6111acbcab260e6ddf3f148ac5a0eef77d0bf4106498ebedbf.
//
// Solidity: event RegisterItem(uint256 _id, address _exhibitor, uint256 _price)
func (_GameShop *GameShopFilterer) WatchRegisterItem(opts *bind.WatchOpts, sink chan<- *GameShopRegisterItem) (event.Subscription, error) {

	logs, sub, err := _GameShop.contract.WatchLogs(opts, "RegisterItem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameShopRegisterItem)
				if err := _GameShop.contract.UnpackLog(event, "RegisterItem", log); err != nil {
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

// ParseRegisterItem is a log parse operation binding the contract event 0xba5b480d77b78f6111acbcab260e6ddf3f148ac5a0eef77d0bf4106498ebedbf.
//
// Solidity: event RegisterItem(uint256 _id, address _exhibitor, uint256 _price)
func (_GameShop *GameShopFilterer) ParseRegisterItem(log types.Log) (*GameShopRegisterItem, error) {
	event := new(GameShopRegisterItem)
	if err := _GameShop.contract.UnpackLog(event, "RegisterItem", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameShopUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the GameShop contract.
type GameShopUnpauseIterator struct {
	Event *GameShopUnpause // Event containing the contract specifics and raw log

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
func (it *GameShopUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameShopUnpause)
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
		it.Event = new(GameShopUnpause)
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
func (it *GameShopUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameShopUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameShopUnpause represents a Unpause event raised by the GameShop contract.
type GameShopUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_GameShop *GameShopFilterer) FilterUnpause(opts *bind.FilterOpts) (*GameShopUnpauseIterator, error) {

	logs, sub, err := _GameShop.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &GameShopUnpauseIterator{contract: _GameShop.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_GameShop *GameShopFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *GameShopUnpause) (event.Subscription, error) {

	logs, sub, err := _GameShop.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameShopUnpause)
				if err := _GameShop.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_GameShop *GameShopFilterer) ParseUnpause(log types.Log) (*GameShopUnpause, error) {
	event := new(GameShopUnpause)
	if err := _GameShop.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556101cf806100326000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063715018a6146100465780638da5cb5b14610050578063f2fde38b14610074575b600080fd5b61004e61009a565b005b6100586100f9565b604080516001600160a01b039092168252519081900360200190f35b61004e6004803603602081101561008a57600080fd5b50356001600160a01b0316610108565b6000546001600160a01b031633146100b157600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6000546001600160a01b031681565b6000546001600160a01b0316331461011f57600080fd5b6101288161012b565b50565b6001600160a01b03811661013e57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea2646970667358221220c9cc48cacef79e09f9d9d0f3b047201b7e48e70c934f66774c22f8d5c1f72dce64736f6c63430006070033"

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// OwnableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ownable contract.
type OwnableOwnershipRenouncedIterator struct {
	Event *OwnableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipRenounced)
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
		it.Event = new(OwnableOwnershipRenounced)
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
func (it *OwnableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipRenounced represents a OwnershipRenounced event raised by the Ownable contract.
type OwnableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OwnableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipRenouncedIterator{contract: _Ownable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipRenounced)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipRenounced(log types.Log) (*OwnableOwnershipRenounced, error) {
	event := new(OwnableOwnershipRenounced)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// PausableBin is the compiled bytecode used for deploying new contracts.
var PausableBin = "0x6080604052600080546001600160a81b031916331790556102f9806100256000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80633f4ba83a146100675780635c975abb14610071578063715018a61461008d5780638456cb59146100955780638da5cb5b1461009d578063f2fde38b146100c1575b600080fd5b61006f6100e7565b005b61007961014a565b604080519115158252519081900360200190f35b61006f61015a565b61006f6101b9565b6100a5610223565b604080516001600160a01b039092168252519081900360200190f35b61006f600480360360208110156100d757600080fd5b50356001600160a01b0316610232565b6000546001600160a01b031633146100fe57600080fd5b600054600160a01b900460ff1661011457600080fd5b6000805460ff60a01b191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b600054600160a01b900460ff1681565b6000546001600160a01b0316331461017157600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6000546001600160a01b031633146101d057600080fd5b600054600160a01b900460ff16156101e757600080fd5b6000805460ff60a01b1916600160a01b1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b6000546001600160a01b031681565b6000546001600160a01b0316331461024957600080fd5b61025281610255565b50565b6001600160a01b03811661026857600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea2646970667358221220baacb20d7b8592f3041def4c581c342d78c77b1b15652cf9257f333a93b82ff764736f6c63430006070033"

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableCallerSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pausable.Contract.RenounceOwnership(&_Pausable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pausable.Contract.RenounceOwnership(&_Pausable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Pausable *PausableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Pausable *PausableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Pausable *PausableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Pausable contract.
type PausableOwnershipRenouncedIterator struct {
	Event *PausableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipRenounced)
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
		it.Event = new(PausableOwnershipRenounced)
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
func (it *PausableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipRenounced represents a OwnershipRenounced event raised by the Pausable contract.
type PausableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Pausable *PausableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PausableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipRenouncedIterator{contract: _Pausable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Pausable *PausableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PausableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipRenounced)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Pausable *PausableFilterer) ParseOwnershipRenounced(log types.Log) (*PausableOwnershipRenounced, error) {
	event := new(PausableOwnershipRenounced)
	if err := _Pausable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PausableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pausable contract.
type PausableOwnershipTransferredIterator struct {
	Event *PausableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipTransferred)
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
		it.Event = new(PausableOwnershipTransferred)
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
func (it *PausableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipTransferred represents a OwnershipTransferred event raised by the Pausable contract.
type PausableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pausable *PausableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipTransferredIterator{contract: _Pausable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pausable *PausableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipTransferred)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pausable *PausableFilterer) ParseOwnershipTransferred(log types.Log) (*PausableOwnershipTransferred, error) {
	event := new(PausableOwnershipTransferred)
	if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PausablePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Pausable contract.
type PausablePauseIterator struct {
	Event *PausablePause // Event containing the contract specifics and raw log

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
func (it *PausablePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePause)
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
		it.Event = new(PausablePause)
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
func (it *PausablePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) FilterPause(opts *bind.FilterOpts) (*PausablePauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausablePauseIterator{contract: _Pausable.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausablePause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePause)
				if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) ParsePause(log types.Log) (*PausablePause, error) {
	event := new(PausablePause)
	if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PausableUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Pausable contract.
type PausableUnpauseIterator struct {
	Event *PausableUnpause // Event containing the contract specifics and raw log

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
func (it *PausableUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpause)
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
		it.Event = new(PausableUnpause)
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
func (it *PausableUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableUnpauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableUnpauseIterator{contract: _Pausable.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableUnpause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpause)
				if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) ParseUnpause(log types.Log) (*PausableUnpause, error) {
	event := new(PausableUnpause)
	if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	return event, nil
}
