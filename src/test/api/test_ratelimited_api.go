package api

import (
	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/ratelimiter"
	"github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/utils/healthcheck"
	"github.com/jabong/florest-core/src/core/common/versionmanager"
)

type TestRateLimitedAPI struct {
	v versionmanager.Version
}

func (a *TestRateLimitedAPI) GetVersion() versionmanager.Version {
	return a.v
}

func (a *TestRateLimitedAPI) SetVersion(action string, version string, resource string, path string) {
	a.v = versionmanager.Version{
		Resource: resource,
		Version:  version,
		Action:   action,
		BucketID: constants.OrchestratorBucketDefaultValue, //todo - should it be a constant
		Path:     path,
	}
}

func (a *TestRateLimitedAPI) GetOrchestrator() orchestrator.Orchestrator {
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

func (a *TestRateLimitedAPI) GetHealthCheck() healthcheck.HCInterface {
	h := new(TestHealthCheck)
	h.S = "Hello World"
	return h
}

func (a *TestRateLimitedAPI) Init() {
	//api initialization should come here
}

func (a *TestRateLimitedAPI) GetRateLimiter() ratelimiter.RateLimiter {
	conf := ratelimiter.Config{Type: ratelimiter.GCRA, MaxRate: 1, MaxBurst: 0}
	rl, _ := ratelimiter.New(&conf)
	return rl
}
