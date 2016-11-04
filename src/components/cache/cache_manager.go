package cache

import (
	"github.com/jabong/florest-core/src/common/collections/maps/concurrentmap/concurrenthashmap"
	"reflect"
)

// cacheMap map to store cache interface
var cacheMap = concurrenthashmap.New()

// Set() stores the key with given type post init check
func Set(key string, conf *Config, obj interface{}) error {
	if val, ok := obj.(CInterface); ok {
		if _, ok = cacheMap.Get(key); ok {
			return getErrObj(ErrKeyPresent, "given key:"+key)
		}
		// check error for initialization
		if err := val.Init(conf); err != nil {
			return getErrObj(ErrInitialization, err.Error())
		}
		// store the new key
		cacheMap.Put(key, val)
		return nil
	} else {
		return getErrObj(ErrWrongType, reflect.TypeOf(obj).String()+":does not implement CInterface")
	}

}

// Get() - returns the cache interface for given key
func Get(key string) (CInterface, error) {
	if val, ok := cacheMap.Get(key); !ok {
		return nil, getErrObj(ErrKeyNotPresent, "given key:"+key)
	} else {
		return val.(CInterface), nil
	}
}
