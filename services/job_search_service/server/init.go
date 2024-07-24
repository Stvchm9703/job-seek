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

	s.log.Info("Connecting to database")
	var err error
	s.dbClient, err = database.InitConnection(&s.config.SurrealDBService, "development")
	if err != nil {
		s.log.Fatalf("Failed to connect to database: %v", err)
		return err
	}
	s.log.Info("Connected to database")

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

func InitGrpcServer(config *runConf.ServiceConfig, log *logrus.Logger) *grpc.Server {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opt := service_util.CreateGrpcServerOption(&config.Server, log)
	grpcServer := grpc.NewServer(opt...)
	ssi := JobSearchServiceServerImpl{
		mut:    &sync.Mutex{},
		log:    log,
		config: config,
	}
	ssi.Startup()
	protos.RegisterJobSearchServiceServer(grpcServer, ssi)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	service_util.BeforeGracefulStop(grpcServer, ssi.Shutdown, log)
	return grpcServer

}
