package hello

import (
	errors "common/appconstant"
	florest_constants "github.com/jabong/florest-core/src/common/constants"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
)

type HelloWorld struct {
	id string
}

func (n *HelloWorld) SetID(id string) {
	n.id = id
}

func (n HelloWorld) GetID() (id string, err error) {
	return n.id, nil
}

func (a HelloWorld) Name() string {
	return "HelloWord"
}

func (a HelloWorld) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	//Business Logic
	return io, &florest_constants.AppError{Code: errors.FunctionalityNotImplementedErrorCode, Message: "invalid request"}
}
