// Copyright © 2022 FuryLion Group LLC. All Rights Reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// without the express permission of FuryLion Group LLC.
// Proprietary and confidential.
//
// Created by Vadim Vlasov on 11.10.22

package maintenance

import (
	"reflect"

	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/pubsub"
)

type UpdateMaintenancesHandler struct {
	pubsub.BaseUnsubscribableHandler

	service *Service
}

func (s *UpdateMaintenancesHandler) Handle(_ pubsub.MessageData) {
	s.service.updateMaintenancesInfo(true)
}

func (s *UpdateMaintenancesHandler) GetDataType() reflect.Type {
	return reflect.TypeOf(base.Empty{})
}
