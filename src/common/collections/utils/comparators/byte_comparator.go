package comparators

//ByteComparator for comparing the byte values
type ByteComparator struct {
}

// NewByteComparator returns the new byte comparator
func NewByteComparator() *ByteComparator {
	return &ByteComparator{}
}

// Compare two byte values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *ByteComparator) Compare(a, b interface{}) int {
	aAsserted := a.(byte)
	bAsserted := b.(byte)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
