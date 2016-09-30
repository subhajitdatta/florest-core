package comparators

//UInt32Comparator for comparing the uint32 values
type UInt32Comparator struct {
}

//NewUInt32Comparator returns the new uint32 comparator
func NewUInt32Comparator() *UInt32Comparator {
	return &UInt32Comparator{}
}

// Compare two uint32 values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *UInt32Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(uint32)
	bAsserted := b.(uint32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
