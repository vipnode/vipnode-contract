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
const VipnodePoolABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"withdrawInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_client\",\"type\":\"address\"},{\"name\":\"_minBalance\",\"type\":\"uint256\"}],\"name\":\"checkBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"clients\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"timeLocked\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"forceWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_client\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_release\",\"type\":\"bool\"}],\"name\":\"opSettle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"opWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"forceSettle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addBalance\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timeLocked\",\"type\":\"uint256\"}],\"name\":\"ForceSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Balance\",\"type\":\"event\"}]"

// VipnodePoolBin is the compiled bytecode used for deploying new contracts.
const VipnodePoolBin = `0x608060405262093a8060015534801561001757600080fd5b506040516020806105b58339810160405251600160a060020a038116151561003e57600080fd5b60008054600160a060020a03909216600160a060020a03199092169190911790556105478061006e6000396000f3006080604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663162075d8811461009d5780634d7b9bd5146100c4578063570ca735146100fc57806361eed2a91461012d5780637be80b3914610167578063a73e679a1461017e578063ae5fa7e7146101a7578063aeea1f7b146101bf578063b163cc38146101d4575b600080fd5b3480156100a957600080fd5b506100b26101dc565b60408051918252519081900360200190f35b3480156100d057600080fd5b506100e8600160a060020a03600435166024356101e2565b604080519115158252519081900360200190f35b34801561010857600080fd5b5061011161022c565b60408051600160a060020a039092168252519081900360200190f35b34801561013957600080fd5b5061014e600160a060020a036004351661023b565b6040805192835260208301919091528051918290030190f35b34801561017357600080fd5b5061017c610254565b005b34801561018a57600080fd5b5061017c600160a060020a03600435166024356044351515610311565b3480156101b357600080fd5b5061017c6004356103f3565b3480156101cb57600080fd5b5061017c610447565b61017c6104ae565b60015481565b60006101ec610504565b5050600160a060020a0391909116600090815260026020908152604091829020825180840190935280548084526001909101549290910191909152101590565b600054600160a060020a031681565b6002602052600090815260409020805460019091015482565b3360009081526002602052604081208054909190811061027357600080fd5b600182015460001061028457600080fd5b600182015442101561029557600080fd5b508054600080835560018301819055604051339183156108fc02918491818181858888f193505050501580156102cf573d6000803e3d6000fd5b50815460408051338152602081019290925280517f134e340554ff8a7d64280a2a28b982df554e2595e5bf45cd39368f31099172a69281900390910190a15050565b600080548190600160a060020a0316331461032b57600080fd5b600160a060020a0385166000908152600260205260409020805490925084111561035457600080fd5b8154849003825582156103a657508054600080835560018301819055604051600160a060020a0387169183156108fc02918491818181858888f193505050501580156103a4573d6000803e3d6000fd5b505b815460408051600160a060020a0388168152602081019290925280517f134e340554ff8a7d64280a2a28b982df554e2595e5bf45cd39368f31099172a69281900390910190a15050505050565b600054600160a060020a0316331461040a57600080fd5b60008054604051600160a060020a039091169183156108fc02918491818181858888f19350505050158015610443573d6000803e3d6000fd5b5050565b336000908152600260205260408120805490911061046457600080fd5b60018054420190820181905560408051338152602081019290925280517fde112f62320281f04efec46e09ed286851fdd0b428dfc7214fe0e64f9364323b9281900390910190a150565b3360008181526002602090815260409182902080543401808255835194855291840191909152815190927f134e340554ff8a7d64280a2a28b982df554e2595e5bf45cd39368f31099172a692908290030190a150565b6040805180820190915260008082526020820152905600a165627a7a723058204151eb23a9a4473dd70413d4ee05dda6a34eb5dea2dd9b0034226efdba3cbd340029`

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

// CheckBalance is a free data retrieval call binding the contract method 0x4d7b9bd5.
//
// Solidity: function checkBalance(_client address, _minBalance uint256) constant returns(bool)
func (_VipnodePool *VipnodePoolCaller) CheckBalance(opts *bind.CallOpts, _client common.Address, _minBalance *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VipnodePool.contract.Call(opts, out, "checkBalance", _client, _minBalance)
	return *ret0, err
}

// CheckBalance is a free data retrieval call binding the contract method 0x4d7b9bd5.
//
// Solidity: function checkBalance(_client address, _minBalance uint256) constant returns(bool)
func (_VipnodePool *VipnodePoolSession) CheckBalance(_client common.Address, _minBalance *big.Int) (bool, error) {
	return _VipnodePool.Contract.CheckBalance(&_VipnodePool.CallOpts, _client, _minBalance)
}

