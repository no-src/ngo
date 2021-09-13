package _contract

// Page 分页数据包装器
type Page struct {
	// 当前页码，从0开始
	PageIndex int
	// 分页总数
	TotalPages int
	// 数据总数
	TotalItems int
	// 页大小
	PageSize int
	// 当前页数据
	Items interface{}
	// 其他用户数据
	Context interface{}
}

// GetPage 分页计算，如果超出索引则统一返回0,0
// total 数据源总数
// pageIndex 分页索引，从0开始
// pageSize 分页大小
// startIndex 分页开始索引,索引从0开始
// endIndex 分页结束索引，结束不包含endIndex
func GetPage(total, pageIndex, pageSize int64) (startIndex, endIndex int64) {
	if total <= 0 || pageIndex < 0 || pageSize <= 0 {
		return startIndex, endIndex
	}
	startIndex = pageIndex * pageSize
	endIndex = startIndex + pageSize

	//检查是否溢出
	if startIndex >= total {
		startIndex = 0
		endIndex = 0
		return startIndex, endIndex
	}
	if endIndex > total {
		endIndex = total
	}
	return startIndex, endIndex
}
