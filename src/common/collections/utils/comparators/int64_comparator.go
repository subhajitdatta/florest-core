package comparators

//Int64Comparator for comparing the int64 values
type Int64Comparator struct {
}

//NewInt64Comparator returns the new int64 comparator
func NewInt64Comparator() *Int64Comparator {
	return &Int64Comparator{}
}

// Compare two int64 values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *Int64Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(int64)
	bAsserted := b.(int64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
