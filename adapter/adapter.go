package adapter

type Adapter struct {
	File *FileConfig
}

func NewFileAdapter(fc *FileConfig) *Adapter {
	return &Adapter{
		File: fc,
	}
}
