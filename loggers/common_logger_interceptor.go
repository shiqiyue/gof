package loggers

import (
	"context"
	"go.uber.org/zap/zapcore"
)

type FieldInterceptor func(ctx context.Context, fields []zapcore.Field) []zapcore.Field

var fieldInterceptors []FieldInterceptor = make([]FieldInterceptor, 0)

func ClearFieldInterceptor() {
	fieldInterceptors = make([]FieldInterceptor, 0)
}

func AddFieldIntercetpor(interceptor FieldInterceptor) {
	fieldInterceptors = append(fieldInterceptors, interceptor)
}
