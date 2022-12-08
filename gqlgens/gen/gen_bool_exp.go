package gen

import "github.com/vektah/gqlparser/v2/ast"

func (c *GenContext) boolExpName() string {
	return c.modelName() + "BoolExp"
}

func (c *GenContext) genBoolExp() {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.boolExpName()
	def.Fields = make([]*ast.FieldDefinition, 0)
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:       "_and",
		Type:       NewListType(c.boolExpName()),
		Directives: nil,
		Position:   nil,
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:       "_not",
		Type:       NewType(c.boolExpName()),
		Directives: nil,
		Position:   nil,
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:       "_or",
		Type:       NewListType(c.boolExpName()),
		Directives: nil,
		Position:   nil,
	})
	for _, field := range c.Fields {
		if !field.IsWhere() {
			continue
		}
		def.Fields = append(def.Fields, &ast.FieldDefinition{
			Name: field.GqlName(),
			Type: NewType(field.Scalar() + "ComparisonExp"),
		})
	}
	c.SchemaDocument.Definitions = append(c.SchemaDocument.Definitions, def)
}
