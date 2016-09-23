package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"

	"github.com/jabong/florest-core/src/common/logger/formatter"
	"github.com/jabong/florest-core/src/common/logger/impls"
	"github.com/jabong/florest-core/src/common/logger/writers/filewriter"
)

//Type of log implementations
const (

	//Syslog is a logger which dumps log to the OS syslog
	//	Syslog  string = "syslog"

	//Filelog is a logger which dumps log to a file. log rotation is not part
	//of file logging. This should not be part of application logger for now.
	//logrotation can be handled by external program like logrotate(8)
	//Sample logrotate config file is placed under logger config directory
	Filelog string = "file"
)

//loggerImpls stores various logger handles mapped by key
var loggerImpls map[string]LogInterface

//conf holds the various logger configs
var conf *Config

//Initialise initialises the logger
func Initialise(confFile string) error {
	conf = new(Config)

	file, err := ioutil.ReadFile(confFile)
	if err != nil {
		msg := fmt.Sprintf("Error loading Logger Config file %s \n %s", confFile, err)
		log.Println(msg)
		return err
	}
	err = json.Unmarshal(file, conf)
	if err != nil {
		msg := fmt.Sprintf("Incorrect Json in %s \n %s", confFile, err)
		log.Println(msg)
		return err
	}
	loggerImpls = make(map[string]LogInterface)
	return initLoggers()
}

//GetLoggerHandle returns a loggerHandle as specified by logType key
func GetLoggerHandle(logType string) (LogInterface, error) {
	loggerHandle, ok := loggerImpls[logType]
	if !ok {
		return nil, errors.New("Undefined log type requested " + logType)
	}
	return loggerHandle, nil
}

//GetDefaultLogTypeKey returns the default logger key
func GetDefaultLogTypeKey() string {
	if conf == nil {
		fmt.Println("Conf is null. Default log type key is empty")
		return ""
	}
	return conf.DefaultLogType
}

//getStackTrace gets the stack trace for a called function.
func getStackTrace() []string {
	var sf []string
	j := 0
	for i := Skip; ; i++ {
		_, filePath, lineNumber, ok := runtime.Caller(i)
		if !ok || j >= CallingDepth {
			break
		}
		sf = append(sf, fmt.Sprintf("%s(%d)", filePath, lineNumber))
		j++
	}
	return sf
}

//initFileLoggers initialises all file loggers
func initFileLoggers() error {
	var tmp LogInterface
	var format formatter.FormatInterface
	for i := 0; i < len(conf.FileLogger); i++ {
		c := conf.FileLogger[i]
		f := c.Path + conf.AppName + c.FileNamePrefix
		fh, err := filewriter.GetNewObj(f)
		if err != nil {
			fmt.Println("Error in initialising file loggers " + err.Error())
			return err
		}
		// check and set sync/asynch logger
		if &conf.AsyncLogger != nil && conf.AsyncLogger.Enabled {
			tmp, err = impls.NewAsynchLogger(&conf.AsyncLogger, fh)
			if err != nil {
				return err
			}
		} else {
			tmp = impls.GetSynchLogger(fh)
		}
		// check and set formatter
		if format, err = formatter.GetFormatter(conf.FileLogger[i].FormatType); err != nil {
			return err
		}
		// attach formatter to writer
		fh.SetFormatter(format)
		// entry for logger
		loggerImpls[c.Key] = tmp
	}
	return nil
}

//initLoggers initialises various type of logger like filelogger, etc
func initLoggers() error {
	return initFileLoggers()
}

//CanLog returns true if the argument logeLevel is smaller than config logLevel
func CanLog(logLevel int) bool {
	return conf != nil && conf.LogLevel >= logLevel
}
