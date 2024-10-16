package in_game_store

import (
	"locgame-mini-server/pkg/dto"
	"locgame-mini-server/pkg/network"
	"locgame-mini-server/pkg/pubsub"
	"reflect"

	storeDto "locgame-mini-server/pkg/dto/store"
)

type CardUpgradeCompleteHandler struct{}

func (h *CardUpgradeCompleteHandler) Handle(
	client *network.Client,
	clientHandler *dto.ClientHandler,
	msg pubsub.MessageData,
) {
	data := msg.(*storeDto.CardUpgradeResult)
	clientHandler.OnCardUpgradeCompleted(
		client,
		&storeDto.CardUpgradeResult{
			OrderID:        data.OrderID,
			OriginalCardID: data.OriginalCardID,
			NewCardId:      data.NewCardId,
		},
		nil,
	)
}

func (h *CardUpgradeCompleteHandler) GetDataType() reflect.Type {
	return reflect.TypeOf(storeDto.CardUpgradeResult{})
}
