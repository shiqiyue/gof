package gorms

import "strings"

func CleanSql(sql string) string {
	//sql = strings.ReplaceAll(sql, "\n", "")
	sql = strings.ReplaceAll(sql, "\r", "")
	sql = strings.ReplaceAll(sql, "\t", "")
	return sql
}
