package sqldb

import (
	"database/sql"
)

// SqlDbInterface is an interface for all sql db implementations. The functions of this interface have to be supported by allsql db implementations
type SDBInterface interface {
	// init initialize db instance
	Init(conf *SDBConfig) *SDBError
	// Query should be used for select purpose
	Query(string, ...interface{}) (*sql.Rows, *SDBError)
	// Execute should be used for data changes
	Execute(string, ...interface{}) (sql.Result, *SDBError)
	// Ping checks the connection
	Ping() *SDBError
	// Close close connection properly
	Close() *SDBError
	// GetTxnObj get transaction object
	GetTxnObj() (*sql.Tx, *SDBError)
}
