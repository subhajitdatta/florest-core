package logger

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/logger/message"
	utilHttp "github.com/jabong/florest-core/src/common/utils/http"
	"time"
)

//DebugSpecific logs a debug to a specific log handle
func DebugSpecific(logType string, a ...interface{}) {
	if !CanLog(DebugLevel) {
		return
	}
	loggerHandle, err := GetLoggerHandle(logType)
	if err != nil {
		fmt.Println("Skipping Log Debug Message. " + err.Error())
		return
	}
	msg := Convert(a...)
	msg.Level = "debug"

	msg.TimeStamp = time.Now().Local().Format(time.RFC3339)
	loggerHandle.Debug(msg)
}

//Debug logs debug to a default log handle
func Debug(a ...interface{}) {
	DebugSpecific(GetDefaultLoggerType(), a...)
}

//InfoSpecific logs a info to a specific log handle
func InfoSpecific(logType string, a ...interface{}) {
	if !CanLog(InfoLevel) {
		return
	}
	loggerHandle, err := GetLoggerHandle(logType)
	if err != nil {
		fmt.Println("Skipping Log Info Message. " + err.Error())
		return
	}
	msg := Convert(a...)
	msg.Level = "info"
	msg.TimeStamp = time.Now().Local().Format(time.RFC3339)
	loggerHandle.Info(msg)
}

//Info logs info to a default log handle
func Info(a ...interface{}) {
	InfoSpecific(GetDefaultLoggerType(), a...)
}

//TraceSpecific logs a trace to a specific log handle
func TraceSpecific(logType string, a ...interface{}) {
	if !CanLog(TraceLevel) {
		return
	}
	loggerHandle, err := GetLoggerHandle(logType)
	if err != nil {
		fmt.Println("Skipping Log Trace Message. " + err.Error())
		return
	}
	msg := Convert(a...)
	msg.Level = "trace"
	msg.StackTraces = getStackTrace()
	msg.TimeStamp = time.Now().Local().Format(time.RFC3339)
	loggerHandle.Trace(msg)
}

//Trace logs Trace to default log handle
func Trace(a ...interface{}) {
	TraceSpecific(GetDefaultLoggerType(), a...)
}

//WarningSpecific logs a warning to a specific log handle
func WarningSpecific(logType string, a ...interface{}) {
	if !CanLog(WarningLevel) {
		return
	}
	loggerHandle, err := GetLoggerHandle(logType)
	if err != nil {
		fmt.Println("Skipping Log Warning Message : " + err.Error())
		return
	}
	msg := Convert(a...)
	msg.Level = "warning"
	msg.StackTraces = getStackTrace()
	msg.TimeStamp = time.Now().Local().Format(time.RFC3339)
	loggerHandle.Warning(msg)
}

//Warning logs warning to default log handle
func Warning(a ...interface{}) {
	WarningSpecific(GetDefaultLoggerType(), a...)
}

//ErrorSpecific logs an error to a specific logger handle
func ErrorSpecific(logType string, a ...interface{}) {
	if !CanLog(ErrLevel) {
		return
	}
	loggerHandle, err := GetLoggerHandle(logType)
	if err != nil {
		fmt.Println("Skipping Log Error Message. " + err.Error())
		return
	}
	msg := Convert(a...)
	msg.Level = "error"
	msg.StackTraces = getStackTrace()
	msg.TimeStamp = time.Now().Local().Format(time.RFC3339)
	loggerHandle.Error(msg)
}

//Error logs error to default log handle
func Error(a ...interface{}) {
	ErrorSpecific(GetDefaultLoggerType(), a...)
}

//ProfileSpecific logs a profile specifying memory and time taken by an execution
//point
func ProfileSpecific(logType string, a ...interface{}) {
	loggerHandle, err := GetLoggerHandle(logType)
	if err != nil {
		fmt.Println("Skipping Log Profile Message : " + err.Error())
		return
	}
	msg := Convert(a...)
	msg.Level = "profile"
	msg.TimeStamp = time.Now().Local().Format(time.RFC3339)
	loggerHandle.Profile(msg)
}

//Profile logs a profile specifying memory and time taken by an execution
//point to a default log handle
func Profile(a ...interface{}) {
	ProfileSpecific(GetDefaultLoggerType(), a...)
}

//GetDefaultLoggerType gets the key of a default logger type
func GetDefaultLoggerType() string {
	return GetDefaultLogTypeKey()
}

// Convert converts application log to LogMsg format suitable for Logger
func Convert(a ...interface{}) message.LogMsg {
	paramLength := len(a)
	if paramLength == 0 {
		return message.LogMsg{
			Message: "Empty log param",
		}
	}
	if paramLength == 1 {
		//Only Log Message string is passed
		return message.LogMsg{
			Message: fmt.Sprintf("%s", a[0]),
		}
	}

	//First param is message string; Second param is request context
	vMsg, msgOk := a[0].(string)
	vRc, rcOk := a[1].(utilHttp.RequestContext)

	if !msgOk || !rcOk {

		return message.LogMsg{
			Message: fmt.Sprintf("Erorr in parsing logging params for %v", a),
		}
	}
	return message.LogMsg{
		Message:       vMsg,
		TransactionID: vRc.TransactionID,
		SessionID:     vRc.SessionID,
		RequestID:     vRc.RequestID,
		AppID:         vRc.AppName,
		UserID:        vRc.UserID,
		URI:           vRc.URI,
	}
}
