package directive

import (
	"context"
	"errors"
	"gitee.com/shiqiyue/go-admin/internal/pkg/auths"
	"github.com/99designs/gqlgen/graphql"
)

// 需要登录
func NeedLogin(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	userId := auths.GetUserId(ctx)
	if userId == nil {
		return nil, errors.New("请先登录")
	}
	return next(ctx)
}
