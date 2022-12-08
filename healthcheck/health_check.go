package healthcheck

import (
	"context"
	"fmt"
	"github.com/shiqiyue/gof/ferror"
)

// 健康检查
func HealthCheck(ctx context.Context) error {
	for _, probe := range probes {
		err := probe.IsHealth(ctx)
		if err != nil {
			return ferror.Wrap(fmt.Sprintf("%s健康检查未通过", probe.Name()), err)
		}
	}
	return nil
}
