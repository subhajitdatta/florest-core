package healthcheck

var healthCheckAPIList []HCInterface

//Initialise initialises an app monitor
func Initialise(apiList []HCInterface) {
	if healthCheckAPIList == nil {
		healthCheckAPIList = apiList
	}
}
