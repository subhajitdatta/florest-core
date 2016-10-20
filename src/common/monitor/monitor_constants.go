package monitor

import ()

const (
	DatadogAgent string = "DatadogAgent"
)

const DefaultMonitor string = DatadogAgent

// Different types of log message tags if verbose option is enabled
const (
	errMsgTag   string = "ERROR"
	infoMsgtag  string = "INFO"
	warnMsgTag  string = "WARNING"
	traceMsgTag string = "TRACE"
)
