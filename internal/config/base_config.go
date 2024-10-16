package config

import (
	"path/filepath"

	"locgame-mini-server/pkg/log"
)

const yamlExt = ".yaml"

var showLog = true
var parentPath string

func SetShowLog(value bool) {
	showLog = value
}

// IConfig contains basic methods for all configurations.
type IConfig interface {
	Unmarshal() error
	Load(path string)
}

// BaseConfig contains basic configuration information.
type BaseConfig struct {
	self IConfig `yaml:"-"`

	filePath     string `yaml:"-"`
	ignoreParent bool
}

// Load loads the config file.
func (c *BaseConfig) Load(path string) {
	c.SetPath(path)

	if err := c.self.Unmarshal(); err != nil {
		panic(err)
	}
}

func (c *BaseConfig) SetPath(path string) {
	if !c.ignoreParent {
		c.filePath = filepath.Join(parentPath, path)
	} else {
		c.filePath = path
	}
	if showLog {
		log.Debug(" - " + c.filePath)
	}
}
