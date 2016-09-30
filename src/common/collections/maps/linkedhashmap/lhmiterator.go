package linkedhashmap

import (
	"github.com/jabong/florest-core/src/common/collections"
)

// Iterator - A stateful iterator for linked hash map
type Iterator struct {
	m       *Map
	current *Link
}

// HasNext method moves the iterator to the next element and returns true if there was a next
// element in the map.
func (iterator *Iterator) HasNext() bool {
	if iterator.current == nil {
		return false
	}
	return true
}

// Next method returns the next element entry if it exists
func (iterator *Iterator) Next() *collections.Entry {
	temp := iterator.current
	if temp == nil {
		return nil
	}
	iterator.current = temp.next
	return collections.NewEntry(temp.key, temp.value)
}

// Reset method resets the iterator to its initial state
// Call Next() to fetch the first element if any.
func (iterator *Iterator) Reset() {
	iterator.current = iterator.m.head
}
