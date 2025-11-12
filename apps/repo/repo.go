package repo

import (
	"github.com/bobacgo/admin-service/apps/repo/data"
)

type Repo struct {
	User *UserRepo
	I18n *I18nRepo
	Menu *MenuRepo
	Role *RoleRepo
}

func NewRepo(data *data.Client) *Repo {
	return &Repo{
		User: NewUserRepo(data),
		I18n: NewI18nRepo(data),
		Menu: NewMenuRepo(data),
		Role: NewRoleRepo(data),
	}
}
