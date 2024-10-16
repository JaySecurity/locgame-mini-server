package store

import (
	"context"
	"fmt"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/utils/bsonutils"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/migrate"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	// Automatic import all migrations.
	_ "locgame-mini-server/migrations"
)

// Store stores references to all other stores required for the current service to run.
type Store struct {
	config *config.Config

	db    *mongo.Client
	redis *redis.Client

	Players          PlayersStoreInterface
	Matchmaking      MatchmakingStoreInterface
	QuickMatch       QuickMatchStoreInterface
	Inventory        InventoryStoreInterface
	Friends          FriendsRepository
	Auth             AuthStoreInterface
	ArenaLeaderboard LeaderboardStore
	Arena            ArenaStoreInterface
	Jobs             *Jobs
	DistributedLocks *DistributedLocks
	Orders           OrdersStoreInterface
	Gifts            GiftsStoreInterface
	Mint             MintStoreInterface
	InGameStore      InGameStoreInterface
	Maintenance      *MaintenanceStore
	DailyRewards     DailyRewardsStoreInterface
	Discounts        *DiscountsStore
}

// NewStore creates a new storage instance.
func NewStore(config *config.Config) *Store {
	s := new(Store)
	s.config = config

	s.db = s.ConnectDatabase()
	s.redis = s.ConnectRedis()

	migrate.SetDatabase(s.db.Database(s.config.Database.Database))
	if err := migrate.Up(-1); err != nil {
		log.Fatal("Error:", err)
	}

	s.Players = NewPlayersStore(s.config, s.db)
	s.Matchmaking = NewMatchmakingStore(s.config, s.redis)
	s.Inventory = NewInventoryStore(s.config, s.db)
	s.Friends = NewMongoFriendsRepository(s.config, s.db)
	s.Auth = NewAuthStore(s.config, s.db, s.redis)
	s.Arena = NewArenaStore(s.config, s.db)
	s.ArenaLeaderboard = NewArenaLeaderboardStore(s.config, s.redis)
	s.Jobs = NewJobs(s.config, s.db)
	s.DistributedLocks = NewDistributedLocks(s.redis)
	s.Orders = NewOrdersStore(s.config, s.db)
	s.Mint = NewMintStore(s.config, s.db)
	s.InGameStore = NewInGameStore(s.config, s.db)
	s.Maintenance = NewMaintenanceStore(s.config, s.db)
	s.Gifts = NewGiftsStore(s.config, s.db)
	s.DailyRewards = NewDailyRewardsStore(s.config, s.db)
	s.Discounts = NewDiscountsStore(s.config, s.db)

	return s
}

func (s *Store) ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", s.config.Redis.Host, s.config.Redis.Port),
		Username: s.config.Redis.Username,
		Password: s.config.Redis.Password,
		DB:       s.config.Redis.Database,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("Unable to connect to Redis:", err)
	}

	log.Debugf("Successful connection to Redis: %s:%d (Database: %d)", s.config.Redis.Host, s.config.Redis.Port, s.config.Redis.Database)
	return client
}

// ConnectDatabase allows connecting to the database.
func (s *Store) ConnectDatabase() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reg := bsonutils.Register(bson.NewRegistryBuilder()).Build()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.config.Database.GetConnectionString()).SetRegistry(reg))
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Cannot connect to database ("+s.config.Database.GetConnectionStringForDisplay()+"):", err)
	}

	log.Debug("Successful connection to the database:", s.config.Database.GetConnectionStringForDisplay(), "("+s.config.Database.Database+")")

	return client
}

// Disconnect allows disconnecting from the database.
func (s *Store) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := s.db.Disconnect(ctx); err != nil {
		panic(err)
	}
}
