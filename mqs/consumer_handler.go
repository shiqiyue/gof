package mqs

import (
	"context"
	"github.com/shiqiyue/gof/asserts"
)

type Handler interface {
	// 处理
	Handle(ctx context.Context, body []byte) error
	// 消费者名称
	ConsumerName() string
	// 监听的队列名称
	QueueName() string
}

type Handlers struct {
	// 消息管理
	MqManagement *Management `inject:""`

	//--------私有属性

	handlers []Handler
}

func (h *Handlers) AddHandler(handler Handler) {
	if h.handlers == nil {
		h.handlers = make([]Handler, 0)
	}
	h.handlers = append(h.handlers, handler)
}

func (h *Handlers) SetUp() {

	ctx := context.Background()
	for _, handler := range h.handlers {
		err := h.MqManagement.SubscribeMessageWithDefaultClient(ctx, &CommonSubscribeConfig{
			QueueOrTopic: handler.QueueName(),
			ConsumerName: handler.ConsumerName(),
			Handle:       handler.Handle,
		})
		asserts.Nil(err, err)
	}

}

// 默认的handlers
var defaultHandlers *Handlers

func RegisterDefaultHandlers(h *Handlers) {
	defaultHandlers = h
}

func AddHandler(h Handler) {
	defaultHandlers.AddHandler(h)
}
