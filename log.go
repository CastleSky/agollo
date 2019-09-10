package agollo

import (
	"fmt"
	)


type splog struct {

}
func (this *splog) Debug(v ...interface{}){
	fmt.Printf("[Debgu] %s \n", v)
}

func (this *splog) Info(v ...interface{}){
	fmt.Printf("[Info] %s \n", v)
}

func (this *splog) Warn(v ...interface{}){
	fmt.Printf("[Warn] %s \n", v)
}

func (this *splog) Error(v ...interface{}){
	fmt.Printf("[Error] %s \n", v)
}

var logger LoggerInterface

func init() {
	initLogger(initgLog())
}

func initLogger(ILogger LoggerInterface) {
	logger = ILogger
}

type LoggerInterface interface {
	// Debugf(format string, params ...interface{})

	// Infof(format string, params ...interface{})

	// Warnf(format string, params ...interface{})

	// Errorf(format string, params ...interface{})

	Debug(v ...interface{})

	Info(v ...interface{})

	Warn(v ...interface{})

	Error(v ...interface{})
}

func initgLog() LoggerInterface {
	logger := &splog{}
	return logger
}
