package config

import (
	"reflect"

	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/pubsub"
)

type ConfigsReloadHandler struct {
	config *Config
}

func (h *ConfigsReloadHandler) GetDataType() reflect.Type {
	return reflect.TypeOf(base.ConfigsReloadRequest{})
}

func (h *ConfigsReloadHandler) Handle(_ pubsub.MessageData) {
	h.config.Reload()
}
