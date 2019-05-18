package coolog

import (
	"fmt"
	"github.com/5046312/coolog/adapter"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

const (
	LEVEL_DEBUG Level = iota
	LEVEL_INFO
	LEVEL_NOTICE
	LEVEL_WARNING
	LEVEL_ERROR
)

var LevelTag = map[Level]string{
	LEVEL_DEBUG:   "Debug",
	LEVEL_INFO:    "Info",
	LEVEL_NOTICE:  "Notice",
	LEVEL_WARNING: "Warning",
	LEVEL_ERROR:   "Error",
}

type Coolog struct {
	Adapter *adapter.Adapter
	Config  *Config
}

// Only Coolog Var
var logger *Coolog

// Get Coolog
func GetCoolog() *Coolog {
	if logger == nil {
		logger = &Coolog{
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
func SetFileLog(fl *adapter.FileLog) *Coolog {
	// Uninitialized
	if GetCoolog().Adapter.FileLog == nil {
		GetCoolog().Adapter.FileLog = fl.InitFileLog()
	} else {
		Warning("Coolog Has Been Initialized For File")
	}
	return logger
}

// Including multiple adapters working at the same time
func (log *Coolog) Write(content string) {
	if log.Adapter.FileLog != nil {
		// Write in files
		fmt.Print(content)
		log.Adapter.FileLog.Write(content)
	}
	// Todo more adapter
}

// Format log content
func (log *Coolog) format(l Level, m ...interface{}) string {
	_, file, line, _ := runtime.Caller(2)
	// Log text temp
	temp := "[ %s ] %s: [%s:%d] "
	filename := filepath.Base(file)
	times := time.Now().Format("2006-01-02 15:04:05.000")
	return fmt.Sprintf(temp, times, LevelTag[l], filename, line) + fmt.Sprintln(m...)
}

// Write Directly Without Initialization
func checkLoggerInit() *Coolog {
	// Initialization of default configuration occurs when global variables are empty
	if GetCoolog().Adapter.FileLog == nil {
		SetFileLog(FileConfig())
	}
	return logger
}

// Debug
func Debug(content ...interface{}) {
	msg := checkLoggerInit().format(LEVEL_DEBUG, content...)
	logger.Write(msg)
}

//
func Info(content ...interface{}) {
	msg := checkLoggerInit().format(LEVEL_INFO, content...)
	logger.Write(msg)
}

//
func Notice(content ...interface{}) {
	msg := checkLoggerInit().format(LEVEL_NOTICE, content...)
	logger.Write(msg)
}

//
func Warning(content ...interface{}) {
	msg := checkLoggerInit().format(LEVEL_WARNING, content...)
	logger.Write(msg)
}

//
func Error(content ...interface{}) {
	msg := checkLoggerInit().format(LEVEL_ERROR, content...)
	logger.Write(msg)
}
