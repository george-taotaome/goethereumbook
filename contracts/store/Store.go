// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516104cc3803806104cc83398101604081905261002e91610054565b5f610039828261019c565b5050610257565b634e487b7160e01b5f52604160045260245ffd5b5f6020808385031215610065575f80fd5b82516001600160401b038082111561007b575f80fd5b818501915085601f83011261008e575f80fd5b8151818111156100a0576100a0610040565b604051601f8201601f19908116603f011681019083821181831017156100c8576100c8610040565b8160405282815288868487010111156100df575f80fd5b5f93505b8284101561010057848401860151818501870152928501926100e3565b5f86848301015280965050505050505092915050565b600181811c9082168061012a57607f821691505b60208210810361014857634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610197575f81815260208120601f850160051c810160208610156101745750805b601f850160051c820191505b8181101561019357828155600101610180565b5050505b505050565b81516001600160401b038111156101b5576101b5610040565b6101c9816101c38454610116565b8461014e565b602080601f8311600181146101fc575f84156101e55750858301515b5f19600386901b1c1916600185901b178555610193565b5f85815260208120601f198616915b8281101561022a5788860151825594840194600190910190840161020b565b508582101561024757878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b610268806102645f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806348f343f31461004357806354fd4d5014610075578063f56256c71461008a575b5f80fd5b610062610051366004610178565b60016020525f908152604090205481565b6040519081526020015b60405180910390f35b61007d61009f565b60405161006c919061018f565b61009d6100983660046101da565b61012a565b005b5f80546100ab906101fa565b80601f01602080910402602001604051908101604052809291908181526020018280546100d7906101fa565b80156101225780601f106100f957610100808354040283529160200191610122565b820191905f5260205f20905b81548152906001019060200180831161010557829003601f168201915b505050505081565b5f8281526001602090815260409182902083905581518481529081018390527fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4910160405180910390a15050565b5f60208284031215610188575f80fd5b5035919050565b5f6020808352835180828501525f5b818110156101ba5785810183015185820160400152820161019e565b505f604082860101526040601f19601f8301168501019250505092915050565b5f80604083850312156101eb575f80fd5b50508035926020909101359150565b600181811c9082168061020e57607f821691505b60208210810361022c57634e487b7160e01b5f52602260045260245ffd5b5091905056fea264697066735822122058312a6736c9140d2afc70c92db2112d440f054df6f3f631c2734dc858212a1164736f6c63430008150033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _version string) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Store *StoreCaller) Items(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "items", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Store *StoreSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Store.Contract.Items(&_Store.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Store *StoreCallerSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Store.Contract.Items(&_Store.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store *StoreCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store *StoreSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store *StoreCallerSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Store *StoreTransactor) SetItem(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setItem", key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Store *StoreSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store.Contract.SetItem(&_Store.TransactOpts, key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Store *StoreTransactorSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store.Contract.SetItem(&_Store.TransactOpts, key, value)
}

// StoreItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Store contract.
type StoreItemSetIterator struct {
	Event *StoreItemSet // Event containing the contract specifics and raw log

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
func (it *StoreItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreItemSet)
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
		it.Event = new(StoreItemSet)
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
func (it *StoreItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreItemSet represents a ItemSet event raised by the Store contract.
type StoreItemSet struct {
	Key   [32]byte
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Store *StoreFilterer) FilterItemSet(opts *bind.FilterOpts) (*StoreItemSetIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &StoreItemSetIterator{contract: _Store.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Store *StoreFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *StoreItemSet) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreItemSet)
				if err := _Store.contract.UnpackLog(event, "ItemSet", log); err != nil {
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

// ParseItemSet is a log parse operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Store *StoreFilterer) ParseItemSet(log types.Log) (*StoreItemSet, error) {
	event := new(StoreItemSet)
	if err := _Store.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
