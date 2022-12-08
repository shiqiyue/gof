package directive

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/shiqiyue/gof/validates"
)

func Validate(ctx context.Context, obj interface{}, next graphql.Resolver, rules string, name *string, msg *string) (res interface{}, err error) {

	val, err := next(ctx)
	if err != nil {
		return val, err
	}

	err = validates.ValidateVar(val, rules)
	if err != nil {
		return nil, errorAddName(err, name, msg)
	}
	return val, nil
}

func errorAddName(err error, name *string, msg *string) error {
	if msg != nil {
		return errors.New(*msg)
	}
	if name == nil {
		return err
	}
	return errors.New(*name + err.Error())
}
