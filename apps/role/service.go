package role

import (
	"context"

	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/bobacgo/admin-service/apps/repo/model"
	"github.com/go-playground/validator/v10"
)

type RoleService struct {
	repo      *RoleRepo
	validator *validator.Validate
}

func NewRoleService(r *RoleRepo, v *validator.Validate) *RoleService {
	return &RoleService{repo: r, validator: v}
}

// GetOne 获取单个角色
func (s *RoleService) GetOne(ctx context.Context, req *GetRoleReq) (*Role, error) {
	return s.repo.FindOne(ctx, req)
}

// GetList 获取角色列表
func (s *RoleService) GetList(ctx context.Context, req *RoleListReq) (*dto.PageResp[Role], error) {
	list, total, err := s.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	return dto.NewPageResp(total, list), nil
}

// PostCreate 创建角色
func (s *RoleService) PostCreate(ctx context.Context, req *RoleCreateReq) (*Role, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	role := &Role{
		Code:        req.Code,
		Description: req.Description,
		Status:      req.Status,
	}

	if err := s.repo.Create(ctx, role); err != nil {
		return nil, err
	}

	return role, nil
}

// PutUpdate 更新角色
func (s *RoleService) PutUpdate(ctx context.Context, req *RoleUpdateReq) (*Role, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	role := &Role{
		Model:       model.Model{ID: req.ID},
		Code:        req.Code,
		Description: req.Description,
		Status:      req.Status,
	}

	if err := s.repo.Update(ctx, role); err != nil {
		return nil, err
	}

	return role, nil
}

// DeleteDel 删除角色
func (s *RoleService) DeleteDel(ctx context.Context, req *DeleteRoleReq) (interface{}, error) {
	return nil, s.repo.Delete(ctx, req.IDs)
}
