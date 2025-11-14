package repo

import (
	"context"
	"database/sql"

	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/bobacgo/admin-service/apps/repo/model"
	. "github.com/bobacgo/orm"
)

type I18nRepo struct {
	clt *data.Client
}

func NewI18nRepo(clt *data.Client) *I18nRepo {
	return &I18nRepo{clt: clt}
}

func (r *I18nRepo) FindOne(ctx context.Context, req *dto.GetI18nReq) (*model.I18n, error) {
	where := map[string]any{}
	if req.ID > 0 {
		where[AND(model.Id)] = req.ID
	}
	if req.Key != "" {
		where[AND(model.Key)] = req.Key
	}
	if req.Class != "" {
		where[AND(model.Class)] = req.Class
	}
	if req.Lang != "" {
		where[AND(model.Lang)] = req.Lang
	}

	row := new(model.I18n)
	err := SELECT1(row).FROM(model.I18nTable).WHERE(where).Query(ctx, r.clt.DB)
	return row, err
}

func (r *I18nRepo) Find(ctx context.Context, req *dto.I18nListReq) ([]*model.I18n, int64, error) {
	where := map[string]any{}
	if req.Key != "" {
		where[AND(model.Key)] = req.Key
	}
	if req.Class != "" {
		where[AND(model.Class)] = req.Class
	}
	if req.Lang != "" {
		where[AND(model.Lang)] = req.Lang
	}

	var (
		list  = make([]*model.I18n, 0)
		total sql.Null[int64]
	)
	if err := SELECT1(COUNT("*", &total)).FROM(model.I18nTable).WHERE(where).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}
	if !total.Valid {
		return list, 0, nil
	}

	rows := make([]*model.I18n, 0)
	offset, limit := req.Limit()
	if err := SELECT2(&rows).FROM(model.I18nTable).WHERE(where).ORDER_BY(DESC(model.Id)).OFFSET(int64(offset)).LIMIT(int64(limit)).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}

	return rows, total.V, nil
}

func (r *I18nRepo) Create(ctx context.Context, row *model.I18n) error {
	id, err := INSERT(row).INTO(model.I18nTable).Omit(model.Id).Exec(ctx, r.clt.DB)
	row.ID = id
	return err
}

func (r *I18nRepo) Update(ctx context.Context, row *model.I18n) error {
	_, err := UPDATE(model.I18nTable).SET1(row).WHERE(M{AND(model.Id): row.ID}).Omit(model.Id).Exec(ctx, r.clt.DB)
	return err
}

func (r *I18nRepo) Delete(ctx context.Context, ids string) error {
	_, err := DELETE().FROM(model.I18nTable).WHERE(M{AND_IN(model.Id): ids}).Exec(ctx, r.clt.DB)
	return err
}
