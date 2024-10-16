package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/store"
	"locgame-mini-server/pkg/log"
)

// OrdersStore stores data for the storage of data about orders.
type OrdersStore struct {
	db     *mongo.Client
	config *config.Config

	collection *mongo.Collection
}

type OrdersStoreInterface interface {
	Create(ctx context.Context, order *store.Order) (string, error)
	Update(ctx context.Context, order *store.Order) error
	Get(ctx context.Context, orderID string) (*store.Order, error)
	GetByHash(ctx context.Context, txHash string) (*store.Order, error)
	GetUnopenedPacks(ctx context.Context, buyerID *base.ObjectID) ([]*store.Order, error)
	DeleteUnpaidOrders(ctx context.Context) (int64, error)
	ClearErrorMessage(ctx context.Context, orderID *base.ObjectID)
}

// NewOrdersStore creates a new instance of the orders store.
func NewOrdersStore(config *config.Config, db *mongo.Client) *OrdersStore {
	s := new(OrdersStore)
	s.config = config
	s.db = db
	s.collection = db.Database(s.config.Database.Database).Collection("orders")

	if _, err := s.collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "payment_hash", Value: -1},
			},
		},
		{
			Keys: bson.D{
				{Key: "created_at", Value: -1},
			},
		},
		{
			Keys: bson.D{
				{Key: "buyer_id", Value: -1},
			},
		},
		{
			Keys: bson.D{
				{Key: "status", Value: -1},
			},
		},
	}); err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *OrdersStore) Create(ctx context.Context, order *store.Order) (string, error) {
	res, err := s.collection.InsertOne(ctx, order)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (s *OrdersStore) Update(ctx context.Context, order *store.Order) error {
	_, err := s.collection.UpdateByID(ctx, order.ID, bson.M{"$set": order})
	return err
}

func (s *OrdersStore) ClearErrorMessage(ctx context.Context, orderID *base.ObjectID) {
	_, _ = s.collection.UpdateByID(ctx, orderID, bson.M{"$unset": bson.M{"error": ""}})
}

func (s *OrdersStore) Get(ctx context.Context, orderID string) (*store.Order, error) {
	var order *store.Order
	id, _ := primitive.ObjectIDFromHex(orderID)
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	return order, err
}

func (s *OrdersStore) GetByHash(ctx context.Context, txHash string) (*store.Order, error) {
	var order *store.Order
	err := s.collection.FindOne(ctx, bson.M{"payment_hash": txHash}).Decode(&order)
	return order, err
}

func (s *OrdersStore) GetUnopenedPacks(ctx context.Context, buyerID *base.ObjectID) ([]*store.Order, error) {
	var orders []*store.Order
	statuses := []store.OrderStatus{
		store.OrderStatus_WaitingForPayment,
		store.OrderStatus_PaymentReceived,
		store.OrderStatus_InProgress,
		store.OrderStatus_Completed,
	}

	cursor, err := s.collection.Find(ctx, bson.M{
		"$and": bson.A{
			bson.M{"buyer_id": buyerID},
			bson.M{"status": bson.M{"$in": statuses}},
		}})

	if err != nil {
		log.Error(err)
		return nil, err
	}

	for cursor.Next(ctx) {
		var order store.Order
		err = cursor.Decode(&order)
		if err == nil {
			if product, ok := s.config.Products.ProductsByID[order.ProductID]; ok && product.Type != store.ProductType_PackOfCoins {
				orders = append(orders, &order)
			}
		}
	}
	_ = cursor.Close(ctx)
	return orders, nil
}

func (s *OrdersStore) DeleteUnpaidOrders(ctx context.Context) (int64, error) {
	statuses := []store.OrderStatus{store.OrderStatus_Unknown, store.OrderStatus_WaitingForPayment, store.OrderStatus_Canceled}
	result, err := s.collection.DeleteMany(ctx, bson.M{"$and": bson.A{bson.M{"status": bson.M{"$in": statuses}}, bson.M{"created_at": bson.M{"$lte": time.Now().UTC().Add(-24 * time.Hour)}}}})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, err
}
