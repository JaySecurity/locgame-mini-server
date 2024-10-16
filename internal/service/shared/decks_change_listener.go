package shared

import (
	"locgame-mini-server/internal/sessions"
	"locgame-mini-server/pkg/dto/cards"
)

type DecksChangeListener interface {
	OnDecksChanged(session *sessions.Session, decks *cards.Decks, defenseDeckChanged bool)
}
