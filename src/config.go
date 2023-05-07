package screend

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Process struct {
	Name    string            `yaml:"name"`
	Command string            `yaml:"command"`
	Args    []string          `yaml:"args"`
	Dir     string            `yaml:"dir"`
	Env     map[string]string `yaml:"env"`

	// internal
	ScreenID string `yaml:"-"`
}

type Config struct {
	Version int `yaml:"version"`
	Bin     struct {
		Screen string `yaml:"screen"`
	} `yaml:"bin"`
	Processes []Process `yaml:"processes"`
}

func readConfig() (config Config, err error) {
	// read config file of YAML from ~/.config/screend/processes.yaml
	config = Config{}

	path := getConfigPath()
	if _, err = os.Stat(path); err != nil {
		config.Version = 1
		config.Bin.Screen, err = exec.LookPath("screen")
		if err != nil {
			return
		}
		err = writeConfig(config)
		if err != nil {
			return
		}
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return
	}

	// parse YAML
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return
	}

	if _, err = os.Stat(config.Bin.Screen); err != nil {
		config.Bin.Screen, err = exec.LookPath("screen")
		if err != nil {
			return
		}
	}

	// check version
	if config.Version != 1 {
		err = errors.New("invalid config version, please update screend")
		return
	}

	return
}

func writeConfig(config Config) (err error) {
	// write config file

	bytes, err := yaml.Marshal(config)
	if err != nil {
		return
	}

	path := getConfigPath()
	err = os.MkdirAll(filepath.Dir(path), 0750)
	if err != nil {
		return
	}
	err = os.WriteFile(path, bytes, 0640)
	if err != nil {
		return
	}

	return
}
