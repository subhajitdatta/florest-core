package comparators

//Int32Comparator for comparing the int32 values
type Int32Comparator struct {
}

//NewInt32Comparator returns the new int32 comparator
func NewInt32Comparator() *Int32Comparator {
	return &Int32Comparator{}
}

// Compare two int32 values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *Int32Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(int32)
	bAsserted := b.(int32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
