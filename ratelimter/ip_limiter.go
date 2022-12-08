package ratelimter

import "time"

// ip流量限制
type IpRateLimiterConf struct {
	// 时间间隔
	PerTime time.Duration
	// 数量
	Num int32
}

// Ip限制数量-提供者
type IpRateLimiterNumProvider interface {
	// 通过Ip获取限制数量
	GetByIp(ip string) int32
}

// ip流量限制
type IpRateLimiter struct {
	limiterNumProvider IpRateLimiterNumProvider
}
