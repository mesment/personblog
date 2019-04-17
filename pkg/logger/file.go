package logger

import (
	"fmt"
	"os"
	"strconv"
)


type FileLog struct {
	logLevel  int
	logPath string
	logName string
	file     *os.File
	warnFile *os.File
	dataChan chan *LogData
}


func NewFileLog(config map[string]string) (Log, error){
	logLevel, ok := config["logLevel"]
	if !ok {
		fmt.Printf("NewFileLog: config missing logLevel,using default Debug")
		logLevel = "Debug"
	}
	level :=LogLevelInt(logLevel)

	logPath, ok := config["logPath"]
	if !ok {
		err := fmt.Errorf("NewFileLog: config missing logPath")
		return nil, err
	}

	logName, ok := config["logName"]
	if !ok {
		err := fmt.Errorf("NewFileLog: config missing logName")
		return nil, err
	}

	dataChanSize, ok := config["dataChanSize"]
	if !ok {
		dataChanSize = "10000"  //如果没有配置，则使用默认值10000
	}
	//转换成数字格式
	chanSize,err := strconv.Atoi(dataChanSize)
	if err != nil {
		chanSize = 10000  //如果配置格式错误，则使用默认值10000
	}

	log:= &FileLog{
		logLevel:level,
		logPath: logPath,
		logName: logName,
		dataChan: make(chan *LogData, chanSize),
	}

	log.Init()

	return log,nil
}


func (f *FileLog) Init() {
	// 一般日志
	filename := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	// os.O_CREATE 创建文件 os.O_APPEND 追加写入 os.O_WRONLY 只写操作
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open log file %s failed, err: %v", filename, err))
	}

	f.file = file

	// 错误日志
	warnFileName := fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	// os.O_CREATE 创建文件 os.O_APPEND 追加写入 os.O_WRONLY 只写操作
	warnFile, err := os.OpenFile(warnFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open warning file %s failed, err: %v", warnFileName, err))
	}

	f.warnFile = warnFile

	//开启一个协程写日志
	go f.writeLogBackGround()
}

func (f *FileLog)writeLogBackGround()  {
	for logData := range f.dataChan {
		var file = f.file
		//切换日志文件
		if logData.IsWarn {
			file = f.warnFile
		}
		fmt.Printf("logdata:%v\n",logData)
		//fmt.Fprintf(os.Stdout,"%s %s[%s/%s:%d] %s\n",logData.Time,logData.LevelStr,
			//logData.FileName,logData.FuncName,logData.LineNo,logData.Msg)
		fmt.Fprintf(file,"%s %s[%s/%s:%d] %s\n",logData.Time,logData.LevelStr,
			logData.FileName,logData.FuncName,logData.LineNo,logData.Msg)


	}
}


func (f *FileLog) Debug(format string, args ...interface{}) {
	if f.logLevel > DebugLevel{
		return
	}
	logData := GetLogData(DebugLevel,format,args...)
	select {
	case f.dataChan <- logData:
	default:

	}
}

func (f *FileLog) Trace(format string, args ...interface{}) {
	if f.logLevel > TraceLevel{
		return
	}
	logData := GetLogData(TraceLevel,format,args...)
	select {
	case f.dataChan <- logData:
	default:

	}
}

func (f *FileLog) Info(format string, args ...interface{}) {
	if f.logLevel > InfoLevel{
		return
	}
	logData := GetLogData(InfoLevel,format,args...)
	select {
	case f.dataChan <- logData:
	default:

	}
}

func (f *FileLog) Warn(format string, args ...interface{}) {
	if f.logLevel > WarnLevel{
		return
	}
	logData := GetLogData(WarnLevel,format,args...)
	select {
	case f.dataChan <- logData:
	default:

	}
}

func (f *FileLog) Error(format string, args ...interface{}) {
	if f.logLevel > ErrorLevel{
		return
	}
	logData := GetLogData(ErrorLevel,format,args...)
	select {
	case f.dataChan <- logData:
	default:

	}
}

func (f *FileLog) Fatal(format string, args ...interface{}) {
	if f.logLevel > FatalLevel{
		return
	}
	logData := GetLogData(FatalLevel,format,args...)
	select {
	case f.dataChan <- logData:
	default:

	}
}

func (f *FileLog) Close() {
	f.file.Close()
	f.warnFile.Close()
}
