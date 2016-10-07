package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/imdario/mergo"
	"github.com/jabong/florest-core/src/common/config"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/components/cache"
)

var cacheImpl cache.CInterface
var refreshInterval int
var configKey string
var initialized = false

type DynamicConfigManager struct {
	applicationConfig interface{}
}

/**
 * Initialize dynamic config manager and start the refresh timer
 */
func (dcm *DynamicConfigManager) Initialize(applicationConfig interface{}, cacheKey string) {
	//Check if the Dynamic config is already initialized
	if initialized {
		return
	}
	dcm.applicationConfig = applicationConfig
	initialized = true
	dynamicConfObj := config.GlobalAppConfig.DynamicConfig
	if !dynamicConfObj.Active {
		logger.Info("Dynamic configuration is not active. Hence, config auto-refresh will not happen at runtime")
		return
	}
	var err error
	cacheImpl, err = cache.Get(cacheKey)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize cache to auto-refresh the config \n %+v \n %s", dynamicConfObj.CacheKey, err))
	}

	refreshInterval = dynamicConfObj.RefreshInterval
	configKey = dynamicConfObj.ConfigKey

	go dcm.refreshConfigAtEveryInterval()
	fmt.Println("Dynamic config is initialized")
}

/**
 * Starts the timer to refresh the config at every refresh interval
 */
func (dcm *DynamicConfigManager) refreshConfigAtEveryInterval() {
	refreshNow := time.NewTicker(time.Second * time.Duration(refreshInterval)).C
	for range refreshNow {
		dcm.refreshConfig()
	}
}

/**
 * Gets the updated config from Blitz and merges with the current application config
 */
func (dcm *DynamicConfigManager) refreshConfig() {
	logger.Info("Refreshing the config. Time now : " + time.Now().String())
	data, _ := cacheImpl.Get(configKey, true, true)
	if data != nil && data.Value != nil {
		var raw json.RawMessage
		newAppConfig := dcm.applicationConfig
		if dataValue, ok := data.Value.(string); ok {
			raw = json.RawMessage(dataValue)
		} else {
			logger.Warning(fmt.Sprintf("Error - cannot convert to type string"))
			return
		}
		dataInBytes, err := json.Marshal(&raw)
		if err != nil {
			logger.Warning(fmt.Sprintf("Incorrect Json. Error - %s", err))
			return
		}
		err = json.Unmarshal(dataInBytes, newAppConfig)
		if err != nil {
			logger.Warning(fmt.Sprintf("Incorrect Json. Error - %s", err))
			return
		}
		configCopy := config.GlobalAppConfig.ApplicationConfig
		// Mergo : Library to merge structs and maps in Golang.
		err = mergo.MergeWithOverwrite(config.GlobalAppConfig.ApplicationConfig, newAppConfig)
		if err != nil {
			logger.Warning(fmt.Sprintf("Failed to merge the application config. Error - %s", err))
			config.GlobalAppConfig.ApplicationConfig = configCopy
		}
	} else {
		logger.Warning("Could not find the dynamic config - key : " + configKey + " in central config cache")
	}
}
