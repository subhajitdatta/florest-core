package treeset

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/collections"
	rbt "github.com/jabong/florest-core/src/common/collections/trees/rbtree"
	"strings"
)

// Set holds elements in a red-black tree
type Set struct {
	tree *rbt.Tree
}

// An empty struct - Refer http://dave.cheney.net/2014/03/25/the-empty-struct
var itemExists = struct{}{}

// NewWith instantiates a new empty set with the custom comparator.
func NewWith(comparator collections.Comparator) *Set {
	return &Set{tree: rbt.New(comparator)}
}

// Add adds one or more items to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.tree.Put(item, itemExists)
	}
}

// Remove removes one or more items from the set.
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		set.tree.Remove(item)
	}
}

// Contains returns true if the given items are found in the set
func (set *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		_, found := set.tree.Get(item)
		if !found {
			return false
		}
	}
	return true
}

// IsEmpty returns true if the set does not contain any elements.
func (set *Set) IsEmpty() bool {
	return set.tree.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	return set.tree.Size()
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	set.tree.Clear()
}

// Values returns all items in the set (In-order).
func (set *Set) Values() []interface{} {
	return set.tree.Keys()
}

// First returns the first(min) entry in the set
func (set *Set) First() interface{} {
	if node := set.tree.Left(); node != nil {
		return node.Key
	}
	return nil
}

// Last returns the last(max) element in the set
func (set *Set) Last() interface{} {
	if node := set.tree.Right(); node != nil {
		return node.Key
	}
	return nil
}

// Iterator returns a stateful iterator used for iterating over all the items of the set
func (set *Set) Iterator() collections.Iterator {
	return &Iterator{index: 0, rbIterator: set.tree.Iterator()}
}

// GetComparator returns the comparator associated with this set
func (set *Set) GetComparator() collections.Comparator {
	return set.tree.GetComparator()
}

// String returns a string representation of container
func (set *Set) String() string {
	str := "TreeSet\n"
	for _, v := range set.Values() {
		str += fmt.Sprintf("%v,", v)
	}
	return strings.TrimRight(str, ",")
}
