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
	adapter *adapter.Adapter
	config  *Config
}

// Get default file log config
func FileConfig() *adapter.FileConfig {
	return adapter.DefaultFileConfig()
}

// Create a file log
func NewFileLog(fc *adapter.FileConfig) *Coolog {
	ad := adapter.NewFileAdapter(fc)
	return &Coolog{adapter: ad}
}

// Including multiple adapters working at the same time
func (log *Coolog) Write(content string) {
	if log.adapter.File != nil {
		// Write in files
		fmt.Print(content)
		log.adapter.File.Write(content)
	}
	// Todo more adapter
}

// Format log content
func (log *Coolog) format(l Level, m ...interface{}) string {
	_, file, line, _ := runtime.Caller(2)
	// Log text temp
	temp := "[ %s ] %s: [%s:%d] "
	filename := filepath.Base(file)
	return fmt.Sprintf(temp, time.Now().Format("2006-01-02 15:04:05"), LevelTag[l], filename, line) + fmt.Sprintln(m...)
}

//
func (log *Coolog) Debug(content ...interface{}) {
	msg := log.format(LEVEL_DEBUG, content...)
	log.Write(msg)
}

//
func (log *Config) Info(content string) {}

//
func (log *Config) Notice(content string) {}

//
func (log *Config) Warning(content string) {}

//
func (log *Config) Error(content string) {}
