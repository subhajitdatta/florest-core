package mongodb

// mongodbb interface
type MDBInterface interface {
	// init initializes the mongodb instance
	Init(*MDBConfig) *MDBError

	// FindOne returns first matching item
	FindOne(string, map[string]interface{}) (interface{}, *MDBError)

	FindOneUsingSession(*MSession, string, map[string]interface{}) (interface{}, *MDBError)

	// FindAll returns all matching items
	FindAll(string, map[string]interface{}) ([]interface{}, *MDBError)

	FindAllUsingSession(*MSession, string, map[string]interface{}) ([]interface{}, *MDBError)

	// Insert add one item
	Insert(string, interface{}) *MDBError

	InsertUsingSession(*MSession, string, interface{}) *MDBError

	// Update modify existing item
	Update(string, map[string]interface{}, interface{}) *MDBError

	UpdateUsingSession(*MSession, string, map[string]interface{}, interface{}) *MDBError

	// Update modify existing item or insert new item if does not exist
	Upsert(string, map[string]interface{}, interface{}) *MDBError

	UpsertUsingSession(*MSession, string, map[string]interface{}, interface{}) *MDBError

	// Remove delete existing item
	Remove(string, map[string]interface{}) *MDBError

	RemoveUsingSession(*MSession, string, map[string]interface{}) *MDBError

	// Close shuts down the current session.
	Close(*MSession)

	// Copy creates a copy of master session
	Copy(safe *Safe) *MSession

	// CloseMasterSession closes the master session
	CloseMasterSession()
}
