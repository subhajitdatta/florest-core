package service

import (
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
)

type RequestStructs struct {
	Headers interface{}
	Params  interface{}
	Body    interface{}
}

type ValidateRequest interface {
	GetStructs(data workflow.WorkFlowData) RequestStructs
}
