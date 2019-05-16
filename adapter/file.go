package adapter

import (
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	DEFAULT_PATH      = "runtime/logs/"
	DEFAULT_FILENAME  = "2006-01-02"
	DEFAULT_EXT       = ".log"
	DEFAULT_SINGLE    = false
	DEFAULT_SIZE      = 5 * 1024
	DEFAULT_Max_Files = 10
)

type FileConfig struct {
	mu *sync.RWMutex
	F  *os.File

	path      string // Log Folder Path
	filename  string // Time Format of File Names
	ext       string // Log File Suffix
	single    bool   // Whether to save logs for a single file
	size      int64  // Upper limit of file capacity when splitting files when non-single file logs
	max_files int    // Early logs that exceed the number of files will be deleted automatically, and no deletions will be made for 0.
}

// Get the default configuration item for the file log
func DefaultFileConfig() *FileConfig {
	fc := &FileConfig{
		mu:        new(sync.RWMutex),
		path:      DEFAULT_PATH,
		filename:  DEFAULT_FILENAME,
		ext:       DEFAULT_EXT,
		single:    DEFAULT_SINGLE,
		size:      DEFAULT_SIZE,
		max_files: DEFAULT_Max_Files,
	}
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
	return dir + "/" + strings.Trim(fc.path, "/") + "/"
}

// Get the log file full path
func (fc *FileConfig) getFullFilePath() string {
	dirPath := fc.getFullDirPath()
	filename := time.Now().Format(fc.filename) + "." + strings.Trim(fc.ext, ".")
	return dirPath + filename
}

// Get log file
func (fc *FileConfig) getFile() *os.File {
	if fc.F == nil {
		filePath := fc.getFullFilePath()
		_, err := os.Stat(filePath)
		switch {
		case os.IsNotExist(err):
			fc.mkLogPATH()
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
func (fc *FileConfig) mkLogPATH() {
	path := fc.getFullDirPath()
	os.MkdirAll(path, os.ModePerm)
}

// Split log file
func (fc *FileConfig) splitLog() {
	filePath := fc.getFullFilePath()
	fileInfo, _ := os.Stat(filePath)
	fileSize := fileInfo.Size()
	// When current file size more than config, split file
	if fileSize > fc.size {

	}
}
