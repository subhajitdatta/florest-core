package formatter

import (
	"encoding/json"
	"fmt"
	"github.com/jabong/florest-core/src/common/logger/message"
)

type jsonFormat struct {
}

//GetFormattedLog returns formatted log
func (jf *jsonFormat) GetFormattedLog(msg *message.LogMsg) interface{} {
	jMsg, err := json.Marshal(msg)
	if err != nil {
		return fmt.Sprintf("\nError In converting to json %+v\n", msg)
	}
	return string(jMsg)
}
