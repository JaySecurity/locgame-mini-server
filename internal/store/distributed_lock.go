package store

import (
	"time"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type DistributedLocks struct {
	rs *redsync.Redsync
}

// NewDistributedLocks creates a new instance of the distributed locks store.
func NewDistributedLocks(client *goredislib.Client) *DistributedLocks {
	return &DistributedLocks{
		rs: redsync.New(goredis.NewPool(client)),
	}
}

func (l *DistributedLocks) NewLock(lockID string, ttl time.Duration) *redsync.Mutex {
	return l.rs.NewMutex("locg.v1.locks:"+lockID, redsync.WithExpiry(ttl), redsync.WithTries(1))
}
