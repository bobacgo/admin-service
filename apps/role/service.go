package role

import (
	"context"

	"github.com/bobacgo/admin-service/apps/repo/model"
)

type RoleService struct {
	repo *RoleRepo
}

func NewRoleService(r *RoleRepo) *RoleService {
	return &RoleService{repo: r}
}

func (s *RoleService) Create(ctx context.Context, req *RoleCreateReq) error {
	return s.repo.Create(ctx, &Role{
		Code:        req.Code,
		Description: req.Description,
		Status:      req.Status,
	})
}

func (s *RoleService) Get(ctx context.Context, req *GetRoleReq) (*Role, error) {
	return s.repo.FindOne(ctx, req)
}

func (s *RoleService) Update(ctx context.Context, req *RoleUpdateReq) error {
	return s.repo.Update(ctx, &Role{
		Model:       model.Model{ID: req.ID},
		Code:        req.Code,
		Description: req.Description,
		Status:      req.Status,
	})
}

func (s *RoleService) List(ctx context.Context, req *RoleListReq) ([]*Role, int64, error) {
	return s.repo.Find(ctx, req)
}

func (s *RoleService) Delete(ctx context.Context, ids string) error {
	return s.repo.Delete(ctx, ids)
}
