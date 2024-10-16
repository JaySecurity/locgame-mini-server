package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	storeDto "locgame-mini-server/pkg/dto/store"
	"locgame-mini-server/pkg/log"

	"gopkg.in/yaml.v3"
)

func init() {
	Register(NewInGameStore)
}

// InGameStore stores in-game in_game_store configuration.
type InGameStore struct {
	BaseConfig

	PackByID   map[string]*storeDto.Pack
	CoinsByID  map[string]*storeDto.CoinsPack
	TokensByID map[string]*storeDto.Token
}

// NewInGameStore creates an instance of the in-game in_game_store configuration.
func NewInGameStore() *InGameStore {
	c := new(InGameStore)
	c.self = c
	c.Load("store/packs")
	return c
}
func (c *InGameStore) Unmarshal() error {
	packID := ""
	c.PackByID = make(map[string]*storeDto.Pack)
	c.CoinsByID = make(map[string]*storeDto.CoinsPack)
	c.TokensByID = make(map[string]*storeDto.Token)
	verbose := showLog
	if showLog {
		SetShowLog(false)
	}
	err := filepath.Walk(c.filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != c.filePath {
			return filepath.SkipDir
		}

		if filepath.Ext(info.Name()) == yamlExt {
			packID = strings.Split(info.Name(), ".")[0]
			if verbose {
				log.Debug("   -", packID)
			}
			data := NewPack()
			data.ignoreParent = true
			data.Load(path)
			data.Pack.ID = packID
			c.PackByID[packID] = data.Pack
		}
		return nil
	})

	SetShowLog(true)

	if err == nil {
		c.SetPath("store/coins.yaml")
		var bytes []byte
		bytes, err = ioutil.ReadFile(c.filePath)
		if err != nil {
			return err
		}
		var data map[string]int32
		err = yaml.Unmarshal(bytes, &data)
		if err == nil {
			for id, value := range data {
				c.CoinsByID[id] = &storeDto.CoinsPack{
					ID:    id,
					Count: value,
				}
			}
		}
	}
	if err == nil {
		c.SetPath("store/tokens.yaml")
		var bytes []byte
		bytes, err = ioutil.ReadFile(c.filePath)
		if err != nil {
			return err
		}
		var data map[string]*storeDto.Token
		err = yaml.Unmarshal(bytes, &data)
		if err == nil {
			for id, value := range data {
				c.TokensByID[id] = &storeDto.Token{
					ID:                          id,
					TokenID:                     value.TokenID,
					Name:                        value.Name,
					Available:                   value.Available,
					QtyPerUnit:                  value.QtyPerUnit,
					MaxSupply:                   value.MaxSupply,
					MaxPurchase:                 value.MaxPurchase,
					MinPurchase:                 value.MinPurchase,
					SaleStart:                   value.SaleStart,
					SaleEnd:                     value.SaleEnd,
					LOCGBonus:                   value.LOCGBonus,
					AdditionalAccountedQuantity: value.AdditionalAccountedQuantity,
					Offers:                      value.Offers,
					PromoCodes:                  value.PromoCodes,
				}
			}
		}
	}

	return err
}
