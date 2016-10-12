package http

import (
	"strings"
)

// CustomHeader customizable headers
type CustomHeader string

const (
	UserID        CustomHeader = "UserId"
	SessionID     CustomHeader = "SessionId"
	RequestID     CustomHeader = "RequestId"
	TransactionID CustomHeader = "TransactionID"
	TokenID       CustomHeader = "TokenID"
	AppID         CustomHeader = "AppID"
	DebugFlag     CustomHeader = "DebugFlag"
)

// CustomHeaderMap map to store custom headers
var CustomHeaderMap = map[CustomHeader]string{
	UserID:        "USER_ID",
	SessionID:     "SESSION_ID",
	RequestID:     "REQUEST_ID",
	TransactionID: "TRANSACTION_ID",
	TokenID:       "TOKEN_ID",
	AppID:         "APP_ID",
	DebugFlag:     "DEBUG",
}

// RegisterCustomHeader registers the map with user defined values
func RegisterCustomHeader(newMap map[CustomHeader]string) {
	for k, v := range newMap {
		// add if value is not empty
		if val := strings.TrimSpace(v); val != "" {
			CustomHeaderMap[k] = val
		}
	}
}
