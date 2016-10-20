package ratelimiter

import (
	"fmt"
)

type Error struct {
	ErrCode          string
	DeveloperMessage string
}

func (e *Error) Error() string {
	return fmt.Sprintf("{ErrCode:%s,DeveloperMessage:%s}", e.ErrCode, e.DeveloperMessage)
}

const (
	ErrInitialization = "Initialization failed"
	ErrFatal          = "Rate Limit fatal error"
)

// getErrObj returns error object with given details
func getErrObj(errCode string, developerMessage string) (ret *Error) {
	ret = new(Error)
	ret.ErrCode = errCode
	ret.DeveloperMessage = developerMessage
	return ret
}
