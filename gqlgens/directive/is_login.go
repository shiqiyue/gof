package directive

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
)

func NeedLogin(isLoginFunc func(ctx context.Context) bool) func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		if isLoginFunc(ctx) {
			return next(ctx)
		}
		return nil, errors.New("请先登录")
	}
}
