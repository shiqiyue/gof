package loggers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

// 配置选项
type ConfigOption struct {
	// 日志保存路径
	Path string
	// 日志等级
	Level string
	// 日志模式
	Mode string
	// 应用名称
	AppName string
}

type Encoder int

var (
	ENCODER_JSON Encoder = 1

	ENCODER_CONSOLE Encoder = 2
)

// 新建日志
func NewLogger(logLevel zapcore.Level, writers []io.Writer, encoderType Encoder, options ...zap.Option) *zap.Logger {
	// 日志格式配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	var encoder zapcore.Encoder
	if encoderType == ENCODER_JSON {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	// 日志输出配置
	var ws zapcore.WriteSyncer
	if writers == nil || len(writers) == 0 {
		writers = make([]io.Writer, 0)
		writers = append(writers, os.Stdout)
	}

	ws = zapcore.AddSync(io.MultiWriter(writers...))
	core := zapcore.NewCore(encoder, ws, logLevel)
	options = append(options, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return zap.New(core, options...)

}

// 获取日志级别
func ParseLevel(lv string) zapcore.Level {
	switch lv {
	case "":
		return zapcore.DebugLevel
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.DebugLevel
	}
}
