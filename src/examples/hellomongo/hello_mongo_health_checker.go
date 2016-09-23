package hellomongo

type HelloMongoHealthCheck struct {
}

func (n HelloMongoHealthCheck) GetName() string {
	return "hello world"
}

func (n HelloMongoHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
	}
}
