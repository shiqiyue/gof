package extension

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

type (
	GqlgenTracer struct{}
)

var _ interface {
	graphql.HandlerExtension
	graphql.OperationInterceptor
	graphql.FieldInterceptor
} = GqlgenTracer{}

func (a GqlgenTracer) ExtensionName() string {
	return "OpenTracing"
}

func (a GqlgenTracer) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (a GqlgenTracer) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	return next(ctx)
}

func (a GqlgenTracer) InterceptField(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc.Object != "Query" && fc.Object != "Mutation" {
		return next(ctx)
	}
	operateName := "GQL: " + fc.Object + "_" + fc.Field.Name
	span, ctx := opentracing.StartSpanFromContext(ctx, operateName)
	span = span.SetTag("resolver.object", fc.Object)
	span = span.SetTag("resolver.field", fc.Field.Name)
	span = span.SetBaggageItem("op", operateName)
	span.LogKV()
	defer span.Finish()

	res, err := next(ctx)
	if err != nil {
		ext.Error.Set(span, true)
		span.LogFields(
			log.String("event", "error"),
		)
		span.LogFields(
			log.String("error.message", fmt.Sprintf("%+v\n", err)),
			log.String("error.kind", fmt.Sprintf("%T", err)),
		)
	}

	errList := graphql.GetFieldErrors(ctx, fc)
	if len(errList) != 0 {
		ext.Error.Set(span, true)
		span.LogFields(
			log.String("event", "error"),
		)

		for idx, err := range errList {
			span.LogFields(
				log.String(fmt.Sprintf("error.%d.message", idx), fmt.Sprintf("%+v\n", err.Unwrap())),
				log.String(fmt.Sprintf("error.%d.kind", idx), fmt.Sprintf("%T", err)),
			)
		}
	}

	return res, err
}
