package role

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bobacgo/admin-service/apps/common/dto"
	"github.com/bobacgo/admin-service/apps/common/model"
	"github.com/bobacgo/admin-service/apps/menu"
	"github.com/bobacgo/admin-service/apps/user"
	"github.com/go-playground/validator/v10"
)

type RoleService struct {
	repo      *RoleRepo
	menuRepo  *menu.MenuRepo
	userRepo  *user.UserRepo
	validator *validator.Validate
}

func NewRoleService(r *RoleRepo, mr *menu.MenuRepo, ur *user.UserRepo, v *validator.Validate) *RoleService {
	return &RoleService{repo: r, menuRepo: mr, userRepo: ur, validator: v}
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
	if len(list) == 0 {
		return dto.NewPageResp(total, list), nil
	}

	ids := make([]int64, 0, len(list))
	for _, r := range list {
		ids = append(ids, r.ID)
	}

	counts, err := s.userRepo.CountByRoleIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, r := range list {
		if cnt, ok := counts[r.ID]; ok {
			r.UserCount = cnt
		}
	}
	return dto.NewPageResp(total, list), nil
}

// Post /role 创建角色
func (s *RoleService) Post(ctx context.Context, req *RoleCreateReq) (any, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	role := &Role{
		RoleName:    req.RoleName,
		Description: req.Description,
		Status:      req.Status,
		Model: model.Model{
			Operator:  req.Operator,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}

	if err := s.repo.Create(ctx, role); err != nil {
		return nil, err
	}

	return struct{}{}, nil
}

// Put /role 更新角色
func (s *RoleService) Put(ctx context.Context, req *RoleUpdateReq) (any, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	// 先查询现有角色数据
	existRole, err := s.repo.FindOne(ctx, &GetRoleReq{ID: req.ID})
	if err != nil {
		return nil, err
	}

	// 只更新前端发送的非空字段
	if req.RoleName != "" {
		existRole.RoleName = req.RoleName
	}
	if req.Description != "" {
		existRole.Description = req.Description
	}
	if req.Status != 0 {
		existRole.Status = req.Status
	}
	if req.Operator != "" {
		existRole.Operator = req.Operator
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

// PostPermissions POST /permissions 保存角色权限（更新菜单的role_codes字段）
func (s *RoleService) PostPermissions(ctx context.Context, req *SaveRolePermissionsReq) (interface{}, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	// 先查询角色信息
	role, err := s.repo.FindOne(ctx, &GetRoleReq{ID: req.RoleId})
	if err != nil {
		// 如果发生错误（通常是角色不存在），直接返回
		return nil, err
	}

	// 删除该角色从所有菜单的role_ids中
	if err := s.menuRepo.RemoveRoleIdFromAllMenus(ctx, fmt.Sprintf("%d", role.ID)); err != nil {
		return nil, err
	}

	// 如果有选中的菜单，添加到这些菜单的role_ids中
	if len(req.MenuIds) > 0 {
		if err := s.menuRepo.AddRoleIdToMenus(ctx, fmt.Sprintf("%d", role.ID), req.MenuIds); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// GetPermissions GET /permissions 获取角色权限（菜单ID列表）
func (s *RoleService) GetPermissions(ctx context.Context, req *GetRolePermissionsReq) (*GetRolePermissionsResp, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	// 先查询角色信息
	role, err := s.repo.FindOne(ctx, &GetRoleReq{ID: req.RoleId})
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("角色不存在")
		}
		return nil, err
	}

	// 获取该角色的所有菜单ID
	ids, err := s.menuRepo.GetMenuIdsByRoleId(ctx, fmt.Sprintf("%d", role.ID))
	if err != nil {
		return nil, err
	}

	return &GetRolePermissionsResp{MenuIds: ids}, nil
}
