package asserts

import (
	"context"
	"errors"
	"github.com/shiqiyue/gof/loggers"
	"go.uber.org/zap"
)

// 断言，不为nil
func NotNil(i interface{}, err error) {
	if i == nil {
		if err != nil {
			loggers.Error(context.Background(), "value is nil", zap.Error(err))
			panic(err)
		} else {
			loggers.Error(context.Background(), "value is nil", zap.Error(errors.New("value is  nil")))
			panic(errors.New("value is  nil"))
		}
	}
}

// 断言，为nil
func Nil(i interface{}, err error) {
	if i != nil {
		if err != nil {
			loggers.Error(context.Background(), "value is not nil", zap.Error(err))
			panic(err)
		} else {
			loggers.Error(context.Background(), "value is not nil", zap.Error(errors.New("value is not nil")))
			panic(errors.New("value is not nil"))
		}

	}
}
