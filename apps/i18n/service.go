package i18n

import (
	"context"
	"time"

	"github.com/bobacgo/admin-service/apps/common/dto"
	"github.com/bobacgo/admin-service/apps/common/model"
	"github.com/go-playground/validator/v10"
)

type I18nService struct {
	repo      *I18nRepo
	validator *validator.Validate
}

func NewI18nService(r *I18nRepo, v *validator.Validate) *I18nService {
	return &I18nService{repo: r, validator: v}
}

// Get /i18n/one 获取单个i18n记录
func (s *I18nService) GetOne(ctx context.Context, req *GetI18nReq) (*I18n, error) {
	return s.repo.FindOne(ctx, req)
}

// Get /i18n/list 获取i18n列表
func (s *I18nService) GetList(ctx context.Context, req *I18nListReq) (*dto.PageResp[I18n], error) {
	list, total, err := s.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	return dto.NewPageResp(total, list), nil
}

// Post /i18n 创建i18n记录
func (s *I18nService) Post(ctx context.Context, req *I18nCreateReq) (*I18n, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	i18n := &I18n{
		Class: req.Class,
		Lang:  req.Lang,
		Key:   req.Key,
		Value: req.Value,
		Model: model.Model{
			Operator:  req.Operator,
			CreatedAt: time.Now().Unix(),
		},
	}

	if err := s.repo.Create(ctx, i18n); err != nil {
		return nil, err
	}

	return i18n, nil
}

// Put /i18n 更新i18n记录
func (s *I18nService) Put(ctx context.Context, req *I18nUpdateReq) (interface{}, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	// 先查询现有i18n数据
	existI18n, err := s.repo.FindOne(ctx, &GetI18nReq{ID: req.ID})
	if err != nil {
		return nil, err
	}

	// 只更新前端发送的非空字段
	if req.Class != "" {
		existI18n.Class = req.Class
	}
	if req.Lang != "" {
		existI18n.Lang = req.Lang
	}
	if req.Value != "" {
		existI18n.Value = req.Value
	}
	if req.Operator != "" {
		existI18n.Operator = req.Operator
	}

	existI18n.UpdatedAt = time.Now().Unix()

	if err := s.repo.Update(ctx, existI18n); err != nil {
		return nil, err
	}

	return nil, nil
}

// Delete /i18n 删除i18n记录
func (s *I18nService) Delete(ctx context.Context, req *DeleteI18nReq) (interface{}, error) {
	return nil, s.repo.Delete(ctx, req.IDs)
}
