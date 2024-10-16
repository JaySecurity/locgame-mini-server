package middleware

import (
	"reflect"

	"go.uber.org/ratelimit"
	"locgame-mini-server/pkg/network"
)

type RateLimitMiddleware struct {
	limiter ratelimit.Limiter
	handler network.ServerHandler
}

func NewRateLimitMiddleware(handler network.ServerHandler, rate int) *RateLimitMiddleware {
	return &RateLimitMiddleware{handler: handler, limiter: ratelimit.New(rate)}
}

func (h *RateLimitMiddleware) Serve(client *network.Client, packet *network.Packet, arg interface{}) {
	h.limiter.Take()
	h.handler.Serve(client, packet, arg)
}

func (h *RateLimitMiddleware) GetArgTypeByMethodID(methodID uint16) reflect.Type {
	return h.handler.GetArgTypeByMethodID(methodID)
}

func (h *RateLimitMiddleware) GetMethodNameByID(methodID uint16) string {
	return h.handler.GetMethodNameByID(methodID)
}

func (h *RateLimitMiddleware) Validate(methodID uint16) bool {
	return h.handler.Validate(methodID)
}
