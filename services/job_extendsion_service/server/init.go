package server

import (
	"fmt"
	"job-seek/pkg/protos"
	"job-seek/pkg/service_util"
	runConf "job-seek/services/job_extendsion_service/config"
	"sync"

	logrus "github.com/sirupsen/logrus"

	"net"

	"google.golang.org/grpc"
)

type JobExtendsionServiceServerImpl struct {
	protos.UnimplementedJobExtendsionServiceServer
	log    *logrus.Logger
	config *runConf.ServiceConfig

	mut *sync.Mutex

	// other implement here
}

func (s JobExtendsionServiceServerImpl) Startup() error {
	s.log.Info("Startup Produle")
	return nil
}

func (s JobExtendsionServiceServerImpl) Shutdown() error {
	s.log.Info("Shutdown")
	return nil
}

func InitGrpcServer(config *runConf.ServiceConfig, log *logrus.Logger) *grpc.Server {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opt := service_util.CreateGrpcServerOption(&config.Server, log)
	grpcServer := grpc.NewServer(opt...)
	ssi := JobExtendsionServiceServerImpl{
		mut:    &sync.Mutex{},
		log:    log,
		config: config,
	}
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
