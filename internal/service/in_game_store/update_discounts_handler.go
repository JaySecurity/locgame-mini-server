// Copyright © 2022 FuryLion Group LLC. All Rights Reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// without the express permission of FuryLion Group LLC.
// Proprietary and confidential.
//
// Created by Vadim Vlasov on 25.11.22

package in_game_store

import (
	"reflect"

	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/pubsub"
)

type UpdateDiscountsHandler struct {
	pubsub.BaseUnsubscribableHandler

	service *Service
}

func (s *UpdateDiscountsHandler) Handle(_ pubsub.MessageData) {
	s.service.updateDiscounts(true)
}

func (s *UpdateDiscountsHandler) GetDataType() reflect.Type {
	return reflect.TypeOf(base.Empty{})
}
