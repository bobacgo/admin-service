package ecode

// Application-level error codes centrally defined here so multiple packages can import without cycles.
var (
	OK            = 0
	ErrCodeParam  = 100001
	ErrCodeServer = 100002
	ErrCodeAuth   = 100003
	ErrCodeNoAuth = 100004
	ErrCodeNoPerm = 100005
	ErrCodeNoData = 100006

	// User-specific
	ErrCodeUsernameOrPassword = 100100
	ErrCodeUserDisabled       = 100101
)
