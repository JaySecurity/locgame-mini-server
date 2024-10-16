package config

import (
	"fmt"
	"io/ioutil"
	"math/rand"

	"gopkg.in/yaml.v3"
	"locgame-mini-server/pkg/dto/cards"
	"locgame-mini-server/pkg/log"
)

func init() {
	Register(NewCardsConfig)
}

// CardsConfig stores cards configuration.
type CardsConfig struct {
	BaseConfig

	cardsWithoutNumbers map[string][]string
	allCardIDs          []string

	CardsByID map[string]*cards.Card
	Sets      map[int32]string
	Packs     map[int32]string
}

// NewCardsConfig creates an instance of the cards' configuration.
func NewCardsConfig() *CardsConfig {
	c := new(CardsConfig)
	c.self = c
	c.cardsWithoutNumbers = make(map[string][]string)
	c.Load("cards.yaml")
	return c
}

func (c *CardsConfig) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	data := struct {
		Sets  map[int32]string `yaml:"Sets"`
		Packs map[int32]string `yaml:"Packs"`
		Cards []*cards.Card    `yaml:"Cards"`
	}{}
	if err == nil {
		if err = yaml.Unmarshal(bytes, &data); err == nil {
			c.CardsByID = make(map[string]*cards.Card)
			c.Packs = data.Packs
			c.Sets = data.Sets
			for _, card := range data.Cards {
				c.CardsByID[card.ArchetypeID] = card
				c.allCardIDs = append(c.allCardIDs, card.ArchetypeID)
				c.cardsWithoutNumbers[card.ArchetypeID[:len(card.ArchetypeID)-4]] = append(c.cardsWithoutNumbers[card.ArchetypeID[:len(card.ArchetypeID)-4]], card.ArchetypeID)
				if card.Image == "" {
					log.Error(card.ArchetypeID)
				}
			}
		}
	}

	return err
}

func (c *CardsConfig) FindRandomCard(set int32, pack int32, gameRarity cards.GameRarity, visualRarity cards.VisualRarity) *cards.Card {
	query := fmt.Sprintf("%03d-%03d-%03d-%03d", set, pack, gameRarity, visualRarity)
	result := c.cardsWithoutNumbers[query]
	if len(result) > 0 {
		return c.CardsByID[result[rand.Intn(len(result))]]
	}
	return nil
}

func (c *CardsConfig) GetAllCardIDs() []string {
	return c.allCardIDs
}
