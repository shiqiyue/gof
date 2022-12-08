package sqls

import (
	"context"
	"errors"
	"github.com/dgryski/trifles/uuid"
	"go.uber.org/zap/buffer"
	"sync"
	"text/template"
)

var sqlCache = &sync.Map{}

type RawSql struct {
	Id  string
	Sql string
}

type cacheItem struct {
	sqlReport *RawSql
	t         *template.Template
}

// 执行sql模板
func ExecuteSqlTemplate(ctx context.Context, sqlReport *RawSql, params map[string]interface{}) (string, error) {
	return executeTemplate(ctx, sqlReport.Sql, sqlReport, params, sqlCache)
}

func executeTemplate(ctx context.Context, sql string, sqlReport *RawSql, params map[string]interface{}, cache *sync.Map) (string, error) {
	t := getFromCache(ctx, sqlReport, cache)
	if t == nil {
		t2, err := template.New(uuid.UUIDv4()).Parse(sql)
		if err != nil {
			return "", err
		}
		if t2 == nil {
			return "", errors.New("template create fail")
		}
		t = t2
		setToCache(ctx, sqlReport, t, cache)
	}
	bs := buffer.Buffer{}
	err := t.Execute(&bs, params)
	if err != nil {
		return "", err
	}
	return bs.String(), nil
}

func getFromCache(ctx context.Context, sqlReport *RawSql, cache *sync.Map) *template.Template {
	cacheValue, ok := cache.Load(sqlReport.Id)
	if ok {
		t, tOk := cacheValue.(*cacheItem)
		if tOk {
			return t.t
		}
	}
	return nil
}

func setToCache(ctx context.Context, sqlReport *RawSql, t *template.Template, cache *sync.Map) {
	cache.Store(sqlReport.Id, &cacheItem{
		sqlReport: sqlReport,
		t:         t,
	})

}