// CheckBalance is a free data retrieval call binding the contract method 0x4d7b9bd5.
//
// Solidity: function checkBalance(_client address, _minBalance uint256) constant returns(bool)
func (_VipnodePool *VipnodePoolCallerSession) CheckBalance(_client common.Address, _minBalance *big.Int) (bool, error) {
	return _VipnodePool.Contract.CheckBalance(&_VipnodePool.CallOpts, _client, _minBalance)
}

// Clients is a free data retrieval call binding the contract method 0x61eed2a9.
//
// Solidity: function clients( address) constant returns(balance uint256, timeLocked uint256)
func (_VipnodePool *VipnodePoolCaller) Clients(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance    *big.Int
	TimeLocked *big.Int
}, error) {
	ret := new(struct {
		Balance    *big.Int
		TimeLocked *big.Int
	})
	out := ret
	err := _VipnodePool.contract.Call(opts, out, "clients", arg0)
	return *ret, err
}

// Clients is a free data retrieval call binding the contract method 0x61eed2a9.
//
// Solidity: function clients( address) constant returns(balance uint256, timeLocked uint256)
func (_VipnodePool *VipnodePoolSession) Clients(arg0 common.Address) (struct {
	Balance    *big.Int
	TimeLocked *big.Int
}, error) {
	return _VipnodePool.Contract.Clients(&_VipnodePool.CallOpts, arg0)
}

// Clients is a free data retrieval call binding the contract method 0x61eed2a9.
//
// Solidity: function clients( address) constant returns(balance uint256, timeLocked uint256)
func (_VipnodePool *VipnodePoolCallerSession) Clients(arg0 common.Address) (struct {
	Balance    *big.Int
	TimeLocked *big.Int
}, error) {
	return _VipnodePool.Contract.Clients(&_VipnodePool.CallOpts, arg0)
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

// OpSettle is a paid mutator transaction binding the contract method 0xa73e679a.
//
// Solidity: function opSettle(_client address, _amount uint256, _release bool) returns()
func (_VipnodePool *VipnodePoolTransactor) OpSettle(opts *bind.TransactOpts, _client common.Address, _amount *big.Int, _release bool) (*types.Transaction, error) {
	return _VipnodePool.contract.Transact(opts, "opSettle", _client, _amount, _release)
}

// OpSettle is a paid mutator transaction binding the contract method 0xa73e679a.
//
// Solidity: function opSettle(_client address, _amount uint256, _release bool) returns()
func (_VipnodePool *VipnodePoolSession) OpSettle(_client common.Address, _amount *big.Int, _release bool) (*types.Transaction, error) {
	return _VipnodePool.Contract.OpSettle(&_VipnodePool.TransactOpts, _client, _amount, _release)
}

// OpSettle is a paid mutator transaction binding the contract method 0xa73e679a.
//
// Solidity: function opSettle(_client address, _amount uint256, _release bool) returns()
func (_VipnodePool *VipnodePoolTransactorSession) OpSettle(_client common.Address, _amount *big.Int, _release bool) (*types.Transaction, error) {
	return _VipnodePool.Contract.OpSettle(&_VipnodePool.TransactOpts, _client, _amount, _release)
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
	Client  common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBalance is a free log retrieval operation binding the contract event 0x134e340554ff8a7d64280a2a28b982df554e2595e5bf45cd39368f31099172a6.
//
// Solidity: e Balance(client address, balance uint256)
func (_VipnodePool *VipnodePoolFilterer) FilterBalance(opts *bind.FilterOpts) (*VipnodePoolBalanceIterator, error) {

	logs, sub, err := _VipnodePool.contract.FilterLogs(opts, "Balance")
	if err != nil {
		return nil, err
	}
	return &VipnodePoolBalanceIterator{contract: _VipnodePool.contract, event: "Balance", logs: logs, sub: sub}, nil
}

// WatchBalance is a free log subscription operation binding the contract event 0x134e340554ff8a7d64280a2a28b982df554e2595e5bf45cd39368f31099172a6.
//
// Solidity: e Balance(client address, balance uint256)
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
	Client     common.Address
	TimeLocked *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterForceSettle is a free log retrieval operation binding the contract event 0xde112f62320281f04efec46e09ed286851fdd0b428dfc7214fe0e64f9364323b.
//
// Solidity: e ForceSettle(client address, timeLocked uint256)
func (_VipnodePool *VipnodePoolFilterer) FilterForceSettle(opts *bind.FilterOpts) (*VipnodePoolForceSettleIterator, error) {

	logs, sub, err := _VipnodePool.contract.FilterLogs(opts, "ForceSettle")
	if err != nil {
		return nil, err
	}
	return &VipnodePoolForceSettleIterator{contract: _VipnodePool.contract, event: "ForceSettle", logs: logs, sub: sub}, nil
}

// WatchForceSettle is a free log subscription operation binding the contract event 0xde112f62320281f04efec46e09ed286851fdd0b428dfc7214fe0e64f9364323b.
//
// Solidity: e ForceSettle(client address, timeLocked uint256)
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
