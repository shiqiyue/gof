package mqs

import "context"

type MqMessage struct {
	ClientName string
}

type Storage interface {
	// StorageMessage 将消息持久化, 返回消息ID或者异常
	StorageMessage(ctx context.Context, clientName string, message *CommonMessage) (int64, error)
	// MarkMessageSendSuccess 标志消息发送成功
	MarkMessageSendSuccess(ctx context.Context, ids []int64) error
	// MarkMessageSendFail 标志消息发送失败
	MarkMessageSendFail(ctx context.Context, id int64, failReason string) error
	// MarkMessageConsumeSuccess 标志消息消费成功
	MarkMessageConsumeSuccess(ctx context.Context, id int64) error
	// MarkMessageConsumeFail 标志消息消费失败
	MarkMessageConsumeFail(ctx context.Context, id int64, failReason string) error
}
