package network

import (
	"encoding/binary"
	"errors"
	"net"

	"github.com/gobwas/ws/wsutil"
	"locgame-mini-server/pkg/log"
)

var (
	// ErrWrongPacketType indicates an invalid network packet type.
	ErrWrongPacketType = errors.New("WRONG_PACKET_TYPE")

	// ErrPacketSizeExceeded indicates that the network packet size has been exceeded.
	ErrPacketSizeExceeded = errors.New("PACKET_SIZE_EXCEEDED")

	// ErrHeaderSizeExceeded indicates no packet header size.
	ErrHeaderSizeExceeded = errors.New("HEADER_SIZE_EXCEEDED")
)

// Stream is a websocket stream.
type Stream struct {
	net.Conn

	isClosed          bool
	OnCloseConnection func()
	ValidateFunc      func(methodID uint16) bool
}

// Init allows to initialize the network stream.
func (s *Stream) Init(onCloseConnection func()) {
	s.OnCloseConnection = onCloseConnection
}

// ReadPacket allows to read a network packet from a network stream.
func (s *Stream) ReadPacket() (*Packet, error) {
	data, err := wsutil.ReadClientBinary(s.Conn)
	if err != nil {
		return nil, err
	}

	if len(data) < 10 {
		return nil, ErrHeaderSizeExceeded
	}

	methodID := binary.LittleEndian.Uint16(data)
	seq := binary.LittleEndian.Uint32(data[2:])
	size := binary.LittleEndian.Uint32(data[6:])

	if s.ValidateFunc != nil && !s.ValidateFunc(methodID) {
		return nil, ErrWrongPacketType
	}

	if len(data) < int(10+size) {
		log.Error("Packet read error: The number of bytes read does not matchmaking the packet size.")
		return nil, ErrPacketSizeExceeded
	}

	return &Packet{
		MethodID: methodID,
		Seq:      seq,
		Payload:  data[10 : 10+size],
	}, nil
}

// WritePacket allows to write a network packet to the network stream.
func (s *Stream) WritePacket(packet *Packet) error {
	return wsutil.WriteServerBinary(s.Conn, packet.GetBytes())
}

// Close allows to close the network stream.
func (s *Stream) Close() {
	if s.isClosed {
		return
	}
	s.isClosed = true
	_ = s.Conn.Close()

	if s.OnCloseConnection != nil {
		s.OnCloseConnection()
	}
}
