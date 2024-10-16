package mint

import (
	"context"
	"fmt"
	"runtime"

	"locgame-mini-server/internal/service/cards"
	storeDto "locgame-mini-server/pkg/dto/store"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/pubsub"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func (w *Worker) onMintOrderRequestReceived(request *storeDto.MintJobRequest) {
	ctx := context.Background()
	order, err := w.GetStore().Orders.Get(ctx, request.ID.Value)
	if err != nil {
		log.Error("Failed to get order:", request.ID, "Error:", err)
		return
	}

	product, ok := w.GetConfig().Products.ProductsByID[order.ProductID]
	if !ok {
		w.setOrderFailedState(ctx, order, errors.New("Product not found:"+order.ProductID))
		return
	}

	switch product.Type {
	case storeDto.ProductType_SpecialOffer:
		fallthrough
	case storeDto.ProductType_PackOfCards:
		pack, ok := w.GetConfig().InGameStore.PackByID[product.Value]
		if !ok {
			w.setOrderFailedState(ctx, order, errors.New("Pack not found: "+product.Value))
			return
		}
		wallet, err := w.GetStore().Players.GetWalletByPlayerID(ctx, order.BuyerID)
		if err != nil {
			w.setOrderFailedState(ctx, order, errors.Wrap(err, "Failed to get wallet by buyer id"))
			return
		}
		tokens := w.generateTokens(ctx, order.Quantity, pack.Items)
		w.uploadAllMetadata(tokens)
		tx, err := w.mint(ctx, wallet, tokens)
		if err != nil {
			w.setOrderFailedState(ctx, order, errors.Wrap(err, "Failed to send transaction for minting cards"))
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

					err = pubsub.SendToPlayer(order.BuyerID.Value, &storeDto.PackPurchaseResult{OrderID: order.ID, Status: order.Status})
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

func (w *Worker) setOrderFailedState(ctx context.Context, order *storeDto.Order, err error) {
	order.Error = err.Error()
	order.Status = storeDto.OrderStatus_Failed
	_ = w.GetStore().Orders.Update(ctx, order)
}

func (w *Worker) uploadAllMetadata(tokens []*storeDto.TokenInfo) {
	for _, token := range tokens {
		if token.Token == "" {
			continue
		}
		cardID := cards.GetCardArchetypeByToken(token.Token, true)
		card, ok := w.GetConfig().Cards.CardsByID[cardID]
		if ok {
			metaData := w.getMetadata(card)
			err := w.uploadMetaDataToS3(token.Token, metaData)
			if err != nil {
				log.Error("Failed to upload metadata to AWS S3:", token.Token, "Error:", err)
			}
		}
	}
}
