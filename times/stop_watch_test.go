package times

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewStopWatch(t *testing.T) {
	mainStopWatch := NewStopWatch("图片处理")
	image1 := mainStopWatch.NewStopWatch("图片1处理")
	image2 := mainStopWatch.NewStopWatch("图片2处理")
	w := sync.WaitGroup{}
	w.Add(2)
	go func() {
		defer func() {
			image1.End()
			w.Done()
		}()
		move := image1.NewStopWatch("移动")
		time.Sleep(time.Second * 2)
		move.End()
		resize := image1.NewStopWatch("变换大小")
		time.Sleep(time.Second * 1)
		resize.End()
	}()
	go func() {
		defer func() {
			image2.End()
			w.Done()
		}()
		move := image2.NewStopWatch("移动")
		time.Sleep(time.Second * 3)
		move.End()
		resize := image2.NewStopWatch("变换大小")
		time.Sleep(time.Second * 2)
		resize.End()
	}()
	w.Wait()
	mainStopWatch.End()
	fmt.Println(mainStopWatch.ToString())
}
