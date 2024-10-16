// Copyright © 2022 FuryLion Group LLC. All Rights Reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// without the express permission of FuryLion Group LLC.
// Proprietary and confidential.
//
// Created by Vadim Vlasov on 25.11.22

package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/store"
)

type DiscountsStore struct {
	db     *mongo.Client
	config *config.Config

	collection *mongo.Collection
}

func NewDiscountsStore(config *config.Config, db *mongo.Client) *DiscountsStore {
	s := new(DiscountsStore)
	s.db = db
	s.config = config

	s.collection = db.Database(s.config.Database.Database).Collection("discounts")
	return s
}

func (s *DiscountsStore) Get(ctx context.Context) ([]*store.Discount, error) {
	cursor, err := s.collection.Find(ctx, bson.D{})
	var data []*store.Discount
	if err == nil {
		err = cursor.All(ctx, &data)
	}
	return data, err
}
