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

// LOCGamePlayNFTMetaData contains all meta data concerning the LOCGamePlayNFT contract.
var LOCGamePlayNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"BaseURISet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"extractCardId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cardId\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cardId\",\"type\":\"uint256\"}],\"name\":\"extractNumberOfCard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenCounter\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cardId\",\"type\":\"uint256\"}],\"name\":\"getCardTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cardTotalSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LOCGamePlayNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use LOCGamePlayNFTMetaData.ABI instead.
var LOCGamePlayNFTABI = LOCGamePlayNFTMetaData.ABI

// LOCGamePlayNFT is an auto generated Go binding around an Ethereum contract.
type LOCGamePlayNFT struct {
	LOCGamePlayNFTCaller     // Read-only binding to the contract
	LOCGamePlayNFTTransactor // Write-only binding to the contract
	LOCGamePlayNFTFilterer   // Log filterer for contract events
}

// LOCGamePlayNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type LOCGamePlayNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LOCGamePlayNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LOCGamePlayNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LOCGamePlayNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LOCGamePlayNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LOCGamePlayNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LOCGamePlayNFTSession struct {
	Contract     *LOCGamePlayNFT   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LOCGamePlayNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LOCGamePlayNFTCallerSession struct {
	Contract *LOCGamePlayNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LOCGamePlayNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LOCGamePlayNFTTransactorSession struct {
	Contract     *LOCGamePlayNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LOCGamePlayNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type LOCGamePlayNFTRaw struct {
	Contract *LOCGamePlayNFT // Generic contract binding to access the raw methods on
}

// LOCGamePlayNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LOCGamePlayNFTCallerRaw struct {
	Contract *LOCGamePlayNFTCaller // Generic read-only contract binding to access the raw methods on
}

// LOCGamePlayNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LOCGamePlayNFTTransactorRaw struct {
	Contract *LOCGamePlayNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLOCGamePlayNFT creates a new instance of LOCGamePlayNFT, bound to a specific deployed contract.
func NewLOCGamePlayNFT(address common.Address, backend bind.ContractBackend) (*LOCGamePlayNFT, error) {
	contract, err := bindLOCGamePlayNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LOCGamePlayNFT{LOCGamePlayNFTCaller: LOCGamePlayNFTCaller{contract: contract}, LOCGamePlayNFTTransactor: LOCGamePlayNFTTransactor{contract: contract}, LOCGamePlayNFTFilterer: LOCGamePlayNFTFilterer{contract: contract}}, nil
}

// NewLOCGamePlayNFTCaller creates a new read-only instance of LOCGamePlayNFT, bound to a specific deployed contract.
func NewLOCGamePlayNFTCaller(address common.Address, caller bind.ContractCaller) (*LOCGamePlayNFTCaller, error) {
	contract, err := bindLOCGamePlayNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LOCGamePlayNFTCaller{contract: contract}, nil
}

// NewLOCGamePlayNFTTransactor creates a new write-only instance of LOCGamePlayNFT, bound to a specific deployed contract.
func NewLOCGamePlayNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*LOCGamePlayNFTTransactor, error) {
	contract, err := bindLOCGamePlayNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LOCGamePlayNFTTransactor{contract: contract}, nil
}

// NewLOCGamePlayNFTFilterer creates a new log filterer instance of LOCGamePlayNFT, bound to a specific deployed contract.
func NewLOCGamePlayNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*LOCGamePlayNFTFilterer, error) {
	contract, err := bindLOCGamePlayNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LOCGamePlayNFTFilterer{contract: contract}, nil
}

