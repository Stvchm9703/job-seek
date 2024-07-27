package server

import (
	"fmt"
	"job-seek/pkg/database"
	"job-seek/pkg/protos"
	"job-seek/pkg/service_util"
	runConf "job-seek/services/job_search_service/config"
	"sync"

	surrealdb "github.com/surrealdb/surrealdb.go"

	logrus "github.com/sirupsen/logrus"

	"net"

	"google.golang.org/grpc"
)

type JobSearchServiceServerImpl struct {
	protos.UnimplementedJobSearchServiceServer
	log    *logrus.Logger
	config *runConf.ServiceConfig

	mut      *sync.Mutex
	dbClient *surrealdb.DB
	// other implement here
}

func (s JobSearchServiceServerImpl) Startup() error {
	s.log.Info("Startup Produle")

	return nil
}

func (s JobSearchServiceServerImpl) Shutdown() error {
	s.log.Info("Shutdown")

	s.log.Info("Closing database connection")
	if s.dbClient != nil {
		s.dbClient.Close()
	} else {
		s.log.Warn("Database connection is nil pointer, check the memmory leak")
	}
	s.log.Info("Database connection closed")

	return nil
}

func InitGrpcServer(config *runConf.ServiceConfig, log *logrus.Logger) (*grpc.Server, *JobSearchServiceServerImpl) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opt := service_util.CreateGrpcServerOption(&config.Server, log)
	grpcServer := grpc.NewServer(opt...)
	ssi := InitService(config, log)
	protos.RegisterJobSearchServiceServer(grpcServer, ssi)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	service_util.BeforeGracefulStop(grpcServer, ssi.Shutdown, log)
	return grpcServer, &ssi
}

func InitService(config *runConf.ServiceConfig, log *logrus.Logger) JobSearchServiceServerImpl {
	dbClient, err := database.InitConnection(&config.SurrealDBService, "development")
	log.Info("Connecting to database")
	if err != nil {
		log.WithFields(map[string]interface{}{
			"error": err,
			"host":  config.SurrealDBService.Host,
		}).Fatal("Failed to connect to database in InitService")
	}
	log.Info("Connected to database")

	ssi := JobSearchServiceServerImpl{
		mut:      &sync.Mutex{},
		log:      log,
		config:   config,
		dbClient: dbClient,
	}
	ssi.Startup()
	return ssi
}
