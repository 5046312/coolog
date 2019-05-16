package adapter

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	DEFAULT_PATH      = "runtime/logs/"
	DEFAULT_FILENAME  = "2006-01-02"
	DEFAULT_EXT       = ".log"
	DEFAULT_SINGLE    = false
	DEFAULT_Max_SIZE  = 5 * 1024
	DEFAULT_Max_Files = 10
	DEFAULT_JSON      = false
)

type FileConfig struct {
	mu *sync.RWMutex
	F  *os.File

	path      string // Log Folder Path
	filename  string // Time Format of File Names
	ext       string // Log File Suffix
	single    bool   // Whether to save logs for a single file
	max_size  int64  // Upper limit of file capacity when splitting files when non-single file logs
	max_files int    // Early logs that exceed the number of files will be deleted automatically, and no deletions will be made for 0.
	json      bool   // JSON format
}

// Get the default configuration item for the file log
func DefaultFileConfig() *FileConfig {
	fc := &FileConfig{
		mu:        new(sync.RWMutex),
		path:      DEFAULT_PATH,
		filename:  DEFAULT_FILENAME,
		ext:       DEFAULT_EXT,
		single:    DEFAULT_SINGLE,
		max_size:  DEFAULT_Max_SIZE,
		max_files: DEFAULT_Max_Files,
		json:      DEFAULT_JSON,
	}
	fc.getFile()
	return fc
}

// Write a line of string to the log file
func (fc *FileConfig) Write(content string) {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	// TODO: Determine whether it has been initialized and locked
	fc.getFile().Write([]byte(content + "\n"))
	fc.splitLog()
}

// Open the log folder full path
func (fc *FileConfig) getFullDirPath() string {
	dir, _ := os.Getwd()
	return dir + "/" + strings.Trim(fc.path, "/") + "/"
}

// Get the main log file name without ext in the corresponding format for today
func (fc *FileConfig) getFilename() string {
	return time.Now().Format(fc.filename)
}

// Get log file ext
func (fc *FileConfig) getExt() string {
	return "." + strings.Trim(fc.ext, ".")
}

// Get the log file full path
func (fc *FileConfig) getFullFilePath() string {
	return fc.getFullDirPath() + fc.getFilename() + fc.getExt()
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
	fmt.Println("main log file size:", fileSize, ", max size:", fc.max_size)

	// When current file size more than config, split file
	if fileSize >= fc.max_size {
		// Create new main log file
		defer fc.getFile()
		fc.close()
		todayFiles := fc.getTodayFiles()
		filesCount := len(todayFiles)
		// Rename main log file
		newName := fc.getFullDirPath() + fc.getFilename() + "-p" + strconv.Itoa(filesCount) + fc.getExt()
		os.Rename(fc.getFullFilePath(), newName)
	}
}

// Check whether the number of log files exceeds the maximum
func (fc *FileConfig) getAllFiles() []string {
	path := fc.getFullDirPath()
	files, _ := filepath.Glob(path + "*" + fc.getExt())
	return files
}

// Get the number of log files for today
func (fc *FileConfig) getTodayFiles() []string {
	path := fc.getFullDirPath()
	files, _ := filepath.Glob(path + fc.getFilename() + "*" + fc.getExt())
	return files
}

// Close file
func (fc *FileConfig) close() {
	fc.F.Close()
	fc.F = nil
}
