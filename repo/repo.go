package repo

import "github.com/bobacgo/admin-service/repo/data"

type Repo struct {
	User *UserRepo
}

func NewRepo(data *data.Client) *Repo {
	return &Repo{
		User: NewUserRepo(data),
	}
}