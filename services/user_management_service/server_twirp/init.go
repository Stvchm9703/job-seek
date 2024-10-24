package server

import (
	"fmt"
	database "job-seek/pkg/database_v1"
	"job-seek/pkg/protos"
	"job-seek/pkg/service_util"
	runConf "job-seek/services/user_management_service/config"
	"net"
	"net/http"
	"sync"

	logrus "github.com/sirupsen/logrus"
	"github.com/surrealdb/surrealdb.go"
)

type UserManagementServiceServerImpl struct {
	protos.UnimplementedUserManagementServiceServer
	log    *logrus.Logger
	config *runConf.ServiceConfig

	mut      *sync.Mutex
	dbClient *surrealdb.DB
	// dbClient *gorm.DB
	// other implement here
}

func (s *UserManagementServiceServerImpl) Startup() error {
	s.log.Info("Startup Produle")
	return nil
}

func (s *UserManagementServiceServerImpl) Shutdown() error {
	s.log.Info("Shutdown")

	s.log.Info("Closing database connection")
	if s.dbClient != nil {
		s.dbClient.Close()
		// dbInstance, _ := s.dbClient.DB()
		// dbInstance.Close()
	} else {
		s.log.Warn("Database connection is nil pointer, check the memmory leak")
	}
	s.log.Info("Database connection closed")

	return nil
}

func InitService(config *runConf.ServiceConfig, log *logrus.Logger) UserManagementServiceServerImpl {
	dbClient, err := database.InitConnection(&config.SurrealDBService, "development")
	log.Info("Connecting to database")
	if err != nil {
		log.WithFields(map[string]interface{}{
			"error": err,
			"host":  config.SurrealDBService,
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

func InitTwirpServer(config *runConf.ServiceConfig, log *logrus.Logger) (*protos.TwirpServer, *UserManagementServiceServerImpl) {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	println("Listening on", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	ssi := InitService(config, log)
	println("service initialized")
	twirpHandler := protos.NewUserManagementServiceServer(&ssi, nil)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if http.Serve(lis, twirpHandler); err != nil {
			log.Fatalf("failed to serve: %v", err)
			println("failed to serve: %v", err)
		}
	}()
	service_util.BeforeTwirpGracefulStop(&twirpHandler, ssi.Shutdown, log)
	return &twirpHandler, &ssi

}
