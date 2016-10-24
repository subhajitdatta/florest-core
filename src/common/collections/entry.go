package collections

// Entry is a Key-Value pair. Key can either an
// index in case of list/set
// an element (interface{}) in case of maps
type Entry struct {
	key   interface{}
	value interface{}
}

// NewEntry instantiates a new immutable entry instance
func NewEntry(key interface{}, value interface{}) *Entry {
	return &Entry{key, value}
}

// GetKey returns the key associated with this entry
func (e *Entry) GetKey() interface{} {
	return e.key
}

// GetValue returns the key associated with this entry
func (e *Entry) GetValue() interface{} {
	return e.value
}
