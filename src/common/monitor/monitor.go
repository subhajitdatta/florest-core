package monitor

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
