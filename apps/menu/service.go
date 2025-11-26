package menu

import (
	"context"
	"encoding/json"

	"github.com/bobacgo/admin-service/apps/repo/dto"
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

// GetOne 获取单个菜单
func (s *MenuService) GetOne(ctx context.Context, req *GetMenuReq) (*Menu, error) {
	return s.repo.FindOne(ctx, req)
}

// GetList 获取菜单列表
func (s *MenuService) GetList(ctx context.Context, req *MenuListReq) (*dto.PageResp[Menu], error) {
	list, total, err := s.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	return dto.NewPageResp(total, list), nil
}

// PostCreate 创建菜单
func (s *MenuService) PostCreate(ctx context.Context, req *MenuCreateReq) (*Menu, error) {
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

	if err := s.repo.Create(ctx, menu); err != nil {
		return nil, err
	}

	return menu, nil
}

// PutUpdate 更新菜单
func (s *MenuService) PutUpdate(ctx context.Context, req *MenuUpdateReq) (*Menu, error) {
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

// DeleteDel 删除菜单
func (s *MenuService) DeleteDel(ctx context.Context, req *DeleteMenuReq) (interface{}, error) {
	return nil, s.repo.Delete(ctx, req.IDs)
}

// GetTree 获取菜单树
func (s *MenuService) GetTree(ctx context.Context, req interface{}) ([]*MenuItem, error) {
	menuList, _, err := s.repo.Find(ctx, &MenuListReq{})
	if err != nil {
		return nil, err
	}
	return s.buildTree(menuList), nil
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
