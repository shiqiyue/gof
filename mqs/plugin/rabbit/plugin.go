package rabbit

import (
	"context"
	"github.com/isayme/go-amqp-reconnect/rabbitmq"
	"github.com/shiqiyue/gof/mqs"
	"sync"
)

func init() {
	mqs.RegisterPlugin(&rabbitMqPlugin{})
}

type rabbitMqPlugin struct {
}

func (r rabbitMqPlugin) NewClient(ctx context.Context, url string, storage mqs.Storage, extra map[string]string) (mqs.Client, error) {
	connection, err := rabbitmq.Dial(url)
	if err != nil {
		return nil, err
	}
	client := &rabbitMqClient{connection: connection, lock: &sync.RWMutex{}, publisherMap: map[string]*Publisher{}, storage: storage}
	return client, nil
}

func (r rabbitMqPlugin) MqType() string {
	return "rabbitmq"
}
