// Package workerpool provides a buffered channel of workers that can be used to perform
// task asynchronously in the background.
//
// NOTE
//
// 1. taskqueue internally uses buffered channel. So when the task queue is already full
// and if the caller is trying to submit a new task, then the caller will be blocked.
//
// 2. workerpool is mainly designed to execute task in the background asynchronously. So
// right now, it is not possible to return some value (including errors) to the job submitter.
//
// Refer https://github.com/jabong/florest-core/wiki/Worker-Pool for more details
package workerpool
