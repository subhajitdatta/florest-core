package api

type TestHealthCheck struct {
	S string
}

func (n TestHealthCheck) GetName() string {
	return "test"
}

func (n TestHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
	}
}
