package service

import (
	"github.com/jabong/florest-core/src/common/constants"
)

const (
	SwaggerAllowedHeaders = constants.SessionID + ", Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, " + constants.TransactionID + ", " + constants.UserID
)
