package caches

import (
	"fmt"
	"github.com/shiqiyue/gof/asserts"
	"go/parser"
	"go/token"
)

func Gen() {
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "D:\\project\\go-admin\\internal\\module\\sys\\service\\bg_user.go", nil, parser.ParseComments)
	asserts.Nil(err, err)
	fmt.Printf("%v\n", astFile)
	fmt.Printf("%v\n", fileSet)
}
