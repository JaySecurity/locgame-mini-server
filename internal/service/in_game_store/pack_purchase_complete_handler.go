package in_game_store

import (
	"reflect"

	"locgame-mini-server/pkg/dto"
	storeDto "locgame-mini-server/pkg/dto/store"
	"locgame-mini-server/pkg/network"
	"locgame-mini-server/pkg/pubsub"
)

type PackPurchaseCompleteHandler struct{}

func (h *PackPurchaseCompleteHandler) Handle(client *network.Client, clientHandler *dto.ClientHandler, msg pubsub.MessageData) {
	data := msg.(*storeDto.PackPurchaseResult)
	clientHandler.OnMintOfPackCompleted(client, &storeDto.PackPurchaseResult{OrderID: data.OrderID, Status: data.Status}, nil)
}

func (h *PackPurchaseCompleteHandler) GetDataType() reflect.Type {
	return reflect.TypeOf(storeDto.PackPurchaseResult{})
}
