package gorms

import (
	"context"
	"github.com/shiqiyue/gof/loggers"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type gormLogger struct {
	SlowThreshold time.Duration
	// 白名单单词,用于打印执行sql的文件名和行数
	whiteWords []string
}

func NewGormLogger(slowthreshold time.Duration, whiteWords []string) *gormLogger {

	if whiteWords == nil {
		whiteWords = make([]string, 0)
	}

	return &gormLogger{
		SlowThreshold: slowthreshold,
		whiteWords:    whiteWords,
	}
}

func (g gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g gormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	loggers.Info(ctx, g.warpMsg(s), zap.Any("infos", i))

}

func (g gormLogger) Warn(ctx context.Context, s string, i ...interface{}) {

	loggers.Warn(ctx, g.warpMsg(s), zap.Any("infos", i))

}

func (g gormLogger) Error(ctx context.Context, s string, i ...interface{}) {

	loggers.Error(ctx, g.warpMsg(s), zap.Any("infos", i))

}

func (g gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	elapsed := time.Since(begin)
	sql, rows := fc()
	if rows == -1 {
		loggers.Debug(ctx, g.warpMsg(sql), zap.String("times", elapsed.String()), zap.String("rows", "-"), zap.Error(err))
	} else {
		loggers.Debug(ctx, g.warpMsg(sql), zap.String("times", elapsed.String()), zap.Int64("rows", rows), zap.Error(err))
	}
}

// 获取打印日志当前应用的文件名和行号
func (g gormLogger) FileWithLineNum() string {
	if g.whiteWords == nil || len(g.whiteWords) == 0 {
		return ""
	}
	for i := 3; i < 15; i++ {
		_, filePath, line, ok := runtime.Caller(i)
		if !ok {
			continue
		}
		for _, whiteWord := range g.whiteWords {
			if strings.Index(filePath, whiteWord) >= 0 {
				base := filepath.Base(filePath)
				return whiteWord + base + ":" + strconv.FormatInt(int64(line), 10)
			}
		}
	}
	return ""
}

func (g gormLogger) warpMsg(msg string) string {
	return g.FileWithLineNum() + "  " + msg
}
