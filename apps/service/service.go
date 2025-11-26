package service

import (
	"github.com/bobacgo/admin-service/apps/i18n"
	"github.com/bobacgo/admin-service/apps/menu"
	"github.com/bobacgo/admin-service/apps/role"
	"github.com/bobacgo/admin-service/apps/user"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	Validator *validator.Validate
	User      *user.UserService
	I18n      *i18n.I18nService
	Menu      *menu.MenuService
	Role      *role.RoleService
}

func GetValidator() *validator.Validate {
	return validator.New()
}
