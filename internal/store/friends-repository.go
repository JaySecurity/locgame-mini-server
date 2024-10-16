package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/log"
)

// MongoFriendsRepository stores data for the storage of data about friends and friends requests.
type MongoFriendsRepository struct {
	db     *mongo.Client
	config *config.Config

	collection *mongo.Collection
}

type FriendsRepository interface {
	GetRequests(ctx context.Context, accountID primitive.ObjectID) ([]*FriendRequestsInfo, error)
	CreateRequest(ctx context.Context, senderID primitive.ObjectID, receiverID *base.ObjectID) error
	AcceptRequest(ctx context.Context, accountID primitive.ObjectID, senderID *base.ObjectID) error
	DeleteRequest(ctx context.Context, accountID primitive.ObjectID, senderID *base.ObjectID) error
	GetFriendsCount(ctx context.Context, accountID *base.ObjectID) (int32, error)
}

type FriendRequestStatus int32

const (
	FriendRequestSent = iota
	FriendRequestAccepted
)

type FriendRequestsInfo struct {
	SenderID   *base.ObjectID `bson:"sender_id"`
	ReceiverID *base.ObjectID `bson:"receiver_id"`
	Status     int32          `bson:"status"`
}

// NewMongoFriendsRepository creates a new instance of the friends' repository.
func NewMongoFriendsRepository(config *config.Config, db *mongo.Client) *MongoFriendsRepository {
	repo := new(MongoFriendsRepository)
	repo.config = config
	repo.db = db
	repo.collection = db.Database(repo.config.Database.Database).Collection("friends")

	if _, err := repo.collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "receiver_id", Value: 1},
			},
		}, {
			Keys: bson.D{
				{Key: "sender_id", Value: 1},
			},
		},
		{
			Keys: bson.D{
				{Key: "status", Value: 1},
			},
		},
	}); err != nil {
		log.Fatal(err)
	}

	return repo
}

func (repo *MongoFriendsRepository) GetRequests(ctx context.Context, accountID primitive.ObjectID) ([]*FriendRequestsInfo, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{
		"$and": bson.A{
			bson.M{"$or": bson.A{bson.M{"receiver_id": accountID}, bson.M{"sender_id": accountID}}},
		},
	})
	var requests []*FriendRequestsInfo
	err = cursor.All(ctx, &requests)
	if err != nil {
		return nil, err
	}
	_ = cursor.Close(ctx)
	return requests, nil
}

func (repo *MongoFriendsRepository) CreateRequest(ctx context.Context, senderID primitive.ObjectID, receiverID *base.ObjectID) error {
	_, err := repo.collection.InsertOne(ctx, &FriendRequestsInfo{
		SenderID:   &base.ObjectID{Value: senderID.Hex()},
		ReceiverID: receiverID,
		Status:     FriendRequestSent,
	})
	return err
}

func (repo *MongoFriendsRepository) AcceptRequest(ctx context.Context, accountID primitive.ObjectID, senderID *base.ObjectID) error {
	_, err := repo.collection.UpdateOne(ctx, bson.M{"sender_id": senderID, "receiver_id": accountID}, bson.M{"$set": bson.M{"status": FriendRequestAccepted}})
	return err
}

func (repo *MongoFriendsRepository) DeleteRequest(ctx context.Context, accountID primitive.ObjectID, senderID *base.ObjectID) error {
	_, err := repo.collection.DeleteOne(ctx, bson.M{"$or": bson.A{
		bson.M{"$and": bson.A{bson.M{"sender_id": senderID}, bson.M{"receiver_id": accountID}}},
		bson.M{"$and": bson.A{bson.M{"sender_id": accountID}, bson.M{"receiver_id": senderID}}},
	}})
	return err
}

func (repo *MongoFriendsRepository) GetFriendsCount(ctx context.Context, accountID *base.ObjectID) (int32, error) {
	count, err := repo.collection.CountDocuments(ctx, bson.M{
		"$or": bson.A{
			bson.M{"receiver_id": accountID, "status": FriendRequestAccepted},
			bson.M{"sender_id": accountID, "status": FriendRequestAccepted},
		},
	})

	return int32(count), err
}
