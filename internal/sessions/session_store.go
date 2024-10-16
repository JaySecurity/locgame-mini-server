package sessions

import (
	"context"
	"runtime"
	"sync"

	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/store"
	"locgame-mini-server/pkg/log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SessionStore is the player's session data store.
type SessionStore struct {
	sessions            map[string]*Session
	sessionsByAccountID map[primitive.ObjectID]*Session

	sessionMutex             *sync.RWMutex
	sessionsByAccountIDMutex *sync.RWMutex

	playerDataStore store.PlayersStoreInterface

	onGetDataFunc       func(id string)
	onRemoveSessionFunc func(id string)
	config              *config.Config
}

// NewSessionStore allows to create a new instance of the session storage.
func NewSessionStore(playerDataStore store.PlayersStoreInterface, cfg *config.Config) *SessionStore {
	return &SessionStore{
		sessions:            make(map[string]*Session),
		sessionsByAccountID: make(map[primitive.ObjectID]*Session),

		playerDataStore: playerDataStore,

		sessionMutex:             &sync.RWMutex{},
		sessionsByAccountIDMutex: &sync.RWMutex{},
		config:                   cfg,
	}
}

// Create allows to create a new session in the session storage.
func (s *SessionStore) Create(accountID primitive.ObjectID) *Session {
	session := new(Session)
	session.Context = context.Background()
	session.SessionID = uuid.New().String()

	s.sessionMutex.Lock()
	s.sessions[session.SessionID] = session
	s.sessionMutex.Unlock()
	s.sessionsByAccountIDMutex.Lock()
	s.sessionsByAccountID[accountID] = session
	s.sessionsByAccountIDMutex.Unlock()

	session.AccountID = accountID

	session.PlayerData = store.NewPlayerData(s.playerDataStore, accountID, s.config)
	session.MatchData = new(store.MatchData)

	session.OnGetData = s.onGetDataFunc
	return session
}

// OnGetData allows to specify a function that will be called when trying to get session data (initial data) from the database.
func (s *SessionStore) OnGetData(fn func(id string)) {
	s.onGetDataFunc = fn
}

// OnRemoveSession allows to specify a function that will be called when the session of their session storage is deleted.
func (s *SessionStore) OnRemoveSession(fn func(id string)) {
	s.onRemoveSessionFunc = fn
}

// Get allows to get a session from a specific context.
func (s *SessionStore) get(id string, warn bool) *Session {
	s.sessionMutex.RLock()
	if session, ok := s.sessions[id]; ok {
		s.sessionMutex.RUnlock()
		return session
	}

	const size = 64 << 10
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]

	if warn {
		log.Warning("A session was accessed which did not exist before. Before accessing the session, create in advance by calling Create(ctx, accountID)!\n", string(buf))
	}
	s.sessionMutex.RUnlock()
	return nil
}

// Get allows to get a session from a specific context.
func (s *SessionStore) Get(id string) *Session {
	return s.get(id, true)
}

// TryGet allows to get a session from a specific context.
func (s *SessionStore) TryGet(id string) *Session {
	return s.get(id, false)
}

// TryGetByAccountID allows to get a session by account ID.
func (s *SessionStore) TryGetByAccountID(accountID primitive.ObjectID) (*Session, bool) {
	s.sessionsByAccountIDMutex.RLock()
	session, ok := s.sessionsByAccountID[accountID]
	s.sessionsByAccountIDMutex.RUnlock()
	return session, ok
}

// Remove allows to remove a session from the session storage for a specific context.
func (s *SessionStore) Remove(id string) {
	s.sessionMutex.Lock()
	s.sessionsByAccountIDMutex.Lock()
	if session, ok := s.sessions[id]; ok {
		if s.onRemoveSessionFunc != nil {
			s.onRemoveSessionFunc(id)
		}
		delete(s.sessionsByAccountID, session.AccountID)
	}
	delete(s.sessions, id)
	s.sessionMutex.Unlock()
	s.sessionsByAccountIDMutex.Unlock()
}
