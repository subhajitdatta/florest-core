package collections

// Collection is base interface that all data structures implement.
type Collection interface {
	// IsEmpty returns true if the collection does not contain any elements
	IsEmpty() bool
	// Size returns number of elements in the collection.
	Size() int
	// Clear removes all elements from the collection.
	Clear()
	// Values returns all values either in
	// random order or
	// insertion order if the collection is iterable or
	// Sorted order if the collection is comparable
	Values() []interface{}
	// Contains returns true if the given keys are found in the collection
	Contains(keys ...interface{}) bool
}
