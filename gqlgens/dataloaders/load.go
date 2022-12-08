package dataloaders

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/iancoleman/strcase"

	"github.com/shiqiyue/gof/loggers"
	"go.uber.org/zap"

	"github.com/shiqiyue/gof/gorms"

	"gorm.io/gorm"
)

type LoadKey interface {
	~int | ~int32 | ~int64 | ~string
}

type LoadFunc[Key any, Value any] func(ctx context.Context, order []Key) ([]*Value, []error)

func GenLoad[Key any, Value any](loader LoadFunc[Key, Value], tip *string) LoadFunc[Key, Value] {
	return func(ctx context.Context, order []Key) (rData []*Value, rErr []error) {
		defer func() {
			if r := recover(); r != nil {
				loggers.Error(ctx, "loader异常",
					zap.String("model", fmt.Sprintf("%T", *new(Value))),
					zap.Stringp("tip", tip),
					zap.Any("recover", r),
				)
				rErr = append([]error(nil), errors.New("加载数据异常"))
				return
			}
		}()
		return loader(ctx, order)
	}
}

type Load2DFunc[Key any, Value any] func(ctx context.Context, order []Key) ([][]*Value, []error)

func GenLoad2D[Key any, Value any](loader Load2DFunc[Key, Value], tip *string) Load2DFunc[Key, Value] {
	return func(ctx context.Context, order []Key) (rData [][]*Value, rErr []error) {
		defer func() {
			if r := recover(); r != nil {
				loggers.Error(ctx, "loader异常",
					zap.String("model", fmt.Sprintf("%T", *new(Value))),
					zap.Stringp("tip", tip),
					zap.Any("recover", r),
				)
				rErr = append([]error(nil), errors.New("加载数据异常"))
				return
			}
		}()
		return loader(ctx, order)
	}
}

func GenLoadIn[Key LoadKey, Value any](db *gorm.DB, getter func(data *Value) Key) LoadFunc[Key, Value] {
	filed := getFiledNameByGetter(getter)
	loader := func(ctx context.Context, order []Key) ([]*Value, []error) {
		db := gorms.GetDb(ctx, db)
		data := make([]*Value, 0, len(order))
		where := filed + " in ?"
		err := db.Where(where, order).Find(&data).Error
		if err != nil {
			return nil, []error{err}
		}
		return SortByOrder(order, data, getter), nil
	}
	return GenLoad(loader, &filed)
}

func GenLoad2DIn[Key LoadKey, Value any](db *gorm.DB, getter func(data *Value) Key, less func(a, b *Value) bool) Load2DFunc[Key, Value] {
	filed := getFiledNameByGetter(getter)
	loader := func(ctx context.Context, order []Key) ([][]*Value, []error) {
		db := gorms.GetDb(ctx, db)
		data := make([]*Value, 0, len(order))
		where := filed + " in ?"
		err := db.Where(where, order).Find(&data).Error
		if err != nil {
			return nil, []error{err}
		}
		return Sort2DByOrder(order, data, getter, less), nil
	}
	return GenLoad2D(loader, &filed)
}

func getFiledNameByGetter[Key LoadKey, Value any](getter func(data *Value) Key) string {
	value := new(Value)
	refValue := reflect.ValueOf(value)
	// 必须是结构体
	refValue = refValue.Elem()
	if refValue.Kind() != reflect.Struct {
		panic(fmt.Sprintf("loads %s is not struct", refValue.String()))
	}

	// 查找结构体的指定字段名
	refValueFieldNum := refValue.NumField()
	for i := 0; i < refValueFieldNum; i++ {
		isMatched := matchFiled(getter, value, refValue.Field(i))
		if isMatched {
			refType := reflect.TypeOf(value)
			refType = refType.Elem()
			return strcase.ToSnake(refType.Field(i).Name)
		}
	}
	panic(fmt.Sprintf("loads struct %s field is not found", refValue.String()))
}

func matchFiled[Key LoadKey, Value any](getter func(data *Value) Key, value *Value, refValueField reflect.Value) bool {
	valueField := refValueField.Interface()
	_, ok := valueField.(Key)
	if ok == false {
		return false
	}
	switch valueField.(type) {
	case int, int32, int64:
		return matchFiledInt(getter, value, refValueField)
	case string:
		return matchFiledString(getter, value, refValueField)
	default:
		return false
	}
}

func matchFiledInt[Key LoadKey, Value any](getter func(data *Value) Key, value *Value, refValueField reflect.Value) bool {
	temp := int64(100)
	refValueField.SetInt(temp)
	return (interface{})(getter(value)).(int64) == temp
}

func matchFiledString[Key LoadKey, Value any](getter func(data *Value) Key, value *Value, refValueField reflect.Value) bool {
	temp := "match"
	refValueField.SetString(temp)
	return (interface{})(getter(value)).(string) == temp
}
