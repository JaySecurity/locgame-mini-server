package mint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"locgame-mini-server/pkg/dto/cards"
)

type MetaData struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Image       string           `json:"image"`
	ExternalURL string           `json:"external_url"`
	Attributes  []CardAttributes `json:"attributes"`
}

func (m MetaData) GetJson() string {
	data, _ := json.Marshal(m)
	return string(data)
}

type CardAttributes struct {
	DisplayType string         `json:"display_type,omitempty"`
	TraitType   string         `json:"trait_type"`
	Value       AttributeValue `json:"value"`
}

type AttributeValue struct {
	IntValue int32
	StrValue string
}

func (v *AttributeValue) MarshalJSON() ([]byte, error) {
	if v.StrValue != "" {
		return json.Marshal(v.StrValue)
	}

	return json.Marshal(v.IntValue)
}

func (w *Worker) getMetadata(card *cards.Card) MetaData {
	set, pack := w.extractSetAndPackFromArchetypeID(card.ArchetypeID)
	return MetaData{
		Name:        card.Name,
		Description: card.Description,
		Image:       card.Image,
		ExternalURL: "https://locgame.io",
		Attributes: []CardAttributes{
			{
				TraitType: "Edition Set",
				Value:     AttributeValue{StrValue: w.GetConfig().Cards.Sets[set]},
			},
			{
				TraitType: "Pack",
				Value:     AttributeValue{StrValue: w.GetConfig().Cards.Packs[pack]},
			},
			{
				TraitType: "Game Rarity",
				Value:     AttributeValue{StrValue: fmt.Sprintf("%03d", card.Properties.GameRarity)},
			},
			{
				TraitType: "Visual Rarity",
				Value:     AttributeValue{StrValue: fmt.Sprintf("%03d", card.Properties.VisualRarity)},
			},
			{
				DisplayType: "boost_number",
				TraitType:   "Influence",
				Value:       AttributeValue{IntValue: card.Boosts.Influence},
			},
			{
				DisplayType: "boost_number",
				TraitType:   "Innovation",
				Value:       AttributeValue{IntValue: card.Boosts.Innovation},
			},
			{
				DisplayType: "boost_number",
				TraitType:   "Dev Skills",
				Value:       AttributeValue{IntValue: card.Boosts.DevSkills},
			},
			{
				DisplayType: "boost_number",
				TraitType:   "Community",
				Value:       AttributeValue{IntValue: card.Boosts.Community},
			},
			{
				DisplayType: "boost_number",
				TraitType:   "Top Wealth",
				Value:       AttributeValue{IntValue: card.Boosts.TopWealth},
			},
		},
	}
}

func (w *Worker) extractSetAndPackFromArchetypeID(cardID string) (int32, int32) {
	parts := strings.Split(cardID, "-")
	set, _ := strconv.Atoi(parts[0])
	pack, _ := strconv.Atoi(parts[1])
	return int32(set), int32(pack)
}

const metaDataBucketName = "meta.locgame.io"

func (w *Worker) uploadMetaDataToS3(token string, metadata MetaData) error {
	data := []byte(metadata.GetJson())
	params := &s3.PutObjectInput{
		Bucket:        aws.String(metaDataBucketName),
		Key:           aws.String(w.GetConfig().Blockchain.Contracts.NFT + "/" + token),
		Body:          bytes.NewReader(data),
		ContentLength: aws.Int64(int64(len(data))),
	}

	_, err := w.s3.PutObject(params)
	return err
}
