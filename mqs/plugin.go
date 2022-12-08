package mqs

import "context"

var plugins []Plugin = make([]Plugin, 0)

func RegisterPlugin(plugin Plugin) {
	plugins = append(plugins, plugin)
}

type Plugin interface {
	// 新建客户端
	NewClient(ctx context.Context, url string, storage Storage, extra map[string]string) (Client, error)
	// 返回mq类型
	MqType() string
}
