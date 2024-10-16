package mint

import (
	"testing"

	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/cards"
	"locgame-mini-server/pkg/log"
)

func TestWorker_randomizeCard(t *testing.T) {
	cfg := config.Init("../../../../configs")
	w := Worker{}
	w.Init(cfg, nil, nil)

	for _, pack := range cfg.InGameStore.PackByID {
		for _, item := range pack.Items {
			if item.RandomizedCard == nil && item.PredefinedCardID != "" {
				return
			}
			i := 0
			var (
				card *cards.Card
				err  error
			)
			for i < 10000 {
				card, err = w.randomizeCard(item)
				if err == nil {
					break
				}
				_ = item
				i++
			}

			if err != nil {
				log.Error("Pack:", pack.ID, "Error:", err)
			}
			_ = card
		}
	}
}

func TestWorker_randomizeCard2(t *testing.T) {
	cfg := config.Init("../../../../configs")
	w := Worker{}
	w.Init(cfg, nil, nil)
	item := cfg.InGameStore.PackByID["dj_soda_legendary"].Items[len(cfg.InGameStore.PackByID["dj_soda_legendary"].Items)-1]
	log.Debug(item)
	card, err := w.randomizeCard(item)
	if err != nil {
		t.Error(err)
	}
	_ = card
}
