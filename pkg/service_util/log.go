package service_util

import (
	"github.com/sirupsen/logrus"
)

type Loggable[T any] struct {
	log *logrus.Logger
}

func createLogField(method string, fields map[string]interface{}) logrus.Fields {
	fieldset := logrus.Fields{
		"method": method,
	}
	for k, v := range fields {
		fieldset[k] = v
	}
	return fieldset
}

func (s Loggable[T]) LogTrace(method string, message string, fields map[string]interface{}) {
	s.log.WithFields(createLogField(method, fields)).Trace(message)
}

func (s Loggable[T]) LogDebug(method string, message string, fields map[string]interface{}) {
	s.log.WithFields(createLogField(method, fields)).Debug(message)
}

func (s Loggable[T]) LogInfo(method string, message string, fields map[string]interface{}) {
	s.log.WithFields(createLogField(method, fields)).Info(message)
}

func (s Loggable[T]) LogWarn(method string, message string, fields map[string]interface{}) {
	s.log.WithFields(createLogField(method, fields)).Warn(message)
}

func (s Loggable[T]) LogError(method string, message string, fields map[string]interface{}) {
	s.log.WithFields(createLogField(method, fields)).Error(message)
}
