package api

import (
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
)

type testExecutor struct {
	id string
}

func (n testExecutor) Name() string {
	return "Test Executor"
}

func (n *testExecutor) SetID(id string) {
	n.id = id
}

func (n testExecutor) GetID() (id string, err error) {
	return n.id, nil
}

func (n testExecutor) Execute(data workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	return data, nil
}