// bindLOCGamePlayNFT binds a generic wrapper to an already deployed contract.
func bindLOCGamePlayNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LOCGamePlayNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LOCGamePlayNFT *LOCGamePlayNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LOCGamePlayNFT.Contract.LOCGamePlayNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LOCGamePlayNFT *LOCGamePlayNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.LOCGamePlayNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LOCGamePlayNFT *LOCGamePlayNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.LOCGamePlayNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LOCGamePlayNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LOCGamePlayNFT.Contract.DEFAULTADMINROLE(&_LOCGamePlayNFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LOCGamePlayNFT.Contract.DEFAULTADMINROLE(&_LOCGamePlayNFT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) MINTERROLE() ([32]byte, error) {
	return _LOCGamePlayNFT.Contract.MINTERROLE(&_LOCGamePlayNFT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) MINTERROLE() ([32]byte, error) {
	return _LOCGamePlayNFT.Contract.MINTERROLE(&_LOCGamePlayNFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.BalanceOf(&_LOCGamePlayNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.BalanceOf(&_LOCGamePlayNFT.CallOpts, owner)
}

// ExtractCardId is a free data retrieval call binding the contract method 0xe03d2b47.
//
// Solidity: function extractCardId(uint256 tokenId) pure returns(uint256 cardId)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) ExtractCardId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "extractCardId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtractCardId is a free data retrieval call binding the contract method 0xe03d2b47.
//
// Solidity: function extractCardId(uint256 tokenId) pure returns(uint256 cardId)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) ExtractCardId(tokenId *big.Int) (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.ExtractCardId(&_LOCGamePlayNFT.CallOpts, tokenId)
}

// ExtractCardId is a free data retrieval call binding the contract method 0xe03d2b47.
//
// Solidity: function extractCardId(uint256 tokenId) pure returns(uint256 cardId)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) ExtractCardId(tokenId *big.Int) (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.ExtractCardId(&_LOCGamePlayNFT.CallOpts, tokenId)
}

// ExtractNumberOfCard is a free data retrieval call binding the contract method 0xf825b6ed.
//
// Solidity: function extractNumberOfCard(uint256 tokenId, uint256 cardId) pure returns(uint256 tokenCounter)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) ExtractNumberOfCard(opts *bind.CallOpts, tokenId *big.Int, cardId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "extractNumberOfCard", tokenId, cardId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtractNumberOfCard is a free data retrieval call binding the contract method 0xf825b6ed.
//
// Solidity: function extractNumberOfCard(uint256 tokenId, uint256 cardId) pure returns(uint256 tokenCounter)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) ExtractNumberOfCard(tokenId *big.Int, cardId *big.Int) (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.ExtractNumberOfCard(&_LOCGamePlayNFT.CallOpts, tokenId, cardId)
}

// ExtractNumberOfCard is a free data retrieval call binding the contract method 0xf825b6ed.
//
// Solidity: function extractNumberOfCard(uint256 tokenId, uint256 cardId) pure returns(uint256 tokenCounter)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) ExtractNumberOfCard(tokenId *big.Int, cardId *big.Int) (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.ExtractNumberOfCard(&_LOCGamePlayNFT.CallOpts, tokenId, cardId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LOCGamePlayNFT.Contract.GetApproved(&_LOCGamePlayNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LOCGamePlayNFT.Contract.GetApproved(&_LOCGamePlayNFT.CallOpts, tokenId)
}

// GetCardTotalSupply is a free data retrieval call binding the contract method 0x8e7fb9ce.
//
// Solidity: function getCardTotalSupply(uint256 cardId) view returns(uint256 cardTotalSupply)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) GetCardTotalSupply(opts *bind.CallOpts, cardId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "getCardTotalSupply", cardId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCardTotalSupply is a free data retrieval call binding the contract method 0x8e7fb9ce.
//
// Solidity: function getCardTotalSupply(uint256 cardId) view returns(uint256 cardTotalSupply)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) GetCardTotalSupply(cardId *big.Int) (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.GetCardTotalSupply(&_LOCGamePlayNFT.CallOpts, cardId)
}

