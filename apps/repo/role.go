package repo

import (
	"context"
	"database/sql"

	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/bobacgo/admin-service/apps/repo/model"
	. "github.com/bobacgo/orm"
)

type RoleRepo struct {
	clt *data.Client
}

func NewRoleRepo(clt *data.Client) *RoleRepo {
	return &RoleRepo{clt: clt}
}

func (r *RoleRepo) FindOne(ctx context.Context, req *dto.GetRoleReq) (*model.Role, error) {
	row := new(model.Role)
	where := make(map[string]any)
	if req.ID > 0 {
		where[AND(model.Id)] = req.ID
	}
	if req.Code != "" {
		where[AND(model.Code)] = req.Code
	}

	err := SELECT1(row).FROM(model.RoleTable).WHERE(where).Query(ctx, r.clt.DB)
	return row, err
}

func (r *RoleRepo) Find(ctx context.Context, req *dto.RoleListReq) ([]*model.Role, int64, error) {
	where := map[string]any{}
	if req.Code != "" {
		where[AND_LIKE(model.Code)] = req.Code + "%"
	}
	if req.Status != "" {
		where[AND_IN(model.Status)] = req.Status
	}

	var (
		list  = make([]*model.Role, 0)
		total sql.Null[int64]
	)
	if err := SELECT1(COUNT("*", &total)).FROM(model.RoleTable).WHERE(where).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}
	if !total.Valid {
		return list, 0, nil
	}

	offset, limit := req.Limit()
	if err := SELECT2(&list).FROM(model.RoleTable).WHERE(where).ORDER_BY(DESC(model.Id)).OFFSET(int64(offset)).LIMIT(int64(limit)).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}

	return list, total.V, nil
}

func (r *RoleRepo) Create(ctx context.Context, row *model.Role) error {
	id, err := INSERT(row).INTO(model.RoleTable).Exec(ctx, r.clt.DB)
	row.ID = id
	return err
}

func (r *RoleRepo) Update(ctx context.Context, row *model.Role) error {
	_, err := UPDATE(model.RoleTable).SET1(row).WHERE(M{AND(model.Id): row.ID}).Exec(ctx, r.clt.DB)
	return err
}

func (r *RoleRepo) Delete(ctx context.Context, ids string) error {
	_, err := DELETE().FROM(model.RoleTable).WHERE(M{AND_IN(model.Id): ids}).Exec(ctx, r.clt.DB)
	return err
}
