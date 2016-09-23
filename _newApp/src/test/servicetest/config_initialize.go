package servicetest

import (
	"github.com/jabong/florest-core/src/core/service"
)

func initTestConfig() {
	cm := new(service.ConfigManager)
	cm.InitializeGlobalConfig("testdata/testconf.json")
}
