package logger

import (
	"github.com/jabong/florest-core/src/common/logger/impls"
)

//FileWriterConfig specifies the configs for file logger
type FileWriterConfig struct {
	//FileNamePrefix specifies the prefix for a log file name
	FileNamePrefix string `json:"FileNamePrefix"`

	//Key specifies a unique name for a particular logger implementation
	Key string `json:"Key"`

	//Path specifies the file logger path
	Path string `json:"Path"`

	//FormatType is the format in which logging is done.Currently It can be string or JSON.
	FormatType string `json:"FormatType"`
}

//Config specifies the logger config
type Config struct {

	//FileLogger stores all the configs for file loggers
	FileLogger []FileWriterConfig `json:"FileLogger"`

	//LogLevel specifies the minimum log level to write (1 - Info, 2 - Trace,
	//3 - Warning, 4 - Error)
	LogLevel int `json:"LogLevel"`

	//DefaultLogType specifies the default logger type. This should match
	//one of keys specified in logger configs
	DefaultLogType string `json:"DefaultLogType"`

	//AppName is the application name for which logging is done. AppName is
	//prefixed with each log file name
	AppName string `json:"AppName"`

	// AysncLogger config for async logging
	AsyncLogger impls.AsynchLoggerConfig
}
