package network

import (
	"context"
	"reflect"
	"time"

	"github.com/gobwas/ws"
	protobuf "google.golang.org/protobuf/proto"
	"locgame-mini-server/pkg/log"
)

// Client represents a network client.
type Client struct {
	ctx    context.Context
	Stream *Stream

	OnCloseConnection func(*Client)

	DeferredCalls Queue

	isBusy *AtomicBool
}

// ClientContextKey points to the client's context key.
type ClientContextKey struct{}
type ClientGeoRegionKey struct{}

// NewClient allows you to create a new instance of the network client.
func NewClient(stream *Stream, geoRegion string) *Client {
	c := &Client{
		Stream: stream,
		isBusy: new(AtomicBool),
	}
	ctx, cancel := context.WithCancel(context.WithValue(
		context.WithValue(context.Background(), ClientContextKey{}, c),
		ClientGeoRegionKey{},
		geoRegion))

	c.ctx = ctx

	stream.Init(func() {
		if c.OnCloseConnection != nil {
			c.OnCloseConnection(c)
		}
		cancel()
	})

	return c
}

// Context returns the client context.
func (c *Client) Context() context.Context {
	return c.ctx
}

// Defer allows to make a deferred call that will occur after sending data to the client, if there was a call during any RPC.
// For example:
// RPC Request ->
// <- RPC Response
// <- Deferred Call
//
// Accepts a context as an argument to determine if an RPC belongs to a specific client.
// If the deferred call occurs within the processing of the current RPC,
// then the deferred call will be called after the RPC, otherwise the call will occur immediately.
func (c *Client) Defer(ctx context.Context, fn func()) {
	if ctx != c.ctx {
		fn()
		return
	}

	if c.isBusy.IsSet() {
		c.DeferredCalls.Enqueue(fn)
	} else {
		fn()
	}
}

// Handle to handle the client's network stream.
func (c *Client) Handle(handler ServerHandler) error {
	ticker := time.NewTicker(20 * time.Second)

	defer func() {
		ticker.Stop()
		c.Stream.Close()
	}()

	go func() {
		for {
			select {
			case <-ticker.C:
				_, _ = c.Stream.Write(ws.CompiledPing)
			case <-c.ctx.Done():
				return
			}
		}
	}()

	for {
		packet, err := c.Stream.ReadPacket()
		if err != nil {
			return err
		}
		protoType := handler.GetArgTypeByMethodID(packet.MethodID)
		if protoType != nil {
			arg := reflect.New(protoType).Interface()
			err := protobuf.Unmarshal(packet.Payload, arg.(protobuf.Message))
			if err == nil {
				if Verbose {
					log.Debugf("--> %s (%s) (%v)", handler.GetMethodNameByID(packet.MethodID), arg, c.Stream.Conn.RemoteAddr())
				}
				c.isBusy.Set()
				handler.Serve(c, packet, arg)
				c.isBusy.UnSet()
			} else {
				log.Error("Payload read error:", err)
			}
		} else {
			log.Error("Error finding argument type for method:", packet.MethodID)
		}
	}
}

// Shutdown allows you to close the client connection.
func (c *Client) Shutdown() {
	_, _ = c.Stream.Conn.Write(ws.CompiledClose)
	c.Stream.Close()
}
