package mongodb

import ()

// mongodbb interface
type MDBInterface interface {
	// init initializes the mongodb instance
	init(*MDBConfig) *MDBError
	// FindOne returns first matching item
	FindOne(string, map[string]interface{}) (interface{}, *MDBError)
	// FindAll returns all matching items
	FindAll(string, map[string]interface{}) ([]interface{}, *MDBError)
	// Insert add one item
	Insert(string, interface{}) *MDBError
	// Update modify existing item
	Update(string, map[string]interface{}, interface{}) *MDBError
	// Remove delete existing item
	Remove(string, map[string]interface{}) *MDBError
}
