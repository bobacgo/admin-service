package menu

import (
	"context"

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

func (r *MenuRepo) Find(ctx context.Context) ([]*Menu, error) {
	list := make([]*Menu, 0)
	if err := SELECT2(&list).FROM(MenuTable).ORDER_BY(repo.DESC(model.Id)).Query(ctx, r.clt.DB); err != nil {
		return nil, err
	}
	return list, nil
}

func (r *MenuRepo) FindOne(ctx context.Context, id int64) (*Menu, error) {
	row := new(Menu)
	if err := SELECT1(row).FROM(MenuTable).WHERE(M{repo.AND(model.Id): id}).Query(ctx, r.clt.DB); err != nil {
		return nil, err
	}
	return row, nil
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
