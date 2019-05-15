package test

import (
	"coolog/adapter"
	"testing"
)

func TestMkFileDir(t *testing.T) {
	config := adapter.DefaultFileConfig()
	config.MkDir()
}
