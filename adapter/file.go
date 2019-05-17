package adapter

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	DEFAULT_PATH     = "runtime/logs/" // Log File Dir
	DEFAULT_FILENAME = "2006-01-02"    // Log File Name Time Format
	DEFAULT_EXT      = ".log"          // Log File Suffix
	DEFAULT_SINGLE   = false           // Split Files
	DEFAULT_MAX_SIZE = 5 * 1024        // Byte
	DEFAULT_MAX_TIME = 5               // Hour
	DEFAULT_JSON     = false           // Todo: Output Json Format
)

type FileLog struct {
	init     bool        // The initialization status
	logChan  chan string // Log Content Channel
	f        *os.File    //
	Path     string      // Log Folder Path
	Filename string      // Time Format of File Names
	Ext      string      // Log File Suffix
	Single   bool        // Whether to save logs for a single file
	MaxSize  int64       // Upper limit of file capacity when splitting files when non-single file logs
	MaxTime  int         // Files that exceed the maximum retention time(hour) will be deleted, and no deletions will be made for 0.
	Json     bool        // JSON format
}

// Get the default configuration item for the file log
func DefaultFileLogConfig() *FileLog {
	return &FileLog{
		logChan:  make(chan string),
		Path:     DEFAULT_PATH,
		Filename: DEFAULT_FILENAME,
		Ext:      DEFAULT_EXT,
		Single:   DEFAULT_SINGLE,
		MaxSize:  DEFAULT_MAX_SIZE,
		MaxTime:  DEFAULT_MAX_TIME,
		Json:     DEFAULT_JSON,
	}
}

// Init File Log
func (fl *FileLog) InitFileLog() *FileLog {
	if !fl.init {
		fl.init = true
		go func() {
			// fmt.Println("File Log Init")
			fl.InitMainFile()
			for {
				select {
				// Write File Log
				case content := <-fl.logChan:
					_, err := fl.getFile().Write([]byte(content))
					if err != nil {
						log.Fatalf("Write In File Err :%v", err)
					}
					fl.splitLog()
				}
			}
		}()
	}
	return fl
}

// Create Main Log File
func (fl *FileLog) InitMainFile() *os.File {
	// fmt.Println("InitMainFile")
	defer fl.limitFiles()
	filePath := fl.getFullFilePath()
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		fl.mkLogDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
	if fl.f != nil {
		fl.close()
	}
	fl.f, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail To Open Log File :%v", err)
	}
	return fl.f
}

// Write a line of string to the log file
func (fl *FileLog) Write(content string) {
	// TODO: Determine whether it has been initialized and locked
	fl.logChan <- content
}

// Open the log folder full path
func (fl *FileLog) getFullDirPath() string {
	dir, _ := os.Getwd()
	return dir + "/" + strings.Trim(fl.Path, "/") + "/"
}

// Get the main log file name without ext in the corresponding format for today
func (fl *FileLog) getFilename() string {
	return time.Now().Format(fl.Filename)
}

// Get log file ext
func (fl *FileLog) getExt() string {
	return "." + strings.Trim(fl.Ext, ".")
}

// Get the log file full path
func (fl *FileLog) getFullFilePath() string {
	return fl.getFullDirPath() + fl.getFilename() + fl.getExt()
}

// Get log file
func (fl *FileLog) getFile() *os.File {
	if fl.f == nil {
		fl.InitMainFile()
	} else {
		// If file exists, Check date with file name
		filename := filepath.Base(fl.f.Name())
		if filename != fl.getFilename()+fl.getExt() {
			// Recreate a log file
			fl.InitMainFile()
		}
	}
	return fl.f
}

// Create log dir
func (fl *FileLog) mkLogDir() {
	path := fl.getFullDirPath()
	os.MkdirAll(path, os.ModePerm)
}

// Split log file
func (fl *FileLog) splitLog() {
	// Single File Not Split Files
	if fl.Single {
		return
	}
	filePath := fl.getFullFilePath()
	fileInfo, _ := os.Stat(filePath)
	fileSize := fileInfo.Size()
	// When current file size more than config, split file
	// fmt.Println("Compare Size with Max", fileSize, fl.MaxSize)
	if fileSize >= fl.MaxSize {
		// Create new main log file
		fl.close()
		// Rename main log file
		newName := fl.getFullDirPath() + fl.getFilename() + "P" + time.Now().Format("150405") + fl.getExt()
		err := os.Rename(fl.getFullFilePath(), newName)
		if err != nil {
			fmt.Println("Rename Err", err.Error())
		}
		fl.InitMainFile()
	}
}

// Limit the max number of log files
// When Create a File
func (fl *FileLog) limitFiles() {
	if fl.MaxTime > 0 {
		path := fl.getFullDirPath()
		filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
			if info == nil || info.IsDir() {
				return nil
			}
			// TODO:Compare ModTime with MaxTime
			if info.ModTime().Add(time.Hour * time.Duration(fl.MaxTime)).Before(time.Now()) {
				//fmt.Println("Del Old Log:", file)
				os.Remove(file)
			}
			return nil
		})
	}
}

// Check whether the number of log files exceeds the maximum
func (fl *FileLog) getAllFiles() []string {
	path := fl.getFullDirPath()
	files, _ := filepath.Glob(path + "*" + fl.getExt())
	return files
}

// Get the number of log files for today
func (fl *FileLog) getTodayFiles() []string {
	path := fl.getFullDirPath()
	files, _ := filepath.Glob(path + fl.getFilename() + "*" + fl.getExt())
	return files
}

// Close file
func (fl *FileLog) close() {
	fl.f.Close()
	fl.f = nil
}
