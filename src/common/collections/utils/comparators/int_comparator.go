package comparators

//IntComparator for comparing the int values
type IntComparator struct {
}

//NewIntComparator returns the new int comparator
func NewIntComparator() *IntComparator {
	return &IntComparator{}
}

// Compare two int values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *IntComparator) Compare(a, b interface{}) int {
	aAsserted := a.(int)
	bAsserted := b.(int)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
