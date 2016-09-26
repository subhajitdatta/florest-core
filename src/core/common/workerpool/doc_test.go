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
	c := workerpool.Config{NWorkers: 10, TaskQueueSize: 15}

	//Instantiate a WPExecutor which internally creates a taskqueue and workerchannelpool :
	wp, err := workerpool.NewWPExecutor(c)
	if err != nil {
		log.Printf("Failed to instantiate WPExecutor")
	}

	// Submit a task
	hello := new(Hello)
	b := make([]interface{}, 1)
	b[0] = "Florest"
	task := workerpool.Task{hello, "DisplayMsg", b}

	// Execute the task
	wp.ExecuteTask(task)
}
