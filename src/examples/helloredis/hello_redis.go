package helloredis

import (
	"fmt"

	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/components/cache"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
	expConf "github.com/jabong/florest-core/src/examples/config"
)

type redisNode struct {
	id string
}

func (a *redisNode) SetID(id string) {
	a.id = id
}

func (a redisNode) GetID() (id string, err error) {
	return a.id, nil
}

func (a redisNode) Name() string {
	return "redisNode"
}

func (a redisNode) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	appConfig, err := expConf.GetExampleAppConfig()
	if err != nil {
		msg := "Redis App Config Not Correct"
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	cacheConf := appConfig.Cache
	if cacheConf == nil {
		msg := "No Cache Config Specified"
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	redisConf := cacheConf.Redis
	if redisConf == nil {
		msg := "No Redis Config Specified"
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	cacheObj, errG := cache.Get(*redisConf)
	if errG != nil {
		msg := fmt.Sprintf("Redis Config Error - %v", errG)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	if errI := cacheObj.Init(redisConf); errI != nil {
		msg := fmt.Sprintf("Error in initializing cache object %v", errI)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	item := cache.Item{
		Key:   "TestKey",
		Value: "TestValue",
	}
	if errS := cacheObj.Set(item, false, false); errS != nil {
		msg := fmt.Sprintf("Error in setting an item in redis %v", errS)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	v, gerr := cacheObj.Get("TestKey", false, false)
	if gerr != nil {
		msg := fmt.Sprintf("Failed to get item from cache %v", gerr)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}

	}
	logger.Info("Received a value from cache %v", v.Value)
	//Business Logic
	io.IOData.Set(constants.Result, "Get & Set in Redis Successful")
	return io, nil
}
