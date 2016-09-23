package formatter

import (
	"github.com/jabong/florest-core/src/common/logger/message"
)

// FormatInterface interface methods for formatterss
type FormatInterface interface {
	GetFormattedLog(msg *message.LogMsg) interface{}
}
