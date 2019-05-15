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
	filename string
	dir      string
	size     int
	ext      string
}

// 获得默认配置
func DefaultFileConfig() *FileConfig {
	return &FileConfig{DEFAULT_DIR, DEFAULT_FILENAME, DEFAULT_SIZE, DEFAULT_EXT}
}

// 初始化file，创建文件
func (file *FileConfig) InitFileLog() {

}

func (file *FileConfig) Write(content string) {
	// 1. 判断文件大小，是否需要拆分

}

func (file *FileConfig) getFullPath() string {
	dir := strings.TrimRight(file.dir, "/") + "/"
	filename := time.Now().Format(file.filename) + file.ext
	return dir + filename
}

func (file *FileConfig) openFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		file.MkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

// create log dir
func (file *FileConfig) MkDir() {
	dir, _ := os.Getwd()
	logPath := dir + strings.TrimRight(dir, "/")
	fmt.Printf("create log dir:%s", logPath)
	fmt.Println(logPath)
	os.MkdirAll(logPath, os.ModePerm)
}
