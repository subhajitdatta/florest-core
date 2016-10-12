package http

import (
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
	chkNSetMap(CustomHeaderMap, headerMap, UserID, rc.UserID)
	chkNSetMap(CustomHeaderMap, headerMap, SessionID, rc.SessionID)
	chkNSetMap(CustomHeaderMap, headerMap, RequestID, rc.RequestID)
	chkNSetMap(CustomHeaderMap, headerMap, TransactionID, GetTransactionID())
	return headerMap
}

//GetTransactionID returns a new v4 UUID
func GetTransactionID() string {
	return uuid.NewV4().String()
}

// chkNSetMap check given key in input map and set in output map
func chkNSetMap(input map[CustomHeader]string, output map[string]string, key CustomHeader, value string) {
	if val, ok := input[key]; ok && val != "" {
		output[val] = value
	}
}
