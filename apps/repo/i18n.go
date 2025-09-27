package repo

import (
	"context"
	"log/slog"

	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/bobacgo/admin-service/apps/repo/model"
	"github.com/bobacgo/admin-service/pkg/kit/orm"
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
		where["AND id = ?"] = req.ID
	}
	if req.Key != "" {
		where["AND key = ?"] = req.Key
	}
	if req.Class != "" {
		where["AND class = ?"] = req.Class
	}
	if req.Lang != "" {
		where["AND lang = ?"] = req.Lang
	}

	row := new(model.I18n)
	err := orm.NewDB(r.clt.DB).Select(row).From(row.TableName()).Where(where).QueryRowContext(ctx)
	return row, err
}

func (r *I18nRepo) Find(ctx context.Context, req *dto.I18nListReq) ([]*model.I18n, int64, error) {
	where := map[string]any{}
	if req.Key != "" {
		where["AND key = ?"] = req.Key
	}
	if req.Class != "" {
		where["AND class = ?"] = req.Class
	}
	if req.Lang != "" {
		where["AND lang = ?"] = req.Lang
	}

	var count int64
	db := orm.NewDB(r.clt.DB)
	if err := db.Debug().Select(db.Count("*", &count)).From(model.I18nTable).Where(where).QueryContext(ctx); err != nil {
		return nil, 0, err
	}

	slog.InfoContext(ctx, "i18n find count", "count", count)

	return orm.FindPage[model.I18n](ctx, r.clt.DB, where, req.Page, req.PageSize)
}

func (r *I18nRepo) Create(ctx context.Context, row *model.I18n) error {
	_, err := orm.NewDB(r.clt.DB).INSERT(row).INTO(model.I18nTable).Exec(ctx)
	return err
}

func (r *I18nRepo) Update(ctx context.Context, row *model.I18n) error {
	_, err := orm.Update(ctx, r.clt.DB, row.ID, row)
	return err
}

func (r *I18nRepo) Delete(ctx context.Context, ids string) error {
	_, err := orm.Delete(ctx, r.clt.DB, model.I18nTable, ids)
	return err
}
