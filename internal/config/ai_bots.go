package config

import (
	"os"
	"path/filepath"
	"strings"

	"locgame-mini-server/pkg/log"
)

func init() {
	Register(NewAiBots)
}

// AiBots stores ai bots configuration.
type AiBots struct {
	BaseConfig

	botById map[string]*AiBot
}

// NewAiBots creates an instance of the ai bots configuration.
func NewAiBots() map[string]*AiBot {
	c := new(AiBots)
	c.self = c
	c.Load("ai/bots")
	return c.botById
}

func (c *AiBots) Unmarshal() error {
	botID := ""
	c.botById = make(map[string]*AiBot)
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
			botID = strings.Split(info.Name(), ".")[0]
			if verbose {
				log.Debug("   -", botID)
			}
			data := NewAiBot()
			data.ignoreParent = true
			data.Load(path)
			c.botById[botID] = data
		}
		return nil
	})

	return err
}
