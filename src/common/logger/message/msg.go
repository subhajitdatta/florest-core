package message

import ()

/*
LogMsg struct is used for bundling message/requestcontext attributes
for dumping into log
*/
type LogMsg struct {
	Level         string   `json:"level"`
	Message       string   `json:"message"`
	TransactionID string   `json:"tId,omitempty"`
	RequestID     string   `json:"reqId,omitempty"`
	AppID         string   `json:"appId,omitempty"`
	SessionID     string   `json:"sessionId,omitempty"`
	UserID        string   `json:"userId,omitempty"`
	StackTraces   []string `json:"stackTraces,omitempty"`
	TimeStamp     string   `json:"timestamp"`
	URI           string   `json:"uri,omitempty"`
}
