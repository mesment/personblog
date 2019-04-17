package logger

import (
	"fmt"
	"github.com/mesment/personblog/pkg/setting"
)

type Log interface {
	Init()
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

var (
	log  Log
)

func Setup()  {

	config := make(map[string]string)

	config["logLevel"] = setting.LogCfg.Level
	config["logPath"] = setting.LogCfg.LogPath
	config["logName"] = setting.LogCfg.LogName
	config["logType"] = setting.LogCfg.LogType

	 _, err := InitLog(config)
	 if err != nil {

	 	fmt.Printf("初始化logger配置信息失败:%v\n",err)
		 return
	 }
	
}


func InitLog( config map[string]string) (Log, error) {

	logType, ok := config["logType"]
	if !ok {
		fmt.Printf("LogType is not found in config. using default console")
		config["logType"] = "console"
	}
	var err error
	switch logType {
	case "file":
		log,err = NewFileLog(config)
		return log, nil
	case "console":
		log,err = NewConsoleLog()
		return log, nil
	default:
		err = fmt.Errorf("InitLog failed: type error,%s unspport.",logType)
}
	return log, err
}

func  Debug(format string, args ...interface{}) {
	log.Debug(format,args...)
}

func  Trace(format string, args ...interface{}) {
	log.Trace(format,args...)
}

func Info(format string, args ...interface{}) {
	log.Info(format, args...)
}

func  Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}

func Fatal(format string, args ...interface{}) {
	log.Fatal(format,args...)
}

func  Close() {
	log.Close()
}
