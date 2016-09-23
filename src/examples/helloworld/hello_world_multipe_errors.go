package helloworld

import (
	"github.com/jabong/florest-core/src/common/constants"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
)

type HelloWorldMultiError struct {
	id string
}

func (h *HelloWorldMultiError) SetID(id string) {
	h.id = id
}

func (h HelloWorldMultiError) GetID() (id string, err error) {
	return h.id, nil
}

func (h HelloWorldMultiError) Name() string {
	return "HelloWord"
}

func (h HelloWorldMultiError) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	//Business Logic

	err := new(constants.AppErrors)
	err.Errors = append(err.Errors, constants.AppError{
		Code:             constants.ParamsInSufficientErrorCode,
		Message:          "In Sufficient Params Error Code",
		DeveloperMessage: "In Sufficient Params Error Code",
	})
	err.Errors = append(err.Errors, constants.AppError{
		Code:             constants.ResourceErrorCode,
		Message:          "Resource Error Code",
		DeveloperMessage: "Resource Error Code",
	})
	return io, err
}
