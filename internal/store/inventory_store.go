package store

import (
	"context"
	"locgame-mini-server/internal/utils/metrics"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/resources"
)

// InventoryStore stores data for the storage of data about inventory.
type InventoryStore struct {
	db     *mongo.Client
	config *config.Config

	collection *mongo.Collection
}

type InventoryStoreInterface interface {
	IncrementResources(ctx context.Context, accountID primitive.ObjectID, adjustments []*resources.ResourceAdjustment, reason string) error
}

// NewInventoryStore creates a new instance of the inventory store.
func NewInventoryStore(config *config.Config, db *mongo.Client) *InventoryStore {
	s := new(InventoryStore)
	s.config = config
	s.db = db
	s.collection = db.Database(s.config.Database.Database).Collection("players")

	return s
}

// IncrementResources allows to apply a specific reward to a specific player.
func (s *InventoryStore) IncrementResources(ctx context.Context, accountID primitive.ObjectID, adjustments []*resources.ResourceAdjustment, reason string) error {
	defer logResource(adjustments, accountID, reason)

	var err error
	_, err = s.collection.UpdateByID(ctx, accountID, bson.M{"$inc": s.getBsonMap(adjustments)})
	if err != nil {
		return err
	}
	return err
}

func logResource(adjustments []*resources.ResourceAdjustment, accountID primitive.ObjectID, reason string) {
	for _, adjustment := range adjustments {
		if adjustment.ResourceID == 1 {
			if adjustment.Quantity > 0 {
				metrics.GetDefault().LCEarned(accountID, adjustment.Quantity, reason)
			} else {
				metrics.GetDefault().LCSpent(accountID, adjustment.Quantity, reason)
			}
		}
	}
}

func (s *InventoryStore) getBsonMap(adjustment []*resources.ResourceAdjustment) bson.M {
	result := bson.M{}

	const prefix = "resources."

	for _, adjustment := range adjustment {
		key := prefix + strconv.Itoa(int(adjustment.ResourceID))
		if result[key] == nil {
			result[key] = adjustment.Quantity
		} else {
			result[key] = adjustment.Quantity + result[key].(int32)
		}
	}
	return result
}
