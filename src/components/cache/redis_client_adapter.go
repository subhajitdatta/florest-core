package cache

import (
	"strings"
	"time"

	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/common/utils/misc"
	"gopkg.in/redis.v3"
)

type mgetResult struct {
	Keys           []string
	SliceCmdOutput *redis.SliceCmd
}

type mdelResult struct {
	Keys         []string
	IntCmdOutput *redis.IntCmd
}

type RedisClientAdapter struct {
	client redisClientInterface
	hashes []string
}

func (ra *RedisClientAdapter) getHashKey(key string) string {
	hash := ra.getHash(key)
	return ra.getHashKeyFromHash(key, hash)
}

func (ra *RedisClientAdapter) getHashKeyFromHash(key string, hash string) string {
	return "{" + hash + "}" + key
}

func (ra *RedisClientAdapter) getHash(key string) string {
	hash := misc.GetHash(key, len(ra.hashes))
	return ra.hashes[hash]
}

func (ra *RedisClientAdapter) Init(conf *Config) error {
	if conf.Cluster {
		ra.client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    strings.Split(conf.ConnStr, ","),
			Password: conf.Password,
		})
	} else {
		ra.client = redis.NewClient(&redis.Options{
			Addr:     conf.ConnStr,
			Password: conf.Password,
		})
	}
	ra.hashes = conf.BucketHashes
	return nil
}

func (ra *RedisClientAdapter) Get(key string, serialize bool, compress bool) (item *Item, err error) {
	hashKey := ra.getHashKey(key)
	val, getErr := ra.client.Get(hashKey).Result()
	if getErr != nil {
		return nil, getErrObj(ErrGetFailure, "Getting key failed with error : "+getErr.Error())
	}
	item = new(Item)
	item.Key = key
	item.Value = val
	return item, nil
}

func (ra *RedisClientAdapter) Set(item Item, serialize bool, compress bool) error {
	hashKey := ra.getHashKey(item.Key)
	err := ra.client.Set(hashKey, item.Value, 0).Err()
	if err != nil {
		return getErrObj(ErrSetFailure, "Setting key failed with error : "+err.Error())
	}
	return nil
}

func (ra *RedisClientAdapter) SetWithTimeout(item Item, serialize bool, compress bool, ttl int32) error {
	hashKey := ra.getHashKey(item.Key)
	err := ra.client.Set(hashKey, item.Value, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		return getErrObj(ErrSetFailure, "Setting key with ttl failed with error : "+err.Error())
	}
	return nil
}

func (ra *RedisClientAdapter) Delete(key string) error {
	hashKey := ra.getHashKey(key)
	val, err := ra.client.Del(hashKey).Result()

	if err != nil {
		return getErrObj(ErrDeleteFailure, "Deleting key failed with error : "+err.Error())
	}

	if val != 1 {
		logger.Info("Delete failed because the key does not exist on the server")
	}

	return nil
}

func (ra *RedisClientAdapter) DeleteBatch(keys []string) error {
	hashCountTemp := 0
	hashKeysMap := make(map[string][]string)
	keysMap := make(map[string][]string)
	valuesChannel := make(chan *mdelResult)
	defer close(valuesChannel)

	for _, key := range keys {
		hash := ra.getHash(key)
		if hashKeysMap[hash] == nil {
			hashCountTemp++
			hashKeysMap[hash] = make([]string, 0)
			keysMap[hash] = make([]string, 0)
		}
		hashKeysMap[hash] = append(hashKeysMap[hash], ra.getHashKeyFromHash(key, hash))
		keysMap[hash] = append(keysMap[hash], key)
	}

	for hash, hashKeys := range hashKeysMap {
		go func(keys []string, hashKeys []string) {
			result := new(mdelResult)
			result.Keys = keys
			result.IntCmdOutput = ra.client.Del(hashKeys...)
			valuesChannel <- result
		}(keysMap[hash], hashKeys)
	}

	for i := 0; i < hashCountTemp; i++ {
		result := <-valuesChannel
		val, err := result.IntCmdOutput.Result()

		if err != nil {
			return getErrObj(ErrDeleteBatchFailure, "Deleting bulk keys failed with error : "+err.Error())
		}

		if val != int64(len(result.Keys)) {
			logger.Info("Delete failed because the keys does not exists in the server")
		}
	}

	return nil
}

func (ra *RedisClientAdapter) GetBatch(keys []string, serialize bool, compress bool) (items map[string]*Item, err error) {
	resMap := make(map[string]*Item, len(keys))
	hashCountTemp := 0
	hashKeysMap := make(map[string][]string)
	keysMap := make(map[string][]string)
	valuesChannel := make(chan *mgetResult)
	defer close(valuesChannel)

	for _, key := range keys {
		hash := ra.getHash(key)
		if hashKeysMap[hash] == nil {
			hashCountTemp++
			hashKeysMap[hash] = make([]string, 0)
			keysMap[hash] = make([]string, 0)
		}
		hashKeysMap[hash] = append(hashKeysMap[hash], ra.getHashKeyFromHash(key, hash))
		keysMap[hash] = append(keysMap[hash], key)
	}

	for hash, hashKeys := range hashKeysMap {
		go func(keys []string, hashKeys []string) {
			result := new(mgetResult)
			result.Keys = keys
			result.SliceCmdOutput = ra.client.MGet(hashKeys...)
			valuesChannel <- result
		}(keysMap[hash], hashKeys)
	}

	for i := 0; i < hashCountTemp; i++ {
		result := <-valuesChannel
		vals, err := result.SliceCmdOutput.Result()
		if err != nil {
			return nil, getErrObj(ErrGetBatchFailure, "Getting bulk keys failed with error : "+err.Error())
		}

		for index, val := range vals {
			item := new(Item)
			item.Key = result.Keys[index]
			if val != nil {
				item.Value = val
			} else {
				item.Error = "The key does not exist on the server"
			}
			resMap[item.Key] = item
		}
	}
	return resMap, nil
}
