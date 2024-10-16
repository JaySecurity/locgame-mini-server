package store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// InGameStore stores data for the storage of data about in-game store.
type InGameStore struct {
	db     *mongo.Client
	config *config.Config

	collection           *mongo.Collection
	salesCollection      *mongo.Collection
	promoCodesCollection *mongo.Collection
}

type PromoCodeData struct {
	IsOwner         bool
	PromoCodeTypeId string `bson:"product_id"`
	PromoCode       string `bson:"promo_code"`
}

type InGameStoreInterface interface {
	IncrementPurchases(ctx context.Context, accountID *base.ObjectID, productID string, qty int64) error
	SetLastWithdrawalAt(ctx context.Context, accountID *base.ObjectID, lastWithdrawal *base.Timestamp) error
	GetProductsSold(ctx context.Context) (map[string]int64, error)
	GetProductSold(ctx context.Context, productID string) (int64, error)
	IncrementSold(ctx context.Context, productID string, qty int64) error
	GetPromoCodeData(ctx context.Context, promoCode string, playerID primitive.ObjectID) (*PromoCodeData, error)
}

// NewInGameStore creates a new instance of the in-game store.
func NewInGameStore(config *config.Config, db *mongo.Client) *InGameStore {
	s := new(InGameStore)
	s.config = config
	s.db = db
	s.collection = db.Database(s.config.Database.Database).Collection("players")
	s.salesCollection = db.Database(s.config.Database.Database).Collection("sales")
	s.promoCodesCollection = db.Database(s.config.Database.Database).Collection("promo_codes")

	return s
}

// IncrementPurchases allows to increment purchases a specific product to a specific player.
func (s *InGameStore) IncrementPurchases(ctx context.Context, accountID *base.ObjectID, productID string, qty int64) error {
	var err error
	_, err = s.collection.UpdateByID(ctx, accountID, bson.M{"$inc": bson.M{"player_store_data.purchased_products." + productID: qty}})
	if err != nil {
		log.Error(err)
		return err
	}
	return err
}

func (s *InGameStore) SetLastWithdrawalAt(ctx context.Context, accountID *base.ObjectID, lastWithdrawal *base.Timestamp) error {
	_, err := s.collection.UpdateByID(ctx, accountID, bson.M{"$set": bson.M{"player_store_data.last_withdrawal_at": lastWithdrawal}})
	return err
}

func (s *InGameStore) IncrementSold(ctx context.Context, productID string, qty int64) error {
	_, err := s.salesCollection.UpdateOne(ctx, bson.M{"product_id": productID}, bson.M{"$inc": bson.M{"sold": qty}}, options.Update().SetUpsert(true))
	return err
}

func (s *InGameStore) GetProductsSold(ctx context.Context) (map[string]int64, error) {
	cursor, err := s.salesCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var data []struct {
		ProductID string `bson:"product_id"`
		Sold      int64  `bson:"sold"`
	}
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	sales := make(map[string]int64)

	for _, sale := range data {
		sales[sale.ProductID] = sale.Sold
	}
	return sales, nil
}

func (s *InGameStore) GetProductSold(ctx context.Context, productID string) (int64, error) {
	var data struct {
		ProductID string `bson:"product_id"`
		Sold      int64  `bson:"sold"`
	}
	err := s.salesCollection.FindOne(ctx, bson.M{"product_id": productID}).Decode(&data)
	if err == mongo.ErrNoDocuments {
		return 0, nil
	}

	return data.Sold, nil
}

func (s *InGameStore) GetPromoCodeData(ctx context.Context, promoCode string, playerId primitive.ObjectID) (*PromoCodeData, error) {
	var data struct {
		PlayerId        primitive.ObjectID `bson:"player_id"`
		PromoCodeTypeId string             `bson:"promo_code_type_id"`
		PromoCode       string             `bson:"promo_code"`
	}

	err := s.promoCodesCollection.FindOne(ctx, bson.M{"promo_code": promoCode}).Decode(&data)

	if err != nil {
		return nil, err
	}

	promoCodeData := PromoCodeData{
		IsOwner:         playerId == data.PlayerId,
		PromoCodeTypeId: data.PromoCodeTypeId,
		PromoCode:       data.PromoCode,
	}

	return &promoCodeData, nil
}
