package orchestrator

import (
	"reflect"
	"testing"
)

const (
	testKey1   = "TEST_KEY_1"
	testValue1 = "TEST_VALUE_1"

	testKey2   = "TEST_KEY_2"
	testValue2 = "TEST_VALUE_2"
)

/*
Test workflow execution context inmemory implementation Get and Set
*/
func TestWorkflowECInmemoryImplGetSet(t *testing.T) {
	testWorkflowECInstance := new(WorkFlowECInMemoryImpl)

	t1err := testWorkflowECInstance.Set(testKey1, testValue1)
	if t1err != nil {
		t.Error("Failed to Set Workflow Execution Context Key")
	}
	t2err := testWorkflowECInstance.Set(testKey2, testValue2)
	if t2err != nil {
		t.Error("Failed to Set Workflow Execution Context Key")
	}

	_, t1verr := testWorkflowECInstance.Get(testKey1)
	if t1verr != nil {
		t.Error("Failed to Get Workflow Execution Context Key")
	}
	_, t2verr := testWorkflowECInstance.Get(testKey2)
	if t2verr != nil {
		t.Error("Failed to Get Workflow Execution Context Key")
	}
}

/*
Test workflow execution context inmemory implementation Get and Set value
*/
func TestWorkflowECInmemoryImplGetSetValue(t *testing.T) {
	testWorkflowECInstance := new(WorkFlowECInMemoryImpl)

	t1err := testWorkflowECInstance.Set(testKey1, testValue1)
	if t1err != nil {
		t.Error("Failed to Set Workflow Execution Context Key")
	}

	value, t1verr := testWorkflowECInstance.Get(testKey1)
	if t1verr != nil {
		t.Error("Failed to Get Workflow Execution Context Key")
	}

	_, ok := value.(string)
	if !ok {
		t.Error("Failed to match value data type from Get key in Workflow Execution Context")
	}
}

/*
Test workflow execution context inmemory implementation Get and Set Buckets
*/
func TestWorkflowECInmemoryImplGetSetBuckets(t *testing.T) {
	testWorkflowECInstance := new(WorkFlowECInMemoryImpl)

	bucketsList := map[string]string{"bucket1": "value1", "bucket2": "value"}
	t1err := testWorkflowECInstance.SetBuckets(bucketsList)
	if t1err != nil {
		t.Error("Failed to Set Workflow Execution Context Buckets")
	}

	buckets, t1verr := testWorkflowECInstance.GetBuckets()
	if t1verr != nil {
		t.Error("Failed to Get Workflow Execution Context Buckets")
	}

	eq := reflect.DeepEqual(buckets, bucketsList)

	if !eq {
		t.Error("Failed to Get Buckets in Workflow Execution Context")
	}

}

/*
Test workflow execution context inmemory implementation Debug information
*/
func TestWorkflowECInmemoryImplDebugMsg(t *testing.T) {
	testWorkflowECInstance := new(WorkFlowECInMemoryImpl)

	testWorkflowECInstance.SetDebugFlag(true)
	testWorkflowECInstance.SetDebugMsg("TEST KEY 1", "TEST MESSAGE")

	_, merr := testWorkflowECInstance.GetDebugMsg()
	if merr != nil {
		t.Error("Error in Workflow Definition Execution Context Inmemory implementation debug message")
	}
}

/*
Test workflow execution context inmemory implementation ThreadId
*/
func TestWorkflowECInmemoryImplThreadId(t *testing.T) {
	testWorkflowECInstance := new(WorkFlowECInMemoryImpl)

	testWorkflowECInstance.Set(threadID, "TestThread1")
	tID, terr := testWorkflowECInstance.GetExecuteThreadID()
	if terr != nil || tID != "TestThread1" {
		t.Error("Error in workflow definition Execution Context Inmemory implementation theread id retrieval")
	}
}
