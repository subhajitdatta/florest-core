package hello

type HelloWorldHealthCheck struct {
}

func (n HelloWorldHealthCheck) GetName() string {
	return "hello world"
}

func (n HelloWorldHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
	}
}
