package stime

import (
	"context"
	"time"

	"locgame-mini-server/pkg/log"
)

var clock Clock

func Init(c Clock) {
	clock = c
}

func Now(ctx context.Context) time.Time {
	return clock.Now(ctx)
}

func RealTime() time.Time {
	if clock == nil {
		log.Warning("Clock instance not found. Fallback to default time")
		return time.Now().UTC()
	}

	return clock.RealTime()
}
