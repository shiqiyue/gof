package rabbit

import "github.com/streadway/amqp"

// rabbitmq消息
type RabbitmqMessage struct {
	Exchange string

	Matatory bool

	Immediate bool

	// key
	Key string
	// 消息
	Message *amqp.Publishing
}

func (r RabbitmqMessage) IsMqMessage() bool {
	return true
}

func (r *RabbitmqMessage) AddHeader(name string, key interface{}) {
	if r.Message.Headers == nil {
		r.Message.Headers = make(map[string]interface{}, 0)
	}
	r.Message.Headers[name] = key
}
