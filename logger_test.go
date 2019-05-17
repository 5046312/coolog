package coolog

import (
	"testing"
	"time"
)

func Test_Coolog_FileLog_Debug(t *testing.T) {
	conf := FileConfig()
	conf.MaxSize = 500
	conf.MaxTime = 1 // Hour
	conf.Single = false
	log := NewFileLog(conf)
	// for {
	log.Debug("Write Debug in file")
	time.Sleep(time.Microsecond * 300)
	// }
}
