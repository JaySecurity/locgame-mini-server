package config

import (
	"locgame-mini-server/pkg/dto/cards"
	"locgame-mini-server/pkg/dto/errors"
	"os"

	"gopkg.in/yaml.v3"
)

func init() {
	Register(NewVirtualCardsConfig)
}

// VirtualCardsConfig stores FreeCards configuration.
type VirtualCardsConfig struct {
	BaseConfig

	// freeCardsWithoutNumbers map[string][]string
	allFreeCardsIDs []string

	CardsByID map[string]*cards.VirtualCard
	Pack      string
	DeckID    string
}

// NewFreeCardsConfig creates an instance of the FreeCards' configuration.
func NewVirtualCardsConfig() *VirtualCardsConfig {
	c := new(VirtualCardsConfig)
	c.self = c
	// c.freeCardsWithoutNumbers = make(map[string][]string)
	c.Load("virtual_cards.yaml")
	return c
}

func (c *VirtualCardsConfig) Unmarshal() error {
	bytes, err := os.ReadFile(c.filePath)

	data := struct {
		Pack   string               `yaml:"Pack"`
		DeckID string               `yaml:"DeckID"`
		Cards  []*cards.VirtualCard `yaml:"VirtualCards"`
	}{}
	if err == nil {
		if err = yaml.Unmarshal(bytes, &data); err == nil {
			c.CardsByID = make(map[string]*cards.VirtualCard)
			c.Pack = data.Pack
			c.DeckID = data.DeckID
			for _, card := range data.Cards {
				c.CardsByID[card.ArchetypeID] = card
				if card.ArchetypeID[:7] == "999-999" {
					c.allFreeCardsIDs = append(c.allFreeCardsIDs, card.ArchetypeID)
				}
				// c.freeCardsWithoutNumbers[card.ArchetypeID[:len(card.ArchetypeID)-4]] = append(c.freeCardsWithoutNumbers[card.ArchetypeID[:len(card.ArchetypeID)-4]], card.ArchetypeID)
			}
		}
	}

	return err
}

func (c *VirtualCardsConfig) GetAllFreeCardsIDs() []string {
	return c.allFreeCardsIDs
}

func (c *VirtualCardsConfig) GetVirtualCards() map[string]*cards.VirtualCard {
	return c.CardsByID
}

func (c *VirtualCardsConfig) GetCardById(cardId string) (*cards.VirtualCard, error) {
	card, ok := c.CardsByID[cardId]
	if ok {
		return card, nil
	} else {
		return nil, errors.ErrInvalidCard
	}
}
