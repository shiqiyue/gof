package gen

import "github.com/vektah/gqlparser/v2/ast"

func (c *GenContext) orderByName() string {
	return c.modelName() + "OrderBy"
}

func (c *GenContext) genOrderBy() {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.orderByName()
	for _, field := range c.Fields {
		if !field.IsOrder() {
			continue
		}
		def.Fields = append(def.Fields, &ast.FieldDefinition{
			Name: field.GqlName(),
			Type: NewType(field.Scalar()),
		})
	}
	c.SchemaDocument.Definitions = append(c.SchemaDocument.Definitions, def)

}
