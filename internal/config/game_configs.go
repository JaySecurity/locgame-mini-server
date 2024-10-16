package config

import (
	"log"
	"reflect"

	"locgame-mini-server/pkg/dto/arena"
	"locgame-mini-server/pkg/dto/friends"
	gameDto "locgame-mini-server/pkg/dto/game"
	"locgame-mini-server/pkg/dto/resources"
	storeDto "locgame-mini-server/pkg/dto/store"
)

var container *Container

func Register(constructor interface{}) {
	if container == nil {
		container = NewConfigsContainer()
	}
	container.Register(constructor)
}

// GameConfigs contains matchmaking configurations.
type GameConfigs struct {
	Cards               *CardsConfig
	VirtualCards        *VirtualCardsConfig
	Metrics             *MetricsConfig
	AiProfiles          map[string]*AiProfile
	AiDecks             map[string]*AiDeck
	AiBots              map[string]*AiBot
	StoryMode           map[int32]*gameDto.StoryModeMissionData
	Leagues             map[int32]*arena.LeagueData
	Arena               *ArenaConfig
	TrophyRoadRewards   []*arena.TrophyRoadRewards
	ResettableResources map[int32]*resources.ResettableResource
	CappedResources     *CappedResources
	Blockchain          *Blockchain
	InGameStore         *InGameStore
	Products            *Products
	Aws                 *AwsConfig
	Withdraw            *WithdrawConfig
	MatchTimerConfig    *MatchTimerConfig
	FriendlyMatch       *friends.FriendlyMatchConfig
	DailyRewards        []*resources.ResourceAdjustment
	CognitoEnv          *CognitoConfig
	Ses                 *SesConfig
	ExtraRewards        *ExtraRewardsConfig
}

// Load allows to load and parse all YAML configuration files.
func (c *GameConfigs) Load() {
	configValue := reflect.ValueOf(c).Elem()
	configType := reflect.TypeOf(c).Elem()

	for i := 0; i < configValue.NumField(); i++ {
		if v, ok := configType.Field(i).Tag.Lookup("config"); ok && v == "ignore" {
			continue
		}
		field := configValue.Field(i)
		value := container.Get(field.Type())
		field.Set(value)
	}

	c.checkProducts()
}

func (c *GameConfigs) checkProducts() {
	for _, product := range c.Products.ProductsByID {
		if product.Type == storeDto.ProductType_PackOfCards {
			if c.InGameStore.PackByID[product.Value] == nil {
				log.Fatal("Product not found:", product.Value)
			}
		}
		if product.Type == storeDto.ProductType_PackOfCoins {
			if c.InGameStore.CoinsByID[product.Value] == nil {
				log.Fatal("Product not found:", product.Value)
			}
		}
		if product.Type == storeDto.ProductType_VToken {
			if c.InGameStore.TokensByID[product.Value] == nil {
				log.Fatal("Product not found:", product.Value)
			}
		}
	}
}
