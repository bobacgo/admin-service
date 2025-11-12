package service

import (
	"context"

	"github.com/bobacgo/admin-service/apps/repo"
	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/bobacgo/admin-service/apps/repo/model"
)

type MenuService struct {
	repo *repo.Repo
}

func NewMenuService(repo *repo.Repo) *MenuService {
	return &MenuService{repo: repo}
}

func (s *MenuService) Create(ctx context.Context, req *dto.MenuCreateReq) error {
	return s.repo.Menu.Create(ctx, &model.Menu{
		Path:      req.Path,
		Name:      req.Name,
		Component: req.Component,
		Redirect:  req.Redirect,
		Meta:      req.Meta,
		Icon:      req.Icon,
	})
}

func (s *MenuService) Get(ctx context.Context, req *dto.GetMenuReq) (*model.Menu, error) {
	return s.repo.Menu.FindOne(ctx, req)
}

func (s *MenuService) Update(ctx context.Context, req *dto.MenuUpdateReq) error {
	return s.repo.Menu.Update(ctx, &model.Menu{
		Model:     model.Model{ID: req.ID},
		Path:      req.Path,
		Name:      req.Name,
		Component: req.Component,
		Redirect:  req.Redirect,
		Meta:      req.Meta,
		Icon:      req.Icon,
	})
}

func (s *MenuService) List(ctx context.Context, req *dto.MenuListReq) ([]*model.Menu, int64, error) {
	return s.repo.Menu.Find(ctx, req)
}

func (s *MenuService) Delete(ctx context.Context, ids string) error {
	return s.repo.Menu.Delete(ctx, ids)
}
