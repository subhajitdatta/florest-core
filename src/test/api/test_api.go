package api

import (
	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/ratelimiter"
	"github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/utils/healthcheck"
	"github.com/jabong/florest-core/src/core/common/versionmanager"
)

type TestAPI struct {
	v versionmanager.Version
}

func (a *TestAPI) GetVersion() versionmanager.Version {
	return a.v
}

func (a *TestAPI) SetVersion(action string, version string, resource string, path string) {
	a.v = versionmanager.Version{
		Resource: resource,
		Version:  version,
		Action:   action,
		BucketID: constants.OrchestratorBucketDefaultValue, //todo - should it be a constant
		Path:     path,
	}
}

func (a *TestAPI) GetOrchestrator() orchestrator.Orchestrator {
	testOrchestrator := new(orchestrator.Orchestrator)
	testOrchestratorWorkflow := new(orchestrator.WorkFlowDefinition)
	testOrchestratorWorkflow.Create()

	testExecutor := new(testExecutor)
	testExecutor.SetID("1")
	testOrchestratorWorkflow.AddExecutionNode(testExecutor)

	testOrchestratorWorkflow.SetStartNode(testExecutor)
	testOrchestrator.Create(testOrchestratorWorkflow)

	return *testOrchestrator
}

func (a *TestAPI) GetHealthCheck() healthcheck.HCInterface {
	h := new(TestHealthCheck)
	h.S = "Hello World"
	return h
}

func (a *TestAPI) Init() {
	//api initialization should come here
}

func (a *TestAPI) GetRateLimiter() ratelimiter.RateLimiter {
	return nil
}
