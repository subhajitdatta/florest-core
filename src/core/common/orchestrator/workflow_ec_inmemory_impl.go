package orchestrator

import (
	"errors"
	"fmt"
)

const (
	buckets   string = "BUCKETS"
	threadID  string = "THREAD_ID"
	debugFlag string = "IS_DEBUG"
	debugMsg  string = "DEBUG_MSG"
)

type WorkFlowECInMemoryImpl struct {
	store map[string]interface{}
}

type WorkflowDebugDataInMemory struct {
	Key   string
	Value string
}

func (ec *WorkFlowECInMemoryImpl) Get(key string) (value interface{}, err error) {
	//Check if the key is already present
	res, found := ec.store[key]
	if !found {
		errString := fmt.Sprintln("In memory ExecutionContext store does not contain key: ", key)
		return nil, errors.New(errString)
	}
	return res, nil
}

func (ec *WorkFlowECInMemoryImpl) Set(key string, value interface{}) (err error) {
	if ec.store == nil {
		ec.store = make(map[string]interface{})
	}
	ec.store[key] = value
	return nil
}

func (ec *WorkFlowECInMemoryImpl) SetBuckets(bucketIDMap map[string]string) (err error) {
	return ec.Set(buckets, bucketIDMap)
}

func (ec *WorkFlowECInMemoryImpl) GetBuckets() (bucketIDMap map[string]string, err error) {
	res, err := ec.Get(buckets)
	if v, ok := res.(map[string]string); ok {
		bucketIDMap = v
	}
	return bucketIDMap, err
}

func (ec *WorkFlowECInMemoryImpl) GetExecuteThreadID() (executeThreadID string, err error) {
	res, err := ec.Get(threadID)
	if v, ok := res.(string); ok {
		executeThreadID = v
	}
	return executeThreadID, err
}

func (ec *WorkFlowECInMemoryImpl) SetDebugFlag(flag bool) (err error) {
	return ec.Set(debugFlag, flag)
}

func (ec *WorkFlowECInMemoryImpl) SetDebugMsg(msgkey string, msgData string) (err error) {
	if ec.store == nil {
		ec.store = make(map[string]interface{})
	}

	isDebugVal, isDebugSet := ec.store[debugFlag]
	if !isDebugSet {
		return nil
	}
	isDebug, ok := isDebugVal.(bool)
	if !ok {
		return errors.New("Incorrect data stored in debug flag")
	}
	if !isDebug {
		return nil
	}

	dMsg, found := ec.store[debugMsg]

	//This is the first debug msg
	if !found {
		newDebugMsg := WorkflowDebugDataInMemory{Key: msgkey, Value: msgData}
		ec.store[debugMsg] = []WorkflowDebugDataInMemory{newDebugMsg}
		return nil
	}

	//Debug Messages are already present
	v, ok := dMsg.([]WorkflowDebugDataInMemory)
	if ok {
		v = append(v, WorkflowDebugDataInMemory{Key: msgkey, Value: msgData})
	}
	ec.store[debugMsg] = v
	return nil
}

func (ec *WorkFlowECInMemoryImpl) GetDebugMsg() (msg []interface{}, err error) {
	msgData, err := ec.Get(debugMsg)
	if v, ok := msgData.([]WorkflowDebugDataInMemory); ok {
		msg = make([]interface{}, len(v))
		for i, val := range v {
			msg[i] = val
		}
	}
	return msg, err
}
