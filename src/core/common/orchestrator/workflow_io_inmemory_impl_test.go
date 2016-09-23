package orchestrator

import (
	"testing"
)

func TestWorkFlowIOInMemoryImplGetSet(t *testing.T) {
	testWorkFlowIOInMemoryImpl := new(WorkFlowIOInMemoryImpl)

	serr := testWorkFlowIOInMemoryImpl.Set("TEST_KEY", "TEST_VALUE")
	if serr != nil {
		t.Error("Failed to Set Workflow IO key in IO inmemory implementation")
	}

	value, gerr := testWorkFlowIOInMemoryImpl.Get("TEST_KEY")
	if gerr != nil {
		t.Error("Failed to Get Workflow IO key in IO inmemory implementation")
	}

	keyValue, ok := value.(string)
	if !ok {
		t.Error("Data type mismatch for Get IO key in Workflow IO inmemory implementation")
	}

	if keyValue != "TEST_VALUE" {
		t.Error("Key Value mismatch for Get IO key in Workflow IO inmemory implementation")
	}
}
