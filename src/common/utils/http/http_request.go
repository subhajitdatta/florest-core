package http

import (
	"errors"
	"net/http"
	"strings"
)

type Method string

const (
	GET    Method = "GET"
	PUT    Method = "PUT"
	POST   Method = "POST"
	DELETE Method = "DELETE"
	PATCH  Method = "PATCH"
)

// Request represents all the request related data
type Request struct {
	HTTPVerb        Method
	URI             string
	OriginalRequest *http.Request
	Headers         RequestHeader
	PathParameters  *map[string]string
}

func getMethod(method string) (Method, error) {
	switch strings.ToUpper(method) {
	case "GET":
		return GET, nil
	case "PUT":
		return PUT, nil
	case "POST":
		return POST, nil
	case "DELETE":
		return DELETE, nil
	case "PATCH":
		return PATCH, nil
	}
	return "", errors.New("Incorrect HTTP Method")
}

func GetRequest(r *http.Request) (Request, error) {
	httpVerb, verr := getMethod(r.Method)
	if verr != nil {
		return Request{}, verr
	}

	return Request{
		HTTPVerb:        httpVerb,
		URI:             r.URL.String(),
		OriginalRequest: r,
		Headers:         GetReqHeader(r)}, nil
}

// GetBodyParameter returns the Body Parameters from the OriginalRequest as string
func (req *Request) GetBodyParameter() (string, error) {
	return getBodyParam(req.OriginalRequest)
}

// GetPathParameter returns the value corresponding to the key in Path parameter map as string
func (req *Request) GetPathParameter(key string) string {
	pathParams := req.PathParameters
	if pathParams != nil {
		if val, ok := (*pathParams)[key]; ok {
			return val
		}
	}
	return ""
}

// GetHeaderParameter returns the value for given header key
func (req *Request) GetHeaderParameter(key string) string {
	return req.OriginalRequest.Header.Get(key)
}
