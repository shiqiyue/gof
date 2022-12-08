package gqlgens

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/shiqiyue/gof/gqlgens/extension"
	"github.com/shiqiyue/gof/gqlgens/middle"
)

func CommonMiddle(srv *handler.Server) {
	srv.SetRecoverFunc(middle.DefaultRecover)
	//srv.Use(extension.Prome{})
	//srv.Use(extension.GqlgenTracer{})
	srv.Use(extension.Logger{IsPrintLog: false})

}
