package loggers

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 简单日志实例

var l *iLogger

var containerId string

var loggerKey struct{}

var opKey struct{}

func getLogger(ctx context.Context) *iLogger {
	return l
}

func init() {
	l = &iLogger{ls: map[string]*zap.Logger{"stdout": NewLogger(zapcore.InfoLevel, nil, ENCODER_CONSOLE)}}

}

// 注册实例到简单日志
func RegisterLogger(name string, logger *zap.Logger) {
	l.ls[name] = logger
}

func UnRegisterLoggers() {
	l.ls = make(map[string]*zap.Logger, 0)
}

func AppendFields(ctx context.Context, fields ...zapcore.Field) []zapcore.Field {
	for _, interceptor := range fieldInterceptors {
		fields = interceptor(ctx, fields)
	}
	return fields

}

func Info(ctx context.Context, msg string, fields ...zapcore.Field) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getLogger(ctx).Info(msg, AppendFields(ctx, fields...)...)
}

func InfoF(ctx context.Context, msg string, a ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	Info(ctx, fmt.Sprintf(msg, a...))
}

func Debug(ctx context.Context, msg string, fields ...zapcore.Field) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getLogger(ctx).Debug(msg, AppendFields(ctx, fields...)...)
}

func DebugF(ctx context.Context, msg string, a ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	Debug(ctx, fmt.Sprintf(msg, a...))
}

func Warn(ctx context.Context, msg string, fields ...zapcore.Field) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getLogger(ctx).Warn(msg, AppendFields(ctx, fields...)...)
}

func WarnF(ctx context.Context, msg string, a ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	Warn(ctx, fmt.Sprintf(msg, a...))
}

func Error(ctx context.Context, msg string, fields ...zapcore.Field) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getLogger(ctx).Error(msg, AppendFields(ctx, fields...)...)
}

func ErrorF(ctx context.Context, msg string, a ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	Error(ctx, fmt.Sprintf(msg, a...))
}

func Fatal(ctx context.Context, msg string, fields ...zapcore.Field) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getLogger(ctx).Fatal(msg, AppendFields(ctx, fields...)...)
}

func FatalF(ctx context.Context, msg string, a ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	Fatal(ctx, fmt.Sprintf(msg, a...))
}
