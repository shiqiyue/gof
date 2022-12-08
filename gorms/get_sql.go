package gorms

import (
	"gorm.io/gorm"
)

func GetSql(d *gorm.DB) string {
	sql := d.Statement.SQL.String()
	vars := d.Statement.Vars
	return d.Dialector.Explain(sql, vars...)
}
