package middle

import (
	"context"
	"fmt"
	"github.com/shiqiyue/gof/caches"
	"time"
)

type ApqApqRedisCache struct {
	cache     *caches.Cache
	ttl       time.Duration
	keyPrefix string
}

func NewApqRedisCache(client redis.UniversalClient, ttl time.Duration, keyPrefix string) (*ApqApqRedisCache, error) {
	ctx := context.Background()
	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("could not create cache: %w", err)
	}
	cache := caches.New(&caches.Options{
		Redis:        client,
		LocalCache:   caches.NewTinyLFU(30000, time.Minute*10),
		StatsEnabled: false,
	})
	return &ApqApqRedisCache{ttl: ttl, cache: cache, keyPrefix: keyPrefix}, nil
}

func (c *ApqApqRedisCache) Add(ctx context.Context, key string, value interface{}) {
	_ = c.cache.Once(&caches.Item{
		Ctx:   ctx,
		Key:   c.keyPrefix + key,
		Value: value,
		TTL:   c.ttl,
	})
}

func (c *ApqApqRedisCache) Get(ctx context.Context, key string) (interface{}, bool) {
	var r string
	err := c.cache.Get(ctx, c.keyPrefix+key, &r)

	if err != nil {
		return struct{}{}, false
	}
	return r, true
}
