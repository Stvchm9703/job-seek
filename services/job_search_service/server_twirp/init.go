package server

import (
	"fmt"
	"job-seek/pkg/database"
	"job-seek/pkg/protos"
	"job-seek/pkg/service_util"
	runConf "job-seek/services/job_search_service/config"
	"net"
	"net/http"
	"sync"

	logrus "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JobSearchServiceServerImpl struct {
	log    *logrus.Logger
	config *runConf.ServiceConfig

	mut *sync.Mutex
	// dbClient *surrealdb.DB
	dbClient *gorm.DB
	// other implement here
}

func (s *JobSearchServiceServerImpl) Startup() error {
	s.log.Info("Startup Produle")

	return nil
}

func (s *JobSearchServiceServerImpl) Shutdown() error {
	s.log.Info("Shutdown")

	s.log.Info("Closing database connection")
	if s.dbClient != nil {
		// s.dbClient.Close()
		dbInstance, _ := s.dbClient.DB()
		dbInstance.Close()
		// s.dbClient = nil
	} else {
		s.log.Warn("Database connection is nil pointer, check the memmory leak")
	}
	s.log.Info("Database connection closed")

	return nil
}

func InitService(config *runConf.ServiceConfig, log *logrus.Logger) JobSearchServiceServerImpl {

	dbClient, err := database.InitConnection(&config.DBService, "development")
	log.Info("Connecting to database")
	if err != nil {
		println("Failed to connect to database in InitService")
		println(err.Error())
		log.WithFields(map[string]interface{}{
			"error": err,
			"host":  config.DBService,
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

func InitTwirpServer(config *runConf.ServiceConfig, log *logrus.Logger) (*protos.TwirpServer, *JobSearchServiceServerImpl) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	println("Listening on", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// opt := service_util.CreateGrpcServerOption(&config.Server, log)
	// grpcServer := grpc.NewServer(opt...)
	ssi := InitService(config, log)
	println("service initialized")
	// protos.RegisterJobSearchServiceServer(grpcServer, ssi)
	twirpHandler := protos.NewJobSearchServiceServer(&ssi, nil)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		// http.ListenAndServe(lis, twirpHandler)
		if http.Serve(lis, twirpHandler); err != nil {
			log.Fatalf("failed to serve: %v", err)
			println("failed to serve: %v", err)
		}
	}()
	service_util.BeforeTwirpGracefulStop(&twirpHandler, ssi.Shutdown, log)
	return &twirpHandler, &ssi

}
