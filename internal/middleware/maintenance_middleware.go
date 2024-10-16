// Copyright © 2022 FuryLion Group LLC. All Rights Reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// without the express permission of FuryLion Group LLC.
// Proprietary and confidential.
//
// Created by Vadim Vlasov on 11.10.22

package middleware

import (
	"context"
	"reflect"

	"locgame-mini-server/internal/service/maintenance"
	"locgame-mini-server/pkg/dto/errors"
	maintenanceDto "locgame-mini-server/pkg/dto/maintenance"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/network"
)

type MaintenanceMiddleware struct {
	handler network.ServerHandler
	service *maintenance.Service
}

type MaintenanceStore interface {
	Get(ctx context.Context) ([]*maintenanceDto.MaintenanceData, error)
}

func NewMaintenanceMiddleware(handler network.ServerHandler, service *maintenance.Service) *MaintenanceMiddleware {
	return &MaintenanceMiddleware{handler: handler, service: service}
}

func (h *MaintenanceMiddleware) Serve(client *network.Client, packet *network.Packet, arg interface{}) {
	if !h.service.IsMaintenance() {
		h.handler.Serve(client, packet, arg)
		return
	}

	err := errors.ErrMaintenanceMode
	packet.Payload = nil
	packet.Error = &network.Error{ErrorCode: errors.ErrorsByCode[err], Description: err.Error()}
	err = client.Stream.WritePacket(packet)
	if err != nil {
		log.Error("Failed to send error to client:", err)
		client.Stream.Close()
	}
}

func (h *MaintenanceMiddleware) GetArgTypeByMethodID(methodID uint16) reflect.Type {
	return h.handler.GetArgTypeByMethodID(methodID)
}

func (h *MaintenanceMiddleware) GetMethodNameByID(methodID uint16) string {
	return h.handler.GetMethodNameByID(methodID)
}

func (h *MaintenanceMiddleware) Validate(methodID uint16) bool {
	return h.handler.Validate(methodID)
}
