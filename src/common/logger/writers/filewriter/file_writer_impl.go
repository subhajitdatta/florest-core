package filewriter

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/logger/formatter"
	"github.com/jabong/florest-core/src/common/logger/message"
	"os"
	"sync"
	"time"
)

//FileWriter is a file logger structure
type FileWriter struct {
	//filename is the log file name with absolute path
	filename string

	//file is the file handle for log filename
	file *os.File

	// formatter
	myFormat formatter.FormatInterface

	// mutex lock used to check and open new file
	mutex *sync.Mutex
}

//GetNewObj returns a file logger with log file name fname, having configuration
//specified in conf and allowedLogLevel specifies the log level that are actually to
//be logged
func GetNewObj(fname string) (*FileWriter, error) {
	obj := new(FileWriter)
	obj.filename = fname
	obj.mutex = new(sync.Mutex)
	err := obj.setLogFileHandle()
	if err != nil {
		return nil, err
	}
	return obj, nil
}

//setLogFileHandle opens a file and assigns the file handle to file in fileLogger
func (fw *FileWriter) setLogFileHandle() error {
	logName := fw.getLogFileExt()
	f, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("\nError In Opening log file %s: %+v\n", fw.filename, err)
		return err
	}
	// close old file if any
	if fw.file != nil {
		fw.file.Close()
	}
	fw.file = f
	return nil
}

//getLogFileExt returns the log file ext name after appending
//today's date and ".log" to the filename
func (fw *FileWriter) getLogFileExt() string {
	t := time.Now().Local()
	tf := t.Format("2006-01-02")
	return fw.filename + tf + ".log"
}

// Write write message to file
func (fw *FileWriter) Write(msg *message.LogMsg) {
	str, _ := fw.myFormat.GetFormattedLog(msg).(string)
	fw.chkNSetFile()
	if _, err := fw.file.WriteString(str + "\n"); err != nil {
		fmt.Println("error in writing to " + fw.getLogFileExt() + ":" + err.Error())
	}
}

// chkNSetFile checks current file and changes if date has changed
func (fw *FileWriter) chkNSetFile() {
	if fw.file.Name() != fw.getLogFileExt() { // open new date file
		fw.mutex.Lock()
		defer fw.mutex.Unlock()
		if fw.file.Name() != fw.getLogFileExt() {
			fw.setLogFileHandle()
		}
	}
}

// SetFormatter get formatted object
func (fw *FileWriter) SetFormatter(format formatter.FormatInterface) {
	fw.myFormat = format
}
