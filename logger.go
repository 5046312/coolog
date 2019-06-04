package coolog

import (
	"fmt"
	"github.com/5046312/coolog/adapter"
	"path/filepath"
	"runtime"
	"time"
)

type level int

const (
	levelDebug level = iota
	levelInfo
	levelNotice
	levelWarning
	levelError
)

var levelTags = map[level]string{
	levelDebug:   "Debug",
	levelInfo:    "Info",
	levelNotice:  "Notice",
	levelWarning: "Warning",
	levelError:   "Error",
}

type coolog struct {
	Adapter *adapter.Adapter
	Config  *Config
}

// Only coolog Var
var logger *coolog

// Get coolog
func getCoolog() *coolog {
	if logger == nil {
		logger = &coolog{
			Config: new(Config),
			Adapter: &adapter.Adapter{
				FileLog: nil,
			},
		}
	}
	return logger
}

// Get default file log config
func FileConfig() *adapter.FileLog {
	return adapter.DefaultFileLogConfig()
}

// Set File Log Config
func SetFile(fl *adapter.FileLog) *coolog {
	// Uninitialized
	if getCoolog().Adapter.FileLog == nil {
		getCoolog().Adapter.FileLog = fl.InitFileLog()
	} else {
		Warning("Has Been Initialized For File")
	}
	return logger
}

// Including multiple adapters working at the same time
func (log *coolog) Write(content string) {
	if log.Adapter.FileLog != nil {
		// Write in files
		fmt.Print(content)
		log.Adapter.FileLog.Write(content)
	}
	// Todo more adapter
}

// Format log content
func format(l level, m ...interface{}) string {
	_, file, line, _ := runtime.Caller(2)
	// Log text temp
	temp := "[ %s ] %s: [%s:%d] "
	filename := filepath.Base(file)
	times := time.Now().Format("2006-01-02 15:04:05.000")
	return fmt.Sprintf(temp, times, levelTags[l], filename, line) + fmt.Sprintln(m...)
}

// Write Directly Without Initialization
func checkLoggerInit() *coolog {
	// Initialization of default configuration occurs when global variables are empty
	if getCoolog().Adapter.FileLog == nil {
		SetFile(FileConfig())
	}
	return logger
}

func Record(content string) {
	checkLoggerInit().Write(content)
}

// Debug
func Debug(content ...interface{}) {
	msg := format(levelDebug, content...)
	checkLoggerInit().Write(msg)
}

//
func Info(content ...interface{}) {
	msg := format(levelInfo, content...)
	checkLoggerInit().Write(msg)
}

//
func Notice(content ...interface{}) {
	msg := format(levelNotice, content...)
	checkLoggerInit().Write(msg)
}

//
func Warning(content ...interface{}) {
	msg := format(levelWarning, content...)
	checkLoggerInit().Write(msg)
}

//
func Error(content ...interface{}) {
	msg := format(levelError, content...)
	checkLoggerInit().Write(msg)
}
