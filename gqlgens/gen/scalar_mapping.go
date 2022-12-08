package gen

var ScalarMapping map[string]string

const (
	SCALAR_INT    = "Int"
	SCALAR_BIGINT = "Bigint"
	SCALAR_INT32  = "Int32"
	SCALAR_STRING = "String"
	SCALAR_FLOAT  = "Float"
	SCALAR_TIME   = "Time"

	SCALAR_BOOLEAN = "Boolean"

	SCALAR_ANY = "Any"
)

// 默认的gql类型
var DefaultGqlTypeNamed = SCALAR_STRING

func init() {
	ScalarMapping = make(map[string]string)
	ScalarMapping["int"] = SCALAR_INT
	ScalarMapping["int64"] = SCALAR_BIGINT
	ScalarMapping["int32"] = SCALAR_INT32
	ScalarMapping["bool"] = SCALAR_BOOLEAN
	ScalarMapping["time.Time"] = SCALAR_TIME
	ScalarMapping["string"] = SCALAR_STRING
	ScalarMapping["pq.StringArray"] = SCALAR_STRING
	ScalarMapping["pq.Int64Array"] = SCALAR_BIGINT
	ScalarMapping["datatypes.JSON"] = SCALAR_ANY

}

func GetScalarByType(t string) string {
	gqlFieldTypeNamed, ok := ScalarMapping[t]
	if !ok {
		return DefaultGqlTypeNamed
	}
	return gqlFieldTypeNamed
}
