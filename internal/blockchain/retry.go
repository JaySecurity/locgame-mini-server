package blockchain

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"locgame-mini-server/pkg/log"
)

// Retry invoke a function with exponential backoff when a retryable error is encountered.
// If fn returns an err that is wrapped in common.ErrRetyable, fn will be retried with exponential backoff.
// This was originally implemented to solve getting rate limited by the alchemy APIs
//
// See: https://docs.alchemy.com/alchemy/documentation/rate-limits#retries
func Retry[T any](ctx context.Context, fn func() (T, error)) (result T, err error) {
	var tries float64 = 0
	for {
		select {
		case <-ctx.Done():
			return result, fmt.Errorf("exceeded context deadline: %w", err)
		default:
			if tries >= 5 {
				return result, fmt.Errorf("exceeded max retries: %w", err)
			}

			result, err = fn()

			isTooManyRequests := false
			if httpErr, ok := err.(rpc.HTTPError); ok {
				if httpErr.StatusCode == 429 {
					isTooManyRequests = true
					log.Warning("Too many request. Retry RPC request.")
				}
			}
			if err == nil || !isTooManyRequests {
				return result, err
			}

			// sleep a power of two seconds + a random number of seconds between 0 and 1
			sleep := (2 * time.Second * time.Duration(math.Pow(2, tries))) + time.Duration(rand.Float64()*float64(time.Second))

			tries++

			time.Sleep(sleep)
		}
	}
}
