package mongodb

import (
	"testing"
)

func TestMongodb(t *testing.T) {
	conf := new(MDBConfig)

	// fill invalid url
	conf.URL = "invalid"
	if _, err := Get(conf); err == nil {
		t.Fatal("invalid url must throw error")
	}

	// Test methods: valid url,db and collection
	conf.URL = "mongodb://localhost:27017"
	conf.DbName = "flashback"
	if _, err := Get(conf); err != nil {
		t.Fatalf("Failed to get config %v", err)
	}
}
