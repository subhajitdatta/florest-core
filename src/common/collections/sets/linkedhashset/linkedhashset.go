package linkedhashset

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/collections"
	lhmap "github.com/jabong/florest-core/src/common/collections/maps/linkedhashmap"
	"strings"
)

// Set holds elements in linked hash map
type Set struct {
	items *lhmap.Map
}

// An empty struct - Refer http://dave.cheney.net/2014/03/25/the-empty-struct
var itemExists = struct{}{}

// New instantiates a new empty set
func New() *Set {
	return &Set{items: lhmap.New()}
}

// Add adds one or more items to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.items.Put(item, itemExists)
	}
}

// Remove removes one or more items from the set.
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		set.items.Remove(item)
	}
}

// Contains returns true if the given items are found in the set
func (set *Set) Contains(items ...interface{}) bool {
	return set.items.Contains(items...)
}

// IsEmpty returns true if the set does not contain any elements.
func (set *Set) IsEmpty() bool {
	return set.items.IsEmpty()
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	return set.items.Size()
}

// Values returns all items in the set (Insertion order).
func (set *Set) Values() []interface{} {
	return set.items.Keys()
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	set.items.Clear()
}

// Iterator returns a stateful iterator used for iterating over all the items of the set
func (set *Set) Iterator() collections.Iterator {
	return &Iterator{iterator: set.items.Iterator(), index: 0}
}

// String returns a string representation of container
func (set *Set) String() string {
	str := "LinkedHashSet\n"
	for _, v := range set.Values() {
		str += fmt.Sprintf("%v,", v)
	}
	return strings.TrimRight(str, ",")
}
