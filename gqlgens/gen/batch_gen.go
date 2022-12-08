package gen

import (
	"github.com/iancoleman/strcase"
	"github.com/shiqiyue/gof/files"
	"github.com/shiqiyue/gof/loggers"
	"io/ioutil"
	"path/filepath"
	"reflect"
)

func BatchGen(dirname string, ms ...interface{}) {
	files.EnsureDirExist(dirname)

	for _, m := range ms {
		graphqlFileName := strcase.ToSnake(reflect.ValueOf(m).Elem().Type().Name()) + ".graphql"
		grapqhlFilePath := filepath.Join(dirname, graphqlFileName)
		if files.IsPathExists(grapqhlFilePath) {
			continue
		}
		r := Gen(m)

		err := ioutil.WriteFile(grapqhlFilePath, []byte(r), 0777)
		if err != nil {
			loggers.Error(nil, err.Error())
		}
	}
}
