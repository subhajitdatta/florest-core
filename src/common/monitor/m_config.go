package monitor

import ()

type MConf struct {

	//APPName acts as a namespace which is prefixed with every custom metric
	APPName string `json:"AppName"`

	// APIKey is an optional parameter used for authenticating a write request in Datadog server.
	// This is needed only for the Http Apis ( which are not supported yet )
	APIKey string `json:"ApiKey"`

	// APPKey is an optional parameter used for authenticating a read request from Datadog server
	// along with ApiKey. This is needed only for the Http Apis ( which are not supported yet )
	APPKey string `json:"AppKey"`

	// AgentServer is the montoring server ip and port
	AgentServer string

	// Platform specifies monitoring platform that is being used. For now only
	// Datadog is supported in agent mode
	Platform string

	// Verbose option if set to true prints down some information for debugging purpose
	Verbose bool

	// Enabled determines whether to disable or enable sending metrics to a monitoring server
	Enabled bool

	// MetricsServer to send server stats, e.g mem usage, disk usage, etc.
	MetricsServer string
}
