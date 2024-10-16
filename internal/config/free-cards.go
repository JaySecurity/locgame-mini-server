package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"locgame-mini-server/pkg/dto/cards"
)

func init() {
	Register(NewFreeCardsConfig)
}

// FreeCardsConfig stores FreeCards configuration.
type FreeCardsConfig struct {
	BaseConfig

	freeCardsWithoutNumbers map[string][]string
	allFreeCardsIDs         []string

	freeCardsByID map[string]*cards.Card
	Pack          string
	DeckID        string
}

// NewFreeCardsConfig creates an instance of the FreeCards' configuration.
func NewFreeCardsConfig() *FreeCardsConfig {
	c := new(FreeCardsConfig)
	c.self = c
	c.freeCardsWithoutNumbers = make(map[string][]string)
	c.Load("free-cards.yaml")
	return c
}

func (c *FreeCardsConfig) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	data := struct {
		Pack   string        `yaml:"Pack"`
		DeckID string        `yaml:"DeckID"`
		Cards  []*cards.Card `yaml:"Cards"`
	}{}
	if err == nil {
		if err = yaml.Unmarshal(bytes, &data); err == nil {
			c.freeCardsByID = make(map[string]*cards.Card)
			c.Pack = data.Pack
			c.DeckID = data.DeckID
			for _, card := range data.Cards {
				c.freeCardsByID[card.ArchetypeID] = card
				c.allFreeCardsIDs = append(c.allFreeCardsIDs, card.ArchetypeID)
				c.freeCardsWithoutNumbers[card.ArchetypeID[:len(card.ArchetypeID)-4]] = append(c.freeCardsWithoutNumbers[card.ArchetypeID[:len(card.ArchetypeID)-4]], card.ArchetypeID)
			}
		}
	}

	return err
}

func (c *FreeCardsConfig) GetAllFreeCardsIDs() []string {
	return c.allFreeCardsIDs
}
