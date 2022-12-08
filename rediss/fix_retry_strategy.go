package rediss

import "time"

type FixRetryStrategy struct {
	Backoff time.Duration
}

func (f FixRetryStrategy) NextBackoff() time.Duration {
	return f.Backoff
}
