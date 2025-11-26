package menu

import (
	"context"
	"database/sql"

	"github.com/bobacgo/admin-service/apps/repo"
	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/repo/model"
	. "github.com/bobacgo/orm"
)

type MenuRepo struct {
	clt *data.Client
}

func NewMenuRepo(clt *data.Client) *MenuRepo {
	return &MenuRepo{clt: clt}
}

func (r *MenuRepo) FindOne(ctx context.Context, req *GetMenuReq) (*Menu, error) {
	row := new(Menu)
	where := make(map[string]any)
	if req.ID > 0 {
		where[repo.AND(model.Id)] = req.ID
	}
	if req.Path != "" {
		where[repo.AND(Path)] = req.Path
	}

	err := SELECT1(row).FROM(MenuTable).WHERE(where).Query(ctx, r.clt.DB)
	return row, err
}

func (r *MenuRepo) Find(ctx context.Context, req *MenuListReq) ([]*Menu, int64, error) {
	where := map[string]any{}
	if req.Path != "" {
		where[repo.AND_LIKE(Path)] = req.Path + "%"
	}
	if req.Name != "" {
		where[repo.AND_LIKE(Name)] = req.Name + "%"
	}

	var (
		list  = make([]*Menu, 0)
		total sql.Null[int64]
	)
	if err := SELECT1(COUNT("*", &total)).FROM(MenuTable).WHERE(where).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}
	if !total.Valid {
		return list, 0, nil
	}

	offset, limit := req.Limit()
	if err := SELECT2(&list).FROM(MenuTable).WHERE(where).ORDER_BY(repo.DESC(model.Id)).OFFSET(int64(offset)).LIMIT(int64(limit)).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}

	return list, total.V, nil
}

func (r *MenuRepo) Create(ctx context.Context, row *Menu) error {
	id, err := INSERT(row).INTO(MenuTable).Omit(model.Id).Exec(ctx, r.clt.DB)
	row.ID = id
	return err
}

func (r *MenuRepo) Update(ctx context.Context, row *Menu) error {
	_, err := UPDATE(MenuTable).SET1(row).WHERE(M{repo.AND(model.Id): row.ID}).Omit(model.Id).Exec(ctx, r.clt.DB)
	return err
}

func (r *MenuRepo) Delete(ctx context.Context, ids string) error {
	_, err := DELETE().FROM(MenuTable).WHERE(M{repo.AND_IN(model.Id): ids}).Exec(ctx, r.clt.DB)
	return err
}
