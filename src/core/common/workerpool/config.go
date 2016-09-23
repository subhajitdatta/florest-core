package workerpool

type Config struct {
	// NWorkers specifies the numbers of workers to be used in an instance of workerpool
	NWorkers int

	// TaskQueueSize specifies the size of the task queue.
	TaskQueueSize int
}
