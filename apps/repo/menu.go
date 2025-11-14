package repo

import (
	"context"
	"database/sql"

	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/bobacgo/admin-service/apps/repo/model"
	. "github.com/bobacgo/orm"
)

type MenuRepo struct {
	clt *data.Client
}

func NewMenuRepo(clt *data.Client) *MenuRepo {
	return &MenuRepo{clt: clt}
}

func (r *MenuRepo) FindOne(ctx context.Context, req *dto.GetMenuReq) (*model.Menu, error) {
	row := new(model.Menu)
	where := make(map[string]any)
	if req.ID > 0 {
		where[AND(model.Id)] = req.ID
	}
	if req.Path != "" {
		where[AND(model.Path)] = req.Path
	}

	err := SELECT1(row).FROM(model.MenuTable).WHERE(where).Query(ctx, r.clt.DB)
	return row, err
}

func (r *MenuRepo) Find(ctx context.Context, req *dto.MenuListReq) ([]*model.Menu, int64, error) {
	where := map[string]any{}
	if req.Path != "" {
		where[AND_LIKE(model.Path)] = req.Path + "%"
	}
	if req.Name != "" {
		where[AND_LIKE(model.Name)] = req.Name + "%"
	}

	var (
		list  = make([]*model.Menu, 0)
		total sql.Null[int64]
	)
	if err := SELECT1(COUNT("*", &total)).FROM(model.MenuTable).WHERE(where).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}
	if !total.Valid {
		return list, 0, nil
	}

	offset, limit := req.Limit()
	if err := SELECT2(&list).FROM(model.MenuTable).WHERE(where).ORDER_BY(DESC(model.Id)).OFFSET(int64(offset)).LIMIT(int64(limit)).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}

	return list, total.V, nil
}

func (r *MenuRepo) Create(ctx context.Context, row *model.Menu) error {
	id, err := INSERT(row).INTO(model.MenuTable).Omit(model.Id).Exec(ctx, r.clt.DB)
	row.ID = id
	return err
}

func (r *MenuRepo) Update(ctx context.Context, row *model.Menu) error {
	_, err := UPDATE(model.MenuTable).SET1(row).WHERE(M{AND(model.Id): row.ID}).Omit(model.Id).Exec(ctx, r.clt.DB)
	return err
}

func (r *MenuRepo) Delete(ctx context.Context, ids string) error {
	_, err := DELETE().FROM(model.MenuTable).WHERE(M{AND_IN(model.Id): ids}).Exec(ctx, r.clt.DB)
	return err
}
