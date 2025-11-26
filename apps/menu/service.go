package menu

import (
	"context"
	"encoding/json"

	"github.com/bobacgo/admin-service/apps/repo/model"
)

type MenuService struct {
	repo *MenuRepo
}

func NewMenuService(r *MenuRepo) *MenuService {
	return &MenuService{repo: r}
}

func (s *MenuService) Create(ctx context.Context, req *MenuCreateReq) error {
	return s.repo.Create(ctx, &Menu{
		Path:      req.Path,
		Name:      req.Name,
		Component: req.Component,
		Redirect:  req.Redirect,
		Meta:      req.Meta,
		Icon:      req.Icon,
	})
}

func (s *MenuService) Get(ctx context.Context, req *GetMenuReq) (*Menu, error) {
	return s.repo.FindOne(ctx, req)
}

func (s *MenuService) Update(ctx context.Context, req *MenuUpdateReq) error {
	return s.repo.Update(ctx, &Menu{
		Model:     model.Model{ID: req.ID},
		Path:      req.Path,
		Name:      req.Name,
		Component: req.Component,
		Redirect:  req.Redirect,
		Meta:      req.Meta,
		Icon:      req.Icon,
	})
}

func (s *MenuService) List(ctx context.Context, req *MenuListReq) ([]*Menu, int64, error) {
	return s.repo.Find(ctx, req)
}

func (s *MenuService) Delete(ctx context.Context, ids string) error {
	return s.repo.Delete(ctx, ids)
}

func (s *MenuService) Tree(ctx context.Context) ([]*MenuItem, error) {
	menuList, _, err := s.List(ctx, &MenuListReq{})
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
