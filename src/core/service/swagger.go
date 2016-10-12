package service

import (
	utilhttp "github.com/jabong/florest-core/src/common/utils/http"
)

// swaggerAllowedHeaders headers allowed for swagger
var (
	swaggerAllowedHeaders = utilhttp.CustomHeaderMap[utilhttp.SessionID] + ", Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, " + utilhttp.CustomHeaderMap[utilhttp.TransactionID] + ", " + utilhttp.CustomHeaderMap[utilhttp.UserID]
)

// RegisterSwaggerHeader registes the user given header for swagger
func RegisterSwaggerHeader(newHeader string) {
	swaggerAllowedHeaders = newHeader
}
