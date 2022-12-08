package healthcheck

import "context"

// 健康检查-探针
type Probe interface {
	// 是否健康
	IsHealth(ctx context.Context) error
	// 探针名称
	Name() string
}

type Probes []Probe

var probes Probes = make([]Probe, 0)

// 添加探针
func AddProbe(probe Probe) {
	probes = append(probes, probe)
}
