package main

import (
	"context"
	"errors"
	"fmt"
	"locgame-mini-server/internal/blockchain"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/router"
	"locgame-mini-server/internal/service/accounts"
	inGameStore "locgame-mini-server/internal/service/in_game_store"
	"locgame-mini-server/internal/service/payments"
	"locgame-mini-server/internal/sessions"
	"locgame-mini-server/internal/store"
	"locgame-mini-server/internal/version"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/network"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "locgame-mini-server/internal/webhooks"
)

type Service struct {
	store  *store.Store
	config *config.Config

	blockchain *blockchain.Blockchain

	Sessions    *sessions.SessionStore
	Accounts    *accounts.Service
	InGameStore *inGameStore.Service
	Payments    *payments.Service
}

func main() {
	fmt.Println(os.Args)

	log.Init()
	log.SetLogLevel(log.LevelDebug)

	log.Info("Version:", version.RELEASE, "Commit:", version.COMMIT, "Repo:", version.REPO, "Build:", version.BUILD)

	cfg := config.Init("configs/")
	network.Verbose = cfg.NetworkVerboseMode

	dataStore := store.NewStore(cfg)

	r := router.New(cfg, dataStore)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	mux := r.Mux

	addr := fmt.Sprintf(":%d", cfg.HttpPort)
	log.Debugf("API running at https://0.0.0.0%s ...", addr)

	s := &http.Server{Addr: addr, Handler: mux}

	go func() {
		<-c
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()
		err := s.Shutdown(context.Background())
		if err != nil {
			log.Error("Server forced stop. Error:", err)
		}
	}()

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
