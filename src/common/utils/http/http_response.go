package http

import (
	"github.com/jabong/florest-core/src/common/constants"
)

type Debug struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ResponseMetaData struct {
	URLParams   map[string]interface{} `json:"urlParams"`
	APIMetaData map[string]interface{} `json:"apiMetaData"`
}

// NewResponseMetaData creates and returns an instance of ResponseMetaData
func NewResponseMetaData() *ResponseMetaData {
	r := new(ResponseMetaData)
	r.URLParams = make(map[string]interface{})
	r.APIMetaData = make(map[string]interface{})
	return r
}

type Response struct {
	Status    constants.APPHttpStatus `json:"status"`
	Data      interface{}             `json:"data"`
	DebugData []Debug                 `json:"debugData,omitempty"`
	MetaData  *ResponseMetaData       `json:"_metaData,omitempty"`
}

// APIResponse represents a complete response containing HTTPStatus, Headers and Body
type APIResponse struct {
	HTTPStatus constants.HTTPCode
	Headers    map[string]string
	Body       []byte
}

// NewAPIResponse returns a creates and returns an instance of APIResponse
func NewAPIResponse() APIResponse {
	a := APIResponse{}
	a.Headers = make(map[string]string)
	return a
}
