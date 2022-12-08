package middle

import (
	"context"
	"github.com/pkg/errors"
	"github.com/shiqiyue/gof/loggers"
	"go.uber.org/zap"
)

func DefaultRecover(ctx context.Context, err interface{}) error {

	loggers.Error(ctx, "panic", zap.Any("err", err), zap.Stack("stack"))

	return errors.New("internal system error")
}
