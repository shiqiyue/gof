package middle

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/go-redis/redis/v9"
	"github.com/shiqiyue/gof/caches"
	"github.com/shiqiyue/gof/loggers"
	"github.com/shiqiyue/gof/reflects"
	"github.com/vektah/gqlparser/v2/ast"
	"go.uber.org/zap"
	"io"
	"reflect"
	"strconv"
	"sync"
	"time"
)

var (
	// cacheable 指令
	CACHE_ABLE_DIRECTIVE = "cacheable"
	// cacheable 指令-参数 缓存名称
	CACHE_ABLE_DIRECTIVE_ARG_CACHENAME = "cacheName"
	// cacheable 指令-参数 缓存名称
	CACHE_ABLE_DIRECTIVE_ARG_KEY = "key"
	// cacheable 指令-参数 过期时间
	CACHE_ABLE_DIRECTIVE_ARG_TTL = "ttl"

	// cacheEvict 指令
	CACHE_EVICT_DIRECTIVE = "cacheEvict"
)

// 缓存中间件
type (
	CacheMiddleware struct {
		client  redis.UniversalClient
		cache   *caches.Cache
		typeMap *sync.Map
	}
)

var _ interface {
	graphql.HandlerExtension
	graphql.OperationInterceptor
	graphql.FieldInterceptor
} = CacheMiddleware{}

func NewCacheMiddleware(client redis.UniversalClient) (*CacheMiddleware, error) {
	ctx := context.Background()
	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("could not create cache: %w", err)
	}
	cache := caches.New(&caches.Options{
		Redis:        client,
		StatsEnabled: false,
		Marshal: func(i interface{}) ([]byte, error) {
			return json.Marshal(i)
		},
		Unmarshal: func(bytes []byte, i interface{}) error {
			return json.Unmarshal(bytes, i)
		},
	})
	return &CacheMiddleware{
		client:  client,
		cache:   cache,
		typeMap: &sync.Map{},
	}, nil
}

func (a CacheMiddleware) ExtensionName() string {
	return "Cache"
}

func (a CacheMiddleware) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (a CacheMiddleware) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	/*	operationContext := graphql.GetOperationContext(ctx)
		if a.IntrospectionEnable && operationContext.OperationName == "IntrospectionQuery" {
			newCtx := context.WithValue(ctx, notNeedAuth, true)
			return next(newCtx)
		}*/
	return next(ctx)
}

func (a CacheMiddleware) InterceptField(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc.Object != "Query" && fc.Object != "Mutation" {
		return next(ctx)
	}

	directives := fc.Field.Field.Definition.Directives
	if directives != nil && len(directives) > 0 {
		// 缓存指令
		cacheAbleDirective := directives.ForName(CACHE_ABLE_DIRECTIVE)
		if cacheAbleDirective != nil {
			// 缓存名称
			cacheName := a.getCacheName(ctx, cacheAbleDirective)
			// key
			key, err := a.getKey(ctx, cacheAbleDirective, fc)
			if err != nil {
				return nil, err
			}
			// 操作名称
			operName := a.getOperName(ctx, fc)
			realKey := fmt.Sprintf("%s:%s", cacheName, key)
			// ttl
			ttl, err := a.getTtl(ctx, cacheAbleDirective)
			if err != nil {
				loggers.Error(ctx, "解析缓存ttl异常", zap.Error(err))
				return nil, err
			}
			// 缓存保存的类型信息
			resultType, resultTypeLoadOk := a.typeMap.Load(operName)
			if resultTypeLoadOk {
				// 如果类型信息，存在的话，则尝试获取缓存结果
				fc.Result = reflects.NewByType(resultType.(reflect.Type))
				err = a.cache.Get(ctx, realKey, fc.Result)
				if err == nil {
					return fc.Result, nil
				}
			}
			res, err := next(ctx)
			if !resultTypeLoadOk {
				// 如果类型信息不存在的话，则将类型信息存储起来
				if err == nil && res != nil {
					a.typeMap.Store(operName, reflect.TypeOf(res))
				}
			}
			if err == nil && res != nil {
				// 设置结果到缓存中
				_ = a.cache.Set(&caches.Item{
					Ctx:   ctx,
					Key:   realKey,
					Value: res,
					TTL:   time.Duration(ttl) * time.Millisecond,
				})
			}
			return res, err
		}
	}

	return next(ctx)
}

func (a CacheMiddleware) getCacheName(ctx context.Context, cacheAbleDirective *ast.Directive) string {
	return cacheAbleDirective.Arguments.ForName(CACHE_ABLE_DIRECTIVE_ARG_CACHENAME).Value.String()
}

func (a CacheMiddleware) getOperName(ctx context.Context, fc *graphql.FieldContext) string {
	return fc.Object + "_" + fc.Field.Name
}

func (a CacheMiddleware) getTtl(ctx context.Context, cacheAbleDirective *ast.Directive) (int64, error) {
	ttl, err := strconv.ParseInt(cacheAbleDirective.Arguments.ForName(CACHE_ABLE_DIRECTIVE_ARG_TTL).Value.String(), 10, 64)
	if err != nil {
		return 0, err
	}
	return ttl, nil
}

func (a CacheMiddleware) getKey(ctx context.Context, cacheAbleDirective *ast.Directive, fc *graphql.FieldContext) (string, error) {
	keyArg := cacheAbleDirective.Arguments.ForName(CACHE_ABLE_DIRECTIVE_ARG_KEY)
	if keyArg != nil {
		return keyArg.Value.String(), nil
	} else {
		// key为空，则自动生成key
		key := a.getOperName(ctx, fc)
		operateArgs := fc.Args
		if operateArgs != nil {
			operateArgJsonBs, err := json.Marshal(operateArgs)
			if err != nil {
				return "", err
			}
			hash := sha512.New()
			r := bytes.NewBuffer(operateArgJsonBs)
			_, err = io.Copy(hash, r)
			if err != nil {
				return "", err
			}
			hashData := hash.Sum(nil)
			hashValue := hex.EncodeToString(hashData)
			key = key + ":" + hashValue
		}
		return key, nil

	}
}
