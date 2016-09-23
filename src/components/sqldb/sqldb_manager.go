package sqldb

// Get() - Creates, initializes and returns the mysql instance based on the given config
func Get(conf *SDBConfig) (ret SDBInterface, err *SDBError) {
	if conf.DriverName == MYSQL {
		ret = new(mysqlDriver)
		err = ret.init(conf)
	} else {
		err = getErrObj(ErrNoDriver, conf.DriverName+" is not supported")
	}
	// return
	return ret, err
}
