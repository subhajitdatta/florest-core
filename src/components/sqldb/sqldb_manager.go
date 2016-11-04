package sqldb

import (
	"github.com/jabong/florest-core/src/common/collections/maps/concurrentmap/concurrenthashmap"
	"reflect"
)

// sdbMap map to store cache interface
var sdbMap = concurrenthashmap.New()

// Set() stores the key with given type post init check
func Set(key string, conf *SDBConfig, obj interface{}) *SDBError {
	if val, ok := obj.(SDBInterface); ok {
		if _, ok = sdbMap.Get(key); ok {
			return getErrObj(ErrKeyPresent, "given key:"+key)
		}
		// check error for initialization
		if err := val.Init(conf); err != nil {
			return getErrObj(ErrInitialization, err.Error())
		}
		// store the new key
		sdbMap.Put(key, val)
		return nil
	} else {
		return getErrObj(ErrWrongType, reflect.TypeOf(obj).String()+":does not implement SDBInterface")
	}

}

// Get() - returns the sql db interface for given key
func Get(key string) (SDBInterface, *SDBError) {
	if val, ok := sdbMap.Get(key); !ok {
		return nil, getErrObj(ErrKeyNotPresent, "given key:"+key)
	} else {
		return val.(SDBInterface), nil
	}
}
