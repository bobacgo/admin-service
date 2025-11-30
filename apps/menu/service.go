package menu

import (
	"context"
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type MenuService struct {
	repo      *MenuRepo
	validator *validator.Validate
}

func NewMenuService(r *MenuRepo, v *validator.Validate) *MenuService {
	return &MenuService{repo: r, validator: v}
}

// Get /menu/list 获取菜单列表
func (s *MenuService) GetList(ctx context.Context, req *MenuListReq) (*MenuListResp, error) {
	list, err := s.repo.Find(ctx)
	if err != nil {
		return nil, err
	}
	return &MenuListResp{List: list}, nil
}

// Post /menu 创建菜单
func (s *MenuService) Post(ctx context.Context, req *MenuCreateReq) (any, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	menu := &Menu{
		Path:      req.Path,
		Name:      req.Name,
		Component: req.Component,
		Redirect:  req.Redirect,
		Meta:      req.Meta,
		Icon:      req.Icon,
	}

	err := s.repo.Create(ctx, menu)
	return nil, err
}

// Put /menu 更新菜单
func (s *MenuService) Put(ctx context.Context, req *MenuUpdateReq) (any, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	// 先查询现有菜单数据
	existMenu, err := s.repo.FindOne(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// 只更新前端发送的非空字段
	if req.ParentID != 0 {
		existMenu.ParentID = req.ParentID
	}
	if req.Path != "" {
		existMenu.Path = req.Path
	}
	if req.Name != "" {
		existMenu.Name = req.Name
	}
	if req.Component != "" {
		existMenu.Component = req.Component
	}
	if req.Redirect != "" {
		existMenu.Redirect = req.Redirect
	}
	if req.Meta != "" {
		existMenu.Meta = req.Meta
	}
	if req.Icon != "" {
		existMenu.Icon = req.Icon
	}
	if req.Sort != 0 {
		existMenu.Sort = req.Sort
	}

	if err := s.repo.Update(ctx, existMenu); err != nil {
		return nil, err
	}

	return nil, nil
}

// Delete /menu 删除菜单
func (s *MenuService) Delete(ctx context.Context, req *DeleteMenuReq) (any, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}
	return nil, s.repo.Delete(ctx, req.IDs)
}

// Get /menu/tree 获取菜单树
func (s *MenuService) GetTree(ctx context.Context, _ any) (*MenuTreeResp, error) {
	menuList, err := s.repo.Find(ctx)
	if err != nil {
		return nil, err
	}
	return &MenuTreeResp{List: s.buildTree(menuList)}, nil
}

func (s *MenuService) buildTree(menuList []*Menu) []*MenuItem {
	var tree []*MenuItem
	for _, menu := range menuList {
		if menu.ParentID == 0 {
			meta := make(map[string]any)
			if menu.Meta != "" {
				_ = json.Unmarshal([]byte(menu.Meta), &meta)
			}
			tree = append(tree, &MenuItem{
				ID:        menu.ID,
				ParentID:  menu.ParentID,
				Path:      menu.Path,
				Name:      menu.Name,
				Component: menu.Component,
				Redirect:  menu.Redirect,
				Meta:      meta,
				Icon:      menu.Icon,
				Sort:      menu.Sort,
				Children:  make([]*MenuItem, 0),
			})
		}
	}
	for _, menu := range menuList {
		if menu.ParentID != 0 {
			for _, item := range tree {
				if item.ID == menu.ParentID {
					meta := make(map[string]any)
					if menu.Meta != "" {
						_ = json.Unmarshal([]byte(menu.Meta), &meta)
					}
					item.Children = append(item.Children, &MenuItem{
						ID:        menu.ID,
						ParentID:  menu.ParentID,
						Path:      menu.Path,
						Name:      menu.Name,
						Component: menu.Component,
						Redirect:  menu.Redirect,
						Meta:      meta,
						Icon:      menu.Icon,
						Sort:      menu.Sort,
						Children:  make([]*MenuItem, 0),
					})
				}
			}
		}
	}
	return tree
}
