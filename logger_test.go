package coolog

import (
	"testing"
	"time"
)

func Test_Coolog_FileLog_Write(t *testing.T) {
	conf := FileConfig()
	conf.Max_size = 200
	log := NewFileLog(conf)
	log.Write("Write in file")
	log.Write("Write in file")
	log.Write("Write in file")
	log.Write("Write in file")
	log.Write("Write in file")
	log.Write("Write in file")
	log.Write("Write in file")
	log.Write("Write in file")
	log.Write("Write in file")
	log.Write("Write in file")
}

func Test_Coolog_FileLog_Debug(t *testing.T) {
	conf := FileConfig()
	conf.Max_size = 500
	conf.Max_time = 100
	log := NewFileLog(conf)
	for {
		log.Debug("Write Debug in file")
		time.Sleep(time.Second * 1)
	}
}
