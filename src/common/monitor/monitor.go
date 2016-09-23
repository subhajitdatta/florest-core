package monitor

import ()

// myMonitor contains an instance of implementation of MonitorInterface
type myMonitor struct {
	agent     MInterface
	isEnabled bool
}

// monitorObj singleton object
var monitorObj MInterface

// GetInstance returns a singleton instance of type MInterface
func GetInstance() MInterface {
	// return
	return monitorObj
}

// Initialize initialize the monitor object
func Initialize(cnfg *MConf) (err error) {
	monitorObj, err = Get(cnfg)
	return err
}
