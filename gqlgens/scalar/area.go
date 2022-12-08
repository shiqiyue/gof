package scalar

import (
	"encoding/json"
	"github.com/99designs/gqlgen/graphql"
	"github.com/pkg/errors"
	"io"
)

type Area struct {
	Province string
	City     string
	County   string
}

func MarshalArea(area Area) graphql.Marshaler {

	return graphql.WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(area)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalArea(v interface{}) (Area, error) {
	area, ok := v.(Area)
	if !ok {
		return Area{}, errors.New("should be area struct")
	}
	return area, nil
}
