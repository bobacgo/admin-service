package service

import "github.com/bobacgo/admin-service/repo"

type Service struct {
	User *UserService
	I18n *I18nService
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		User: NewUserService(repo),
		I18n: NewI18nService(repo),
	}
}
