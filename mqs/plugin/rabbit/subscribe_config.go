package rabbit

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	amqptransport "github.com/go-kit/kit/transport/amqp"
	"github.com/shiqiyue/gof/ferror"
	"github.com/shiqiyue/gof/mqs"
	"github.com/streadway/amqp"
)

type RabbitMqSubscribeConfig struct {
	// --- 以下是qos配置
	PrefetchCount int

	PrefetchSize int

	Global bool
	// --- 以下是consumer配置
	Queue string

	ConsumerName string

	Exclusive bool

	NoLocal bool

	NoWait bool

	Args amqp.Table

	// --- 以下配置
	// go kit
	Handler endpoint.Endpoint
	// 入参解析
	DecodeRequestFunc amqptransport.DecodeRequestFunc
	// 出参渲染
	EncodeResponseFunc amqptransport.EncodeResponseFunc
	// 返回消息发布到新的队列
	ResponsePublisher amqptransport.ResponsePublisher

	Storage mqs.Storage
}

func (r RabbitMqSubscribeConfig) IsSubscribeConfig() bool {
	return true
}

func warpCommonHandle(handle func(ctx context.Context, body []byte) error) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		body, ok := request.([]byte)
		if !ok {
			return nil, ferror.New("request is not byte array")
		}
		err = handle(ctx, body)
		return nil, err
	}
}

func defaultDecodeRequestFunc(ctx context.Context, d *amqp.Delivery) (request interface{}, err error) {
	return d.Body, nil
}

func EncodeJSONResponse(
	ctx context.Context,
	pub *amqp.Publishing,
	response interface{},
) error {
	if response == nil {
		return nil
	}
	b, err := json.Marshal(response)
	if err != nil {
		return err
	}
	pub.Body = b
	return nil
}
