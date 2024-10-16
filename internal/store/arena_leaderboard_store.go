package store

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/arena"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/log"
)

type ScoreData struct {
	Rank  int32
	Score int32
}

type LeaderboardStore interface {
	// Exists allows checking if the given leaderboard exists in the database.
	Exists(ctx context.Context) bool

	// GetLeaderboard returns a leaderboard with information about the players.
	GetLeaderboard(ctx context.Context, count int32) ([]*arena.ArenaLeaderboardPlayer, error)

	// AddMembers adds a list of players to the leaderboard.
	AddMembers(ctx context.Context, players ...*arena.ArenaLeaderboardPlayer) error

	// UpdateMemberData updates the player data required for the leaderboard.
	UpdateMemberData(ctx context.Context, accountID primitive.ObjectID, player *arena.ArenaMemberData, force bool) error

	// UpdateMembersData allows updating information for multiple players.
	// If a certain player is not in the table, he will be skipped.
	UpdateMembersData(ctx context.Context, players map[string]*arena.ArenaMemberData) error

	// MemberExists allows you to check if the given player is on the leaderboard.
	MemberExists(ctx context.Context, accountID string) bool

	// SetRating allows to specify the player's current rating in the leaderboard.
	SetRating(ctx context.Context, accountID primitive.ObjectID, rating int32) error

	// GetRank allows to get the player's current rank in the leaderboard (1-based).
	GetRank(ctx context.Context, id primitive.ObjectID) (int64, error)

	// GetMaxRating allows to get max rating in the leaderboard.
	GetMaxRating(ctx context.Context) (int32, error)

	// GetMembersByScoreRange returns a list of players with ratings in the specified range.
	GetMembersByScoreRange(ctx context.Context, min, max int32) ([]string, error)

	// DeleteMember delete all user information from leaderboard.
	DeleteMember(ctx context.Context, accountID primitive.ObjectID) error

	MembersCount(ctx context.Context) (int64, error)
}

const (
	leaderboardKey           = "locg.v1.arena:global-leaderboard:scores"
	leaderboardMemberDataKey = "locg.v1.arena:global-leaderboard:member_data"
)

type ArenaLeaderboardStore struct {
	rdb    *redis.Client
	config *config.Config
}

// NewArenaLeaderboardStore creates a new instance of the arena leaderboard store.
func NewArenaLeaderboardStore(config *config.Config, rdb *redis.Client) *ArenaLeaderboardStore {
	s := new(ArenaLeaderboardStore)
	s.config = config
	s.rdb = rdb
	return s
}

// Exists allows checking if the given leaderboard exists in the database.
func (s *ArenaLeaderboardStore) Exists(ctx context.Context) bool {
	res, _ := s.rdb.Exists(ctx, leaderboardKey).Result()
	return res != 0
}

// GetLeaderboard returns a leaderboard (100 players) with information about the players.
func (s *ArenaLeaderboardStore) GetLeaderboard(ctx context.Context, count int32) ([]*arena.ArenaLeaderboardPlayer, error) {
	scores, err := s.rdb.ZRangeArgsWithScores(ctx, redis.ZRangeArgs{
		Key:   leaderboardKey,
		Start: 0,
		Stop:  count,
		Rev:   true,
	}).Result()

	if err != nil {
		return nil, err
	}

	ranks := make(map[string]ScoreData)
	members := make([]string, 0, len(scores))

	for i, z := range scores {
		ranks[z.Member.(string)] = ScoreData{Rank: int32(i + 1), Score: int32(z.Score)}
		members = append(members, z.Member.(string))
	}

	membersData, err := s.rdb.HMGet(ctx, leaderboardMemberDataKey, members...).Result()
	if err != nil {
		return nil, err
	}

	players := make([]*arena.ArenaLeaderboardPlayer, 0, len(members))

	for i, member := range members {
		if membersData[i] == nil {
			continue
		}
		var player arena.ArenaLeaderboardPlayer
		err = json.Unmarshal([]byte(membersData[i].(string)), &player)

		if err != nil {
			log.Error(err)
			continue
		}

		player.ID = &base.ObjectID{Value: member}
		player.Rating = ranks[member].Score
		player.Rank = ranks[member].Rank
		players = append(players, &player)
	}

	return players, nil
}

