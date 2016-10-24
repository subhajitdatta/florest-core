package comparators

//UInt64Comparator for comparing the uint64 values
type UInt64Comparator struct {
}

//NewUInt64Comparator returns the new uint64 comparator
func NewUInt64Comparator() *UInt64Comparator {
	return &UInt64Comparator{}
}

// Compare two uint64 values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *UInt64Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(uint64)
	bAsserted := b.(uint64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
