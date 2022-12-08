package gen

import "github.com/vektah/gqlparser/v2/ast"

func (c *GenContext) editReqName() string {
	return c.modelName() + "EditReq"
}

func (c *GenContext) genEditReq() {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.editReqName()
	def.Directives = []*ast.Directive{c.modelDirective()}
	def.Fields = make([]*ast.FieldDefinition, 0)
	for _, field := range c.Fields {
		if !field.IsEdit() {
			continue
		}
		def.Fields = append(def.Fields, &ast.FieldDefinition{
			Name: field.GqlName(),
			Type: &ast.Type{
				NamedType: field.Scalar(),
				NonNull:   !field.Nullable,
			},
		})
	}
	c.SchemaDocument.Definitions = append(c.SchemaDocument.Definitions, def)
}
