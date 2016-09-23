package cache

//Item repesents the structure of an item to be stored in some cache data store
type Item struct {
	Key   string
	Value interface{}
	Error string
}

// CInterface enforces standard cache methods to be implemented
type CInterface interface {
	// Init initialises a connection to a cache server where all keys are to be namespaced by keyPrefix
	//	and the dump of all key-val pair in the cache to be stored in dumpFilePath (in case Dump is called)
	Init(conf *Config) error

	// Get gets an item from a cache store indexed with key. serialize and compress indicates if the cache implementation
	// has to undergo some serialization or compression before returning the item
	Get(key string, serialize bool, compress bool) (item *Item, err error)

	// Set sets an item into a cache store. serialise and compress indicates if the cache implementation
	// has to undergo some serialization or compression before setting the item in cache
	Set(item Item, serialize bool, compress bool) error

	// SetWithTimeout sets the item into cache, same as Set but this function takes an extra add_argument
	// which sets the timeout for the particular item, does not take expirySec from config
	SetWithTimeout(item Item, serialize bool, compress bool, ttl int32) error

	// Delete deletes a key from cache
	Delete(key string) error

	// DeleteBatch deletes an array of keys from cache
	DeleteBatch(keys []string) error

	// GetBatch gets a list of all items indexed with keys. serialize and compress indicates if the
	// cache implementation has to undergo some serialization or compression before returning the items
	GetBatch(keys []string, serialize bool, compress bool) (items map[string]*Item, err error)
}
