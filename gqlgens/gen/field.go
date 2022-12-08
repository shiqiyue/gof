package gen

import (
	"github.com/iancoleman/strcase"
	"strings"
)

type FieldInfo struct {
	// 名称
	Name string
	// 类型
	Type string
	// 是否可以为空
	Nullable bool
	// tag
	Tag string
	// Gorm Tag
	GormTag string
}

func (i FieldInfo) GqlName() string {
	return strcase.ToLowerCamel(i.Name)
}

func (i FieldInfo) DBName() string {
	return strcase.ToSnake(i.Name)
}

func (i FieldInfo) IsDetail() bool {
	return strings.Contains(i.Tag, "detail")
}

func (i FieldInfo) IsAdd() bool {
	return strings.Contains(i.Tag, "add")
}

func (i FieldInfo) IsEdit() bool {
	return strings.Contains(i.Tag, "edit")
}

func (i FieldInfo) IsList() bool {
	return strings.Contains(i.Tag, "list")
}

func (i FieldInfo) IsWhere() bool {
	return strings.Contains(i.Tag, "where")
}

func (i FieldInfo) IsOrder() bool {
	return strings.Contains(i.Tag, "order")

}

func (i FieldInfo) Scalar() string {
	return GetScalarByType(i.Type)
}

func (i FieldInfo) IsArray() bool {
	if strings.Contains(i.Type, "Array") {
		return true
	}
	return false
}

func (i FieldInfo) Description() string {
	commentIndex := strings.LastIndex(i.GormTag, "comment:")
	if commentIndex >= 0 {
		endIndex := len(i.GormTag) - 1
		for j := endIndex; j > commentIndex; j-- {
			if i.GormTag[j] == ';' {
				endIndex = j
				break
			}
		}
		return i.GormTag[commentIndex+8 : endIndex+1]
	}
	return ""
}
