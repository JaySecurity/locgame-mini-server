package jobs

import (
	"reflect"

	"locgame-mini-server/pkg/dto/jobs"
	"locgame-mini-server/pkg/pubsub"
)

type TriggerNowHandler struct{}

func (h *TriggerNowHandler) Handle(msg pubsub.MessageData) {
	data := msg.(*jobs.TriggerNowMessage)

	triggerNow(data.JobName)
}

func (h *TriggerNowHandler) GetDataType() reflect.Type {
	return reflect.TypeOf(jobs.TriggerNowMessage{})
}
