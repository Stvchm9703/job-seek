package server

import (
	"context"
	"fmt"
	"job-seek/pkg/protos"
	runConf "job-seek/services/fetch_job_service/config"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	logrus "github.com/sirupsen/logrus"

	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
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
	opt := CreateGrpcServerOption(&config.Server, log)
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
	beforeGracefulStop(grpcServer, &fjss)
	return grpcServer

}

func CreateGrpcServerOption(config *runConf.ServerConfig, log *logrus.Logger) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionIdle:     time.Duration(config.MaxConnectionIdle) * time.Second,
				MaxConnectionAge:      time.Duration(config.MaxConnectionAge) * time.Second,
				MaxConnectionAgeGrace: time.Duration(config.MaxConnectionAgeGrace) * time.Second,
				Time:                  time.Duration(config.Time) * time.Second,
				Timeout:               time.Duration(config.Timeout) * time.Second,
			},
		),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             time.Duration(config.MinTime) * time.Second,
				PermitWithoutStream: config.PermitWithoutStream,
			},
		),
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(
				messageLoggerInterceptors(log),
				logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
			),
		),

		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(
				messageLoggerInterceptors(log),
				logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
			),
		),
	}
}

func beforeGracefulStop(gs *grpc.Server, fjss *FetchJobServiceServer) {
	log := fjss.log
	log.Info("BeforeGracefulStop")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGABRT)
	aa := <-c
	log.Info("OS.signal", aa)
	log.Info(gs.GetServiceInfo())
	fjss.Shutdown()
	gs.GracefulStop()
	log.Info("os GracefulStop")
	os.Exit(0)
}

func messageLoggerInterceptors(log *logrus.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make(map[string]any, len(fields)/2)
		i := logging.Fields(fields).Iterator()
		for i.Next() {
			k, v := i.At()
			f[k] = v
		}
		l := log.WithFields(f)

		switch lvl {
		case logging.LevelDebug:
			l.Debug(msg)
		case logging.LevelInfo:
			l.Info(msg)
		case logging.LevelWarn:
			l.Warn(msg)
		case logging.LevelError:
			l.Error(msg)
		default:
			l.Trace(msg)
		}
	})
}
