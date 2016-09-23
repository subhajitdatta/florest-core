package mongodb

// Get() - Creates, initializes and returns the mongo instance based on the given config
func Get(conf *MDBConfig) (MDBInterface, *MDBError) {
	ret := new(mongoDriver)
	if err := ret.init(conf); err != nil {
		return nil, err
	}
	return ret, nil
}
