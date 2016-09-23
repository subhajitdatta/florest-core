package http

import (
	"github.com/jabong/florest-core/src/common/constants"
	"github.com/twinj/uuid"
)

//GetHTTPHeaders returns a map of required headers
//GetHTTPHeaders reads the header values from input request context. A new transaction id is
//created for each call to this method
func GetHTTPHeaders(rc *RequestContext) map[string]string {
	if rc == nil {
		return nil
	}

	headerMap := make(map[string]string, 4)
	headerMap[constants.UserID] = rc.UserID
	headerMap[constants.SessionID] = rc.SessionID
	headerMap[constants.RequestID] = rc.RequestID
	headerMap[constants.TransactionID] = GetTransactionID()
	return headerMap
}

//GetTransactionID returns a new v4 UUID
func GetTransactionID() string {
	return uuid.NewV4().String()
}
