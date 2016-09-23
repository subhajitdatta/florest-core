package collections

//Element for the Stack
type element struct {
	value interface{}
	next  *element
}

//Nodes Stack structure
type Stack struct {
	top  *element
	size int
}

//Push a Element in the Stack
func (s *Stack) Push(v interface{}) {
	s.top = &element{value: v, next: s.top}
	s.size++
}

//Remove the top Element type from the Stack
//Return nil if the stack is empty
func (s *Stack) Pop() (v interface{}) {
	if s.size <= 0 {
		return nil
	}
	v, s.top = s.top.value, s.top.next
	s.size--

	return v
}

//Check if Stack is empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

//Create a deepcopy of the node stack
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
