package directive

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
)

// 不需要登录-指令
// 没有操作逻辑，具体操作交给GqlgenMiddileware
func CacheableDirective(ctx context.Context, obj interface{}, next graphql.Resolver, cacheName string, key *string, ttl int64) (res interface{}, err error) {
	return next(ctx)
}

// 不需要登录-指令
// 没有操作逻辑，具体操作交给GqlgenMiddileware
func CacheEvictDirective(ctx context.Context, obj interface{}, next graphql.Resolver, cacheName string, keys []string, allEntries *bool) (res interface{}, err error) {
	return next(ctx)
}
