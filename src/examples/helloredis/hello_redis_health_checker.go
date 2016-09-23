package helloredis

type HelloRedisHealthCheck struct {
}

func (n HelloRedisHealthCheck) GetName() string {
	return "hello world"
}

func (n HelloRedisHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
	}
}
