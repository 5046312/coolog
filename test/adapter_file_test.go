package test

import (
	"coolog/adapter"
	"testing"
)

func Test_InitFileLog(t *testing.T) {
	adapter.DefaultFileConfig()
}

func Test_WriteFileLog(t *testing.T) {
	config := adapter.DefaultFileConfig()
	config.Write("Write Test Text")
	config.Write("Write Test Text")
	config.Write("Write Test Text")
	config.Write("Write Test Text")
	config.Write("Write Test Text")
	config.Write("Write Test Text")
	config.Write("Write Test Text")
	config.Write("Write Test Text")
}

func Test_SplitLogFile(t *testing.T) {
	config := adapter.DefaultFileConfig()
	config.Write("Write Test Text")
}
