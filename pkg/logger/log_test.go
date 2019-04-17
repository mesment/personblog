package logger

import (
	"fmt"
	"testing"
	"time"
)

func TestFileLog(t *testing.T)  {

	config := make(map[string]string)
	config["logPath"] = "."
	config["logName"] = "testlog"
	config["dataChanSize"] = "5000"

	log,err:= NewFileLog(config)
	fmt.Printf("%v\n",log)
	if err != nil {
		t.Errorf("init faied")
	}
	log.Debug("file this  a test")

	log.Debug("file this  a test")
	log.Debug("file this  a test")
	log.Debug("file this  a test")
	log.Debug("file this  a test")
	log.Debug("file this  a test")
	log.Debug("file this  a test")
	log.Debug("file this  a test")
	log.Warn("file this  a warning")
	log.Error("file this  a error")
	log.Fatal("file this is a fatal log")

	time.Sleep(10*time.Second)
	log.Close()
}

/*
func TestConsoleLog(t *testing.T) {
	log,err := NewConsoleLog()
	if err != nil {
		t.Errorf("init faied")
	}
	log.Debug("console log")
	log.Warn("console warn log")
	log.Error("console err log")
	log.Fatal("console fatal log")
}
*/