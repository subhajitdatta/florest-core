package orchestrator

import (
	"testing"
)

func TestWorkflowStateCreate(t *testing.T) {
	testWorkflowState := new(workFlowState)
	testWorkflowState.create()

	if testWorkflowState.value == nil {
		t.Error("Failed to create Workflow State")
	}
}

func TestWorkflowStateSetGet(t *testing.T) {
	testWorkflowState := new(workFlowState)
	testWorkflowState.create()

	if testWorkflowState.value == nil {
		t.Error("Failed to create Workflow State")
	}

	serr := testWorkflowState.set("TEST_KEY", "TEST_VALUE")
	if serr != nil {
		t.Error("Failed to set workflow state")
	}

	value, gerr := testWorkflowState.Get("TEST_KEY")
	if gerr != nil {
		t.Error("Failed to get workflow state")
	}

	if value != "TEST_VALUE" {
		t.Error("Mismatch in the workflow state get")
	}

}

func TestWorkflowStateGetAll(t *testing.T) {
	testWorkflowState := new(workFlowState)
	testWorkflowState.create()

	if testWorkflowState.value == nil {
		t.Error("Failed to create Workflow State")
	}

	serr := testWorkflowState.set("TEST_KEY", "TEST_VALUE")
	if serr != nil {
		t.Error("Failed to set workflow state")
	}

	s1err := testWorkflowState.set("TEST_KEY1", "TEST_VALUE1")
	if s1err != nil {
		t.Error("Failed to set workflow state")
	}

	statesMap := testWorkflowState.GetAll()
	if statesMap == nil {
		t.Error("Failed to get all workflow state")
	}

	v1, v1Found := statesMap["TEST_KEY"]
	if !v1Found {
		t.Error("Failed to get workflow state key value")
	}
	if v1 != "TEST_VALUE" {
		t.Error("Mismatch in value returned for workflow state key")
	}

	v2, v2Found := statesMap["TEST_KEY1"]
	if !v2Found {
		t.Error("Failed to get workflow state key value")
	}
	if v2 != "TEST_VALUE1" {
		t.Error("Mismatch in value returned for workflow state key")
	}

}
