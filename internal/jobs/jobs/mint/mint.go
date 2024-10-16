package mint

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strings"

	"locgame-mini-server/pkg/dto/cards"
	storeDto "locgame-mini-server/pkg/dto/store"
	"locgame-mini-server/pkg/log"
)

func (w *Worker) mint(ctx context.Context, wallet string, tokens []*storeDto.TokenInfo) (string, error) {
	tx, err := w.blockchain.NftMint(ctx, wallet, tokens)
	return tx, err
}

func (w *Worker) generateTokens(ctx context.Context, qty int64, items []*storeDto.PackItem) []*storeDto.TokenInfo {
	var tokens []*storeDto.TokenInfo

	for i := 0; i < int(qty); i++ {
		for _, item := range items {
			tokenID, err := w.generateCard(ctx, item)
			token := &storeDto.TokenInfo{}
			if err != nil {
				log.Error("Error:", err)
				token.Status = storeDto.TokenStatus_TokenError
			} else {
				token.Token = tokenID
				token.Status = storeDto.TokenStatus_TokenWaitingForMint
			}
			tokens = append(tokens, token)
		}
	}
	return tokens
}

func (w *Worker) generateCard(ctx context.Context, item *storeDto.PackItem) (string, error) {
	var (
		card *cards.Card
		err  error
	)
	if item.PredefinedCardID != "" {
		var ok bool
		card, ok = w.GetConfig().Cards.CardsByID[item.PredefinedCardID]
		if !ok {
			return "", errors.New("card not found: " + item.PredefinedCardID)
		}
	} else {
		iter := 0
		for iter < 10000 {
			card, err = w.randomizeCard(item)
			if err == nil {
				break
			}
			iter++
		}
		if err != nil {
			return "", err
		}
	}

	var mintedCards int64
	for {
		mintedCards, err = w.GetStore().Mint.IncrementMintedCards(ctx, card.ArchetypeID)
		if err != nil {
			log.Error("Failed to get the number of minted cards:", err)
			return "", err
		}
		tokenId := fmt.Sprintf("1%s%09d", strings.Replace(card.ArchetypeID, "-", "", -1), mintedCards)
		tokenIdBigInt := new(big.Int)
		tokenIdBigInt.SetString(tokenId, 10)
		owner, err := w.blockchain.GetTokenOwner(tokenIdBigInt)
		if err != nil {
			log.Error("Failed to get token owner:", err)
		}
		if owner == "0x0000000000000000000000000000000000000000" {
			break
		}
	}

	return fmt.Sprintf("1%s%09d", strings.Replace(card.ArchetypeID, "-", "", -1), mintedCards), nil
}

func (w *Worker) randomizeCard(item *storeDto.PackItem) (*cards.Card, error) {
	var (
		randomGameRarity   cards.GameRarity   = 1
		randomVisualRarity cards.VisualRarity = 1
		probability        float32
	)

	for i := 1; i <= 5; i++ {
		probability += item.RandomizedCard.GameRarities[int32(i)]
	}
	if probability != 1 {
		return nil, errors.New(fmt.Sprint("Game Rarity probability does not add up to 1: ", probability, "Item: ", item))
	}

	probability = 0
	randomValue := rand.Float32()

	for i := 5; i > 1; i-- {
		if item.RandomizedCard.GameRarities[int32(i)] > 0 && randomValue < probability+item.RandomizedCard.GameRarities[int32(i)] {
			randomGameRarity = cards.GameRarity(i)
			probability += item.RandomizedCard.GameRarities[int32(i)]
		}
	}

	probability = 0

	for i := 1; i <= 5; i++ {
		probability += item.RandomizedCard.VisualRarities[int32(i)]
	}
	if probability != 1 {
		return nil, errors.New(fmt.Sprint("Visual Rarity probability does not add up to 1: ", probability, "Item: ", item))
	}

	probability = 0
	randomValue = rand.Float32()

	for i := 5; i > 1; i-- {
		if item.RandomizedCard.VisualRarities[int32(i)] > 0 && randomValue < probability+item.RandomizedCard.VisualRarities[int32(i)] {
			randomVisualRarity = cards.VisualRarity(i)
			probability += item.RandomizedCard.VisualRarities[int32(i)]
		}
	}

	card := w.GetConfig().Cards.FindRandomCard(item.RandomizedCard.Set, item.RandomizedCard.Pack, randomGameRarity, randomVisualRarity)
	if card == nil {
		return nil, errors.New(fmt.Sprintf("Card not found (%v): %v(%v) %v(%v) GameRarity: %v VisualRarity: %v", fmt.Sprintf("%03d-%03d-%03d-%03d", item.RandomizedCard.Set, item.RandomizedCard.Pack, randomGameRarity, randomVisualRarity), w.GetConfig().Cards.Sets[item.RandomizedCard.Set], item.RandomizedCard.Set, w.GetConfig().Cards.Packs[item.RandomizedCard.Pack], item.RandomizedCard.Pack, randomGameRarity, randomVisualRarity))
	}
	return card, nil
}
