package adapter

import (
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	DEFAULT_DIR      = "runtime/logs/"
	DEFAULT_SIZE     = 5
	DEFAULT_FILENAME = "2006-01-02"
	DEFAULT_EXT      = ".log"
)

type FileConfig struct {
	mu        *sync.RWMutex
	F 	 *os.File
	filename string
	dir      string
	size     int
	ext      string
}

// Get the default configuration item for the file log
func DefaultFileConfig() *FileConfig {
	fc := &FileConfig{new(sync.RWMutex), nil,DEFAULT_FILENAME,DEFAULT_DIR, DEFAULT_SIZE, DEFAULT_EXT}
	fc.getFile()
	return fc
}

// Write a line of string to the log file
func (fc *FileConfig) Write(content string) error {
	defer fc.splitLog()
	// TODO: Determine whether it has been initialized and locked
	_, err := fc.getFile().Write([]byte(content + "\n"))
	return err
}

// Open the log folder full path
func (fc *FileConfig) getFullDirPath() string {
	dir, _ := os.Getwd()
	return dir + "/" + strings.Trim(fc.dir, "/") + "/"
}

// Get the log file full path
func (fc *FileConfig) getFullFilePath() string {
	dirPath := fc.getFullDirPath()
	filename := time.Now().Format(fc.filename) + fc.ext
	return dirPath + filename
}

// Get log file
func (fc *FileConfig) getFile() *os.File {
	if fc.F == nil {
		filePath := fc.getFullFilePath()
		_, err := os.Stat(filePath)
		switch {
		case os.IsNotExist(err):
			fc.mkLogDir()
		case os.IsPermission(err):
			log.Fatalf("Permission :%v", err)
		}

		fc.F, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Fail To Open Log File :%v", err)
		}
	}
	return fc.F
}

// Create log dir
func (fc *FileConfig) mkLogDir() {
	path := fc.getFullDirPath()
	os.MkdirAll(path, os.ModePerm)
}

// Split log file
func (fc *FileConfig) splitLog(){

}