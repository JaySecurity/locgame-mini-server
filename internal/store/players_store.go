package store

import (
	"context"
	"locgame-mini-server/pkg/dto/arena"
	"regexp"
	"strings"

	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/utils/bsonutils"
	"locgame-mini-server/pkg/dto/accounts"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/errors"
	"locgame-mini-server/pkg/dto/game"
	"locgame-mini-server/pkg/dto/player"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/stime"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PlayersStore stores data for the operation of the account store.
type PlayersStore struct {
	db     *mongo.Client
	config *config.Config

	collection *mongo.Collection
}

type WatchStreamEvent struct {
	ID                bson.M             `bson:"_id" json:"_id"`
	OperationType     string             `bson:"operationType" json:"operationType"`
	FullDocument      bson.M             `bson:"fullDocument,omitempty" json:"fullDocument,omitempty"`
	Namespace         Namespace          `bson:"ns" json:"ns"`
	DocumentKey       bson.M             `bson:"documentKey" json:"documentKey"`
	UpdateDescription *UpdateDescription `bson:"updateDescription,omitempty" json:"updateDescription,omitempty"`
}

// Namespace represents the namespace of the document that was changed
type Namespace struct {
	DB   string `bson:"db" json:"db"`
	Coll string `bson:"coll" json:"coll"`
}

// UpdateDescription represents the description of updates made to the document
type UpdateDescription struct {
	UpdatedFields bson.M   `bson:"updatedFields" json:"updatedFields"`
	RemovedFields []string `bson:"removedFields" json:"removedFields"`
}

// PlayersStoreInterface provides an interface for the ability to later mock
//
//go:generate mockery -name PlayersStoreInterface
type PlayersStoreInterface interface {
	SetOnlineState(ctx context.Context, accountID primitive.ObjectID, isOnline bool) error

	GetAccountIDByWallet(context.Context, string) (primitive.ObjectID, error)

	RegisterAccount(context.Context, *player.PlayerData) (primitive.ObjectID, error)

	GetPlayerDataByID(context.Context, primitive.ObjectID) (*player.PlayerData, error)

	SetDataWithProj(context.Context, primitive.ObjectID, *player.PlayerData, *player.PlayerData) error
	SetData(context.Context, primitive.ObjectID, *player.PlayerData) error
	ForceSetData(ctx context.Context, accountID primitive.ObjectID, data *player.PlayerData) error

	SetArenaData(ctx context.Context, accountID primitive.ObjectID, data *arena.ArenaData) error

	GetPlayerInfoWithDeckData(ctx context.Context, accountID *base.ObjectID) (*game.PlayerInfo, []string, error)

	GetUsersInfo(ctx context.Context, ids ...*base.ObjectID) ([]*accounts.UserInfo, error)
	GetUsersInfoAsMap(ctx context.Context, ids ...*base.ObjectID) (map[string]*accounts.UserInfo, error)

	Find(ctx context.Context, accountID *base.ObjectID, query string) ([]*accounts.UserInfo, error)

	IsOnline(ctx context.Context, accountID *base.ObjectID) bool

	GetWalletByPlayerID(ctx context.Context, playerID *base.ObjectID) (string, error)

	GetAccountIDByCognitoUsername(context.Context, string) (primitive.ObjectID, error)
	GetAccountIDByEmail(context.Context, string) (primitive.ObjectID, error)
	SetActiveWallet(ctx context.Context, id primitive.ObjectID, in *accounts.SetActiveWalletRequest) error
}

// NewPlayersStore creates a new instance of the account store.
func NewPlayersStore(config *config.Config, db *mongo.Client) *PlayersStore {
	s := new(PlayersStore)
	s.db = db
	s.config = config

	s.collection = db.Database(s.config.Database.Database).Collection("players")

	if _, err := s.collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "external_wallet", Value: 1},
		},
		Options: options.Index().SetUnique(true).SetSparse(true),
	}); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			log.Info(err, "Indexes already exists")
		} else {
			log.Fatal(err)
		}
	}

	if _, err := s.collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "last_activity", Value: 1},
		},
	}); err != nil {
		log.Fatal(err)
	}

	// Uncomment this line to enable change stream logging
	// go watchChangeStream(s.collection)

	return s
}

func watchChangeStream(collection *mongo.Collection) {
	// Set up a change stream to watch for updates
	pipeline := mongo.Pipeline{bson.D{{"$match", bson.D{{"operationType", "update"}}}}}
	opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)
	changeStream, err := collection.Watch(context.TODO(), pipeline, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer changeStream.Close(context.Background())

	// Iterate over the change stream
	for changeStream.Next(context.Background()) {
		var event WatchStreamEvent
		if err := changeStream.Decode(&event); err != nil {
			log.Fatal(err)
		}
		log.Debugf("Change detected Id: %v\n", event.ID)
		log.Debugf("Change detected Updates: %v\n", event.UpdateDescription.UpdatedFields)
		log.Debugf("Change detected Removed: %v\n", event.UpdateDescription.RemovedFields)
	}

	if err := changeStream.Err(); err != nil {
		log.Fatal(err)
	}

	log.Debug("Finished watching the change stream.")
}

