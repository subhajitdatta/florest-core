package mongodb

import mgo "gopkg.in/mgo.v2"

type MSession struct {
	mgoSession *mgo.Session
}

type Safe struct {
	// Min # of servers to ack before success
	W int

	// Write mode for MongoDB 2.0+ (e.g. "majority")
	WMode string

	// Milliseconds to wait for W before timing out
	WTimeout int

	// Should servers sync to disk before returning success
	FSync bool

	// Wait for next group commit if journaling; no effect otherwise
	J bool
}
