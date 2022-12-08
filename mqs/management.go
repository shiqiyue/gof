package mqs

import (
	"context"
	"github.com/shiqiyue/gof/ferror"
	"strings"
)

// 默认客户端名称
var defaultClientName = "default"

var prepareMessageInfosKey = struct {
}{}

type prepareMessageInfo struct {
	messageId  int64
	message    *CommonMessage
	clientName string
}

type PublishStorageMessageError struct {
	Errs []error
}

func (p *PublishStorageMessageError) Error() string {
	ss := make([]string, 0)
	for _, err := range p.Errs {
		ss = append(ss, err.Error())
	}
	return strings.Join(ss, "\n")
}

// 准备发送的消息列表
type prepareMessageInfos *[]*prepareMessageInfo

// mq管理
// 保存mq客户端示例
type Management struct {
	// 客户端
	clientHolder map[string]Client
	// 持久化存储
	Storage Storage
}

func NewManagement() *Management {
	return &Management{clientHolder: map[string]Client{}}
}

func NewManagementWithOption(options ...Option) (*Management, error) {
	cfg, err := newManagementCfg(options...)
	if err != nil {
		return nil, err
	}
	management := &Management{
		clientHolder: map[string]Client{},
		Storage:      cfg.Storage,
	}
	return management, nil
}

// 新建客户端
func (m *Management) NewClient(ctx context.Context, cfg ClientCfg) error {
	for _, plugin := range plugins {
		if plugin.MqType() == cfg.Type {
			client, err := plugin.NewClient(ctx, cfg.Url, m.Storage, cfg.Extra)
			if err != nil {
				return err
			}
			m.addClientToHolder(cfg.Name, client)
			return nil
		}
	}

	return ferror.New("不支持的mq类型")
}

// PrepareStorageMessage 准备持久化消息,设置messageIds到上下文中
func PrepareStorageMessage(ctx context.Context) context.Context {
	_, exist := ctx.Value(prepareMessageInfosKey).(prepareMessageInfos)
	if exist {
		return ctx
	}
	pMsgInfos := make([]*prepareMessageInfo, 0)
	return context.WithValue(ctx, prepareMessageInfosKey, prepareMessageInfos(&pMsgInfos))
}

// getPrepareMessageInfos 从上下文中获取prepareMessageInfos对象
func getPrepareMessageInfos(ctx context.Context) (prepareMessageInfos, error) {
	mids, exist := ctx.Value(prepareMessageInfosKey).(prepareMessageInfos)
	if exist {
		return mids, nil
	}
	return nil, ferror.New("不存在messageIds, 请事先调用PrepareStorageMessage")

}

// StorageMessageWithDefaultClient 使用默认客户端持久化消息
func (m *Management) StorageMessageWithDefaultClient(ctx context.Context, message *CommonMessage) error {
	return m.StorageMessage(ctx, defaultClientName, message)
}

// StorageMessage 持久化消息
func (m *Management) StorageMessage(ctx context.Context, clientName string, message *CommonMessage) error {
	if m.Storage == nil {
		return ferror.New("Storage不存在，请先设置Storage")
	}
	prepareMsgs, err := getPrepareMessageInfos(ctx)
	if err != nil {
		return err
	}
	msgId, err := m.Storage.StorageMessage(ctx, clientName, message)
	if err != nil {
		return err
	}
	*prepareMsgs = append(*prepareMsgs, &prepareMessageInfo{
		messageId:  msgId,
		message:    message,
		clientName: clientName,
	})
	return nil
}

// PublishStorageMessage 发送持久化的消息
func (m *Management) PublishStorageMessage(ctx context.Context) error {
	prepareMsgs, _ := getPrepareMessageInfos(ctx)
	if prepareMsgs == nil || len(*prepareMsgs) == 0 {
		return nil
	}
	errs := make([]error, 0)
	successIds := make([]int64, 0)
	for _, prepareMsg := range *prepareMsgs {
		prepareMsg.message.MsgId = &prepareMsg.messageId
		err := m.PublishMessage(ctx, prepareMsg.clientName, prepareMsg.message)
		if err != nil {
			errs = append(errs, err)
			err := m.Storage.MarkMessageSendFail(ctx, prepareMsg.messageId, err.Error())
			if err != nil {
				errs = append(errs, err)
			}
		} else {
			successIds = append(successIds, prepareMsg.messageId)
		}
	}
	if len(successIds) > 0 {
		err := m.Storage.MarkMessageSendSuccess(ctx, successIds)
		if err != nil {
			return err
		}
	}
	if len(errs) > 0 {
		return &PublishStorageMessageError{Errs: errs}
	}
	return nil
}

// 发布消息
func (m *Management) PublishMessageWithDefaultClient(ctx context.Context, message Message) error {
	return m.PublishMessage(ctx, defaultClientName, message)
}

// 发布消息
func (m *Management) PublishMessage(ctx context.Context, clientName string, message Message) error {
	client := m.getClient(ctx, clientName)
	if client == nil {
		return ferror.New("客户端不存在")
	}
	err := client.Publish(ctx, message)
	return err
}

// 订阅消息
func (m *Management) SubscribeMessageWithDefaultClient(ctx context.Context, config SubscribeConfig) error {

	return m.SubscribeMessage(ctx, defaultClientName, config)
}

// 订阅消息
func (m *Management) SubscribeMessage(ctx context.Context, clientName string, config SubscribeConfig) error {
	client := m.getClient(ctx, clientName)
	if client == nil {
		return ferror.New("客户端不存在")
	}
	err := client.Subscribe(ctx, config)
	return err
}

func (m *Management) addClientToHolder(name string, client Client) {
	m.clientHolder[name] = client
}

func (m *Management) getClient(ctx context.Context, clientName string) Client {
	return m.clientHolder[clientName]
}
