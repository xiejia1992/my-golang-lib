package mylog

import (
	"time"
	"os"
	"sync"
	"io"
	"fmt"
)


type myLog struct {
	logFileName string
}

var (
	fileWriteMutex *sync.Mutex
	wg sync.WaitGroup
)

func NewMyLog(logFileName string) *myLog {
	mylog := &myLog{
		logFileName: logFileName,
	}
	fileWriteMutex = new(sync.Mutex)
	return mylog
}

func (self *myLog) Debug (message string) {
	wg.Add(1)
	go self.writer("Debug", message)
	wg.Wait()
}

func (self *myLog) Info (message string) {
	wg.Add(1)
	go self.writer("Info", message)
	wg.Wait()
}

func (self *myLog) Warning (message string) {
	wg.Add(1)
	go self.writer("Warning", message)
	wg.Wait()
}

func (self *myLog) Error (message string) {
	wg.Add(1)
	go self.writer("Error", message)
	wg.Wait()
}

func currentTime() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	return timeStr
}

func (self *myLog) openFile() (*os.File, error) {
	file, err := os.OpenFile(self.logFileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return file, err
}

func (self *myLog) writer(level, message string) {
	ct := currentTime()
	defer wg.Done()
	fileWriteMutex.Lock()
	defer fileWriteMutex.Unlock()

	file, err := self.openFile()
	if err != nil {
		return
	}
	defer file.Close()
	logContent := ct + " [" + level + "] \"" + message + "\"\n"
	_, err = io.WriteString(file, logContent)
	if err != nil {
		fmt.Println(err)
	}
}