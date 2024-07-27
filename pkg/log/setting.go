package log

import (
	"fmt"
	"io"
	"os"
	"time"

	logrus "github.com/sirupsen/logrus"
	writer "github.com/sirupsen/logrus/hooks/writer"
)

func InitLog(service_name string, level int) *logrus.Logger {
	log := logrus.New()
	log.SetOutput(io.Discard)
	log.SetFormatter(&logrus.JSONFormatter{})
	if level == 0 {
		log.SetLevel(logrus.PanicLevel)
	} else if level == 1 {
		log.SetLevel(logrus.ErrorLevel)

	} else if level == 2 {
		log.SetLevel(logrus.WarnLevel)

	} else if level == 3 {
		log.SetLevel(logrus.InfoLevel)

	} else if level == 4 {
		log.SetLevel(logrus.TraceLevel)
	}

	os.MkdirAll(fmt.Sprintf("log/%s", service_name), os.ModePerm)

	errFile, errFileErr := os.OpenFile(
		fmt.Sprintf("log/%s/error__%s.log", service_name, time.Now().Format("2006-01-02T15:04:05")),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666,
	)
	if errFileErr != nil {
		log.Error("Failed to log to file, using default stderr")
	}
	log.AddHook(&writer.Hook{ // Send logs with level higher than warning to stderr
		Writer: errFile,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
		},
	})

	logFile, logFileErr := os.OpenFile(
		fmt.Sprintf("log/%s/debug__%s.log", service_name, time.Now().Format("2006-01-02T15:04:05")),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666,
	)
	if logFileErr != nil {
		log.Error("Failed to log to file, using default stderr")
	}
	log.AddHook(&writer.Hook{ // Send info and debug logs to stdout
		Writer: logFile,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
			logrus.TraceLevel,
		},
	})

	log.WithFields(logrus.Fields{
		"service_name": service_name,
	}).Info("Logrus initialized")
	return log

}
