package mqs

type ClientCfg struct {
	// 类型，必填，现在支持rocketmq和rabbitmq
	Type string `mapstructure:"type"`
	// 客户端名称
	Name string `mapstructure:"name"`
	// mq地址
	Url string `mapstructure:"url"`
	// 额外配置，具体看各个mq的实现
	Extra map[string]string `mapstructure:"额外配置"`
}
