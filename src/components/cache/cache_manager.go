package cache

import (
	"reflect"
	"sync"
)

// cacheMap map to store cache interface
var cacheMap = make(map[string]CInterface)

// mutex used to edit the map with lock
var mutex *sync.Mutex = new(sync.Mutex)

// Set() stores the key with given type post init check
func Set(key string, conf *Config, obj interface{}) error {
	if val, ok := obj.(CInterface); ok {
		mutex.Lock()
		defer mutex.Unlock()
		if _, ok = cacheMap[key]; ok {
			return getErrObj(ErrKeyPresent, "given key:"+key)
		}
		// check error for initialization
		if err := val.Init(conf); err != nil {
			return getErrObj(ErrInitialization, err.Error())
		}
		// store the new key
		cacheMap[key] = val
		return nil
	} else {
		return getErrObj(ErrWrongType, reflect.TypeOf(obj).String()+":does not implement CInterface")
	}

}

// Get() - returns the cache interface for given key
func Get(key string) (CInterface, error) {
	mutex.Lock() // lock required as of go 1.6 concurrent read and write are not safe in map
	defer mutex.Unlock()
	if val, ok := cacheMap[key]; !ok {
		return nil, getErrObj(ErrKeyNotPresent, "given key:"+key)
	} else {
		return val, nil
	}
}
