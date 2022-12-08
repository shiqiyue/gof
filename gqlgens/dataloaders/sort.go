package dataloaders

import "golang.org/x/exp/slices"

func SortByOrder[Key comparable, Value any](order []Key, src []*Value, getKey func(data *Value) Key) []*Value {
	temp := make(map[Key]*Value, len(src))
	for _, value := range src {
		temp[getKey(value)] = value
	}
	dst := make([]*Value, 0, len(order))
	for _, key := range order {
		dst = append(dst, temp[key])
	}
	return dst
}

func Sort2DByOrder[Key comparable, Value any](order []Key, src []*Value, getKey func(data *Value) Key, less func(a, b *Value) bool) [][]*Value {
	temp := make(map[Key][]*Value, len(src))
	for _, value := range src {
		tempKey := getKey(value)
		temp[tempKey] = append(temp[tempKey], value)
	}
	dst := make([][]*Value, 0, len(order))
	for _, key := range order {
		dstItem := temp[key]
		slices.SortStableFunc(dstItem, less)
		dst = append(dst, dstItem)
	}
	return dst
}
