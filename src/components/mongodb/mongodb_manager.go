package mongodb

import (
	"reflect"
	"sync"
)

// mongoMap map to store cache interface
var mongoMap = make(map[string]MDBInterface)

// mutex used to edit the map with lock
var mutex *sync.Mutex = new(sync.Mutex)

// Set() stores the key with given type post init check
func Set(key string, conf *MDBConfig, obj interface{}) *MDBError {
	if val, ok := obj.(MDBInterface); ok {
		mutex.Lock()
		defer mutex.Unlock()
		if _, ok = mongoMap[key]; ok {
			return getErrObj(ErrKeyPresent, "given key:"+key)
		}
		// check error for initialization
		if err := val.Init(conf); err != nil {
			return getErrObj(ErrInitialization, err.Error())
		}
		// store the new key
		mongoMap[key] = val
		return nil
	} else {
		return getErrObj(ErrWrongType, reflect.TypeOf(obj).String()+":does not implement MDBInterface")
	}

}

// Get() - returns the mongodb interface for given key
func Get(key string) (MDBInterface, *MDBError) {
	mutex.Lock() // lock required as of go 1.6 concurrent read and write are not safe in map
	defer mutex.Unlock()
	if val, ok := mongoMap[key]; !ok {
		return nil, getErrObj(ErrKeyNotPresent, "given key:"+key)
	} else {
		return val, nil
	}
}
