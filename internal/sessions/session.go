package sessions

import (
	"context"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/store"
	"locgame-mini-server/pkg/dto/base"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Session stores session data.
type Session struct {
	Context   context.Context
	AccountID primitive.ObjectID

	PlayerData  *store.PlayerData
	MatchData   *store.MatchData
	FriendsData *store.FriendsData

	OwnedCards []string

	LastArenaOpponent *base.ObjectID
	FriendlyMatchInfo struct {
		IsInitiator bool
		Opponent    *base.ObjectID
		Stake       int32
	}

	QuickMatchInfo struct {
		Opponent    *base.ObjectID
		Stake       int32
		IsInitiator bool
	}

	SessionID         string
	DuplicatedSession bool

	OnGetData      func(id string)
	Config         *config.Config
	IsMetaMaskUser bool
}

// GetData allows getting data for the current session from the database.
func (s *Session) GetData(id string) error {
	err := s.PlayerData.Update(s.Context)

	if err == nil && s.OnGetData != nil {
		s.OnGetData(id)
	}
	return err
}
