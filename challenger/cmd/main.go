package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	"github.com/Layr-Labs/incredible-sorting-avs/challenger"
	commonincredible "github.com/Layr-Labs/incredible-sorting-avs/common"
	cstaskmanager "github.com/Layr-Labs/incredible-sorting-avs/contracts/bindings/IncredibleSortingTaskManager"
	"github.com/Layr-Labs/incredible-sorting-avs/core/config"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string
)

func main() {

	app := cli.NewApp()
	app.Flags = config.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "credible-squaring-challenger"
	app.Usage = "Credible Squaring Challenger"
	app.Description = "Service that challenges wrong response to the task."

	app.Action = challengerMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func challengerMain(ctx *cli.Context) error {

	log.Println("Initializing Challenger...")
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	challengerConfig := &challenger.Config{}
	err := commonincredible.ReadTomlConfig(configPath, challengerConfig)

	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		return err
	}

	ethRpcClient, err := ethclient.Dial(challengerConfig.EthHttpUrl)

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSortingTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	privateKeyHex := ctx.String("ecdsa-private-key")
	if privateKeyHex == "" {
		logger.Fatal("Missing required flag: --ecdsa-private-key")
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		logger.Errorf("Cannot parse ECDSA private key", "err", err)
		return err
	}

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethRpcClient, ecdsaPrivateKey)

	cfg := sdkchallenger.Config{
		EthWsUrl:   challengerConfig.EthWsUrl,
		EthHttpUrl: challengerConfig.EthHttpUrl,
	}

	taskManagerAddr := challengerConfig.TaskManagerAddress
	challengerRaiser, err := taskmanager.NewTaskManagerFromAbi[[]uint32, []uint32](
		common.HexToAddress(taskManagerAddr),
		taskManagerAbi,
		txMgr,
		ethRpcClient,
	)

	calculator := sdkoperator.NewFunctionResponseCalculator(SortNumbers)
	squareValidation := sdkchallenger.ResponseValidationFunctionFromResponseCalculator(calculator, slices.Equal)

	indexingChallengerProcessor, err := sdkchallenger.NewIndexingProcessor(
		logger,
		squareValidation,
		challengerRaiser,
	)

	challenger, err := sdkchallenger.NewChallenger(
		logger,
		cfg,
		taskManagerAbi,
		indexingChallengerProcessor,
	)
	if err != nil {
		logger.Fatalf("Failed to create challenger from config: %v", err)
	}

	err = <-challenger.Start(context.Background())
	if err != nil {
		return err
	}

	return nil

}

func SortNumbers(taskIndex uint32, numbersToBeSorted []uint32) ([]uint32, error) {
	sorted := make([]uint32, len(numbersToBeSorted))
	copy(sorted, numbersToBeSorted)

	slices.Sort(sorted)
	return sorted, nil
}
