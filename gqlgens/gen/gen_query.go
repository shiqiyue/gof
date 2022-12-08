package gen

import "github.com/vektah/gqlparser/v2/ast"

func (c *GenContext) genQuery() {
	def := &ast.Definition{
		Kind: ast.Object,
		Name: "Query",
	}
	def.Fields = make([]*ast.FieldDefinition, 0)
	// 列表查询
	listQueryArguments := []*ast.ArgumentDefinition{
		NewListAndItemNotNullArgument(c.columnEnumName(), "distinct_on", ""),
		NewArgument(SCALAR_INT, "limit", ""),
		NewArgument(SCALAR_INT, "offset", ""),
		NewListAndItemNotNullArgument(c.orderByName(), "order_by", ""),
		NewArgument(c.boolExpName(), "where", ""),
	}
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Description: "列表查询",
		Name:        c.modelSneakName() + "s",
		Arguments:   listQueryArguments,
		Type:        ast.NonNullListType(ast.NonNullNamedType(c.modelName(), nil), nil),
		Directives:  nil,
		Position:    nil,
	})

	// 数量查询
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Description: "数量查询",
		Name:        c.modelSneakName() + "_count",
		Arguments:   listQueryArguments,
		Type:        NewNotNullType(SCALAR_INT),
	})

	// 主键查询
	pkArguments := []*ast.ArgumentDefinition{
		NewNotNullArgument(SCALAR_INT, "id", ""),
	}
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Description: "根据ID查询",
		Name:        c.modelSneakName(),
		Arguments:   pkArguments,
		Type:        NewNotNullType(c.modelName()),
	})

	c.SchemaDocument.Extensions = append(c.SchemaDocument.Extensions, def)

}
