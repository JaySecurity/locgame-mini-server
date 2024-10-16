package pubsub

import (
	"errors"
	"reflect"
	"runtime"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	protobuf "google.golang.org/protobuf/proto"
	"locgame-mini-server/pkg/dto"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/log"
)

// PubSub is a message system for communication between services.
type PubSub struct {
	nats          NatsBroker
	addr          string
	clientHandler *dto.ClientHandler

	mutex              sync.Mutex
	subscriptions      map[*nats.Subscription]*Subscription
	replySubscriptions map[*nats.Subscription]*ReplySubscription
}

// Subscription stores information about the event being called when a new message is received.
type Subscription struct {
	Func     func(data MessageData)
	DataType reflect.Type
}

// ReplySubscription stores information about the event being called when a new message is received.
// Used to reply on income messages.
type ReplySubscription struct {
	Func     func(data MessageData) (MessageData, error)
	DataType reflect.Type
}

type MessageData interface {
	protobuf.Message
}

// Broker is the message broker interface.
type Broker interface {
	Subscribe(subjectName string, subscription *Subscription) (*nats.Subscription, error)
	QueueSubscribe(subject string, subscription *Subscription) (err error)
	Request(subj string, data []byte, timeout time.Duration) (*nats.Msg, error)
}

// NatsBroker is the NATS message broker interface.
type NatsBroker interface {
	Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error)
	Publish(subj string, data []byte) error
	QueueSubscribe(subj, queue string, cb nats.MsgHandler) (*nats.Subscription, error)
	Request(subj string, data []byte, timeout time.Duration) (*nats.Msg, error)
}

var pubSub *PubSub

func init() {
	pubSub = &PubSub{
		subscriptions:      make(map[*nats.Subscription]*Subscription),
		replySubscriptions: make(map[*nats.Subscription]*ReplySubscription),
	}
}

func SetBroker(broker NatsBroker) {
	pubSub.nats = broker
}

// Connect allows you to connect to the message broker.
func Connect(clientHandler *dto.ClientHandler, addr string) {
	pubSub.addr = addr
	pubSub.clientHandler = clientHandler
	var err error

	var address = "nats://" + pubSub.addr
	log.Info("Connect to NATS:", address)

	pubSub.nats, err = nats.Connect(
		address,
		nats.PingInterval(2*time.Second),
		nats.MaxPingsOutstanding(5),

		nats.MaxReconnects(-1),
		nats.ReconnectWait(5*time.Second),

		nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
			log.Error("Lost connection to NATS:", err)
		}),
		nats.ReconnectHandler(func(conn *nats.Conn) {
			log.Info("NATS connection restored from URL:", conn.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Fatal("The connection to NATS is closed. Reason:", nc.LastError())
		}))

	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}

	log.Info("Connection to NATS is successful.")

	for subject, handler := range GetHandlers() {
		h := handler
		_, err = pubSub.Subscribe(subject, &Subscription{
			Func: func(data MessageData) {
				defer func() {
					if err := recover(); err != nil {
						const size = 64 << 10
						buf := make([]byte, size)
						buf = buf[:runtime.Stack(buf, false)]
						log.Errorf("Panic: %v\n%s", err, buf)
					}
				}()
				h.Handle(data)
			},
			DataType: h.GetDataType(),
		})
		if err != nil {
			log.Fatal()
		}
	}

	for subject, handler := range GetQueueHandlers() {
		h := handler
		_, err = pubSub.QueueSubscribe(subject, &Subscription{
			Func: func(data MessageData) {
				defer func() {
					if err := recover(); err != nil {
						const size = 64 << 10
						buf := make([]byte, size)
						buf = buf[:runtime.Stack(buf, false)]
						log.Errorf("Panic: %v\n%s", err, buf)
					}
				}()
				h.Handle(data)
			},
			DataType: h.GetDataType(),
		})
		if err != nil {
			log.Fatal()
		}
	}
}

// Subscribe allows you to subscribe to specific subject in the message broker.
// The message will be received by everyone who subscribes to it.
func (b *PubSub) Subscribe(subjectName string, subscription *Subscription) (*nats.Subscription, error) {
	subs, err := b.nats.Subscribe(subjectName, b.messagesHandler)
	if err != nil {
		log.Error("Failed to subscribe to events:", subjectName)
	} else {
		b.mutex.Lock()
		b.subscriptions[subs] = subscription
		b.mutex.Unlock()
	}
	return subs, err
}

// QueueSubscribe allows you to subscribe to specific subject in the message broker.
// Only one subscriber from the group will receive the message.
func (b *PubSub) QueueSubscribe(subjectName string, subscription *Subscription) (*nats.Subscription, error) {
	subs, err := b.nats.QueueSubscribe(subjectName, "GLOBAL", b.messagesHandler)
	if err != nil {
		log.Error("Failed to subscribe to events:", subjectName)
	} else {
		b.mutex.Lock()
		b.subscriptions[subs] = subscription
		b.mutex.Unlock()
	}

	return subs, err
}

