package stime

import (
	"context"
	"time"
)

type Clock interface {
	Now(context.Context) time.Time
	RealTime() time.Time
}
