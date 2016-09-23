package http

import (
	"github.com/jabong/florest-core/src/common/constants"
	"net/http"
	"strconv"
)

type RequestHeader struct {
	ContentType   string
	Accept        string
	UserID        string
	SessionID     string
	AuthToken     string
	TransactionID string
	RequestID     string
	Timestamp     string
	UserAgent     string
	Referrer      string
	BucketsList   string
	Debug         bool
	ClientAppID   string
}

func GetReqHeader(req *http.Request) RequestHeader {
	return RequestHeader{
		ContentType:   req.Header.Get("Content-Type"),
		Accept:        req.Header.Get("Accept"),
		UserID:        req.Header.Get(constants.UserID),
		SessionID:     req.Header.Get(constants.SessionID),
		AuthToken:     req.Header.Get(constants.AuthToken),
		TransactionID: req.Header.Get(constants.TransactionID),
		RequestID:     req.Header.Get(constants.RequestID),
		Timestamp:     req.Header.Get("ts"),
		UserAgent:     req.Header.Get("User-Agent"),
		Referrer:      req.Header.Get("Referer"),
		BucketsList:   req.Header.Get("bucket"),
		Debug:         getBoolHeaderField(req, constants.Debug),
		ClientAppID:   req.Header.Get(constants.APPID),
	}
}

func getBoolHeaderField(req *http.Request, headerKey string) bool {
	value, err := strconv.ParseBool(req.Header.Get(headerKey))
	if err != nil {
		return false
	}
	return value
}
