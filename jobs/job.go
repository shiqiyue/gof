package jobs

import (
	"context"
	"github.com/shiqiyue/gof/loggers"
	"go.uber.org/zap"
)

type JobFunc func(ctx context.Context) error

var jobs map[string]*jobWrapper = make(map[string]*jobWrapper, 0)

// 注册任务
func RegisterJob(jobName string, jobFunc JobFunc) {
	jobs[jobName] = &jobWrapper{jobFunc: jobFunc, jobName: jobName}
}

// 获取任务
func GetJobs() map[string]*jobWrapper {
	return jobs
}

type jobWrapper struct {
	jobFunc JobFunc
	jobName string
}

func (j jobWrapper) Run() {
	ctx := context.Background()
	defer func() {
		if err := recover(); err != nil {
			loggers.Error(ctx, "执行任务panic", zap.String("job", j.jobName), zap.Any("err", err))
		}
	}()
	err := j.jobFunc(ctx)
	if err != nil {
		loggers.Error(ctx, "执行任务异常", zap.String("job", j.jobName), zap.Error(err))
	}

}
