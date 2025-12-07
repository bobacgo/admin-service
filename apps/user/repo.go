package user

import (
	"context"
	"database/sql"

	"github.com/bobacgo/admin-service/apps/repo"
	"github.com/bobacgo/admin-service/apps/repo/data"
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
func (r *UserRepo) Create(ctx context.Context, row *User) error {
	id, err := INSERT(row).INTO(UsersTable).Omit(model.Id).Exec(ctx, r.clt.DB)
	row.ID = id
	return err
}

func (r *UserRepo) FindOne(ctx context.Context, req *GetUserReq) (*User, error) {
	row := new(User)

	where := make(map[string]any)
	if req.ID > 0 {
		where[repo.AND(model.Id)] = req.ID
	}
	if req.Account != "" {
		where[repo.AND(Account)] = req.Account
	}
	if req.Phone != "" {
		where[repo.AND(Phone)] = req.Phone
	}
	if req.Email != "" {
		where[repo.AND(Email)] = req.Email
	}

	err := SELECT1(row).FROM(UsersTable).WHERE(where).Query(ctx, r.clt.DB)
	return row, err
}

func (r *UserRepo) Find(ctx context.Context, req *UserListReq) ([]*User, int64, error) {
	where := map[string]any{}
	if req.Account != "" {
		where[repo.AND_LIKE(Account)] = req.Account + "%" // 右模糊查询
	}
	if req.Phone != "" {
		where[repo.AND_LIKE(Phone)] = req.Phone + "%" // 右模糊查询
	}
	if req.Email != "" {
		where[repo.AND_LIKE(Email)] = req.Email + "%" // 右模糊查询
	}
	if req.Status != "" {
		where[repo.AND_IN(model.Status)] = req.Status
	}

	var (
		list  = make([]*User, 0)
		total sql.Null[int64]
	)
	if err := SELECT1(COUNT("*", &total)).FROM(UsersTable).WHERE(where).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}
	if !total.Valid {
		return list, 0, nil
	}
	offset, limit := req.Limit()
	if err := SELECT2(&list).FROM(UsersTable).WHERE(where).ORDER_BY(repo.DESC(model.Id)).OFFSET(int64(offset)).LIMIT(int64(limit)).Query(ctx, r.clt.DB); err != nil {
		return nil, 0, err
	}

	return list, total.V, nil
}

func (r *UserRepo) Update(ctx context.Context, row *User) error {
	_, err := UPDATE(UsersTable).SET1(row).WHERE(M{repo.AND(model.Id): row.ID}).Omit(model.Id).Exec(ctx, r.clt.DB)
	return err
}

func (r *UserRepo) Delete(ctx context.Context, ids string) error {
	_, err := DELETE().FROM(UsersTable).WHERE(M{repo.AND_IN(model.Id): ids}).Exec(ctx, r.clt.DB)
	return err
}

// CountByRoleId 返回 role id 在 users.role_ids 字段中出现的用户数
func (r *UserRepo) CountByRoleId(ctx context.Context, id string) (int64, error) {
	var cnt int64
	// 使用 MySQL 的 FIND_IN_SET 来匹配以逗号分隔的 role_ids 字段
	row := r.clt.DB.QueryRowContext(ctx, "SELECT COUNT(*) FROM "+UsersTable+" WHERE FIND_IN_SET(?, role_ids)", id)
	if err := row.Scan(&cnt); err != nil {
		return 0, err
	}
	return cnt, nil
}
