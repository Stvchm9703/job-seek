package cmd

import (
	"job-seek/services/job_extendsion_service/config"
	"job-seek/services/job_extendsion_service/server"

	logger "job-seek/pkg/log"

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
	log = logger.InitLog("job_extendsion_service", verbose)

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

	runtimeServer = server.InitGrpcServer(&config.RuntimeConfig, log)

}

func ServerTestRun(verboseLevel int) {
	initLogger(3, verboseLevel)
	log.Info("Start Server Test Run")
}

func ServerRun(verboseLevel int) {
	initLogger(0, verboseLevel)
	log.Info("Start Server")

}
