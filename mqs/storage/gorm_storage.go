package storage

import (
	"context"
	"encoding/json"
	"github.com/shiqiyue/gof/ferror"
	"github.com/shiqiyue/gof/gorms"
	"github.com/shiqiyue/gof/mqs"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type gormStorage struct {
	Db *gorm.DB
}

//go:generate goqueryset -in $GOFILE
// gen:qs
// gorm存储消息
type gormStorageMsg struct {
	Id           int64          `gorm:"primaryKey;comment:主键"`
	CreatedAt    time.Time      `gorm:"not null;comment:创建时间"`
	UpdatedAt    time.Time      `gorm:"not null;comment:更新时间"`
	MqClientName string         `gorm:"not null;comment:MQ客户端名称"`
	Message      datatypes.JSON `gorm:"not null;comment:消息"`
	Status       int            `gorm:"not null; comment:状态"`
	FailReason   *string        `gorm:"comment:失败原因"`
}

func (g *gormStorageMsg) TableName() string {
	return "mq_message"
}

func NewGormStorage(db *gorm.DB) (*gormStorage, error) {
	err := db.AutoMigrate(&gormStorageMsg{})
	if err != nil {
		return nil, err
	}
	return &gormStorage{Db: db}, nil
}

func (g *gormStorage) StorageMessage(ctx context.Context, clientName string, message *mqs.CommonMessage) (int64, error) {
	db := gorms.GetDb(ctx, g.Db)
	bs, err := json.Marshal(message)
	if err != nil {
		return 0, ferror.Wrap("序列化消息异常", err)
	}
	msg := &gormStorageMsg{
		MqClientName: clientName,
		Message:      bs,
		Status:       int(GormStorageMsgStatusNOTSEND),
	}
	err = msg.Create(db)
	if err != nil {
		return 0, ferror.Wrap("持久化消息异常", err)
	}
	return msg.Id, nil
}

func (g *gormStorage) MarkMessageSendSuccess(ctx context.Context, ids []int64) error {
	db := gorms.GetDb(ctx, g.Db)
	err := NewgormStorageMsgQuerySet(db).IdIn(ids...).GetUpdater().SetStatus(int(GormStorageMsgStatusSENDSUCCESS)).Update()
	if err != nil {
		return ferror.Wrap("设置消息发送成功异常", err)
	}
	return nil
}

func (g *gormStorage) MarkMessageSendFail(ctx context.Context, id int64, failReason string) error {
	db := gorms.GetDb(ctx, g.Db)
	err := NewgormStorageMsgQuerySet(db).IdEq(id).GetUpdater().SetStatus(int(GormStorageMsgStatusSENDFAIL)).SetFailReason(&failReason).Update()
	if err != nil {
		return ferror.Wrap("设置消息发送失败异常", err)
	}
	return nil
}

func (g *gormStorage) MarkMessageConsumeSuccess(ctx context.Context, id int64) error {
	db := gorms.GetDb(ctx, g.Db)
	err := NewgormStorageMsgQuerySet(db).IdEq(id).GetUpdater().SetStatus(int(GormStorageMsgStatusCONSUMESUCCESS)).Update()
	if err != nil {
		return ferror.Wrap("设置消息消费成功异常", err)
	}
	return nil
}

func (g *gormStorage) MarkMessageConsumeFail(ctx context.Context, id int64, failReason string) error {
	db := gorms.GetDb(ctx, g.Db)
	err := NewgormStorageMsgQuerySet(db).IdEq(id).GetUpdater().SetStatus(int(GormStorageMsgStatusCONSUMEFAIL)).SetFailReason(&failReason).Update()
	if err != nil {
		return ferror.Wrap("设置消息消费失败异常", err)
	}
	return nil
}
