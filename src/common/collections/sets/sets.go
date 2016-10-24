package sets

import "github.com/jabong/florest-core/src/common/collections"

// Set interface that all sets implement
type Set interface {
	// Add adds one or more items to the set.
	Add(items ...interface{})
	// Remove removes one or more items from the set.
	Remove(items ...interface{})

	// extends Collection interface
	collections.Collection
}
