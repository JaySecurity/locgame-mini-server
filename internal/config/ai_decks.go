package config

import (
	"os"
	"path/filepath"
	"strings"

	"locgame-mini-server/pkg/log"
)

func init() {
	Register(NewAiDecks)
}

// AiDecks stores ai decks configuration.
type AiDecks struct {
	BaseConfig

	deckByID map[string]*AiDeck
}

// NewAiDecks creates an instance of the ai decks configuration.
func NewAiDecks() map[string]*AiDeck {
	c := new(AiDecks)
	c.self = c
	c.Load("ai/decks")
	return c.deckByID
}

func (c *AiDecks) Unmarshal() error {
	deckID := ""
	c.deckByID = make(map[string]*AiDeck)
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
			deckID = strings.Split(info.Name(), ".")[0]
			if verbose {
				log.Debug("   -", deckID)
			}
			data := NewAiDeck()
			data.ignoreParent = true
			data.Load(path)
			c.deckByID[deckID] = data
		}
		return nil
	})

	return err
}
