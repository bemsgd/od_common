package logging

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	AppName string
	LogDir  string
}

type LogLevel string

const (
	Info    LogLevel = "INFO"
	Warning LogLevel = "WARN"
	Error   LogLevel = "ERROR"
)

var logging *Logger
var lock = &sync.Mutex{}

func GetLogger(appName, logDir string) *Logger {
	if logging == nil {
		lock.Lock()

		defer lock.Unlock()
		logging = &Logger{
			AppName: appName,
			LogDir:  logDir,
		}
		return logging
	}
	return logging
}

func (log *Logger) Add(ll LogLevel, message string, opt ...string) {
	logTime := time.Now().Format("20060102 15:04:05.000")
	logFileTime := time.Now().Format("2006010215")
	logFileName := fmt.Sprintf("%v%v-%v.log", log.LogDir, log.AppName, logFileTime)
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	datawriter := bufio.NewWriter(logFile)
	logMsg := ""
	if len(opt) > 0 {
		logMsg = fmt.Sprintf("%v %v	%v %v\n", logTime, ll, message, opt)
	} else {
		logMsg = fmt.Sprintf("%v %v	%v \n", logTime, ll, message)
	}

	datawriter.WriteString(logMsg)
	fmt.Print(logMsg)
	datawriter.Flush()
	logFile.Close()
}
