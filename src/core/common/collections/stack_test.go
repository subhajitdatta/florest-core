package collections

import (
	"testing"
)

/*
Test Stack isEmpty
*/
func TestStackIsEmpty(t *testing.T) {

	stck := new(Stack)

	if !stck.IsEmpty() {
		t.Error("IsEmpty check failed for empty stack")
	}

	stck.Push(10)
	if stck.IsEmpty() {
		t.Error("IsEmpty check failed after stack push")
	}

	stck.Pop()
	if !stck.IsEmpty() {
		t.Error("IsEmpty check failed after stack pop")
	}
}

/*
Test Stack push
*/
func TestStackPush(t *testing.T) {

	stck := new(Stack)
	stck.Push(10)

	if stck.size != 1 {
		t.Error("Stack size check failed for push")
	}
	if stck.top == nil {
		t.Error("Stack top element not set for push")
	}
	v := stck.top.value
	if _, ok := v.(int); !ok {
		t.Error("Type assestion failed for the stack top element")
	}
}

/*
Test Stack pop
*/
func TestStackPop(t *testing.T) {

	stck := new(Stack)
	stck.Push(10)
	stck.Push(20)

	v1 := stck.Pop()
	if val, ok := v1.(int); !ok || val != 20 {
		t.Error("Type assestion/Incorrect value for the stack pop")
	}

	v2 := stck.Pop()
	if val, ok := v2.(int); !ok || val != 10 {
		t.Error("Type assestion/Incorrect value for the stack pop")
	}
}

/*
Test Stack Clone for empty stack
*/
func TestEmptyStackClone(t *testing.T) {
	stck := new(Stack)
	cloneStck := stck.Clone()

	if cloneStck.size != 0 || cloneStck.top != nil {
		t.Error("Failed to clone empty stack")
	}
}

/*
Test Stack Clone
*/
func TestStackClone(t *testing.T) {
	type V struct {
		x int
	}

	stck := new(Stack)
	stck.Push(V{x: 10})
	stck.Push(V{x: 20})

	cloneStck := stck.Clone()

	if cloneStck.size != stck.size {
		t.Error("Failed to clone empty stack")
	}

	// 	if cloneStck.Pop() == stck.Pop() {
	// 		t.Error("Stack Clone is not deep copy")
	// 	}
}

/*
Test Stack Push and Pop with nil values
*/
func TestStackNilPushPop(t *testing.T) {
	stck := new(Stack)
	stck.Push(nil)

	if stck.size != 1 {
		t.Error("Size check failed after stack push with nil value")
	}

	v := stck.Pop()
	if v != nil {
		t.Error("Stack pop failed for nil value pushed in stack")
	}
}

/*
Test Stack Clone for stack which contains nil values
*/
func TestStackNilClone(t *testing.T) {
	stck := new(Stack)
	stck.Push(nil)
	stck.Push(10)

	if stck.size != 2 {
		t.Error("Size check failed after stack push with nil value")
	}

	cloneStck := stck.Clone()

	v1 := cloneStck.Pop()
	if val, ok := v1.(int); !ok || val != 10 {
		t.Error("Type assestion/Incorrect value for the stack pop with nil value")
	}

	v2 := cloneStck.Pop()
	if v2 != nil {
		t.Error("Type assestion/Incorrect value for the stack pop with nil value")
	}

}