func (s *PlayersStore) GetAccountIDByWallet(ctx context.Context, address string) (primitive.ObjectID, error) {
	identifier := struct {
		ID primitive.ObjectID `bson:"_id"`
	}{}

	err := s.collection.FindOne(ctx, bson.M{"$or": []bson.M{{"external_wallet": address}, {"wallet": address}}}, options.FindOne().SetProjection(bson.M{"_id": 1})).Decode(&identifier)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return identifier.ID, nil
}

func (s *PlayersStore) GetAccountIDByEmail(ctx context.Context, email string) (primitive.ObjectID, error) {
	identifier := struct {
		ID primitive.ObjectID `bson:"_id"`
	}{}

	err := s.collection.FindOne(ctx, bson.M{"email": email}, options.FindOne().SetProjection(bson.M{"_id": 1})).Decode(&identifier)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return identifier.ID, nil
}

// GetPlayerDataByID returns specific player data from the database by account id.
func (s *PlayersStore) GetPlayerDataByID(ctx context.Context, accountID primitive.ObjectID) (*player.PlayerData, error) {
	var data player.PlayerData
	err := s.collection.FindOne(ctx, bson.M{"_id": accountID}).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Debug("User not found: " + accountID.Hex()) // Convert accountID to string using Hex() method
			err = errors.ErrUserNotFound
			return nil, err
		}
		log.Errorf("Error while getting player data: %v", err)
		return nil, err
	}
	return &data, err
}

func (s *PlayersStore) GetWalletByPlayerID(ctx context.Context, playerID *base.ObjectID) (string, error) {
	var data struct {
		Wallet string `bson:"active_wallet"`
	}
	err := s.collection.FindOne(ctx, bson.M{"_id": playerID}).Decode(&data)
	return data.Wallet, err
}

// RegisterAccount registers a new account in the database.
func (s *PlayersStore) RegisterAccount(ctx context.Context, initialData *player.PlayerData) (primitive.ObjectID, error) {
	res, err := s.collection.InsertOne(ctx, initialData)
	if err == nil {
		accountID := res.InsertedID.(primitive.ObjectID)
		return accountID, nil
	}

	log.Error("Failed to create new player:", err)
	return primitive.ObjectID{}, err
}

func (s *PlayersStore) SetOnlineState(ctx context.Context, accountID primitive.ObjectID, isOnline bool) error {
	_, err := s.collection.UpdateByID(ctx, accountID, bson.M{"$set": bson.M{"online": isOnline, "last_activity": stime.RealTime()}})
	return err
}

func (s *PlayersStore) SetDataWithProj(ctx context.Context, accountID primitive.ObjectID, data *player.PlayerData, proj *player.PlayerData) error {
	opts := new(bsonutils.BSONOptions)
	opts.SetProjection(proj)
	_, err := s.collection.UpdateByID(ctx, accountID, bson.M{"$set": bsonutils.ToBSONMap(data, opts)})
	return err
}

func (s *PlayersStore) SetData(ctx context.Context, accountID primitive.ObjectID, data *player.PlayerData) error {
	_, err := s.collection.UpdateByID(ctx, accountID, bson.M{"$set": bsonutils.ToBSONMap(data)})
	return err
}

func (s *PlayersStore) ForceSetData(ctx context.Context, accountID primitive.ObjectID, data *player.PlayerData) error {

	_, err := s.collection.UpdateByID(ctx, accountID, bson.M{"$set": bsonutils.ToBSONMap(data)})
	return err
}

func (s *PlayersStore) SetArenaData(ctx context.Context, accountID primitive.ObjectID, data *arena.ArenaData) error {
	_, err := s.collection.UpdateByID(ctx, accountID, bson.M{"$set": &player.PlayerData{ArenaData: data}})
	return err
}

func (s *PlayersStore) GetPlayerInfoWithDeckData(ctx context.Context, accountID *base.ObjectID) (*game.PlayerInfo, []string, error) {
	var data player.PlayerData
	projection := bson.D{
		{"_id", 1},
		{"name", 1},
		{"avatar_id", 1},
		{"decks", 1},
	}
	opts := options.FindOne().SetProjection(projection)
	err := s.collection.FindOne(ctx, bson.M{"_id": accountID}, opts).Decode(&data)
	if err == nil {
		return &game.PlayerInfo{
			ID:         data.ID,
			PlayerType: game.PlayerType_Real,
			Name:       data.Name,
			AvatarID:   data.AvatarID,
		}, data.Decks.Decks[data.Decks.Active].Cards, nil
	}
	return nil, nil, err
}

func (s *PlayersStore) GetUsersInfo(ctx context.Context, ids ...*base.ObjectID) ([]*accounts.UserInfo, error) {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "avatar_id", Value: 1},
		{Key: "arena_data.rating", Value: 1},
		{Key: "arena_data.league", Value: 1},
		{Key: "online", Value: 1},
	}

	var result []*accounts.UserInfo
	cursor, err := s.collection.Find(ctx, bson.M{"_id": bson.M{"$in": ids}}, options.Find().SetProjection(projection))
	if err == nil {
		for cursor.Next(ctx) {
			var data *player.PlayerData
			err = cursor.Decode(&data)
			if err != nil {
				log.Error(err)
				continue
			}

			result = append(result, s.convertPlayerDataToUserInfo(data))
		}
		_ = cursor.Close(ctx)
		return result, err
	}
	return nil, err
}

