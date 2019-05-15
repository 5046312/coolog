package adapter

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	DEFAULT_DIR      = "runtime/logs/"
	DEFAULT_SIZE     = 5
	DEFAULT_FILENAME = "2016-01-02"
	DEFAULT_EXT      = ".log"
)

type FileConfig struct {
	file 	 *os.File
	filename string
	dir      string
	size     int
	ext      string
}

// 获得默认配置
func DefaultFileConfig() *FileConfig {
	return &FileConfig{nil,DEFAULT_FILENAME,DEFAULT_DIR, DEFAULT_SIZE, DEFAULT_EXT}
}

// Generate log files based on configuration
func (file *FileConfig) InitFileLog() {

}

func (file *FileConfig) Write(content string) {
	// 1. 判断文件大小，是否需要拆分

}

// Open the log folder full path
func (file *FileConfig) getFullDirPath() string {
	dir, _ := os.Getwd()
	return dir + "/" + strings.Trim(file.dir, "/") + "/"
}

// Get the log file full path
func (file *FileConfig) getFullPath() string {
	dirPath := file.getFullDirPath()
	filename := time.Now().Format(file.filename) + file.ext
	return dirPath + filename
}

// Get log file
func (file *FileConfig) getFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		file.MkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to Open Log File :%v", err)
	}

	return handle
}

// Create log dir
func (file *FileConfig) MkDir() {
	path := file.getFullDirPath()
	fmt.Printf("create log dir:%s", path)
	os.MkdirAll(path, os.ModePerm)
}
