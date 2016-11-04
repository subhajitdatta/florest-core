package mongodb

import (
	"github.com/jabong/florest-core/src/common/collections/maps/concurrentmap/concurrenthashmap"
	"reflect"
)

// mongoMap map to store cache interface
var mongoMap = concurrenthashmap.New()

// Set() stores the key with given type post init check
func Set(key string, conf *MDBConfig, obj interface{}) *MDBError {
	if val, ok := obj.(MDBInterface); ok {
		if _, ok = mongoMap.Get(key); ok {
			return getErrObj(ErrKeyPresent, "given key:"+key)
		}
		// check error for initialization
		if err := val.Init(conf); err != nil {
			return getErrObj(ErrInitialization, err.Error())
		}
		// store the new key
		mongoMap.Put(key, val)
		return nil
	} else {
		return getErrObj(ErrWrongType, reflect.TypeOf(obj).String()+":does not implement MDBInterface")
	}

}

// Get() - returns the mongodb interface for given key
func Get(key string) (MDBInterface, *MDBError) {
	if val, ok := mongoMap.Get(key); !ok {
		return nil, getErrObj(ErrKeyNotPresent, "given key:"+key)
	} else {
		return val.(MDBInterface), nil
	}
}
