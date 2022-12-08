package gen

import (
	"bytes"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
	"reflect"
)

func Gen(m interface{}) string {
	t := reflect.ValueOf(m).Elem().Type()
	context := &GenContext{
		SchemaDocument: &ast.SchemaDocument{},
		T:              t,
	}

	context.resolveType(t)
	context.genSchema()

	var buf bytes.Buffer
	f := formatter.NewFormatter(&buf)
	f.FormatSchemaDocument(context.SchemaDocument)
	return buf.String()
}

func GenType(t reflect.Type) (def *ast.Definition) {

	return nil
}

func GenAddInput(t reflect.Type) (def *ast.Definition) {

	return nil
}

func GenEditInput(t reflect.Type) (def *ast.Definition) {

	return nil
}

func GenRemoveInput(t reflect.Type) (def *ast.Definition) {

	return nil
}

func GenAddMuation(t reflect.Type) (def *ast.Definition) {
	return nil
}

func GenEditMutation(t reflect.Type) (def *ast.Definition) {
	return nil
}

func GenRemoveMutation(t reflect.Type) (def *ast.Definition) {
	return nil
}
