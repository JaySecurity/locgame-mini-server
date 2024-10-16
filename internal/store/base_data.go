package store

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type baseData struct {
	accountID primitive.ObjectID
}

func (d *baseData) init(accountID primitive.ObjectID) {
	d.accountID = accountID
}
