// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vipnodepool

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

// VipnodePoolABI is the input ABI used to generate the binding from.
const VipnodePoolABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"withdrawInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"accounts\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"timeLocked\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"forceWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"opWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"forceSettle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addBalance\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_releaseAmount\",\"type\":\"uint256\"},{\"name\":\"_newBalance\",\"type\":\"uint256\"}],\"name\":\"opSettle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timeLocked\",\"type\":\"uint256\"}],\"name\":\"ForceSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Balance\",\"type\":\"event\"}]"

// VipnodePoolBin is the compiled bytecode used for deploying new contracts.
const VipnodePoolBin = `0x608060405262093a8060015534801561001757600080fd5b506040516020806104f48339810160405251600160a060020a038116151561003e57600080fd5b60008054600160a060020a03909216600160a060020a03199092169190911790556104868061006e6000396000f30060806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663162075d88114610092578063570ca735146100b95780635e5c06e2146100ea5780637be80b3914610124578063ae5fa7e71461013b578063aeea1f7b14610153578063b163cc3814610168578063b27d1ce314610170575b600080fd5b34801561009e57600080fd5b506100a7610197565b60408051918252519081900360200190f35b3480156100c557600080fd5b506100ce61019d565b60408051600160a060020a039092168252519081900360200190f35b3480156100f657600080fd5b5061010b600160a060020a03600435166101ac565b6040805192835260208301919091528051918290030190f35b34801561013057600080fd5b506101396101c5565b005b34801561014757600080fd5b50610139600435610282565b34801561015f57600080fd5b506101396102d6565b61013961033d565b34801561017c57600080fd5b50610139600160a060020a0360043516602435604435610393565b60015481565b600054600160a060020a031681565b6002602052600090815260409020805460019091015482565b336000908152600260205260408120805490919081106101e457600080fd5b60018201546000106101f557600080fd5b600182015442101561020657600080fd5b508054600080835560018301819055604051339183156108fc02918491818181858888f19350505050158015610240573d6000803e3d6000fd5b50815460408051338152602081019290925280517f134e340554ff8a7d64280a2a28b982df554e2595e5bf45cd39368f31099172a69281900390910190a15050565b600054600160a060020a0316331461029957600080fd5b60008054604051600160a060020a039091169183156108fc02918491818181858888f193505050501580156102d2573d6000803e3d6000fd5b5050565b33600090815260026020526040812080549091106102f357600080fd5b60018054420190820181905560408051338152602081019290925280517fde112f62320281f04efec46e09ed286851fdd0b428dfc7214fe0e64f9364323b9281900390910190a150565b3360008181526002602090815260409182902080543401808255835194855291840191909152815190927f134e340554ff8a7d64280a2a28b982df554e2595e5bf45cd39368f31099172a692908290030190a150565b60008054600160a060020a031633146103ab57600080fd5b50600160a060020a0383166000908152600260205260408120828155600181018290559083111561040e57604051600160a060020a0385169084156108fc029085906000818181858888f1935050505015801561040c573d6000803e3d6000fd5b505b805460408051600160a060020a0387168152602081019290925280517f134e340554ff8a7d64280a2a28b982df554e2595e5bf45cd39368f31099172a69281900390910190a1505050505600a165627a7a72305820164ef50c7c781b7be9c350aa8b9f016a386726ec906580e11587a885cf429ca50029`

