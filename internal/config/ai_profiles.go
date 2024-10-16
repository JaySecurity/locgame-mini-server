package config

import (
	"os"
	"path/filepath"
	"strings"

	"locgame-mini-server/pkg/log"
)

func init() {
	Register(NewAiProfiles)
}

// AiProfiles stores cards configuration.
type AiProfiles struct {
	BaseConfig

	profileByID map[string]*AiProfile
}

// NewAiProfiles creates an instance of the ai profiles configuration.
func NewAiProfiles() map[string]*AiProfile {
	c := new(AiProfiles)
	c.self = c
	c.Load("ai/profiles")
	return c.profileByID
}

func (c *AiProfiles) Unmarshal() error {
	profileID := ""
	c.profileByID = make(map[string]*AiProfile)
	verbose := showLog
	if showLog {
		SetShowLog(false)
		defer SetShowLog(true)
	}
	err := filepath.Walk(c.filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != c.filePath {
			return filepath.SkipDir
		}

		if filepath.Ext(info.Name()) == yamlExt {
			profileID = strings.Split(info.Name(), ".")[0]
			if verbose {
				log.Debug("   -", profileID)
			}
			data := NewAiProfile()
			data.ignoreParent = true
			data.Load(path)
			c.profileByID[profileID] = data
		}
		return nil
	})

	return err
}
