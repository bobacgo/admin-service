package user

import (
	"context"
)

type UserService struct {
	repo *UserRepo
}

func NewUserService(r *UserRepo) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(ctx context.Context, user *User) error {
	return s.repo.Create(ctx, user)
}

func (s *UserService) Get(ctx context.Context, req *GetUserReq) (*User, error) {
	return s.repo.FindOne(ctx, req)
}

func (s *UserService) Update(ctx context.Context, user *User) error {
	return s.repo.Update(ctx, user)
}

func (s *UserService) List(ctx context.Context, req *UserListReq) ([]*User, int64, error) {
	rows, total, err := s.repo.Find(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	for _, row := range rows {
		row.Password = ""
	}
	return rows, total, nil
}

func (s *UserService) Delete(ctx context.Context, ids string) error {
	return s.repo.Delete(ctx, ids)
}
