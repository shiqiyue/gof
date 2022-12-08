package mqs

import "context"

// 客户端
type Client interface {
	// 发送记录
	Publish(ctx context.Context, message Message) error
	// 订阅
	Subscribe(ctx context.Context, config SubscribeConfig) error
}
