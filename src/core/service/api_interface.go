package service

import (
	"github.com/jabong/florest-core/src/common/ratelimiter"
	"github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/utils/healthcheck"
	"github.com/jabong/florest-core/src/core/common/versionmanager"
)

type APIInterface interface {
	GetVersion() versionmanager.Version

	GetOrchestrator() orchestrator.Orchestrator

	GetHealthCheck() healthcheck.HCInterface

	GetRateLimiter() ratelimiter.RateLimiter

	Init()
}
