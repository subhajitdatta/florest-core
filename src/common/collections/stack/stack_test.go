package stack

import (
	"fmt"
	"reflect"
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

	if actualSize := stck.Size(); actualSize != 0 {
		t.Errorf("Size() - Got %v expected %v", actualSize, 0)
	}

	stck.Push(10)
	if stck.IsEmpty() {
		t.Error("IsEmpty check failed after stack push")
	}

	if actualSize := stck.Size(); actualSize != 1 {
		t.Errorf("Size() - Got %v expected %v", actualSize, 1)
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
Test Stack Contains Clear
*/
func TestStackContainsClear(t *testing.T) {

	stck := new(Stack)
	stck.Push(10)
	stck.Push(20)

	// Testing contains method
	if contains := stck.Contains(10, 20); !contains {
		t.Errorf("Contains() - Got %v expected %v", contains, true)
	}

	if contains := stck.Contains(15, 20); contains {
		t.Errorf("!Contains() - Got %v expected %v", contains, false)
	}

	// Testing Clear method
	stck.Clear()

	if actualValues, expectedValues := fmt.Sprintf("%s", stck.Values()), "[]"; actualValues != expectedValues {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}
	if actualSize := stck.Size(); actualSize != 0 {
		t.Errorf("Size() - Got %v expected %v", actualSize, 0)
	}
	if isEmpty := stck.IsEmpty(); !isEmpty {
		t.Errorf("IsEmpty() - Got %v expected %v", isEmpty, true)
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

	v1 := V{x: 10}
	v2 := V{x: 20}

	stck := new(Stack)
	stck.Push(v1)
	stck.Push(v2)

	if actualValues, expectedValues := stck.Values(), []interface{}{v2, v1}; !Equals(actualValues, expectedValues) {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}

	cloneStck := stck.Clone()

	if cloneStck.size != stck.size {
		t.Error("Failed to clone empty stack")
	}

	if actualValues, expectedValues := cloneStck.Values(), []interface{}{v2, v1}; !Equals(actualValues, expectedValues) {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}
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

func Equals(a []interface{}, b []interface{}) bool {
	return reflect.DeepEqual(a, b)
}
