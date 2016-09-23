package hellorediscluster

type HelloRedisClusterHealthCheck struct {
}

func (n HelloRedisClusterHealthCheck) GetName() string {
	return "hello world"
}

func (n HelloRedisClusterHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
	}
}
