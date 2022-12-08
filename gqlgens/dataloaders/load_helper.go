package dataloaders

import "sync"

type LoadHelper[Value any] struct {
	sync.Mutex
	value *Value
}

func (helper *LoadHelper[Value]) Get(initer func() *Value) *Value {
	if helper.value == nil {
		helper.Lock()
		defer helper.Unlock()
		if helper.value == nil {
			helper.value = initer()
		}
	}
	return helper.value
}
