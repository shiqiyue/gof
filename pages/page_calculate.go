package pages

// 分页参数转化成数据库limit,offset
func ToDbQuery(currentPage, pageSize int) (limit, offset int) {
	if currentPage < 1 {
		currentPage = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	if currentPage == 1 {
		offset = 0
	} else {
		offset = (currentPage - 1) * pageSize
	}
	limit = pageSize
	return
}
