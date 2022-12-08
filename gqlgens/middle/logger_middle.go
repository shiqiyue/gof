package middle

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/shiqiyue/gof/loggers"
)

func LoggerMiddle(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	response := next(ctx)
	fc := graphql.GetFieldContext(ctx)
	if fc == nil || (fc.Object != "Query" && fc.Object != "Mutation") {
		return response
	}
	operateName := "GQL: " + fc.Object + "_" + fc.Field.Name
	loggers.Info(ctx, operateName)
	errors := graphql.GetErrors(ctx)
	if errors != nil && len(errors) > 0 {
		for _, e := range errors {
			if e != nil {
				// TODO 暂时日志到控制台
				loggers.Error(ctx, fmt.Sprintf("%+v\n", e.Unwrap()))
			}
		}
	}

	return response

}
