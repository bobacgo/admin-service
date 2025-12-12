package user

import (
	"context"
	"database/sql"
	"strings"

	"github.com/bobacgo/admin-service/apps/common/model"
	"github.com/bobacgo/admin-service/apps/common/repo"
	"github.com/bobacgo/admin-service/apps/common/repo/data"
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
	if req.Status > 0 {
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

func (r *UserRepo) Update(ctx context.Context, req *UpdateUserReq) error {
	set := M{
		Phone:           req.Phone,
		Email:           req.Email,
		model.UpdatedAt: req.UpdatedAt,
		model.Operator:  req.Operator,
	}
	_, err := UPDATE(UsersTable).SET(set).WHERE(M{repo.AND(model.Id): req.Id}).Exec(ctx, r.clt.DB)
	return err
}

func (r *UserRepo) UpdateLoginInfo(ctx context.Context, req *UpdateLoginInfoReq) error {
	set := M{
		LoginAt: req.LoginAt,
		LoginIp: req.LoginIp,
	}
	_, err := UPDATE(UsersTable).SET(set).WHERE(M{repo.AND(model.Id): req.Id}).Exec(ctx, r.clt.DB)
	return err
}

func (r *UserRepo) UpdateStatus(ctx context.Context, req *UpdateUserStatusReq) error {
	set := M{
		model.Status:    req.Status,
		model.UpdatedAt: req.UpdatedAt,
		model.Operator:  req.Operator,
	}
	_, err := UPDATE(UsersTable).SET(set).WHERE(M{repo.AND(model.Id): req.Id}).Exec(ctx, r.clt.DB)
	return err
}

func (r *UserRepo) UpdateRole(ctx context.Context, req *UpdateUserRoleReq) error {
	set := M{
		RoleIds:         req.RoleIds,
		model.UpdatedAt: req.UpdatedAt,
		model.Operator:  req.Operator,
	}
	_, err := UPDATE(UsersTable).SET(set).WHERE(M{repo.AND(model.Id): req.Id}).Exec(ctx, r.clt.DB)
	return err
}

func (r *UserRepo) UpdatePassword(ctx context.Context, req *UpdateUserPasswordReq) error {
	set := M{
		Password:        req.NewPassword,
		model.UpdatedAt: req.UpdatedAt,
		model.Operator:  req.Operator,
	}
	_, err := UPDATE(UsersTable).SET(set).WHERE(M{repo.AND(model.Id): req.Id}).Exec(ctx, r.clt.DB)
	return err
}

func (r *UserRepo) Delete(ctx context.Context, ids string) error {
	_, err := DELETE().FROM(UsersTable).WHERE(M{repo.AND_IN(model.Id): ids}).Exec(ctx, r.clt.DB)
	return err
}

// CountByRoleId 返回 role id 在 users.role_ids 字段中出现的用户数
func (r *UserRepo) CountByRoleId(ctx context.Context, id int64) (int64, error) {
	var cnt int64
	// 使用 MySQL 的 FIND_IN_SET 来匹配以逗号分隔的 role_ids 字段
	row := r.clt.DB.QueryRowContext(ctx, "SELECT COUNT(*) FROM "+UsersTable+" WHERE FIND_IN_SET(?, role_ids)", id)
	if err := row.Scan(&cnt); err != nil {
		return 0, err
	}
	return cnt, nil
}

// CountByRoleIds 返回多个角色下的用户数量，结果以角色ID为键
func (r *UserRepo) CountByRoleIds(ctx context.Context, ids []int64) (map[int64]int64, error) {
	res := make(map[int64]int64)
	if len(ids) == 0 {
		return res, nil
	}

	placeholders := make([]string, 0, len(ids))
	args := make([]any, 0, len(ids))
	for _, id := range ids {
		placeholders = append(placeholders, "?")
		args = append(args, id)
	}

	query := "SELECT r.id, COUNT(u.id) AS cnt " +
		"FROM roles r " +
		"LEFT JOIN " + UsersTable + " u ON FIND_IN_SET(r.id, u.role_ids) " +
		"WHERE r.id IN (" + strings.Join(placeholders, ",") + ") " +
		"GROUP BY r.id"

	rows, err := r.clt.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var roleID int64
		var cnt sql.NullInt64
		if err := rows.Scan(&roleID, &cnt); err != nil {
			return nil, err
		}
		if cnt.Valid {
			res[roleID] = cnt.Int64
		} else {
			res[roleID] = 0
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// 确保所有角色ID都返回结果
	for _, id := range ids {
		if _, ok := res[id]; !ok {
			res[id] = 0
		}
	}

	return res, nil
}
