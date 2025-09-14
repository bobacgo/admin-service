package api

// 10 服务 | 03：03号模块 | 25：错误码为25
var (
	// 服务 10
	// 通用模块 00
	OK            = 0
	ErrCodeParam  = 100001
	ErrCodeServer = 100002
	ErrCodeAuth   = 100003
	ErrCodeNoAuth = 100004
	ErrCodeNoPerm = 100005
	ErrCodeNoData = 100006

	// 业务模块
	// User 01
	ErrCodeUsernameOrPassword = 100100
	ErrCodeUserDisabled       = 100101
)
