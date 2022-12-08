package mqs

import (
	"context"
)

type SubscribeConfig interface {
	IsSubscribeConfig() bool
}

type CommonSubscribeConfig struct {
	// 队列或者主题名称
	QueueOrTopic string

	// 消费者名称
	ConsumerName string

	// 处理器
	Handle func(ctx context.Context, body []byte) error
}

func (c CommonSubscribeConfig) IsSubscribeConfig() bool {
	return true
}
