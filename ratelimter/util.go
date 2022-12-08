package ratelimter

import (
	"github.com/juju/ratelimit"
	"time"
)

func NewEmptyBucket(fillTime time.Duration, capacity int64) *ratelimit.Bucket {
	bucket := ratelimit.NewBucket(fillTime/time.Duration(capacity), capacity)
	bucket.TakeAvailable(capacity)
	return bucket
}

func NewFullBucket(fillTime time.Duration, capacity int64) *ratelimit.Bucket {
	bucket := ratelimit.NewBucket(fillTime/time.Duration(capacity), capacity)
	return bucket
}
