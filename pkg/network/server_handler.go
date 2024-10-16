package network

import "reflect"

// ServerHandler server call handler interface.
type ServerHandler interface {
	Serve(client *Client, packet *Packet, arg interface{})
	GetArgTypeByMethodID(methodID uint16) reflect.Type
	GetMethodNameByID(methodID uint16) string
	Validate(methodID uint16) bool
}
