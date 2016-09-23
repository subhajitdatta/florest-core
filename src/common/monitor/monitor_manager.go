package monitor

import (
	"fmt"
	"log"
)

// Get gets a monitor implementation as specified in the platform in conf
func Get(conf *MConf) (MInterface, error) {
	switch conf.Platform {
	case DatadogAgent:
		datadogAgentClient, err := newDogstatsdAgent(conf)
		if err != nil {
			return nil, err
		}
		return datadogAgentClient, nil
	}
	return nil, fmt.Errorf("Unknown Monitor Type %s requested", conf.Platform)
}

// logMsg logs a log message prefixed with msgType if the verbose option is enabled in conf
func logMsg(msgType string, msg interface{}, conf *MConf) {
	if !conf.Verbose {
		return
	}
	log.Printf("\n%s: %v", msgType, msg)
}

// recoverFromPanic recovers from a panic from the surrounding function
// , creates an error from the panic and returns it
func recoverFromPanic(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("%s", r)
	}
}
