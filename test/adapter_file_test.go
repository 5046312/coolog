package test

import (
	"coolog/adapter"
	"testing"
)


func Test_WriteFileLog(t *testing.T){
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