// DeployVipnodePool deploys a new Ethereum contract, binding an instance of VipnodePool to it.
func DeployVipnodePool(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address) (common.Address, *types.Transaction, *VipnodePool, error) {
	parsed, err := abi.JSON(strings.NewReader(VipnodePoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VipnodePoolBin), backend, _operator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VipnodePool{VipnodePoolCaller: VipnodePoolCaller{contract: contract}, VipnodePoolTransactor: VipnodePoolTransactor{contract: contract}, VipnodePoolFilterer: VipnodePoolFilterer{contract: contract}}, nil
}

// VipnodePool is an auto generated Go binding around an Ethereum contract.
type VipnodePool struct {
	VipnodePoolCaller     // Read-only binding to the contract
	VipnodePoolTransactor // Write-only binding to the contract
	VipnodePoolFilterer   // Log filterer for contract events
}

// VipnodePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type VipnodePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VipnodePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VipnodePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VipnodePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VipnodePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VipnodePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VipnodePoolSession struct {
	Contract     *VipnodePool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VipnodePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VipnodePoolCallerSession struct {
	Contract *VipnodePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VipnodePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VipnodePoolTransactorSession struct {
	Contract     *VipnodePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VipnodePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type VipnodePoolRaw struct {
	Contract *VipnodePool // Generic contract binding to access the raw methods on
}

// VipnodePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VipnodePoolCallerRaw struct {
	Contract *VipnodePoolCaller // Generic read-only contract binding to access the raw methods on
}

// VipnodePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VipnodePoolTransactorRaw struct {
	Contract *VipnodePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVipnodePool creates a new instance of VipnodePool, bound to a specific deployed contract.
func NewVipnodePool(address common.Address, backend bind.ContractBackend) (*VipnodePool, error) {
	contract, err := bindVipnodePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VipnodePool{VipnodePoolCaller: VipnodePoolCaller{contract: contract}, VipnodePoolTransactor: VipnodePoolTransactor{contract: contract}, VipnodePoolFilterer: VipnodePoolFilterer{contract: contract}}, nil
}

// NewVipnodePoolCaller creates a new read-only instance of VipnodePool, bound to a specific deployed contract.
func NewVipnodePoolCaller(address common.Address, caller bind.ContractCaller) (*VipnodePoolCaller, error) {
	contract, err := bindVipnodePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VipnodePoolCaller{contract: contract}, nil
}

// NewVipnodePoolTransactor creates a new write-only instance of VipnodePool, bound to a specific deployed contract.
func NewVipnodePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*VipnodePoolTransactor, error) {
	contract, err := bindVipnodePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VipnodePoolTransactor{contract: contract}, nil
}

// NewVipnodePoolFilterer creates a new log filterer instance of VipnodePool, bound to a specific deployed contract.
func NewVipnodePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*VipnodePoolFilterer, error) {
	contract, err := bindVipnodePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VipnodePoolFilterer{contract: contract}, nil
}

