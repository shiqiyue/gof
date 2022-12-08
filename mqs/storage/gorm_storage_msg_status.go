package storage

//go:generate go-enum -f=$GOFILE --marshal --names --ptr
// gormStorageMsgStatus 消息状态
/*
ENUM(
NOT_SEND = 1// 未发送
SEND_SUCCESS = 2// 发送成功
SEND_FAIL = 3// 发送失败
CONSUME_SUCCESS = 4// 消费成功
CONSUME_FAIL = 5// 消费失败
)
*/
type gormStorageMsgStatus int32
