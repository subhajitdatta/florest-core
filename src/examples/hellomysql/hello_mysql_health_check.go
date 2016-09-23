package hellomysql

type HelloMySQLHealthCheck struct {
}

func (n HelloMySQLHealthCheck) GetName() string {
	return "hello world"
}

func (n HelloMySQLHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
	}
}
