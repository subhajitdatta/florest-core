package sqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

type MysqlDriver struct {
	db *sql.DB
}

//init intializes and create a mysql connection
func (obj *MysqlDriver) Init(conf *SDBConfig) (aerr *SDBError) {
	var err error
	// open connection
	obj.db, err = sql.Open(MYSQL, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Dbname,
		url.QueryEscape(conf.Timezone),
	))
	if err == nil {
		// set max open
		obj.db.SetMaxOpenConns(conf.MaxOpenCon)
		// set max idle
		obj.db.SetMaxIdleConns(conf.MaxIdleCon)
		// try a ping
		err = obj.db.Ping()
	}
	// set error if needed
	if err != nil {
		aerr = getErrObj(ErrInitialization, err.Error())
	}
	// return
	return aerr
}

//Query executes the query on mysql DB and returns the pointer to rows
func (obj *MysqlDriver) Query(query string, args ...interface{}) (*sql.Rows, *SDBError) {
	rows, err := obj.db.Query(query, args...)
	if err != nil {
		return nil, getErrObj(ErrQueryFailure, err.Error())
	}
	return rows, nil
}

//Execute executes the query and returns the pointer to the sql Result
func (obj *MysqlDriver) Execute(query string, args ...interface{}) (sql.Result, *SDBError) {
	res, err := obj.db.Exec(query, args...)
	if err != nil {
		return nil, getErrObj(ErrExecuteFailure, err.Error())
	}
	return res, nil
}

func (obj *MysqlDriver) GetTxnObj() (*sql.Tx, *SDBError) {
	txn, err := obj.db.Begin()
	if err != nil {
		return nil, getErrObj(ErrGetTxnFailure, err.Error())
	}
	return txn, nil
}

func (obj *MysqlDriver) Ping() *SDBError {
	err := obj.db.Ping()
	if err != nil {
		return getErrObj(ErrPingFailure, err.Error())
	}
	return nil
}

func (obj *MysqlDriver) Close() *SDBError {
	err := obj.db.Close()
	if err != nil {
		return getErrObj(ErrCloseFailure, err.Error())
	}
	return nil
}
