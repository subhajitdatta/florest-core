package hystrix

import (
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix/metric_collector"
)

const (
	dmCircuitOpen       = "hystrix.%s.circuitOpen"
	dmAttempts          = "hystrix.%s.attempts"
	dmErrors            = "hystrix.%s.errors"
	dmSuccesses         = "hystrix.%s.successes"
	dmFailures          = "hystrix.%s.failures"
	dmRejects           = "hystrix.%s.rejects"
	dmShortCircuits     = "hystrix.%s.shortCircuits"
	dmTimeouts          = "hystrix.%s.timeouts"
	dmFallbackSuccesses = "hystrix.%s.fallbackSuccesses"
	dmFallbackFailures  = "hystrix.%s.fallbackFailures"
	dmTotalDuration     = "hystrix.%s.totalDuration"
	dmRunDuration       = "hystrix.%s.runDuration"
)

type IDatadogClient interface {
	Gauge(name string, value float64, tags []string, rate float64) error
	Count(name string, value int64, tags []string, rate float64) error
	Histogram(name string, value float64, tags []string, rate float64) error
}

type HDatadogMetricsCollector struct {
	client      IDatadogClient
	circuitName string
}

// IncrementAttempts increments the number of calls to this circuit.
func (dc *HDatadogMetricsCollector) IncrementAttempts() {
	dc.client.Count(dc.getMetricWidCircuitName(dmAttempts), 1, nil, 1.0)
}

// IncrementErrors increments the number of unsuccessful attempts.
// Attempts minus Errors will equal successes within a time range.
// Errors are any result from an attempt that is not a success.
func (dc *HDatadogMetricsCollector) IncrementErrors() {
	dc.client.Count(dc.getMetricWidCircuitName(dmErrors), 1, nil, 1.0)
}

// IncrementSuccesses increments the number of requests that succeed.
func (dc *HDatadogMetricsCollector) IncrementSuccesses() {
	dc.client.Gauge(dc.getMetricWidCircuitName(dmCircuitOpen), 0, nil, 1.0)
	dc.client.Count(dc.getMetricWidCircuitName(dmSuccesses), 1, nil, 1.0)
}

// IncrementFailures increments the number of requests that fail.
func (dc *HDatadogMetricsCollector) IncrementFailures() {
	dc.client.Count(dc.getMetricWidCircuitName(dmFailures), 1, nil, 1.0)
}

// IncrementRejects increments the number of requests that are rejected.
func (dc *HDatadogMetricsCollector) IncrementRejects() {
	dc.client.Count(dc.getMetricWidCircuitName(dmRejects), 1, nil, 1.0)
}

// IncrementShortCircuits increments the number of requests that short circuited
// due to the circuit being open.
func (dc *HDatadogMetricsCollector) IncrementShortCircuits() {
	dc.client.Gauge(dc.getMetricWidCircuitName(dmCircuitOpen), 1, nil, 1.0)
	dc.client.Count(dc.getMetricWidCircuitName(dmShortCircuits), 1, nil, 1.0)
}

// IncrementTimeouts increments the number of timeouts that occurred in the
// circuit breaker.
func (dc *HDatadogMetricsCollector) IncrementTimeouts() {
	dc.client.Count(dc.getMetricWidCircuitName(dmTimeouts), 1, nil, 1.0)
}

// IncrementFallbackSuccesses increments the number of successes that occurred
// during the execution of the fallback function.
func (dc *HDatadogMetricsCollector) IncrementFallbackSuccesses() {
	dc.client.Count(dc.getMetricWidCircuitName(dmFallbackSuccesses), 1, nil, 1.0)
}

// IncrementFallbackFailures increments the number of failures that occurred
// during the execution of the fallback function.
func (dc *HDatadogMetricsCollector) IncrementFallbackFailures() {
	dc.client.Count(dc.getMetricWidCircuitName(dmFallbackFailures), 1, nil, 1.0)
}

// UpdateTotalDuration updates the internal counter of how long we've run for.
func (dc *HDatadogMetricsCollector) UpdateTotalDuration(timeSinceStart time.Duration) {
	ms := float64(timeSinceStart.Nanoseconds() / 1000000)
	dc.client.Histogram(dc.getMetricWidCircuitName(dmTotalDuration), ms, nil, 1.0)
}

// UpdateRunDuration updates the internal counter of how long the last run took.
func (dc *HDatadogMetricsCollector) UpdateRunDuration(runDuration time.Duration) {
	ms := float64(runDuration.Nanoseconds() / 1000000)
	dc.client.Histogram(dc.getMetricWidCircuitName(dmRunDuration), ms, nil, 1.0)
}

// Reset is a noop operation in this collector.
func (dc *HDatadogMetricsCollector) Reset() {}

// getMetricWidCircuitName adds circuitName to the metric submitted to datadog
func (dc *HDatadogMetricsCollector) getMetricWidCircuitName(metricName string) string {
	return fmt.Sprintf(metricName, dc.circuitName)
}

// Register registers datadog as the metrics collector
func RegisterDatadogCollector(client IDatadogClient) {
	metricCollector.Registry.Register(newDatadogMetricCollector(client))
}

func newDatadogMetricCollector(client IDatadogClient) func(string) metricCollector.MetricCollector {
	return func(name string) metricCollector.MetricCollector {

		return &HDatadogMetricsCollector{
			client:      client,
			circuitName: name,
		}
	}
}
