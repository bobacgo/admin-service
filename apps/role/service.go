package role

import (
	"context"
	"time"

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

// Get /role/one 获取单个角色
func (s *RoleService) GetOne(ctx context.Context, req *GetRoleReq) (*Role, error) {
	return s.repo.FindOne(ctx, req)
}

// Get /role/list 获取角色列表
func (s *RoleService) GetList(ctx context.Context, req *RoleListReq) (*dto.PageResp[Role], error) {
	list, total, err := s.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	return dto.NewPageResp(total, list), nil
}

// Post /role 创建角色
func (s *RoleService) Post(ctx context.Context, req *RoleCreateReq) (*Role, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	role := &Role{
		Code:        req.Code,
		Description: req.Description,
		Status:      req.Status,
		Model: model.Model{
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}

	if err := s.repo.Create(ctx, role); err != nil {
		return nil, err
	}

	return role, nil
}

// Put /role 更新角色
func (s *RoleService) Put(ctx context.Context, req *RoleUpdateReq) (interface{}, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	// 先查询现有角色数据
	existRole, err := s.repo.FindOne(ctx, &GetRoleReq{ID: req.ID})
	if err != nil {
		return nil, err
	}

	// 只更新前端发送的非空字段
	if req.Code != "" {
		existRole.Code = req.Code
	}
	if req.Description != "" {
		existRole.Description = req.Description
	}
	if req.Status != 0 {
		existRole.Status = req.Status
	}

	existRole.UpdatedAt = time.Now().Unix()

	if err := s.repo.Update(ctx, existRole); err != nil {
		return nil, err
	}

	return nil, nil
}

// Delete /role 删除角色
func (s *RoleService) Delete(ctx context.Context, req *DeleteRoleReq) (interface{}, error) {
	return nil, s.repo.Delete(ctx, req.IDs)
}
