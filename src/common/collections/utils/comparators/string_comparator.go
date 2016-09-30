package comparators

//StringComparator for comparing the string values
type StringComparator struct {
}

//NewStringComparator returns the new string comparator
func NewStringComparator() *StringComparator {
	return &StringComparator{}
}

// Compare two string values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *StringComparator) Compare(a, b interface{}) int {
	s1 := a.(string)
	s2 := b.(string)
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}
