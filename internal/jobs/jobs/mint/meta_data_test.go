package mint

import (
	"testing"

	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/cards"
)

func TestWorker_getMetadataJson(t *testing.T) {
	cfg := config.Init("../../../../configs")
	tests := []struct {
		name string
		card *cards.Card
		want string
	}{
		{
			name: "Card 003-006-001-001-035",
			card: cfg.Cards.CardsByID["003-006-001-001-035"],
			want: "{\"name\":\"Regulator\",\"description\":\"Everything is a security\",\"image\":\"https://static-files.locgame.io/cards/8c3e1c53-75bf-4367-90e6-98eda5c62170\",\"external_url\":\"https://locgame.io\",\"attributes\":[{\"trait_type\":\"Edition Set\",\"value\":\"Legends of Crypto Apollo\"},{\"trait_type\":\"Pack\",\"value\":\"Apollo\"},{\"trait_type\":\"Game Rarity\",\"value\":\"001\"},{\"trait_type\":\"Visual Rarity\",\"value\":\"001\"},{\"display_type\":\"boost_number\",\"trait_type\":\"Influence\",\"value\":0},{\"display_type\":\"boost_number\",\"trait_type\":\"Innovation\",\"value\":20},{\"display_type\":\"boost_number\",\"trait_type\":\"Dev Skills\",\"value\":10},{\"display_type\":\"boost_number\",\"trait_type\":\"Community\",\"value\":0},{\"display_type\":\"boost_number\",\"trait_type\":\"Top Wealth\",\"value\":10}]}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Worker{}
			s.Init(cfg, nil, nil)
			metaData := s.getMetadata(tt.card)
			if got := metaData.GetJson(); got != tt.want {
				t.Errorf("GetMetadata():\n%v\n%v", got, tt.want)
			}
		})
	}
}
