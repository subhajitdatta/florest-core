package sqldb

import (
	"testing"
)

func TestSqlDb(t *testing.T) {
	var dbObj SDBInterface
	conf := new(SDBConfig)
	// fill invalid driver name
	conf.DriverName = "invalid" // set driver name as mysql
	if _, err := Get(conf); err == nil {
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
	dbObj, errC := Get(conf)
	if errC == nil {
		t.Fatal("Failed to get myql config")
	}
	// As invalid db object, assert error for all methods
	if _, err := dbObj.Query("invalid query"); err == nil {
		t.Fatal("query must fail for this invalid db")
	}

	if _, err := dbObj.Execute("invalid execute"); err == nil {
		t.Fatal("execute must fail for this invalid db")
	}

	if _, err := dbObj.GetTxnObj(); err == nil {
		t.Fatal("get txn object must fail for this invalid db")
	}
	if err := dbObj.Close(); err != nil {
		t.Fatalf("close must fail for this invalid db %v", err)
	}
}
