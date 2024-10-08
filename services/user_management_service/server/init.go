package server

import (
	"fmt"
	"job-seek/pkg/database"
	"job-seek/pkg/protos"
	"job-seek/pkg/service_util"
	runConf "job-seek/services/user_management_service/config"
	"net"
	"sync"

	logrus "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type UserManagementServiceServerImpl struct {
	protos.UnimplementedUserManagementServiceServer
	log    *logrus.Logger
	config *runConf.ServiceConfig

	mut      *sync.Mutex
	dbClient *gorm.DB
	// other implement here
}

func (s UserManagementServiceServerImpl) Startup() error {
	s.log.Info("Startup Produle")
	return nil
}

func (s UserManagementServiceServerImpl) Shutdown() error {
	s.log.Info("Shutdown")

	s.log.Info("Closing database connection")
	if s.dbClient != nil {
		dbInstance, _ := s.dbClient.DB()
		dbInstance.Close()
	} else {
		s.log.Warn("Database connection is nil pointer, check the memmory leak")
	}
	s.log.Info("Database connection closed")

	return nil
}

func InitGrpcServer(config *runConf.ServiceConfig, log *logrus.Logger) (*grpc.Server, *UserManagementServiceServerImpl) {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opt := service_util.CreateGrpcServerOption(&config.Server, log)
	grpcServer := grpc.NewServer(opt...)
	ssi := InitService(config, log)
	protos.RegisterUserManagementServiceServer(grpcServer, ssi)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	service_util.BeforeGracefulStop(grpcServer, ssi.Shutdown, log)
	return grpcServer, &ssi

}

func InitService(config *runConf.ServiceConfig, log *logrus.Logger) UserManagementServiceServerImpl {
	dbClient, err := database.InitConnection(&config.DBService, "development")
	log.Info("Connecting to database")
	if err != nil {
		log.WithFields(map[string]interface{}{
			"error": err,
			"host":  config.DBService,
		}).Fatal("Failed to connect to database in InitService")
	}
	log.Info("Connected to database")

	ssi := UserManagementServiceServerImpl{
		mut:      &sync.Mutex{},
		log:      log,
		config:   config,
		dbClient: dbClient,
	}
	ssi.Startup()
	return ssi
}
