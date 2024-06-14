// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier

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

// IAvailBridgeMerkleProofInput is an auto generated low-level Go binding around an user-defined struct.
type IAvailBridgeMerkleProofInput struct {
	DataRootProof [][32]byte
	LeafProof     [][32]byte
	RangeHash     [32]byte
	DataRootIndex *big.Int
	BlobRoot      [32]byte
	BridgeRoot    [32]byte
	Leaf          [32]byte
	LeafIndex     *big.Int
}

// VectorverifierMetaData contains all meta data concerning the Vectorverifier contract.
var VectorverifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"verifyDataAvailability\",\"inputs\":[{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structIAvailBridge.MerkleProofInput\",\"components\":[{\"name\":\"dataRootProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"leafProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"rangeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"dataRootIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blobRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leafIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyDataInclusion\",\"inputs\":[{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blobData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structIAvailBridge.MerkleProofInput\",\"components\":[{\"name\":\"dataRootProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"leafProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"rangeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"dataRootIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blobRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leafIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"}]",
}

// VectorverifierABI is the input ABI used to generate the binding from.
// Deprecated: Use VectorverifierMetaData.ABI instead.
var VectorverifierABI = VectorverifierMetaData.ABI

// Vectorverifier is an auto generated Go binding around an Ethereum contract.
type Vectorverifier struct {
	VectorverifierCaller     // Read-only binding to the contract
	VectorverifierTransactor // Write-only binding to the contract
	VectorverifierFilterer   // Log filterer for contract events
}

// VectorverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VectorverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VectorverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VectorverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VectorverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VectorverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VectorverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VectorverifierSession struct {
	Contract     *Vectorverifier   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VectorverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VectorverifierCallerSession struct {
	Contract *VectorverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VectorverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VectorverifierTransactorSession struct {
	Contract     *VectorverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VectorverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VectorverifierRaw struct {
	Contract *Vectorverifier // Generic contract binding to access the raw methods on
}

// VectorverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VectorverifierCallerRaw struct {
	Contract *VectorverifierCaller // Generic read-only contract binding to access the raw methods on
}

// VectorverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VectorverifierTransactorRaw struct {
	Contract *VectorverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVectorverifier creates a new instance of Vectorverifier, bound to a specific deployed contract.
func NewVectorverifier(address common.Address, backend bind.ContractBackend) (*Vectorverifier, error) {
	contract, err := bindVectorverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vectorverifier{VectorverifierCaller: VectorverifierCaller{contract: contract}, VectorverifierTransactor: VectorverifierTransactor{contract: contract}, VectorverifierFilterer: VectorverifierFilterer{contract: contract}}, nil
}

// NewVectorverifierCaller creates a new read-only instance of Vectorverifier, bound to a specific deployed contract.
func NewVectorverifierCaller(address common.Address, caller bind.ContractCaller) (*VectorverifierCaller, error) {
	contract, err := bindVectorverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VectorverifierCaller{contract: contract}, nil
}

// NewVectorverifierTransactor creates a new write-only instance of Vectorverifier, bound to a specific deployed contract.
func NewVectorverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VectorverifierTransactor, error) {
	contract, err := bindVectorverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VectorverifierTransactor{contract: contract}, nil
}

// NewVectorverifierFilterer creates a new log filterer instance of Vectorverifier, bound to a specific deployed contract.
func NewVectorverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VectorverifierFilterer, error) {
	contract, err := bindVectorverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VectorverifierFilterer{contract: contract}, nil
}

// bindVectorverifier binds a generic wrapper to an already deployed contract.
func bindVectorverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VectorverifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vectorverifier *VectorverifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vectorverifier.Contract.VectorverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vectorverifier *VectorverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vectorverifier.Contract.VectorverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vectorverifier *VectorverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vectorverifier.Contract.VectorverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vectorverifier *VectorverifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vectorverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vectorverifier *VectorverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vectorverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vectorverifier *VectorverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vectorverifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyDataAvailability is a free data retrieval call binding the contract method 0xf9ce76ba.
//
// Solidity: function verifyDataAvailability(address bridge, bytes32 blobHash, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) view returns(bool)
func (_Vectorverifier *VectorverifierCaller) VerifyDataAvailability(opts *bind.CallOpts, bridge common.Address, blobHash [32]byte, input IAvailBridgeMerkleProofInput) (bool, error) {
	var out []interface{}
	err := _Vectorverifier.contract.Call(opts, &out, "verifyDataAvailability", bridge, blobHash, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDataAvailability is a free data retrieval call binding the contract method 0xf9ce76ba.
//
// Solidity: function verifyDataAvailability(address bridge, bytes32 blobHash, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) view returns(bool)
func (_Vectorverifier *VectorverifierSession) VerifyDataAvailability(bridge common.Address, blobHash [32]byte, input IAvailBridgeMerkleProofInput) (bool, error) {
	return _Vectorverifier.Contract.VerifyDataAvailability(&_Vectorverifier.CallOpts, bridge, blobHash, input)
}

// VerifyDataAvailability is a free data retrieval call binding the contract method 0xf9ce76ba.
//
// Solidity: function verifyDataAvailability(address bridge, bytes32 blobHash, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) view returns(bool)
func (_Vectorverifier *VectorverifierCallerSession) VerifyDataAvailability(bridge common.Address, blobHash [32]byte, input IAvailBridgeMerkleProofInput) (bool, error) {
	return _Vectorverifier.Contract.VerifyDataAvailability(&_Vectorverifier.CallOpts, bridge, blobHash, input)
}

// VerifyDataInclusion is a free data retrieval call binding the contract method 0x541f47ea.
//
// Solidity: function verifyDataInclusion(address bridge, bytes blobData, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) view returns(bool)
func (_Vectorverifier *VectorverifierCaller) VerifyDataInclusion(opts *bind.CallOpts, bridge common.Address, blobData []byte, input IAvailBridgeMerkleProofInput) (bool, error) {
	var out []interface{}
	err := _Vectorverifier.contract.Call(opts, &out, "verifyDataInclusion", bridge, blobData, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDataInclusion is a free data retrieval call binding the contract method 0x541f47ea.
//
// Solidity: function verifyDataInclusion(address bridge, bytes blobData, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) view returns(bool)
func (_Vectorverifier *VectorverifierSession) VerifyDataInclusion(bridge common.Address, blobData []byte, input IAvailBridgeMerkleProofInput) (bool, error) {
	return _Vectorverifier.Contract.VerifyDataInclusion(&_Vectorverifier.CallOpts, bridge, blobData, input)
}

// VerifyDataInclusion is a free data retrieval call binding the contract method 0x541f47ea.
//
// Solidity: function verifyDataInclusion(address bridge, bytes blobData, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) view returns(bool)
func (_Vectorverifier *VectorverifierCallerSession) VerifyDataInclusion(bridge common.Address, blobData []byte, input IAvailBridgeMerkleProofInput) (bool, error) {
	return _Vectorverifier.Contract.VerifyDataInclusion(&_Vectorverifier.CallOpts, bridge, blobData, input)
}
