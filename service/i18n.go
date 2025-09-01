package service

import (
	"context"
	"time"

	"github.com/bobacgo/admin-service/repo"
	"github.com/bobacgo/admin-service/repo/dto"
	"github.com/bobacgo/admin-service/repo/model"
)

type I18nService struct {
	repo *repo.Repo
}

func NewI18nService(repo *repo.Repo) *I18nService {
	return &I18nService{
		repo: repo,
	}
}

func (svc *I18nService) Get(ctx context.Context, req *dto.GetI18nReq) (*model.I18n, error) {
	return svc.repo.I18n.FindOne(ctx, req)
}

func (svc *I18nService) List(ctx context.Context, req *dto.I18nListReq) ([]*model.I18n, error) {
	return svc.repo.I18n.Find(ctx, req)
}

func (svc *I18nService) Create(ctx context.Context, req *dto.I18nCreateReq) error {
	return svc.repo.I18n.Create(ctx, &model.I18n{
		Class: req.Class,
		Lang:  req.Lang,
		Key:   req.Key,
		Value: req.Value,
		Model: model.Model{
			CreatedAt: time.Now().Unix(),
		},
	})
}

func (svc *I18nService) Update(ctx context.Context, row *model.I18n) error {
	return svc.repo.I18n.Update(ctx, row)
}

func (svc *I18nService) Delete(ctx context.Context, ids string) error {
	return svc.repo.I18n.Delete(ctx, ids)
}
