package service

import (
	"github.com/jabong/florest-core/src/common/config"
	"github.com/jabong/florest-core/src/common/constants"
)

var apiList []APIInterface
var resourceBucketMapping map[string]string
var apiCustomInitFunc func()
var configEnvUpdateMap map[string]string
var globalEnvUpdateMap map[string]string

func RegisterAPI(apiInstance APIInterface) {
	apiList = append(apiList, apiInstance)
}

func RegisterConfig(applicationConfig interface{}) {
	config.GlobalAppConfig.ApplicationConfig = applicationConfig
}

func RegisterHTTPErrors(appErrorCodeMap map[constants.APPErrorCode]constants.HTTPCode) {
	constants.UpdateAppHTTPError(appErrorCodeMap)
}

func RegisterResourceBucketMapping(resource, bucketID string) {
	if len(resourceBucketMapping) == 0 {
		resourceBucketMapping = make(map[string]string)
	}
	resourceBucketMapping[resource] = bucketID
}

func RegisterCustomAPIInitFunc(f func()) {
	apiCustomInitFunc = f
}

func RegisterConfigEnvUpdateMap(a map[string]string) {
	configEnvUpdateMap = a
}

func RegisterGlobalEnvUpdateMap(a map[string]string) {
	globalEnvUpdateMap = a
}
