package logging

import (
	"fmt"
)

type loggingModule interface {
	Log(msg string)
}

type loggingImpl struct {
	logLevel string
}

func (lm *loggingImpl) Log(msg string) {
	if lm.logLevel == "debug" {
		fmt.Printf("%s", msg)
	}
}

var loggingInstance *loggingImpl

func GetLoggingModule(logLevel string) loggingModule {
	if loggingInstance == nil {
		loggingInstance = &loggingImpl{
			logLevel: logLevel,
		}
	}
	return loggingInstance
}
