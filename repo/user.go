package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/bobacgo/admin-service/pkg/util"
	"github.com/bobacgo/admin-service/repo/data"
	"github.com/bobacgo/admin-service/repo/dto"
	"github.com/bobacgo/admin-service/repo/model"
)

type UserRepo struct {
	clt *data.Client
}

func NewUserRepo(data *data.Client) *UserRepo {
	return &UserRepo{clt: data}
}

// 创建
func (r *UserRepo) Create(ctx context.Context, row *model.User) error {
	mapping := row.Mapping()

	var (
		columns = make([]string, 0, len(mapping))
		pos     = make([]string, 0, len(mapping))
		values  = make([]any, 0, len(mapping))
	)
	for k, fn := range mapping {
		columns = append(columns, k)
		pos = append(pos, "?")
		values = append(values, fn(row))
	}

	sqlText := "INSERT INTO " + row.TableName() + " (" + strings.Join(columns, ",") + ") VALUES (" + strings.Join(pos, ",") + ")"
	_, err := r.clt.DB.ExecContext(ctx, sqlText, values...)
	return err
}

func (r *UserRepo) FindOne(ctx context.Context, req *dto.GetUserReq) (*model.User, error) {
	row := new(model.User)
	mapping := row.MappingSelect()
	var (
		columns = make([]string, 0, len(mapping))
		values  = make([]any, 0, len(mapping))
	)
	for k, fn := range mapping {
		columns = append(columns, k)
		values = append(values, fn(row))
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

func (r *UserRepo) Find(ctx context.Context, req *dto.UserListReq) ([]*model.User, error) {
	u := new(model.User)
	mapping := u.MappingSelect()

	columns := make([]string, 0, len(mapping))
	for col := range mapping {
		columns = append(columns, col)
	}

	cond := []string{
		"1 = 1", // Always true condition to simplify appending
	}
	// if registerTime > 0 {
	// 	cond = append(cond, "AND registerTime > ?")
	// }

	sqlText := "SELECT " + strings.Join(columns, ",") + " FROM " + u.TableName() + " WHERE " + strings.Join(cond, " ") + " ORDER BY id DESC"
	rows, err := r.clt.DB.QueryContext(ctx, sqlText)
	if err != nil {
		return nil, fmt.Errorf("db.Query: %w", err)
	}
	defer rows.Close()

	var list []*model.User
	for rows.Next() {
		record := new(model.User)
		values := make([]any, 0, len(columns))
		for _, col := range columns {
			values = append(values, mapping[col](record))
		}
		if err = rows.Scan(values...); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		list = append(list, record)
	}
	return list, nil
}

func (r *UserRepo) Update(ctx context.Context, row *model.User) error {
	mapping := row.Mapping()
	var (
		columns = make([]string, 0, len(mapping))
		values  = make([]any, 0, len(mapping))
	)
	for k, fn := range mapping {
		v := fn(row)
		if util.IsZero(v) {
			continue
		}
		columns = append(columns, k+" = ?")
		values = append(values, v)
	}
	sqlText := "UPDATE " + row.TableName() + " SET " + strings.Join(columns, ",") + " WHERE id = ?"
	values = append(values, row.ID)
	_, err := r.clt.DB.ExecContext(ctx, sqlText, values...)
	return err
}

func (r *UserRepo) Delete(ctx context.Context, ids string) error {
	sqlText := "DELETE FROM " + new(model.User).TableName() + " WHERE id IN (?)"
	_, err := r.clt.DB.ExecContext(ctx, sqlText, ids)
	return err
}