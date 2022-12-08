package loggers

import "go.uber.org/zap"

type iLogger struct {
	ls map[string]*zap.Logger
}

func (log *iLogger) Info(msg string, fields ...zap.Field) {
	for _, l := range log.ls {
		l.Info(msg, fields...)
	}
}

func (log *iLogger) Debug(msg string, fields ...zap.Field) {
	for _, l := range log.ls {
		l.Debug(msg, fields...)
	}
}

func (log *iLogger) Warn(msg string, fields ...zap.Field) {
	for _, l := range log.ls {
		l.Warn(msg, fields...)
	}
}

func (log *iLogger) Error(msg string, fields ...zap.Field) {
	for _, l := range log.ls {
		l.Error(msg, fields...)
	}
}

func (log *iLogger) Fatal(msg string, fields ...zap.Field) {
	for _, l := range log.ls {
		l.Fatal(msg, fields...)
	}
}
