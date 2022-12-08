package rabbit

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	kitOpentracing "github.com/go-kit/kit/tracing/opentracing"
	amqptransport "github.com/go-kit/kit/transport/amqp"
	"github.com/isayme/go-amqp-reconnect/rabbitmq"
	"github.com/shiqiyue/gof/loggers"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// consumer 消费者
type consumer struct {
	connection *rabbitmq.Connection

	channel *rabbitmq.Channel

	config *RabbitMqSubscribeConfig

	subscriber *amqptransport.Subscriber
}

// startConsume 开始消费
func (c *consumer) startConsume() error {
	msgs, err := c.channel.Consume(
		c.config.Queue,        // queue
		c.config.ConsumerName, // consumer
		false,                 // auto-ack
		c.config.Exclusive,    // exclusive
		c.config.NoLocal,      // no-local
		c.config.NoWait,       // no-wait
		c.config.Args,         // args
	)
	if err != nil {
		return err
	}

	// opentracing
	e := c.config.Handler
	tracer := opentracing.GlobalTracer()
	if tracer != nil {
		e = kitOpentracing.TraceServer(tracer, c.config.ConsumerName)(e)
	}

	subscriber := amqptransport.NewSubscriber(e,
		c.config.DecodeRequestFunc,
		c.config.EncodeResponseFunc,
		amqptransport.SubscriberErrorEncoder(c.AckAndRepublish5TimesErrorEncoder),
		amqptransport.SubscriberBefore(c.NewOpentracingBefore(), c.NewLoggerDelivery()),
		amqptransport.SubscriberAfter(amqptransport.SetAckAfterEndpoint(false), c.OpentracingAfter),
		amqptransport.SubscriberAfter(c.MarkMessageConsumeSuccess),
		amqptransport.SubscriberResponsePublisher(c.config.ResponsePublisher))

	serveDelivery := subscriber.ServeDelivery(c.channel)
	go func() {
		for d := range msgs {
			c.handleDelivery(d, c.channel, serveDelivery)
		}
	}()
	return nil
}

// 新建Config
func NewConfig(queue, consumerName string, Handler endpoint.Endpoint, DecodeRequestFunc amqptransport.DecodeRequestFunc, EncodeResponseFunc amqptransport.EncodeResponseFunc, ResponsePublisher amqptransport.ResponsePublisher) RabbitMqSubscribeConfig {
	return RabbitMqSubscribeConfig{
		PrefetchCount:      1,
		PrefetchSize:       0,
		Global:             false,
		Queue:              queue,
		ConsumerName:       consumerName,
		Exclusive:          false,
		NoLocal:            false,
		NoWait:             false,
		Args:               nil,
		Handler:            Handler,
		DecodeRequestFunc:  DecodeRequestFunc,
		EncodeResponseFunc: EncodeResponseFunc,
		ResponsePublisher:  ResponsePublisher,
	}
}

func NewCommonConfig(queue, consumerName string, Handler endpoint.Endpoint, DecodeRequestFunc amqptransport.DecodeRequestFunc) RabbitMqSubscribeConfig {
	return NewConfig(queue, consumerName, Handler, DecodeRequestFunc, EncodeJSONResponse, amqptransport.NopResponsePublisher)
}

// 新建消费者
func NewConsumer(connection *rabbitmq.Connection, config *RabbitMqSubscribeConfig) error {
	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	// 设置请求参数，比如每次拉取多少数据
	err = channel.Qos(
		config.PrefetchCount, // prefetch count
		config.PrefetchSize,  // prefetch size
		config.Global,        // global
	)
	if err != nil {
		return err
	}

	c := &consumer{
		connection: connection,
		channel:    channel,
		config:     config,
		subscriber: nil,
	}
	err = c.startConsume()
	if err != nil {
		return err
	}
	return nil

}

func (c *consumer) handleDelivery(d amqp.Delivery, channel *rabbitmq.Channel, serveDelivery func(deliv *amqp.Delivery)) {
	defer func() {
		// panic处理
		if err := recover(); err != nil {
			e, ok := err.(error)
			var finalErr error
			if !ok {
				finalErr = errors.New(fmt.Sprint(err))
			} else {
				finalErr = e
			}
			c.AckAndRepublish5TimesErrorEncoder(context.Background(), finalErr, &d, channel, nil)
		}
	}()
	serveDelivery(&d)
}

// opentracing集成 / mq消费前执行
func (c *consumer) NewOpentracingBefore() func(ctx context.Context, publishing *amqp.Publishing, delivery *amqp.Delivery) context.Context {
	return func(ctx context.Context, publishing *amqp.Publishing, delivery *amqp.Delivery) context.Context {
		spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, delivery.Headers)
		var sp opentracing.Span
		operateName := "MQ: " + c.config.Queue

		if spanCtx == nil {
			sp = opentracing.GlobalTracer().StartSpan(operateName)
		} else {
			sp = opentracing.GlobalTracer().StartSpan(operateName, opentracing.ChildOf(spanCtx))
		}
		sp = sp.SetTag("exchange", delivery.Exchange)
		sp = sp.SetBaggageItem("op", operateName)

		return ctx
	}
}