func (s *PlayersStore) GetUsersInfoAsMap(ctx context.Context, ids ...*base.ObjectID) (map[string]*accounts.UserInfo, error) {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "avatar_id", Value: 1},
		{Key: "arena_data.rating", Value: 1},
		{Key: "arena_data.league", Value: 1},
		{Key: "online", Value: 1},
	}

	result := make(map[string]*accounts.UserInfo)

	cursor, err := s.collection.Find(ctx, bson.M{"_id": bson.M{"$in": ids}}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var data *player.PlayerData
		err = cursor.Decode(&data)
		if err != nil {
			log.Error(err)
			continue
		}

		result[data.ID.Value] = s.convertPlayerDataToUserInfo(data)
	}
	_ = cursor.Close(ctx)

	return result, err
}

func (s *PlayersStore) convertPlayerDataToUserInfo(data *player.PlayerData) *accounts.UserInfo {
	result := &accounts.UserInfo{
		ID:       data.ID,
		Name:     data.Name,
		AvatarID: data.AvatarID,
		IsOnline: data.Online,
	}
	if data.ArenaData != nil {
		result.Rating = data.ArenaData.Rating
		result.League = int32(data.ArenaData.League)
	}
	return result
}

func (s *PlayersStore) Find(ctx context.Context, accountID *base.ObjectID, query string) ([]*accounts.UserInfo, error) {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "avatar_id", Value: 1},
		{Key: "arena_data.rating", Value: 1},
		{Key: "arena_data.league", Value: 1},
		{Key: "online", Value: 1},
	}

	if accountID.Value == query {
		return nil, nil
	}

	id, err := primitive.ObjectIDFromHex(query)
	var users []*accounts.UserInfo
	if err == nil {
		var playerData *player.PlayerData
		err = s.collection.FindOne(ctx, bson.M{"_id": id}, options.FindOne().SetProjection(projection)).Decode(&playerData)
		if err == nil {
			users = append(users, s.convertPlayerDataToUserInfo(playerData))
		}
	} else {
		cursor, err := s.collection.Find(ctx, bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: "^" + regexp.QuoteMeta(query) + ".*", Options: "i"}}}, options.Find().SetProjection(projection).SetLimit(10))
		if err != nil {
			return nil, err
		}
		for cursor.Next(ctx) {
			var playerData *player.PlayerData
			err = cursor.Decode(&playerData)
			if err == nil {
				if playerData.ID.Value == accountID.Value {
					continue
				}
				users = append(users, s.convertPlayerDataToUserInfo(playerData))
			}
		}
		_ = cursor.Close(ctx)
	}
	return users, nil
}

func (s *PlayersStore) IsOnline(ctx context.Context, accountID *base.ObjectID) bool {
	projection := bson.D{
		{Key: "online", Value: 1},
	}
	data := struct {
		Online bool `bson:"online"`
	}{}
	err := s.collection.FindOne(ctx, bson.M{"_id": accountID}, options.FindOne().SetProjection(projection)).Decode(&data)
	if err != nil {
		return false
	}

	return data.Online
}

func (s *PlayersStore) GetAccountIDByCognitoUsername(ctx context.Context, CognitoUsername string) (primitive.ObjectID, error) {
	identifier := struct {
		ID primitive.ObjectID `bson:"_id"`
	}{}

	err := s.collection.FindOne(ctx, bson.M{"cognito_username": CognitoUsername}, options.FindOne().SetProjection(bson.M{"_id": 1})).Decode(&identifier)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return identifier.ID, nil
}

func (s *PlayersStore) SetActiveWallet(ctx context.Context, accountID primitive.ObjectID, in *accounts.SetActiveWalletRequest) error {
	wallet := strings.ToLower(in.Wallet)
	provider := in.Provider

	switch provider {
	case accounts.ProviderType_Particle:
		_, err := s.collection.UpdateByID(ctx, accountID, bson.M{"$set": bson.M{"particle_wallet": wallet, "active_wallet": wallet}})
		if mongo.IsDuplicateKeyError(err) {
			err = errors.ErrDuplicateWallet
		}
		return err
	case accounts.ProviderType_MetaMask:
		_, err := s.collection.UpdateByID(ctx, accountID, bson.M{"$set": bson.M{"external_wallet": wallet, "active_wallet": wallet}})
		if mongo.IsDuplicateKeyError(err) {
			err = errors.ErrDuplicateWallet
		}
		return err
	case accounts.ProviderType_None:
		_, err := s.collection.UpdateByID(ctx, accountID, bson.M{"$unset": bson.M{"external_wallet": wallet, "active_wallet": wallet}})
		if mongo.IsDuplicateKeyError(err) {
			err = errors.ErrDuplicateWallet
		}
		return err
	}

	return nil
}
