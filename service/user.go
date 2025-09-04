package service

import (
	"context"

	"github.com/bobacgo/admin-service/repo"
	"github.com/bobacgo/admin-service/repo/dto"
	"github.com/bobacgo/admin-service/repo/model"
)

type UserService struct {
	repo *repo.Repo
}

func NewUserService(repo *repo.Repo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, user *model.User) error {
	return s.repo.User.Create(ctx, user)
}

func (s *UserService) Get(ctx context.Context, req *dto.GetUserReq) (*model.User, error) {
	return s.repo.User.FindOne(ctx, req)
}

func (s *UserService) Update(ctx context.Context, user *model.User) error {
	return s.repo.User.Update(ctx, user)
}

func (s *UserService) List(ctx context.Context, req *dto.UserListReq) ([]*model.User, int64, error) {
	rows, total, err := s.repo.User.Find(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	for _, row := range rows {
		row.Password = ""
	}

	return rows, total, nil
}

func (s *UserService) Delete(ctx context.Context, ids string) error {
	return s.repo.User.Delete(ctx, ids)
}

// 其他业务
