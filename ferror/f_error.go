package ferror

import (
	"github.com/shiqiyue/gof/errors"
)

// Wrap error, include current caller
func Wrap(msg string, subError error) error {
	if subError == nil {
		return nil
	}
	return errors.Wrap(subError, msg)
}

func WrapWithCode(msg string, code string, subError error) error {
	if subError == nil {
		return NewWithCode(code, msg)
	}
	return Wrap(code+":"+msg, subError)
}

func WrapCode(code string, subError error) error {
	if subError == nil {
		panic("sub error can not nil")
	}

	return Wrap(code+":"+cause(subError).Error(), subError)
}

func cause(err error) error {
	type causer interface {
		Cause() error
	}
	preErr := err
	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		preErr = err

		err = cause.Cause()
	}
	return preErr
}

func New(msg string) error {
	return errors.New(msg)
}

func NewWithCode(code, msg string) error {
	if code == "" {
		return errors.New(msg)
	}
	return errors.New(code + ":" + msg)
}
