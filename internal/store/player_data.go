package store

import (
	"context"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/cards"
	"locgame-mini-server/pkg/dto/player"
	"locgame-mini-server/pkg/dto/tutorial"

	"go.mongodb.org/mongo-driver/bson/primitive"

	storeDto "locgame-mini-server/pkg/dto/store"
)

type PlayerData struct {
	baseData
	*player.PlayerData

	store  PlayersStoreInterface
	config *config.Config
}

func NewPlayerData(store PlayersStoreInterface, accountID primitive.ObjectID, cfg *config.Config) *PlayerData {
	d := new(PlayerData)
	d.init(accountID)
	d.store = store
	d.config = cfg

	return d
}

func (d *PlayerData) Update(ctx context.Context) error {
	data, err := d.store.GetPlayerDataByID(ctx, d.accountID)
	if err == nil {
		d.PlayerData = data

		if d.PlayerData.TutorialData == nil {
			d.PlayerData.TutorialData = &tutorial.TutorialData{}
		}

		if d.PlayerData.Resources == nil {
			d.PlayerData.Resources = make(map[int32]int32)
		}

		if d.PlayerData.Decks == nil || d.PlayerData.Decks.Decks == nil {
			d.PlayerData.Decks = &cards.Decks{
				Decks:   make(map[string]*cards.Deck),
				Active:  "",
				Defense: "",
			}
		}

		if d.PlayerData.PlayerStoreData == nil {
			d.PlayerData.PlayerStoreData = &storeDto.PlayerStoreData{PurchasedProducts: make(map[string]int32)}
		}

		d.ID = &base.ObjectID{Value: d.accountID.Hex()}
	}
	return err
}
