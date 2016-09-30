package treemap

import (
	"github.com/jabong/florest-core/src/common/collections"
)

// Iterator - A stateful iterator for tree map
type Iterator struct {
	rbIterator collections.Iterator
}

// HasNext method moves the iterator to the next element and returns true if there was a next
// element in the map.
func (iterator *Iterator) HasNext() bool {
	return iterator.rbIterator.HasNext()
}

// Next method returns the next element entry if it exists
func (iterator *Iterator) Next() *collections.Entry {
	return iterator.rbIterator.Next()
}

// Reset method resets the iterator to its initial state
// Call Next() to fetch the first element if any.
func (iterator *Iterator) Reset() {
	iterator.rbIterator.Reset()
}
