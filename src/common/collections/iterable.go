package collections

// Iterable defines the elements in the collection are iterable.
// To be implemented by all ordered collections.
type Iterable interface {
	// GetIterator returns the stateful iterator defined for the collection
	Iterator() *Iterator
}
