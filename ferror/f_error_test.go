package ferror

import (
	errors2 "errors"
	"fmt"
	"github.com/pkg/errors"
	"github.com/shiqiyue/gof/resultcode"
	"testing"
)

func TestWrap(t *testing.T) {
	err1 := errors.New("err1")
	err2 := errors.Wrap(err1, "err2")
	fmt.Println(err2)
}

func TestWrapWithCode(t *testing.T) {
	err := errors.New("err1")
	err2 := Wrap("err2", err)
	err3 := WrapWithCode("err3", resultcode.FAIL, err2)
	fmt.Println(err3)
}

func TestWrapCode(t *testing.T) {
	err := errors2.New("net error")
	err2 := Wrap("err2", err)
	err3 := WrapCode(resultcode.FAIL, err2)
	fmt.Println(err3)
}
