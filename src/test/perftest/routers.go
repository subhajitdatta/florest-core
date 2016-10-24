package perftest

import (
	"net/http"
	"strings"

	"github.com/jabong/florest-core/src/core/service"
	"github.com/jabong/florest-core/src/test/api"
	testService "github.com/jabong/florest-core/src/test/servicetest"
)

type route struct {
	method string
	path   string
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {}

func loadHTTPServeMux(routes []route) http.Handler {
	serveMux := http.NewServeMux()
	for _, route := range routes {
		serveMux.HandleFunc(route.path, httpHandlerFunc)
	}
	return serveMux
}

func loadFlorestSingle(action string, version string, resource string, path string) {
	testAPI := new(api.TestAPI)
	testAPI.SetVersion(action, version, resource, path)
	service.RegisterAPI(testAPI)
	testService.InitializeTestService()
}

func loadFlorest(version string, resource string, routes []route) {
	for _, route := range routes {
		testAPI := new(api.TestAPI)
		testAPI.SetVersion(route.method, version, resource, route.path[1:])
		service.RegisterAPI(testAPI)
	}
	testService.InitializeTestService()
}

func convertColonToBraces(routes []route) []route {
	result := make([]route, len(routes))

	for index, oldRoute := range routes {
		path := oldRoute.path[1:]
		pathArr := strings.Split(path, "/")
		convertedPath := ""
		for _, pathParam := range pathArr {
			if strings.Index(pathParam, ":") == 0 {
				pathParam = "{" + pathParam[1:] + "}"
			}
			convertedPath = convertedPath + "/" + pathParam
		}
		newRoute := new(route)
		newRoute.path = convertedPath
		newRoute.method = oldRoute.method
		result[index] = *newRoute
	}

	return result
}
