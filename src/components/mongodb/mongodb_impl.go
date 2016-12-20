package mongodb

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// mongodb driver
type MongoDriver struct {
	conn    *mgo.Database
	session *mgo.Session
	conf    *MDBConfig
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

func (obj *MongoDriver) Copy(safe *Safe) *MSession {
	ms := new(MSession)
	ms.mgoSession = obj.session.Copy()
	if safe != nil {
		ms.mgoSession.SetSafe(obj.getMSafeFrom(safe))
	}
	return ms
}

func (obj *MongoDriver) getMSafeFrom(safe *Safe) *mgo.Safe {
	mgoSafe := new(mgo.Safe)
	mgoSafe.FSync = safe.FSync
	mgoSafe.J = safe.J
	mgoSafe.W = safe.W
	mgoSafe.WMode = safe.WMode
	mgoSafe.WTimeout = safe.WTimeout
	return mgoSafe
}

// Close shuts down the current session.
func (obj *MongoDriver) Close(session *MSession) {
	if session == nil {
		return
	}
	session.mgoSession.Close()
}

// CloseMasterSession closes the master session
func (obj *MongoDriver) CloseMasterSession() {
	obj.session.Close()
}

// FindOne queries the mongo DB using the session and returns only single result/collection
func (obj *MongoDriver) FindOneUsingSession(session *MSession, collection string, query map[string]interface{}) (ret interface{}, aerr *MDBError) {
	sess := session.mgoSession
	return obj.findOne(sess.DB(obj.conf.DbName).C(collection), bson.M(query))
}

// FindOne queries the mongo DB and returns only single result/collection
func (obj *MongoDriver) FindOne(collection string, query map[string]interface{}) (ret interface{}, aerr *MDBError) {
	obj.session.Refresh()
	return obj.findOne(obj.conn.C(collection), bson.M(query))
}

func (obj *MongoDriver) findOne(collection *mgo.Collection, query bson.M) (ret interface{}, aerr *MDBError) {
	if err := collection.Find(query).One(&ret); err != nil {
		return nil, getErrObj(ErrFindOneFailure, err.Error())
	}
	return ret, nil
}

// FindAll queries the mongo DB using session and returns all the results
func (obj *MongoDriver) FindAllUsingSession(session *MSession, collection string, query map[string]interface{}) (ret []interface{}, aerr *MDBError) {
	sess := session.mgoSession
	return obj.findAll(sess.DB(obj.conf.DbName).C(collection), bson.M(query))
}

// FindAll queries the mongo DB and returns all the results
func (obj *MongoDriver) FindAll(collection string, query map[string]interface{}) (ret []interface{}, aerr *MDBError) {
	obj.session.Refresh()
	return obj.findAll(obj.conn.C(collection), bson.M(query))
}

func (obj *MongoDriver) findAll(collection *mgo.Collection, query bson.M) (ret []interface{}, aerr *MDBError) {
	if err := collection.Find(query).All(&ret); err != nil {
		return nil, getErrObj(ErrFindAllFailure, err.Error())
	}
	return ret, nil
}

// Insert
func (obj *MongoDriver) InsertUsingSession(session *MSession, collection string, value interface{}) *MDBError {
	sess := session.mgoSession
	return obj.insert(sess.DB(obj.conf.DbName).C(collection), value)
}

func (obj *MongoDriver) Insert(collection string, value interface{}) *MDBError {
	obj.session.Refresh()
	return obj.insert(obj.conn.C(collection), value)
}

func (obj *MongoDriver) insert(collection *mgo.Collection, value interface{}) *MDBError {
	if err := collection.Insert(value); err != nil {
		return getErrObj(ErrInsertFailure, err.Error())
	}
	return nil
}

// Update updates the mongo DB collection passed as an argument
func (obj *MongoDriver) UpdateUsingSession(session *MSession, collection string, query map[string]interface{}, value interface{}) *MDBError {
	sess := session.mgoSession
	return obj.update(sess.DB(obj.conf.DbName).C(collection), bson.M(query), value)
}

// Update updates the mongo DB collection passed as an argument
func (obj *MongoDriver) Update(collection string, query map[string]interface{}, value interface{}) *MDBError {
	obj.session.Refresh()
	return obj.update(obj.conn.C(collection), bson.M(query), value)
}

func (obj *MongoDriver) update(collection *mgo.Collection, query bson.M, value interface{}) *MDBError {
	if err := collection.Update(query, value); err != nil {
		return getErrObj(ErrUpdateFailure, err.Error())
	}
	return nil
}

// Upsert updates/inserts the mongo DB collection passed as an argument
func (obj *MongoDriver) Upsert(collection string, query map[string]interface{}, value interface{}) *MDBError {
	obj.session.Refresh()
	return obj.upsert(obj.conn.C(collection), bson.M(query), value)
}

// UpsertUsingSession updates/inserts the mongo DB collection passed as an argument using the session passed as argument
func (obj *MongoDriver) UpsertUsingSession(session *MSession, collection string, query map[string]interface{}, value interface{}) *MDBError {
	sess := session.mgoSession
	return obj.upsert(sess.DB(obj.conf.DbName).C(collection), bson.M(query), value)
}

func (obj *MongoDriver) upsert(collection *mgo.Collection, query bson.M, value interface{}) *MDBError {
	if _, err := collection.Upsert(query, value); err != nil {
		return getErrObj(ErrUpdateFailure, err.Error())
	}
	return nil
}

// Remove deletes the documents using session from the collection passed in the argument
func (obj *MongoDriver) RemoveUsingSession(session *MSession, collection string, query map[string]interface{}) *MDBError {
	sess := session.mgoSession
	return obj.remove(sess.DB(obj.conf.DbName).C(collection), bson.M(query))
}

// Remove deletes the documents from the collection passed in the argument
func (obj *MongoDriver) Remove(collection string, query map[string]interface{}) *MDBError {
	obj.session.Refresh()
	return obj.remove(obj.conn.C(collection), bson.M(query))
}

func (obj *MongoDriver) remove(collection *mgo.Collection, query bson.M) *MDBError {
	if err := collection.Remove(query); err != nil {
		return getErrObj(ErrRemoveFailure, err.Error())
	}
	return nil
}
