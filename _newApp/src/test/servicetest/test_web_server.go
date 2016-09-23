package servicetest

import (
	"github.com/jabong/florest-core/src/core/service"
	"net/http"
	"net/http/httptest"
)

type testWebserver struct {
}

func (ws testWebserver) Response(req *http.Request) *httptest.ResponseRecorder {

	w := httptest.NewRecorder()
	webServer := new(service.Webserver)
	webServer.ServiceHandler(w, req)

	return w

}

var testHttpServer *testWebserver = nil

func initialiseTestWebServer() {
	if testHttpServer == nil {
		testHttpServer = new(testWebserver)
	}
}

func GetResponse(req *http.Request) *httptest.ResponseRecorder {
	return testHttpServer.Response(req)
}
