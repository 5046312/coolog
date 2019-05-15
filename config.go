package coolog

type Adapter string

const (
	Adapter_File Adapter = "file"
)

const (
	DEFAULT_ADAPTER = Adapter_File
	DEFAULT_LEVEL   = LEVEL_DEBUG
)

type Config struct {
	adapter Adapter
	level   Level
}

// 获得默认配置
func defaultConfig() *Config {
	return &Config{DEFAULT_ADAPTER, DEFAULT_LEVEL}
}
