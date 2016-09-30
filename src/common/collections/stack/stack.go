package stack

import (
	"github.com/jabong/florest-core/src/common/collections/utils"
)

//Element for the Stack
type element struct {
	value interface{}
	next  *element
}

//Stack structure
type Stack struct {
	top  *element
	size int
}

//Push a Element in the Stack
func (s *Stack) Push(v interface{}) {
	s.top = &element{value: v, next: s.top}
	s.size++
}

//Pop the top Element type from the Stack
//Return nil if the stack is empty
func (s *Stack) Pop() (v interface{}) {
	if s.size <= 0 {
		return nil
	}
	v, s.top = s.top.value, s.top.next
	s.size--

	return v
}

//IsEmpty returns true if Stack is empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Size returns number of elements in the stack.
func (s *Stack) Size() int {
	return s.size
}

// Clear removes all elements from the stack.
func (s *Stack) Clear() {
	s.top = nil
	s.size = 0
}

// Values returns all values in the stack from top to bottom
func (s *Stack) Values() []interface{} {
	current := s.top
	values := make([]interface{}, s.size)
	index := 0
	for current != nil {
		values[index] = current.value
		current = current.next
		index++
	}
	return values
}

// Contains returns true if the given keys are found in the collection
func (s *Stack) Contains(keys ...interface{}) bool {
	elements := s.Values()
	elementsMap := utils.ConvertArrayToMap(elements)
	for _, key := range keys {
		if _, found := elementsMap[key]; !found {
			return false
		}
	}
	return true
}

//Clone creates a deepcopy of the node stack
func (s *Stack) Clone() (r *Stack) {

	//Check for the empty stack
	if s.size == 0 || s.top == nil {
		return new(Stack)
	}
	newTop := &element{s.top.value, nil}
	r = &Stack{newTop, s.size}

	var sCurrElem = s.top.next
	var rPrevElem = r.top

	for sCurrElem != nil {
		newElem := &element{value: sCurrElem.value}

		rPrevElem.next = newElem
		rPrevElem = rPrevElem.next

		sCurrElem = sCurrElem.next
	}

	return r
}
