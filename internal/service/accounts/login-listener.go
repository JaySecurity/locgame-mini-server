package accounts

import (
	"context"
	"locgame-mini-server/internal/utils/metrics"
	"locgame-mini-server/pkg/dto/player"
	"locgame-mini-server/pkg/log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LoginListener defines the interface for listening to user login events
type LoginListener interface {
	OnLogin(ctx context.Context, accountID primitive.ObjectID, user *player.PlayerData)
}

// RegisterListener adds a login listener to the UserService
func (s *Service) RegisterListener(listener LoginListener) {
	s.UserLoginListener = append(s.UserLoginListener, listener)
}

// MetricListener is an implementation of LoginListener that sends data to Opensearch
type MetricListener struct{}

// OnLogin handles the login event and send it to Opensearch
func (ml *MetricListener) OnLogin(ctx context.Context, accountID primitive.ObjectID, user *player.PlayerData) {
	metrics.GetDefault().LastLogin(user.ID.Value, user.LastActivity.ToTime())
}

func (s *Service) OnLogin(ctx context.Context, accountID primitive.ObjectID, user *player.PlayerData) {
	// s.triggerWalletObservers(accountID, user.ActiveWallet)
	log.Debug("User logged in: ", user.ID.Value)
}
