package sqldb

import (
	"testing"
)

func TestSqlDb(t *testing.T) {
	conf := new(SDBConfig)
	// fill invalid driver name
	if _, err := Get("invalid"); err == nil {
		t.Fatal("invalid driver must throw error")
	}

	// fill valid driver, invalid db details
	conf.DriverName = MYSQL // set driver name as mysql
	conf.Username = "root"
	conf.Password = ""
	conf.Host = "localhost"
	conf.Port = "3306"
	conf.Dbname = "invalid"
	conf.Timezone = "Local"
	conf.MaxOpenCon = 2
	conf.MaxIdleCon = 1
	Set("mysdb", conf, new(MysqlDriver))
	_, errC := Get("mysdb")
	if errC == nil {
		t.Fatal("Failed to get myql config")
	}
}
