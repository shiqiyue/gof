package beans

import (
	"fmt"
	"github.com/shiqiyue/inject"
	"reflect"
)

var (
	injectGraph = &inject.Graph{}
)

// 将对象注册到injectGraph
func ProvideBean(v interface{}) error {
	return Provide(&inject.Object{Value: v})
}

// 将对象注册到injectGraph
func Provide(objects ...*inject.Object) error {
	err := injectGraph.Provide(objects...)
	return err
}

func PrintAll() {
	objects := injectGraph.Objects()
	fmt.Println(objects)
}

// 根据对象的标签，完成对象属性的注入
func Populate() error {
	return injectGraph.Populate()
}

// 通过名称获取对象
func GetObjectByName(name string) interface{} {
	for _, o := range injectGraph.Objects() {
		if o.Name == name {
			return o.Value
		}

	}
	return nil
}

// 通过类型获取对象
func GetObjectByType(value interface{}) (res []interface{}) {
	rv := reflect.TypeOf(value)
	for _, o := range injectGraph.Objects() {
		if reflect.TypeOf(o.Value) == rv {
			res = append(res, o.Value)
		}
	}
	return
}

// 通过类型获取对象
func GetOneObjectByType(value interface{}) interface{} {
	rv := reflect.TypeOf(value)
	for _, o := range injectGraph.Objects() {
		if reflect.TypeOf(o.Value) == rv {
			return o.Value
		}
	}
	return nil
}
