package menu

import (
	"context"
	"encoding/json"

	"github.com/bobacgo/admin-service/apps/repo/model"
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
func (s *MenuService) Put(ctx context.Context, req *MenuUpdateReq) (*Menu, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	menu := &Menu{
		Model:     model.Model{ID: req.ID},
		Path:      req.Path,
		Name:      req.Name,
		Component: req.Component,
		Redirect:  req.Redirect,
		Meta:      req.Meta,
		Icon:      req.Icon,
	}

	if err := s.repo.Update(ctx, menu); err != nil {
		return nil, err
	}

	return menu, nil
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
