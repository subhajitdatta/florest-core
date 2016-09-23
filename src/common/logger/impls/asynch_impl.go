package impls

import (
	"errors"
	"fmt"
	"github.com/jabong/florest-core/src/common/logger/message"
	"github.com/jabong/florest-core/src/common/logger/writers"
	"time"
)

// AsynchLoggerConfig config for async logger
type AsynchLoggerConfig struct {
	/*
	 * The variables are user configurable
	 */
	Enabled bool
	// requestQSize is the size of the queue for holding the log messages till
	// they are processed by a watcher
	LogQSize int

	// FlushInterval time in seconds
	FlushInterval int
}

// AsynchLogger implements logger interface with async functionality
type AsynchLogger struct {

	/*
	 * The variables are user configurable
	 */

	// requestQSize is the size of the queue for holding the log messages till
	// they are processed by a watcher
	requestQSize int

	// flushInterval time in seconds
	flushInterval int

	// requestQ is the queue used in logging
	requestQ chan *message.LogMsg

	// flush signal for the watcher to write
	flush chan bool

	// quit signal for the watcher to quit
	quit chan bool

	// writer interface for writing logs
	writer writers.LogWriter
}

// Destroy sends quit signal to watcher and releases all the resources.
func (logr *AsynchLogger) Destroy() {
	// quit watcher
	logr.quit <- true
	// wait for watcher quit
	<-logr.quit
}

// Flush the writer
func (logr *AsynchLogger) Flush() {
	// send flush signal
	logr.flush <- true
	// wait for flush finish
	<-logr.flush
}

// flushBuf flushes the content of buffer to out and reset the buffer
func (logr *AsynchLogger) flushBuf(msg *message.LogMsg) {
	if msg != nil {
		logr.writer.Write(msg)
	}
}

// recoverFromPanic recovers from a panic from the surrounding function
// , creates an error from the panic and returns it
func (logr AsynchLogger) recoverFromPanic(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("%s", r)
	}
}

// log records log v... with level `level'.
func (logr *AsynchLogger) Log(msg *message.LogMsg) (err error) {
	defer logr.recoverFromPanic(&err)
	logr.requestQ <- msg
	return nil
}

// watcher watches the logr.queue channel, and writes the logs to output
func (logr *AsynchLogger) watcher() {
	for {
		timeout := time.After(time.Second / time.Duration(logr.flushInterval))

		select {
		case req := <-logr.requestQ:
			logr.flushBuf(req)
		case <-timeout:
		// nothing
		case <-logr.flush:
			finish := false
			for {
				select {
				case req := <-logr.requestQ:
					logr.flushBuf(req)
				default:
					finish = true
				}
				if finish {
					logr.flush <- true
					break
				}
			}

		case <-logr.quit:
			// If quit signal received, cleans the channel
			// and writes all of them to io.Writer.
			for {
				select {
				case req := <-logr.requestQ:
					logr.flushBuf(req)
				default:
					logr.quit <- true
					return
				}
			}
		}
	}
}

// newLogger creates & returns a new instance of logger from the supplied conf
func NewAsynchLogger(conf *AsynchLoggerConfig, writer writers.LogWriter) (*AsynchLogger, error) {
	l := new(AsynchLogger)

	l.flushInterval = conf.FlushInterval
	if l.flushInterval <= 0 {
		return nil, errors.New("Incorrect flush interval for aynch logging")
	}

	l.requestQSize = conf.LogQSize
	if l.requestQSize <= 0 {
		return nil, errors.New("Incorrect log queue size for aynch logging")
	}
	l.requestQ = make(chan *message.LogMsg, l.requestQSize)

	l.flush = make(chan bool)
	l.quit = make(chan bool)

	l.writer = writer
	go l.watcher()
	return l, nil
}

// Debug log the debug message in queue
func (logr *AsynchLogger) Debug(msg message.LogMsg) {
	logr.Log(&msg)
}

// Info log the info message in queue
func (logr *AsynchLogger) Info(msg message.LogMsg) {
	logr.Log(&msg)
}

// Trace log the trace message in queue
func (logr *AsynchLogger) Trace(msg message.LogMsg) {
	logr.Log(&msg)
}

// Warning log the warning message in queue
func (logr *AsynchLogger) Warning(msg message.LogMsg) {
	logr.Log(&msg)
}

// Error log the error message in queue
func (logr *AsynchLogger) Error(msg message.LogMsg) {
	logr.Log(&msg)
}

// Profile log the profile message in queue
func (logr *AsynchLogger) Profile(msg message.LogMsg) {
	logr.Log(&msg)
}
