package server

import (
	"fmt"
	"job-seek/pkg/database"
	"job-seek/pkg/protos"
	"job-seek/pkg/service_util"
	runConf "job-seek/services/prediction_service/config"
	"sync"

	"gorm.io/gorm"

	logrus "github.com/sirupsen/logrus"

	"net"

	"google.golang.org/grpc"
)

type PredictionServiceServerImpl struct {
	protos.UnimplementedPredictionServiceServer
	log    *logrus.Logger
	config *runConf.ServiceConfig

	mut      *sync.Mutex
	dbClient *gorm.DB
	// other implement here
}

func (s PredictionServiceServerImpl) Startup() error {
	s.log.Info("Startup Produle")
	return nil
}

func (s PredictionServiceServerImpl) Shutdown() error {
	s.log.Info("Shutdown")
	return nil
}

func InitGrpcServer(config *runConf.ServiceConfig, log *logrus.Logger) (*grpc.Server, *PredictionServiceServerImpl) {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opt := service_util.CreateGrpcServerOption(&config.Server, log)
	grpcServer := grpc.NewServer(opt...)
	ssi := InitService(config, log)
	protos.RegisterPredictionServiceServer(grpcServer, ssi)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	service_util.BeforeGracefulStop(grpcServer, ssi.Shutdown, log)
	return grpcServer, &ssi

}

func InitService(config *runConf.ServiceConfig, log *logrus.Logger) PredictionServiceServerImpl {
	dbClient, err := database.InitConnection(&config.SurrealDBService, "development")
	log.Info("Connecting to database")
	if err != nil {
		log.WithFields(map[string]interface{}{
			"error": err,
			"host":  config.SurrealDBService.Host,
		}).Fatal("Failed to connect to database in InitService")
	}
	log.Info("Connected to database")

	ssi := PredictionServiceServerImpl{
		mut:      &sync.Mutex{},
		log:      log,
		config:   config,
		dbClient: dbClient,
	}
	ssi.Startup()
	return ssi
}
