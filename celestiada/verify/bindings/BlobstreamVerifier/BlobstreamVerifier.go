// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blobstreamverifier

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

// AttestationProof is an auto generated low-level Go binding around an user-defined struct.
type AttestationProof struct {
	TupleRootNonce *big.Int
	Tuple          DataRootTuple
	Proof          BinaryMerkleProof
}

// BinaryMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type BinaryMerkleProof struct {
	SideNodes [][32]byte
	Key       *big.Int
	NumLeaves *big.Int
}

// DataRootTuple is an auto generated low-level Go binding around an user-defined struct.
type DataRootTuple struct {
	Height   *big.Int
	DataRoot [32]byte
}

// Namespace is an auto generated low-level Go binding around an user-defined struct.
type Namespace struct {
	Version [1]byte
	Id      [28]byte
}

// NamespaceMerkleMultiproof is an auto generated low-level Go binding around an user-defined struct.
type NamespaceMerkleMultiproof struct {
	BeginKey  *big.Int
	EndKey    *big.Int
	SideNodes []NamespaceNode
}

// NamespaceNode is an auto generated low-level Go binding around an user-defined struct.
type NamespaceNode struct {
	Min    Namespace
	Max    Namespace
	Digest [32]byte
}

// SharesProof is an auto generated low-level Go binding around an user-defined struct.
type SharesProof struct {
	Data             [][]byte
	ShareProofs      []NamespaceMerkleMultiproof
	Namespace        Namespace
	RowRoots         []NamespaceNode
	RowProofs        []BinaryMerkleProof
	AttestationProof AttestationProof
}

// SpanSequence is an auto generated low-level Go binding around an user-defined struct.
type SpanSequence struct {
	Height *big.Int
	Index  *big.Int
	Length *big.Int
}

// BlobstreamverifierMetaData contains all meta data concerning the Blobstreamverifier contract.
var BlobstreamverifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"verifyDataAvailability\",\"inputs\":[{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"span\",\"type\":\"tuple\",\"internalType\":\"structSpanSequence\",\"components\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_rowRoots\",\"type\":\"tuple[]\",\"internalType\":\"structNamespaceNode[]\",\"components\":[{\"name\":\"min\",\"type\":\"tuple\",\"internalType\":\"structNamespace\",\"components\":[{\"name\":\"version\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"id\",\"type\":\"bytes28\",\"internalType\":\"bytes28\"}]},{\"name\":\"max\",\"type\":\"tuple\",\"internalType\":\"structNamespace\",\"components\":[{\"name\":\"version\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"id\",\"type\":\"bytes28\",\"internalType\":\"bytes28\"}]},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_rowProofs\",\"type\":\"tuple[]\",\"internalType\":\"structBinaryMerkleProof[]\",\"components\":[{\"name\":\"sideNodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"key\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numLeaves\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_attestationProof\",\"type\":\"tuple\",\"internalType\":\"structAttestationProof\",\"components\":[{\"name\":\"tupleRootNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tuple\",\"type\":\"tuple\",\"internalType\":\"structDataRootTuple\",\"components\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structBinaryMerkleProof\",\"components\":[{\"name\":\"sideNodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"key\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numLeaves\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyDataInclusion\",\"inputs\":[{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sharesProof\",\"type\":\"tuple\",\"internalType\":\"structSharesProof\",\"components\":[{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"shareProofs\",\"type\":\"tuple[]\",\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"components\":[{\"name\":\"beginKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sideNodes\",\"type\":\"tuple[]\",\"internalType\":\"structNamespaceNode[]\",\"components\":[{\"name\":\"min\",\"type\":\"tuple\",\"internalType\":\"structNamespace\",\"components\":[{\"name\":\"version\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"id\",\"type\":\"bytes28\",\"internalType\":\"bytes28\"}]},{\"name\":\"max\",\"type\":\"tuple\",\"internalType\":\"structNamespace\",\"components\":[{\"name\":\"version\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"id\",\"type\":\"bytes28\",\"internalType\":\"bytes28\"}]},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"namespace\",\"type\":\"tuple\",\"internalType\":\"structNamespace\",\"components\":[{\"name\":\"version\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"id\",\"type\":\"bytes28\",\"internalType\":\"bytes28\"}]},{\"name\":\"rowRoots\",\"type\":\"tuple[]\",\"internalType\":\"structNamespaceNode[]\",\"components\":[{\"name\":\"min\",\"type\":\"tuple\",\"internalType\":\"structNamespace\",\"components\":[{\"name\":\"version\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"id\",\"type\":\"bytes28\",\"internalType\":\"bytes28\"}]},{\"name\":\"max\",\"type\":\"tuple\",\"internalType\":\"structNamespace\",\"components\":[{\"name\":\"version\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"id\",\"type\":\"bytes28\",\"internalType\":\"bytes28\"}]},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"rowProofs\",\"type\":\"tuple[]\",\"internalType\":\"structBinaryMerkleProof[]\",\"components\":[{\"name\":\"sideNodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"key\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numLeaves\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"attestationProof\",\"type\":\"tuple\",\"internalType\":\"structAttestationProof\",\"components\":[{\"name\":\"tupleRootNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tuple\",\"type\":\"tuple\",\"internalType\":\"structDataRootTuple\",\"components\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structBinaryMerkleProof\",\"components\":[{\"name\":\"sideNodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"key\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numLeaves\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"span\",\"type\":\"tuple\",\"internalType\":\"structSpanSequence\",\"components\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"}]",
}

// BlobstreamverifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BlobstreamverifierMetaData.ABI instead.
var BlobstreamverifierABI = BlobstreamverifierMetaData.ABI

// Blobstreamverifier is an auto generated Go binding around an Ethereum contract.
type Blobstreamverifier struct {
	BlobstreamverifierCaller     // Read-only binding to the contract
	BlobstreamverifierTransactor // Write-only binding to the contract
	BlobstreamverifierFilterer   // Log filterer for contract events
}

// BlobstreamverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlobstreamverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlobstreamverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlobstreamverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlobstreamverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlobstreamverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlobstreamverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlobstreamverifierSession struct {
	Contract     *Blobstreamverifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BlobstreamverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlobstreamverifierCallerSession struct {
	Contract *BlobstreamverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BlobstreamverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlobstreamverifierTransactorSession struct {
	Contract     *BlobstreamverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BlobstreamverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlobstreamverifierRaw struct {
	Contract *Blobstreamverifier // Generic contract binding to access the raw methods on
}

// BlobstreamverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlobstreamverifierCallerRaw struct {
	Contract *BlobstreamverifierCaller // Generic read-only contract binding to access the raw methods on
}

// BlobstreamverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlobstreamverifierTransactorRaw struct {
	Contract *BlobstreamverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlobstreamverifier creates a new instance of Blobstreamverifier, bound to a specific deployed contract.
func NewBlobstreamverifier(address common.Address, backend bind.ContractBackend) (*Blobstreamverifier, error) {
	contract, err := bindBlobstreamverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blobstreamverifier{BlobstreamverifierCaller: BlobstreamverifierCaller{contract: contract}, BlobstreamverifierTransactor: BlobstreamverifierTransactor{contract: contract}, BlobstreamverifierFilterer: BlobstreamverifierFilterer{contract: contract}}, nil
}

// NewBlobstreamverifierCaller creates a new read-only instance of Blobstreamverifier, bound to a specific deployed contract.
func NewBlobstreamverifierCaller(address common.Address, caller bind.ContractCaller) (*BlobstreamverifierCaller, error) {
	contract, err := bindBlobstreamverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlobstreamverifierCaller{contract: contract}, nil
}

// NewBlobstreamverifierTransactor creates a new write-only instance of Blobstreamverifier, bound to a specific deployed contract.
func NewBlobstreamverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BlobstreamverifierTransactor, error) {
	contract, err := bindBlobstreamverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlobstreamverifierTransactor{contract: contract}, nil
}

// NewBlobstreamverifierFilterer creates a new log filterer instance of Blobstreamverifier, bound to a specific deployed contract.
func NewBlobstreamverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BlobstreamverifierFilterer, error) {
	contract, err := bindBlobstreamverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlobstreamverifierFilterer{contract: contract}, nil
}

