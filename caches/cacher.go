package caches

type Cacher interface {
	// 缓存配置
	CacheConfig() CacheConfig
}
