package http

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jabong/florest-core/src/common/constants"
)

//Get makes an http get request with default header parameters
func Get(url string, headers map[string]string,
	timeOut time.Duration) (ret *APIResponse, err error) {
	var resp *http.Response // response
	defer recoverFromPanic("Get()", resp, &err)
	req, err := getReqWithoutBody("GET", url)
	if err != nil {
		return nil, err
	}
	return httpExecuter(req, resp, headers, timeOut)
}

//Post makes an http post request with given parameters
func Post(url string, headers map[string]string, body string,
	timeOut time.Duration) (ret *APIResponse, err error) {
	var resp *http.Response // response
	defer recoverFromPanic("Post()", resp, &err)
	req, err := getReqWithBody("POST", url, body)
	if err != nil {
		return nil, err
	}
	return httpExecuter(req, resp, headers, timeOut)
}

//Put makes an http put request with given parameters
func Put(url string, headers map[string]string, body string,
	timeOut time.Duration) (ret *APIResponse, err error) {
	var resp *http.Response // response
	defer recoverFromPanic("Put()", resp, &err)
	req, err := getReqWithBody("PUT", url, body)
	if err != nil {
		return nil, err
	}
	return httpExecuter(req, resp, headers, timeOut)
}

//Delete makes an http delete request with given parameters
func Delete(url string, headers map[string]string, body string,
	timeOut time.Duration) (ret *APIResponse, err error) {
	var resp *http.Response // response
	defer recoverFromPanic("Delete()", resp, &err)
	req, err := getReqWithBody("DELETE", url, body)
	if err != nil {
		return nil, err
	}
	return httpExecuter(req, resp, headers, timeOut)
}

//Patch makes an http patch request with given parameters
func Patch(url string, headers map[string]string, body string,
	timeOut time.Duration) (ret *APIResponse, err error) {
	var resp *http.Response // response
	defer recoverFromPanic("Patch()", resp, &err)
	req, err := getReqWithBody("PATCH", url, body)
	if err != nil {
		return nil, err
	}
	return httpExecuter(req, resp, headers, timeOut)
}

// recoverFromPanic closes any open http response, recovers with panic details
func recoverFromPanic(name string, resp *http.Response, err *error) {
	if resp != nil {
		resp.Body.Close()
	}
	if r := recover(); r != nil {
		*err = errors.New(name + ":" + fmt.Sprintf("%s", r))
	}
}

// httpExecuter executes the http call with given headers and timeout
func httpExecuter(req *http.Request, resp *http.Response, headers map[string]string, timeOut time.Duration) (*APIResponse, error) {
	ret := new(APIResponse)
	var err error

	// add client headers
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	var client *http.Client
	// set client
	if isPoolSet() {
		if err = incNumCon(); err != nil {
			return ret, err
		}
		defer func() {
			if derr := decNumCon(); derr != nil {
				err = derr // set decrement error
			}
		}()
		client = &http.Client{Transport: poolObj.transport, Timeout: timeOut}
	} else {
		client = &http.Client{Timeout: timeOut}
	}
	resp, err = client.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()

	// read http status
	ret.HTTPStatus = constants.HTTPCode(resp.StatusCode)

	// read headers
	ret.Headers = make(map[string]string)
	for h, v := range resp.Header {
		ret.Headers[h] = v[0]
	}

	// read body
	body, berr := ioutil.ReadAll(resp.Body)
	if berr != nil {
		return nil, berr
	}
	ret.Body = body
	// return
	return ret, err
}

// getReqWithBody returns http request for given url and body
func getReqWithBody(name string, url string, body string) (req *http.Request, err error) {
	req, err = http.NewRequest(name, url, strings.NewReader(body))
	if err == nil {
		req.Close = true
		// must for body
		req.Header.Set("Content-Type", "application/json")
	}
	// return
	return req, err
}

// getReqWithoutBody returns http request for given url
func getReqWithoutBody(name string, url string) (req *http.Request, err error) {
	req, err = http.NewRequest(name, url, nil)
	if err == nil {
		req.Close = true
	}
	// return
	return req, err
}
