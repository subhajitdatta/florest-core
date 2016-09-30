package concurrenthashmap

import (
	"fmt"
	"strings"
	"sync"
)

// Map holds the elements in go's native map and uses RWMutex to
// guard concurrent access to it
type Map struct {
	items map[interface{}]interface{}
	// extends Read Write mutex, guards concurrent access to items.
	sync.RWMutex
}

// New instantiates a concurrent hash map
func New() *Map {
	return &Map{items: make(map[interface{}]interface{})}
}

// Put inserts an entry into the map.
func (m *Map) Put(key interface{}, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.items[key] = value
}

// PutIfAbsent inserts an entry into the map, if the key doesn't exists
func (m *Map) PutIfAbsent(key interface{}, value interface{}) {
	m.Lock()
	defer m.Unlock()
	_, found := m.items[key]
	if !found {
		m.items[key] = value
	}
}

// Get searches the element in the map by key and returns its value or nil if key doesn't exists.
// Second return parameter is true if key was found, otherwise false.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	m.RLock()
	defer m.RUnlock()
	value, found = m.items[key]
	return
}

// Remove removes the element from the map by key.
func (m *Map) Remove(key interface{}) {
	m.Lock()
	defer m.Unlock()
	delete(m.items, key)
}

// IsEmpty returns true if map does not contain any elements
func (m *Map) IsEmpty() bool {
	return m.Size() == 0
}

// Size returns number of el
// ements in the map.
func (m *Map) Size() int {
	m.RLock()
	defer m.RUnlock()
	size := len(m.items)
	return size
}

// Keys returns all keys of the map(random order).
func (m *Map) Keys() []interface{} {
	keys := make([]interface{}, m.Size())
	index := 0
	m.RLock()
	defer m.RUnlock()
	for key := range m.items {
		keys[index] = key
		index++
	}
	return keys
}

// Values returns all values of the map (random order).
func (m *Map) Values() []interface{} {
	values := make([]interface{}, m.Size())
	index := 0
	m.RLock()
	defer m.RUnlock()
	for _, value := range m.items {
		values[index] = value
		index++
	}
	return values
}

// Contains returns true if the given keys are found in the map
func (m *Map) Contains(keys ...interface{}) bool {
	m.RLock()
	defer m.RUnlock()
	for _, key := range keys {
		_, found := m.items[key]
		if !found {
			return false
		}
	}
	return true
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.Lock()
	defer m.Unlock()
	m.items = make(map[interface{}]interface{})
}

// String returns a string representation of container
func (m *Map) String() string {
	str := "ConcurrentHashMap\nMap["
	m.RLock()
	defer m.RUnlock()
	for key, value := range m.items {
		str += fmt.Sprintf("%v:%v ", key, value)
	}
	return strings.TrimRight(str, " ") + "]"
}
