package coolog

import (
	"coolog/adapter"
	"fmt"
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
var log *Coolog

// Get Coolog
func GetCoolog() *Coolog {
	if log == nil {
		log = &Coolog{
			Adapter: new(adapter.Adapter),
			Config:  new(Config),
		}
	}
	return log
}

// Get default file log config
func FileConfig() *adapter.FileLog {
	return adapter.DefaultFileLogConfig()
}

// Create a file log
func NewFileLog(fl *adapter.FileLog) *Coolog {
	GetCoolog().Adapter.FileLog = fl.InitFileLog()
	return log
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
	time := time.Now().String()[:23]
	return fmt.Sprintf(temp, time, LevelTag[l], filename, line) + fmt.Sprintln(m...)
}

//
func (log *Coolog) Debug(content ...interface{}) {
	msg := log.format(LEVEL_DEBUG, content...)
	log.Write(msg)
}

//
func (log *Coolog) Info(content ...interface{}) {
	msg := log.format(LEVEL_INFO, content...)
	log.Write(msg)
}

//
func (log *Coolog) Notice(content ...interface{}) {
	msg := log.format(LEVEL_NOTICE, content...)
	log.Write(msg)
}

//
func (log *Coolog) Warning(content ...interface{}) {
	msg := log.format(LEVEL_WARNING, content...)
	log.Write(msg)
}

//
func (log *Coolog) Error(content ...interface{}) {
	msg := log.format(LEVEL_ERROR, content...)
	log.Write(msg)
}
