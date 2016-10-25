package hello

import (
	"fmt"

	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/utils/healthcheck"
	"github.com/jabong/florest-core/src/core/common/versionmanager"
	"github.com/jabong/florest-core/src/common/ratelimiter"
)

type HelloAPI struct {
}

func (a *HelloAPI) GetVersion() versionmanager.Version {
	return versionmanager.Version{
		Resource: "HELLO",
		Version:  "V1",
		Action:   "GET",
		BucketID: constants.OrchestratorBucketDefaultValue, //todo - should it be a constant
		Path:     "",
	}
}

func (a *HelloAPI) GetOrchestrator() orchestrator.Orchestrator {
	logger.Info("Hello World Pipeline Creation begin")

	helloWorldOrchestrator := new(orchestrator.Orchestrator)
	helloWorldWorkflow := new(orchestrator.WorkFlowDefinition)
	helloWorldWorkflow.Create()

	//Creation of the nodes in the workflow definition
	helloWorldNode := new(HelloWorld)
	helloWorldNode.SetID("hello world node 1")
	eerr := helloWorldWorkflow.AddExecutionNode(helloWorldNode)
	if eerr != nil {
		logger.Error(fmt.Sprintln(eerr))
	}

	//Set start node for the search workflow
	helloWorldWorkflow.SetStartNode(helloWorldNode)

	//Assign the workflow definition to the Orchestrator
	helloWorldOrchestrator.Create(helloWorldWorkflow)

	logger.Info(helloWorldOrchestrator.String())
	logger.Info("Hello World Pipeline Created")
	logger.Info("Hello World Pipeline Created")
	return *helloWorldOrchestrator
}

func (a *HelloAPI) GetHealthCheck() healthcheck.HCInterface {
	return new(HelloWorldHealthCheck)
}

func (a *HelloAPI) Init() {
	//api initialization should come here
}

func (a *HelloAPI) GetRateLimiter() ratelimiter.RateLimiter {
	return nil
}
