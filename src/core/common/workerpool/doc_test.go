package workerpool_test

import (
	"fmt"
	"log"

	"github.com/jabong/florest-core/src/core/common/workerpool"
)

type Hello struct {
}

func (h *Hello) DisplayMsg(msg string) {
	fmt.Println("Hello " + msg)
}

func Example() {
	c := Config{NWorkers: 10, TaskQueueSize: 15}

	//Instantiate a WPExecutor which internally creates a taskqueue and workerchannelpool :
	wp, err := NewWPExecutor(c)
	if err != nil {
		log.Printf("Failed to instantiate WPExecutor")
	}

	// Submit a task
	hello := new(Hello)
	task := workerpool.Task{hello, "DisplayMsg", "Florest"}

	// Execute the task
	wp.ExecuteTask(task)
}
