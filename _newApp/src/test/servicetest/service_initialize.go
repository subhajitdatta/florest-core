package servicetest

import (
	"common/appconstant"
	"github.com/jabong/florest-core/src/core/service"
	"hello"
)

func InitializeTestService() {
	service.RegisterHTTPErrors(appconstant.APPErrorCodeToHTTPCodeMap)
	service.RegisterAPI(new(hello.HelloAPI))
	initTestLogger()

	initTestConfig()

	service.InitVersionManager()

	initialiseTestWebServer()

	service.InitHealthCheck()

}

func PurgeTestService() {

}
