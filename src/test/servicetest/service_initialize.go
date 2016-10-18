package servicetest

import (
	"github.com/jabong/florest-core/src/core/common/env"
	"github.com/jabong/florest-core/src/core/service"
	"github.com/jabong/florest-core/src/test/api"
)

func InitializeTestService() {

	service.RegisterAPI(new(api.TestAPI))

	reqVNAPI := new(api.ReqVNAPI)
	reqVNAPI.SetVersion("GET", "V1", "REQVD", "")

	service.RegisterAPI(reqVNAPI)

	env.GetOsEnviron()

	initTestConfig()

	initTestLogger()

	service.InitMonitor()

	service.InitVersionManager()

	service.InitCustomAPIInit()

	service.InitApis()

	service.InitHealthCheck()

	initialiseTestWebServer()

}

func PurgeTestService() {

}
