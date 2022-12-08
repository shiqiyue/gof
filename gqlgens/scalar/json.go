package scalar

import (
	"encoding/json"
	"fmt"
	"io"
)

type JSONScalar json.RawMessage

//
func (y *JSONScalar) UnmarshalGQL(v interface{}) error {
	data, ok := v.(string)
	if !ok {
		return fmt.Errorf("Scalar must be a string ")
	}

	*y = []byte(data)
	return nil
}

//
func (y JSONScalar) MarshalGQL(w io.Writer) {
	_, _ = w.Write(y)
}
