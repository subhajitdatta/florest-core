package treeset

import (
	"github.com/jabong/florest-core/src/common/collections"
)

// Iterator - A stateful iterator for tree set
type Iterator struct {
	index      int
	rbIterator collections.Iterator
}

// HasNext method moves the iterator to the next element and returns true if there was a next
// element in the set.
func (iterator *Iterator) HasNext() bool {
	return iterator.rbIterator.HasNext()
}

// Next method returns the next element entry if it exists
func (iterator *Iterator) Next() *collections.Entry {
	temp := iterator.rbIterator.Next()
	if temp == nil {
		return nil
	}
	index := iterator.index
	iterator.index++
	return collections.NewEntry(index, temp.GetKey())
}

// Reset method resets the iterator to its initial state
// Call Next() to fetch the first element if any.
func (iterator *Iterator) Reset() {
	iterator.rbIterator.Reset()
	iterator.index = 0
}
