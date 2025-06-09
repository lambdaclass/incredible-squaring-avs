package main

import (
	"context"
	"log"
	"os"
	"slices"

	"github.com/urfave/cli"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	commonincredible "github.com/Layr-Labs/incredible-sorting-avs/common"
	cstaskmanager "github.com/Layr-Labs/incredible-sorting-avs/contracts/bindings/IncredibleSortingTaskManager"
	"github.com/Layr-Labs/incredible-sorting-avs/core/config"
	"github.com/Layr-Labs/incredible-sorting-avs/operator"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{config.ConfigFileFlag}
	app.Name = "credible-squaring-operator"
	app.Usage = "Credible Squaring Operator"
	app.Description = "Service that reads numbers onchain, squares, signs, and sends them to the aggregator."

	app.Action = operatorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed. Message:", err)
	}
}

func operatorMain(ctx *cli.Context) error {

	log.Println("Initializing Operator")
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)

	opConfig := &operator.Config{}
	err := commonincredible.ReadTomlConfig(configPath, opConfig)

	logger, err := sdklogging.NewZapLogger(sdklogging.Production) // Change here if want to change logging level
	if err != nil {
		return err
	}

	logger.Infof("Config is %#v", opConfig)

	logger.Info("initializing operator")

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSortingTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	operatorConfig := opConfig.Config

	calculator := sdkoperator.NewFunctionResponseCalculator(SortNumbers)

	failingFunction, err := sdkoperator.NewFailingResponseCalculator(calculator, 10, []uint32{})
	if err != nil {
		logger.Fatalf(err.Error())
	}

	if err != nil {
		logger.Fatalf(err.Error())
	}

	operator, err := sdkoperator.NewOperator(
		logger,
		operatorConfig,
		taskManagerAbi,
		failingFunction,
		nil,
	)
	if err != nil {
		logger.Fatalf(err.Error())
	}
	log.Println("initialized operator")

	log.Println("starting operator")
	err = <-operator.Start(context.Background())
	if err != nil {
		logger.Fatalf(err.Error())
	}
	log.Println("started operator")

	return nil

}

func SortNumbers(taskIndex uint32, numbersToBeSorted []uint32) ([]uint32, error) {
	sorted := make([]uint32, len(numbersToBeSorted))
	copy(sorted, numbersToBeSorted)

	slices.Sort(sorted)
	return sorted, nil
}
