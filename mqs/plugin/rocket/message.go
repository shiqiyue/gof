package rocket

import "github.com/apache/rocketmq-client-go/v2/primitive"

// rocketmq消息
type RocketMqMessage struct {
	// 消息
	Message *primitive.Message
	// oneway
	OneWay bool
}

func (r *RocketMqMessage) isMqMessage() bool {
	return true
}
