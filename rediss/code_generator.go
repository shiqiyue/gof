package rediss

import (
	"context"
	errors2 "errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

// 编号生成规则
const (
	/***
	 * 日期+redis自增
	 */
	DATE_AND_REDIS_INC = iota
)

const (
	DEFAULT_REDIS_KEY_PREFIX = "codeGenerate:"
)

type RedisCodeGenerator struct {
	r redis.UniversalClient

	// 编号前缀
	noPrefix string

	// 编号名称
	noName string

	// redis key前缀
	redisKeyPrefix string

	// 编号生成规则
	genType int
}

func NewRedisCodeGenerator(r redis.UniversalClient, noPrefix, noName, redisKeyPrefix string, genType int) (*RedisCodeGenerator, error) {
	if noName == "" {
		return nil, errors2.New("noName can not be empty")
	}
	if redisKeyPrefix == "" {
		redisKeyPrefix = DEFAULT_REDIS_KEY_PREFIX
	}
	if genType != DATE_AND_REDIS_INC {
		return nil, errors2.New("genType is not valid")
	}
	return &RedisCodeGenerator{
		r:              r,
		noPrefix:       noPrefix,
		noName:         noName,
		redisKeyPrefix: redisKeyPrefix,
		genType:        genType,
	}, nil
}

func (r *RedisCodeGenerator) Gen(ctx context.Context) (string, error) {
	switch r.genType {
	case DATE_AND_REDIS_INC:
		return r.genByDateAndRedisInc(ctx)
	default:
		panic(fmt.Sprintf("不支持的编号生成规则:%d", r.genType))
	}
}

func (r *RedisCodeGenerator) genByDateAndRedisInc(ctx context.Context) (string, error) {
	now := time.Now()
	dateStr := now.Format("20060102")
	redisKey := r.redisKeyPrefix + r.noName + ":" + dateStr
	incr := r.r.Incr(ctx, redisKey)
	if incr.Err() != nil {
		return "", incr.Err()
	}
	r.r.Expire(ctx, redisKey, time.Hour*24)
	val := incr.Val()
	code := fmt.Sprintf("%s%s%04d", r.noPrefix, dateStr, val)
	return code, nil

}
