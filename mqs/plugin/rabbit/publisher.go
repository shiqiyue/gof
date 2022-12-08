package rabbit

import (
	"context"
	"github.com/isayme/go-amqp-reconnect/rabbitmq"
	uuid2 "github.com/pborman/uuid"
	"github.com/shiqiyue/gof/asserts"
	"github.com/streadway/amqp"
)

type Publisher struct {
	// 渠道
	Channel *rabbitmq.Channel
	// exchange
	Exchange string
	// mandatory
	Matatory bool
	// Immediate
	Immediate bool
}

// 新建发送者
func NewPublisher(conn *rabbitmq.Connection, exchange string, matatory, immediate bool) *Publisher {
	channel, err := conn.Channel()
	asserts.Nil(err, err)
	return &Publisher{
		Channel:   channel,
		Exchange:  exchange,
		Matatory:  matatory,
		Immediate: immediate,
	}
}

// 新建默认配置的发送者
func NewDefaultPublisher(conn *rabbitmq.Connection, exchange string) *Publisher {

	return NewPublisher(conn, exchange, false, false)
}

// 发送
func (p *Publisher) Publish(ctx context.Context, key string, publishing *amqp.Publishing) error {
	if publishing.MessageId == "" {
		publishing.MessageId = uuid2.NewUUID().String()
	}
	if publishing.Headers == nil {
		publishing.Headers = map[string]interface{}{}
	}
	sp := opentracing.SpanFromContext(ctx)
	if sp != nil {
		opentracing.GlobalTracer().Inject(sp.Context(), opentracing.TextMap, publishing.Headers)
	}
	return p.Channel.Publish(p.Exchange, key, p.Matatory, p.Immediate, *publishing)
}
