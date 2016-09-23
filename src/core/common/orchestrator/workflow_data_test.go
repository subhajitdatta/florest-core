package orchestrator

import (
	"testing"
)

func TestWorkflowDataCreate(t *testing.T) {
	testWorkFlowData := new(WorkFlowData)
	testWorkflowIOData := new(WorkFlowIOInMemoryImpl)
	testWorkflowECData := new(WorkFlowECInMemoryImpl)

	testWorkFlowData.Create(testWorkflowIOData, testWorkflowECData)

	if testWorkFlowData.IOData == nil {
		t.Error("Failed to initialize workflow Input Output Data")
	}

	if testWorkFlowData.ExecContext == nil {
		t.Error("Failed to initialize workflow Execution Context Data")
	}
}

func TestWorkflowDataSetGetSate(t *testing.T) {
	testWorkFlowData := new(WorkFlowData)
	testWorkflowIOData := new(WorkFlowIOInMemoryImpl)
	testWorkflowECData := new(WorkFlowECInMemoryImpl)

	testWorkFlowData.Create(testWorkflowIOData, testWorkflowECData)

	t1err := testWorkFlowData.setWorkflowState("TEST NODE 1", "TEST NODE 1 STATE")
	if t1err != nil {
		t.Error("Failed to Set Workflow State from Workflow Data")
	}
	t2err := testWorkFlowData.setWorkflowState("TEST NODE 2", "TEST NODE 2 STATE")
	if t2err != nil {
		t.Error("Failed to Set Workflow State from Workflow Data")
	}

	states := testWorkFlowData.GetWorkflowState()

	_, foundTestNode1State := states["TEST NODE 1"]
	if !foundTestNode1State {
		t.Error("Failed to Get Workflow State")
	}
	_, foundTestNode2State := states["TEST NODE 2"]
	if !foundTestNode2State {
		t.Error("Failed to Get Workflow State")
	}

}
