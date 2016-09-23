package servicetest

import (
	"github.com/jabong/florest-core/src/common/config"
	"github.com/jabong/florest-core/src/core/service"
)

func initTestConfig() {
	service.RegisterConfig(new(TestAPPConfig))

	cm := new(service.ConfigManager)
	cm.InitializeGlobalConfig("testdata/testconf.json")
	cm.UpdateConfigFromEnv(config.GlobalAppConfig, "global")

}
