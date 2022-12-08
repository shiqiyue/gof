package gen

import "github.com/vektah/gqlparser/v2/ast"

// 新建[item!]! type
func NewNotNullListAndItemNotNullType(typename string) *ast.Type {
	return ast.NonNullListType(ast.NonNullNamedType(typename, nil), nil)
}

// 新建 [item!] type
func NewListAndItemNotNullType(typename string) *ast.Type {
	return ast.ListType(ast.NonNullNamedType(typename, nil), nil)
}

func NewListType(typename string) *ast.Type {
	return ast.ListType(ast.NamedType(typename, nil), nil)
}

// 新建 item! type
func NewNotNullType(typename string) *ast.Type {
	return ast.NonNullNamedType(typename, nil)
}

// 新建 item type
func NewType(typename string) *ast.Type {
	return ast.NamedType(typename, nil)
}

// 新建 [item!]! argument
func NewNotNullListAndItemNotNullArgument(typename string, argumentname string, description string) *ast.ArgumentDefinition {
	return &ast.ArgumentDefinition{
		Description: description,
		Name:        argumentname,
		Type:        NewNotNullListAndItemNotNullType(typename),
	}
}

// 新建 [item!] argument
func NewListAndItemNotNullArgument(typename string, argumentname string, description string) *ast.ArgumentDefinition {
	return &ast.ArgumentDefinition{
		Description: description,
		Name:        argumentname,
		Type:        NewListAndItemNotNullType(typename),
	}
}

// 新建 [item] argument
func NewListArgument(typename string, argumentname string, description string) *ast.ArgumentDefinition {
	return &ast.ArgumentDefinition{
		Description: description,
		Name:        argumentname,
		Type:        NewListType(typename),
	}
}

// 新建 item! argument
func NewNotNullArgument(typename string, argumentname string, description string) *ast.ArgumentDefinition {
	return &ast.ArgumentDefinition{
		Description: description,
		Name:        argumentname,
		Type:        NewNotNullType(typename),
	}
}

// 新建 item argument
func NewArgument(typename string, argumentname string, description string) *ast.ArgumentDefinition {
	return &ast.ArgumentDefinition{
		Description: description,
		Name:        argumentname,
		Type:        NewType(typename),
	}
}
