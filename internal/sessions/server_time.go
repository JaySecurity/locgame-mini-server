package sessions

import (
	"context"
	"time"
)

type ServerTime struct {
	Session *SessionStore
}

func (s *ServerTime) Now(_ context.Context) time.Time {
	return time.Now().UTC()
}

func (s *ServerTime) RealTime() time.Time {
	return time.Now().UTC()
}
