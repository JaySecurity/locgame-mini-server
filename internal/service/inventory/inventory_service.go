package inventory

import (
	"locgame-mini-server/pkg/dto"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/player"
	"locgame-mini-server/pkg/dto/resources"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/stime"

	"github.com/robfig/cron/v3"

	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/sessions"
	"locgame-mini-server/internal/store"
)

type Service struct {
	store  *store.Store
	config *config.Config

	parser cron.Parser

	clientHandler *dto.ClientHandler
}

// New creates a new instance of the reward services.
func New(config *config.Config, store *store.Store) *Service {
	s := new(Service)
	s.config = config
	s.store = store
	s.parser = cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)

	return s
}

func (s *Service) Adjust(session *sessions.Session, reason string, adjustments ...*resources.ResourceAdjustment) ([]*resources.ResourceAdjustment, error) {
	return s.adjust(session, 1, reason, adjustments...)
}

func (s *Service) adjust(session *sessions.Session, multiply int32, reason string, adjustments ...*resources.ResourceAdjustment) ([]*resources.ResourceAdjustment, error) {
	if adjustments == nil || (len(adjustments) == 1 && adjustments[0] == nil) {
		return nil, nil
	}

	var err error
	var result []*resources.ResourceAdjustment
	for _, adjustment := range adjustments {
		result = s.getAdjustResult(adjustment, result, multiply)
	}

	for _, adjustment := range result {
		err = s.processResettableResource(session, adjustment.ResourceID)
		if err != nil {
			return nil, err
		}
	}

	for _, adjustment := range result {
		if s.config.CappedResources.HasHardCapacity(adjustment.ResourceID) {
			capacity := s.config.CappedResources.GetCapacity(adjustment.ResourceID, 1)
			if capacity < session.PlayerData.Resources[adjustment.ResourceID]+adjustment.Quantity {
				adjustment.Quantity += capacity - (session.PlayerData.Resources[adjustment.ResourceID] + adjustment.Quantity)
			}
		}
		// Checking so as not to go into negative values.
		if session.PlayerData.Resources[adjustment.ResourceID]+adjustment.Quantity < 0 {
			adjustment.Quantity -= session.PlayerData.Resources[adjustment.ResourceID] + adjustment.Quantity
		}

		session.PlayerData.Resources[adjustment.ResourceID] += adjustment.Quantity
	}

	err = s.store.Inventory.IncrementResources(session.Context, session.AccountID, result, reason)

	return result, err
}

func (s *Service) UpdateResettableResources(session *sessions.Session) {
	if session.PlayerData.ResettableResources == nil {
		session.PlayerData.ResettableResources = make(map[int32]*resources.ResettableResourceData)
	}
	for _, resource := range s.config.ResettableResources {
		_ = s.processResettableResource(session, resource.ResourceID)
	}
}

func (s *Service) processResettableResource(session *sessions.Session, resourceID int32) error {
	if _, exists := s.config.ResettableResources[resourceID]; !exists {
		return nil
	}

	if _, exists := s.config.CappedResources.Resources[resourceID]; !exists {
		log.Warning("Capped config  for resource ", resourceID, " not found. Reset process skipped.")
		return nil
	}

	cronTime, _ := s.parser.Parse(s.config.ResettableResources[resourceID].ResetTime)
	resourceData, exists := session.PlayerData.ResettableResources[resourceID]

	if !exists {
		resourceData = &resources.ResettableResourceData{
			NextResetTime: &base.Timestamp{Seconds: cronTime.Next(stime.Now(session.Context)).Unix()},
		}
		session.PlayerData.ResettableResources[resourceID] = resourceData
		return s.store.Players.SetData(session.Context, session.AccountID, &player.PlayerData{
			ResettableResources: map[int32]*resources.ResettableResourceData{
				resourceID: session.PlayerData.ResettableResources[resourceID],
			},
		})
	}

	if resourceData.NextResetTime.Seconds > stime.Now(session.Context).Unix() {
		return nil
	}

	capacity := s.config.CappedResources.GetCapacity(resourceID, 1)
	session.PlayerData.ResettableResources[resourceID].NextResetTime = &base.Timestamp{Seconds: cronTime.Next(stime.Now(session.Context)).Unix()}

	playerData := &player.PlayerData{
		ResettableResources: map[int32]*resources.ResettableResourceData{
			resourceID: session.PlayerData.ResettableResources[resourceID],
		},
	}

	if val, exists := session.PlayerData.Resources[resourceID]; !exists || val < capacity {
		session.PlayerData.Resources[resourceID] = capacity
		playerData.Resources = session.PlayerData.Resources
	}

	return s.store.Players.SetData(session.Context, session.AccountID, playerData)
}

func (s *Service) getAdjustResult(adjustment *resources.ResourceAdjustment, result []*resources.ResourceAdjustment, multiply int32) []*resources.ResourceAdjustment {
	result = append(result, &resources.ResourceAdjustment{
		ResourceID: adjustment.ResourceID,
		Quantity:   adjustment.Quantity * multiply,
	})

	return result
}

func (s *Service) InverseAdjust(session *sessions.Session, reason string, adjustments ...*resources.ResourceAdjustment) ([]*resources.ResourceAdjustment, error) {
	return s.adjust(session, -1, reason, adjustments...)
}

func (s *Service) IsEnough(session *sessions.Session, adjustments ...*resources.ResourceAdjustment) bool {
	for _, adjustment := range adjustments {
		if session.PlayerData.Resources[adjustment.ResourceID] < adjustment.Quantity {
			return false
		}
	}

	return true
}

func (s *Service) InverseIsEnough(session *sessions.Session, adjustments ...*resources.ResourceAdjustment) bool {
	for _, adjustment := range adjustments {
		if session.PlayerData.Resources[adjustment.ResourceID] < -adjustment.Quantity {
			return false
		}
	}

	return true
}
