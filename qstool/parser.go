package qstool

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

const (
	tagSql    = "qssql"
	tagFormat = "qsformat"
)

func ParseWhere(where interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return parseWhere(db, where)
	}
}

func ParseOrder(key string, reverse bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return parseOrder(db, key, reverse)
	}
}

func ParsePage(pageIndex, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return parsePage(db, pageIndex, pageSize)
	}
}

func parseWhere(db *gorm.DB, where interface{}) *gorm.DB {
	if where == nil {
		return db
	}
	rValue := reflect.ValueOf(where)
	if rValue.Kind() == reflect.Ptr {
		if rValue.IsNil() {
			return db
		}
		rValue = rValue.Elem()
	}
	if rValue.Kind() != reflect.Struct {
		panic("ParseWhere: 入参的元素必须是结构体")
	}
	rType := rValue.Type()
	rValueNum := rValue.NumField()
	for i := 0; i < rValueNum; i++ {
		itemValue := rValue.Field(i)
		if itemValue.IsZero() {
			continue
		}
		if itemValue.Kind() == reflect.Ptr {
			itemValue = itemValue.Elem()
		}
		itemTag := rType.Field(i).Tag
		itemTagSql := itemTag.Get(tagSql)
		itemTagFormat := itemTag.Get(tagFormat)
		itemInterface := itemValue.Interface()
		if itemTagFormat != "" {
			itemInterface = fmt.Sprintf(itemTagFormat, itemInterface)
		}
		db = db.Where(itemTagSql, itemInterface)
	}
	return db
}

func parseOrder(db *gorm.DB, key string, reverse bool) *gorm.DB {
	key = strings.ToLower(key)
	if reverse {
		db = db.Order(fmt.Sprintf("%s ASC", key)) // 反向
	} else {
		db = db.Order(fmt.Sprintf("%s DESC", key)) // 正向
	}
	return db
}

func parsePage(db *gorm.DB, pageIndex, pageSize int) *gorm.DB {
	if pageIndex < 1 {
		pageIndex = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	limit := pageSize
	offset := (pageIndex - 1) * pageSize
	return db.Limit(limit).Offset(offset)
}
