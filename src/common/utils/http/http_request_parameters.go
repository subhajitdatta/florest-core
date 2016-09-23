package http

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// GetStringParamFields returns a string type fieldName from http req
func GetStringParamFields(req *http.Request, fieldName string) string {
	return strings.TrimSpace(req.FormValue(fieldName))
}

// GetIntParamFields returns a int type fieldName from http req
func GetIntParamFields(req *http.Request, fieldName string) (int, error) {
	s := GetStringParamFields(req, fieldName)
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// GetBoolParamFields returns a boolean type fieldName from http req
func GetBoolParamFields(req *http.Request, fieldName string) (bool, error) {
	value, err := strconv.ParseBool(GetStringParamFields(req, fieldName))
	if err != nil {
		return false, err
	}
	return value, nil
}

// GetIntParamNUpdateMetaData returns an int value of field name req and updates the response metadata node. If the field
// is not present in the req then it returns the defaultVal
func GetIntParamNUpdateMetaData(req *http.Request, fieldName string, defaultVal int, md *ResponseMetaData) (int, error) {
	s := GetStringParamFields(req, fieldName)
	md.URLParams[fieldName] = s
	if s == "" {
		return defaultVal, nil
	}
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// getBodyParam returns body param for the request
func getBodyParam(req *http.Request) (string, error) {
	bytAry, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	return string(bytAry), nil
}
