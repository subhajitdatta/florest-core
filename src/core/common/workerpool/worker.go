package workerpool

import (
	"fmt"
	"reflect"

	"github.com/jabong/florest-core/src/common/logger"
)

// NewWorker creates, and returns a new worker instance. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// task.
func NewWorker(id int, workerPool chan chan Task) Worker {
	// Creates and return the worker.
	worker := Worker{
		ID:         id,
		WorkerChan: make(chan Task),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool)}

	return worker
}

type Worker struct {
	ID         int
	WorkerChan chan Task
	WorkerPool chan chan Task
	QuitChan   chan bool
}

// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (t Worker) Start() {
	go func() {
		defer func() {
			// do recovery if error
			if r := recover(); r != nil {
				logger.Error(fmt.Sprintf("Error in starting the worker, Error:%s", r))
				return
			}
		}()
		for {
			// Add its own task channel into the worker pool.
			t.WorkerPool <- t.WorkerChan

			select {
			case task := <-t.WorkerChan:
				t.execute(task)
			case <-t.QuitChan:
				// We have been asked to stop.
				logger.Debug(fmt.Sprintf("worker%d stopping\n", t.ID), false)
				return
			}
		}
	}()
}

// Stop tells the worker to stop listening for task requests.
// Note that the worker will only stop *after* it has finished its current task.
func (t Worker) Stop() {
	t.QuitChan <- true
}

//execute executes the task to be done by the worker
func (t Worker) execute(task Task) {
	t.callMethod(task.Instance, task.MethodName, task.Args)
}

// Using reflection, it calls the method on the given instance with the given arguments
func (t Worker) callMethod(instance interface{}, methodName string, args []interface{}) {
	defer func() {
		// do recovery if error
		if r := recover(); r != nil {
			logger.Error(fmt.Sprintf("Error in calling the method-%s, Error:%s", methodName, r))
			return
		}
	}()

	var ptr reflect.Value
	var value reflect.Value
	var finalMethod reflect.Value

	value = reflect.ValueOf(instance)

	// if we start with a pointer, we need to get value pointed to
	// if we start with a value, we need to get a pointer to that value
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(instance))
		temp := ptr.Elem()
		temp.Set(value)
	}

	// check for method on value
	method := value.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}
	// check for method on pointer
	method = ptr.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}

	inputs := make([]reflect.Value, len(args))
	for i, v := range args {
		inputs[i] = reflect.ValueOf(v)
	}

	if finalMethod.IsValid() {
		finalMethod.Call(inputs)
	} else {
		logger.Error("This method - " + methodName + " is not valid. Hence, Ignoring.")
	}
}
