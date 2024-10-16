package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/player"
)

// DailyRewardsStore stores data for the operation of the alchemy laboratory store.
type DailyRewardsStore struct {
	db     *mongo.Client
	config *config.Config

	collection *mongo.Collection
}

// DailyRewardsStoreInterface provides an interface for the ability to later mock
//
//go:generate mockery -name DailyRewardsStoreInterface
type DailyRewardsStoreInterface interface {
	Get(ctx context.Context, accountID primitive.ObjectID) (*player.DailyRewardData, error)
	Update(ctx context.Context, data *player.DailyRewardData) error
}

// NewDailyRewardsStore creates a new instance of the daily rewards store.
func NewDailyRewardsStore(config *config.Config, db *mongo.Client) *DailyRewardsStore {
	s := new(DailyRewardsStore)
	s.db = db
	s.config = config

	s.collection = db.Database(s.config.Database.Database).Collection("daily_rewards")

	return s
}

func (s *DailyRewardsStore) Get(ctx context.Context, accountID primitive.ObjectID) (*player.DailyRewardData, error) {
	var data player.DailyRewardData
	err := s.collection.FindOne(ctx, bson.M{"_id": accountID}).Decode(&data)
	if err == mongo.ErrNoDocuments {
		data.ID = &base.ObjectID{Value: accountID.Hex()}
		data.Counter = 1
		_, err = s.collection.InsertOne(ctx, &data)
	}
	return &data, err
}

func (s *DailyRewardsStore) Update(ctx context.Context, data *player.DailyRewardData) error {
	_, err := s.collection.UpdateByID(ctx, data.ID, bson.M{"$set": data})
	return err
}
