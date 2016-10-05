package mongodb

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// mongodb driver
type MongoDriver struct {
	conn    *mgo.Database
	session *mgo.Session
}

// init method
func (obj *MongoDriver) Init(conf *MDBConfig) *MDBError {
	// set the connection
	tmp, err := mgo.Dial(conf.URL)
	if err != nil {
		return getErrObj(ErrInitialization, err.Error()+"-connection url:"+conf.URL)
	}
	obj.session = tmp
	obj.conn = tmp.DB(conf.DbName)
	return nil

}

// FindOne queries the mongo DB and returns only single result/collection
func (obj *MongoDriver) FindOne(collection string, query map[string]interface{}) (ret interface{}, aerr *MDBError) {
	obj.session.Refresh()
	err := obj.conn.C(collection).Find(bson.M(query)).One(&ret)
	if err != nil {
		return nil, getErrObj(ErrFindOneFailure, err.Error())
	}
	return ret, nil

}

// FindAll queries the mongo DB and returns all the results
func (obj *MongoDriver) FindAll(collection string, query map[string]interface{}) (ret []interface{}, aerr *MDBError) {
	obj.session.Refresh()
	err := obj.conn.C(collection).Find(bson.M(query)).All(&ret)
	if err != nil {
		return nil, getErrObj(ErrFindAllFailure, err.Error())
	}
	return ret, nil

}

func (obj *MongoDriver) Insert(collection string, value interface{}) *MDBError {
	obj.session.Refresh()
	err := obj.conn.C(collection).Insert(value)
	if err != nil {
		return getErrObj(ErrInsertFailure, err.Error())
	}
	return nil

}

// Update updates the mongo DB collection passed as an argument
func (obj *MongoDriver) Update(collection string, query map[string]interface{}, value interface{}) *MDBError {
	obj.session.Refresh()
	err := obj.conn.C(collection).Update(bson.M(query), value)
	if err != nil {
		return getErrObj(ErrUpdateFailure, err.Error())
	}
	return nil
}

// Remove deletes the documents from the collection passed in the argument
func (obj *MongoDriver) Remove(collection string, query map[string]interface{}) *MDBError {
	obj.session.Refresh()
	err := obj.conn.C(collection).Remove(bson.M(query))
	if err != nil {
		return getErrObj(ErrRemoveFailure, err.Error())
	}
	return nil
}
