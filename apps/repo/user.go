package repo

import (
	"context"
	"database/sql"

	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/bobacgo/admin-service/apps/repo/model"
	. "github.com/bobacgo/orm"
)

type UserRepo struct {
	clt *data.Client
}

func NewUserRepo(data *data.Client) *UserRepo {
	return &UserRepo{clt: data}
}

// 创建
func (r *UserRepo) Create(ctx context.Context, row *model.User) error {
	id, err := INSERT(row).INTO(model.UsersTable).Exec(ctx, r.clt.DB)
	row.ID = id
	return err
}

func (r *UserRepo) FindOne(ctx context.Context, req *dto.GetUserReq) (*model.User, error) {
	row := new(model.User)

	where := make(map[string]any)
	if req.ID > 0 {
		where[AND(model.Id)] = req.ID
	}
	if req.Account != "" {
		where[AND(model.Account)] = req.Account
	}
	if req.Phone != "" {
		where[AND(model.Phone)] = req.Phone
	}
	if req.Email != "" {
		where[AND(model.Email)] = req.Email
	}

	err := SELECT1(row).FROM(model.UsersTable).WHERE(where).Query(ctx, r.clt.DB)
	return row, err
}

func (r *UserRepo) Find(ctx context.Context, req *dto.UserListReq) ([]*model.User, int64, error) {
	where := map[string]any{}
	if req.Account != "" {
		where[AND_LIKE(model.Account)] = req.Account + "%" // 右模糊查询
	}
	if req.Phone != "" {
		where[AND_LIKE(model.Phone)] = req.Phone + "%" // 右模糊查询
	}
	if req.Email != "" {
		where[AND_LIKE(model.Email)] = req.Email + "%" // 右模糊查询
	}
	if req.Status != "" {
		where[AND_IN(model.Status)] = req.Status
	}

	var (
		list  = make([]*model.User, 0)
		total sql.Null[int64]
	)
	if err := SELECT1(COUNT("*", &total)).FROM(model.UsersTable).WHERE(where).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}
	if !total.Valid {
		return list, 0, nil
	}
	offset, limit := req.Limit()
	if err := SELECT2(&list).FROM(model.UsersTable).WHERE(where).ORDER_BY(DESC(model.Id)).OFFSET(int64(offset)).LIMIT(int64(limit)).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}

	return list, total.V, nil
}

func (r *UserRepo) Update(ctx context.Context, row *model.User) error {
	_, err := UPDATE(model.UsersTable).SET1(row).WHERE(M{AND(model.Id): row.ID}).Exec(ctx, r.clt.DB)
	return err
}

func (r *UserRepo) Delete(ctx context.Context, ids string) error {
	_, err := DELETE().FROM(model.UsersTable).WHERE(M{AND_IN(model.Id): ids}).Exec(ctx, r.clt.DB)
	return err
}