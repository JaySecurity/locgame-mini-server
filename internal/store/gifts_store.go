package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/store"
	"locgame-mini-server/pkg/log"
)

// GiftsStore stores data for the storage of data about gifts.
type GiftsStore struct {
	db     *mongo.Client
	config *config.Config

	collection *mongo.Collection
}

type GiftsStoreInterface interface {
	Update(ctx context.Context, gift *store.Gift) error
	Get(ctx context.Context, giftID string) (*store.Gift, error)
	ClearErrorMessage(ctx context.Context, giftID *base.ObjectID)
}

// NewGiftsStore creates a new instance of the gifts store.
func NewGiftsStore(config *config.Config, db *mongo.Client) *GiftsStore {
	s := new(GiftsStore)
	s.config = config
	s.db = db
	s.collection = db.Database(s.config.Database.Database).Collection("gifts")

	if _, err := s.collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "created_at", Value: -1},
			},
		},
		{
			Keys: bson.D{
				{Key: "status", Value: -1},
			},
		},
	}); err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *GiftsStore) Update(ctx context.Context, gift *store.Gift) error {
	_, err := s.collection.UpdateByID(ctx, gift.ID, bson.M{"$set": gift})
	return err
}

func (s *GiftsStore) ClearErrorMessage(ctx context.Context, giftID *base.ObjectID) {
	_, _ = s.collection.UpdateByID(ctx, giftID, bson.M{"$unset": bson.M{"error": ""}})
}

func (s *GiftsStore) Get(ctx context.Context, giftID string) (*store.Gift, error) {
	var gift *store.Gift
	id, _ := primitive.ObjectIDFromHex(giftID)
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&gift)
	return gift, err
}
