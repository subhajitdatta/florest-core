package comparators

//Int8Comparator for comparing the int8 values
type Int8Comparator struct {
}

//NewInt8Comparator returns the new int8 comparator
func NewInt8Comparator() *Int8Comparator {
	return &Int8Comparator{}
}

// Compare two int8 values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *Int8Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(int8)
	bAsserted := b.(int8)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
