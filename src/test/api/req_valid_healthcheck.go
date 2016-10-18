package api

type ReqValidHealthCheck struct {
	S string
}

func (n ReqValidHealthCheck) GetName() string {
	return "reqValid"
}

func (n ReqValidHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
	}
}
