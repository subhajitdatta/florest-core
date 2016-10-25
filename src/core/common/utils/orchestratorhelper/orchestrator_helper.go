package orchestratorhelper

import (
	"fmt"

	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/common/ratelimiter"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/versionmanager"
)

func GetOrchestrator(resource string,
	version string,
	action string,
	bucketID string, pathParams string) (*workflow.Orchestrator, *ratelimiter.RateLimiter, *map[string]string, error) {

	logFormatter := "GetOrchestrator ==== Resource: %v === Version: %v === Action: %v === BucketId: %v === PathParams: %v"
	logger.Info(fmt.Sprintf(logFormatter, resource, version, action, bucketID, pathParams))
	orchestratorVersion, ratelimiter, parameters, gerr := versionmanager.Get(resource, version, action, bucketID, pathParams)
	if gerr != nil {
		return nil, nil, nil, &constants.AppError{Code: constants.InvalidRequestURI, Message: gerr.Error()}
	}

	orchestrator, ok := orchestratorVersion.(workflow.Orchestrator)
	if !ok {
		return nil, nil, nil, &constants.AppError{Code: constants.ResourceErrorCode,
			Message: "Error retrieving orchestrator"}
	}

	return &orchestrator, ratelimiter, &parameters, nil

}

func ExecuteOrchestrator(input *workflow.WorkFlowData,
	orchestrator *workflow.Orchestrator) (interface{}, error) {

	output := orchestrator.Start(input)
	res, _ := output.IOData.Get(constants.Result)

	orchestratorStates := output.GetWorkflowState()
	var orchestratorError error
	for _, err := range orchestratorStates {
		if v, ok := err.(error); ok {
			orchestratorError = v
		}
	}

	rc, _ := input.ExecContext.Get(constants.RequestContext)
	logger.Info(fmt.Sprintf("%v", orchestratorError), rc)
	return res, orchestratorError
}
