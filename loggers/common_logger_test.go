package loggers

import (
	"context"
	"go.uber.org/zap"
	"testing"
)

func TestInfo(t *testing.T) {
	SetAppName("t")
	Info(context.Background(), "你好", zap.String("testString", "value1"))
}
