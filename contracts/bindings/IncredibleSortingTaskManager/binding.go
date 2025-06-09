// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIncredibleSortingTaskManager

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerTypesNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerTypesNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerTypesQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerTypesQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IIncredibleSortingTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSortingTaskManagerTask struct {
	NumbersToBeSorted         []uint32
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IIncredibleSortingTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSortingTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	SortedNumbers      []uint32
}

// IIncredibleSortingTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSortingTaskManagerTaskResponseMetadata struct {
	TaskRespondedBlock uint32
	HashOfNonSigners   [32]byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractIncredibleSortingTaskManagerMetaData contains all meta data concerning the ContractIncredibleSortingTaskManager contract.
var ContractIncredibleSortingTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WADS_TO_SLASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureCheckerTypes.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureCheckerTypes.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"numbersToBeSorted\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_slasher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_serviceManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"instantSlasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSortingTaskManager.Task\",\"components\":[{\"name\":\"numbersToBeSorted\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSortingTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sortedNumbers\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSortingTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSortingTaskManager.Task\",\"components\":[{\"name\":\"numbersToBeSorted\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSortingTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sortedNumbers\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureCheckerTypes.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serviceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSortingTaskManager.Task\",\"components\":[{\"name\":\"numbersToBeSorted\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSortingTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sortedNumbers\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSortingTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BitmapValueTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayLengthTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayNotOrdered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECAddFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECMulFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpModFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputEmptyQuorumNumbers\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputNonSignerLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSPairingKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidQuorumApkHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReferenceBlocknumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonSignerPubkeysNotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRegistryCoordinatorOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ScalarTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StaleStakesForbidden\",\"inputs\":[]}]",
	Bin: "0x61014080604052346101d857606081614b5380380380916100208285610283565b8339810103126101d85780516001600160a01b038116908181036101d85760208301516001600160a01b038116938482036101d857604001519363ffffffff851685036101d857156102745760805260a052604051636830483560e01b8152602081600481855afa9081156101e4575f91610231575b5060c052604051632efa2ca360e11b815290602090829060049082905afa9081156101e4575f916101ef575b5060e05260c05160405163df5cf72360e01b815290602090829060049082906001600160a01b03165afa9081156101e4575f9161019e575b50610100526101205260405161489890816102bb82396080518181816102ca015281816111eb015281816119f801526122ba015260a0518181816109630152818161151b01528181612f7c0152818161306501526135f7015260c0518181816114d70152818161336701526134b2015260e05181818161149301526132a30152610100518181816121b601526131780152610120518181816105950152611d6f0152f35b90506020813d6020116101dc575b816101b960209383610283565b810103126101d857516001600160a01b03811681036101d8575f6100fa565b5f80fd5b3d91506101ac565b6040513d5f823e3d90fd5b90506020813d602011610229575b8161020a60209383610283565b810103126101d857516001600160a01b03811681036101d8575f6100c2565b3d91506101fd565b90506020813d60201161026c575b8161024c60209383610283565b810103126101d857516001600160a01b03811681036101d8576004610096565b3d915061023f565b6339b190bb60e11b5f5260045ffd5b601f909101601f19168101906001600160401b038211908210176102a657604052565b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f3560e01c8063136439dd1461029a578063171f1d5b146102955780631ad43189146101e6578063245a7bfc146102905780632cb223d51461028b5780632d89f6fc1461028657806331b36bd9146102815780633563b0d11461027c5780633998fdd314610277578063416c7e5e146102725780634d2b57fe1461026d5780634f739f7414610268578063595c6a67146102635780635a2d7f021461025e5780635ac86ab7146102595780635c155662146102545780635c975abb1461024f5780635decc3f51461024a5780635df459461461024557806368304835146102405780636d14a9871461023b5780636efb463614610236578063715018a61461023157806372d18e8d1461021d5780637afa1eed1461022c578063877cd13014610227578063886f1195146102225780638b00ce7c1461021d5780638da5cb5b146102185780639b290e9814610213578063b98d09081461020e578063bb0c914114610209578063ca02d91514610204578063ca8aa7c7146101ff578063cc2a9a5b146101fa578063cefdc1d4146101f5578063df5cf723146101f0578063f2fde38b146101eb578063f5c9899d146101e6578063f63c5bab146101e15763fabc1cbc146101dc575f80fd5b612291565b612276565b610579565b6121e5565b6121a1565b61205d565b611f1d565b611ef5565b611c50565b611a99565b611a77565b611a4f565b611a27565b6118f5565b6119e3565b61195c565b611918565b61189a565b6117ed565b611506565b6114c2565b61147e565b611440565b611423565b6112cc565b611260565b611233565b6111c0565b610ccc565b610a91565b610931565b6108ff565b610885565b6106db565b610633565b6105fa565b6105b9565b610511565b3461035a57602036600319011261035a5760043560405163237dfb4760e11b8152336004820152906020826024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9182156103555761032492610310915f91610326575b50612393565b61031f606654828116146123a9565b613ff1565b005b610348915060203d60201161034e575b61034081836103ad565b810190612373565b5f61030a565b503d610336565b612388565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b604081019081106001600160401b0382111761038d57604052565b61035e565b608081019081106001600160401b0382111761038d57604052565b90601f801991011681019081106001600160401b0382111761038d57604052565b604051906103de610100836103ad565b565b604051906103de6040836103ad565b604051906103de6060836103ad565b906103de60405192836103ad565b60409060e319011261035a576040519061042582610372565b60e4358252610104356020830152565b919082604091031261035a5760405161044d81610372565b6020808294803584520135910152565b9080601f8301121561035a57604051916104786040846103ad565b82906040810192831161035a57905b8282106104945750505090565b8135815260209182019101610487565b90608060631983011261035a576040516104bd81610372565b60206104d882946104cf81606461045d565b845260a461045d565b910152565b919060808382031261035a5760206104d8604051926104fb84610372565b60408496610509838261045d565b86520161045d565b3461035a5761012036600319011261035a57600435604036602319011261035a57610569604091825161054381610372565b60243581526044356020820152610559366104a4565b906105633661040c565b926123fd565b8251911515825215156020820152f35b3461035a575f36600319011261035a57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461035a575f36600319011261035a5760cd546040516001600160a01b039091168152602090f35b63ffffffff81160361035a57565b35906103de826105e1565b3461035a57602036600319011261035a5763ffffffff60043561061c816105e1565b165f5260cb602052602060405f2054604051908152f35b3461035a57602036600319011261035a5763ffffffff600435610655816105e1565b165f5260ca602052602060405f2054604051908152f35b6001600160a01b0381160361035a57565b6001600160401b03811161038d5760051b60200190565b90602080835192838152019201905f5b8181106106b15750505090565b82518452602093840193909201916001016106a4565b9060206106d8928181520190610694565b90565b3461035a57604036600319011261035a576004356106f88161066c565b602435906001600160401b03821161035a573660238301121561035a578160040135916107248361067d565b9261073260405194856103ad565b8084526024602085019160051b8301019136831161035a57602401905b8282106107735761076f6107638686612567565b604051918291826106c7565b0390f35b6020809183356107828161066c565b81520191019061074f565b6001600160401b03811161038d57601f01601f191660200190565b9291926107b48261078d565b916107c260405193846103ad565b82948184528183011161035a578281602093845f960137010152565b9080602083519182815201916020808360051b8301019401925f915b83831061080957505050505090565b9091929394601f19828203018352855190602080835192838152019201905f905b80821061084957505050602080600192970193019301919392906107fa565b909192602060606001926001600160601b0360408851868060a01b0381511684528581015186850152015116604082015201940192019061082a565b3461035a57606036600319011261035a576004356108a28161066c565b6024356001600160401b03811161035a573660238201121561035a5761076f916108d96108eb9236906024816004013591016107a8565b604435916108e6836105e1565b6127a5565b6040519182916020835260208301906107de565b3461035a575f36600319011261035a5760d1546040516001600160a01b039091168152602090f35b8015150361035a57565b3461035a57602036600319011261035a5760043561094e81610927565b604051638da5cb5b60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610355575f916109c3575b506001600160a01b031633036109b45761032490614484565b637070f3b160e11b5f5260045ffd5b6109e5915060203d6020116109eb575b6109dd81836103ad565b81019061262c565b5f61099b565b503d6109d3565b9080601f8301121561035a578135610a098161067d565b92610a1760405194856103ad565b81845260208085019260051b82010192831161035a57602001905b828210610a3f5750505090565b8135815260209182019101610a32565b60206040818301928281528451809452019201905f5b818110610a725750505090565b82516001600160a01b0316845260209384019390920191600101610a65565b3461035a57604036600319011261035a57600435610aae8161066c565b6024356001600160401b03811161035a57610acd9036906004016109f2565b610ad78151612505565b916001600160a01b03165f5b8251811015610b7457806020610afc610b1c9386612544565b5160405180948192630a5aec1960e21b8352600483019190602083019252565b0381865afa91821561035557600192610b50915f91610b56575b50610b418388612544565b6001600160a01b039091169052565b01610ae3565b610b6e915060203d81116109eb576109dd81836103ad565b5f610b36565b6040518061076f8682610a4f565b9181601f8401121561035a578235916001600160401b03831161035a576020838186019501011161035a57565b9181601f8401121561035a578235916001600160401b03831161035a576020808501948460051b01011161035a57565b90602080835192838152019201905f5b818110610bfc5750505090565b825163ffffffff16845260209384019390920191600101610bef565b90602082526060610c66610c51610c3b84516080602088015260a0870190610bdf565b6020850151868203601f19016040880152610bdf565b6040840151858203601f190184870152610bdf565b910151916080601f1982840301910152815180825260208201916020808360051b8301019401925f915b838310610c9f57505050505090565b9091929394602080610cbd600193601f198682030187528951610bdf565b97019301930191939290610c90565b3461035a57608036600319011261035a57600435610ce98161066c565b602435610cf5816105e1565b6044356001600160401b03811161035a57610d14903690600401610b82565b92906064356001600160401b03811161035a57610d35903690600401610baf565b939092610d40612acf565b50604051636830483560e01b8152926001600160a01b039190911690602084600481855afa938415610355575f9461119f575b50610d7c612acf565b936040516361c8a12f60e11b81525f8180610d9c8b8b8a60048501612b76565b0381875afa908115610355575f91611185575b5085526040516340e03a8160e11b81526001600160a01b039190911691905f8180610ddf8c868a60048501612bcd565b0381865afa908115610355575f9161116b575b506040860152610e0188612641565b97606086019889525f5b60ff8116828110156110aa575f610e39828d610e268e612505565b905190610e338383612544565b52612544565b505f5b8b8110610ebf5750610e4d81612505565b905f5b8d828210610e7c57610e7795949250610e719391505190610e338383612544565b50612bfe565b610e0b565b90610eb9610ea4610e9a83610e948960019751612544565b51612544565b5163ffffffff1690565b610eae8387612544565b9063ffffffff169052565b01610e50565b610ed2818d8d959d9e969e9c949c612c14565b3560208a610ee4610e9a858751612544565b6040516304ec635160e01b8152600481019490945263ffffffff9182166024850152166044830152816064818c5afa801561035557610f8b6001610f7e8f9493858b8d85975f9461106a575b50610f6892610f6292610f5492610f4f898060c01b0388161515612c43565b612c59565b356001600160f81b03191690565b60f81c90565b6001600160c01b0391821660ff919091161c1690565b166001600160c01b031690565b14610fa3575b506001019a929a999199989098610e3c565b996020829b610fc6610f62610f54610fbd8f978b8b612c14565b35938b8d612c59565b60405163dd9846b960e01b8152600481019290925260ff16602482015263ffffffff939093166044840152826064818b5afa908115610355576110248f611029938f84906001975f93611032575b50610e9490610eae939451612544565b612c7a565b9990508a610f91565b610eae93509061105b610e949260203d8111611063575b61105381836103ad565b810190612c65565b935090611014565b503d611049565b610f5491945092610f6292611098610f689560203d81116110a3575b61109081836103ad565b810190612c24565b959250925092610f30565b503d611086565b604051632efa2ca360e11b81528490899089866020836004818e5afa908115610355576110f8955f94859361114a575b5060405163354952a360e21b81529687948593849360048501612c88565b03916001600160a01b03165afa80156103555761076f925f91611128575b50602082015260405191829182610c18565b61114491503d805f833e61113c81836103ad565b810190612af3565b83611116565b61116491935060203d6020116109eb576109dd81836103ad565b91876110da565b61117f91503d805f833e61113c81836103ad565b5f610df2565b61119991503d805f833e61113c81836103ad565b5f610daf565b6111b991945060203d6020116109eb576109dd81836103ad565b925f610d73565b3461035a575f36600319011261035a5760405163237dfb4760e11b81523360048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156103555761122b915f916103265750612393565b610324613fbd565b3461035a575f36600319011261035a57602060405167016345785d8a00008152f35b60ff81160361035a57565b3461035a57602036600319011261035a576020600160ff60043561128381611255565b161b806066541614604051908152f35b60206040818301928281528451809452019201905f5b8181106112b65750505090565b82518452602093840193909201916001016112a9565b3461035a57606036600319011261035a576004356112e98161066c565b6024356001600160401b03811161035a576113089036906004016109f2565b60443591611315836105e1565b6040516361c8a12f60e11b8152906001600160a01b03165f828061133d868860048401612cad565b0381845afa918215610355575f92611407575b5061135b8351612505565b935f5b84518110156113f9576113718186612544565b5190602083611383610e9a8489612544565b6040516304ec635160e01b8152600481019590955263ffffffff918216602486015216604484015282606481875afa8015610355576001925f916113db575b50828060c01b03166113d48289612544565b520161135e565b6113f3915060203d81116110a35761109081836103ad565b5f6113c2565b6040518061076f8882611293565b61141c9192503d805f833e61113c81836103ad565b905f611350565b3461035a575f36600319011261035a576020606654604051908152f35b3461035a57602036600319011261035a5763ffffffff600435611462816105e1565b165f5260cc602052602060ff60405f2054166040519015158152f35b3461035a575f36600319011261035a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461035a575f36600319011261035a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461035a575f36600319011261035a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b9291906115568161067d565b9361156460405195866103ad565b602085838152019160051b810192831161035a57905b82821061158657505050565b602080918335611595816105e1565b81520191019061157a565b9080601f8301121561035a578160206106d89335910161154a565b81601f8201121561035a5780356115d18161067d565b926115df60405194856103ad565b81845260208085019260061b8401019281841161035a57602001915b838310611609575050505090565b60206040916116188486610435565b8152019201916115fb565b9080601f8301121561035a57813561163a8161067d565b9261164860405194856103ad565b81845260208085019260051b8201019183831161035a5760208201905b83821061167457505050505090565b81356001600160401b03811161035a57602091611696878480948801016115a0565b815201910190611665565b9190916101808184031261035a576116b76103ce565b9281356001600160401b03811161035a57816116d49184016115a0565b845260208201356001600160401b03811161035a57816116f59184016115bb565b602085015260408201356001600160401b03811161035a57816117199184016115bb565b604085015261172b81606084016104dd565b606085015261173d8160e08401610435565b60808501526101208201356001600160401b03811161035a57816117629184016115a0565b60a08501526101408201356001600160401b03811161035a57816117879184016115a0565b60c08501526101608201356001600160401b03811161035a576117aa9201611623565b60e0830152565b90602080835192838152019201905f5b8181106117ce5750505090565b82516001600160601b03168452602093840193909201916001016117c1565b3461035a57608036600319011261035a576004356024356001600160401b03811161035a57611820903690600401610b82565b909160443561182e816105e1565b606435926001600160401b03841161035a576118909461185561185b9536906004016116a1565b93612e9f565b60405192839260408452602061187c825160408088015260808701906117b1565b910151848203603f190160608601526117b1565b9060208301520390f35b3461035a575f36600319011261035a576118b2614660565b603380546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461035a575f36600319011261035a57602063ffffffff60c95416604051908152f35b3461035a575f36600319011261035a5760ce546040516001600160a01b039091168152602090f35b9081608091031261035a5790565b9081604091031261035a5790565b3461035a5760a036600319011261035a576004356001600160401b03811161035a5761198c903690600401611940565b6024356001600160401b03811161035a576119ab90369060040161194e565b90604036604319011261035a576084356001600160401b03811161035a57610324926119dd60449236906004016115bb565b92613a91565b3461035a575f36600319011261035a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461035a575f36600319011261035a576033546040516001600160a01b039091168152602090f35b3461035a575f36600319011261035a5760cf546040516001600160a01b039091168152602090f35b3461035a575f36600319011261035a57602060ff609754166040519015158152f35b3461035a57606036600319011261035a576004356001600160401b03811161035a57611ac9903690600401610baf565b9060243591611ad7836105e1565b6044356001600160401b03811161035a57611af6903690600401610b82565b60ce54909491906001600160a01b03163303611c015761032494611b49611b5093611b2e611beb97611b26613b91565b97369161154a565b86524363ffffffff16602087015263ffffffff166060860152565b36916107a8565b60408201526040516020810190611b7981611b6b8585613bb5565b03601f1981018352826103ad565b519020611ba2611b8e60c95463ffffffff1690565b63ffffffff165f5260ca60205260405f2090565b5560c95463ffffffff16907f493f5d8b10b0662a82ac5d061fc5deeaaf6522ae2deff4840e4f02aa3556052e60405180611be363ffffffff86169482613bb5565b0390a26139a3565b63ffffffff1663ffffffff1960c954161760c955565b60405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608490fd5b3461035a57606036600319011261035a576004356001600160401b03811161035a57611c80903690600401611940565b6024356001600160401b03811161035a57611c9f90369060040161194e565b906044356001600160401b03811161035a57611cbf9036906004016116a1565b60cd549092906001600160a01b03163303611eb057611ce260208394930161373c565b91611dc3611cf36040860186613a5f565b929094611d31611d056060890161373c565b97604051611d1b81611b6b602082019485613c1b565b519020611d2a611b8e8861373c565b5414613cbd565b611d5b611d54611d408761373c565b63ffffffff165f5260cb60205260405f2090565b5415613d2f565b8363ffffffff431696611da5611d9d611d947f0000000000000000000000000000000000000000000000000000000000000000866139d3565b63ffffffff1690565b891115613d90565b6040516020810190611dbb81611b6b8b85613df2565b519020612e9f565b919060ff5f9616955b828110611e4e577fa85d3cc594af947f29831473e254ff1342149a6a4001ce7ed5d3cfea37d9348e868686611e0e611e026103e0565b63ffffffff9094168452565b60208301526040516020810190611e2a81611b6b868686613ec3565b519020611e39611d408361373c565b55611e4960405192839283613ec3565b0390a1005b80611eaa611e86611e81611e75611e686001968851612544565b516001600160601b031690565b6001600160601b031690565b613e03565b611ea3611e758b611e9e611e688760208b0151612544565b613e2f565b1115613e52565b01611dcc565b60405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606490fd5b3461035a575f36600319011261035a5760d0546040516001600160a01b039091168152602090f35b3461035a5760c036600319011261035a57600435611f3a8161066c565b611fba602435611f498161066c565b604435611f558161066c565b606435611f618161066c565b60843591611f6e8361066c565b60a43593611f7b8561066c565b5f5496611fa060ff60088a901c16158099819a612038575b8115612018575b50613ef1565b87611fb1600160ff195f5416175f55565b61200157613f54565b611fc057005b611fce61ff00195f54165f55565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498908060208101611e49565b61201361010061ff00195f5416175f55565b613f54565b303b1591508161202a575b505f611f9a565b60ff1660011490505f612023565b600160ff8216109150611f93565b6040906106d89392815281602082015201906107de565b3461035a57606036600319011261035a5760043561207a8161066c565b602435604435612089816105e1565b6120ca6120946124e3565b928061209f85612537565b526040516361c8a12f60e11b81526001600160a01b0386169490925f91849182918760048401612cad565b0381875afa93841561035557836120f4611d94610e9a612129986020975f91612187575b50612537565b92604051968794859384936304ec635160e01b85526004850163ffffffff604092959493606083019683521660208201520152565b03915afa801561035557612158925f91612168575b506001600160c01b03169261215284614700565b906127a5565b9061076f60405192839283612046565b612181915060203d6020116110a35761109081836103ad565b5f61213e565b61219b91503d805f833e61113c81836103ad565b5f6120ee565b3461035a575f36600319011261035a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461035a57602036600319011261035a576004356122028161066c565b61220a614660565b6001600160a01b0381161561222257610324906146b8565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b3461035a575f36600319011261035a57602060405160648152f35b3461035a57602036600319011261035a5760043560405163755b36bd60e11b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610355575f91612354575b506001600160a01b03163303612345576123136066541982198116146123a9565b806066556040519081527f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c60203392a2005b63794821ff60e01b5f5260045ffd5b61236d915060203d6020116109eb576109dd81836103ad565b5f6122f2565b9081602091031261035a57516106d881610927565b6040513d5f823e3d90fd5b1561239a57565b631d77d47760e21b5f5260045ffd5b156123b057565b63c61dca5d60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b9060028110156123e45760051b0190565b6123bf565b634e487b7160e01b5f52601260045260245ffd5b6124d96124b66124df956124b06124a985875160208901518a515160208c51015160208d016020815151915101519189519360208b0151956040519760208901998a5260208a015260408901526060880152608087015260a086015260c085015260e084015261010083015261248081610120840103601f1981018352826103ad565b5190207f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001900690565b8096614067565b906140ad565b926124b06124cb6124c561410f565b94614206565b916124d4614322565b614067565b91614356565b9091565b604080519091906124f483826103ad565b6001815291601f1901366020840137565b9061250f8261067d565b61251c60405191826103ad565b828152809261252d601f199161067d565b0190602036910137565b8051156123e45760200190565b80518210156123e45760209160051b010190565b9081602091031261035a575190565b9190916125748351612505565b925f5b8151811015612627578060206125a06125936125c99486612544565b516001600160a01b031690565b6040516309aa152760e11b81526001600160a01b03909116600482015292839081906024820190565b03816001600160a01b0388165afa8015610355576001925f916125f9575b506125f28288612544565b5201612577565b61261a915060203d8111612620575b61261281836103ad565b810190612558565b5f6125e7565b503d612608565b505050565b9081602091031261035a57516106d88161066c565b9061264b8261067d565b61265860405191826103ad565b8281526020819361266b601f199161067d565b0191015f5b82811061267c57505050565b606082820152602001612670565b9081518110156123e4570160200190565b60208183031261035a578051906001600160401b03821161035a57019080601f8301121561035a5781516126ce8161067d565b926126dc60405194856103ad565b81845260208085019260051b82010192831161035a57602001905b8282106127045750505090565b81518152602091820191016126f7565b9061271e8261067d565b61272b60405191826103ad565b828152809261273c601f199161067d565b015f5b81811061274b57505050565b6040519060608201918083106001600160401b0384111761038d576020926040525f81525f838201525f60408201528282860101520161273f565b9081602091031261035a57516001600160601b038116810361035a5790565b604051636830483560e01b815293919291906001600160a01b0316602085600481845afa948515610355575f95612aae575b50604051634f4c91e160e11b815294602086600481855afa918215610355576004965f93612a8c575b5060209060405197888092632efa2ca360e11b82525afa958615610355575f96612a6b575b506128338593929551612641565b945f935b8051851015612a615761285e610f62612850878461268a565b516001600160f81b03191690565b604051638902624560e01b815260ff8216600482015263ffffffff88166024820152909490925f846044816001600160a01b0385165afa938415610355575f94612a3d575b506128ae8451612714565b6128b8888b612544565b526128c3878a612544565b505f5b8451811015612a2c578060206128df6129019388612544565b518d60405180809681946308f6629d60e31b8352600483019190602083019252565b03916001600160a01b03165afa918215610355575f92612a0c575b506129278187612544565b518a60208a612936858b612544565b5160405163fa28c62760e01b8152600481019190915260ff91909116602482015263ffffffff929092166044830152816064816001600160a01b038d165afa938415610355576129c98c8f6129c46001986129d59789975f926129dc575b506129af6129a06103ef565b6001600160a01b039098168852565b60208701526001600160601b03166040860152565b612544565b5190610e338383612544565b50016128c6565b6129fe91925060203d8111612a05575b6129f681836103ad565b810190612786565b905f612994565b503d6129ec565b612a2591925060203d81116109eb576109dd81836103ad565b905f61291c565b506001909601959094509150612837565b612a5a9194503d805f833e612a5281836103ad565b81019061269b565b925f6128a3565b5050509350505090565b612a8591965060203d6020116109eb576109dd81836103ad565b945f612825565b6020919350612aa790823d84116109eb576109dd81836103ad565b9290612800565b612ac891955060203d6020116109eb576109dd81836103ad565b935f6127d7565b60405190612adc82610392565b606080838181528160208201528160408201520152565b60208183031261035a578051906001600160401b03821161035a57019080601f8301121561035a578151612b268161067d565b92612b3460405194856103ad565b81845260208085019260051b82010192831161035a57602001905b828210612b5c5750505090565b602080918351612b6b816105e1565b815201910190612b4f565b63ffffffff909116815260406020820181905281018390526001600160fb1b03831161035a5760609260051b809284830137010190565b908060209392818452848401375f828201840152601f01601f1916010190565b60409063ffffffff6106d895931681528160208201520191612bad565b634e487b7160e01b5f52601160045260245ffd5b60ff1660ff8114612c0f5760010190565b612bea565b91908110156123e45760051b0190565b9081602091031261035a57516001600160c01b038116810361035a5790565b15612c4a57565b6325ec6c1f60e01b5f5260045ffd5b908210156123e4570190565b9081602091031261035a57516106d8816105e1565b5f198114612c0f5760010190565b91612ca660209263ffffffff92969596604086526040860191612bad565b9416910152565b60409063ffffffff6106d894931681528160208201520190610694565b60405190612cd782610372565b60606020838281520152565b15612cea57565b62f8202d60e51b5f5260045ffd5b15612cff57565b6343714afd60e01b5f5260045ffd5b15612d1557565b635f832f4160e01b5f5260045ffd5b15612d2b57565b634b874f4560e01b5f5260045ffd5b9081602091031261035a57516106d881611255565b5f19810191908211612c0f57565b15612d6457565b633fdc650560e21b5f5260045ffd5b9060018201809211612c0f57565b9060028201809211612c0f57565b9060038201809211612c0f57565b9060048201809211612c0f57565b9060058201809211612c0f57565b91908201809211612c0f57565b15612dcd57565b63affc5edb60e01b5f5260045ffd5b9081602091031261035a575167ffffffffffffffff198116810361035a5790565b15612e0457565b63e1310aed60e01b5f5260045ffd5b906001600160601b03809116911603906001600160601b038211612c0f57565b15612e3a57565b6367988d3360e01b5f5260045ffd5b15612e5057565b63ab1b236b60e01b5f5260045ffd5b60049163ffffffff60e01b9060e01b1681520160208251919201905f5b818110612e895750505090565b8251845260209384019390920191600101612e7c565b949392909193612ead612cca565b50612eb9851515612ce3565b60408401515185148061372e575b80613720575b80613712575b612edc90612cf8565b612eee60208501515185515114612d0e565b612f0563ffffffff431663ffffffff841610612d24565b612f0d6103e0565b5f81525f602082015292612f1f612cca565b612f2887612505565b6020820152612f3687612505565b8152612f40612cca565b92612f4f602088015151612505565b8452612f5f602088015151612505565b602085810191909152604051639aa1653d60e01b815290816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561035557612fc8915f916136e3575b50612fc3368b876107a8565b6144c2565b985f965b6020890151805189101561313a5760208861302f610e9a8c6130278f96868e61300c612ff9868095612544565b5180515f526020015160205260405f2090565b6130198484840151612544565b5282613107575b0151612544565b519551612544565b6040516304ec635160e01b8152600481019490945263ffffffff9182166024850152166044830152816064816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa918215610355576124b08a6130dc8f6130d58f8460208f926130cc936130c48460019e6130e29e5f916130ea575b508f8060c01b03169251612544565b520151612544565b51938d51612544565b51166144ed565b9061451e565b970196612fcc565b6131019150863d81116110a35761109081836103ad565b5f6130b5565b6131356131178484840151612544565b5161312e8484015161312887612d4f565b90612544565b5110612d5d565b613020565b5090959794965061314f9198939299506145db565b9161315c60975460ff1690565b9081156136db576040516318891fd760e31b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610355575f916136bc575b5091905b5f925b81841061320d575050505050926131f46131ef6131e86132079585611b6b98608060606020990151920151926123fd565b9190612e33565b612e49565b0151604051928391602083019586612e5f565b51902090565b92989596909399919794878b888c888d6135b6575b610e9a8260a0613262610f62610f548461326a9761325c61324e612ff98f9c604060209f9e0151612544565b67ffffffffffffffff191690565b9b612c59565b970151612544565b604051631a2f32ab60e21b815260ff95909516600486015263ffffffff9182166024860152166044840152826064816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156103555761332e610e9a8f958f906133268f978f96848f61332060c096613319848f60209f90613020610f5499604093610f629c5f91613588575b5067ffffffffffffffff19918216911614612dfd565b51906140ad565b9c612c59565b960151612544565b604051636414a62b60e11b815260ff94909416600485015263ffffffff9182166024850152166044830152816064816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610355576133bb918c8f925f92613564575b5060206133ad92930151612544565b906001600160601b03169052565b6133db8c6133ad8c6133d4611e68826020860151612544565b9251612544565b5f985f5b60208a01515181101561354b578b8d61341d89613410610f62610f54868f896134089151612544565b519487612c59565b60ff161c60019081161490565b61342c575b50506001016133df565b8a8a6134ae859f948f9686610e948f9360e0613465610e9a95602061345d610f62610f54839f61346e9c8991612c59565b9a0151612544565b519b0151612544565b60405163795f4a5760e11b815260ff909316600484015263ffffffff93841660248401526044830196909652919094166064850152839081906084820190565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610355578f61351a908f936001959486955f92613525575b506135146133ad9293519361350f611e688487612544565b612e13565b92612544565b019a90508b8d613422565b6133ad92506135446135149160203d8111612a05576129f681836103ad565b92506134f7565b5093919796996001919699509a94929a019291906131b7565b6133ad9250613581602091823d8111612a05576129f681836103ad565b925061339e565b60206135a992503d81116135af575b6135a181836103ad565b810190612ddc565b5f613303565b503d613597565b6135f394506135d09250610f6291610f5491602095612c59565b60405163124d062160e11b815260ff909116600482015291829081906024820190565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156103555760208961326a8f938f60a08f97610f62610f548f8f9061325c61324e612ff98f60408b96918f8893610e9a9f6136779061367d936132629f5f92613693575b5063ffffffff809116931690612db9565b11612dc6565b5050505050509750505050505092935050613222565b602063ffffffff92935082916136b4913d81116126205761261281836103ad565b929150613666565b6136d5915060203d6020116110635761105381836103ad565b5f6131b0565b5f91906131b4565b613705915060203d60201161370b575b6136fd81836103ad565b810190612d3a565b5f612fb7565b503d6136f3565b5060e0840151518514612ed3565b5060c0840151518514612ecd565b5060a0840151518514612ec7565b356106d8816105e1565b903590601e198136030182121561035a57018035906001600160401b03821161035a57602001918160051b3603831361035a57565b1561378257565b60405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608490fd5b9035601e198236030181121561035a5701602081359101916001600160401b03821161035a578160051b3603831361035a57565b916020908281520191905f5b81811061381e5750505090565b90919260208060019263ffffffff8735613837816105e1565b168152019401929101613811565b90604061386c6106d89363ffffffff813561385f816105e1565b16845260208101906137d1565b9190928160208201520191613805565b906020613896604092959495606085526060850190613845565b9463ffffffff81356138a7816105e1565b16828501520135910152565b156138ba57565b60405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608490fd5b1561392c57565b60405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a490fd5b63ffffffff60019116019063ffffffff8211612c0f57565b63ffffffff60649116019063ffffffff8211612c0f57565b9063ffffffff8091169116019063ffffffff8211612c0f57565b156139f457565b60405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608490fd5b903590601e198136030182121561035a57018035906001600160401b03821161035a5760200191813603831361035a57565b9250611d94613b50613b6393613b20613b5594613ac3613abc823599613ab68b6105e1565b80613746565b369161154a565b50613ae7613adf8963ffffffff165f5260cb60205260405f2090565b54151561377b565b613aff8863ffffffff165f5260cb60205260405f2090565b5490604051613b1781611b6b8760208301958661387c565b519020146138b3565b613b4b613b45613b3e8863ffffffff165f5260cc60205260405f2090565b5460ff1690565b15613925565b61373c565b6139bb565b63ffffffff431611156139ed565b63ffffffff3391167ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb055f80a3565b60405190613b9e82610392565b5f6060838181528260208201528160408201520152565b60208152608063ffffffff60606020613bd88651858388015260a0870190610bdf565b8382880151166040870152816040880151601f1988840301858901528051918291828552018484015e5f838284010152601f801991011601019401511691015290565b60208152613c3d613c2c83806137d1565b6080602085015260a0840191613805565b9063ffffffff6020840135613c51816105e1565b1660408201526040830135601e198436030181121561035a57830191602083359301906001600160401b03841161035a57833603821361035a576060613cab613cb2926106d896608095601f198884030185890152612bad565b95016105ef565b63ffffffff16910152565b15613cc457565b60405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608490fd5b15613d3657565b60405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608490fd5b15613d9757565b60405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608490fd5b9060206106d8928181520190613845565b90606482029180830460641490151715612c0f57565b90600682029180830460061490151715612c0f57565b906001600160601b03809116911602906001600160601b038216918203612c0f57565b15613e5957565b608460405162461bcd60e51b815260206004820152604060248201527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152fd5b906020613edd604092959495606085526060850190613845565b9463ffffffff815116828501520151910152565b15613ef857565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b613f5d906146b8565b60cd80546001600160a01b03199081166001600160a01b039384161790915560ce805482169383169390931790925560d0805483169382169390931790925560cf805482169383169390931790925560d180549092169216919091179055565b5f196066556040515f1981527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2565b806066556040519081527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2565b6040519061403082610372565b5f6020838281520152565b6040519061018061404c81846103ad565b368337565b604051906140606020836103ad565b6020368337565b91906040906060614076614023565b948592602085519261408885856103ad565b8436853780518452015160208301528482015260076107cf195a01fa156140ab57565bfe5b6020929160806040926140be614023565b958693818651936140cf86866103ad565b85368637805185520151828401528051868401520151606082015260066107cf195a01fa80156140ab571561410057565b63d4b68fd760e01b5f5260045ffd5b60405161411b81610372565b604090815161412a83826103ad565b823682378152602082519161413f84846103ad565b833684370152805161415182826103ad565b7f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c281527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed60208201528151906141a783836103ad565b7f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208301526141fc835193846103ad565b8252602082015290565b5f5160206148435f395f51905f529061421d614023565b505f919006602060c0835b61431d575f935f5160206148435f395f51905f526003818681818009090860405161425385826103ad565b8436823784818560405161426782826103ad565b813682378381528360208201528360408201528560608201527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808201525f5160206148435f395f51905f5260a082015260056107cf195a01fa80156140ab576142d19061482c565b519161431d575f5160206148435f395f51905f528280091461430857505f5160206148435f395f51905f5260015f94089293614228565b929350506143146103e0565b92835282015290565b6123e9565b61432a614023565b5060405161433781610372565b600181526002602082015290565b90600c8110156123e45760051b0190565b9392909161436460406103fe565b948552602085015261437660406103fe565b918252602082015261438661403b565b925f5b600281106143b3575050506020610180926143a2614051565b93849160086201d4c0fa9151151590565b806143bf600192613e19565b6143c982856123d3565b51516143d58289614345565b5260206143e283866123d3565b5101516143f76143f183612d73565b89614345565b5261440282866123d3565b5151516144116143f183612d81565b5261442761441f83876123d3565b515160200190565b516144346143f183612d8f565b52602061444183876123d3565b510151516144516143f183612d9d565b5261447d6144776144706020614467868a6123d3565b51015160200190565b5192612dab565b88614345565b5201614389565b60207f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc91151560ff196097541660ff821617609755604051908152a1565b9060016144d060ff936147b4565b928392161b11156144de5790565b63ca95733360e01b5f5260045ffd5b805f915b6144f9575090565b5f198101818111612c0f5761ffff9116911661ffff8114612c0f5760010190806144f1565b90614527614023565b5061ffff8116906102008210156145cc57600182146145c7576145486103e0565b5f81525f602082015292906001905f925b61ffff831685101561456d57505050505090565b600161ffff831660ff86161c8116146145a7575b600161459d6145928360ff946140ad565b9460011b61fffe1690565b9401169291614559565b94600161459d6145926145bc8960ff956140ad565b989350505050614581565b505090565b637fc4ea7d60e11b5f5260045ffd5b6145e3614023565b50805190811580614654575b156146105750506040516146046040826103ad565b5f81525f602082015290565b60205f5160206148435f395f51905f52910151065f5160206148435f395f51905f52035f5160206148435f395f51905f528111612c0f57604051916141fc83610372565b506020810151156145ef565b6033546001600160a01b0316330361467457565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b61ffff61470c826144ed565b166147168161078d565b9061472460405192836103ad565b808252614733601f199161078d565b013660208301375f5f5b8251821080614793575b1561478c576001811b8416614765575b61476090612c7a565b61473d565b9060016147609160ff60f81b8460f81b165f1a614782828761268a565b5301919050614757565b5050905090565b506101008110614747565b156147a557565b631019106960e31b5f5260045ffd5b9061010082511161481d5781511561481857602082015160019060f81c81901b5b8351821015614813576001906147fe6147f4610f62612850868961268a565b60ff600191161b90565b9061480a81831161479e565b179101906147d5565b925050565b5f9150565b637da54e4760e11b5f5260045ffd5b1561483357565b63d51edae360e01b5f5260045ffdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a26469706673582212207bea0e9df8efa186fc55f7cb35a2ef9454185b478e7ced4615c0af5265a011a364736f6c634300081b0033",
}

// ContractIncredibleSortingTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIncredibleSortingTaskManagerMetaData.ABI instead.
var ContractIncredibleSortingTaskManagerABI = ContractIncredibleSortingTaskManagerMetaData.ABI

// ContractIncredibleSortingTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIncredibleSortingTaskManagerMetaData.Bin instead.
var ContractIncredibleSortingTaskManagerBin = ContractIncredibleSortingTaskManagerMetaData.Bin

// DeployContractIncredibleSortingTaskManager deploys a new Ethereum contract, binding an instance of ContractIncredibleSortingTaskManager to it.
func DeployContractIncredibleSortingTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _pauserRegistry common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractIncredibleSortingTaskManager, error) {
	parsed, err := ContractIncredibleSortingTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIncredibleSortingTaskManagerBin), backend, _registryCoordinator, _pauserRegistry, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIncredibleSortingTaskManager{ContractIncredibleSortingTaskManagerCaller: ContractIncredibleSortingTaskManagerCaller{contract: contract}, ContractIncredibleSortingTaskManagerTransactor: ContractIncredibleSortingTaskManagerTransactor{contract: contract}, ContractIncredibleSortingTaskManagerFilterer: ContractIncredibleSortingTaskManagerFilterer{contract: contract}}, nil
}

// ContractIncredibleSortingTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractIncredibleSortingTaskManager struct {
	ContractIncredibleSortingTaskManagerCaller     // Read-only binding to the contract
	ContractIncredibleSortingTaskManagerTransactor // Write-only binding to the contract
	ContractIncredibleSortingTaskManagerFilterer   // Log filterer for contract events
}

// ContractIncredibleSortingTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIncredibleSortingTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSortingTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIncredibleSortingTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSortingTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIncredibleSortingTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSortingTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIncredibleSortingTaskManagerSession struct {
	Contract     *ContractIncredibleSortingTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                         // Call options to use throughout this session
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractIncredibleSortingTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIncredibleSortingTaskManagerCallerSession struct {
	Contract *ContractIncredibleSortingTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                               // Call options to use throughout this session
}

// ContractIncredibleSortingTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIncredibleSortingTaskManagerTransactorSession struct {
	Contract     *ContractIncredibleSortingTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                               // Transaction auth options to use throughout this session
}

// ContractIncredibleSortingTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIncredibleSortingTaskManagerRaw struct {
	Contract *ContractIncredibleSortingTaskManager // Generic contract binding to access the raw methods on
}

// ContractIncredibleSortingTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIncredibleSortingTaskManagerCallerRaw struct {
	Contract *ContractIncredibleSortingTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIncredibleSortingTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIncredibleSortingTaskManagerTransactorRaw struct {
	Contract *ContractIncredibleSortingTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIncredibleSortingTaskManager creates a new instance of ContractIncredibleSortingTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSortingTaskManager(address common.Address, backend bind.ContractBackend) (*ContractIncredibleSortingTaskManager, error) {
	contract, err := bindContractIncredibleSortingTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManager{ContractIncredibleSortingTaskManagerCaller: ContractIncredibleSortingTaskManagerCaller{contract: contract}, ContractIncredibleSortingTaskManagerTransactor: ContractIncredibleSortingTaskManagerTransactor{contract: contract}, ContractIncredibleSortingTaskManagerFilterer: ContractIncredibleSortingTaskManagerFilterer{contract: contract}}, nil
}

// NewContractIncredibleSortingTaskManagerCaller creates a new read-only instance of ContractIncredibleSortingTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSortingTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIncredibleSortingTaskManagerCaller, error) {
	contract, err := bindContractIncredibleSortingTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerCaller{contract: contract}, nil
}

