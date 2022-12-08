package redis_zap

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewRedisWriter(key string, cli redis.UniversalClient) *redisWriter {
	return &redisWriter{
		cli: cli, listKey: key,
	}
}

// 为 logger 提供写入 redis 队列的 io 接口
type redisWriter struct {
	cli     redis.UniversalClient
	listKey string
}

func (w *redisWriter) Write(p []byte) (int, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// 将日志写入 redis 队列
	_, err := w.cli.RPush(context.Background(), w.listKey, p).Result()
	return len(p), err
}

func NewLogger(writer *redisWriter) *zap.Logger {
	// 限制日志输出级别, >= DebugLevel 会打印所有级别的日志
	// 生产环境中一般使用 >= ErrorLevel
	lowPriority := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv >= zapcore.DebugLevel
	})

	// 使用 JSON 格式日志
	jsonEnc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	stdCore := zapcore.NewCore(jsonEnc, zapcore.Lock(os.Stdout), lowPriority)

	// addSync 将 io.Writer 装饰为 WriteSyncer
	// 故只需要一个实现 io.Writer 接口的对象即可
	syncer := zapcore.AddSync(writer)
	redisCore := zapcore.NewCore(jsonEnc, syncer, lowPriority)

	// 集成多个 core
	core := zapcore.NewTee(stdCore, redisCore)

	// logger 输出到 console 且标识调用代码行
	return zap.New(core).WithOptions(zap.AddCaller())
}
