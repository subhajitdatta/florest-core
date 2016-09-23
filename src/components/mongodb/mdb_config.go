package mongodb

import ()

// config from mongo db
type MDBConfig struct {
	URL    string // e.g. mongodb://localhost:27017
	DbName string
}
