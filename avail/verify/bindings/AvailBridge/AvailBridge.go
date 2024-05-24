// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package availbridge

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

// IAvailBridgeMessage is an auto generated low-level Go binding around an user-defined struct.
type IAvailBridgeMessage struct {
	MessageType       [1]byte
	From              [32]byte
	To                [32]byte
	OriginDomain      uint32
	DestinationDomain uint32
	Data              []byte
	MessageId         uint64
}

// AvailbridgeMetaData contains all meta data concerning the Availbridge contract.
var AvailbridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyBridged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlobRootEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeRootEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataRootCommitmentEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAssetId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataRootProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationOrAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFungibleTokenTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLeaf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnlockFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageId\",\"type\":\"uint256\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageId\",\"type\":\"uint256\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avail\",\"outputs\":[{\"internalType\":\"contractIAvail\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePerByte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeePerByte\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"contractIAvail\",\"name\":\"newAvail\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"},{\"internalType\":\"contractIVectorx\",\"name\":\"newVectorx\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isBridged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isSent\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"messageType\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"messageId\",\"type\":\"uint64\"}],\"internalType\":\"structIAvailBridge.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"dataRootProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"leafProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rangeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dataRootIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blobRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIAvailBridge.MerkleProofInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"receiveAVAIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"messageType\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"messageId\",\"type\":\"uint64\"}],\"internalType\":\"structIAvailBridge.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"dataRootProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"leafProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rangeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dataRootIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blobRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIAvailBridge.MerkleProofInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"receiveERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"messageType\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"messageId\",\"type\":\"uint64\"}],\"internalType\":\"structIAvailBridge.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"dataRootProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"leafProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rangeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dataRootIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blobRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIAvailBridge.MerkleProofInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"receiveETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"messageType\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"messageId\",\"type\":\"uint64\"}],\"internalType\":\"structIAvailBridge.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"dataRootProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"leafProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rangeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dataRootIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blobRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIAvailBridge.MerkleProofInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"receiveMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendAVAIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"}],\"name\":\"sendETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeePerByte\",\"type\":\"uint256\"}],\"name\":\"updateFeePerByte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"assetIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenAddresses\",\"type\":\"address[]\"}],\"name\":\"updateTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVectorx\",\"name\":\"newVectorx\",\"type\":\"address\"}],\"name\":\"updateVectorx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vectorx\",\"outputs\":[{\"internalType\":\"contractIVectorx\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"dataRootProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"leafProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rangeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dataRootIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blobRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIAvailBridge.MerkleProofInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"verifyBlobLeaf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"dataRootProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"leafProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rangeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dataRootIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blobRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIAvailBridge.MerkleProofInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"verifyBridgeLeaf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AvailbridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use AvailbridgeMetaData.ABI instead.
var AvailbridgeABI = AvailbridgeMetaData.ABI

// Availbridge is an auto generated Go binding around an Ethereum contract.
type Availbridge struct {
	AvailbridgeCaller     // Read-only binding to the contract
	AvailbridgeTransactor // Write-only binding to the contract
	AvailbridgeFilterer   // Log filterer for contract events
}

// AvailbridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvailbridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvailbridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvailbridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvailbridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvailbridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvailbridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvailbridgeSession struct {
	Contract     *Availbridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AvailbridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvailbridgeCallerSession struct {
	Contract *AvailbridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AvailbridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvailbridgeTransactorSession struct {
	Contract     *AvailbridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AvailbridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvailbridgeRaw struct {
	Contract *Availbridge // Generic contract binding to access the raw methods on
}

// AvailbridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvailbridgeCallerRaw struct {
	Contract *AvailbridgeCaller // Generic read-only contract binding to access the raw methods on
}

// AvailbridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvailbridgeTransactorRaw struct {
	Contract *AvailbridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvailbridge creates a new instance of Availbridge, bound to a specific deployed contract.
func NewAvailbridge(address common.Address, backend bind.ContractBackend) (*Availbridge, error) {
	contract, err := bindAvailbridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Availbridge{AvailbridgeCaller: AvailbridgeCaller{contract: contract}, AvailbridgeTransactor: AvailbridgeTransactor{contract: contract}, AvailbridgeFilterer: AvailbridgeFilterer{contract: contract}}, nil
}

// NewAvailbridgeCaller creates a new read-only instance of Availbridge, bound to a specific deployed contract.
func NewAvailbridgeCaller(address common.Address, caller bind.ContractCaller) (*AvailbridgeCaller, error) {
	contract, err := bindAvailbridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvailbridgeCaller{contract: contract}, nil
}

// NewAvailbridgeTransactor creates a new write-only instance of Availbridge, bound to a specific deployed contract.
func NewAvailbridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*AvailbridgeTransactor, error) {
	contract, err := bindAvailbridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvailbridgeTransactor{contract: contract}, nil
}

// NewAvailbridgeFilterer creates a new log filterer instance of Availbridge, bound to a specific deployed contract.
func NewAvailbridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*AvailbridgeFilterer, error) {
	contract, err := bindAvailbridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvailbridgeFilterer{contract: contract}, nil
}

// bindAvailbridge binds a generic wrapper to an already deployed contract.
func bindAvailbridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AvailbridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Availbridge *AvailbridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Availbridge.Contract.AvailbridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Availbridge *AvailbridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Availbridge.Contract.AvailbridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Availbridge *AvailbridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Availbridge.Contract.AvailbridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Availbridge *AvailbridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Availbridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Availbridge *AvailbridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Availbridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Availbridge *AvailbridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Availbridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Availbridge *AvailbridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Availbridge *AvailbridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Availbridge.Contract.DEFAULTADMINROLE(&_Availbridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Availbridge *AvailbridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Availbridge.Contract.DEFAULTADMINROLE(&_Availbridge.CallOpts)
}

// Avail is a free data retrieval call binding the contract method 0x5bdeac42.
//
// Solidity: function avail() view returns(address)
func (_Availbridge *AvailbridgeCaller) Avail(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "avail")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avail is a free data retrieval call binding the contract method 0x5bdeac42.
//
// Solidity: function avail() view returns(address)
func (_Availbridge *AvailbridgeSession) Avail() (common.Address, error) {
	return _Availbridge.Contract.Avail(&_Availbridge.CallOpts)
}

