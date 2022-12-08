package mqs

//go:generate go-options managementCfg
// managementCfg mq管理器配置
type managementCfg struct {
	// 持久化存储
	Storage Storage
}
