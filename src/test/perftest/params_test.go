package perftest

import (
	"net/http"
	"testing"
)

var parseAPIRoutes = []route{
	// Objects
	{"POST", "/1/classes/:className"},
	{"GET", "/1/classes/:className/:objectId"},
	{"PUT", "/1/classes/:className/:objectId"},
	{"GET", "/1/classes/:className"},
	{"DELETE", "/1/classes/:className/:objectId"},

	// Users
	{"POST", "/1/users"},
	{"GET", "/1/login"},
	{"GET", "/1/users/:objectId"},
	{"PUT", "/1/users/:objectId"},
	{"GET", "/1/users"},
	{"DELETE", "/1/users/:objectId"},
	{"POST", "/1/requestPasswordReset"},

	// Roles
	{"POST", "/1/roles"},
	{"GET", "/1/roles/:objectId"},
	{"PUT", "/1/roles/:objectId"},
	{"GET", "/1/roles"},
	{"DELETE", "/1/roles/:objectId"},

	// Files
	{"POST", "/1/files/:fileName"},

	// Analytics
	{"POST", "/1/events/:eventName"},

	// Push Notifications
	{"POST", "/1/push"},

	// Installations
	{"POST", "/1/installations"},
	{"GET", "/1/installations/:objectId"},
	{"PUT", "/1/installations/:objectId"},
	{"GET", "/1/installations"},
	{"DELETE", "/1/installations/:objectId"},

	// Cloud Functions
	{"POST", "/1/functions"},
}

var parseAPIRoutesFlorest = convertColonToBraces(parseAPIRoutes)

var (
	parseHTTPRouter http.Handler
)

func init() {
	println("#Parse Routes:", len(parseAPIRoutes))

	calcMem("Florest_Parse", func() {
		loadFlorest("v1", "test", parseAPIRoutesFlorest)
	})
}

func BenchmarkFlorestParseStatic(b *testing.B) {
	benchFlorestRequest(b, "GET", "/1/users")
}

func BenchmarkFlorestParseParam(b *testing.B) {
	benchFlorestRequest(b, "GET", "/1/classes/go")
}

func BenchmarkFlorestParse2Params(b *testing.B) {
	benchFlorestRequest(b, "GET", "/1/classes/go/123456789")
}

func BenchmarkFlorestParamAll(b *testing.B) {
	benchFlorestRoutes(b, parseAPIRoutes)
}
