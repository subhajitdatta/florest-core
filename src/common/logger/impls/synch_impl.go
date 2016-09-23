package impls

import (
	"github.com/jabong/florest-core/src/common/logger/message"
	"github.com/jabong/florest-core/src/common/logger/writers"
)

// SynchLogger implements logger interface with sync functionality
type SynchLogger struct {
	writer writers.LogWriter
}

// GetSynchLogger returns the SynchLogger instance
func GetSynchLogger(writer writers.LogWriter) *SynchLogger {
	ret := new(SynchLogger)
	ret.writer = writer
	return ret
}

// Debug write debug message
func (logr *SynchLogger) Debug(msg message.LogMsg) {
	logr.writer.Write(&msg)
}

// Info write info message
func (logr *SynchLogger) Info(msg message.LogMsg) {
	logr.writer.Write(&msg)
}

// Trace write trace message
func (logr *SynchLogger) Trace(msg message.LogMsg) {
	logr.writer.Write(&msg)
}

// Warning write warning message
func (logr *SynchLogger) Warning(msg message.LogMsg) {
	logr.writer.Write(&msg)
}

// Error write error message
func (logr *SynchLogger) Error(msg message.LogMsg) {
	logr.writer.Write(&msg)
}

// Profile write profile message
func (logr *SynchLogger) Profile(msg message.LogMsg) {
	logr.writer.Write(&msg)
}
