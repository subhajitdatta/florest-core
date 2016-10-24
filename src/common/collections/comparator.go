package collections

// Comparator is used for comparing the elements of the collection.
// All sorted collections should have a comparator.
type Comparator interface {
	// Compare compares the elements and returns
	// 0 if they are equal
	// -1 if a < b
	// 1 if a > b
	Compare(a, b interface{}) int
}
