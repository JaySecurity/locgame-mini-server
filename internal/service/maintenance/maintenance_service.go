// Copyright © 2022 FuryLion Group LLC. All Rights Reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// without the express permission of FuryLion Group LLC.
// Proprietary and confidential.
//
// Created by Vadim Vlasov on 11.10.22

package maintenance

import (
	"context"
	"sync"
	"time"

	protobuf "google.golang.org/protobuf/proto"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/store"
	"locgame-mini-server/pkg/dto"
	"locgame-mini-server/pkg/dto/maintenance"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/network"
	"locgame-mini-server/pkg/pubsub"
)

// Service is responsible for handling maintenance.
type Service struct {
	config *config.Config
	store  *store.Store

	data            []*maintenance.MaintenanceData
	nextMaintenance *maintenance.MaintenanceData
	server          *network.Server
	clientHandler   *dto.ClientHandler

	mutex sync.Mutex
}

// New creates a new instance of the market services.
func New(config *config.Config, store *store.Store, server *network.Server, clientHandler *dto.ClientHandler) *Service {
	s := new(Service)
	s.config = config
	s.store = store
	s.server = server
	s.clientHandler = clientHandler

	return s
}

func (s *Service) Init() {
	s.updateMaintenancesInfo(false)

	err := pubsub.Subscribe("update-maintenance-info", &UpdateMaintenancesHandler{service: s})
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Service) NextMaintenance() *maintenance.MaintenanceData {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.nextMaintenance
}

func (s *Service) IsMaintenance() bool {
	now := time.Now().UTC().Unix()

	s.mutex.Lock()
	if s.nextMaintenance != nil && s.nextMaintenance.StartDate.Seconds < now && s.nextMaintenance.EndDate.Seconds > now {
		s.mutex.Unlock()
		return true
	}

	s.mutex.Unlock()
	return false
}

func (s *Service) updateMaintenancesInfo(notifyClients bool) {
	s.mutex.Lock()
	s.data, _ = s.store.Maintenance.Get(context.Background())
	s.nextMaintenance = s.findCloserMaintenance(s.data, time.Now().UTC().Unix())

	if notifyClients {
		var data *maintenance.MaintenanceData
		if s.nextMaintenance != nil {
			data = protobuf.Clone(s.nextMaintenance).(*maintenance.MaintenanceData)
		}
		s.mutex.Unlock()

		clients := s.server.GetClients()
		for _, client := range clients {
			s.clientHandler.OnMaintenanceInfoChanged(client, data, nil)
		}
	} else {
		s.mutex.Unlock()
	}
}

func (s *Service) findCloserMaintenance(data []*maintenance.MaintenanceData, now int64) *maintenance.MaintenanceData {
	var closerIndex = -1

	if len(data) > 0 {
		closerIndex = 0
	}

	for i, maintenanceData := range data {
		startDiff := maintenanceData.StartDate.Seconds - now

		if startDiff >= 0 {
			currentStartDiff := data[closerIndex].StartDate.Seconds - now

			if currentStartDiff >= 0 {
				if startDiff < currentStartDiff {
					closerIndex = i
				}
			} else {
				closerIndex = i
			}
		} else if startDiff < 0 && maintenanceData.EndDate.Seconds > now {
			closerIndex = i
		}
	}

	if closerIndex >= 0 {
		return data[closerIndex]
	}

	return nil
}
