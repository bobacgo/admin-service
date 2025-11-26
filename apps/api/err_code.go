package api

import "github.com/bobacgo/admin-service/apps/ecode"

// Re-export ecode constants for existing code that imports apps/api
var (
	OK                        = ecode.OK
	ErrCodeParam              = ecode.ErrCodeParam
	ErrCodeServer             = ecode.ErrCodeServer
	ErrCodeAuth               = ecode.ErrCodeAuth
	ErrCodeNoAuth             = ecode.ErrCodeNoAuth
	ErrCodeNoPerm             = ecode.ErrCodeNoPerm
	ErrCodeNoData             = ecode.ErrCodeNoData
	ErrCodeUsernameOrPassword = ecode.ErrCodeUsernameOrPassword
	ErrCodeUserDisabled       = ecode.ErrCodeUserDisabled
)
