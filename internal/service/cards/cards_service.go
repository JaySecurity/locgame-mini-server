package cards

import (
	"context"
	"locgame-mini-server/internal/blockchain"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/service/shared"
	"locgame-mini-server/internal/sessions"
	"locgame-mini-server/internal/store"
	"locgame-mini-server/pkg/dto"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/cards"
	"locgame-mini-server/pkg/dto/player"
	"locgame-mini-server/pkg/log"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	store         *store.Store
	config        *config.Config
	sessions      *sessions.SessionStore
	blockchain    *blockchain.Blockchain
	clientHandler *dto.ClientHandler

	onDecksChanged []shared.DecksChangeListener
}

// New creates a new instance of the cards services.
func New(config *config.Config, sessions *sessions.SessionStore, store *store.Store, blockchain *blockchain.Blockchain, clientHandler *dto.ClientHandler) *Service {
	s := new(Service)
	s.config = config
	s.store = store
	s.sessions = sessions
	s.blockchain = blockchain
	s.clientHandler = clientHandler

	return s
}

func (s *Service) GetOwnedTokens(id string) ([]string, error) {
	session := s.sessions.Get(id)
	if session.PlayerData.ActiveWallet == "" {
		return []string{}, nil
	}
	tokens, err := s.blockchain.GetTokens(session.PlayerData.ActiveWallet)
	if err != nil {
		return nil, err
	}

	var ownedCards []string
	for _, token := range tokens {
		ownedCards = append(ownedCards, token.String())
	}
	return ownedCards, nil
}

func (s *Service) GetOwnedCards(id string) []string {
	session := s.sessions.Get(id)
	tokens, err := s.GetOwnedTokens(id)
	ownedCards := make([]string, 0, len(tokens))
	if err != nil {
		log.Error("Failed to get tokens from the blockchain network:", err, "Wallet:", session.PlayerData.ActiveWallet)
		return ownedCards
	}
	for _, token := range tokens {
		if len(token) >= 16 {
			archetypeId := GetCardArchetypeByToken(token, true)
			if card := s.config.Cards.CardsByID[archetypeId]; card != nil {
				ownedCards = append(ownedCards, card.ArchetypeID)
			}
		}
	}
	ownedCards = append(ownedCards, session.PlayerData.VirtualCards...)
	return ownedCards
}

func GetCardArchetypeByToken(token string, withEdition bool) string {
	if withEdition {
		return token[1:4] + "-" + token[4:7] + "-" + token[7:10] + "-" + token[10:13] + "-" + token[13:16]
	}
	return token[4:7] + "-" + token[7:10] + "-" + token[10:13] + "-" + token[13:16]
}

func (s *Service) CreateDeck(id string, in *cards.Deck) (*cards.Deck, error) {
	session := s.sessions.Get(id)
	in.ID = &base.ObjectID{Value: primitive.NewObjectID().Hex()}
	if session.PlayerData.ActiveWallet != "" {
		if session.PlayerData.Decks.Decks == nil {
			session.PlayerData.Decks.Decks = make(map[string]*cards.Deck)
		}
		session.PlayerData.Decks.Decks[in.ID.Value] = in
	}
	session.PlayerData.Decks.Decks[in.ID.Value] = in
	ctx := context.Background()
	return in, s.store.Players.SetData(ctx, session.AccountID, &player.PlayerData{Decks: session.PlayerData.Decks})
}

func (s *Service) SetDecks(id string, in *cards.Decks) (*cards.DecksChanges, error) {
	// TODO Validate cards (Unique, Owned)
	session := s.sessions.Get(id)
	session.PlayerData.Decks = in

	defenseDeckChanged := false
	defenseDeck, exists := session.PlayerData.Decks.Decks[session.PlayerData.Decks.Defense]

	if !exists || len(defenseDeck.Cards) != 15 {
		session.PlayerData.Decks.Defense = ""

		for id, deck := range in.Decks {
			if len(deck.Cards) == 15 {
				defenseDeckChanged = true
				session.PlayerData.Decks.Defense = id
				break
			}
		}
	}
	ctx := context.Background()
	err := s.store.Players.ForceSetData(ctx, session.AccountID, &player.PlayerData{Decks: in})

	if len(s.onDecksChanged) > 0 {
		for _, listener := range s.onDecksChanged {
			listener.OnDecksChanged(session, in, defenseDeckChanged)
		}
	}

	if defenseDeckChanged {
		return &cards.DecksChanges{DefenseDeckChanged: defenseDeckChanged, NewDefenseDeck: session.PlayerData.Decks.Defense}, err
	}
	return &cards.DecksChanges{}, err
}

func (s *Service) GetUniqueIdFromArchetypeId(archetypeID string) string {
	parts := strings.Split(archetypeID, "-")
	if len(parts) == 4 {
		return parts[1] + "-" + parts[3] // Unique ID => 'pack'-'cardId'
	}

	if len(parts) == 5 {
		return parts[2] + "-" + parts[4]
	}

	log.Error("Unable to get UniqueID from ArchetypeID")
	return ""
}

func (s *Service) SubscribeDecksChange(listener shared.DecksChangeListener) {
	s.onDecksChanged = append(s.onDecksChanged, listener)
}
