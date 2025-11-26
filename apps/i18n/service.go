package i18n

import (
	"context"
	"time"

	"github.com/bobacgo/admin-service/apps/repo/model"
)

type I18nService struct {
	repo *I18nRepo
}

func NewI18nService(r *I18nRepo) *I18nService {
	return &I18nService{repo: r}
}

func (s *I18nService) Get(ctx context.Context, req *GetI18nReq) (*I18n, error) {
	return s.repo.FindOne(ctx, req)
}

func (s *I18nService) List(ctx context.Context, req *I18nListReq) ([]*I18n, int64, error) {
	return s.repo.Find(ctx, req)
}

func (s *I18nService) Create(ctx context.Context, req *I18nCreateReq) error {
	return s.repo.Create(ctx, &I18n{
		Class: req.Class,
		Lang:  req.Lang,
		Key:   req.Key,
		Value: req.Value,
		Model: model.Model{
			CreatedAt: time.Now().Unix(),
		},
	})
}

func (s *I18nService) Update(ctx context.Context, req *I18nUpdateReq) error {
	return s.repo.Update(ctx, &I18n{
		Class: req.Class,
		Lang:  req.Lang,
		Value: req.Value,
		Model: model.Model{
			ID: req.ID,
		},
	})
}

func (s *I18nService) Delete(ctx context.Context, ids string) error {
	return s.repo.Delete(ctx, ids)
}
