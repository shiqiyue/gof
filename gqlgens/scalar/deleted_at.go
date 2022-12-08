package scalar

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm"
	"io"
	"strconv"
	"time"
)

func MarshalDeletedAt(t gorm.DeletedAt) graphql.Marshaler {
	if !t.Valid {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Time.Format(time.RFC3339)))
	})
}

func UnmarshalDeletedAt(v interface{}) (gorm.DeletedAt, error) {
	if tmpStr, ok := v.(string); ok {
		t, err := time.Parse(time.RFC3339, tmpStr)
		if err != nil {
			return gorm.DeletedAt{}, errors.New("time should be RFC3339 formatted string")
		}
		return gorm.DeletedAt{Time: t, Valid: true}, nil
	}
	return gorm.DeletedAt{}, errors.New("time should be RFC3339 formatted string")
}
