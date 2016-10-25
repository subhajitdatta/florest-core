package ratelimiter

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/constants"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testWebserver struct {
	requestHandler http.HandlerFunc
}

func (o *testWebserver) Init(fn http.HandlerFunc) {
	o.requestHandler = fn
}

func (o *testWebserver) Response(req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	o.requestHandler(w, req)
	return w

}

func TestHTTPRateLimiterNotSet(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi there")
	}

	conf := new(Config)
	rateLimiter, err := New(conf)
	if err == nil && rateLimiter != nil {
		t.Fatal("Should not initialize rate limiter")
	}

	testHTTPServer := new(testWebserver)
	testHTTPServer.Init(MakeRateLimitedHTTPHandler(handler, rateLimiter, "MY APP"))

	req, _ := http.NewRequest("GET", "hello", nil)

	for i := 1; i <= 10; i++ {
		resp := testHTTPServer.Response(req)
		if constants.HTTPCode(resp.Code) != constants.HTTPStatusSuccessCode {
			t.Fatalf("Iteration %v : Rate limited http handler should be success %v, but got %v",
				i, constants.HTTPStatusSuccessCode, resp.Code)
		}
	}

}

func TestHTTPRateLimiter(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi there")
	}

	conf := new(Config)
	conf.Type = GCRA
	//2 request per second without burst
	conf.MaxRate = 2
	conf.MaxBurst = 5

	rateLimiter, err := New(conf)
	if err != nil {
		t.Fatal(err.Error())
	}

	testHTTPServer := new(testWebserver)
	testHTTPServer.Init(MakeRateLimitedHTTPHandler(handler, rateLimiter, "MY APP"))

	var i = 0
	for i = 1; i <= 6; i++ {
		req, _ := http.NewRequest("GET", "hello", nil)
		resp := testHTTPServer.Response(req)
		if constants.HTTPCode(resp.Code) != constants.HTTPStatusSuccessCode {
			t.Fatalf("Iteration %v : Rate limited http handler should be success %v, but got %v",
				i, constants.HTTPStatusSuccessCode, resp.Code)
		}
	}

	//Rate limit should now be exceeded
	req, _ := http.NewRequest("GET", "hello", nil)
	resp := testHTTPServer.Response(req)
	if constants.HTTPCode(resp.Code) != constants.HTTPRateLimitExceeded {
		t.Fatalf("Iteration %v : Rate limited http handler should have rate limited %v, but got %v",
			i, constants.HTTPRateLimitExceeded, resp.Code)
	}

}
