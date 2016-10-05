package mongodb

import (
	"testing"
)

func TestMongodb(t *testing.T) {
	conf := new(MDBConfig)

	// fill invalid url
	conf.URL = "invalid"
	Set("invalid", conf, new(MongoDriver))
	if _, err := Get("invalid"); err == nil {
		t.Fatal("invalid url must throw error")
	}

	// Test methods: valid url,db and collection
	conf.URL = "mongodb://localhost:27017"
	conf.DbName = "flashback"
	Set("mymongo", conf, new(MongoDriver))
	if _, err := Get("mymongo"); err != nil {
		t.Fatalf("Failed to get config %v", err)
	}
}
