package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

//日志数据
type LogData struct {
	Msg		 string
	Time	 string
	LevelStr string
	FileName string
	FuncName string
	LineNo   int
	IsWarn   bool //是否写入错误日志文件
}



//日志级别字符串
func LogLevelString(level int) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case TraceLevel:
		return "TRACE"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "DEBUG"
	}
	return "DEBUG"
}

func LogLevelInt(levelStr string) int {
	switch levelStr {
	case "DEBUG":case "Debug": case "debug":
		return DebugLevel
	case "TRACE":case "Trace": case "trace":
		return TraceLevel
	case "INFO" : case "Info": case "info":
		return InfoLevel
	case "WARN": case "Warn": case "warn":
		return WarnLevel
	case "ERROR":case "Error":case "error":
		return ErrorLevel
	case "FATAL":case "Fatal":case "fatal":
		return FatalLevel
	default:
		return DebugLevel
	}
	return DebugLevel
}

//GetLineInfo：获取行号
func GetLineInfo() (fileName, funcName string, lineNo int) {
	//pc 计数器， file 文件名， line 行号， ok 是否
	// runtime.Caller(4)这里的4是一个层级关系，可以尝试使用0 1 2 3来看看
	// 4 在其他项目中使用的时候，在test中，使用3
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name() // 获取当前的方法
		lineNo = line
	}
	return
}

//获取时间
func GetTime() string {
	now := time.Now().Format("2006-01-02 15:04:05.999")
	return now
}


func writeLog(file *os.File, logLevel int, format string, args ...interface{})  {
	logData := GetLogData(logLevel,format,args...)
	fmt.Fprintf(file,"%s %s [%s/%s:%d] %s\n",logData.Time, logData.LevelStr,
		logData.FileName, logData.FuncName, logData.LineNo,logData.Msg)
}


func GetLogData(logLevel int, format string, args ...interface{}) *LogData {
	//时间
	now := GetTime()
	//日志级别
	levelStr :=LogLevelString(logLevel)
	//文件名， 函数名， 行号
	fileName, funcName, lineNo := GetLineInfo()
	//返回的是带路径的，只需要文件名以及函数名
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format, args...)
	isWarn := logLevel >= WarnLevel && logLevel <= FatalLevel
	return &LogData{
		Msg:  msg,
		Time:  now,
		LevelStr: levelStr,
		FileName: fileName,
		FuncName: funcName,
		LineNo:   lineNo,
		IsWarn:   isWarn,
	}
}


