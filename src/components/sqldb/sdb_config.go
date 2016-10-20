package sqldb

import ()

// Config struct contains all cofnig data necessary to connect to mysql DB
type SDBConfig struct {
	DriverName string
	Username   string
	Password   string
	Host       string
	Port       string
	Dbname     string
	Timezone   string
	MaxOpenCon int
	MaxIdleCon int
}
