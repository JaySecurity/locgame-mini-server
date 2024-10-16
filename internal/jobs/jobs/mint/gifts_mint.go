package mint

import (
	"context"
	"fmt"
	"runtime"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	storeDto "locgame-mini-server/pkg/dto/store"
	"locgame-mini-server/pkg/log"
)

func (w *Worker) onMintGiftRequestReceived(request *storeDto.MintJobRequest) {
	ctx := context.Background()
	gift, err := w.GetStore().Gifts.Get(ctx, request.ID.Value)
	if err != nil {
		log.Error("Failed to get gift:", request.ID, "Error:", err)
		return
	}

	if !common.IsHexAddress(gift.Wallet) {
		w.setGiftFailedState(ctx, gift, errors.New("Invalid wallet address"))
		return
	}

	var items []*storeDto.PackItem
	for _, card := range gift.Cards {
		items = append(items, &storeDto.PackItem{PredefinedCardID: card})
	}
	tokens := w.generateTokens(ctx, 1, items)
	w.uploadAllMetadata(tokens)
	tx, err := w.mint(ctx, gift.Wallet, tokens)
	if err != nil {
		w.setGiftFailedState(ctx, gift, errors.Wrap(err, "Failed to send transaction for minting cards"))
	} else {
		gift.Status = storeDto.OrderStatus_InProgress
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
				gift.Status = storeDto.OrderStatus_Completed
				if gift.Error != "" {
					gift.Error = ""
					w.GetStore().Gifts.ClearErrorMessage(ctx, gift.ID)
				}

				_ = w.GetStore().Gifts.Update(ctx, gift)
			} else {
				gift.Status = storeDto.OrderStatus_Failed
				if err == nil {
					msg, err := w.blockchain.GetFailingMessage(transaction)
					if err != nil {
						msg = err.Error()
					}
					gift.Error = msg
					log.Error("Failed to mint cards.", msg)
				} else {
					gift.Error = err.Error()
					log.Error("Failed to mint cards: ", err)
				}
				_ = w.GetStore().Gifts.Update(ctx, gift)
			}
		}()
	}

	gift.OperationHash = tx
	_ = w.GetStore().Gifts.Update(ctx, gift)
}

func (w *Worker) setGiftFailedState(ctx context.Context, gift *storeDto.Gift, err error) {
	gift.Error = err.Error()
	gift.Status = storeDto.OrderStatus_Failed
	_ = w.GetStore().Gifts.Update(ctx, gift)
}
