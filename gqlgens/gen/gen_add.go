package gen

import "github.com/vektah/gqlparser/v2/ast"

func (c *GenContext) addReqName() string {
	return c.modelName() + "AddReq"
}

func (c *GenContext) genAddReq() {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.addReqName()
	def.Directives = []*ast.Directive{c.modelDirective()}
	def.Fields = make([]*ast.FieldDefinition, 0)
	for _, field := range c.Fields {
		if !field.IsAdd() {
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
