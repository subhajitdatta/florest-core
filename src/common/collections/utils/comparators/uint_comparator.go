package comparators

//UIntComparator for comparing the uint values
type UIntComparator struct {
}

//NewUIntComparator returns the new uint comparator
func NewUIntComparator() *UIntComparator {
	return &UIntComparator{}
}

// Compare two uint values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *UIntComparator) Compare(a, b interface{}) int {
	aAsserted := a.(uint)
	bAsserted := b.(uint)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
