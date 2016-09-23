package formatter

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/logger/message"
)

type stringFormat struct {
}

//formatString is format of the log in string formattype configuration
var formatString = "[level : %s, message : %s, tId : %s, reqId : %s, appId : %s, sessionId : %s, userId : %s, stackTraces : %s, timestamp : %s, uri : %s]"

//GetFormattedLog returns formatted log as a string interface
func (sf *stringFormat) GetFormattedLog(msg *message.LogMsg) interface{} {
	return fmt.Sprintf(formatString, msg.Level,
		msg.Message,
		msg.TransactionID,
		msg.RequestID,
		msg.AppID,
		msg.SessionID,
		msg.UserID,
		msg.StackTraces,
		msg.TimeStamp,
		msg.URI)
}
