package rediss

import (
	"context"
	"fmt"
	"gitee.com/shiqiyue/redislock"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewLocker(t *testing.T) {
	client := redis.NewClient(&redis.Options{Network: "tcp", Addr: "127.0.0.1:6379", DB: 3})
	defer client.Close()
	locker, err := NewLocker("sys", client, nil)
	assert.Nil(t, err)
	bgCtx := context.Background()
	_, err = locker.GetLock(bgCtx, "enterprise:create:1", &redislock.Options{Wait: time.Second * 30, KeyTTL: time.Second * 30})
	assert.Nil(t, err)
	_, err = locker.GetLock(bgCtx, "enterprise:create:1", &redislock.Options{Wait: time.Second * 30, KeyTTL: time.Second * 30})
	assert.NotNil(t, err)
	fmt.Println(err)
	//lock.Release(bgCtx)
}
