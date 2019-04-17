package logger

import (
	"os"
)

type ConsoleLog struct {
}


func NewConsoleLog() (Log,error) {
	return &ConsoleLog{},nil
}


func (f *ConsoleLog)Init()  {}

func (f *ConsoleLog) Debug(format string, args ...interface{}) {
	writeLog(os.Stdout,DebugLevel,format,args...)
}

func (f *ConsoleLog) Trace(format string, args ...interface{}) {
	writeLog(os.Stdout,TraceLevel,format,args...)
}

func (f *ConsoleLog) Info(format string, args ...interface{}) {
	writeLog(os.Stdout,InfoLevel,format,args...)
}

func (f *ConsoleLog) Warn(format string, args ...interface{}) {
	writeLog(os.Stdout,WarnLevel,format,args...)
}

func (f *ConsoleLog) Error(format string, args ...interface{}) {
	writeLog(os.Stdout,ErrorLevel,format,args...)
}

func (f *ConsoleLog) Fatal(format string, args ...interface{}) {
	writeLog(os.Stdout,FatalLevel,format,args...)
}

func (f *ConsoleLog) Close() {
}
