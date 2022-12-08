package sqls

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/shiqiyue/gof/ferror"
	"github.com/shiqiyue/gof/gorms"
	"github.com/shiqiyue/gof/loggers"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type PageParam struct {
	ListSql     *RawSql
	CountSql    *RawSql
	M           map[string]interface{}
	CurrentPage int
	PageSize    int
}

// 列表查询
func List(ctx context.Context, db *gorm.DB, rawSql *RawSql, m map[string]interface{}, rs interface{}) error {
	sql, err := ExecuteSqlTemplate(ctx, rawSql, m)
	if err != nil {
		return err
	}
	resultMap := make([]map[string]interface{}, 0)
	err, _ = executeSql(ctx, db, sql, m, &resultMap)
	if err != nil {
		return err
	}
	err = mapstructure.Decode(&resultMap, rs)
	if err != nil {
		return err
	}
	return nil
}

// 数量查询
func Count(ctx context.Context, db *gorm.DB, rawSql *RawSql, m map[string]interface{}, c *int) error {
	sql, err := ExecuteSqlTemplate(ctx, rawSql, m)
	if err != nil {
		return err
	}
	err, _ = executeSql(ctx, db, sql, m, c)
	if err != nil {
		return err
	}

	return nil
}

// 分页查询
func Page(ctx context.Context, db *gorm.DB, pageParam PageParam, c *int, rs interface{}) error {
	err := Count(ctx, db, pageParam.CountSql, pageParam.M, c)
	if err != nil {
		return err
	}
	limit, offset := gorms.GetPage(pageParam.CurrentPage, pageParam.PageSize)
	pageParam.M["limit"] = limit
	pageParam.M["offset"] = offset
	err = List(ctx, db, pageParam.ListSql, pageParam.M, rs)
	if err != nil {
		return err
	}
	return nil
}

func GetSql(d *gorm.DB) string {
	sql := d.Statement.SQL.String()
	vars := d.Statement.Vars
	return d.Dialector.Explain(sql, vars...)
}

func CleanSql(sql string) string {
	sql = strings.ReplaceAll(sql, "\r", "")
	sql = strings.ReplaceAll(sql, "\t", "")
	return sql
}

// 执行sql
func executeSql(ctx context.Context, conn *gorm.DB, sql string, queryParam map[string]interface{}, target interface{}) (err error, rawsql string) {
	sql = CleanSql(sql)
	if len(queryParam) > 0 {
		loggers.Debug(ctx, "执行报表SQL:"+sql, zap.Any("参数", queryParam), zap.String("SQL", sql))
		d := conn.Raw(sql, queryParam)
		rawsql = GetSql(d)
		d = conn.Raw(rawsql)
		d.Scan(target)
		err = d.Error
		if err != nil {
			err = ferror.Wrap("执行SQL异常,错误的SQL为"+rawsql, err)
		}
		return
	}
	d := conn.Raw(sql)
	rawsql = GetSql(d)
	d = conn.Raw(rawsql)
	d.Scan(target)
	err = d.Error
	if err != nil {
		err = ferror.Wrap("执行SQL异常,错误的SQL为"+rawsql, err)

	}
	return

}
