package service

import (
	"github.com/jabong/florest-core/src/common/constants"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
)

func getCustomMetricPrefix(data workflow.WorkFlowData) string {
	var monitorMetricPrefix string
	monitorCustomMetricPrefix, mcmpError := data.ExecContext.Get(constants.MonitorCustomMetric)
	if mcmpError == nil {
		monitorMetricPrefix, _ = monitorCustomMetricPrefix.(string)
	}
	return monitorMetricPrefix
}