// opentracing集成 / mq消费后执行
func (c *consumer) OpentracingAfter(ctx context.Context,
	deliv *amqp.Delivery,
	ch amqptransport.Channel,
	pub *amqp.Publishing) context.Context {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.Finish()
	}
	return ctx
}

// MarkMessageConsumeSuccess 标识消息消费成功
func (c *consumer) MarkMessageConsumeSuccess(ctx context.Context, delivery *amqp.Delivery, channel amqptransport.Channel, publishing *amqp.Publishing) context.Context {
	// 如果消息头包含有消费ID，则调用存储接口将消费设置成消费成功
	if delivery.Headers != nil {
		msgId, msgIdExist := delivery.Headers[MSG_ID_HEADER_NAME].(int64)
		if msgIdExist {
			if c.config.Storage != nil {
				err := c.config.Storage.MarkMessageConsumeSuccess(ctx, msgId)
				if err != nil {
					loggers.Error(ctx, "设置消息消费成功异常", zap.Error(err))
				}
			}
		}
	}
	return ctx
}

// MarkMessageConsumeSuccess 标识消息消费失败
func (c *consumer) MarkMessageConsumeFail(ctx context.Context, delivery *amqp.Delivery, channel amqptransport.Channel, publishing *amqp.Publishing) context.Context {
	// 如果消息头包含有消费ID，则调用存储接口将消费设置成消费成功
	if delivery.Headers != nil {
		msgId, msgIdExist := delivery.Headers[MSG_ID_HEADER_NAME].(int64)
		if msgIdExist {
			if c.config.Storage != nil {
				err := c.config.Storage.MarkMessageConsumeSuccess(ctx, msgId)
				if err != nil {
					loggers.Error(ctx, "设置消息消费成功异常", zap.Error(err))
				}
			}
		}
	}
	return ctx
}

func (c *consumer) NewLoggerDelivery() func(ctx context.Context, publishing *amqp.Publishing, delivery *amqp.Delivery) context.Context {
	return func(ctx context.Context, publishing *amqp.Publishing, delivery *amqp.Delivery) context.Context {
		loggers.Info(ctx, "MQ: "+c.config.Queue)
		loggers.Info(ctx, "消息内容", zap.String("MessageId", delivery.MessageId), zap.String("exchange", delivery.Exchange), zap.Any("Headers", delivery.Headers), zap.Uint64("deliveryTag", delivery.DeliveryTag))
		return ctx
	}
}

// ack，并且重新发送到队列5次
func (c *consumer) AckAndRepublish5TimesErrorEncoder(ctx context.Context, err error, deliv *amqp.Delivery, ch amqptransport.Channel, pub *amqp.Publishing) {
	loggers.Error(ctx, "消费消息异常", zap.Any("headers", deliv.Headers), zap.String("MessageId", deliv.MessageId), zap.String("exchange", deliv.Exchange), zap.Error(err), zap.Stack("stack"))
	if deliv.Headers == nil {
		deliv.Headers = make(map[string]interface{}, 0)
	}
	verrorCount := deliv.Headers[ERROR_COUNT_HEADER_NAME]
	var errcount int32 = 0
	if verrorCount != nil {
		errcount, _ = verrorCount.(int32)
	}
	errcount++
	if errcount >= 5 {
		loggers.Error(ctx, "消费消息严重异常，失败五次",
			zap.Any("headers", deliv.Headers),
			zap.String("MessageId", deliv.MessageId),
			zap.String("exchange", deliv.Exchange),
			zap.Any("body", deliv.Body),
			zap.Error(err),
			zap.Stack("stack"))
		c.MarkMessageConsumeFail(ctx, deliv, ch, pub)
		_ = deliv.Nack(false, false)
	} else {
		_ = deliv.Ack(false)
		deliv.Headers[ERROR_COUNT_HEADER_NAME] = errcount
		deliv.Headers[DELAY_HEADER_NAME] = c.getDelayByErrorCount(errcount)
		_ = ch.Publish(deliv.Exchange, deliv.RoutingKey, false, false, amqp.Publishing{
			Headers:      deliv.Headers,
			DeliveryMode: deliv.DeliveryMode,
			Priority:     deliv.Priority,
			Expiration:   deliv.Expiration,
			MessageId:    deliv.MessageId,
			Timestamp:    deliv.Timestamp,
			Type:         deliv.Type,
			UserId:       deliv.UserId,
			AppId:        deliv.AppId,
			Body:         deliv.Body,
		})
	}

}

// 通过错误次数，获取延迟时间
func (c *consumer) getDelayByErrorCount(errcount int32) int32 {
	switch errcount {
	case 1:
		return 1000
	case 2:
		return 5000
	case 3:
		return 10000
	case 4:
		return 20000
	case 5:
		return 30000
	default:
		return 30000
	}
}
