package sqls

import (
	"gitee.com/shiqiyue/postgresql-parser/pkg/sql/parser"
	"gitee.com/shiqiyue/postgresql-parser/pkg/sql/sem/tree"
)

// 解析sql，获取select
func ParseSelectColumns(sql string) ([]string, error) {

	parse, err := parser.Parse(sql)
	if err != nil {
		return nil, err
	}
	var columns = []string{}
	for _, statement := range parse {
		t, ok := statement.AST.(*tree.Select)
		if !ok {
			continue
		}
		sc, ok := t.Select.(*tree.SelectClause)
		if !ok {
			continue
		}
		for _, expr := range sc.Exprs {
			if expr.As != "" {
				columns = append(columns, string(expr.As))
				continue
			}
			unresolvedName, ok := expr.Expr.(*tree.UnresolvedName)
			if ok {
				if unresolvedName.Parts[0] != "" {
					columns = append(columns, string(unresolvedName.Parts[0]))
				}
			}
		}

	}
	return columns, nil
}
