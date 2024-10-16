package pubsub

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"

	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	protobuf "google.golang.org/protobuf/proto"
	"locgame-mini-server/pkg/dto"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/network"
)

var playerChannelsMutex sync.Mutex
var playerChannels = make(map[string]*playerSubscription)
var playerHandlers = make(map[string]PlayerHandler)

type PlayerHandler interface {
	Handle(client *network.Client, clientHandler *dto.ClientHandler, data MessageData)
	GetDataType() reflect.Type
}

type playerSubscription struct {
	Client       *network.Client
	Subscription *nats.Subscription
}

func RegisterPlayerChannel(id primitive.ObjectID, client *network.Client) {
	hex := id.Hex()

	playerChannelsMutex.Lock()
	defer playerChannelsMutex.Unlock()

	if subscription, ok := playerChannels[hex]; ok {
		_ = subscription.Subscription.Unsubscribe()
	}

	subscription, err := pubSub.nats.Subscribe(fmt.Sprintf("player.%s.*", hex), playerMessageHandler)
	if err != nil {
		log.Error("Failed to create a subscription to the player's channel:", err)
		return
	}
	playerChannels[hex] = &playerSubscription{
		Client:       client,
		Subscription: subscription,
	}
}

func UnregisterPlayerChannel(id primitive.ObjectID, client *network.Client) {
	hex := id.Hex()
	playerChannelsMutex.Lock()

	found := false
	if channel, ok := playerChannels[hex]; ok {
		if channel.Client == client {
			found = true
			_ = channel.Subscription.Unsubscribe()
		}
	}

	if found {
		delete(playerChannels, hex)
	}
	playerChannelsMutex.Unlock()
}

func RegisterPlayerHandler(handler PlayerHandler) {
	playerChannelsMutex.Lock()
	playerHandlers[handler.GetDataType().Name()] = handler
	playerChannelsMutex.Unlock()
}

func SendToPlayer(accountID string, data MessageData) error {
	return Send(fmt.Sprintf("player.%s.%s", accountID, getMessageDataName(data)), data)
}

func getMessageDataName(data MessageData) string {
	return reflect.TypeOf(data).Elem().Name()
}

func playerMessageHandler(msg *nats.Msg) {
	accountID, subSubject, err := parseSubject(msg.Subject)
	if err == nil {
		playerChannelsMutex.Lock()

		if subscription, ok := playerChannels[accountID]; ok {
			if handler, ok := playerHandlers[subSubject]; ok {
				playerChannelsMutex.Unlock()

				arg := reflect.New(handler.GetDataType()).Interface()
				if err := protobuf.Unmarshal(msg.Data, arg.(MessageData)); err == nil {
					defer func() {
						if err := recover(); err != nil {
							const size = 64 << 10
							buf := make([]byte, size)
							buf = buf[:runtime.Stack(buf, false)]
							log.Errorf("Panic: %v\n%s", err, buf)
						}
					}()
					handler.Handle(subscription.Client, pubSub.clientHandler, arg.(protobuf.Message))
				}
			} else {
				playerChannelsMutex.Unlock()
			}
		} else {
			playerChannelsMutex.Unlock()
		}
	}
}

func parseSubject(subject string) (accountID string, subSubject string, err error) {
	splits := strings.Split(subject, ".")
	if len(splits) == 3 {
		accountID = splits[1]
		subSubject = splits[2]
	} else {
		err = errors.New("unknown subject")
	}
	return
}
