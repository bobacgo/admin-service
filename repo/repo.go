package repo

import "github.com/bobacgo/admin-service/repo/data"

type Repo struct {
	User *UserRepo
	I18n *I18nRepo
}

func NewRepo(data *data.Client) *Repo {
	return &Repo{
		User: NewUserRepo(data),
		I18n: NewI18nRepo(data),
	}
}
