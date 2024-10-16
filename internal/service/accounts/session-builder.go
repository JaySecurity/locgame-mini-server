package accounts

import (
	"context"
	"fmt"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/sessions"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/cards"
	"locgame-mini-server/pkg/dto/player"
	"locgame-mini-server/pkg/stime"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionBuilder struct {
	ctx             context.Context
	s               *Service
	cfg             *config.Config
	accountID       primitive.ObjectID
	session         *sessions.Session
	err             error
	CognitoUsername string
	Email           string
	Wallet          string
	IsMetaMaskUser  bool
}

func (s *Service) NewSessionBuilder(ctx context.Context, cfg *config.Config) *SessionBuilder {
	return &SessionBuilder{
		ctx: ctx,
		s:   s,
		cfg: cfg,
	}
}

func (b *SessionBuilder) SetCognitoUsername(username string) *SessionBuilder {
	b.CognitoUsername = username
	return b
}

func (b *SessionBuilder) SetIsMetaMaskUser(isMetaMask bool) *SessionBuilder {
	b.IsMetaMaskUser = isMetaMask
	return b
}

func (b *SessionBuilder) SetEmailAddress(email string) *SessionBuilder {
	b.Email = email
	return b
}

func (b *SessionBuilder) SetWallet(wallet string) *SessionBuilder {
	b.Wallet = wallet
	return b
}

func (b *SessionBuilder) CreateSession() *SessionBuilder {
	if b.err != nil {
		return b
	}
	if b.CognitoUsername != "" {
		b.accountID, b.err = b.s.store.Players.GetAccountIDByCognitoUsername(b.ctx, b.CognitoUsername)
	} else if b.Wallet != "" {
		b.accountID, b.err = b.s.store.Players.GetAccountIDByWallet(b.ctx, b.Wallet)
	} else if b.Email != "" {
		b.accountID, b.err = b.s.store.Players.GetAccountIDByEmail(b.ctx, b.Email)
	}
	if b.err == mongo.ErrNoDocuments {
		b.createAccount()
	}
	b.createSession()
	return b
}

// TODO: Add Beginner deck  Here
func (b *SessionBuilder) createAccount() {
	v := b.cfg.VirtualCards
	virtualCards := v.GetAllFreeCardsIDs()
	deckId := v.DeckID
	beginnerDeck := &cards.Deck{
		Name: "Beginner",
		ID: &base.ObjectID{
			Value: deckId,
		},
		Cards:    virtualCards,
		DeckType: cards.DeckType_Free,
	}
	data := &player.PlayerData{
		CognitoUsername: b.CognitoUsername,
		Email:           b.Email,
		ActiveWallet:    b.Wallet,
		Decks:           &cards.Decks{Decks: make(map[string]*cards.Deck), Active: beginnerDeck.ID.Value, Defense: beginnerDeck.ID.Value},
		Name:            fmt.Sprintf("Legend_%06d", rand.Intn(999999)),
		LastActivity:    &base.Timestamp{Seconds: stime.RealTime().Unix()},
		CreatedAt:       &base.Timestamp{Seconds: stime.RealTime().Unix()},
		VirtualCards:    virtualCards,
	}
	data.Decks.Decks[deckId] = beginnerDeck
	b.accountID, b.err = b.s.store.Players.RegisterAccount(b.ctx, data)
	if b.err == nil {
		err := b.s.store.Players.SetData(b.ctx, b.accountID, data)
		if err != nil {
			b.err = err
		}
	}
}

func (b *SessionBuilder) createSession() {
	if b.err == nil {
		b.session = b.s.sessions.Create(b.accountID)
		b.session.IsMetaMaskUser = b.IsMetaMaskUser
		b.err = b.session.GetData(b.session.SessionID)
	}

	for _, listener := range b.s.UserLoginListener {
		listener.OnLogin(b.ctx, b.accountID, b.session.PlayerData.PlayerData)
	}
}

func (b *SessionBuilder) Build() (*sessions.Session, error) {
	return b.session, b.err
}
