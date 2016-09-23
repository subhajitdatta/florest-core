package http

import (
	"fmt"
)

/*
The Request Execution Context is used for tracking a particular request processing
*/
type RequestContext struct {
	AppName       string
	UserID        string
	SessionID     string
	RequestID     string
	TransactionID string
	URI           string
	ClientAppID   string
	TokenID       string
}

//Implements the Stringer interface
func (t RequestContext) String() string {
	format := "[AppName : %s, UserID : %s, SessionID : %s, RequestID : %s, TransactionID : %s, TokenId : %s, URI : %s, ClientAppId : %s]"
	return fmt.Sprintf(format,
		t.AppName,
		t.UserID,
		t.SessionID,
		t.RequestID,
		t.TransactionID,
		t.TokenID,
		t.URI,
		t.ClientAppID,
	)
}
