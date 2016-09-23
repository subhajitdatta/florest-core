package monitor

import ()

// MData is the logging/monitoring data that application has to sent

type MData struct {
	// Title denotes the title of a log message and event
	Title string

	// Body denotes the body of a log message and event
	Body string

	// Tags adds dimension to a event or log. This should be of the form <key:value> e.g env:prod
	Tags map[string]string

	// SentEvent controls whether the data has to be sent as event to some monitoring platform like datadog
	ToSendEvent bool
}
