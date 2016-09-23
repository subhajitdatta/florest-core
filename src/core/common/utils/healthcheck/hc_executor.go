package healthcheck

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
)

type HCExecutor struct {
	id string
}

func (n HCExecutor) Name() string {
	return "Health Check Executor"
}

func (n *HCExecutor) SetID(id string) {
	n.id = id
}

func (n HCExecutor) GetID() (id string, err error) {
	return n.id, nil
}

func (n HCExecutor) Execute(data workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	rc, _ := data.ExecContext.Get(constants.RequestContext)
	logger.Info(fmt.Sprintln("entered ", n.Name()), rc)

	if healthCheckAPIList == nil {
		return data, &constants.AppError{Code: constants.ResourceErrorCode, Message: "Health Chech Api not Initialized"}
	}

	var res = make(map[string]interface{})

	for _, apiResource := range healthCheckAPIList {
		res[apiResource.GetName()] = apiResource.GetHealth()
	}

	data.IOData.Set(constants.Result, res)

	logger.Info(fmt.Sprintln("exiting ", n.Name()), rc)
	return data, nil

}
