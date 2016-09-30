package comparators

//UInt16Comparator for comparing the uint16 values
type UInt16Comparator struct {
}

//NewUInt16Comparator returns the new uint16 comparator
func NewUInt16Comparator() *UInt16Comparator {
	return &UInt16Comparator{}
}

// Compare two uint16 values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *UInt16Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(uint16)
	bAsserted := b.(uint16)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
