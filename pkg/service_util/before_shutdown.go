package service_util

import (
	"os"
	"os/signal"
	"syscall"

	logrus "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type ShudownCallback = func() error

func BeforeGracefulStop(gs *grpc.Server, shudownCB ShudownCallback, log *logrus.Logger) {
	// log := fjss.log
	log.Info("BeforeGracefulStop")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGABRT)
	aa := <-c
	log.Info("OS.signal", aa)
	log.Info(gs.GetServiceInfo())
	// fjss.Shutdown()
	shudownCB()
	gs.GracefulStop()
	log.Info("os GracefulStop")
	os.Exit(0)
}
