package service

import "github.com/bobacgo/admin-service/repo"

type Service struct {
	User *UserService
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		User: NewUserService(repo),
	}
}