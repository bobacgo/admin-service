package repo

import (
	"context"
	"strings"

	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/bobacgo/admin-service/apps/repo/model"
	"github.com/bobacgo/admin-service/pkg/util"
)

type UserRepo struct {
	clt *data.Client
}

func NewUserRepo(data *data.Client) *UserRepo {
	return &UserRepo{clt: data}
}

// 创建
func (r *UserRepo) Create(ctx context.Context, row *model.User) error {
	mapping := row.Mapping(false)

	var (
		columns = make([]string, 0, len(mapping))
		pos     = make([]string, 0, len(mapping))
		values  = make([]any, 0, len(mapping))
	)
	for k, v := range mapping {
		if k == "id" {
			continue
		}
		columns = append(columns, k)
		pos = append(pos, "?")
		values = append(values, v)
	}

	sqlText := "INSERT INTO " + row.TableName() + " (" + strings.Join(columns, ",") + ") VALUES (" + strings.Join(pos, ",") + ")"
	_, err := r.clt.DB.ExecContext(ctx, sqlText, values...)
	return err
}

func (r *UserRepo) FindOne(ctx context.Context, req *dto.GetUserReq) (*model.User, error) {
	row := new(model.User)
	mapping := row.Mapping(true)
	var (
		columns = make([]string, 0, len(mapping))
		values  = make([]any, 0, len(mapping))
	)
	for k, v := range mapping {
		columns = append(columns, k)
		values = append(values, v)
	}

	condValues := make([]interface{}, 0)
	cond := []string{
		"1 = 1", // Always true condition to simplify appending
	}
	if req.ID > 0 {
		cond = append(cond, "AND id = ?")
		condValues = append(condValues, req.ID)
	}
	if req.Account != "" {
		cond = append(cond, "AND account = ?")
		condValues = append(condValues, req.Account)
	}
	if req.Phone != "" {
		cond = append(cond, "AND phone = ?")
		condValues = append(condValues, req.Phone)
	}
	if req.Email != "" {
		cond = append(cond, "AND email = ?")
		condValues = append(condValues, req.Email)
	}

	sqlText := "SELECT " + strings.Join(columns, ",") + " FROM " + row.TableName() + " WHERE " + strings.Join(cond, " ")
	res := r.clt.DB.QueryRowContext(ctx, sqlText, condValues...)
	err := res.Scan(values...)
	return row, err
}

func (r *UserRepo) Find(ctx context.Context, req *dto.UserListReq) ([]*model.User, int64, error) {
	where := map[string]any{
		"AND account LIKE ?": req.Account + "%",
		"AND phone LIKE ?":   req.Phone + "%",
		"AND email LIKE ?":   req.Email + "%",
		"AND status IN (?)":  req.Status,
	}
	if req.Account != "" {
		where["AND account LIKE ?"] = req.Account + "%"
	}
	if req.Phone != "" {
		where["AND phone LIKE ?"] = req.Phone + "%"
	}
	if req.Email != "" {
		where["AND email LIKE ?"] = req.Email + "%"
	}
	if req.Status != "" {
		where["AND status IN (?)"] = req.Status
	}

	//return orm.FindPage[model.User](ctx, r.clt.DB, where, req.Page, req.PageSize)
	// TODO: 分页查询
	return nil, 0, nil
}

func (r *UserRepo) Update(ctx context.Context, row *model.User) error {
	mapping := row.Mapping(false)
	var (
		columns = make([]string, 0, len(mapping))
		values  = make([]any, 0, len(mapping))
	)
	for k, v := range mapping {
		if k == "id" {
			continue
		}
		if util.IsZero(v) {
			continue
		}
		columns = append(columns, k+" = ?")
		values = append(values, v)
	}
	if len(columns) == 0 {
		return nil
	}
	sqlText := "UPDATE " + row.TableName() + " SET " + strings.Join(columns, ",") + " WHERE id = ?"
	values = append(values, row.ID)
	_, err := r.clt.DB.ExecContext(ctx, sqlText, values...)
	return err
}

func (r *UserRepo) Delete(ctx context.Context, ids string) error {
	sqlText := "DELETE FROM " + model.UsersTable + " WHERE id IN (?)"
	_, err := r.clt.DB.ExecContext(ctx, sqlText, ids)
	return err
}