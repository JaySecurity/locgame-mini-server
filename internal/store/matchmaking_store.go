package store

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/game"
)

// MatchmakingStore stores data for the operation of the matchmaking store.
type MatchmakingStore struct {
	redis  *redis.Client
	config *config.Config
}

// MatchmakingStoreInterface provides an interface for the ability to later mock
//
//go:generate mockery -name MatchmakingStoreInterface
type MatchmakingStoreInterface interface {
	AddPlayer(ctx context.Context, accountID *base.ObjectID, mmr float32, gameType game.GameType) error
	RemovePlayer(ctx context.Context, accountID *base.ObjectID, gameType game.GameType) error
	FindOpponent(ctx context.Context, mmr float32, gameType game.GameType) string
}

const (
	quickMatchMatchmakingKey = "locg.v1.matchmaking:quick-matchmaking"
)

// NewMatchmakingStore creates a new instance of the matchmaking store.
func NewMatchmakingStore(config *config.Config, redis *redis.Client) *MatchmakingStore {
	s := new(MatchmakingStore)
	s.redis = redis
	s.config = config

	return s
}

func (s *MatchmakingStore) AddPlayer(ctx context.Context, accountID *base.ObjectID, mmr float32, gameType game.GameType) error {
	switch gameType {
	case game.GameType_QuickMatch:
		res := s.redis.RPush(ctx, quickMatchMatchmakingKey, accountID.Value)
		return res.Err()
	default:
		return errors.New("not implemented")
	}
}

func (s *MatchmakingStore) RemovePlayer(ctx context.Context, accountID *base.ObjectID, gameType game.GameType) error {
	switch gameType {
	case game.GameType_QuickMatch:
		res := s.redis.LRem(ctx, quickMatchMatchmakingKey, 0, accountID.Value)
		return res.Err()
	default:
		return errors.New("not implemented")
	}
}

func (s *MatchmakingStore) FindOpponent(ctx context.Context, mmr float32, gameType game.GameType) string {
	return s.redis.LPop(ctx, quickMatchMatchmakingKey).Val() // TODO
}
