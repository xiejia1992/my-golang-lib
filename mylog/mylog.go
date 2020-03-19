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

func (self *myLog) Debug (s string) {
	ct := currentTime()
	wg.Add(1)
	go func(){
		defer wg.Done()
		fileWriteMutex.Lock()
		defer fileWriteMutex.Unlock()

		writer, err := self.writer()
		if err != nil {
			return
		}
		defer writer.Close()
		logContent := ct + " [Debug] \"" + s + "\"\n"
		_, err = io.WriteString(writer, logContent)
		if err != nil {
			fmt.Println(err)
		}
	}()
	wg.Wait()
}

func (self *myLog) Info (s string) {
	ct := currentTime()
	wg.Add(1)
	go func(){
		defer wg.Done()
		fileWriteMutex.Lock()
		defer fileWriteMutex.Unlock()

		writer, err := self.writer()
		if err != nil {
			return
		}
		defer writer.Close()
		logContent := ct + " [Info] \"" + s + "\"\n"
		_, err = io.WriteString(writer, logContent)
		if err != nil {
			fmt.Println(err)
		}
	}()
	wg.Wait()
}

func (self *myLog) Warning (s string) {
	ct := currentTime()
	wg.Add(1)
	go func(){
		defer wg.Done()
		fileWriteMutex.Lock()
		defer fileWriteMutex.Unlock()

		writer, err := self.writer()
		if err != nil {
			return
		}
		defer writer.Close()
		logContent := ct + " [Warning] \"" + s + "\"\n"
		_, err = io.WriteString(writer, logContent)
		if err != nil {
			fmt.Println(err)
		}
	}()
	wg.Wait()
}

func (self *myLog) Error (s string) {
	ct := currentTime()
	wg.Add(1)
	go func(){
		defer wg.Done()
		fileWriteMutex.Lock()
		defer fileWriteMutex.Unlock()

		writer, err := self.writer()
		if err != nil {
			return
		}
		defer writer.Close()
		logContent := ct + " [Error] \"" + s + "\"\n"
		_, err = io.WriteString(writer, logContent)
		if err != nil {
			fmt.Println(err)
		}
	}()
	wg.Wait()
}

func currentTime() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	return timeStr
}

func (self *myLog) writer() (*os.File, error) {
	writer, err := os.OpenFile(self.logFileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return writer, err
}