package comparators

//Float64Comparator for comparing the float64 values
type Float64Comparator struct {
}

//NewFloat64Comparator returns the new float64 comparator
func NewFloat64Comparator() *Float64Comparator {
	return &Float64Comparator{}
}

// Compare two float64 values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *Float64Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(float64)
	bAsserted := b.(float64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
