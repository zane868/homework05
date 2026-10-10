// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// BeggingContractMetaData contains all meta data concerning the BeggingContract contract.
var BeggingContractMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoBalanceToWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDonated\",\"type\":\"uint256\"}],\"name\":\"Donated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"GetDonation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "BeggingContract",
}

// BeggingContract is an auto generated Go binding around an Ethereum contract.
type BeggingContract struct {
	abi abi.ABI
}

// GetABI returns the ABI associated with this contract binding.
func (c *BeggingContract) GetABI() abi.ABI {
	return c.abi
}

// NewBeggingContract creates a new instance of BeggingContract.
func NewBeggingContract() *BeggingContract {
	parsed, err := BeggingContractMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &BeggingContract{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *BeggingContract) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDonate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd917deb5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function Donate() payable returns()
func (beggingContract *BeggingContract) PackDonate() []byte {
	enc, err := beggingContract.abi.Pack("Donate")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDonate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd917deb5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function Donate() payable returns()
func (beggingContract *BeggingContract) TryPackDonate() ([]byte, error) {
	return beggingContract.abi.Pack("Donate")
}

// PackGetDonation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x21cf4a12.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function GetDonation(address _address) view returns(uint256)
func (beggingContract *BeggingContract) PackGetDonation(address common.Address) []byte {
	enc, err := beggingContract.abi.Pack("GetDonation", address)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetDonation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x21cf4a12.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function GetDonation(address _address) view returns(uint256)
func (beggingContract *BeggingContract) TryPackGetDonation(address common.Address) ([]byte, error) {
	return beggingContract.abi.Pack("GetDonation", address)
}

// UnpackGetDonation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x21cf4a12.
//
// Solidity: function GetDonation(address _address) view returns(uint256)
func (beggingContract *BeggingContract) UnpackGetDonation(data []byte) (*big.Int, error) {
	out, err := beggingContract.abi.Unpack("GetDonation", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (beggingContract *BeggingContract) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := beggingContract.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (beggingContract *BeggingContract) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return beggingContract.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (beggingContract *BeggingContract) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := beggingContract.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf6761151.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function Withdraw(address recipient) returns()
func (beggingContract *BeggingContract) PackWithdraw(recipient common.Address) []byte {
	enc, err := beggingContract.abi.Pack("Withdraw", recipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf6761151.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function Withdraw(address recipient) returns()
func (beggingContract *BeggingContract) TryPackWithdraw(recipient common.Address) ([]byte, error) {
	return beggingContract.abi.Pack("Withdraw", recipient)
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address initialOwner) returns()
func (beggingContract *BeggingContract) PackInitialize(initialOwner common.Address) []byte {
	enc, err := beggingContract.abi.Pack("initialize", initialOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(address initialOwner) returns()
func (beggingContract *BeggingContract) TryPackInitialize(initialOwner common.Address) ([]byte, error) {
	return beggingContract.abi.Pack("initialize", initialOwner)
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (beggingContract *BeggingContract) PackOwner() []byte {
	enc, err := beggingContract.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (beggingContract *BeggingContract) TryPackOwner() ([]byte, error) {
	return beggingContract.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (beggingContract *BeggingContract) UnpackOwner(data []byte) (common.Address, error) {
	out, err := beggingContract.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (beggingContract *BeggingContract) PackProxiableUUID() []byte {
	enc, err := beggingContract.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (beggingContract *BeggingContract) TryPackProxiableUUID() ([]byte, error) {
	return beggingContract.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (beggingContract *BeggingContract) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := beggingContract.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (beggingContract *BeggingContract) PackRenounceOwnership() []byte {
	enc, err := beggingContract.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceOwnership() returns()
func (beggingContract *BeggingContract) TryPackRenounceOwnership() ([]byte, error) {
	return beggingContract.abi.Pack("renounceOwnership")
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (beggingContract *BeggingContract) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := beggingContract.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (beggingContract *BeggingContract) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return beggingContract.abi.Pack("transferOwnership", newOwner)
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (beggingContract *BeggingContract) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := beggingContract.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (beggingContract *BeggingContract) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return beggingContract.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// BeggingContractDonated represents a Donated event raised by the BeggingContract contract.
type BeggingContractDonated struct {
	Donor        common.Address
	Amount       *big.Int
	TotalDonated *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const BeggingContractDonatedEventName = "Donated"

// ContractEventName returns the user-defined event name.
func (BeggingContractDonated) ContractEventName() string {
	return BeggingContractDonatedEventName
}

// UnpackDonatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Donated(address indexed donor, uint256 amount, uint256 totalDonated)
func (beggingContract *BeggingContract) UnpackDonatedEvent(log *types.Log) (*BeggingContractDonated, error) {
	event := "Donated"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != beggingContract.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(BeggingContractDonated)
	if len(log.Data) > 0 {
		if err := beggingContract.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range beggingContract.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BeggingContractInitialized represents a Initialized event raised by the BeggingContract contract.
type BeggingContractInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const BeggingContractInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (BeggingContractInitialized) ContractEventName() string {
	return BeggingContractInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (beggingContract *BeggingContract) UnpackInitializedEvent(log *types.Log) (*BeggingContractInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != beggingContract.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(BeggingContractInitialized)
	if len(log.Data) > 0 {
		if err := beggingContract.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range beggingContract.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BeggingContractOwnershipTransferred represents a OwnershipTransferred event raised by the BeggingContract contract.
type BeggingContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const BeggingContractOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (BeggingContractOwnershipTransferred) ContractEventName() string {
	return BeggingContractOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (beggingContract *BeggingContract) UnpackOwnershipTransferredEvent(log *types.Log) (*BeggingContractOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != beggingContract.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(BeggingContractOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := beggingContract.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range beggingContract.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BeggingContractUpgraded represents a Upgraded event raised by the BeggingContract contract.
type BeggingContractUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const BeggingContractUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (BeggingContractUpgraded) ContractEventName() string {
	return BeggingContractUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (beggingContract *BeggingContract) UnpackUpgradedEvent(log *types.Log) (*BeggingContractUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != beggingContract.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(BeggingContractUpgraded)
	if len(log.Data) > 0 {
		if err := beggingContract.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range beggingContract.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BeggingContractWithdrawn represents a Withdrawn event raised by the BeggingContract contract.
type BeggingContractWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const BeggingContractWithdrawnEventName = "Withdrawn"

// ContractEventName returns the user-defined event name.
func (BeggingContractWithdrawn) ContractEventName() string {
	return BeggingContractWithdrawnEventName
}

// UnpackWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Withdrawn(address indexed recipient, uint256 amount)
func (beggingContract *BeggingContract) UnpackWithdrawnEvent(log *types.Log) (*BeggingContractWithdrawn, error) {
	event := "Withdrawn"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != beggingContract.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(BeggingContractWithdrawn)
	if len(log.Data) > 0 {
		if err := beggingContract.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range beggingContract.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (beggingContract *BeggingContract) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], beggingContract.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return beggingContract.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], beggingContract.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return beggingContract.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], beggingContract.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return beggingContract.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], beggingContract.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return beggingContract.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], beggingContract.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return beggingContract.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], beggingContract.abi.Errors["NoBalanceToWithdraw"].ID.Bytes()[:4]) {
		return beggingContract.UnpackNoBalanceToWithdrawError(raw[4:])
	}
	if bytes.Equal(raw[:4], beggingContract.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return beggingContract.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], beggingContract.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return beggingContract.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], beggingContract.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return beggingContract.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], beggingContract.abi.Errors["TransferFailed"].ID.Bytes()[:4]) {
		return beggingContract.UnpackTransferFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], beggingContract.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return beggingContract.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], beggingContract.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return beggingContract.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// BeggingContractAddressEmptyCode represents a AddressEmptyCode error raised by the BeggingContract contract.
type BeggingContractAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func BeggingContractAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (beggingContract *BeggingContract) UnpackAddressEmptyCodeError(raw []byte) (*BeggingContractAddressEmptyCode, error) {
	out := new(BeggingContractAddressEmptyCode)
	if err := beggingContract.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BeggingContractERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the BeggingContract contract.
type BeggingContractERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func BeggingContractERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (beggingContract *BeggingContract) UnpackERC1967InvalidImplementationError(raw []byte) (*BeggingContractERC1967InvalidImplementation, error) {
	out := new(BeggingContractERC1967InvalidImplementation)
	if err := beggingContract.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BeggingContractERC1967NonPayable represents a ERC1967NonPayable error raised by the BeggingContract contract.
type BeggingContractERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func BeggingContractERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (beggingContract *BeggingContract) UnpackERC1967NonPayableError(raw []byte) (*BeggingContractERC1967NonPayable, error) {
	out := new(BeggingContractERC1967NonPayable)
	if err := beggingContract.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BeggingContractFailedCall represents a FailedCall error raised by the BeggingContract contract.
type BeggingContractFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func BeggingContractFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (beggingContract *BeggingContract) UnpackFailedCallError(raw []byte) (*BeggingContractFailedCall, error) {
	out := new(BeggingContractFailedCall)
	if err := beggingContract.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BeggingContractInvalidInitialization represents a InvalidInitialization error raised by the BeggingContract contract.
type BeggingContractInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func BeggingContractInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (beggingContract *BeggingContract) UnpackInvalidInitializationError(raw []byte) (*BeggingContractInvalidInitialization, error) {
	out := new(BeggingContractInvalidInitialization)
	if err := beggingContract.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BeggingContractNoBalanceToWithdraw represents a NoBalanceToWithdraw error raised by the BeggingContract contract.
type BeggingContractNoBalanceToWithdraw struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoBalanceToWithdraw()
func BeggingContractNoBalanceToWithdrawErrorID() common.Hash {
	return common.HexToHash("0xbbd817087f4149ecf85d6d8465bfb82c4318225eec681f0abdb3144233370d14")
}

// UnpackNoBalanceToWithdrawError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoBalanceToWithdraw()
func (beggingContract *BeggingContract) UnpackNoBalanceToWithdrawError(raw []byte) (*BeggingContractNoBalanceToWithdraw, error) {
	out := new(BeggingContractNoBalanceToWithdraw)
	if err := beggingContract.abi.UnpackIntoInterface(out, "NoBalanceToWithdraw", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BeggingContractNotInitializing represents a NotInitializing error raised by the BeggingContract contract.
type BeggingContractNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func BeggingContractNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (beggingContract *BeggingContract) UnpackNotInitializingError(raw []byte) (*BeggingContractNotInitializing, error) {
	out := new(BeggingContractNotInitializing)
	if err := beggingContract.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BeggingContractOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the BeggingContract contract.
type BeggingContractOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func BeggingContractOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (beggingContract *BeggingContract) UnpackOwnableInvalidOwnerError(raw []byte) (*BeggingContractOwnableInvalidOwner, error) {
	out := new(BeggingContractOwnableInvalidOwner)
	if err := beggingContract.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BeggingContractOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the BeggingContract contract.
type BeggingContractOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func BeggingContractOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (beggingContract *BeggingContract) UnpackOwnableUnauthorizedAccountError(raw []byte) (*BeggingContractOwnableUnauthorizedAccount, error) {
	out := new(BeggingContractOwnableUnauthorizedAccount)
	if err := beggingContract.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BeggingContractTransferFailed represents a TransferFailed error raised by the BeggingContract contract.
type BeggingContractTransferFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferFailed()
func BeggingContractTransferFailedErrorID() common.Hash {
	return common.HexToHash("0x90b8ec1877afffd816d05d9b13947f3ff18ec5851c38bad15ec2b710f92391b1")
}

// UnpackTransferFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferFailed()
func (beggingContract *BeggingContract) UnpackTransferFailedError(raw []byte) (*BeggingContractTransferFailed, error) {
	out := new(BeggingContractTransferFailed)
	if err := beggingContract.abi.UnpackIntoInterface(out, "TransferFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BeggingContractUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the BeggingContract contract.
type BeggingContractUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func BeggingContractUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (beggingContract *BeggingContract) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*BeggingContractUUPSUnauthorizedCallContext, error) {
	out := new(BeggingContractUUPSUnauthorizedCallContext)
	if err := beggingContract.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BeggingContractUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the BeggingContract contract.
type BeggingContractUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func BeggingContractUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (beggingContract *BeggingContract) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*BeggingContractUUPSUnsupportedProxiableUUID, error) {
	out := new(BeggingContractUUPSUnsupportedProxiableUUID)
	if err := beggingContract.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
