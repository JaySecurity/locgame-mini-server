package network

import (
	"fmt"
)

// Packet stores information about the network packet.
type Packet struct {
	MethodID uint16
	Seq      uint32
	Payload  []byte
	Error    *Error
}

// CalculateSize returns the calculated network packet size.
func (p *Packet) CalculateSize() uint32 {
	var length uint32
	length += 2 /* MethodID (2 bytes) */
	length += 4 /* Seq (4 bytes) */
	length += 4 /* PayloadSize (4 bytes) */
	length += 2 /* ErrorSize (2 bytes) */
	length += uint32(len(p.Payload))
	if p.Error != nil {
		length += p.Error.GetSize()
	}
	return length
}

// GetBytes returns an array of network packet bytes.
func (p *Packet) GetBytes() []byte {
	data := make([]byte, p.CalculateSize())

	data[0] = byte(p.MethodID)
	data[1] = byte(p.MethodID >> 8)

	data[2] = byte(p.Seq)
	data[3] = byte(p.Seq >> 8)
	data[4] = byte(p.Seq >> 16)
	data[5] = byte(p.Seq >> 24)

	payloadSize := len(p.Payload)
	data[6] = byte(payloadSize)
	data[7] = byte(payloadSize >> 8)
	data[8] = byte(payloadSize >> 16)
	data[9] = byte(payloadSize >> 24)

	if p.Error == nil {
		data[10] = 0
		data[11] = 0
	} else {
		errorSize := p.Error.GetSize()
		data[10] = byte(errorSize)
		data[11] = byte(errorSize >> 8)

		copy(data[12+payloadSize:], p.Error.GetBytes())
	}

	copy(data[12:], p.Payload)

	return data
}

// String represents the network packet as text.
func (p *Packet) String() string {
	return fmt.Sprintf("Type: %d, Seq: %d, Length: %d", p.MethodID, p.Seq, len(p.Payload))
}
