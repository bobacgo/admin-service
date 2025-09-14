package service

import (
	"github.com/bobacgo/admin-service/apps/repo"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	Validator *validator.Validate
	User      *UserService
	I18n      *I18nService
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		Validator: validator.New(),
		User:      NewUserService(repo),
		I18n:      NewI18nService(repo),
	}
}