// Avail is a free data retrieval call binding the contract method 0x5bdeac42.
//
// Solidity: function avail() view returns(address)
func (_Availbridge *AvailbridgeCallerSession) Avail() (common.Address, error) {
	return _Availbridge.Contract.Avail(&_Availbridge.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_Availbridge *AvailbridgeCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_Availbridge *AvailbridgeSession) DefaultAdmin() (common.Address, error) {
	return _Availbridge.Contract.DefaultAdmin(&_Availbridge.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_Availbridge *AvailbridgeCallerSession) DefaultAdmin() (common.Address, error) {
	return _Availbridge.Contract.DefaultAdmin(&_Availbridge.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_Availbridge *AvailbridgeCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_Availbridge *AvailbridgeSession) DefaultAdminDelay() (*big.Int, error) {
	return _Availbridge.Contract.DefaultAdminDelay(&_Availbridge.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_Availbridge *AvailbridgeCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _Availbridge.Contract.DefaultAdminDelay(&_Availbridge.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_Availbridge *AvailbridgeCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_Availbridge *AvailbridgeSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _Availbridge.Contract.DefaultAdminDelayIncreaseWait(&_Availbridge.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_Availbridge *AvailbridgeCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _Availbridge.Contract.DefaultAdminDelayIncreaseWait(&_Availbridge.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_Availbridge *AvailbridgeCaller) FeePerByte(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "feePerByte")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_Availbridge *AvailbridgeSession) FeePerByte() (*big.Int, error) {
	return _Availbridge.Contract.FeePerByte(&_Availbridge.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_Availbridge *AvailbridgeCallerSession) FeePerByte() (*big.Int, error) {
	return _Availbridge.Contract.FeePerByte(&_Availbridge.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Availbridge *AvailbridgeCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Availbridge *AvailbridgeSession) FeeRecipient() (common.Address, error) {
	return _Availbridge.Contract.FeeRecipient(&_Availbridge.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Availbridge *AvailbridgeCallerSession) FeeRecipient() (common.Address, error) {
	return _Availbridge.Contract.FeeRecipient(&_Availbridge.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_Availbridge *AvailbridgeCaller) Fees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "fees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_Availbridge *AvailbridgeSession) Fees() (*big.Int, error) {
	return _Availbridge.Contract.Fees(&_Availbridge.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_Availbridge *AvailbridgeCallerSession) Fees() (*big.Int, error) {
	return _Availbridge.Contract.Fees(&_Availbridge.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xfcee45f4.
//
// Solidity: function getFee(uint256 length) view returns(uint256)
func (_Availbridge *AvailbridgeCaller) GetFee(opts *bind.CallOpts, length *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "getFee", length)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xfcee45f4.
//
// Solidity: function getFee(uint256 length) view returns(uint256)
func (_Availbridge *AvailbridgeSession) GetFee(length *big.Int) (*big.Int, error) {
	return _Availbridge.Contract.GetFee(&_Availbridge.CallOpts, length)
}

// GetFee is a free data retrieval call binding the contract method 0xfcee45f4.
//
// Solidity: function getFee(uint256 length) view returns(uint256)
func (_Availbridge *AvailbridgeCallerSession) GetFee(length *big.Int) (*big.Int, error) {
	return _Availbridge.Contract.GetFee(&_Availbridge.CallOpts, length)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Availbridge *AvailbridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Availbridge *AvailbridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Availbridge.Contract.GetRoleAdmin(&_Availbridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Availbridge *AvailbridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Availbridge.Contract.GetRoleAdmin(&_Availbridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Availbridge *AvailbridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Availbridge *AvailbridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Availbridge.Contract.HasRole(&_Availbridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Availbridge *AvailbridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Availbridge.Contract.HasRole(&_Availbridge.CallOpts, role, account)
}

// IsBridged is a free data retrieval call binding the contract method 0x701bbfe8.
//
// Solidity: function isBridged(bytes32 ) view returns(bool)
func (_Availbridge *AvailbridgeCaller) IsBridged(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "isBridged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBridged is a free data retrieval call binding the contract method 0x701bbfe8.
//
// Solidity: function isBridged(bytes32 ) view returns(bool)
func (_Availbridge *AvailbridgeSession) IsBridged(arg0 [32]byte) (bool, error) {
	return _Availbridge.Contract.IsBridged(&_Availbridge.CallOpts, arg0)
}

// IsBridged is a free data retrieval call binding the contract method 0x701bbfe8.
//
// Solidity: function isBridged(bytes32 ) view returns(bool)
func (_Availbridge *AvailbridgeCallerSession) IsBridged(arg0 [32]byte) (bool, error) {
	return _Availbridge.Contract.IsBridged(&_Availbridge.CallOpts, arg0)
}

// IsSent is a free data retrieval call binding the contract method 0x63d91f40.
//
// Solidity: function isSent(uint256 ) view returns(bytes32)
func (_Availbridge *AvailbridgeCaller) IsSent(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "isSent", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IsSent is a free data retrieval call binding the contract method 0x63d91f40.
//
// Solidity: function isSent(uint256 ) view returns(bytes32)
func (_Availbridge *AvailbridgeSession) IsSent(arg0 *big.Int) ([32]byte, error) {
	return _Availbridge.Contract.IsSent(&_Availbridge.CallOpts, arg0)
}

// IsSent is a free data retrieval call binding the contract method 0x63d91f40.
//
// Solidity: function isSent(uint256 ) view returns(bytes32)
func (_Availbridge *AvailbridgeCallerSession) IsSent(arg0 *big.Int) ([32]byte, error) {
	return _Availbridge.Contract.IsSent(&_Availbridge.CallOpts, arg0)
}

// MessageId is a free data retrieval call binding the contract method 0x669f618b.
//
// Solidity: function messageId() view returns(uint256)
func (_Availbridge *AvailbridgeCaller) MessageId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "messageId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageId is a free data retrieval call binding the contract method 0x669f618b.
//
// Solidity: function messageId() view returns(uint256)
func (_Availbridge *AvailbridgeSession) MessageId() (*big.Int, error) {
	return _Availbridge.Contract.MessageId(&_Availbridge.CallOpts)
}

// MessageId is a free data retrieval call binding the contract method 0x669f618b.
//
// Solidity: function messageId() view returns(uint256)
func (_Availbridge *AvailbridgeCallerSession) MessageId() (*big.Int, error) {
	return _Availbridge.Contract.MessageId(&_Availbridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Availbridge *AvailbridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Availbridge *AvailbridgeSession) Owner() (common.Address, error) {
	return _Availbridge.Contract.Owner(&_Availbridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Availbridge *AvailbridgeCallerSession) Owner() (common.Address, error) {
	return _Availbridge.Contract.Owner(&_Availbridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Availbridge *AvailbridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Availbridge *AvailbridgeSession) Paused() (bool, error) {
	return _Availbridge.Contract.Paused(&_Availbridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Availbridge *AvailbridgeCallerSession) Paused() (bool, error) {
	return _Availbridge.Contract.Paused(&_Availbridge.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_Availbridge *AvailbridgeCaller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(struct {
		NewAdmin common.Address
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_Availbridge *AvailbridgeSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _Availbridge.Contract.PendingDefaultAdmin(&_Availbridge.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_Availbridge *AvailbridgeCallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _Availbridge.Contract.PendingDefaultAdmin(&_Availbridge.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_Availbridge *AvailbridgeCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(struct {
		NewDelay *big.Int
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_Availbridge *AvailbridgeSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _Availbridge.Contract.PendingDefaultAdminDelay(&_Availbridge.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_Availbridge *AvailbridgeCallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _Availbridge.Contract.PendingDefaultAdminDelay(&_Availbridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Availbridge *AvailbridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Availbridge *AvailbridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Availbridge.Contract.SupportsInterface(&_Availbridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Availbridge *AvailbridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Availbridge.Contract.SupportsInterface(&_Availbridge.CallOpts, interfaceId)
}

// Tokens is a free data retrieval call binding the contract method 0x904194a3.
//
// Solidity: function tokens(bytes32 ) view returns(address)
func (_Availbridge *AvailbridgeCaller) Tokens(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x904194a3.
//
// Solidity: function tokens(bytes32 ) view returns(address)
func (_Availbridge *AvailbridgeSession) Tokens(arg0 [32]byte) (common.Address, error) {
	return _Availbridge.Contract.Tokens(&_Availbridge.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x904194a3.
//
// Solidity: function tokens(bytes32 ) view returns(address)
func (_Availbridge *AvailbridgeCallerSession) Tokens(arg0 [32]byte) (common.Address, error) {
	return _Availbridge.Contract.Tokens(&_Availbridge.CallOpts, arg0)
}

// Vectorx is a free data retrieval call binding the contract method 0xcc778e84.
//
// Solidity: function vectorx() view returns(address)
func (_Availbridge *AvailbridgeCaller) Vectorx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "vectorx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vectorx is a free data retrieval call binding the contract method 0xcc778e84.
//
// Solidity: function vectorx() view returns(address)
func (_Availbridge *AvailbridgeSession) Vectorx() (common.Address, error) {
	return _Availbridge.Contract.Vectorx(&_Availbridge.CallOpts)
}

// Vectorx is a free data retrieval call binding the contract method 0xcc778e84.
//
// Solidity: function vectorx() view returns(address)
func (_Availbridge *AvailbridgeCallerSession) Vectorx() (common.Address, error) {
	return _Availbridge.Contract.Vectorx(&_Availbridge.CallOpts)
}

// VerifyBlobLeaf is a free data retrieval call binding the contract method 0x06fc6005.
//
// Solidity: function verifyBlobLeaf((bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) view returns(bool)
func (_Availbridge *AvailbridgeCaller) VerifyBlobLeaf(opts *bind.CallOpts, input IAvailBridgeMerkleProofInput) (bool, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "verifyBlobLeaf", input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBlobLeaf is a free data retrieval call binding the contract method 0x06fc6005.
//
// Solidity: function verifyBlobLeaf((bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) view returns(bool)
func (_Availbridge *AvailbridgeSession) VerifyBlobLeaf(input IAvailBridgeMerkleProofInput) (bool, error) {
	return _Availbridge.Contract.VerifyBlobLeaf(&_Availbridge.CallOpts, input)
}

// VerifyBlobLeaf is a free data retrieval call binding the contract method 0x06fc6005.
//
// Solidity: function verifyBlobLeaf((bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) view returns(bool)
func (_Availbridge *AvailbridgeCallerSession) VerifyBlobLeaf(input IAvailBridgeMerkleProofInput) (bool, error) {
	return _Availbridge.Contract.VerifyBlobLeaf(&_Availbridge.CallOpts, input)
}

// VerifyBridgeLeaf is a free data retrieval call binding the contract method 0xaf53dade.
//
// Solidity: function verifyBridgeLeaf((bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) view returns(bool)
func (_Availbridge *AvailbridgeCaller) VerifyBridgeLeaf(opts *bind.CallOpts, input IAvailBridgeMerkleProofInput) (bool, error) {
	var out []interface{}
	err := _Availbridge.contract.Call(opts, &out, "verifyBridgeLeaf", input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBridgeLeaf is a free data retrieval call binding the contract method 0xaf53dade.
//
// Solidity: function verifyBridgeLeaf((bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) view returns(bool)
func (_Availbridge *AvailbridgeSession) VerifyBridgeLeaf(input IAvailBridgeMerkleProofInput) (bool, error) {
	return _Availbridge.Contract.VerifyBridgeLeaf(&_Availbridge.CallOpts, input)
}

// VerifyBridgeLeaf is a free data retrieval call binding the contract method 0xaf53dade.
//
// Solidity: function verifyBridgeLeaf((bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) view returns(bool)
func (_Availbridge *AvailbridgeCallerSession) VerifyBridgeLeaf(input IAvailBridgeMerkleProofInput) (bool, error) {
	return _Availbridge.Contract.VerifyBridgeLeaf(&_Availbridge.CallOpts, input)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_Availbridge *AvailbridgeTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_Availbridge *AvailbridgeSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _Availbridge.Contract.AcceptDefaultAdminTransfer(&_Availbridge.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_Availbridge *AvailbridgeTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _Availbridge.Contract.AcceptDefaultAdminTransfer(&_Availbridge.TransactOpts)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_Availbridge *AvailbridgeTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_Availbridge *AvailbridgeSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.BeginDefaultAdminTransfer(&_Availbridge.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_Availbridge *AvailbridgeTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.BeginDefaultAdminTransfer(&_Availbridge.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_Availbridge *AvailbridgeTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_Availbridge *AvailbridgeSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _Availbridge.Contract.CancelDefaultAdminTransfer(&_Availbridge.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_Availbridge *AvailbridgeTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _Availbridge.Contract.CancelDefaultAdminTransfer(&_Availbridge.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_Availbridge *AvailbridgeTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_Availbridge *AvailbridgeSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _Availbridge.Contract.ChangeDefaultAdminDelay(&_Availbridge.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_Availbridge *AvailbridgeTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _Availbridge.Contract.ChangeDefaultAdminDelay(&_Availbridge.TransactOpts, newDelay)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Availbridge *AvailbridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Availbridge *AvailbridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.GrantRole(&_Availbridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Availbridge *AvailbridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.GrantRole(&_Availbridge.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x4df23081.
//
// Solidity: function initialize(uint256 newFeePerByte, address newFeeRecipient, address newAvail, address governance, address pauser, address newVectorx) returns()
func (_Availbridge *AvailbridgeTransactor) Initialize(opts *bind.TransactOpts, newFeePerByte *big.Int, newFeeRecipient common.Address, newAvail common.Address, governance common.Address, pauser common.Address, newVectorx common.Address) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "initialize", newFeePerByte, newFeeRecipient, newAvail, governance, pauser, newVectorx)
}

// Initialize is a paid mutator transaction binding the contract method 0x4df23081.
//
// Solidity: function initialize(uint256 newFeePerByte, address newFeeRecipient, address newAvail, address governance, address pauser, address newVectorx) returns()
func (_Availbridge *AvailbridgeSession) Initialize(newFeePerByte *big.Int, newFeeRecipient common.Address, newAvail common.Address, governance common.Address, pauser common.Address, newVectorx common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.Initialize(&_Availbridge.TransactOpts, newFeePerByte, newFeeRecipient, newAvail, governance, pauser, newVectorx)
}

// Initialize is a paid mutator transaction binding the contract method 0x4df23081.
//
// Solidity: function initialize(uint256 newFeePerByte, address newFeeRecipient, address newAvail, address governance, address pauser, address newVectorx) returns()
func (_Availbridge *AvailbridgeTransactorSession) Initialize(newFeePerByte *big.Int, newFeeRecipient common.Address, newAvail common.Address, governance common.Address, pauser common.Address, newVectorx common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.Initialize(&_Availbridge.TransactOpts, newFeePerByte, newFeeRecipient, newAvail, governance, pauser, newVectorx)
}

// ReceiveAVAIL is a paid mutator transaction binding the contract method 0xa25a59cc.
//
// Solidity: function receiveAVAIL((bytes1,bytes32,bytes32,uint32,uint32,bytes,uint64) message, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) returns()
func (_Availbridge *AvailbridgeTransactor) ReceiveAVAIL(opts *bind.TransactOpts, message IAvailBridgeMessage, input IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "receiveAVAIL", message, input)
}

// ReceiveAVAIL is a paid mutator transaction binding the contract method 0xa25a59cc.
//
// Solidity: function receiveAVAIL((bytes1,bytes32,bytes32,uint32,uint32,bytes,uint64) message, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) returns()
func (_Availbridge *AvailbridgeSession) ReceiveAVAIL(message IAvailBridgeMessage, input IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availbridge.Contract.ReceiveAVAIL(&_Availbridge.TransactOpts, message, input)
}

// ReceiveAVAIL is a paid mutator transaction binding the contract method 0xa25a59cc.
//
// Solidity: function receiveAVAIL((bytes1,bytes32,bytes32,uint32,uint32,bytes,uint64) message, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) returns()
func (_Availbridge *AvailbridgeTransactorSession) ReceiveAVAIL(message IAvailBridgeMessage, input IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availbridge.Contract.ReceiveAVAIL(&_Availbridge.TransactOpts, message, input)
}

// ReceiveERC20 is a paid mutator transaction binding the contract method 0x3a066f6e.
//
// Solidity: function receiveERC20((bytes1,bytes32,bytes32,uint32,uint32,bytes,uint64) message, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) returns()
func (_Availbridge *AvailbridgeTransactor) ReceiveERC20(opts *bind.TransactOpts, message IAvailBridgeMessage, input IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "receiveERC20", message, input)
}

// ReceiveERC20 is a paid mutator transaction binding the contract method 0x3a066f6e.
//
// Solidity: function receiveERC20((bytes1,bytes32,bytes32,uint32,uint32,bytes,uint64) message, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) returns()
func (_Availbridge *AvailbridgeSession) ReceiveERC20(message IAvailBridgeMessage, input IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availbridge.Contract.ReceiveERC20(&_Availbridge.TransactOpts, message, input)
}

// ReceiveERC20 is a paid mutator transaction binding the contract method 0x3a066f6e.
//
// Solidity: function receiveERC20((bytes1,bytes32,bytes32,uint32,uint32,bytes,uint64) message, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) returns()
func (_Availbridge *AvailbridgeTransactorSession) ReceiveERC20(message IAvailBridgeMessage, input IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availbridge.Contract.ReceiveERC20(&_Availbridge.TransactOpts, message, input)
}

// ReceiveETH is a paid mutator transaction binding the contract method 0x4d5dc664.
//
// Solidity: function receiveETH((bytes1,bytes32,bytes32,uint32,uint32,bytes,uint64) message, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) returns()
func (_Availbridge *AvailbridgeTransactor) ReceiveETH(opts *bind.TransactOpts, message IAvailBridgeMessage, input IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "receiveETH", message, input)
}

// ReceiveETH is a paid mutator transaction binding the contract method 0x4d5dc664.
//
// Solidity: function receiveETH((bytes1,bytes32,bytes32,uint32,uint32,bytes,uint64) message, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) returns()
func (_Availbridge *AvailbridgeSession) ReceiveETH(message IAvailBridgeMessage, input IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availbridge.Contract.ReceiveETH(&_Availbridge.TransactOpts, message, input)
}

// ReceiveETH is a paid mutator transaction binding the contract method 0x4d5dc664.
//
// Solidity: function receiveETH((bytes1,bytes32,bytes32,uint32,uint32,bytes,uint64) message, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) returns()
func (_Availbridge *AvailbridgeTransactorSession) ReceiveETH(message IAvailBridgeMessage, input IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availbridge.Contract.ReceiveETH(&_Availbridge.TransactOpts, message, input)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0xc07703c9.
//
// Solidity: function receiveMessage((bytes1,bytes32,bytes32,uint32,uint32,bytes,uint64) message, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) returns()
func (_Availbridge *AvailbridgeTransactor) ReceiveMessage(opts *bind.TransactOpts, message IAvailBridgeMessage, input IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "receiveMessage", message, input)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0xc07703c9.
//
// Solidity: function receiveMessage((bytes1,bytes32,bytes32,uint32,uint32,bytes,uint64) message, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) returns()
func (_Availbridge *AvailbridgeSession) ReceiveMessage(message IAvailBridgeMessage, input IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availbridge.Contract.ReceiveMessage(&_Availbridge.TransactOpts, message, input)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0xc07703c9.
//
// Solidity: function receiveMessage((bytes1,bytes32,bytes32,uint32,uint32,bytes,uint64) message, (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) input) returns()
func (_Availbridge *AvailbridgeTransactorSession) ReceiveMessage(message IAvailBridgeMessage, input IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availbridge.Contract.ReceiveMessage(&_Availbridge.TransactOpts, message, input)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Availbridge *AvailbridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Availbridge *AvailbridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.RenounceRole(&_Availbridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Availbridge *AvailbridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.RenounceRole(&_Availbridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Availbridge *AvailbridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Availbridge *AvailbridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.RevokeRole(&_Availbridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Availbridge *AvailbridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.RevokeRole(&_Availbridge.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_Availbridge *AvailbridgeTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_Availbridge *AvailbridgeSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _Availbridge.Contract.RollbackDefaultAdminDelay(&_Availbridge.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_Availbridge *AvailbridgeTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _Availbridge.Contract.RollbackDefaultAdminDelay(&_Availbridge.TransactOpts)
}

// SendAVAIL is a paid mutator transaction binding the contract method 0x1ae2ff72.
//
// Solidity: function sendAVAIL(bytes32 recipient, uint256 amount) returns()
func (_Availbridge *AvailbridgeTransactor) SendAVAIL(opts *bind.TransactOpts, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "sendAVAIL", recipient, amount)
}

// SendAVAIL is a paid mutator transaction binding the contract method 0x1ae2ff72.
//
// Solidity: function sendAVAIL(bytes32 recipient, uint256 amount) returns()
func (_Availbridge *AvailbridgeSession) SendAVAIL(recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Availbridge.Contract.SendAVAIL(&_Availbridge.TransactOpts, recipient, amount)
}

// SendAVAIL is a paid mutator transaction binding the contract method 0x1ae2ff72.
//
// Solidity: function sendAVAIL(bytes32 recipient, uint256 amount) returns()
func (_Availbridge *AvailbridgeTransactorSession) SendAVAIL(recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Availbridge.Contract.SendAVAIL(&_Availbridge.TransactOpts, recipient, amount)
}

// SendERC20 is a paid mutator transaction binding the contract method 0x758f666f.
//
// Solidity: function sendERC20(bytes32 assetId, bytes32 recipient, uint256 amount) returns()
func (_Availbridge *AvailbridgeTransactor) SendERC20(opts *bind.TransactOpts, assetId [32]byte, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "sendERC20", assetId, recipient, amount)
}

// SendERC20 is a paid mutator transaction binding the contract method 0x758f666f.
//
// Solidity: function sendERC20(bytes32 assetId, bytes32 recipient, uint256 amount) returns()
func (_Availbridge *AvailbridgeSession) SendERC20(assetId [32]byte, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Availbridge.Contract.SendERC20(&_Availbridge.TransactOpts, assetId, recipient, amount)
}

// SendERC20 is a paid mutator transaction binding the contract method 0x758f666f.
//
// Solidity: function sendERC20(bytes32 assetId, bytes32 recipient, uint256 amount) returns()
func (_Availbridge *AvailbridgeTransactorSession) SendERC20(assetId [32]byte, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Availbridge.Contract.SendERC20(&_Availbridge.TransactOpts, assetId, recipient, amount)
}

// SendETH is a paid mutator transaction binding the contract method 0x0846691f.
//
// Solidity: function sendETH(bytes32 recipient) payable returns()
func (_Availbridge *AvailbridgeTransactor) SendETH(opts *bind.TransactOpts, recipient [32]byte) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "sendETH", recipient)
}

// SendETH is a paid mutator transaction binding the contract method 0x0846691f.
//
// Solidity: function sendETH(bytes32 recipient) payable returns()
func (_Availbridge *AvailbridgeSession) SendETH(recipient [32]byte) (*types.Transaction, error) {
	return _Availbridge.Contract.SendETH(&_Availbridge.TransactOpts, recipient)
}

// SendETH is a paid mutator transaction binding the contract method 0x0846691f.
//
// Solidity: function sendETH(bytes32 recipient) payable returns()
func (_Availbridge *AvailbridgeTransactorSession) SendETH(recipient [32]byte) (*types.Transaction, error) {
	return _Availbridge.Contract.SendETH(&_Availbridge.TransactOpts, recipient)
}

// SendMessage is a paid mutator transaction binding the contract method 0x23c640e7.
//
// Solidity: function sendMessage(bytes32 recipient, bytes data) payable returns()
func (_Availbridge *AvailbridgeTransactor) SendMessage(opts *bind.TransactOpts, recipient [32]byte, data []byte) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "sendMessage", recipient, data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x23c640e7.
//
// Solidity: function sendMessage(bytes32 recipient, bytes data) payable returns()
func (_Availbridge *AvailbridgeSession) SendMessage(recipient [32]byte, data []byte) (*types.Transaction, error) {
	return _Availbridge.Contract.SendMessage(&_Availbridge.TransactOpts, recipient, data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x23c640e7.
//
// Solidity: function sendMessage(bytes32 recipient, bytes data) payable returns()
func (_Availbridge *AvailbridgeTransactorSession) SendMessage(recipient [32]byte, data []byte) (*types.Transaction, error) {
	return _Availbridge.Contract.SendMessage(&_Availbridge.TransactOpts, recipient, data)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool status) returns()
func (_Availbridge *AvailbridgeTransactor) SetPaused(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "setPaused", status)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool status) returns()
func (_Availbridge *AvailbridgeSession) SetPaused(status bool) (*types.Transaction, error) {
	return _Availbridge.Contract.SetPaused(&_Availbridge.TransactOpts, status)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool status) returns()
func (_Availbridge *AvailbridgeTransactorSession) SetPaused(status bool) (*types.Transaction, error) {
	return _Availbridge.Contract.SetPaused(&_Availbridge.TransactOpts, status)
}

// UpdateFeePerByte is a paid mutator transaction binding the contract method 0x5e6c8c6c.
//
// Solidity: function updateFeePerByte(uint256 newFeePerByte) returns()
func (_Availbridge *AvailbridgeTransactor) UpdateFeePerByte(opts *bind.TransactOpts, newFeePerByte *big.Int) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "updateFeePerByte", newFeePerByte)
}

// UpdateFeePerByte is a paid mutator transaction binding the contract method 0x5e6c8c6c.
//
// Solidity: function updateFeePerByte(uint256 newFeePerByte) returns()
func (_Availbridge *AvailbridgeSession) UpdateFeePerByte(newFeePerByte *big.Int) (*types.Transaction, error) {
	return _Availbridge.Contract.UpdateFeePerByte(&_Availbridge.TransactOpts, newFeePerByte)
}

// UpdateFeePerByte is a paid mutator transaction binding the contract method 0x5e6c8c6c.
//
// Solidity: function updateFeePerByte(uint256 newFeePerByte) returns()
func (_Availbridge *AvailbridgeTransactorSession) UpdateFeePerByte(newFeePerByte *big.Int) (*types.Transaction, error) {
	return _Availbridge.Contract.UpdateFeePerByte(&_Availbridge.TransactOpts, newFeePerByte)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address newFeeRecipient) returns()
func (_Availbridge *AvailbridgeTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "updateFeeRecipient", newFeeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address newFeeRecipient) returns()
func (_Availbridge *AvailbridgeSession) UpdateFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.UpdateFeeRecipient(&_Availbridge.TransactOpts, newFeeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address newFeeRecipient) returns()
func (_Availbridge *AvailbridgeTransactorSession) UpdateFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.UpdateFeeRecipient(&_Availbridge.TransactOpts, newFeeRecipient)
}

// UpdateTokens is a paid mutator transaction binding the contract method 0xd590bd34.
//
// Solidity: function updateTokens(bytes32[] assetIds, address[] tokenAddresses) returns()
func (_Availbridge *AvailbridgeTransactor) UpdateTokens(opts *bind.TransactOpts, assetIds [][32]byte, tokenAddresses []common.Address) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "updateTokens", assetIds, tokenAddresses)
}

// UpdateTokens is a paid mutator transaction binding the contract method 0xd590bd34.
//
// Solidity: function updateTokens(bytes32[] assetIds, address[] tokenAddresses) returns()
func (_Availbridge *AvailbridgeSession) UpdateTokens(assetIds [][32]byte, tokenAddresses []common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.UpdateTokens(&_Availbridge.TransactOpts, assetIds, tokenAddresses)
}

// UpdateTokens is a paid mutator transaction binding the contract method 0xd590bd34.
//
// Solidity: function updateTokens(bytes32[] assetIds, address[] tokenAddresses) returns()
func (_Availbridge *AvailbridgeTransactorSession) UpdateTokens(assetIds [][32]byte, tokenAddresses []common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.UpdateTokens(&_Availbridge.TransactOpts, assetIds, tokenAddresses)
}

// UpdateVectorx is a paid mutator transaction binding the contract method 0x23ecb4f2.
//
// Solidity: function updateVectorx(address newVectorx) returns()
func (_Availbridge *AvailbridgeTransactor) UpdateVectorx(opts *bind.TransactOpts, newVectorx common.Address) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "updateVectorx", newVectorx)
}

// UpdateVectorx is a paid mutator transaction binding the contract method 0x23ecb4f2.
//
// Solidity: function updateVectorx(address newVectorx) returns()
func (_Availbridge *AvailbridgeSession) UpdateVectorx(newVectorx common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.UpdateVectorx(&_Availbridge.TransactOpts, newVectorx)
}

// UpdateVectorx is a paid mutator transaction binding the contract method 0x23ecb4f2.
//
// Solidity: function updateVectorx(address newVectorx) returns()
func (_Availbridge *AvailbridgeTransactorSession) UpdateVectorx(newVectorx common.Address) (*types.Transaction, error) {
	return _Availbridge.Contract.UpdateVectorx(&_Availbridge.TransactOpts, newVectorx)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x476343ee.
//
// Solidity: function withdrawFees() returns()
func (_Availbridge *AvailbridgeTransactor) WithdrawFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Availbridge.contract.Transact(opts, "withdrawFees")
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x476343ee.
//
// Solidity: function withdrawFees() returns()
func (_Availbridge *AvailbridgeSession) WithdrawFees() (*types.Transaction, error) {
	return _Availbridge.Contract.WithdrawFees(&_Availbridge.TransactOpts)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x476343ee.
//
// Solidity: function withdrawFees() returns()
func (_Availbridge *AvailbridgeTransactorSession) WithdrawFees() (*types.Transaction, error) {
	return _Availbridge.Contract.WithdrawFees(&_Availbridge.TransactOpts)
}

// AvailbridgeDefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the Availbridge contract.
type AvailbridgeDefaultAdminDelayChangeCanceledIterator struct {
	Event *AvailbridgeDefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *AvailbridgeDefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailbridgeDefaultAdminDelayChangeCanceled)
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
		it.Event = new(AvailbridgeDefaultAdminDelayChangeCanceled)
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
func (it *AvailbridgeDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailbridgeDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailbridgeDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the Availbridge contract.
type AvailbridgeDefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_Availbridge *AvailbridgeFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*AvailbridgeDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _Availbridge.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &AvailbridgeDefaultAdminDelayChangeCanceledIterator{contract: _Availbridge.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_Availbridge *AvailbridgeFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *AvailbridgeDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _Availbridge.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailbridgeDefaultAdminDelayChangeCanceled)
				if err := _Availbridge.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

// ParseDefaultAdminDelayChangeCanceled is a log parse operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_Availbridge *AvailbridgeFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*AvailbridgeDefaultAdminDelayChangeCanceled, error) {
	event := new(AvailbridgeDefaultAdminDelayChangeCanceled)
	if err := _Availbridge.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvailbridgeDefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the Availbridge contract.
type AvailbridgeDefaultAdminDelayChangeScheduledIterator struct {
	Event *AvailbridgeDefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *AvailbridgeDefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailbridgeDefaultAdminDelayChangeScheduled)
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
		it.Event = new(AvailbridgeDefaultAdminDelayChangeScheduled)
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
func (it *AvailbridgeDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailbridgeDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailbridgeDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the Availbridge contract.
type AvailbridgeDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_Availbridge *AvailbridgeFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*AvailbridgeDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _Availbridge.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &AvailbridgeDefaultAdminDelayChangeScheduledIterator{contract: _Availbridge.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_Availbridge *AvailbridgeFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *AvailbridgeDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _Availbridge.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailbridgeDefaultAdminDelayChangeScheduled)
				if err := _Availbridge.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

// ParseDefaultAdminDelayChangeScheduled is a log parse operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_Availbridge *AvailbridgeFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*AvailbridgeDefaultAdminDelayChangeScheduled, error) {
	event := new(AvailbridgeDefaultAdminDelayChangeScheduled)
	if err := _Availbridge.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvailbridgeDefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the Availbridge contract.
type AvailbridgeDefaultAdminTransferCanceledIterator struct {
	Event *AvailbridgeDefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *AvailbridgeDefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailbridgeDefaultAdminTransferCanceled)
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
		it.Event = new(AvailbridgeDefaultAdminTransferCanceled)
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
func (it *AvailbridgeDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailbridgeDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailbridgeDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the Availbridge contract.
type AvailbridgeDefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_Availbridge *AvailbridgeFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*AvailbridgeDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _Availbridge.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &AvailbridgeDefaultAdminTransferCanceledIterator{contract: _Availbridge.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_Availbridge *AvailbridgeFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *AvailbridgeDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _Availbridge.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailbridgeDefaultAdminTransferCanceled)
				if err := _Availbridge.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

// ParseDefaultAdminTransferCanceled is a log parse operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_Availbridge *AvailbridgeFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*AvailbridgeDefaultAdminTransferCanceled, error) {
	event := new(AvailbridgeDefaultAdminTransferCanceled)
	if err := _Availbridge.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvailbridgeDefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the Availbridge contract.
type AvailbridgeDefaultAdminTransferScheduledIterator struct {
	Event *AvailbridgeDefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *AvailbridgeDefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailbridgeDefaultAdminTransferScheduled)
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
		it.Event = new(AvailbridgeDefaultAdminTransferScheduled)
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
func (it *AvailbridgeDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailbridgeDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailbridgeDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the Availbridge contract.
type AvailbridgeDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_Availbridge *AvailbridgeFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*AvailbridgeDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Availbridge.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &AvailbridgeDefaultAdminTransferScheduledIterator{contract: _Availbridge.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_Availbridge *AvailbridgeFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *AvailbridgeDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Availbridge.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailbridgeDefaultAdminTransferScheduled)
				if err := _Availbridge.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

// ParseDefaultAdminTransferScheduled is a log parse operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_Availbridge *AvailbridgeFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*AvailbridgeDefaultAdminTransferScheduled, error) {
	event := new(AvailbridgeDefaultAdminTransferScheduled)
	if err := _Availbridge.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvailbridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Availbridge contract.
type AvailbridgeInitializedIterator struct {
	Event *AvailbridgeInitialized // Event containing the contract specifics and raw log

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
func (it *AvailbridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailbridgeInitialized)
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
		it.Event = new(AvailbridgeInitialized)
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
func (it *AvailbridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailbridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailbridgeInitialized represents a Initialized event raised by the Availbridge contract.
type AvailbridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Availbridge *AvailbridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*AvailbridgeInitializedIterator, error) {

	logs, sub, err := _Availbridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AvailbridgeInitializedIterator{contract: _Availbridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Availbridge *AvailbridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AvailbridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Availbridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailbridgeInitialized)
				if err := _Availbridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Availbridge *AvailbridgeFilterer) ParseInitialized(log types.Log) (*AvailbridgeInitialized, error) {
	event := new(AvailbridgeInitialized)
	if err := _Availbridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvailbridgeMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the Availbridge contract.
type AvailbridgeMessageReceivedIterator struct {
	Event *AvailbridgeMessageReceived // Event containing the contract specifics and raw log

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
func (it *AvailbridgeMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailbridgeMessageReceived)
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
		it.Event = new(AvailbridgeMessageReceived)
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
func (it *AvailbridgeMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailbridgeMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailbridgeMessageReceived represents a MessageReceived event raised by the Availbridge contract.
type AvailbridgeMessageReceived struct {
	From      [32]byte
	To        common.Address
	MessageId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0x4ad8286366216a121ffbecdd11163a134fc364cdf7cc99aae4cc3221d8d92269.
//
// Solidity: event MessageReceived(bytes32 indexed from, address indexed to, uint256 messageId)
func (_Availbridge *AvailbridgeFilterer) FilterMessageReceived(opts *bind.FilterOpts, from [][32]byte, to []common.Address) (*AvailbridgeMessageReceivedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Availbridge.contract.FilterLogs(opts, "MessageReceived", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AvailbridgeMessageReceivedIterator{contract: _Availbridge.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0x4ad8286366216a121ffbecdd11163a134fc364cdf7cc99aae4cc3221d8d92269.
//
// Solidity: event MessageReceived(bytes32 indexed from, address indexed to, uint256 messageId)
func (_Availbridge *AvailbridgeFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *AvailbridgeMessageReceived, from [][32]byte, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Availbridge.contract.WatchLogs(opts, "MessageReceived", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailbridgeMessageReceived)
				if err := _Availbridge.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0x4ad8286366216a121ffbecdd11163a134fc364cdf7cc99aae4cc3221d8d92269.
//
// Solidity: event MessageReceived(bytes32 indexed from, address indexed to, uint256 messageId)
func (_Availbridge *AvailbridgeFilterer) ParseMessageReceived(log types.Log) (*AvailbridgeMessageReceived, error) {
	event := new(AvailbridgeMessageReceived)
	if err := _Availbridge.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvailbridgeMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the Availbridge contract.
type AvailbridgeMessageSentIterator struct {
	Event *AvailbridgeMessageSent // Event containing the contract specifics and raw log

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
func (it *AvailbridgeMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailbridgeMessageSent)
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
		it.Event = new(AvailbridgeMessageSent)
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
func (it *AvailbridgeMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailbridgeMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailbridgeMessageSent represents a MessageSent event raised by the Availbridge contract.
type AvailbridgeMessageSent struct {
	From      common.Address
	To        [32]byte
	MessageId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x06fd209663be9278f96bc53dfbf6cf3cdcf2172c38b5de30abff93ba443d653a.
//
// Solidity: event MessageSent(address indexed from, bytes32 indexed to, uint256 messageId)
func (_Availbridge *AvailbridgeFilterer) FilterMessageSent(opts *bind.FilterOpts, from []common.Address, to [][32]byte) (*AvailbridgeMessageSentIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Availbridge.contract.FilterLogs(opts, "MessageSent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AvailbridgeMessageSentIterator{contract: _Availbridge.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x06fd209663be9278f96bc53dfbf6cf3cdcf2172c38b5de30abff93ba443d653a.
//
// Solidity: event MessageSent(address indexed from, bytes32 indexed to, uint256 messageId)
func (_Availbridge *AvailbridgeFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *AvailbridgeMessageSent, from []common.Address, to [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Availbridge.contract.WatchLogs(opts, "MessageSent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailbridgeMessageSent)
				if err := _Availbridge.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x06fd209663be9278f96bc53dfbf6cf3cdcf2172c38b5de30abff93ba443d653a.
//
// Solidity: event MessageSent(address indexed from, bytes32 indexed to, uint256 messageId)
func (_Availbridge *AvailbridgeFilterer) ParseMessageSent(log types.Log) (*AvailbridgeMessageSent, error) {
	event := new(AvailbridgeMessageSent)
	if err := _Availbridge.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvailbridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Availbridge contract.
type AvailbridgePausedIterator struct {
	Event *AvailbridgePaused // Event containing the contract specifics and raw log

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
func (it *AvailbridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailbridgePaused)
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
		it.Event = new(AvailbridgePaused)
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
func (it *AvailbridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailbridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailbridgePaused represents a Paused event raised by the Availbridge contract.
type AvailbridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Availbridge *AvailbridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*AvailbridgePausedIterator, error) {

	logs, sub, err := _Availbridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AvailbridgePausedIterator{contract: _Availbridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Availbridge *AvailbridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AvailbridgePaused) (event.Subscription, error) {

	logs, sub, err := _Availbridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailbridgePaused)
				if err := _Availbridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Availbridge *AvailbridgeFilterer) ParsePaused(log types.Log) (*AvailbridgePaused, error) {
	event := new(AvailbridgePaused)
	if err := _Availbridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvailbridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Availbridge contract.
type AvailbridgeRoleAdminChangedIterator struct {
	Event *AvailbridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AvailbridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailbridgeRoleAdminChanged)
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
		it.Event = new(AvailbridgeRoleAdminChanged)
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
func (it *AvailbridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailbridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailbridgeRoleAdminChanged represents a RoleAdminChanged event raised by the Availbridge contract.
type AvailbridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Availbridge *AvailbridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AvailbridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Availbridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AvailbridgeRoleAdminChangedIterator{contract: _Availbridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Availbridge *AvailbridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AvailbridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Availbridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailbridgeRoleAdminChanged)
				if err := _Availbridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Availbridge *AvailbridgeFilterer) ParseRoleAdminChanged(log types.Log) (*AvailbridgeRoleAdminChanged, error) {
	event := new(AvailbridgeRoleAdminChanged)
	if err := _Availbridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvailbridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Availbridge contract.
type AvailbridgeRoleGrantedIterator struct {
	Event *AvailbridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *AvailbridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailbridgeRoleGranted)
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
		it.Event = new(AvailbridgeRoleGranted)
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
func (it *AvailbridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailbridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailbridgeRoleGranted represents a RoleGranted event raised by the Availbridge contract.
type AvailbridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Availbridge *AvailbridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AvailbridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _Availbridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AvailbridgeRoleGrantedIterator{contract: _Availbridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Availbridge *AvailbridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AvailbridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Availbridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailbridgeRoleGranted)
				if err := _Availbridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Availbridge *AvailbridgeFilterer) ParseRoleGranted(log types.Log) (*AvailbridgeRoleGranted, error) {
	event := new(AvailbridgeRoleGranted)
	if err := _Availbridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvailbridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Availbridge contract.
type AvailbridgeRoleRevokedIterator struct {
	Event *AvailbridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AvailbridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailbridgeRoleRevoked)
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
		it.Event = new(AvailbridgeRoleRevoked)
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
func (it *AvailbridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailbridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailbridgeRoleRevoked represents a RoleRevoked event raised by the Availbridge contract.
type AvailbridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Availbridge *AvailbridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AvailbridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _Availbridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AvailbridgeRoleRevokedIterator{contract: _Availbridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Availbridge *AvailbridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AvailbridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Availbridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailbridgeRoleRevoked)
				if err := _Availbridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Availbridge *AvailbridgeFilterer) ParseRoleRevoked(log types.Log) (*AvailbridgeRoleRevoked, error) {
	event := new(AvailbridgeRoleRevoked)
	if err := _Availbridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvailbridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Availbridge contract.
type AvailbridgeUnpausedIterator struct {
	Event *AvailbridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *AvailbridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailbridgeUnpaused)
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
		it.Event = new(AvailbridgeUnpaused)
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
func (it *AvailbridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailbridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailbridgeUnpaused represents a Unpaused event raised by the Availbridge contract.
type AvailbridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Availbridge *AvailbridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AvailbridgeUnpausedIterator, error) {

	logs, sub, err := _Availbridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AvailbridgeUnpausedIterator{contract: _Availbridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Availbridge *AvailbridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AvailbridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Availbridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailbridgeUnpaused)
				if err := _Availbridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Availbridge *AvailbridgeFilterer) ParseUnpaused(log types.Log) (*AvailbridgeUnpaused, error) {
	event := new(AvailbridgeUnpaused)
	if err := _Availbridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
