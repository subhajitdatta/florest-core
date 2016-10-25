package hellomysql

import (
	"fmt"

	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/common/ratelimiter"
	"github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/utils/healthcheck"
	"github.com/jabong/florest-core/src/core/common/versionmanager"
)

type HelloMySQLAPI struct {
}

func (a *HelloMySQLAPI) GetVersion() versionmanager.Version {
	return versionmanager.Version{
		Resource: "MYSQL",
		Version:  "V1",
		Action:   "GET",
		BucketID: constants.OrchestratorBucketDefaultValue, //todo - should it be a constant
		Path:     "",                                       // path can be entered in the form buckets/{bucketId}/keys/{keyId}
	}
}

func (a *HelloMySQLAPI) GetOrchestrator() orchestrator.Orchestrator {
	logger.Info("Hello MySql Pipeline Creation begin")

	helloMySqlOrchestrator := new(orchestrator.Orchestrator)
	helloMySqlWorkflow := new(orchestrator.WorkFlowDefinition)
	helloMySqlWorkflow.Create()

	//Creation of the nodes in the workflow definition
	helloMySqlNode := new(mysqlNode)
	helloMySqlNode.SetID("hello mongo node 1")
	eerr := helloMySqlWorkflow.AddExecutionNode(helloMySqlNode)
	if eerr != nil {
		logger.Error(fmt.Sprintln(eerr))
	}

	//Set start node for the search workflow
	helloMySqlWorkflow.SetStartNode(helloMySqlNode)

	//Assign the workflow definition to the Orchestrator
	helloMySqlOrchestrator.Create(helloMySqlWorkflow)

	logger.Info(helloMySqlOrchestrator.String())
	logger.Info("Hello MySql Pipeline Created")
	return *helloMySqlOrchestrator
}

func (a *HelloMySQLAPI) GetHealthCheck() healthcheck.HCInterface {
	return new(HelloMySQLHealthCheck)
}

func (a *HelloMySQLAPI) Init() {
	//api initialization should come here
}

func (a *HelloMySQLAPI) GetRateLimiter() ratelimiter.RateLimiter {
	return nil
}