// QueueReplySubscribe allows you to subscribe to specific subject in the message broker and reply on income message.
// Only one subscriber from the group will receive the message.
func (b *PubSub) QueueReplySubscribe(subjectName string, subscription *ReplySubscription) (*nats.Subscription, error) {
	subs, err := b.nats.QueueSubscribe(subjectName, "GLOBAL", b.replyMessagesHandler)
	if err != nil {
		log.Error("Failed to subscribe to events:", subjectName)
	} else {
		b.mutex.Lock()
		b.replySubscriptions[subs] = subscription
		b.mutex.Unlock()
	}

	return subs, err
}

func RegisterReplyHandler(subject string, handler UnsubscribableReplyHandler) error {
	subs, err := pubSub.QueueReplySubscribe(subject, &ReplySubscription{
		Func:     handler.Handle,
		DataType: handler.GetDataType(),
	})
	handler.SetSubscription(subs)
	return err
}

func Subscribe(subject string, handler UnsubscribableHandler) error {
	subs, err := pubSub.Subscribe(subject, &Subscription{
		Func: func(data MessageData) {
			handler.Handle(data)
		},
		DataType: handler.GetDataType(),
	})
	handler.SetSubscription(subs)
	return err
}

func QueueSubscribe(subject string, handler UnsubscribableHandler) error {
	subs, err := pubSub.QueueSubscribe(subject, &Subscription{
		Func: func(data MessageData) {
			handler.Handle(data)
		},
		DataType: handler.GetDataType(),
	})
	handler.SetSubscription(subs)
	return err
}

// Send allows you to send a message to the message broker.
func Send(subject string, message MessageData) error {
	data, err := protobuf.Marshal(message)
	if err == nil {
		err = pubSub.nats.Publish(subject, data)
		if err != nil {
			log.Error("Error posting message to NATS:", err)
		}
	} else {
		log.Error("Message serialization error:", err)
	}
	return err
}

func Request[T MessageData](subject string, message MessageData, timeout time.Duration) (res T, err error) {
	var data []byte
	data, err = protobuf.Marshal(message)
	if err == nil {
		msg, err := pubSub.nats.Request(subject, data, timeout)
		if err != nil {
			log.Error("Error posting message to NATS:", err)
			return getZero[T](), err
		}
		resp := new(base.PubSubReply)
		err = protobuf.Unmarshal(msg.Data, resp)
		if err == nil {
			if resp.Error == "" {
				t := reflect.TypeOf(res)
				if t.Kind() == reflect.Pointer {
					t = t.Elem()
				}
				result := reflect.New(t).Interface().(T)
				err = protobuf.Unmarshal(resp.Data, result)
				if err != nil {
					log.Error("Message deserialization error:", err)

					return getZero[T](), err
				}

				return result, nil
			}

			return getZero[T](), errors.New(resp.Error)
		}
	}

	log.Error("Message serialization error:", err)
	return getZero[T](), err
}

func getZero[T any]() T {
	var result T
	return result
}

// SendEvent allows you to send a event to the message broker.
func SendEvent(message MessageData) error {
	return Send(reflect.TypeOf(message).Elem().String(), message)
}

func (b *PubSub) messagesHandler(msg *nats.Msg) {
	b.mutex.Lock()
	subscription, ok := b.subscriptions[msg.Sub]
	b.mutex.Unlock()

	if ok {
		arg := reflect.New(subscription.DataType).Interface()
		if data, ok := arg.(MessageData); ok {
			if err := protobuf.Unmarshal(msg.Data, data); err == nil {
				subscription.Func(data)
			}
		}
	}
}

func (b *PubSub) replyMessagesHandler(msg *nats.Msg) {
	b.mutex.Lock()
	subscription, ok := b.replySubscriptions[msg.Sub]
	b.mutex.Unlock()

	var err error
	if ok {
		arg := reflect.New(subscription.DataType).Interface()
		if data, ok := arg.(MessageData); ok {
			if err = protobuf.Unmarshal(msg.Data, data); err == nil {
				var reply MessageData
				reply, err = subscription.Func(data)
				if err != nil {
					var data []byte
					if data, err = protobuf.Marshal(&base.PubSubReply{Error: err.Error()}); err == nil {
						_ = msg.Respond(data)
						return
					}
				} else {
					var resData []byte
					if resData, err = protobuf.Marshal(reply); err == nil {
						var respondData []byte
						if respondData, err = protobuf.Marshal(&base.PubSubReply{
							Data: resData,
						}); err == nil {
							_ = msg.Respond(respondData)
							return
						}
					}
				}
			}
		}
		if err == nil {
			err = errors.New("PubSub: unable unmarshal respond")
		}
		log.Error(err)
		data, _ := protobuf.Marshal(&base.PubSubReply{
			Error: err.Error(),
		})
		_ = msg.Respond(data)
	}
}
