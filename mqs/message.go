package mqs

type Message interface {
	IsMqMessage() bool
}

// 通用消息
type CommonMessage struct {
	// 主题或者exchange
	TopicOrExchange string
	// key
	Key string
	// 内容
	Body []byte
	// 消息ID，启动Mq持久化后才有，持久化存储返回的
	MsgId *int64

	// rabbitmq使用
	// 是否立即发送
	Immediate bool

	// rocketmq使用
	// 是否oneway
	IsOneway bool
}

func (c CommonMessage) IsMqMessage() bool {
	return true
}
