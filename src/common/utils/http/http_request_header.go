package http

import (
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
		UserID:        req.Header.Get(CustomHeaderMap[UserID]),
		SessionID:     req.Header.Get(CustomHeaderMap[SessionID]),
		AuthToken:     req.Header.Get(CustomHeaderMap[TokenID]),
		TransactionID: req.Header.Get(CustomHeaderMap[TransactionID]),
		RequestID:     req.Header.Get(CustomHeaderMap[RequestID]),
		Timestamp:     req.Header.Get("ts"),
		UserAgent:     req.Header.Get("User-Agent"),
		Referrer:      req.Header.Get("Referer"),
		BucketsList:   req.Header.Get("bucket"),
		Debug:         getBoolHeaderField(req, CustomHeaderMap[DebugFlag]),
		ClientAppID:   req.Header.Get(CustomHeaderMap[AppID]),
	}
}

func getBoolHeaderField(req *http.Request, headerKey string) bool {
	value, err := strconv.ParseBool(req.Header.Get(headerKey))
	if err != nil {
		return false
	}
	return value
}
