package hystrix

// HCommandConf specifies the hystrix settings for each command
type HCommandConf struct {
	// Timeout is how long to wait for command to complete, in milliseconds. If nothing specified it takes
	// value of 1000
	Timeout int

	// MaxConcurrentRequests is how many commands of the same type can run at the same time. If nothing specified it takes
	// value of 10
	MaxConcurrentRequests int

	// RequestVolumeThreshold is the minimum number of requests needed before a circuit can be tripped due to health. If nothing specified it takes
	// value of 20
	RequestVolumeThreshold int

	// SleepWindow is how long, in milliseconds, to wait after a circuit opens before testing for recovery. If nothing specified it takes
	// value of 5000
	SleepWindow int

	// ErrorPercentThreshold causes circuits to open once the rolling measure of errors exceeds this percent of requests. If nothing specified it takes
	// value of 50
	ErrorPercentThreshold int
}
