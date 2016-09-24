// Package logger provides the logger used in florest-core.
//
// Characteristics
//
//
// 1. Supports 5 log levels - Debug, Info, Trace, Warning & Error.
//
// 2. Provides a stack trace for all the error logs
//
// 3. Supports async logging
//
// 4. Log Messages by default are in logstash compatible structured json format
//
// Configuration
//
// Logger configuration can be specified via json configuration file. For example refer
// "florest-core/config/logger/logger.json"
//
// The log level maps to an integer as specified in "florest-core/src/common/logger/logger_constants.go"
// To set the log level as Info for the application. Then "LogLevel" should be set to value 2. The "LogLevel"
// in json configuration file can also be specified when building the florest app by executing the below
// command:-
//		make deploy LOGLEVEL=2
//
// By default logger writes to file which can be configured from the json configuration file using sync
// logging. However async logger is more performant than sync logger but there is a chance that logs
// messages might be lost.
package logger
