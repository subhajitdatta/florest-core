package linkedhashset

import (
	"github.com/jabong/florest-core/src/common/collections"
)

// Iterator - A stateful iterator for linked hash set
type Iterator struct {
	iterator collections.Iterator
	index    int
}

// HasNext method moves the iterator to the next element and returns true if there was a next
// element in the set.
func (iterator *Iterator) HasNext() bool {
	return iterator.iterator.HasNext()
}

// Next method returns the next element entry if it exists
func (iterator *Iterator) Next() *collections.Entry {
	temp := iterator.iterator.Next()
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
	iterator.iterator.Reset()
	iterator.index = 0
}
