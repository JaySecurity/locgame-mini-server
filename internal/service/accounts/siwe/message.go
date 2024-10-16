package siwe

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Message struct {
	Address common.Address `bson:"address,omitempty"`
	Nonce   string         `bson:"nonce,omitempty"`

	IssuedAt       time.Time `bson:"issued_at,omitempty"`
	ExpirationTime time.Time `bson:"expiration_time,omitempty"`

	Client string `bson:"client"`
}
