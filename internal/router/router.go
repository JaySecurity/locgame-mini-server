package router

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"locgame-mini-server/internal/blockchain"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/service/accounts"
	inGameStore "locgame-mini-server/internal/service/in_game_store"
	"locgame-mini-server/internal/service/inventory"
	"locgame-mini-server/internal/service/maintenance"
	"locgame-mini-server/internal/service/payments"
	"locgame-mini-server/internal/sessions"
	"locgame-mini-server/internal/store"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/stime"
)

// Router contains information about all possible routes with the current services.
// Is the central gateway for the services.
// Consolidates domain services.
// Stores connection with storages and configurations to provide them to domain services.
type Router struct {
	store  *store.Store
	config *config.Config
	Mux    *http.ServeMux

	blockchain *blockchain.Blockchain

	Sessions *sessions.SessionStore
	Accounts *accounts.Service

	InGameStore *inGameStore.Service
	Payments    *payments.Service
	Maintenance *maintenance.Service
	Inventory   *inventory.Service
}

// New creates a new instance of Router.
func New(cfg *config.Config, store *store.Store) *Router {
	s := new(Router)
	s.config = cfg
	s.store = store
	s.Mux = http.NewServeMux()

	rand.Seed(time.Now().UTC().Unix())

	var err error
	s.blockchain, err = blockchain.Connect(s.config.Blockchain)
	if err != nil {
		log.Fatal(err)
	}

	s.Sessions = sessions.NewSessionStore(store.Players, s.config)
	s.Sessions.OnGetData(func(id string) {
		session := s.Sessions.Get(id)

		session.PlayerData.Online = true
		s.Accounts.SetOnlineState(id, true)

	})
	s.Inventory = inventory.New(s.config, s.store)
	s.Accounts = accounts.New(s.config, s.Sessions, s.store)
	s.Payments = payments.New(s.config, s.Sessions, s.store, s.blockchain)
	s.InGameStore = inGameStore.New(s.config, s.Sessions, s.store, s.blockchain, s.Payments, s.Inventory)

	stime.Init(&sessions.ServerTime{Session: s.Sessions})

	s.Accounts.OnLogout = func(ctx context.Context) {
		// session := s.Sessions.TryGet(ctx)
		// if session != nil {
		// 	s.OnDisconnectClient(session.GetClient())
		// }
	}

	s.InGameStore.Init()

	// Add Middleware
	s.HandleAccountRoutes()
	s.HandleStoreRoutes()
	s.HandlePaymentRoutes()
	s.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Welcome to LoC REST API Server."))
	})
	return s
}

// OnDisconnectClient is called when there is a disconnect from the client.
func (r *Router) OnDisconnectClient(id string) {
	session := r.Sessions.TryGet(id)
	if session != nil {
		r.Sessions.Remove(id)
	}
}
