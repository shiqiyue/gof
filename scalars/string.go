package scalars

func StringIfNilDefault(str *string, def string) string {
	if str == nil {
		return def
	}
	return *str
}

func GetStringValue(str *string) string {
	return *str
}

// 返回String的引用
func GetStringRef(i string) *string {
	return &i
}
