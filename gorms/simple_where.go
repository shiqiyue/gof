package gorms

import (
	"context"
	"gorm.io/gorm"
)

type Operate int

const (
	// 相等
	OPERATE_EQ = iota + 1

	// in
	OPERATE_IN

	// like
	OPERATE_LIKE

	// >
	OPERATE_GT

	// >=
	OPERATE_GTE

	// <
	OPERATE_LT

	// <=
	OPERATE_LTE

	// !=
	OPERATE_NE
)

// 简单查询
type SimpleWhere struct {
	// 查询项
	Items []SimpleWhereItem
}

// 简单查询-项
type SimpleWhereItem struct {
	// 列名
	Column string
	// 操作
	Operate Operate
	// 值
	Value interface{}
}

// 权限过滤数据
func (w *SimpleWhere) DoAuthFilter(ctx context.Context, db *gorm.DB) *gorm.DB {
	/*	user := auths.GetUser(ctx)
		switch user.DataAccess {
		case int(sys_enum.DATA_ACCESS_PROVINCE):
			return db.Where("auth_province_id = ?", user.ProvinceId)
		case int(sys_enum.DATA_ACCESS_CITY):
			return db.Where("auth_city_id = ?", user.CityId)
		case int(sys_enum.DATA_ACCESS_AREA):
			return db.Where("auth_district_id = ?", user.AreaId)
		case int(sys_enum.DATA_ACCESS_ORG):
			return db.Where("auth_organization_id = ?", user.OrganizationId)
		case int(sys_enum.DATA_ACCESS_DEPT):
			return db.Where("auth_department_id = ?", user.DepartmentId)
		case int(sys_enum.DATA_ACCESS_USER):
			return db.Where("created_by = ?", user.UserId)
		}*/
	return db
}

// 执行过滤
func (w *SimpleWhere) DoWhere(db *gorm.DB) *gorm.DB {
	if len(w.Items) == 0 {
		return db
	}
	for _, item := range w.Items {
		switch item.Operate {
		case OPERATE_EQ:
			db = db.Where(item.Column+" = ?", item.Value)
		case OPERATE_IN:
			db = db.Where(item.Column+" in ?", item.Value)
		case OPERATE_LIKE:
			db = db.Where(item.Column+" like ?", item.Value)
		case OPERATE_GT:
			db = db.Where(item.Column+" > ?", item.Value)
		case OPERATE_GTE:
			db = db.Where(item.Column+" >= ?", item.Value)
		case OPERATE_LT:
			db = db.Where(item.Column+" < ?", item.Value)
		case OPERATE_LTE:
			db = db.Where(item.Column+" <= ?", item.Value)
		case OPERATE_NE:
			db = db.Where(item.Column+" != ?", item.Value)
		}
	}
	return db
}
