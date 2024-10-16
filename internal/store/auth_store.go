package store

import (
	"context"
	"encoding/json"
	"time"

	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/service/accounts/siwe"
	"locgame-mini-server/pkg/dto/errors"
	"locgame-mini-server/pkg/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AuthStore stores data for the operation of the authentication store.
type AuthStore struct {
	db     *mongo.Client
	redis  *redis.Client
	config *config.Config

	collection *mongo.Collection
}

type challengeRecord struct {
	Email  string `json:"email"`
	Code   string `json:"code"`
	Wallet string `json:"wallet"`
}

// AuthStoreInterface provides an interface for the ability to later mock
//
//go:generate mockery -name AuthStoreInterface
type AuthStoreInterface interface {
	GetChallenge(ctx context.Context, wallet string) (*siwe.Message, error)
	AddChallenge(ctx context.Context, walletAddress string, message *siwe.Message) error
	DeleteChallenge(ctx context.Context, wallet string)
	AddEmailChallenge(ctx context.Context, email string, code string, wallet string) error
	GetEmailChallenge(ctx context.Context, email string) (*challengeRecord, error)
}

// NewAuthStore creates a new instance of the auth store.
func NewAuthStore(config *config.Config, db *mongo.Client, redis *redis.Client) *AuthStore {
	s := new(AuthStore)
	s.db = db
	s.config = config
	s.redis = redis

	s.collection = db.Database(s.config.Database.Database).Collection("auth_challenges")

	if _, err := s.collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "address", Value: 1},
		},
		Options: options.Index().SetUnique(true).SetExpireAfterSeconds(300),
	}); err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *AuthStore) GetChallenge(ctx context.Context, wallet string) (*siwe.Message, error) {
	var data siwe.Message
	err := s.collection.FindOne(ctx, bson.M{"address": wallet}).Decode(&data)
	if err == mongo.ErrNoDocuments {
		return nil, errors.ErrAuthChallengeNotFound
	}

	return &data, err
}

func (s *AuthStore) AddChallenge(ctx context.Context, walletAddress string, message *siwe.Message) error {
	opts := options.Update().SetUpsert(true)
	_, err := s.collection.UpdateOne(ctx, bson.M{"address": common.HexToAddress(walletAddress)}, bson.M{"$set": message}, opts)
	return err
}

func (s *AuthStore) DeleteChallenge(ctx context.Context, walletAddress string) {
	_, _ = s.collection.DeleteOne(ctx, bson.M{"address": common.HexToAddress(walletAddress)})
}

func (s *AuthStore) AddEmailChallenge(ctx context.Context, email string, code string, wallet string) error {
	record, err := json.Marshal(challengeRecord{
		Email:  email,
		Code:   code,
		Wallet: wallet,
	})
	if err != nil {
		return err
	}
	err = s.redis.Set(ctx, email, record, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthStore) GetEmailChallenge(ctx context.Context, email string) (*challengeRecord, error) {

	val, err := s.redis.Get(ctx, email).Bytes() // Convert val from string to []byte
	if err != nil {
		return nil, err
	}
	var record = challengeRecord{}
	err = json.Unmarshal(val, &record)
	if err != nil {
		return nil, err
	}
	return &record, nil
}
