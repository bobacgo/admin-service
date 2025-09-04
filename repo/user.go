package repo

import (
	"context"
	"fmt"
	"strings"

	kitorm "github.com/bobacgo/admin-service/pkg/kit-orm"
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
	row := new(model.User)
	mapping := row.Mapping(true)

	columns := make([]string, 0, len(mapping))
	values := make([]any, 0, len(mapping))
	for k, v := range mapping {
		columns = append(columns, k)
		values = append(values, v)
	}

	condValues := make([]interface{}, 0)
	cond := []string{
		"1 = 1", // Always true condition to simplify appending
	}
	if req.Account != "" {
		cond = append(cond, "AND account LIKE ?")
		condValues = append(condValues, req.Account+"%")
	}
	if req.Phone != "" {
		cond = append(cond, "AND phone LIKE ?")
		condValues = append(condValues, req.Phone+"%")
	}
	if req.Email != "" {
		cond = append(cond, "AND email LIKE ?")
		condValues = append(condValues, req.Email+"%")
	}
	if req.Status != "" {
		cond = append(cond, "AND status IN (?)")
		condValues = append(condValues, req.Status)
	}

	var count int64

	// 分页
	offset, limit := req.PageReq.Limit()
	if offset > 0 { // 需要分页
		// 统计
		var err error
		count, err = kitorm.Count(ctx, r.clt.DB, row.TableName(), cond, condValues)
		if err != nil {
			return nil, 0, fmt.Errorf("db.Count: %w", err)
		}
		if count == 0 { // 没有数据
			return make([]*model.User, 0), 0, nil
		}

		cond = append(cond, "LIMIT ?, ?")
		condValues = append(condValues, offset, limit)
	}

	sqlText := "SELECT " + strings.Join(columns, ",") + " FROM " + row.TableName() + " WHERE " + strings.Join(cond, " ") + " ORDER BY id DESC"
	rows, err := r.clt.DB.QueryContext(ctx, sqlText, condValues...)
	if err != nil {
		return nil, 0, fmt.Errorf("db.Query: %w", err)
	}
	defer rows.Close()

	list := make([]*model.User, 0)
	for rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, 0, fmt.Errorf("rows.Scan: %w", err)
		}
		cp := *row
		list = append(list, &cp)
	}
	if offset == 0 { // 不需要分页
		count = int64(len(list))
	}
	return list, count, nil
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