// AddMembers adds a list of players to the leaderboard.
func (s *ArenaLeaderboardStore) AddMembers(ctx context.Context, players ...*arena.ArenaLeaderboardPlayer) error {
	members := make([]*redis.Z, 0, len(players))
	membersData := make([]string, 0, len(players))
	for _, player := range players {
		members = append(members, &redis.Z{
			Score:  float64(player.Rating),
			Member: player.ID.Value,
		})
		membersData = append(membersData, player.ID.Value)

		rank := player.Rank
		rating := player.Rating
		ID := player.ID

		player.Rank = 0
		player.Rating = 0
		player.ID = nil

		data, err := json.Marshal(player)

		player.Rank = rank
		player.Rating = rating
		player.ID = ID

		if err != nil {
			log.Error(err)
			continue
		}

		membersData = append(membersData, string(data))
	}

	s.rdb.ZAdd(ctx, leaderboardKey, members...)
	s.rdb.HMSet(ctx, leaderboardMemberDataKey, membersData)

	return nil
}

// MemberExists allows you to check if the given player is on the leaderboard.
func (s *ArenaLeaderboardStore) MemberExists(ctx context.Context, accountID string) bool {
	return s.rdb.HExists(ctx, leaderboardMemberDataKey, accountID).Val()
}

// UpdateMemberData updates the player data required for the leaderboard.
func (s *ArenaLeaderboardStore) UpdateMemberData(ctx context.Context, accountID primitive.ObjectID, player *arena.ArenaMemberData, force bool) error {
	if force || s.MemberExists(ctx, accountID.Hex()) {
		if data, err := json.Marshal(player); err == nil {
			s.rdb.HSet(ctx, leaderboardMemberDataKey, accountID.Hex(), data)
		}
	}
	return nil
}

// UpdateMembersData allows updating information for multiple players.
func (s *ArenaLeaderboardStore) UpdateMembersData(ctx context.Context, players map[string]*arena.ArenaMemberData) error {
	var membersData []string
	for id, player := range players {
		if s.MemberExists(ctx, id) {
			if data, err := json.Marshal(player); err == nil {
				membersData = append(membersData, id)
				membersData = append(membersData, string(data))
			}
		}
	}
	if len(membersData) > 0 {
		return s.rdb.HMSet(ctx, leaderboardMemberDataKey, membersData).Err()
	}
	return nil
}

// SetRating allows you to specify the player's current rating in the leaderboard.
func (s *ArenaLeaderboardStore) SetRating(ctx context.Context, accountID primitive.ObjectID, rating int32) error {
	return s.rdb.ZAdd(ctx, leaderboardKey, &redis.Z{
		Score:  float64(rating),
		Member: accountID.Hex(),
	}).Err()
}

// GetRank allows to get the player's current rank in the leaderboard (1-based).
func (s *ArenaLeaderboardStore) GetRank(ctx context.Context, accountID primitive.ObjectID) (int64, error) {
	res, err := s.rdb.ZRank(ctx, leaderboardKey, accountID.Hex()).Result()
	return res + 1, err
}

// GetMaxRating allows to get max rating in the leaderboard.
func (s *ArenaLeaderboardStore) GetMaxRating(ctx context.Context) (int32, error) {
	scores, err := s.rdb.ZRangeArgsWithScores(ctx, redis.ZRangeArgs{
		Key:   leaderboardKey,
		Start: 0,
		Stop:  1,
		Rev:   true,
	}).Result()

	if len(scores) > 0 {
		return int32(scores[0].Score), err
	}

	return 0, err
}

// GetMembersByScoreRange returns a list of players with ratings in the specified range.
func (s *ArenaLeaderboardStore) GetMembersByScoreRange(ctx context.Context, min, max int32) ([]string, error) {
	return s.rdb.ZRangeByScore(ctx, leaderboardKey, &redis.ZRangeBy{
		Min:   strconv.Itoa(int(min)),
		Max:   strconv.Itoa(int(max)),
		Count: 10,
	}).Result()
}

// DeleteMember delete all user information from leaderboard.
func (s *ArenaLeaderboardStore) DeleteMember(ctx context.Context, accountID primitive.ObjectID) error {
	err := s.rdb.ZRem(ctx, leaderboardKey, accountID.Hex()).Err()
	if err == nil {
		err = s.rdb.HDel(ctx, leaderboardMemberDataKey, accountID.Hex()).Err()
	}
	return err
}

func (s *ArenaLeaderboardStore) MembersCount(ctx context.Context) (int64, error) {
	return s.rdb.ZCard(ctx, leaderboardKey).Result()
}
