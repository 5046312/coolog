package adapter

import (
	"testing"
)

func Test_InitFileLog(t *testing.T) {
	DefaultFileConfig()
}

func Test_WriteFileLog(t *testing.T) {
	config := DefaultFileConfig()
	config.Write("Write Test Text1")
	config.Write("Write Test Text2")
	config.Write("Write Test Text3")
	config.Write("Write Test Text4")
	config.Write("Write Test Text5")
	config.Write("Write Test Text6")
	config.Write("Write Test Text7")
	config.Write("Write Test Text8")
}

func Test_SplitLogFile(t *testing.T) {
	config := DefaultFileConfig()
	config.Write("Write Test Text")
}
func Test_GetAllFiles(t *testing.T) {
	config := DefaultFileConfig()
	config.getAllFiles()
}

func Test_GetTodayFiles(t *testing.T) {
	config := DefaultFileConfig()
	config.getTodayFiles()
}
