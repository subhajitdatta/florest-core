package workerpool

import (
	"fmt"
	"strconv"

	"github.com/jabong/florest-core/src/common/logger"
)

// WPType specifies the worker pool  where each  worker's channel will be put into
type WPType chan chan Task

type WPExecutor struct {
	nWorkers      int
	taskQueueSize int
	taskQueue     chan Task
	workerPool    WPType
}

// Creates a new worker pool executor and initializes it.
func NewWPExecutor(conf Config) (workerPoolExecutor *WPExecutor, errObj error) {
	defer func() {
		// do recovery if error
		if r := recover(); r != nil {
			errObj = fmt.Errorf("Failed to create workerpool executor, Error:%s", r)
			return
		}
	}()

	workerPoolExecutor = new(WPExecutor)
	workerPoolExecutor.nWorkers = conf.NWorkers
	workerPoolExecutor.taskQueueSize = conf.TaskQueueSize

	// First, initialize the channel we are going to put the worker's worker channels into.
	workerPoolExecutor.workerPool = make(WPType, conf.NWorkers)

	workerPoolExecutor.taskQueue = make(chan Task, conf.TaskQueueSize)

	// Now, create all the workers.
	for i := 0; i < conf.NWorkers; i++ {
		logger.Debug("Starting worker "+strconv.Itoa(i+1), false)
		worker := NewWorker(i+1, workerPoolExecutor.workerPool)
		worker.Start()
	}

	go func() {
		defer func() {
			// do recovery if error
			if r := recover(); r != nil {
				logger.Error(fmt.Sprintf("Error in starting the worker pool, Error:%s", r))
				return
			}
		}()
		for task := range workerPoolExecutor.taskQueue {
			logger.Debug("Received task requeust", false)
			workerChan := <-workerPoolExecutor.workerPool
			logger.Debug("Dispatching task to the worker", false)
			workerChan <- task

		}
	}()

	return workerPoolExecutor, nil
}

// Dispatches the task to an available worker to execute it
func (workerPoolExecutor *WPExecutor) ExecuteTask(t Task) {
	if workerPoolExecutor.taskQueue == nil {
		panic(fmt.Sprintf("WPExecutor is not initalized. Use workerPool.NewWPExecutor method to create & initialize workerPoolExecutor"))
	}
	workerPoolExecutor.taskQueue <- t
}
