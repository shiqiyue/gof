package reflects

import "reflect"

// 只支持一个指针，支持 *People, 不支持**People
func NewByType(t reflect.Type) interface{} {
	if t == nil {
		return nil
	}
	switch t.Kind() {
	case reflect.Ptr:
		return newByType(t.Elem(), true)
	default:
		return newByType(t, false)
	}

}

func newByType(t reflect.Type, isPtr bool) interface{} {
	if t == nil {
		return nil
	}
	switch t.Kind() {

	case reflect.Slice, reflect.Array, reflect.Func, reflect.Interface, reflect.Map, reflect.Struct:
		return newValue(t, isPtr).Interface()
	case reflect.Int:
		return int(newValue(t, isPtr).Int())
	case reflect.Bool:
		return newValue(t, isPtr).Bool()
	case reflect.Int8:
		return int8(newValue(t, isPtr).Int())
	case reflect.Int16:
		return int16(newValue(t, isPtr).Int())
	case reflect.Int32:
		return int32(newValue(t, isPtr).Int())
	case reflect.Int64:
		return int64(newValue(t, isPtr).Int())
	case reflect.Uint:
		return uint(newValue(t, isPtr).Uint())
	case reflect.Uint8:
		return uint8(newValue(t, isPtr).Uint())
	case reflect.Uint16:
		return uint16(newValue(t, isPtr).Uint())
	case reflect.Uint32:
		return uint32(newValue(t, isPtr).Uint())
	case reflect.Uint64:
		return uint64(newValue(t, isPtr).Uint())
	case reflect.Float64:
		return float64(newValue(t, isPtr).Float())
	case reflect.Float32:
		return float32(newValue(t, isPtr).Float())
	case reflect.String:
		return newValue(t, isPtr).String()
	default:
		panic("unsupported type")
	}

}

func newValue(t reflect.Type, isPtr bool) reflect.Value {
	if isPtr {
		return reflect.New(t)
	} else {
		return reflect.New(t).Elem()
	}
}
