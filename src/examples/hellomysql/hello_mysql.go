package hellomysql

import (
	"fmt"

	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/components/sqldb"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
)

type mysqlNode struct {
	id string
}

func (a *mysqlNode) SetID(id string) {
	a.id = id
}

func (a mysqlNode) GetID() (id string, err error) {
	return a.id, nil
}

func (a mysqlNode) Name() string {
	return "mysqlNode"
}

func (a mysqlNode) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	// get db object
	db, errG := sqldb.Get("mysdb") // It should be called only once and can be shared across go routines
	defer func() {
		if db == nil {
			return
		}
		if errC := db.Close(); errC != nil {
			logger.Error(fmt.Sprintf("Failed to close DB Connection - %v", errC))
		}
	}()
	if errG != nil {
		msg := fmt.Sprintf("Failed to get Mysql instance %v", errG)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	// got db object, try methods

	if errP := db.Ping(); errP != nil { // try pinging db
		msg := fmt.Sprintf("Ping to DB Failed %v", errP)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	// execute a statement: create one table
	if _, errC := db.Execute("create table florest_employee (name varchar(255))"); errC != nil {
		msg := fmt.Sprintf("Create Table Failed - %v", errC)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	// raw query: query on table
	rows, qerr := db.Query("SELECT * from florest_employee")
	if qerr != nil {
		msg := fmt.Sprintf("Select from florest_employee test table Failed - %v", qerr)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	rows.Close()

	name := "rajcomics"
	// query with run-time arguments: query on table
	rows, qerrS := db.Query("SELECT * from florest_employee where name=?", &name)
	if qerrS != nil {
		msg := fmt.Sprintf("Select from florest_employee with where clause Failed - %v", qerrS)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	rows.Close()

	// start and commit one txn: insert one row in table
	txObj, terr := db.GetTxnObj()
	if terr != nil {
		msg := fmt.Sprintf("Failed to get Txn Object - %v", terr)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	txObj.Exec("insert into florest_employee (name) values('abc')")
	if cerr := txObj.Commit(); cerr != nil {
		msg := fmt.Sprintf("Insert in DB Failed - %v", cerr)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	// start and rollback one txn
	txObj, terr = db.GetTxnObj()
	if terr != nil {
		msg := fmt.Sprintf("Failed to get Txn Object for testing rollback case - %v", terr)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	txObj.Exec("select * from florest_employee")
	txObj.Rollback()

	// delete the created table
	if _, err := db.Execute("drop table florest_employee"); err != nil {
		msg := fmt.Sprintf("Failed to delete florest_employee - %v", terr)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	//Business Logic
	io.IOData.Set(constants.Result, "Mysql Connection and operation is successful")
	return io, nil
}
