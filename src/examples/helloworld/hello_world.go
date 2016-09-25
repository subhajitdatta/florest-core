package helloworld

import (
	"fmt"

	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/utils/misc"
)

type helloNode struct {
	id string
}

func (h *helloNode) SetID(id string) {
	h.id = id
}

func (h helloNode) GetID() (id string, err error) {
	return h.id, nil
}

func (h helloNode) Name() string {
	return "HelloWord"
}

func (h helloNode) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	//Business Logic
	if _, err := misc.GetRequestFromIO(io); err != nil {
		logger.Error(fmt.Sprintf("Error in getting request from IO - %v", err))
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: "invalid request"}
	}
	io.IOData.Set(constants.Result, "Hello World")
	return io, nil
}
