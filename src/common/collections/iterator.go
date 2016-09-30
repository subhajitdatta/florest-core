package collections

// Iterator is stateful iterator for ordered collections
type Iterator interface {
	// HasNext method moves the iterator to the next element and returns true if there was a next
	// element in the collection.
	//
	// If HasNext() returns true, then next element's entry can be retrieved by Next().
	//
	// If Next() was called for the first time, then it will point the iterator to the first
	// element if it exists.
	//
	// Modifies the state of the iterator.
	HasNext() bool

	// Next method returns the next element entry if it exists
	Next() *Entry

	// Reset method resets the iterator to its initial state. Call Next() to fetch the first
	// element if any.
	Reset()
}
