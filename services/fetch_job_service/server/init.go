package server

import (
	"fmt"
	"job-seek/pkg/protos"
	"job-seek/pkg/service_util"
	runConf "job-seek/services/fetch_job_service/config"
	"sync"

	logrus "github.com/sirupsen/logrus"

	"net"

	"google.golang.org/grpc"
)

type FetchJobServiceServer struct {
	protos.UnimplementedJobSearchServiceServer
	log    *logrus.Logger
	config *runConf.ServiceConfig

	mut      *sync.Mutex
	dsClient *grpc.ClientConn
}

func (s FetchJobServiceServer) Shutdown() error {
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
	fjss := FetchJobServiceServer{
		mut:      &sync.Mutex{},
		log:      log,
		config:   config,
		dsClient: nil,
	}
	protos.RegisterJobSearchServiceServer(grpcServer, fjss)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	service_util.BeforeGracefulStop(grpcServer, fjss.Shutdown, log)
	return grpcServer

}
