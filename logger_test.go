package coolog

import "testing"

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
	conf.Max_size = 200
	log := NewFileLog(conf)
	log.Debug("Write Debug in file")
	log.Debug("Write Debug in file")
	log.Debug("Write Debug in file")
	log.Debug("Write Debug in file")
	log.Debug("Write Debug in file")
}
