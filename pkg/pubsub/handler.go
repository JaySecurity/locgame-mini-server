package pubsub

import (
	"reflect"

	"github.com/nats-io/nats.go"
)

type Handler interface {
	Handle(data MessageData)
	GetDataType() reflect.Type
}

type UnsubscribableReplyHandler interface {
	Handle(data MessageData) (MessageData, error)
	GetDataType() reflect.Type

	SetSubscription(subs *nats.Subscription)
	Unsubscribe() error
}

type UnsubscribableHandler interface {
	Handler

	SetSubscription(subs *nats.Subscription)
	Unsubscribe() error
}

type BaseUnsubscribableHandler struct {
	subscription *nats.Subscription
}

func (h *BaseUnsubscribableHandler) SetSubscription(subs *nats.Subscription) {
	h.subscription = subs
}

func (h *BaseUnsubscribableHandler) Unsubscribe() error {
	return h.subscription.Unsubscribe()
}