// NewContractIncredibleSortingTaskManagerTransactor creates a new write-only instance of ContractIncredibleSortingTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSortingTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIncredibleSortingTaskManagerTransactor, error) {
	contract, err := bindContractIncredibleSortingTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerTransactor{contract: contract}, nil
}

// NewContractIncredibleSortingTaskManagerFilterer creates a new log filterer instance of ContractIncredibleSortingTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSortingTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIncredibleSortingTaskManagerFilterer, error) {
	contract, err := bindContractIncredibleSortingTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerFilterer{contract: contract}, nil
}

// bindContractIncredibleSortingTaskManager binds a generic wrapper to an already deployed contract.
func bindContractIncredibleSortingTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIncredibleSortingTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSortingTaskManager.Contract.ContractIncredibleSortingTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.ContractIncredibleSortingTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.ContractIncredibleSortingTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSortingTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSortingTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSortingTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSortingTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSortingTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// WADSTOSLASH is a free data retrieval call binding the contract method 0x5a2d7f02.
//
// Solidity: function WADS_TO_SLASH() view returns(uint256)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) WADSTOSLASH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "WADS_TO_SLASH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WADSTOSLASH is a free data retrieval call binding the contract method 0x5a2d7f02.
//
// Solidity: function WADS_TO_SLASH() view returns(uint256)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) WADSTOSLASH() (*big.Int, error) {
	return _ContractIncredibleSortingTaskManager.Contract.WADSTOSLASH(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// WADSTOSLASH is a free data retrieval call binding the contract method 0x5a2d7f02.
//
// Solidity: function WADS_TO_SLASH() view returns(uint256)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) WADSTOSLASH() (*big.Int, error) {
	return _ContractIncredibleSortingTaskManager.Contract.WADSTOSLASH(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Aggregator(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Aggregator(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSortingTaskManager.Contract.AllTaskHashes(&_ContractIncredibleSortingTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSortingTaskManager.Contract.AllTaskHashes(&_ContractIncredibleSortingTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSortingTaskManager.Contract.AllTaskResponses(&_ContractIncredibleSortingTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSortingTaskManager.Contract.AllTaskResponses(&_ContractIncredibleSortingTaskManager.CallOpts, arg0)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) AllocationManager() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.AllocationManager(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) AllocationManager() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.AllocationManager(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.BlsApkRegistry(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.BlsApkRegistry(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (IBLSSignatureCheckerTypesQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerTypesQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerTypesQuorumStakeTotals)).(*IBLSSignatureCheckerTypesQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (IBLSSignatureCheckerTypesQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleSortingTaskManager.Contract.CheckSignatures(&_ContractIncredibleSortingTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (IBLSSignatureCheckerTypesQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleSortingTaskManager.Contract.CheckSignatures(&_ContractIncredibleSortingTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Delegation(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Delegation(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) Generator() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Generator(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Generator(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetBatchOperatorFromId(&_ContractIncredibleSortingTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetBatchOperatorFromId(&_ContractIncredibleSortingTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetBatchOperatorId(&_ContractIncredibleSortingTaskManager.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetBatchOperatorId(&_ContractIncredibleSortingTaskManager.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleSortingTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleSortingTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetOperatorState(&_ContractIncredibleSortingTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetOperatorState(&_ContractIncredibleSortingTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSortingTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSortingTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractIncredibleSortingTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractIncredibleSortingTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleSortingTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// InstantSlasher is a free data retrieval call binding the contract method 0x9b290e98.
//
// Solidity: function instantSlasher() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) InstantSlasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "instantSlasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InstantSlasher is a free data retrieval call binding the contract method 0x9b290e98.
//
// Solidity: function instantSlasher() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) InstantSlasher() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.InstantSlasher(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// InstantSlasher is a free data retrieval call binding the contract method 0x9b290e98.
//
// Solidity: function instantSlasher() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) InstantSlasher() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.InstantSlasher(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleSortingTaskManager.Contract.LatestTaskNum(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleSortingTaskManager.Contract.LatestTaskNum(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Owner(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Owner(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Paused(&_ContractIncredibleSortingTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Paused(&_ContractIncredibleSortingTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Paused0(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Paused0(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.PauserRegistry(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.PauserRegistry(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "serviceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) ServiceManager() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.ServiceManager(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) ServiceManager() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.ServiceManager(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.StakeRegistry(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleSortingTaskManager.Contract.StakeRegistry(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractIncredibleSortingTaskManager.Contract.StaleStakesForbidden(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractIncredibleSortingTaskManager.Contract.StaleStakesForbidden(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleSortingTaskManager.Contract.TaskNumber(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleSortingTaskManager.Contract.TaskNumber(&_ContractIncredibleSortingTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleSortingTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractIncredibleSortingTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleSortingTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractIncredibleSortingTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractIncredibleSortingTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractIncredibleSortingTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleSortingTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractIncredibleSortingTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleSortingTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xbb0c9141.
//
// Solidity: function createNewTask(uint32[] numbersToBeSorted, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, numbersToBeSorted []uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.contract.Transact(opts, "createNewTask", numbersToBeSorted, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xbb0c9141.
//
// Solidity: function createNewTask(uint32[] numbersToBeSorted, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) CreateNewTask(numbersToBeSorted []uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.CreateNewTask(&_ContractIncredibleSortingTaskManager.TransactOpts, numbersToBeSorted, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xbb0c9141.
//
// Solidity: function createNewTask(uint32[] numbersToBeSorted, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactorSession) CreateNewTask(numbersToBeSorted []uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.CreateNewTask(&_ContractIncredibleSortingTaskManager.TransactOpts, numbersToBeSorted, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address initialOwner, address _aggregator, address _generator, address _allocationManager, address _slasher, address _serviceManager) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _allocationManager common.Address, _slasher common.Address, _serviceManager common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.contract.Transact(opts, "initialize", initialOwner, _aggregator, _generator, _allocationManager, _slasher, _serviceManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address initialOwner, address _aggregator, address _generator, address _allocationManager, address _slasher, address _serviceManager) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) Initialize(initialOwner common.Address, _aggregator common.Address, _generator common.Address, _allocationManager common.Address, _slasher common.Address, _serviceManager common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Initialize(&_ContractIncredibleSortingTaskManager.TransactOpts, initialOwner, _aggregator, _generator, _allocationManager, _slasher, _serviceManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address initialOwner, address _aggregator, address _generator, address _allocationManager, address _slasher, address _serviceManager) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactorSession) Initialize(initialOwner common.Address, _aggregator common.Address, _generator common.Address, _allocationManager common.Address, _slasher common.Address, _serviceManager common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Initialize(&_ContractIncredibleSortingTaskManager.TransactOpts, initialOwner, _aggregator, _generator, _allocationManager, _slasher, _serviceManager)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Pause(&_ContractIncredibleSortingTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Pause(&_ContractIncredibleSortingTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.PauseAll(&_ContractIncredibleSortingTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.PauseAll(&_ContractIncredibleSortingTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x877cd130.
//
// Solidity: function raiseAndResolveChallenge((uint32[],uint32,bytes,uint32) task, (uint32,uint32[]) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IIncredibleSortingTaskManagerTask, taskResponse IIncredibleSortingTaskManagerTaskResponse, taskResponseMetadata IIncredibleSortingTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x877cd130.
//
// Solidity: function raiseAndResolveChallenge((uint32[],uint32,bytes,uint32) task, (uint32,uint32[]) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) RaiseAndResolveChallenge(task IIncredibleSortingTaskManagerTask, taskResponse IIncredibleSortingTaskManagerTaskResponse, taskResponseMetadata IIncredibleSortingTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleSortingTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x877cd130.
//
// Solidity: function raiseAndResolveChallenge((uint32[],uint32,bytes,uint32) task, (uint32,uint32[]) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactorSession) RaiseAndResolveChallenge(task IIncredibleSortingTaskManagerTask, taskResponse IIncredibleSortingTaskManagerTaskResponse, taskResponseMetadata IIncredibleSortingTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleSortingTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.RenounceOwnership(&_ContractIncredibleSortingTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.RenounceOwnership(&_ContractIncredibleSortingTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xca02d915.
//
// Solidity: function respondToTask((uint32[],uint32,bytes,uint32) task, (uint32,uint32[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IIncredibleSortingTaskManagerTask, taskResponse IIncredibleSortingTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xca02d915.
//
// Solidity: function respondToTask((uint32[],uint32,bytes,uint32) task, (uint32,uint32[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) RespondToTask(task IIncredibleSortingTaskManagerTask, taskResponse IIncredibleSortingTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.RespondToTask(&_ContractIncredibleSortingTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xca02d915.
//
// Solidity: function respondToTask((uint32[],uint32,bytes,uint32) task, (uint32,uint32[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactorSession) RespondToTask(task IIncredibleSortingTaskManagerTask, taskResponse IIncredibleSortingTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.RespondToTask(&_ContractIncredibleSortingTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.SetStaleStakesForbidden(&_ContractIncredibleSortingTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.SetStaleStakesForbidden(&_ContractIncredibleSortingTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.TransferOwnership(&_ContractIncredibleSortingTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.TransferOwnership(&_ContractIncredibleSortingTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Unpause(&_ContractIncredibleSortingTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSortingTaskManager.Contract.Unpause(&_ContractIncredibleSortingTaskManager.TransactOpts, newPausedStatus)
}

// ContractIncredibleSortingTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerInitializedIterator struct {
	Event *ContractIncredibleSortingTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSortingTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSortingTaskManagerInitialized)
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
		it.Event = new(ContractIncredibleSortingTaskManagerInitialized)
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
func (it *ContractIncredibleSortingTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSortingTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSortingTaskManagerInitialized represents a Initialized event raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractIncredibleSortingTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerInitializedIterator{contract: _ContractIncredibleSortingTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSortingTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSortingTaskManagerInitialized)
				if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractIncredibleSortingTaskManagerInitialized, error) {
	event := new(ContractIncredibleSortingTaskManagerInitialized)
	if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSortingTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerNewTaskCreatedIterator struct {
	Event *ContractIncredibleSortingTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSortingTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSortingTaskManagerNewTaskCreated)
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
		it.Event = new(ContractIncredibleSortingTaskManagerNewTaskCreated)
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
func (it *ContractIncredibleSortingTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSortingTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSortingTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IIncredibleSortingTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x493f5d8b10b0662a82ac5d061fc5deeaaf6522ae2deff4840e4f02aa3556052e.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint32[],uint32,bytes,uint32) task)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSortingTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerNewTaskCreatedIterator{contract: _ContractIncredibleSortingTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x493f5d8b10b0662a82ac5d061fc5deeaaf6522ae2deff4840e4f02aa3556052e.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint32[],uint32,bytes,uint32) task)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSortingTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSortingTaskManagerNewTaskCreated)
				if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0x493f5d8b10b0662a82ac5d061fc5deeaaf6522ae2deff4840e4f02aa3556052e.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint32[],uint32,bytes,uint32) task)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractIncredibleSortingTaskManagerNewTaskCreated, error) {
	event := new(ContractIncredibleSortingTaskManagerNewTaskCreated)
	if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSortingTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerOwnershipTransferredIterator struct {
	Event *ContractIncredibleSortingTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSortingTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSortingTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractIncredibleSortingTaskManagerOwnershipTransferred)
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
func (it *ContractIncredibleSortingTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSortingTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSortingTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractIncredibleSortingTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerOwnershipTransferredIterator{contract: _ContractIncredibleSortingTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSortingTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSortingTaskManagerOwnershipTransferred)
				if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractIncredibleSortingTaskManagerOwnershipTransferred, error) {
	event := new(ContractIncredibleSortingTaskManagerOwnershipTransferred)
	if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSortingTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerPausedIterator struct {
	Event *ContractIncredibleSortingTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSortingTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSortingTaskManagerPaused)
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
		it.Event = new(ContractIncredibleSortingTaskManagerPaused)
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
func (it *ContractIncredibleSortingTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSortingTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSortingTaskManagerPaused represents a Paused event raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleSortingTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerPausedIterator{contract: _ContractIncredibleSortingTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSortingTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSortingTaskManagerPaused)
				if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) ParsePaused(log types.Log) (*ContractIncredibleSortingTaskManagerPaused, error) {
	event := new(ContractIncredibleSortingTaskManagerPaused)
	if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdate)
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
func (it *ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractIncredibleSortingTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractIncredibleSortingTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSortingTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractIncredibleSortingTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSortingTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSortingTaskManagerTaskChallengedSuccessfully)
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
		it.Event = new(ContractIncredibleSortingTaskManagerTaskChallengedSuccessfully)
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
func (it *ContractIncredibleSortingTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSortingTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSortingTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleSortingTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractIncredibleSortingTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSortingTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSortingTaskManagerTaskChallengedSuccessfully)
				if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractIncredibleSortingTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractIncredibleSortingTaskManagerTaskChallengedSuccessfully)
	if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfully)
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
		it.Event = new(ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfully)
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
func (it *ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractIncredibleSortingTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
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

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractIncredibleSortingTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSortingTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerTaskCompletedIterator struct {
	Event *ContractIncredibleSortingTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSortingTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSortingTaskManagerTaskCompleted)
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
		it.Event = new(ContractIncredibleSortingTaskManagerTaskCompleted)
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
func (it *ContractIncredibleSortingTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSortingTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSortingTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSortingTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerTaskCompletedIterator{contract: _ContractIncredibleSortingTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSortingTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSortingTaskManagerTaskCompleted)
				if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractIncredibleSortingTaskManagerTaskCompleted, error) {
	event := new(ContractIncredibleSortingTaskManagerTaskCompleted)
	if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSortingTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerTaskRespondedIterator struct {
	Event *ContractIncredibleSortingTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSortingTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSortingTaskManagerTaskResponded)
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
		it.Event = new(ContractIncredibleSortingTaskManagerTaskResponded)
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
func (it *ContractIncredibleSortingTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSortingTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSortingTaskManagerTaskResponded represents a TaskResponded event raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerTaskResponded struct {
	TaskResponse         IIncredibleSortingTaskManagerTaskResponse
	TaskResponseMetadata IIncredibleSortingTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xa85d3cc594af947f29831473e254ff1342149a6a4001ce7ed5d3cfea37d9348e.
//
// Solidity: event TaskResponded((uint32,uint32[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractIncredibleSortingTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerTaskRespondedIterator{contract: _ContractIncredibleSortingTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xa85d3cc594af947f29831473e254ff1342149a6a4001ce7ed5d3cfea37d9348e.
//
// Solidity: event TaskResponded((uint32,uint32[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSortingTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSortingTaskManagerTaskResponded)
				if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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

// ParseTaskResponded is a log parse operation binding the contract event 0xa85d3cc594af947f29831473e254ff1342149a6a4001ce7ed5d3cfea37d9348e.
//
// Solidity: event TaskResponded((uint32,uint32[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractIncredibleSortingTaskManagerTaskResponded, error) {
	event := new(ContractIncredibleSortingTaskManagerTaskResponded)
	if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSortingTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerUnpausedIterator struct {
	Event *ContractIncredibleSortingTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSortingTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSortingTaskManagerUnpaused)
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
		it.Event = new(ContractIncredibleSortingTaskManagerUnpaused)
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
func (it *ContractIncredibleSortingTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSortingTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSortingTaskManagerUnpaused represents a Unpaused event raised by the ContractIncredibleSortingTaskManager contract.
type ContractIncredibleSortingTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleSortingTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSortingTaskManagerUnpausedIterator{contract: _ContractIncredibleSortingTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSortingTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSortingTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSortingTaskManagerUnpaused)
				if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSortingTaskManager *ContractIncredibleSortingTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractIncredibleSortingTaskManagerUnpaused, error) {
	event := new(ContractIncredibleSortingTaskManagerUnpaused)
	if err := _ContractIncredibleSortingTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
