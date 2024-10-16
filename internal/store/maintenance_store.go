// Copyright © 2022 FuryLion Group LLC. All Rights Reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// without the express permission of FuryLion Group LLC.
// Proprietary and confidential.
//
// Created by Vadim Vlasov on 11.10.22

package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/maintenance"
)

type MaintenanceStore struct {
	db     *mongo.Client
	config *config.Config

	collection *mongo.Collection
}

func NewMaintenanceStore(config *config.Config, db *mongo.Client) *MaintenanceStore {
	s := new(MaintenanceStore)
	s.db = db
	s.config = config

	s.collection = db.Database(s.config.Database.Database).Collection("maintenance")
	return s
}

func (s *MaintenanceStore) Get(ctx context.Context) ([]*maintenance.MaintenanceData, error) {
	cursor, err := s.collection.Find(ctx, bson.D{})
	var data []*maintenance.MaintenanceData
	if err == nil {
		err = cursor.All(ctx, &data)
	}
	return data, err
}
