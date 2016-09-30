package hashset

import (
	"fmt"
	"strings"
)

// Set holds elements in go's native map
type Set struct {
	items map[interface{}]struct{}
}

// An empty struct - Refer http://dave.cheney.net/2014/03/25/the-empty-struct
var itemExists = struct{}{}

// New instantiates a new empty set
func New() *Set {
	return &Set{items: make(map[interface{}]struct{})}
}

// Add adds one or more items to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.items[item] = itemExists
	}
}

// Remove removes one or more items from the set.
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains returns true if the given items are found in the set
func (set *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		_, found := set.items[item]
		if !found {
			return false
		}
	}
	return true
}

// IsEmpty returns true if the set does not contain any elements.
func (set *Set) IsEmpty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	set.items = make(map[interface{}]struct{})
}

// Values returns all items in the set (Random order).
func (set *Set) Values() []interface{} {
	values := make([]interface{}, set.Size())
	index := 0
	for item, _ := range set.items {
		values[index] = item
		index++
	}
	return values
}

// String returns a string representation of container
func (set *Set) String() string {
	str := "HashSet\n"
	for item, _ := range set.items {
		str += fmt.Sprintf("%v,", item)
	}
	return strings.TrimRight(str, ",")
}
