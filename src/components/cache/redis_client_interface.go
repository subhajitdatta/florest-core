package cache

import (
	"gopkg.in/redis.v3"
	"time"
)

type redisClientInterface interface {
	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(keys ...string) *redis.IntCmd
	MGet(keys ...string) *redis.SliceCmd
	MSet(keys ...string) *redis.StatusCmd
}
