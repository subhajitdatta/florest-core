package helloredis

import (
	"fmt"

	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/common/ratelimiter"
	"github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/utils/healthcheck"
	"github.com/jabong/florest-core/src/core/common/versionmanager"
)

type HelloRedisAPI struct {
}

func (a *HelloRedisAPI) GetVersion() versionmanager.Version {
	return versionmanager.Version{
		Resource: "REDIS",
		Version:  "V1",
		Action:   "GET",
		BucketID: constants.OrchestratorBucketDefaultValue, //todo - should it be a constant
		Path:     "",                                       // path can be entered in the form buckets/{bucketId}/keys/{keyId}
	}
}

func (a *HelloRedisAPI) GetOrchestrator() orchestrator.Orchestrator {
	logger.Info("Hello Mongo Pipeline Creation begin")

	helloRedisOrchestrator := new(orchestrator.Orchestrator)
	helloRedisWorkflow := new(orchestrator.WorkFlowDefinition)
	helloRedisWorkflow.Create()

	//Creation of the nodes in the workflow definition
	redisNode := new(redisNode)
	redisNode.SetID("hello mongo node 1")
	eerr := helloRedisWorkflow.AddExecutionNode(redisNode)
	if eerr != nil {
		logger.Error(fmt.Sprintln(eerr))
	}

	//Set start node for the search workflow
	helloRedisWorkflow.SetStartNode(redisNode)

	//Assign the workflow definition to the Orchestrator
	helloRedisOrchestrator.Create(helloRedisWorkflow)

	logger.Info(helloRedisOrchestrator.String())
	logger.Info("Hello Redis Pipeline Created")
	return *helloRedisOrchestrator
}

func (a *HelloRedisAPI) GetHealthCheck() healthcheck.HCInterface {
	return new(HelloRedisHealthCheck)
}

func (a *HelloRedisAPI) Init() {
	//api initialization should come here
}

func (a *HelloRedisAPI) GetRateLimiter() ratelimiter.RateLimiter {
	return nil
}
