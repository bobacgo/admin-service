package service

import (
	"context"

	"github.com/bobacgo/admin-service/apps/repo"
	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/bobacgo/admin-service/apps/repo/model"
)

type RoleService struct {
	repo *repo.Repo
}

func NewRoleService(repo *repo.Repo) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) Create(ctx context.Context, req *dto.RoleCreateReq) error {
	return s.repo.Role.Create(ctx, &model.Role{
		Code:        req.Code,
		Description: req.Description,
		Status:      req.Status,
	})
}

func (s *RoleService) Get(ctx context.Context, req *dto.GetRoleReq) (*model.Role, error) {
	return s.repo.Role.FindOne(ctx, req)
}

func (s *RoleService) Update(ctx context.Context, req *dto.RoleUpdateReq) error {
	return s.repo.Role.Update(ctx, &model.Role{
		Model:       model.Model{ID: req.ID},
		Code:        req.Code,
		Description: req.Description,
		Status:      req.Status,
	})
}

func (s *RoleService) List(ctx context.Context, req *dto.RoleListReq) ([]*model.Role, int64, error) {
	return s.repo.Role.Find(ctx, req)
}

func (s *RoleService) Delete(ctx context.Context, ids string) error {
	return s.repo.Role.Delete(ctx, ids)
}
