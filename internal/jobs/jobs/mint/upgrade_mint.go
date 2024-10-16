package mint

import (
	"context"
	"fmt"
	"locgame-mini-server/internal/service/cards"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/pubsub"
	"runtime"
	"strings"
	"time"

	storeDto "locgame-mini-server/pkg/dto/store"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func (w *Worker) onMintUpgradeOrderRequestReceived(request *storeDto.MintJobRequest) {
	ctx := context.Background()
	order, err := w.GetStore().Orders.Get(ctx, request.ID.Value)
	if err != nil {
		log.Error("Failed to get order:", request.ID, "Error:", err)
		return
	}

	switch order.ProductType {
	case storeDto.ProductType_CardUpgrade:
		cardIds := strings.Split(order.ProductID, "|")
		newCardId := cardIds[1]
		items := []*storeDto.PackItem{{PredefinedCardID: newCardId}}
		wallet, err := w.GetStore().Players.GetWalletByPlayerID(ctx, order.BuyerID)
		if err != nil {
			w.setOrderFailedState(ctx, order, errors.Wrap(err, "Failed to get wallet by buyer id"))
			return
		}
		tokens := w.generateTokens(ctx, order.Quantity, items)
		w.uploadAllMetadata(tokens)
		tx, err := w.mint(ctx, wallet, tokens)
		if err != nil {
			w.setOrderFailedState(
				ctx,
				order,
				errors.Wrap(err, "Failed to send transaction for minting cards"),
			)
		} else {
			order.Status = storeDto.OrderStatus_InProgress
			transaction := common.HexToHash(tx)
			go func() {
				defer func() {
					if r := recover(); r != nil {
						const size = 64 << 10
						buf := make([]byte, size)
						buf = buf[:runtime.Stack(buf, false)]
						err, ok := r.(error)
						if !ok {
							err = fmt.Errorf("%v", r)
						}
						w.GetLogger().Errorf("Panic: %v\n%s", err, buf)
					}
				}()
				result, err := w.blockchain.GetTransactionResult(context.Background(), w.blockchain.PolygonClient, transaction, 0)
				if result {
					order.Status = storeDto.OrderStatus_Completed
					if order.Error != "" {
						order.Error = ""
						w.GetStore().Orders.ClearErrorMessage(ctx, order.ID)
					}

					for _, token := range tokens {
						order.Cards = append(order.Cards, cards.GetCardArchetypeByToken(token.Token, true))
					}
					time.Sleep(30 * time.Second)
					err = pubsub.SendToPlayer(order.BuyerID.Value, &storeDto.CardUpgradeResult{OrderID: order.ID, OriginalCardID: cardIds[0], NewCardId: cardIds[1]})
					if err != nil {
						log.Error(err)
					}

					_ = w.GetStore().Orders.Update(ctx, order)
				} else {
					order.Status = storeDto.OrderStatus_Failed
					if err == nil {
						msg, err := w.blockchain.GetFailingMessage(transaction)
						if err != nil {
							msg = err.Error()
						}
						order.Error = msg
						log.Error("Failed to mint cards.", msg)
					} else {
						order.Error = err.Error()
						log.Error("Failed to mint cards: ", err)
					}
					_ = w.GetStore().Orders.Update(ctx, order)
				}
			}()
		}

		order.OperationHash = tx
		_ = w.GetStore().Orders.Update(ctx, order)
	}
}
