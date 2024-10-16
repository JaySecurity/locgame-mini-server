package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/log"
)

// MintStore stores data for the storage of data about mints.
type MintStore struct {
	db     *mongo.Client
	config *config.Config

	collection *mongo.Collection
}

type MintStoreInterface interface {
	GetMintedCardsCount(ctx context.Context, cardID string) (int64, error)
	IncrementMintedCards(ctx context.Context, cardID string) (int64, error)
	SetMintedCards(ctx context.Context, cardID string, value int64) error
}

// NewMintStore creates a new instance of the mint store.
func NewMintStore(config *config.Config, db *mongo.Client) *MintStore {
	s := new(MintStore)
	s.config = config
	s.db = db
	s.collection = db.Database(s.config.Database.Database).Collection("minted_cards")

	if _, err := s.collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "card_id", Value: -1},
			},
		},
	}); err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *MintStore) GetMintedCardsCount(ctx context.Context, cardID string) (int64, error) {
	var data struct {
		CardID string `bson:"card_id"`
		Count  int64  `bson:"count"`
	}
	err := s.collection.FindOne(ctx, bson.M{"card_id": cardID}).Decode(&data)
	if err == mongo.ErrNoDocuments {
		return 0, nil
	}

	return data.Count, err
}

func (s *MintStore) IncrementMintedCards(ctx context.Context, cardID string) (int64, error) {
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	var data struct {
		CardID string `bson:"card_id"`
		Count  int64  `bson:"count"`
	}
	err := s.collection.FindOneAndUpdate(ctx, bson.M{"card_id": cardID}, bson.D{{"$inc", bson.D{{"count", 1}}}}, opts).Decode(&data)
	return data.Count, err
}

func (s *MintStore) SetMintedCards(ctx context.Context, cardID string, value int64) error {
	opts := options.Update().SetUpsert(true)
	_, err := s.collection.UpdateOne(ctx, bson.M{"card_id": cardID}, bson.M{"$set": bson.M{"count": value}}, opts)
	return err
}
