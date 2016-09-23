package monitor

import ()

type MInterface interface {

	// Init initialised Monitor Client with the supplied config
	Init(conf *MConf) error

	// Info records a info type event with title, text and associated tags if any
	Info(data *MData) error

	// Success records a success type event with title, text and associated tags if any
	Success(data *MData) error

	// Warning records a warning type event with title, text and associated tags if any
	Warning(data *MData) error

	// Error records a error type event with title, text and associated tags if any
	Error(data *MData) error

	// Gauge submits a gauge metric to monitoring server with name, value and associated tags.
	// gauge metric measure the value of a particular thing over time. rate denotes the
	// sample rate. If rate is set to 1 that means send metric always.
	Gauge(name string, value float64, tags []string, rate float64) error

	// Count submits a count metric to a monitoring server with name, increment value and associated
	// tags. count metric basically denotes a counter. rate denotes the sample rate. If rate is set
	// to 1 that means send metric always.
	Count(name string, value int64, tags []string, rate float64) error

	// Histogram submits a metric to a monitoring server with name, value and associated
	// tags to create a histogram. rate denotes the sample rate. If rate is set to 1 that means send
	// metric always.
	Histogram(name string, value float64, tags []string, rate float64) error

	// Set  submits a set metric to a monitoring server with name, value and associated tags
	// Sets are used to count the number of unique elements in a group. e.g- to track the number
	// of unique visitor to your site. rate denotes the sample rate. If rate is set to 1 that means send
	// metric always.
	Set(name string, value string, tags []string, rate float64) error

	// SendAppMetrics can be used to send some other custom metrics other than Gauge, Count
	// Histogram & Set exposed as a http endpoint at serverAddr
	SendAppMetrics(serverIP string) error
}
