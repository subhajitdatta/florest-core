package comparators

//UInt8Comparator for comparing the uint8 values
type UInt8Comparator struct {
}

//NewUInt8Comparator returns the new uint8 comparator
func NewUInt8Comparator() *UInt8Comparator {
	return &UInt8Comparator{}
}

// Compare two uint8 values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *UInt8Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(uint8)
	bAsserted := b.(uint8)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
