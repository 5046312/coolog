package coolog

import (
	"testing"
	"time"
)

func Test_Coolog_FileLog_Debug(t *testing.T) {
	conf := FileConfig()
	// conf.MaxSize = 500
	// conf.MaxTime = 1 // Hour
	conf.Single = true
	conf.Ext = ".bin"
	conf.Path = "./runtime/lll/"
	SetFileLog(conf)
	for {
		Debug("Write Debug in file")
		time.Sleep(time.Microsecond * 300)
	}
}

func Test_Coolog_Write_Directly(t *testing.T) {
	for {
		Debug("Write Debug in file")
		Debug("Write Debug in file")
		Debug("Write Debug in file")
		Debug("Write Debug in file")
		time.Sleep(time.Microsecond * 600)
	}
}

func Test_Init_TwiceTime(t *testing.T) {
	Debug("Write Debug in file")

	fc := FileConfig()
	SetFileLog(fc)
	//Debug("初始化Log配置")
	time.Sleep(time.Microsecond * 600)
}

func Test_Init(t *testing.T) {
	fc := FileConfig()
	SetFileLog(fc)
	//Debug("初始化Log配置")
	time.Sleep(time.Microsecond * 600)
}
