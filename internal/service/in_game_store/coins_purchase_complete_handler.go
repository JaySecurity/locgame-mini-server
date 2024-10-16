package in_game_store

// import (
// 	"reflect"

// 	"locgame-mini-server/internal/sessions"
// 	"locgame-mini-server/pkg/dto"
// 	storeDto "locgame-mini-server/pkg/dto/store"
// 	"locgame-mini-server/pkg/network"
// 	"locgame-mini-server/pkg/pubsub"
// )

// type CoinsPurchaseCompleteHandler struct {
// 	Sessions *sessions.SessionStore
// }

// func (h *CoinsPurchaseCompleteHandler) Handle(client *network.Client, clientHandler *dto.ClientHandler, data pubsub.MessageData) {
// 	message := data.(*storeDto.CoinsPurchaseResult)
// 	session := h.Sessions.Get(client.Context())
// 	if session != nil {
// 		session.PlayerData.Resources[message.Coins.ResourceID] += message.Coins.Quantity
// 		clientHandler.OnCoinsPurchaseCompleted(client, message, nil)
// 	}
// }

// func (h *CoinsPurchaseCompleteHandler) GetDataType() reflect.Type {
// 	return reflect.TypeOf(storeDto.CoinsPurchaseResult{})
// }
