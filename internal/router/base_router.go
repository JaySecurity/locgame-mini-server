package router

import (
	"locgame-mini-server/internal/config/resources"
	"locgame-mini-server/pkg/dto/game"
)

// GetConfigs returns all the configurations required for the client to work.
func (r *Router) GetConfigs() (*game.Configs, error) {

	return &game.Configs{
		Resources:           resources.Resources,
		ResourceCategories:  resources.Categories,
		ExtraRewards:        r.config.ExtraRewards.GetDTO(),
		CappedResources:     r.config.CappedResources.Resources,
		ResettableResources: r.config.ResettableResources,
		Cards:               r.config.Cards.CardsByID,
		StoryMode:           r.config.StoryMode,
		TrophyRoadRewards:   r.config.TrophyRoadRewards,
		Leagues:             r.config.Leagues,
		FriendlyMatch:       r.config.FriendlyMatch,
		NextMaintenance:     r.Maintenance.NextMaintenance(),
		MatchTimer:          r.config.MatchTimerConfig.GetDTO(),
	}, nil
}

// GetPlayerData returns all player data
// func (r *Router) GetPlayerData(client *network.Client, _ *base.Empty) (*player.PlayerDataResponse, error) {
// 	ctx := client.Context()
// 	session := r.sessions.Get(id)
// 	session.OnGetData(ctx)
// 	return &player.PlayerDataResponse{
// 		Data:       session.PlayerData.PlayerData,
// 		OwnedCards: session.OwnedCards,
// 	}, nil
// }
