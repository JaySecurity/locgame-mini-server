package store

import (
	"context"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/arena"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/cards"
	"locgame-mini-server/pkg/dto/player"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/stime"
)

type ArenaStore struct {
	db     *mongo.Client
	config *config.Config

	collection           *mongo.Collection
	battleLogsCollection *mongo.Collection
}

type ArenaStoreInterface interface {
	GetLeaderboard(ctx context.Context, count int32) ([]*arena.ArenaLeaderboardPlayer, error)
	GetOpponentData(ctx context.Context, accountID primitive.ObjectID) (*arena.ArenaPlayerData, error)

	GetRating(ctx context.Context, accountID primitive.ObjectID) (int32, int32, error)
	SetRating(ctx context.Context, accountID primitive.ObjectID, rating int32, maxRating int32) error

	SaveBattleLog(ctx context.Context, data *arena.ArenaBattleLog) error
	GetBattleLogs(ctx context.Context, timestamp *base.Timestamp, accountID primitive.ObjectID) ([]*arena.ArenaBattleLog, error)
	DeleteBattleLogs(ctx context.Context, accountID primitive.ObjectID) error
}

// NewArenaStore creates a new instance of the arena store.
func NewArenaStore(config *config.Config, db *mongo.Client) *ArenaStore {
	s := new(ArenaStore)
	s.config = config
	s.db = db
	s.collection = db.Database(s.config.Database.Database).Collection("players")
	s.battleLogsCollection = db.Database(s.config.Database.Database).Collection("battle_logs")

	if _, err := s.collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "arena_data.rating", Value: -1},
		},
	}); err != nil {
		log.Fatal(err)
	}

	if _, err := s.collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "arena_data.state", Value: -1},
		},
	}); err != nil {
		log.Fatal(err)
	}

	if _, err := s.battleLogsCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "player._id", Value: 1},
		},
	}); err != nil {
		log.Fatal(err)
	}

	if _, err := s.battleLogsCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "opponent._id", Value: 1},
		},
	}); err != nil {
		log.Fatal(err)
	}

	if _, err := s.battleLogsCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "created_at", Value: -1},
		},
		Options: options.Index().SetExpireAfterSeconds(604800),
	}); err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *ArenaStore) GetLeaderboard(ctx context.Context, count int32) ([]*arena.ArenaLeaderboardPlayer, error) {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "avatar_id", Value: 1},
		{Key: "arena_data.rating", Value: 1},
	}

	findOptions := options.Find()
	if count > 0 {
		findOptions.SetLimit(int64(count))
	}
	findOptions.SetProjection(projection)
	findOptions.SetSort(bson.M{"arena_data.rating": -1})
	findOptions.SetHint("arena_data.rating_-1")

	cursor, err := s.collection.Find(ctx, bson.M{"arena_data.state": bson.M{"$gt": 0}}, findOptions)

	if err != nil {
		return nil, err
	}

	rank := 1
	playersMap := make(map[string]*arena.ArenaLeaderboardPlayer)

	for cursor.Next(ctx) {
		var data *player.PlayerData
		err = cursor.Decode(&data)
		if err != nil {
			log.Error(err)
			continue
		}
		playersMap[data.ID.Value] = &arena.ArenaLeaderboardPlayer{
			ID:       data.ID,
			Rank:     int32(rank),
			AvatarID: data.AvatarID,
			Name:     data.Name,
			Rating:   data.ArenaData.Rating,
		}
		rank++
	}
	_ = cursor.Close(ctx)

	players := make([]*arena.ArenaLeaderboardPlayer, 0, len(playersMap))
	for _, p := range playersMap {
		players = append(players, p)
	}

	sort.SliceStable(players, func(i, j int) bool {
		return players[i].Rank < players[j].Rank
	})

	return players, err
}

func (s *ArenaStore) GetOpponentData(ctx context.Context, accountID primitive.ObjectID) (*arena.ArenaPlayerData, error) {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "avatar_id", Value: 1},
		{Key: "arena_data.rating", Value: 1},
		{Key: "arena_data.league", Value: 1},
		{Key: "decks", Value: 1},
	}
	var playerData player.PlayerData
	err := s.collection.FindOne(ctx, bson.M{"_id": accountID}, options.FindOne().SetProjection(projection)).Decode(&playerData)
	if err != nil {
		return nil, err
	}

	var defenseDeck *cards.Deck
	if len(playerData.Decks.Decks) > 0 {
		defenseDeck = playerData.Decks.Decks[playerData.Decks.Defense]
	}

	result := &arena.ArenaPlayerData{
		ID:       playerData.ID,
		Name:     playerData.Name,
		AvatarID: playerData.AvatarID,
		Rating:   playerData.ArenaData.Rating,
		League:   playerData.ArenaData.League,
	}

	if defenseDeck != nil {
		result.DefenseDeck = defenseDeck.Cards
	}

	return result, err
}

func (s *ArenaStore) SaveBattleLog(ctx context.Context, data *arena.ArenaBattleLog) error {
	_, err := s.battleLogsCollection.InsertOne(ctx, data)
	return err
}

func (s *ArenaStore) GetBattleLogs(ctx context.Context, timestamp *base.Timestamp, accountID primitive.ObjectID) ([]*arena.ArenaBattleLog, error) {
	if timestamp == nil {
		timestamp = &base.Timestamp{}
	}
	cursor, err := s.battleLogsCollection.Find(ctx, bson.M{
		"$and": bson.A{
			bson.M{"opponent._id": accountID},
			bson.M{"created_at": bson.M{"$gt": timestamp}},
		},
	})

	var result []*arena.ArenaBattleLog
	if err == nil {
		if err = cursor.All(ctx, &result); err != nil {
			log.Error(err)
			return nil, err
		}
		_, _ = s.collection.UpdateByID(ctx, accountID, bson.M{"$set": bson.M{"arena_data.last_seen_battle_log": stime.RealTime()}})
	}

	return result, err
}

func (s *ArenaStore) DeleteBattleLogs(ctx context.Context, accountID primitive.ObjectID) error {
	_, err := s.battleLogsCollection.DeleteMany(ctx, bson.M{"$or": bson.A{bson.M{"player._id": accountID}, bson.M{"opponent._id": accountID}}})
	return err
}

func (s *ArenaStore) GetRating(ctx context.Context, accountID primitive.ObjectID) (int32, int32, error) {
	projection := bson.D{
		{Key: "arena_data.rating", Value: 1},
		{Key: "arena_data.max_rating", Value: 1},
	}
	opts := options.FindOne().SetProjection(projection)
	data := struct {
		ArenaData struct {
			Rating    int32 `bson:"rating"`
			MaxRating int32 `bson:"max_rating"`
		} `bson:"arena_data"`
	}{}
	err := s.collection.FindOne(ctx, bson.M{"_id": accountID}, opts).Decode(&data)
	return data.ArenaData.Rating, data.ArenaData.MaxRating, err
}

func (s *ArenaStore) SetRating(ctx context.Context, accountID primitive.ObjectID, rating int32, maxRating int32) error {
	_, err := s.collection.UpdateByID(ctx, accountID, bson.M{"$set": bson.M{"arena_data.rating": rating, "arena_data.max_rating": maxRating}})
	return err
}