// bindBlobstreamverifier binds a generic wrapper to an already deployed contract.
func bindBlobstreamverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlobstreamverifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blobstreamverifier *BlobstreamverifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blobstreamverifier.Contract.BlobstreamverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blobstreamverifier *BlobstreamverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blobstreamverifier.Contract.BlobstreamverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blobstreamverifier *BlobstreamverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blobstreamverifier.Contract.BlobstreamverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blobstreamverifier *BlobstreamverifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blobstreamverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blobstreamverifier *BlobstreamverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blobstreamverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blobstreamverifier *BlobstreamverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blobstreamverifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyDataAvailability is a free data retrieval call binding the contract method 0x10284901.
//
// Solidity: function verifyDataAvailability(address bridge, (uint256,uint256,uint256) span, ((bytes1,bytes28),(bytes1,bytes28),bytes32)[] _rowRoots, (bytes32[],uint256,uint256)[] _rowProofs, (uint256,(uint256,bytes32),(bytes32[],uint256,uint256)) _attestationProof, bytes32 root) view returns(bool)
func (_Blobstreamverifier *BlobstreamverifierCaller) VerifyDataAvailability(opts *bind.CallOpts, bridge common.Address, span SpanSequence, _rowRoots []NamespaceNode, _rowProofs []BinaryMerkleProof, _attestationProof AttestationProof, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Blobstreamverifier.contract.Call(opts, &out, "verifyDataAvailability", bridge, span, _rowRoots, _rowProofs, _attestationProof, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDataAvailability is a free data retrieval call binding the contract method 0x10284901.
//
// Solidity: function verifyDataAvailability(address bridge, (uint256,uint256,uint256) span, ((bytes1,bytes28),(bytes1,bytes28),bytes32)[] _rowRoots, (bytes32[],uint256,uint256)[] _rowProofs, (uint256,(uint256,bytes32),(bytes32[],uint256,uint256)) _attestationProof, bytes32 root) view returns(bool)
func (_Blobstreamverifier *BlobstreamverifierSession) VerifyDataAvailability(bridge common.Address, span SpanSequence, _rowRoots []NamespaceNode, _rowProofs []BinaryMerkleProof, _attestationProof AttestationProof, root [32]byte) (bool, error) {
	return _Blobstreamverifier.Contract.VerifyDataAvailability(&_Blobstreamverifier.CallOpts, bridge, span, _rowRoots, _rowProofs, _attestationProof, root)
}

// VerifyDataAvailability is a free data retrieval call binding the contract method 0x10284901.
//
// Solidity: function verifyDataAvailability(address bridge, (uint256,uint256,uint256) span, ((bytes1,bytes28),(bytes1,bytes28),bytes32)[] _rowRoots, (bytes32[],uint256,uint256)[] _rowProofs, (uint256,(uint256,bytes32),(bytes32[],uint256,uint256)) _attestationProof, bytes32 root) view returns(bool)
func (_Blobstreamverifier *BlobstreamverifierCallerSession) VerifyDataAvailability(bridge common.Address, span SpanSequence, _rowRoots []NamespaceNode, _rowProofs []BinaryMerkleProof, _attestationProof AttestationProof, root [32]byte) (bool, error) {
	return _Blobstreamverifier.Contract.VerifyDataAvailability(&_Blobstreamverifier.CallOpts, bridge, span, _rowRoots, _rowProofs, _attestationProof, root)
}

// VerifyDataInclusion is a free data retrieval call binding the contract method 0x181ffdc0.
//
// Solidity: function verifyDataInclusion(address bridge, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes32)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) sharesProof, bytes32 root, (uint256,uint256,uint256) span) view returns(bool)
func (_Blobstreamverifier *BlobstreamverifierCaller) VerifyDataInclusion(opts *bind.CallOpts, bridge common.Address, sharesProof SharesProof, root [32]byte, span SpanSequence) (bool, error) {
	var out []interface{}
	err := _Blobstreamverifier.contract.Call(opts, &out, "verifyDataInclusion", bridge, sharesProof, root, span)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDataInclusion is a free data retrieval call binding the contract method 0x181ffdc0.
//
// Solidity: function verifyDataInclusion(address bridge, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes32)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) sharesProof, bytes32 root, (uint256,uint256,uint256) span) view returns(bool)
func (_Blobstreamverifier *BlobstreamverifierSession) VerifyDataInclusion(bridge common.Address, sharesProof SharesProof, root [32]byte, span SpanSequence) (bool, error) {
	return _Blobstreamverifier.Contract.VerifyDataInclusion(&_Blobstreamverifier.CallOpts, bridge, sharesProof, root, span)
}

// VerifyDataInclusion is a free data retrieval call binding the contract method 0x181ffdc0.
//
// Solidity: function verifyDataInclusion(address bridge, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes32)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) sharesProof, bytes32 root, (uint256,uint256,uint256) span) view returns(bool)
func (_Blobstreamverifier *BlobstreamverifierCallerSession) VerifyDataInclusion(bridge common.Address, sharesProof SharesProof, root [32]byte, span SpanSequence) (bool, error) {
	return _Blobstreamverifier.Contract.VerifyDataInclusion(&_Blobstreamverifier.CallOpts, bridge, sharesProof, root, span)
}
