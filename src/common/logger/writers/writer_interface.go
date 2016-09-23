package writers

import (
	"github.com/jabong/florest-core/src/common/logger/formatter"
	"github.com/jabong/florest-core/src/common/logger/message"
)

type LogWriter interface {
	Write(msg *message.LogMsg)
	SetFormatter(formatter.FormatInterface)
}