// bindVipnodePool binds a generic wrapper to an already deployed contract.
func bindVipnodePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VipnodePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VipnodePool *VipnodePoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VipnodePool.Contract.VipnodePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VipnodePool *VipnodePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VipnodePool.Contract.VipnodePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VipnodePool *VipnodePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VipnodePool.Contract.VipnodePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VipnodePool *VipnodePoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VipnodePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VipnodePool *VipnodePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VipnodePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VipnodePool *VipnodePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VipnodePool.Contract.contract.Transact(opts, method, params...)
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts( address) constant returns(balance uint256, timeLocked uint256)
func (_VipnodePool *VipnodePoolCaller) Accounts(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance    *big.Int
	TimeLocked *big.Int
}, error) {
	ret := new(struct {
		Balance    *big.Int
		TimeLocked *big.Int
	})
	out := ret
	err := _VipnodePool.contract.Call(opts, out, "accounts", arg0)
	return *ret, err
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts( address) constant returns(balance uint256, timeLocked uint256)
func (_VipnodePool *VipnodePoolSession) Accounts(arg0 common.Address) (struct {
	Balance    *big.Int
	TimeLocked *big.Int
}, error) {
	return _VipnodePool.Contract.Accounts(&_VipnodePool.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts( address) constant returns(balance uint256, timeLocked uint256)
func (_VipnodePool *VipnodePoolCallerSession) Accounts(arg0 common.Address) (struct {
	Balance    *big.Int
	TimeLocked *big.Int
}, error) {
	return _VipnodePool.Contract.Accounts(&_VipnodePool.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_VipnodePool *VipnodePoolCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VipnodePool.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_VipnodePool *VipnodePoolSession) Operator() (common.Address, error) {
	return _VipnodePool.Contract.Operator(&_VipnodePool.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_VipnodePool *VipnodePoolCallerSession) Operator() (common.Address, error) {
	return _VipnodePool.Contract.Operator(&_VipnodePool.CallOpts)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() constant returns(uint256)
func (_VipnodePool *VipnodePoolCaller) WithdrawInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VipnodePool.contract.Call(opts, out, "withdrawInterval")
	return *ret0, err
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() constant returns(uint256)
func (_VipnodePool *VipnodePoolSession) WithdrawInterval() (*big.Int, error) {
	return _VipnodePool.Contract.WithdrawInterval(&_VipnodePool.CallOpts)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() constant returns(uint256)
func (_VipnodePool *VipnodePoolCallerSession) WithdrawInterval() (*big.Int, error) {
	return _VipnodePool.Contract.WithdrawInterval(&_VipnodePool.CallOpts)
}

// AddBalance is a paid mutator transaction binding the contract method 0xb163cc38.
//
// Solidity: function addBalance() returns()
func (_VipnodePool *VipnodePoolTransactor) AddBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VipnodePool.contract.Transact(opts, "addBalance")
}

// AddBalance is a paid mutator transaction binding the contract method 0xb163cc38.
//
// Solidity: function addBalance() returns()
func (_VipnodePool *VipnodePoolSession) AddBalance() (*types.Transaction, error) {
	return _VipnodePool.Contract.AddBalance(&_VipnodePool.TransactOpts)
}

// AddBalance is a paid mutator transaction binding the contract method 0xb163cc38.
//
// Solidity: function addBalance() returns()
func (_VipnodePool *VipnodePoolTransactorSession) AddBalance() (*types.Transaction, error) {
	return _VipnodePool.Contract.AddBalance(&_VipnodePool.TransactOpts)
}

// ForceSettle is a paid mutator transaction binding the contract method 0xaeea1f7b.
//
// Solidity: function forceSettle() returns()
func (_VipnodePool *VipnodePoolTransactor) ForceSettle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VipnodePool.contract.Transact(opts, "forceSettle")
}

// ForceSettle is a paid mutator transaction binding the contract method 0xaeea1f7b.
//
// Solidity: function forceSettle() returns()
func (_VipnodePool *VipnodePoolSession) ForceSettle() (*types.Transaction, error) {
	return _VipnodePool.Contract.ForceSettle(&_VipnodePool.TransactOpts)
}

// ForceSettle is a paid mutator transaction binding the contract method 0xaeea1f7b.
//
// Solidity: function forceSettle() returns()
func (_VipnodePool *VipnodePoolTransactorSession) ForceSettle() (*types.Transaction, error) {
	return _VipnodePool.Contract.ForceSettle(&_VipnodePool.TransactOpts)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x7be80b39.
//
// Solidity: function forceWithdraw() returns()
func (_VipnodePool *VipnodePoolTransactor) ForceWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VipnodePool.contract.Transact(opts, "forceWithdraw")
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x7be80b39.
//
// Solidity: function forceWithdraw() returns()
func (_VipnodePool *VipnodePoolSession) ForceWithdraw() (*types.Transaction, error) {
	return _VipnodePool.Contract.ForceWithdraw(&_VipnodePool.TransactOpts)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x7be80b39.
//
// Solidity: function forceWithdraw() returns()
func (_VipnodePool *VipnodePoolTransactorSession) ForceWithdraw() (*types.Transaction, error) {
	return _VipnodePool.Contract.ForceWithdraw(&_VipnodePool.TransactOpts)
}

// OpSettle is a paid mutator transaction binding the contract method 0xb27d1ce3.
//
// Solidity: function opSettle(_account address, _releaseAmount uint256, _newBalance uint256) returns()
func (_VipnodePool *VipnodePoolTransactor) OpSettle(opts *bind.TransactOpts, _account common.Address, _releaseAmount *big.Int, _newBalance *big.Int) (*types.Transaction, error) {
	return _VipnodePool.contract.Transact(opts, "opSettle", _account, _releaseAmount, _newBalance)
}

// OpSettle is a paid mutator transaction binding the contract method 0xb27d1ce3.
//
// Solidity: function opSettle(_account address, _releaseAmount uint256, _newBalance uint256) returns()
func (_VipnodePool *VipnodePoolSession) OpSettle(_account common.Address, _releaseAmount *big.Int, _newBalance *big.Int) (*types.Transaction, error) {
	return _VipnodePool.Contract.OpSettle(&_VipnodePool.TransactOpts, _account, _releaseAmount, _newBalance)
}

// OpSettle is a paid mutator transaction binding the contract method 0xb27d1ce3.
//
// Solidity: function opSettle(_account address, _releaseAmount uint256, _newBalance uint256) returns()
func (_VipnodePool *VipnodePoolTransactorSession) OpSettle(_account common.Address, _releaseAmount *big.Int, _newBalance *big.Int) (*types.Transaction, error) {
	return _VipnodePool.Contract.OpSettle(&_VipnodePool.TransactOpts, _account, _releaseAmount, _newBalance)
}

// OpWithdraw is a paid mutator transaction binding the contract method 0xae5fa7e7.
//
// Solidity: function opWithdraw(_amount uint256) returns()
func (_VipnodePool *VipnodePoolTransactor) OpWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _VipnodePool.contract.Transact(opts, "opWithdraw", _amount)
}

// OpWithdraw is a paid mutator transaction binding the contract method 0xae5fa7e7.
//
// Solidity: function opWithdraw(_amount uint256) returns()
func (_VipnodePool *VipnodePoolSession) OpWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _VipnodePool.Contract.OpWithdraw(&_VipnodePool.TransactOpts, _amount)
}

// OpWithdraw is a paid mutator transaction binding the contract method 0xae5fa7e7.
//
// Solidity: function opWithdraw(_amount uint256) returns()
func (_VipnodePool *VipnodePoolTransactorSession) OpWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _VipnodePool.Contract.OpWithdraw(&_VipnodePool.TransactOpts, _amount)
}

// VipnodePoolBalanceIterator is returned from FilterBalance and is used to iterate over the raw logs and unpacked data for Balance events raised by the VipnodePool contract.
type VipnodePoolBalanceIterator struct {
	Event *VipnodePoolBalance // Event containing the contract specifics and raw log

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
func (it *VipnodePoolBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VipnodePoolBalance)
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
		it.Event = new(VipnodePoolBalance)
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
func (it *VipnodePoolBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VipnodePoolBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VipnodePoolBalance represents a Balance event raised by the VipnodePool contract.
type VipnodePoolBalance struct {
	Account common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBalance is a free log retrieval operation binding the contract event 0x134e340554ff8a7d64280a2a28b982df554e2595e5bf45cd39368f31099172a6.
//
// Solidity: e Balance(account address, balance uint256)
func (_VipnodePool *VipnodePoolFilterer) FilterBalance(opts *bind.FilterOpts) (*VipnodePoolBalanceIterator, error) {

	logs, sub, err := _VipnodePool.contract.FilterLogs(opts, "Balance")
	if err != nil {
		return nil, err
	}
	return &VipnodePoolBalanceIterator{contract: _VipnodePool.contract, event: "Balance", logs: logs, sub: sub}, nil
}

// WatchBalance is a free log subscription operation binding the contract event 0x134e340554ff8a7d64280a2a28b982df554e2595e5bf45cd39368f31099172a6.
//
// Solidity: e Balance(account address, balance uint256)
func (_VipnodePool *VipnodePoolFilterer) WatchBalance(opts *bind.WatchOpts, sink chan<- *VipnodePoolBalance) (event.Subscription, error) {

	logs, sub, err := _VipnodePool.contract.WatchLogs(opts, "Balance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VipnodePoolBalance)
				if err := _VipnodePool.contract.UnpackLog(event, "Balance", log); err != nil {
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

// VipnodePoolForceSettleIterator is returned from FilterForceSettle and is used to iterate over the raw logs and unpacked data for ForceSettle events raised by the VipnodePool contract.
type VipnodePoolForceSettleIterator struct {
	Event *VipnodePoolForceSettle // Event containing the contract specifics and raw log

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
func (it *VipnodePoolForceSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VipnodePoolForceSettle)
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
		it.Event = new(VipnodePoolForceSettle)
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
func (it *VipnodePoolForceSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VipnodePoolForceSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VipnodePoolForceSettle represents a ForceSettle event raised by the VipnodePool contract.
type VipnodePoolForceSettle struct {
	Account    common.Address
	TimeLocked *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterForceSettle is a free log retrieval operation binding the contract event 0xde112f62320281f04efec46e09ed286851fdd0b428dfc7214fe0e64f9364323b.
//
// Solidity: e ForceSettle(account address, timeLocked uint256)
func (_VipnodePool *VipnodePoolFilterer) FilterForceSettle(opts *bind.FilterOpts) (*VipnodePoolForceSettleIterator, error) {

	logs, sub, err := _VipnodePool.contract.FilterLogs(opts, "ForceSettle")
	if err != nil {
		return nil, err
	}
	return &VipnodePoolForceSettleIterator{contract: _VipnodePool.contract, event: "ForceSettle", logs: logs, sub: sub}, nil
}

// WatchForceSettle is a free log subscription operation binding the contract event 0xde112f62320281f04efec46e09ed286851fdd0b428dfc7214fe0e64f9364323b.
//
// Solidity: e ForceSettle(account address, timeLocked uint256)
func (_VipnodePool *VipnodePoolFilterer) WatchForceSettle(opts *bind.WatchOpts, sink chan<- *VipnodePoolForceSettle) (event.Subscription, error) {

	logs, sub, err := _VipnodePool.contract.WatchLogs(opts, "ForceSettle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VipnodePoolForceSettle)
				if err := _VipnodePool.contract.UnpackLog(event, "ForceSettle", log); err != nil {
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
