package logger

import (
	//"reflect"
	"testing"
	"time"
)

//startTest does initialisation that is required before each test
func startTest() {
	//TO DO:- Have separate config for the test
	Initialise("../../../config/logger/logger_dev.json")
}

//endTest cleans up all resources that are allocated or initialised
//during startTest
func endTest() {
	conf = nil
}

//TO DO - This test will be fixed once we fix on a test framework
//func Test_GetLoggerHandle(t *testing.T) {
//	_, ok := loggerImpls["123459"]
//	if ok {
//		t.Error("GetLoggerHandle Test Failed - Returned wrong logger handle")
//	}
//
//	startTest()
//
//	f := new(FileLoggerImpl)
//	f.filename = getLogFName(conf.FileLogger[0])
//
//	r, err := GetLoggerHandle("123459")
//	if err == nil {
//		t.Error("GetLoggerHandle Test Failed - Returned logger handle for invalid logtype key", r)
//	}
//
//	r, err = GetLoggerHandle(conf.DefaultLogType)
//	if reflect.DeepEqual(f, r) == false {
//		t.Errorf("GetLoggerHandle Test Failed - Returned wrong logger handle %v", r)
//	}
//
//	endTest()
//}

func Test_GetDefaultLogTypeKey(t *testing.T) {
	if GetDefaultLogTypeKey() != "" {
		t.Error("Returned non-null log type key")
	}

	startTest()

	lType := GetDefaultLogTypeKey()
	if conf.DefaultLogType != lType {
		t.Errorf("Returned invalid log type %s", lType)
	}

	endTest()
}

//TO DO - This test will be fixed once we fix on a test framework
//func Test_newFileLogger(t *testing.T) {
//	startTest()
//
//	f := new(FileLoggerImpl)
//	f.filename = getLogFName(conf.FileLogger[0])
//	f.maxAge = conf.FileLogger[0].MaxAge
//	f.maxBackups = conf.FileLogger[0].MaxBackUp
//	f.maxSize = conf.FileLogger[0].MaxSize
//
//	r := newFileLogger(f.filename, conf.FileLogger[0])
//	if reflect.DeepEqual(f, r) == false {
//		t.Errorf("NewFileLogger Test failed - Returned incorrent file impl %v", *r)
//	}
//
//	endTest()
//}

func getLogFName(c FileWriterConfig) string {
	t := time.Now().Local()
	tf := t.Format("2006-01-02")
	f := c.Path + c.FileNamePrefix + tf + ".log"
	return f
}
