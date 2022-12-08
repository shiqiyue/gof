package slices

// 交集
func Join[V comparable](col1 []V, col2 []V) []V {
	if len(col1) == 0 || len(col2) == 0 {
		return make([]V, 0)
	}
	rs := make([]V, 0)
	for _, col1Item := range col1 {
		for _, col2Item := range col2 {
			if col1Item == col2Item {
				rs = append(rs, col1Item)
			}
		}
	}
	return rs
}
