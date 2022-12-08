package extension

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/shiqiyue/gof/loggers"
	"go.uber.org/zap"
)

type (
	Logger struct {
		IsPrintLog bool
	}
)

var _ interface {
	graphql.HandlerExtension
	graphql.OperationInterceptor
	graphql.FieldInterceptor
	graphql.ResponseInterceptor
} = Logger{}

func (a Logger) ExtensionName() string {
	return "Logger"
}

func (a Logger) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (a Logger) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {

	return next(ctx)
}

func (a Logger) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	response := next(ctx)
	if response != nil {
		if a.IsPrintLog {
			loggers.Info(ctx, "GQL返回结果", zap.Any("res", response.Data))
		}
	}
	errList := graphql.GetErrors(ctx)
	if len(errList) != 0 {
		for _, err := range errList {
			loggers.Error(ctx, err.Error(), zap.Error(err.Unwrap()))
		}
	}
	return response
}

func (a Logger) InterceptField(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc.Object != "Query" && fc.Object != "Mutation" {
		return next(ctx)
	}
	operateName := "GQL: " + fc.Object + " " + fc.Field.Name
	if a.IsPrintLog {
		loggers.Info(ctx, "GQL操作请求参数", zap.Any("args", fc.Args))
	}
	res, err := next(ctx)
	if err != nil {
		loggers.Error(ctx, fmt.Sprintf("%s-异常", operateName), zap.Error(err))
	}
	if a.IsPrintLog {
		loggers.Info(ctx, "GQL操作返回结果", zap.Any("res", res))
	}

	return res, err
}
