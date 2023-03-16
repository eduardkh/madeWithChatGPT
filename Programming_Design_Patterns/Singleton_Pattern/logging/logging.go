package logging

import "fmt"

type loggingModule interface {
	Log(msg string)
}

type loggingImpl struct{}

func (lm *loggingImpl) Log(msg string) {
	fmt.Printf("[LOG] %s\n", msg)
}

var loggingInstance *loggingImpl

func GetLoggingModule() loggingModule {
	if loggingInstance == nil {
		loggingInstance = &loggingImpl{}
	}
	return loggingInstance
}
