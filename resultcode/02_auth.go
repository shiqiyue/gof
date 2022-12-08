package resultcode

// 权限错误, 02+三位数
const (
	// 需要认证
	AUTH_NEED_LOGIN = "02000"
	// 权限不足
	AUTH_DENIED = "02001"

	// 权限-参数验证错误
	AUTH_PARAM_UNVALID = "02997"
	// 权限-参数不足
	AUTH_PARAM_REQUIRED = "02998"
	// 权限-其他异常
	AUTH_OTHER_ERROR = "02999"
)
