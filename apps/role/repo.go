package role

import (
	"context"
	"database/sql"

	"github.com/bobacgo/admin-service/apps/common/model"
	"github.com/bobacgo/admin-service/apps/common/repo"
	"github.com/bobacgo/admin-service/apps/common/repo/data"
	. "github.com/bobacgo/orm"
)

type RoleRepo struct {
	clt *data.Client
}

func NewRoleRepo(clt *data.Client) *RoleRepo {
	return &RoleRepo{clt: clt}
}

func (r *RoleRepo) FindOne(ctx context.Context, req *GetRoleReq) (*Role, error) {
	row := new(Role)
	where := make(map[string]any)
	if req.ID > 0 {
		where[repo.AND(model.Id)] = req.ID
	}
	if req.RoleName != "" {
		where[repo.AND(RoleName)] = req.RoleName
	}

	err := SELECT1(row).FROM(RoleTable).WHERE(where).Query(ctx, r.clt.DB)
	return row, err
}

func (r *RoleRepo) Find(ctx context.Context, req *RoleListReq) ([]*Role, int64, error) {
	where := map[string]any{}
	if req.RoleName != "" {
		where[repo.AND_LIKE(RoleName)] = req.RoleName + "%"
	}
	if req.Status != "" {
		where[repo.AND_IN(model.Status)] = req.Status
	}

	var (
		list  = make([]*Role, 0)
		total sql.Null[int64]
	)
	if err := SELECT1(COUNT("*", &total)).FROM(RoleTable).WHERE(where).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}
	if !total.Valid {
		return list, 0, nil
	}

	offset, limit := req.Limit()
	if err := SELECT2(&list).FROM(RoleTable).WHERE(where).ORDER_BY(repo.DESC(model.Id)).OFFSET(int64(offset)).LIMIT(int64(limit)).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}

	return list, total.V, nil
}

func (r *RoleRepo) Create(ctx context.Context, row *Role) error {
	id, err := INSERT(row).INTO(RoleTable).Omit(model.Id).Exec(ctx, r.clt.DB)
	row.ID = id
	return err
}

func (r *RoleRepo) Update(ctx context.Context, row *Role) error {
	_, err := UPDATE(RoleTable).SET1(row).WHERE(M{repo.AND(model.Id): row.ID}).Omit(model.Id).Exec(ctx, r.clt.DB)
	return err
}

func (r *RoleRepo) Delete(ctx context.Context, ids string) error {
	_, err := DELETE().FROM(RoleTable).WHERE(M{repo.AND_IN(model.Id): ids}).Exec(ctx, r.clt.DB)
	return err
}
