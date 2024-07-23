package service_util

import (
	"context"
	"job-seek/pkg/config"
	"time"

	logrus "github.com/sirupsen/logrus"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func CreateGrpcServerOption(conf *config.ServerConfig, log *logrus.Logger) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionIdle:     time.Duration(conf.MaxConnectionIdle) * time.Second,
				MaxConnectionAge:      time.Duration(conf.MaxConnectionAge) * time.Second,
				MaxConnectionAgeGrace: time.Duration(conf.MaxConnectionAgeGrace) * time.Second,
				Time:                  time.Duration(conf.Time) * time.Second,
				Timeout:               time.Duration(conf.Timeout) * time.Second,
			},
		),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             time.Duration(conf.MinTime) * time.Second,
				PermitWithoutStream: conf.PermitWithoutStream,
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
