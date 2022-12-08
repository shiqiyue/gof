package resultcode

// 数据层错误，01+三位数
const (
	// 数据层-sql错误
	DATA_SQL_ERROR = "01000"
	// 数据层-查询异常
	DATA_QUERY_ERROR = "01001"
	// 数据层-数据不存在
	DATA_DATA_NOT_FOUND = "01002"
	// 数据层-更新异常
	DATA_UPDATE_ERROR = "01003"
	// 数据层-删除异常
	DATA_DELETE_ERROR = "01004"
	// 数据层-事务异常
	DATA_TX_ERROR = "01005"
	// 数据层-操作超时
	DATA_TIMEOUT = "01006"

	// 数据层-参数验证错误
	DATA_PARAM_UNVALID = "01997"
	// 数据层-参数不足
	DATA_PARAM_REQUIRED = "01998"
	// 数据层-其他异常
	DATA_OTHER_ERROR = "01999"
)
