package treemap

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/collections"
	rbt "github.com/jabong/florest-core/src/common/collections/trees/rbtree"
	"strings"
)

// Map holds the elements in a red-black tree
type Map struct {
	tree *rbt.Tree
}

// NewWith instantiates a tree map with the custom comparator.
func NewWith(comparator collections.Comparator) *Map {
	return &Map{tree: rbt.New(comparator)}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map) Put(key interface{}, value interface{}) {
	m.tree.Put(key, value)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	return m.tree.Get(key)
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map) Remove(key interface{}) {
	m.tree.Remove(key)
}

// Keys returns all keys in-order
func (m *Map) Keys() []interface{} {
	return m.tree.Keys()
}

// First returns the first(min) entry in the map
func (m *Map) First() *collections.Entry {
	if node := m.tree.Left(); node != nil {
		return collections.NewEntry(node.Key, node.Value)
	}
	return nil
}

// Last returns the last(max) entry in the map
func (m *Map) Last() *collections.Entry {
	if node := m.tree.Right(); node != nil {
		return collections.NewEntry(node.Key, node.Value)
	}
	return nil
}

// IsEmpty returns true if the map does not contain any elements
func (m *Map) IsEmpty() bool {
	return m.tree.IsEmpty()
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return m.tree.Size()
}

// Values returns all values in-order based on the key.
func (m *Map) Values() []interface{} {
	return m.tree.Values()
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.tree.Clear()
}

// Returns true if the given keys are found in the map
func (m *Map) Contains(keys ...interface{}) bool {
	return m.tree.Contains(keys...)
}

// Iterator returns a stateful iterator used for iterating over all the entries of the map
func (m *Map) Iterator() collections.Iterator {
	return &Iterator{rbIterator: m.tree.Iterator()}
}

// GetComparator returns the comparator associated with this map
func (m *Map) GetComparator() collections.Comparator {
	return m.tree.GetComparator()
}

// String returns a string representation of container
func (m *Map) String() string {
	str := "TreeMap\nmap["
	it := m.Iterator()
	for it.HasNext() {
		entry := it.Next()
		str += fmt.Sprintf("%v:%v ", entry.GetKey(), entry.GetValue())
	}
	return strings.TrimRight(str, " ") + "]"
}
