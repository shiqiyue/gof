package times

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"time"
)

type StopWatchStatus int

const (
	NEW = iota + 1
	END
)

// 秒表
type StopWatch struct {
	// 名称
	Name string
	// 等级
	Level int
	// 开始时间
	StartTime time.Time
	// 结束时间
	EndTime *time.Time
	// 持续时间
	Duration *time.Duration
	// 子秒表
	taskMap map[string]*StopWatch
	// 子秒表
	taskList []*StopWatch
	lock     sync.Mutex
	status   StopWatchStatus
}

// 新建秒表
func NewStopWatch(name string) *StopWatch {
	return &StopWatch{Name: name,
		lock:      sync.Mutex{},
		Level:     1,
		taskMap:   map[string]*StopWatch{},
		taskList:  []*StopWatch{},
		status:    NEW,
		StartTime: time.Now()}
}

// 开始
func (w *StopWatch) NewStopWatch(name string) *StopWatch {
	w.lock.Lock()
	defer w.lock.Unlock()
	item, ok := w.taskMap[name]
	if ok {
		return item
	}
	item = NewStopWatch(name)
	item.Level = w.Level + 1
	w.taskMap[name] = item
	w.taskList = append(w.taskList, item)
	return item

}

// 任务结束
func (t *StopWatch) End() {
	if t.status == END {
		return
	}
	for _, task := range t.taskList {
		task.End()
	}
	now := time.Now()
	t.EndTime = &now
	d := time.Since(t.StartTime)
	t.Duration = &d
	t.status = END

}

func (t *StopWatch) ToString() string {
	buffer := bytes.Buffer{}

	header := strings.Repeat("   ", t.Level-1)
	buffer.WriteString(header)
	buffer.WriteString(t.Name)
	buffer.WriteString(":")
	buffer.WriteString(fmt.Sprint(t.Duration.Milliseconds(), "ms"))
	buffer.WriteString("\n")
	for _, task := range t.taskList {
		buffer.WriteString(task.ToString())
	}
	return buffer.String()
}
