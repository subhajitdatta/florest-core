package hellorediscluster

import (
	"fmt"

	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/common/ratelimiter"
	"github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/utils/healthcheck"
	"github.com/jabong/florest-core/src/core/common/versionmanager"
)

type HelloRedisClusterAPI struct {
}

func (a *HelloRedisClusterAPI) GetVersion() versionmanager.Version {
	return versionmanager.Version{
		Resource: "REDISCLUSTER",
		Version:  "V1",
		Action:   "GET",
		BucketID: constants.OrchestratorBucketDefaultValue, //todo - should it be a constant
		Path:     "",                                       // path can be entered in the form buckets/{bucketId}/keys/{keyId}
	}
}

func (a *HelloRedisClusterAPI) GetOrchestrator() orchestrator.Orchestrator {
	logger.Info("Hello Redis Cluster Pipeline Creation begin")

	helloRedisClusterOrchestrator := new(orchestrator.Orchestrator)
	helloRedisClusterWorkflow := new(orchestrator.WorkFlowDefinition)
	helloRedisClusterWorkflow.Create()

	//Creation of the nodes in the workflow definition
	rcNode := new(redisClusterNode)
	rcNode.SetID("hello Redis Cluster node 1")
	eerr := helloRedisClusterWorkflow.AddExecutionNode(rcNode)
	if eerr != nil {
		logger.Error(fmt.Sprintln(eerr))
	}

	//Set start node for the search workflow
	helloRedisClusterWorkflow.SetStartNode(rcNode)

	//Assign the workflow definition to the Orchestrator
	helloRedisClusterOrchestrator.Create(helloRedisClusterWorkflow)

	logger.Info(helloRedisClusterOrchestrator.String())
	logger.Info("Hello Redis Cluster Pipeline Created")
	return *helloRedisClusterOrchestrator
}

func (a *HelloRedisClusterAPI) GetHealthCheck() healthcheck.HCInterface {
	return new(HelloRedisClusterHealthCheck)
}

func (a *HelloRedisClusterAPI) Init() {
	//api initialization should come here
}

func (a *HelloRedisClusterAPI) GetRateLimiter() ratelimiter.RateLimiter {
	return nil
}
