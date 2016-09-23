package workerpool

type Task struct {
	Instance   interface{}
	MethodName string
	Args       []interface{}
}
