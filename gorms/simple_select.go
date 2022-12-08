package gorms

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 简单查询
type SimpleSelect struct {
	// 查询项
	Where *SimpleWhere
	// 排序
	OrderBy *clause.OrderByColumn
	// offset
	Offset *int
	// limit
	Limit *int
}

// 执行查询
func (s *SimpleSelect) DoSelect(db *gorm.DB) *gorm.DB {
	if s.Where != nil {
		db = s.Where.DoWhere(db)
	}
	if s.OrderBy != nil {
		db = db.Order(*s.OrderBy)
	}
	if s.Offset != nil {
		db = db.Offset(*s.Offset)
	}
	if s.Limit != nil {
		db = db.Limit(*s.Limit)
	}
	return db
}
