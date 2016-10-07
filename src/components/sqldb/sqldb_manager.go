package sqldb

import (
	"reflect"
	"sync"
)

// sdbMap map to store cache interface
var sdbMap = make(map[string]SDBInterface)

// mutex used to edit the map with lock
var mutex *sync.Mutex = new(sync.Mutex)

// Set() stores the key with given type post init check
func Set(key string, conf *SDBConfig, obj interface{}) *SDBError {
	if val, ok := obj.(SDBInterface); ok {
		mutex.Lock()
		defer mutex.Unlock()
		if _, ok = sdbMap[key]; ok {
			return getErrObj(ErrKeyPresent, "given key:"+key)
		}
		// check error for initialization
		if err := val.Init(conf); err != nil {
			return getErrObj(ErrInitialization, err.Error())
		}
		// store the new key
		sdbMap[key] = val
		return nil
	} else {
		return getErrObj(ErrWrongType, reflect.TypeOf(obj).String()+":does not implement SDBInterface")
	}

}

// Get() - returns the sql db interface for given key
func Get(key string) (SDBInterface, *SDBError) {
	mutex.Lock() // lock required as of go 1.6 concurrent read and write are not safe in map
	defer mutex.Unlock()
	if val, ok := sdbMap[key]; !ok {
		return nil, getErrObj(ErrKeyNotPresent, "given key:"+key)
	} else {
		return val, nil
	}
}
