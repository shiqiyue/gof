package rediss

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"strconv"

	"github.com/go-redis/redis/extra/rediscmd/v8"
	"github.com/go-redis/redis/v8"
)

type TracingHook struct{}

var _ redis.Hook = (*TracingHook)(nil)

func NewTracingHook() *TracingHook {
	return new(TracingHook)
}

func (TracingHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {

	serverSpan, newCtx := opentracing.StartSpanFromContext(ctx, cmd.FullName())

	serverSpan.SetTag("db.system", "redis")
	serverSpan.SetTag("db.statement", rediscmd.CmdString(cmd))

	return newCtx, nil
}

func (TracingHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := opentracing.SpanFromContext(ctx)
	defer span.Finish()
	if err := cmd.Err(); err != nil {
		if err == redis.Nil {
			return nil
		}
		ext.Error.Set(span, true)
		span.LogFields(
			log.String("error.message", fmt.Sprintf("%+v\n", err)),
			log.String("error.kind", fmt.Sprintf("%T", err)),
		)
	}
	return nil
}

func (TracingHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {

	summary, cmdsString := rediscmd.CmdsString(cmds)

	serverSpan, newCtx := opentracing.StartSpanFromContext(ctx, "pipeline "+summary)

	serverSpan.SetTag("db.system", "redis")
	serverSpan.SetTag("db.redis.num_cmd", len(cmds))

	serverSpan.SetTag("db.statement", cmdsString)

	return newCtx, nil

}

func (TracingHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span := opentracing.SpanFromContext(ctx)
	defer span.Finish()
	for i, cmd := range cmds {
		if err := cmd.Err(); err != nil {
			if err == redis.Nil {
				continue
			}
			ext.Error.Set(span, true)
			span.LogFields(
				log.String("error.message."+strconv.Itoa(i), fmt.Sprintf("%+v\n", err)),
				log.String("error.kind."+strconv.Itoa(i), fmt.Sprintf("%T", err)),
			)
		}
	}
	return nil
}
