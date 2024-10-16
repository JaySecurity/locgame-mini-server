package in_game_store

// import (
// 	"reflect"

// 	"locgame-mini-server/internal/sessions"
// 	"locgame-mini-server/pkg/dto"
// 	storeDto "locgame-mini-server/pkg/dto/store"
// 	"locgame-mini-server/pkg/network"
// 	"locgame-mini-server/pkg/pubsub"
// )

// type TokenPurchaseCompleteHandler struct {
// 	Sessions *sessions.SessionStore
// }

// func (h *TokenPurchaseCompleteHandler) Handle(client *network.Client, clientHandler *dto.ClientHandler, data pubsub.MessageData) {
// 	message := data.(*storeDto.TokenPurchaseResult)
// 	session := h.Sessions.Get(client.Context())
// 	if session != nil {
// 		session.PlayerData.Resources[message.Tokens.ResourceID] += message.Tokens.Quantity
// 		clientHandler.OnTokenPurchaseCompleted(client, message, nil)
// 	}
// }

// func (h *TokenPurchaseCompleteHandler) GetDataType() reflect.Type {
// 	return reflect.TypeOf(storeDto.TokenPurchaseResult{})
// }
