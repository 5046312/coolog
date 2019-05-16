package adapter

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	DEFAULT_PATH     = "runtime/logs/" // Log File Dir
	DEFAULT_FILENAME = "2006-01-02"    // Log File Name Time Format
	DEFAULT_EXT      = ".log"          // Log File Suffix
	DEFAULT_SINGLE   = false           // Split Files
	DEFAULT_Max_SIZE = 5 * 1024        // Byte
	DEFAULT_Max_Time = 5               // Hour
	DEFAULT_JSON     = false           // Output Json Format
)

type FileConfig struct {
	mu       *sync.RWMutex
	f        *os.File
	files    int    // The number of logs in Log Folder
	Path     string // Log Folder Path
	Filename string // Time Format of File Names
	Ext      string // Log File Suffix
	Single   bool   // Whether to save logs for a single file
	Max_size int64  // Upper limit of file capacity when splitting files when non-single file logs
	Max_time int    // Files that exceed the maximum retention time(hour) will be deleted, and no deletions will be made for 0.
	Json     bool   // JSON format
}

// Get the default configuration item for the file log
func DefaultFileConfig() *FileConfig {
	fc := &FileConfig{
		mu:       new(sync.RWMutex),
		Path:     DEFAULT_PATH,
		Filename: DEFAULT_FILENAME,
		Ext:      DEFAULT_EXT,
		Single:   DEFAULT_SINGLE,
		Max_size: DEFAULT_Max_SIZE,
		Max_time: DEFAULT_Max_Time,
		Json:     DEFAULT_JSON,
	}
	return fc
}

// Write a line of string to the log file
func (fc *FileConfig) Write(content string) {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	// TODO: Determine whether it has been initialized and locked
	_, err := fc.getFile().Write([]byte(content))
	if err != nil {
		log.Fatalf("Write in file err :%v", err)
	}
	fc.splitLog()
}

// Open the log folder full path
func (fc *FileConfig) getFullDirPath() string {
	dir, _ := os.Getwd()
	return dir + "/" + strings.Trim(fc.Path, "/") + "/"
}

// Get the main log file name without ext in the corresponding format for today
func (fc *FileConfig) getFilename() string {
	return time.Now().Format(fc.Filename)
}

// Get log file ext
func (fc *FileConfig) getExt() string {
	return "." + strings.Trim(fc.Ext, ".")
}

// Get the log file full path
func (fc *FileConfig) getFullFilePath() string {
	return fc.getFullDirPath() + fc.getFilename() + fc.getExt()
}

// Get log file
func (fc *FileConfig) getFile() *os.File {
	if fc.f == nil {
		filePath := fc.getFullFilePath()
		_, err := os.Stat(filePath)
		switch {
		case os.IsNotExist(err):
			fc.mkLogPATH()
		case os.IsPermission(err):
			log.Fatalf("Permission :%v", err)
		}

		fc.f, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Fail To Open Log File :%v", err)
		}
		fc.limitFiles()
	}
	return fc.f
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
	if fileSize >= fc.Max_size {
		// Create new main log file
		fc.close()
		// Rename main log file
		newName := fc.getFullDirPath() + fc.getFilename() + "P" + time.Now().Format("150405") + fc.getExt()
		os.Rename(fc.getFullFilePath(), newName)
		fc.getFile()
	}
}

// Limit the max number of log files
// When Create a File
func (fc *FileConfig) limitFiles() {
	if fc.Max_time > 0 {
		path := fc.getFullDirPath()
		filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
			if info == nil || info.IsDir() {
				return nil
			}
			// TODO:Compare ModTime with MaxTime
			if info.ModTime().Add(time.Second * time.Duration(fc.Max_time)).Before(time.Now()) {
				//fmt.Println("Del Old Log:", file)
				os.Remove(file)
			}
			return nil
		})
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
	fc.f.Close()
	fc.f = nil
}
