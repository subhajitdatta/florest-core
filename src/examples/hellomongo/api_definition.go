package hellomongo

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/common/ratelimiter"
	"github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/utils/healthcheck"
	"github.com/jabong/florest-core/src/core/common/versionmanager"
)

type HelloMongoAPI struct {
}

func (a *HelloMongoAPI) GetVersion() versionmanager.Version {
	return versionmanager.Version{
		Resource: "MONGO",
		Version:  "V1",
		Action:   "GET",
		BucketID: constants.OrchestratorBucketDefaultValue, //todo - should it be a constant
		Path:     "",                                       // path can be entered in the form buckets/{bucketId}/keys/{keyId}
	}
}

func (a *HelloMongoAPI) GetOrchestrator() orchestrator.Orchestrator {
	logger.Error("Hello Mongo Pipeline Creation begin")

	helloMongoOrchestrator := new(orchestrator.Orchestrator)
	helloMongoWorkflow := new(orchestrator.WorkFlowDefinition)
	helloMongoWorkflow.Create()

	//Creation of the nodes in the workflow definition
	helloMongoNode := new(mongoNode)
	helloMongoNode.SetID("hello mongo node 1")
	eerr := helloMongoWorkflow.AddExecutionNode(helloMongoNode)
	if eerr != nil {
		logger.Error(fmt.Sprintln(eerr))
	}

	//Set start node for the search workflow
	helloMongoWorkflow.SetStartNode(helloMongoNode)

	//Assign the workflow definition to the Orchestrator
	helloMongoOrchestrator.Create(helloMongoWorkflow)

	logger.Info(helloMongoOrchestrator.String())
	logger.Info("Hello Mongo Pipeline Created")
	return *helloMongoOrchestrator
}

func (a *HelloMongoAPI) GetHealthCheck() healthcheck.HCInterface {
	return new(HelloMongoHealthCheck)
}

func (a *HelloMongoAPI) Init() {
	//api initialization should come here
}

func (a *HelloMongoAPI) GetRateLimiter() ratelimiter.RateLimiter {
	return nil
}
