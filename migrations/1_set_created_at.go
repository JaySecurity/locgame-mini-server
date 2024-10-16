// Copyright © 2022 FuryLion Group LLC. All Rights Reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// without the express permission of FuryLion Group LLC.
// Proprietary and confidential.
//
// Created by Vadim Vlasov on 05.12.22

package migrations

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/migrate"
)

func init() {
	const collectionName = "players"

	err := migrate.Register(func(db *mongo.Database) error {
		upFunc := func(db *mongo.Database) error {
			_, err := db.Collection(collectionName).UpdateMany(context.Background(), bson.M{"created_at": bson.M{"$exists": false}}, bson.M{"$set": bson.M{"created_at": time.Now().UTC()}})
			if err != nil {
				log.Error(err)
			}
			return err
		}
		return upFunc(db)
	}, func(db *mongo.Database) error {
		downFunc := func(db *mongo.Database) error {
			_, err := db.Collection(collectionName).UpdateMany(context.Background(), bson.M{}, bson.M{"$unset": bson.M{"created_at": nil}})
			if err != nil {
				log.Error(err)
			}
			return err
		}
		return downFunc(db)
	})
	if err != nil {
		log.Fatal("An error occurred during migration:", err)
	}
}
