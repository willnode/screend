//go:build linux || darwin
// +build linux darwin

package screend

import (
	"os"
	"path/filepath"
)

const (
	ConfigPath = "/.config/screend/processes.yaml"
)

func getConfigPath() (path string) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	path = filepath.Join(dirname, ConfigPath)
	return
}
