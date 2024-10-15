package cmd

import (
	"fmt"
	logger "job-seek/pkg/log"
	"job-seek/services/prediction_service/config"
	"job-seek/services/prediction_service/server"
	twp "job-seek/services/prediction_service/server_twirp"

	logrus "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	log           *logrus.Logger
	runtimeServer *grpc.Server
)

func initLogger(minLevel int, verboseLevel int) {
	verbose := verboseLevel
	if verbose < minLevel && minLevel != 0 {
		verbose = minLevel
	}
	log = logger.InitLog("prediction_service", verbose)

	// configPrint, _ := json.Marshal(config.RuntimeConfig)
	log.WithFields(logrus.Fields{
		"verbose": verbose,
		"config":  config.RuntimeConfig,
	}).Info("Run config")

}

func ServerDryRun(verboseLevel int) {
	initLogger(3, verboseLevel)
	log.Info("Start Server DryRun")
	// ctx, ctxCancel := context.WithCancel(context.Background())
	// defer ctxCancel()
	log.Info("Created context")

	runtimeServer, _ = server.InitGrpcServer(&config.RuntimeConfig, log)

}

func ServerTestRun(verboseLevel int) {
	initLogger(3, verboseLevel)
	log.Info("Start Server Test Run")
	fmt.Println("Start Server Test Run")
	// var manualShutdown func()
	server.InitService(&config.RuntimeConfig, log)

	log.Info("Server Start Test Run")
	log.Info("run prediction_service")

	// add the test case here

	// manualShutdown()
	log.Info("Server Test Run End")
}

func ServerRun(verboseLevel int) {
	initLogger(0, verboseLevel)
	log.Info("Start Server")
	runtimeServer, _ = server.InitGrpcServer(&config.RuntimeConfig, log)
}

func ServerTwirpRun(verboseLevel int) {
	initLogger(0, verboseLevel)
	log.Info("Start Server")
	twp.InitTwirpServer(&config.RuntimeConfig, log)
}
