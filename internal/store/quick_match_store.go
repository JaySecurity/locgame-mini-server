package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/game"
	"locgame-mini-server/pkg/log"
)

// QuickMatchStore stores data for the operation of the QuickMatch store.
type QuickMatch struct {
	redis  *redis.Client
	config *config.Config
}

// QuickMatchStoreInterface provides an interface for the ability to later mock
//
//go:generate mockery -name QuickMatchStoreInterface
type QuickMatchStoreInterface interface {
	AddPlayer(ctx context.Context, accountID *base.ObjectID, gameType game.GameType, stake int32) error
	RemovePlayer(ctx context.Context, accountID *base.ObjectID, gameType game.GameType) error
	FindOpponent(ctx context.Context, mmr float32) string
	GetOpponentList(ctx context.Context) (map[string]string, error)
}

const (
	quickMatchMatchmakingKeyHashMAP = "locg.v1.matchmaking:quick-matchmakingHSET"
)

// NewQuickMatchStore creates a new instance of the quickMatch store.
func NewQuickMatchStore(config *config.Config, redis *redis.Client) *QuickMatch {
	s := new(QuickMatch)
	s.redis = redis
	s.config = config

	return s
}

func (s *QuickMatch) AddPlayer(ctx context.Context, accountID *base.ObjectID, gameType game.GameType, stake int32) error {
	fmt.Println("add player ", accountID, " ", gameType)
	switch gameType {
	case game.GameType_QuickMatch:
		res := s.redis.HSet(ctx, quickMatchMatchmakingKeyHashMAP, accountID.Value, 0)
		return res.Err()
	case game.GameType_QuickMatchWithStake:
		res := s.redis.HSet(ctx, quickMatchMatchmakingKeyHashMAP, accountID.Value, stake)
		return res.Err()
	default:
		return errors.New("not implemented")
	}
}

func (s *QuickMatch) RemovePlayer(ctx context.Context, accountID *base.ObjectID, gameType game.GameType) error {
	switch gameType {
	case game.GameType_QuickMatch, game.GameType_QuickMatchWithStake:
		res := s.redis.HDel(ctx, quickMatchMatchmakingKeyHashMAP, accountID.Value)
		return res.Err()
	default:
		return errors.New("not implemented")
	}
}

func (s *QuickMatch) FindOpponent(ctx context.Context, mmr float32) string {

	//result, err := s.redis.HGet(ctx, quickMatchMatchmakingKeyHashMAP, "account1").Result()
	result, err := s.redis.HGetAll(ctx, quickMatchMatchmakingKeyHashMAP).Result()
	//result, err := s.redis.HKeys(ctx, quickMatchMatchmakingKeyHashMAP).Result()
	if err != nil {
		log.Debug("failed to remove element from Redis: %w", err)
		return ""
	}
	fmt.Println(result)
	return "result"
}

func (s *QuickMatch) GetOpponentList(ctx context.Context) (map[string]string, error) {
	result, err := s.redis.HGetAll(ctx, quickMatchMatchmakingKeyHashMAP).Result()
	if err != nil {
		log.Debug("failed to GetOpponentList from Redis: %w", err)
		return nil, err
	}

	return result, nil
}
