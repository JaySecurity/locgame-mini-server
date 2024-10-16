package network

// Error stores information about the error.
type Error struct {
	ErrorCode   uint16
	Description string
}

// GetSize returns the size in bytes.
func (e *Error) GetSize() uint32 {
	return uint32(2 + len([]byte(e.Description)))
}

// GetBytes returns an array of bytes.
func (e *Error) GetBytes() []byte {
	data := make([]byte, e.GetSize())
	data[0] = byte(e.ErrorCode)
	data[1] = byte(e.ErrorCode >> 8)
	copy(data[2:], e.Description)
	return data
}