// GetCardTotalSupply is a free data retrieval call binding the contract method 0x8e7fb9ce.
//
// Solidity: function getCardTotalSupply(uint256 cardId) view returns(uint256 cardTotalSupply)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) GetCardTotalSupply(cardId *big.Int) (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.GetCardTotalSupply(&_LOCGamePlayNFT.CallOpts, cardId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LOCGamePlayNFT.Contract.GetRoleAdmin(&_LOCGamePlayNFT.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LOCGamePlayNFT.Contract.GetRoleAdmin(&_LOCGamePlayNFT.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LOCGamePlayNFT.Contract.HasRole(&_LOCGamePlayNFT.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LOCGamePlayNFT.Contract.HasRole(&_LOCGamePlayNFT.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LOCGamePlayNFT.Contract.IsApprovedForAll(&_LOCGamePlayNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LOCGamePlayNFT.Contract.IsApprovedForAll(&_LOCGamePlayNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) Name() (string, error) {
	return _LOCGamePlayNFT.Contract.Name(&_LOCGamePlayNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) Name() (string, error) {
	return _LOCGamePlayNFT.Contract.Name(&_LOCGamePlayNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LOCGamePlayNFT.Contract.OwnerOf(&_LOCGamePlayNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LOCGamePlayNFT.Contract.OwnerOf(&_LOCGamePlayNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LOCGamePlayNFT.Contract.SupportsInterface(&_LOCGamePlayNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LOCGamePlayNFT.Contract.SupportsInterface(&_LOCGamePlayNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) Symbol() (string, error) {
	return _LOCGamePlayNFT.Contract.Symbol(&_LOCGamePlayNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) Symbol() (string, error) {
	return _LOCGamePlayNFT.Contract.Symbol(&_LOCGamePlayNFT.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.TokenByIndex(&_LOCGamePlayNFT.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.TokenByIndex(&_LOCGamePlayNFT.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.TokenOfOwnerByIndex(&_LOCGamePlayNFT.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.TokenOfOwnerByIndex(&_LOCGamePlayNFT.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LOCGamePlayNFT.Contract.TokenURI(&_LOCGamePlayNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LOCGamePlayNFT.Contract.TokenURI(&_LOCGamePlayNFT.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[] ownerTokens)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[] ownerTokens)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _LOCGamePlayNFT.Contract.TokensOfOwner(&_LOCGamePlayNFT.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[] ownerTokens)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _LOCGamePlayNFT.Contract.TokensOfOwner(&_LOCGamePlayNFT.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LOCGamePlayNFT *LOCGamePlayNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LOCGamePlayNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) TotalSupply() (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.TotalSupply(&_LOCGamePlayNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LOCGamePlayNFT *LOCGamePlayNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _LOCGamePlayNFT.Contract.TotalSupply(&_LOCGamePlayNFT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.Approve(&_LOCGamePlayNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.Approve(&_LOCGamePlayNFT.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.Burn(&_LOCGamePlayNFT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.Burn(&_LOCGamePlayNFT.TransactOpts, tokenId)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) Destroy() (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.Destroy(&_LOCGamePlayNFT.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) Destroy() (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.Destroy(&_LOCGamePlayNFT.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.GrantRole(&_LOCGamePlayNFT.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.GrantRole(&_LOCGamePlayNFT.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "mint", to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.Mint(&_LOCGamePlayNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.Mint(&_LOCGamePlayNFT.TransactOpts, to, tokenId)
}

// MintBatch is a paid mutator transaction binding the contract method 0x75ceb341.
//
// Solidity: function mintBatch(address to, uint256[] tokenIds) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) MintBatch(opts *bind.TransactOpts, to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "mintBatch", to, tokenIds)
}

// MintBatch is a paid mutator transaction binding the contract method 0x75ceb341.
//
// Solidity: function mintBatch(address to, uint256[] tokenIds) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) MintBatch(to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.MintBatch(&_LOCGamePlayNFT.TransactOpts, to, tokenIds)
}

// MintBatch is a paid mutator transaction binding the contract method 0x75ceb341.
//
// Solidity: function mintBatch(address to, uint256[] tokenIds) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) MintBatch(to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.MintBatch(&_LOCGamePlayNFT.TransactOpts, to, tokenIds)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.RenounceRole(&_LOCGamePlayNFT.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.RenounceRole(&_LOCGamePlayNFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.RevokeRole(&_LOCGamePlayNFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.RevokeRole(&_LOCGamePlayNFT.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.SafeTransferFrom(&_LOCGamePlayNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.SafeTransferFrom(&_LOCGamePlayNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.SafeTransferFrom0(&_LOCGamePlayNFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.SafeTransferFrom0(&_LOCGamePlayNFT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.SetApprovalForAll(&_LOCGamePlayNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.SetApprovalForAll(&_LOCGamePlayNFT.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) SetBaseURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "setBaseURI", uri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) SetBaseURI(uri string) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.SetBaseURI(&_LOCGamePlayNFT.TransactOpts, uri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) SetBaseURI(uri string) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.SetBaseURI(&_LOCGamePlayNFT.TransactOpts, uri)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.TransferFrom(&_LOCGamePlayNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LOCGamePlayNFT *LOCGamePlayNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LOCGamePlayNFT.Contract.TransferFrom(&_LOCGamePlayNFT.TransactOpts, from, to, tokenId)
}

// LOCGamePlayNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTApprovalIterator struct {
	Event *LOCGamePlayNFTApproval // Event containing the contract specifics and raw log

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
func (it *LOCGamePlayNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LOCGamePlayNFTApproval)
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
		it.Event = new(LOCGamePlayNFTApproval)
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
func (it *LOCGamePlayNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LOCGamePlayNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LOCGamePlayNFTApproval represents a Approval event raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LOCGamePlayNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LOCGamePlayNFTApprovalIterator{contract: _LOCGamePlayNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LOCGamePlayNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LOCGamePlayNFTApproval)
				if err := _LOCGamePlayNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) ParseApproval(log types.Log) (*LOCGamePlayNFTApproval, error) {
	event := new(LOCGamePlayNFTApproval)
	if err := _LOCGamePlayNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LOCGamePlayNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTApprovalForAllIterator struct {
	Event *LOCGamePlayNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LOCGamePlayNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LOCGamePlayNFTApprovalForAll)
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
		it.Event = new(LOCGamePlayNFTApprovalForAll)
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
func (it *LOCGamePlayNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LOCGamePlayNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LOCGamePlayNFTApprovalForAll represents a ApprovalForAll event raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LOCGamePlayNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LOCGamePlayNFTApprovalForAllIterator{contract: _LOCGamePlayNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LOCGamePlayNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LOCGamePlayNFTApprovalForAll)
				if err := _LOCGamePlayNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) ParseApprovalForAll(log types.Log) (*LOCGamePlayNFTApprovalForAll, error) {
	event := new(LOCGamePlayNFTApprovalForAll)
	if err := _LOCGamePlayNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LOCGamePlayNFTBaseURISetIterator is returned from FilterBaseURISet and is used to iterate over the raw logs and unpacked data for BaseURISet events raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTBaseURISetIterator struct {
	Event *LOCGamePlayNFTBaseURISet // Event containing the contract specifics and raw log

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
func (it *LOCGamePlayNFTBaseURISetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LOCGamePlayNFTBaseURISet)
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
		it.Event = new(LOCGamePlayNFTBaseURISet)
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
func (it *LOCGamePlayNFTBaseURISetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LOCGamePlayNFTBaseURISetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LOCGamePlayNFTBaseURISet represents a BaseURISet event raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTBaseURISet struct {
	BaseURI common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaseURISet is a free log retrieval operation binding the contract event 0xf9c7803e94e0d3c02900d8a90893a6d5e90dd04d32a4cfe825520f82bf9f32f6.
//
// Solidity: event BaseURISet(string indexed baseURI)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) FilterBaseURISet(opts *bind.FilterOpts, baseURI []string) (*LOCGamePlayNFTBaseURISetIterator, error) {

	var baseURIRule []interface{}
	for _, baseURIItem := range baseURI {
		baseURIRule = append(baseURIRule, baseURIItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.FilterLogs(opts, "BaseURISet", baseURIRule)
	if err != nil {
		return nil, err
	}
	return &LOCGamePlayNFTBaseURISetIterator{contract: _LOCGamePlayNFT.contract, event: "BaseURISet", logs: logs, sub: sub}, nil
}

// WatchBaseURISet is a free log subscription operation binding the contract event 0xf9c7803e94e0d3c02900d8a90893a6d5e90dd04d32a4cfe825520f82bf9f32f6.
//
// Solidity: event BaseURISet(string indexed baseURI)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) WatchBaseURISet(opts *bind.WatchOpts, sink chan<- *LOCGamePlayNFTBaseURISet, baseURI []string) (event.Subscription, error) {

	var baseURIRule []interface{}
	for _, baseURIItem := range baseURI {
		baseURIRule = append(baseURIRule, baseURIItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.WatchLogs(opts, "BaseURISet", baseURIRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LOCGamePlayNFTBaseURISet)
				if err := _LOCGamePlayNFT.contract.UnpackLog(event, "BaseURISet", log); err != nil {
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

// ParseBaseURISet is a log parse operation binding the contract event 0xf9c7803e94e0d3c02900d8a90893a6d5e90dd04d32a4cfe825520f82bf9f32f6.
//
// Solidity: event BaseURISet(string indexed baseURI)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) ParseBaseURISet(log types.Log) (*LOCGamePlayNFTBaseURISet, error) {
	event := new(LOCGamePlayNFTBaseURISet)
	if err := _LOCGamePlayNFT.contract.UnpackLog(event, "BaseURISet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LOCGamePlayNFTRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTRoleAdminChangedIterator struct {
	Event *LOCGamePlayNFTRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LOCGamePlayNFTRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LOCGamePlayNFTRoleAdminChanged)
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
		it.Event = new(LOCGamePlayNFTRoleAdminChanged)
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
func (it *LOCGamePlayNFTRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LOCGamePlayNFTRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LOCGamePlayNFTRoleAdminChanged represents a RoleAdminChanged event raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LOCGamePlayNFTRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LOCGamePlayNFTRoleAdminChangedIterator{contract: _LOCGamePlayNFT.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LOCGamePlayNFTRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LOCGamePlayNFTRoleAdminChanged)
				if err := _LOCGamePlayNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) ParseRoleAdminChanged(log types.Log) (*LOCGamePlayNFTRoleAdminChanged, error) {
	event := new(LOCGamePlayNFTRoleAdminChanged)
	if err := _LOCGamePlayNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LOCGamePlayNFTRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTRoleGrantedIterator struct {
	Event *LOCGamePlayNFTRoleGranted // Event containing the contract specifics and raw log

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
func (it *LOCGamePlayNFTRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LOCGamePlayNFTRoleGranted)
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
		it.Event = new(LOCGamePlayNFTRoleGranted)
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
func (it *LOCGamePlayNFTRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LOCGamePlayNFTRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LOCGamePlayNFTRoleGranted represents a RoleGranted event raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LOCGamePlayNFTRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LOCGamePlayNFTRoleGrantedIterator{contract: _LOCGamePlayNFT.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LOCGamePlayNFTRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LOCGamePlayNFTRoleGranted)
				if err := _LOCGamePlayNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) ParseRoleGranted(log types.Log) (*LOCGamePlayNFTRoleGranted, error) {
	event := new(LOCGamePlayNFTRoleGranted)
	if err := _LOCGamePlayNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LOCGamePlayNFTRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTRoleRevokedIterator struct {
	Event *LOCGamePlayNFTRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LOCGamePlayNFTRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LOCGamePlayNFTRoleRevoked)
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
		it.Event = new(LOCGamePlayNFTRoleRevoked)
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
func (it *LOCGamePlayNFTRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LOCGamePlayNFTRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LOCGamePlayNFTRoleRevoked represents a RoleRevoked event raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LOCGamePlayNFTRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LOCGamePlayNFTRoleRevokedIterator{contract: _LOCGamePlayNFT.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LOCGamePlayNFTRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LOCGamePlayNFTRoleRevoked)
				if err := _LOCGamePlayNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) ParseRoleRevoked(log types.Log) (*LOCGamePlayNFTRoleRevoked, error) {
	event := new(LOCGamePlayNFTRoleRevoked)
	if err := _LOCGamePlayNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LOCGamePlayNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTTransferIterator struct {
	Event *LOCGamePlayNFTTransfer // Event containing the contract specifics and raw log

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
func (it *LOCGamePlayNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LOCGamePlayNFTTransfer)
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
		it.Event = new(LOCGamePlayNFTTransfer)
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
func (it *LOCGamePlayNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LOCGamePlayNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LOCGamePlayNFTTransfer represents a Transfer event raised by the LOCGamePlayNFT contract.
type LOCGamePlayNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LOCGamePlayNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LOCGamePlayNFTTransferIterator{contract: _LOCGamePlayNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LOCGamePlayNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LOCGamePlayNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LOCGamePlayNFTTransfer)
				if err := _LOCGamePlayNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LOCGamePlayNFT *LOCGamePlayNFTFilterer) ParseTransfer(log types.Log) (*LOCGamePlayNFTTransfer, error) {
	event := new(LOCGamePlayNFTTransfer)
	if err := _LOCGamePlayNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
