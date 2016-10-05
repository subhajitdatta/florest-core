package sqldb

import (
	"fmt"
)

type SDBError struct {
	ErrCode          string
	DeveloperMessage string
}

func (e *SDBError) Error() string {
	return fmt.Sprintf("{ErrCode:%s,DeveloperMessage:%s}", e.ErrCode, e.DeveloperMessage)
}

const (
	ErrNoDriver       = "Driver not found"
	ErrInitialization = "Initialization failed"
	ErrQueryFailure   = "Failure in Query() method"
	ErrExecuteFailure = "Failure in Execute() method"
	ErrPingFailure    = "Failure in Ping() method"
	ErrGetTxnFailure  = "Failure in GetTxnObj() method"
	ErrCloseFailure   = "Failure in Close() method"
	ErrKeyPresent     = "Key is already present"
	ErrKeyNotPresent  = "Key is not present"
	ErrWrongType      = "Incorrect type sent"
)

// getErrObj returns error object with given details
func getErrObj(errCode string, developerMessage string) (ret *SDBError) {
	ret = new(SDBError)
	ret.ErrCode = errCode
	ret.DeveloperMessage = developerMessage
	return ret
}
