package repo

import (
	"context"
	"database/sql"

	"github.com/bobacgo/admin-service/apps/common/model"
	"github.com/bobacgo/admin-service/apps/common/repo"
	"github.com/bobacgo/admin-service/apps/common/repo/data"
	"github.com/bobacgo/admin-service/apps/mgr/dto"
	model2 "github.com/bobacgo/admin-service/apps/mgr/repo/model"
	. "github.com/bobacgo/orm"
)

type I18nRepo struct {
	clt *data.Client
}

func NewI18nRepo(clt *data.Client) *I18nRepo {
	return &I18nRepo{clt: clt}
}

func (r *I18nRepo) FindOne(ctx context.Context, req *dto.GetI18nReq) (*model2.I18n, error) {
	row := new(model2.I18n)
	where := make(map[string]any)
	if req.ID > 0 {
		where[repo.AND(model.Id)] = req.ID
	}
	if req.Key != "" {
		where[repo.AND(model2.Key)] = req.Key
	}
	if req.Class != "" {
		where[repo.AND(model2.Class)] = req.Class
	}
	if req.Lang != "" {
		where[repo.AND(model2.Lang)] = req.Lang
	}

	err := SELECT1(row).FROM(model2.I18nTable).WHERE(where).Query(ctx, r.clt.DB)
	return row, err
}

func (r *I18nRepo) Find(ctx context.Context, req *dto.I18nListReq) ([]*model2.I18n, int64, error) {
	where := map[string]any{}
	if req.Key != "" {
		where[repo.AND(model2.Key)] = req.Key
	}
	if req.Class != "" {
		where[repo.AND(model2.Class)] = req.Class
	}
	if req.Lang != "" {
		where[repo.AND(model2.Lang)] = req.Lang
	}

	var (
		list  = make([]*model2.I18n, 0)
		total sql.Null[int64]
	)
	if err := SELECT1(COUNT("*", &total)).FROM(model2.I18nTable).WHERE(where).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}
	if !total.Valid {
		return list, 0, nil
	}

	offset, limit := req.Limit()
	if err := SELECT2(&list).FROM(model2.I18nTable).WHERE(where).ORDER_BY(repo.DESC(model.Id)).OFFSET(int64(offset)).LIMIT(int64(limit)).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}

	return list, total.V, nil
}

func (r *I18nRepo) Create(ctx context.Context, row *model2.I18n) error {
	id, err := INSERT(row).INTO(model2.I18nTable).Omit(model.Id).Exec(ctx, r.clt.DB)
	row.ID = id
	return err
}

func (r *I18nRepo) Update(ctx context.Context, row *model2.I18n) error {
	_, err := UPDATE(model2.I18nTable).SET1(row).WHERE(M{repo.AND(model.Id): row.ID}).Omit(model.Id).Exec(ctx, r.clt.DB)
	return err
}

func (r *I18nRepo) Delete(ctx context.Context, ids string) error {
	_, err := DELETE().FROM(model2.I18nTable).WHERE(M{repo.AND_IN(model.Id): ids}).Exec(ctx, r.clt.DB)
	return err
}