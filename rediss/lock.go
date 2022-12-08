package rediss

import (
	"context"
	"errors"
	"gitee.com/shiqiyue/redislock"
	"github.com/go-redis/redis/v8"
	"time"
)

// redis分布式锁
type Locker struct {
	client         redis.UniversalClient
	LockerClient   *redislock.Client
	defaultOptions *redislock.Options
	moduleName     string
}

// NewLocker new a redis locker
// moduleName: 模块名称
// client: 客户端
// defaultOptions: 默认配置
func NewLocker(moduleName string, client redis.UniversalClient, defaultOptions *redislock.Options) (*Locker, error) {
	if client == nil {
		return nil, errors.New("redis客户端不能为空")
	}
	l := redislock.New(client)
	return &Locker{
		client:         client,
		LockerClient:   l,
		defaultOptions: defaultOptions,
		moduleName:     moduleName,
	}, nil
}

// 获取锁对象
func (l *Locker) GetLock(ctx context.Context, key string, ttl time.Duration, options *redislock.Options) (*redislock.Lock, error) {
	op := l.defaultOptions
	if options != nil {
		op = options
	}
	return l.LockerClient.Obtain(ctx, l.getKey(key), ttl, op)
}

func (l *Locker) getKey(key string) string {
	return l.moduleName + ":" + key
}
