package hellorediscluster

import (
	"fmt"

	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/components/cache"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
)

type redisClusterNode struct {
	id string
}

func (a *redisClusterNode) SetID(id string) {
	a.id = id
}

func (a redisClusterNode) GetID() (id string, err error) {
	return a.id, nil
}

func (a redisClusterNode) Name() string {
	return "redisClusterNode"
}

func (a redisClusterNode) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	cacheAdapter, errG := cache.Get("myRedisCluster") // It should be called only once and can be shared across go routines
	if errG != nil {
		msg := fmt.Sprintf("Redis Cluster Config Error - %v", errG)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	// Put some items with TTL
	item1 := cache.Item{Key: "somekey1", Value: "somevalue1", Error: ""}
	item2 := cache.Item{Key: "somekey2", Value: "somevalue2", Error: ""}
	item3 := cache.Item{Key: "somekey3", Value: "somevalue3", Error: ""}

	if errT := cacheAdapter.SetWithTimeout(item1, false, false, 1000); errT != nil {
		msg := fmt.Sprintf("Error in setting keys in cache item1. Error - %v", errT)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	if errT2 := cacheAdapter.SetWithTimeout(item2, false, false, 1000); errT2 != nil {
		msg := fmt.Sprintf("Error in setting keys in cache item2. Error - %v", errT2)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	if errT3 := cacheAdapter.SetWithTimeout(item3, false, false, 1000); errT3 != nil {
		msg := fmt.Sprintf("Error in setting keys in cache item3. Error - %v", errT3)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	logger.Info("Setting items are successful")

	// Get an item
	item, errG2 := cacheAdapter.Get("somekey1", false, false)
	if errG2 != nil {
		msg := fmt.Sprintf("Getting item from cache failed. Error %v", errG2)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	logger.Info("Got the item - key : " + item.Key + ", value : " + item.Value.(string))

	// Get bulk items
	keys := []string{"somekey1", "somekey2", "somekey4"}

	items, errGB := cacheAdapter.GetBatch(keys, false, false)
	if errGB != nil {
		msg := fmt.Sprintf("Getting bulk items from cache failed. Error - %v", errGB)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	logger.Info("Got bulk items " + items["somekey1"].Value.(string) + ", " + items["somekey2"].Value.(string) + ", " + items["somekey4"].Error)

	// Delete an item
	if errD := cacheAdapter.Delete("somekey1"); errD != nil {
		msg := fmt.Sprintf("Failed to delete item from redis cluster Error - %v", errD)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	logger.Info("Item deleted successfully..")

	// Delete bulk items
	keysToDelete := []string{"somekey2", "somekey3"}
	if errKD := cacheAdapter.DeleteBatch(keysToDelete); errKD != nil {
		msg := fmt.Sprintf("Error in deleting bulk items from cache. Error - %v", errKD)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	logger.Info("Bulk items deleted successfully..")
	io.IOData.Set(constants.Result, "Get Set Delete and Batch Operation successful")
	//Business Logic
	return io, nil
}
