package coolog

import "coolog/adapter"

type Level int

const (
	LEVEL_ERROR Level = iota
	LEVEL_WARNING
	LEVEL_NOTICE
	LEVEL_INFO
	LEVEL_DEBUG
)

var LevelTag = map[Level]string{
	LEVEL_ERROR:   "Error",
	LEVEL_WARNING: "Warning",
	LEVEL_NOTICE:  "Notice",
	LEVEL_INFO:    "Info",
	LEVEL_DEBUG:   "Debug",
}

type Coolog struct {
	adapter adapter.Adapter
}

// create a file log
func NewFileLog(conf *adapter.FileConfig) *Coolog {
	return &Coolog{
		adapter: adapter.Adapter{conf },
	}
}

// Including multiple adapters working at the same time
func (log *Config) Write(content string){

	
}
