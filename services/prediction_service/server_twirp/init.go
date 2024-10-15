package server

import (
	"fmt"
	"job-seek/pkg/database"
	"job-seek/pkg/protos"
	"job-seek/pkg/service_util"
	runConf "job-seek/services/prediction_service/config"
	"net/http"
	"sync"

	"gorm.io/gorm"

	logrus "github.com/sirupsen/logrus"

	"net"
)

type PredictionServiceServerImpl struct {
	log      *logrus.Logger
	config   *runConf.ServiceConfig
	mut      *sync.Mutex
	dbClient *gorm.DB
	// other implement here
}

func (s *PredictionServiceServerImpl) Startup() error {
	s.log.Info("Startup Produle")
	return nil
}

func (s *PredictionServiceServerImpl) Shutdown() error {
	s.log.Info("Shutdown")
	return nil
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

func InitTwirpServer(config *runConf.ServiceConfig, log *logrus.Logger) (*protos.TwirpServer, *PredictionServiceServerImpl) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	println("Listening on", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ssi := InitService(config, log)
	println("service initialized")

	twirpHandler := protos.NewPredictionServiceServer(&ssi, nil)

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
