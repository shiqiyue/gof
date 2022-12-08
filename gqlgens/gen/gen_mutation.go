package gen

import "github.com/vektah/gqlparser/v2/ast"

func (c *GenContext) genMuation() {
	def := &ast.Definition{
		Kind: ast.Object,
		Name: "Mutation",
	}
	def.Fields = make([]*ast.FieldDefinition, 0)
	// 添加
	addArguments := []*ast.ArgumentDefinition{&ast.ArgumentDefinition{
		Name: "req",
		Type: ast.NonNullNamedType(c.addReqName(), nil),
	}}
	addMutation := &ast.FieldDefinition{
		Description: "添加记录",
		Name:        "add_" + c.modelSneakName(),
		Arguments:   addArguments,
		Type:        NewNotNullType(SCALAR_BOOLEAN),
	}
	def.Fields = append(def.Fields, addMutation)
	// 修改
	editArguments := []*ast.ArgumentDefinition{&ast.ArgumentDefinition{
		Name: "req",
		Type: ast.NonNullNamedType(c.editReqName(), nil),
	}}
	editMutation := &ast.FieldDefinition{
		Description: "修改记录",
		Name:        "edit_" + c.modelSneakName(),
		Arguments:   editArguments,
		Type:        NewNotNullType(SCALAR_BOOLEAN),
	}
	def.Fields = append(def.Fields, editMutation)
	// 删除
	removeArguments := []*ast.ArgumentDefinition{&ast.ArgumentDefinition{
		Name: "id",
		Type: ast.NonNullNamedType("Int", nil),
	}}
	removeMutation := &ast.FieldDefinition{
		Description: "删除记录",
		Name:        "remove_" + c.modelSneakName(),
		Arguments:   removeArguments,
		Type:        NewNotNullType(SCALAR_BOOLEAN),
	}
	def.Fields = append(def.Fields, removeMutation)

	c.SchemaDocument.Extensions = append(c.SchemaDocument.Extensions, def)
}
