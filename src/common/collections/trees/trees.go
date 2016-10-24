package trees

import "github.com/jabong/florest-core/src/common/collections"

// Tree interface that all trees implement
type Tree interface {
	// Put inserts an element into the tree.
	Put(key interface{}, value interface{})
	// Get searches the element in the tree by key and returns its value or nil if key doesn't exists.
	// Second return parameter is true if key was found, otherwise false.
	Get(key interface{}) (value interface{}, found bool)
	// Remove removes the element from the tree by key.
	Remove(key interface{})
	// Keys returns all keys of the element
	// random order or
	// insertion order if the map is iterable or
	// Sorted order if the map is comparable
	Keys() []interface{}

	// extends Collection interface
	collections.Collection
}
