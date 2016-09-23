package utils

import (
	"net/http"
)

func CreateTestRequest(httpMethod string, urlString string) *http.Request {
	request, _ := http.NewRequest(httpMethod, urlString, nil)
	request.RequestURI = urlString
	return request
}
