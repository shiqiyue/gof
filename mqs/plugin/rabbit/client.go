package rabbit

import (
	"context"
	amqptransport "github.com/go-kit/kit/transport/amqp"
	"github.com/isayme/go-amqp-reconnect/rabbitmq"
	"github.com/shiqiyue/gof/ferror"
	"github.com/shiqiyue/gof/mqs"
	"github.com/streadway/amqp"
	"sync"
)

type rabbitMqClient struct {
	connection *rabbitmq.Connection

	publisherMap map[string]*Publisher

	lock *sync.RWMutex

	storage mqs.Storage
}

func (r *rabbitMqClient) Publish(ctx context.Context, message mqs.Message) error {
	rabbitmqMessage, ok := message.(*RabbitmqMessage)
	if !ok {
		commonMessage, ok := message.(*mqs.CommonMessage)
		if !ok {
			return ferror.New("消息类型错误")
		}
		rabbitmqMessage = &RabbitmqMessage{
			Exchange:  commonMessage.TopicOrExchange,
			Matatory:  false,
			Immediate: commonMessage.Immediate,
			Key:       commonMessage.Key,
			Message: &amqp.Publishing{DeliveryMode: amqp.Persistent,
				ContentType: "text/plain",
				Body:        commonMessage.Body},
		}
		if commonMessage.MsgId != nil {
			rabbitmqMessage.AddHeader(MSG_ID_HEADER_NAME, *commonMessage.MsgId)
		}
	}
	publisher, err := r.getPublisher(ctx, rabbitmqMessage)
	if err != nil {
		return err
	}
	err = publisher.Publish(ctx, rabbitmqMessage.Key, rabbitmqMessage.Message)

	return err
}

func (r *rabbitMqClient) Subscribe(ctx context.Context, config mqs.SubscribeConfig) error {
	subscribeConfig, ok := config.(*RabbitMqSubscribeConfig)
	if !ok {
		commonSubscribeConfig, ok := config.(*mqs.CommonSubscribeConfig)
		if !ok {
			return ferror.New("订阅配置类型错误")
		}
		subscribeConfig = &RabbitMqSubscribeConfig{
			PrefetchCount:      5,
			PrefetchSize:       0,
			Global:             false,
			Queue:              commonSubscribeConfig.QueueOrTopic,
			ConsumerName:       commonSubscribeConfig.ConsumerName,
			Exclusive:          false,
			NoLocal:            false,
			NoWait:             false,
			Args:               nil,
			Handler:            warpCommonHandle(commonSubscribeConfig.Handle),
			DecodeRequestFunc:  defaultDecodeRequestFunc,
			EncodeResponseFunc: EncodeJSONResponse,
			ResponsePublisher:  amqptransport.NopResponsePublisher,
			Storage:            r.storage,
		}
	}
	err := NewConsumer(r.connection, subscribeConfig)
	return err
}

// 获取发布者
func (r *rabbitMqClient) getPublisher(ctx context.Context, message *RabbitmqMessage) (*Publisher, error) {
	publisher, publisherExist := r.publisherMap[message.Exchange]
	if publisherExist {
		return publisher, nil
	}
	return r.createPublisher(ctx, message)
}

func (r *rabbitMqClient) createPublisher(ctx context.Context, message *RabbitmqMessage) (*Publisher, error) {
	r.lock.Lock()
	defer func() {
		r.lock.Unlock()
	}()
	publisher := NewPublisher(r.connection, message.Exchange, message.Matatory, message.Immediate)
	r.publisherMap[message.Exchange] = publisher
	return publisher, nil
}
