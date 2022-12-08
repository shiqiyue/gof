package scalars

// 返回int的引用
func GetIntRef(i int) *int {
	return &i
}

func GetIntValue(i *int) int {
	return *i
}

func IntIfNilDefault(i *int, def int) int {
	if i == nil {
		return def
	}
	return *i
}

func Int32IfNilDefault(i *int32, def int32) int32 {
	if i == nil {
		return def
	}
	return *i
}

func Int64IfNilDefault(i *int64, def int64) int64 {
	if i == nil {
		return def
	}
	return *i
}
