package perftest

import (
	"net/http"
	"runtime"
	"testing"

	"github.com/jabong/florest-core/src/core/service"
	testUtil "github.com/jabong/florest-core/src/test/utils"
)

func BenchmarkFlorestParam(b *testing.B) {
	loadFlorestSingle("GET", "v1", "test", "user/{name}")
	benchFlorestRequest(b, "GET", "/user/gordon")
}

func benchFlorestRequest(b *testing.B, action string, path string) {
	request := testUtil.CreateTestRequest("GET", "florest/v1/test"+path)

	w := new(mockResponseWriter)
	webServer := new(service.Webserver)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		webServer.ServiceHandler(w, request)
	}
}

func benchFlorestRoutes(b *testing.B, routes []route) {
	w := new(mockResponseWriter)
	webServer := new(service.Webserver)

	r, _ := http.NewRequest("GET", "/", nil)
	u := r.URL
	rq := u.RawQuery

	b.ReportAllocs()
	b.ResetTimer()

	florestPath := "florest/v1/test"

	for i := 0; i < b.N; i++ {
		for _, route := range routes {
			r.Method = route.method
			r.RequestURI = florestPath + route.path
			u.Path = florestPath + route.path
			u.RawQuery = rq
			webServer.ServiceHandler(w, r)
		}
	}
}

func benchHTTPRouteRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := new(mockResponseWriter)
	u := r.URL
	rq := u.RawQuery
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.RawQuery = rq
		router.ServeHTTP(w, r)
	}
}

func benchHTTPRoutes(b *testing.B, router http.Handler, routes []route) {
	w := new(mockResponseWriter)
	r, _ := http.NewRequest("GET", "/", nil)
	u := r.URL
	rq := u.RawQuery

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, route := range routes {
			r.Method = route.method
			r.RequestURI = route.path
			u.Path = route.path
			u.RawQuery = rq
			router.ServeHTTP(w, r)
		}
	}
}

func calcMem(name string, load func()) {
	m := new(runtime.MemStats)

	// before
	runtime.GC()
	runtime.ReadMemStats(m)
	before := m.HeapAlloc

	load()

	// after
	runtime.GC()
	runtime.ReadMemStats(m)
	after := m.HeapAlloc
	println("   "+name+":", after-before, "Bytes")
}
