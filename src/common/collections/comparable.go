package collections

// Comparable defines the elements in the collection are comparable.
// To be implemented by all sortable collections.
type Comparable interface {
	// GetComparator returns the comparator associated with the collection
	GetComparator() *Comparator
}
