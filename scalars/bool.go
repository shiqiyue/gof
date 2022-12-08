package scalars

// 获取true的引用
func GetTrueRef() *bool {
	t := true
	return &t
}

// 获取false的引用
func GetFalseRef() *bool {
	f := false
	return &f
}

func BoolIfNilDefault(boolean *bool, def bool) bool {
	if boolean == nil {
		return def
	}
	return *boolean
}

// 获取bool值的引用
func GetBoolRef(b bool) *bool {
	return &b
}
