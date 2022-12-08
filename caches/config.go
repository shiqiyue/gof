package caches

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type CacheConfig struct {
	// key前缀
	KeyPrefix string
	// redis
	Redis redis.UniversalClient
	// 是否启用
	Enable bool
	// id记录缓存-默认TTl
	IdCacheTTL time.Time
	// 列表记录缓存-默认TTL
	ListCacheTTL time.Time
}
