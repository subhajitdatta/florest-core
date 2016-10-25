package servicetest

import (
	"github.com/jabong/florest-core/src/core/service"
	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	testUtil "test/utils"
)

type testRateLimitWebserver struct {
	webServer *service.Webserver
}

func (ws *testRateLimitWebserver) Init() {
	ws.webServer = new(service.Webserver)
}

func (ws testRateLimitWebserver) Response(req *http.Request) *httptest.ResponseRecorder {

	w := httptest.NewRecorder()
	ws.webServer.ServiceHandler(w, req)

	return w

}

func testRateLimitedAPI() {
	apiName := "florest"
	gk.Describe("GET"+"/"+apiName+"/V1/TESTRATE/", func() {

		testRateLimitHTTPServer := new(testRateLimitWebserver)
		testRateLimitHTTPServer.Init()

		gk.Context("then the response", func() {
			gk.It("should be rate limited", func() {

				request1 := testUtil.CreateTestRequest("GET", "/"+apiName+"/V1/TESTRATE/")
				response1 := testRateLimitHTTPServer.Response(request1)
				gm.Expect(response1.HeaderMap.Get("Content-Type")).To(gm.Equal("application/json"))
				gm.Expect(response1.HeaderMap.Get("Cache-Control")).To(gm.Equal(""))
				gm.Expect(response1.Code).To(gm.Equal(200))

				request2 := testUtil.CreateTestRequest("GET", "/"+apiName+"/V1/TESTRATE/")
				response2 := testRateLimitHTTPServer.Response(request2)
				gm.Expect(response2.HeaderMap.Get("Content-Type")).To(gm.Equal("application/json"))
				gm.Expect(response2.HeaderMap.Get("Cache-Control")).To(gm.Equal(""))
				gm.Expect(response2.Code).To(gm.Equal(429))
			})
		})
	})
}
