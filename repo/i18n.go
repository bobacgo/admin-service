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

type I18nRepo struct {
	clt *data.Client
}

func NewI18nRepo(clt *data.Client) *I18nRepo {
	return &I18nRepo{clt: clt}
}

func (r *I18nRepo) FindOne(ctx context.Context, req *dto.GetI18nReq) (*model.I18n, error) {
	row := new(model.I18n)
	mapping := row.Mapping(true)
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
	if req.Key != "" {
		cond = append(cond, "AND key = ?")
		condValues = append(condValues, req.Key)
	}
	if req.Lang != "" {
		cond = append(cond, "AND lang = ?")
		condValues = append(condValues, req.Lang)
	}

	sqlText := "SELECT " + strings.Join(columns, ",") + " FROM " + row.TableName() + " WHERE " + strings.Join(cond, " ")
	res := r.clt.DB.QueryRowContext(ctx, sqlText, condValues...)
	err := res.Scan(values...)
	return row, err
}

func (r *I18nRepo) Find(ctx context.Context, req *dto.I18nListReq) ([]*model.I18n, error) {
	u := new(model.I18n)
	mapping := u.Mapping(true)

	columns := make([]string, 0, len(mapping))
	for col := range mapping {
		columns = append(columns, col)
	}

	cond := []string{
		"1 = 1", // Always true condition to simplify appending
	}
	condValues := make([]interface{}, 0)

	if req.Class != "" {
		cond = append(cond, "AND class = ?")
		condValues = append(condValues, req.Class)
	}
	if req.Lang != "" {
		cond = append(cond, "AND lang = ?")
		condValues = append(condValues, req.Lang)
	}
	if req.Key != "" {
		cond = append(cond, "AND key = ?")
		condValues = append(condValues, req.Key)
	}

	sqlText := "SELECT " + strings.Join(columns, ",") + " FROM " + u.TableName() + " WHERE " + strings.Join(cond, " ") + " ORDER BY id DESC"
	rows, err := r.clt.DB.QueryContext(ctx, sqlText, condValues...)
	if err != nil {
		return nil, fmt.Errorf("db.Query: %w", err)
	}
	defer rows.Close()

	list := make([]*model.I18n, 0)
	for rows.Next() {
		record := new(model.I18n)
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

func (r *I18nRepo) Create(ctx context.Context, row *model.I18n) error {
	mapping := row.Mapping(false)

	var (
		columns = make([]string, 0, len(mapping))
		pos     = make([]string, 0, len(mapping))
		values  = make([]any, 0, len(mapping))
	)
	for k, fn := range mapping {
		if k == "id" {
			continue
		}
		columns = append(columns, k)
		pos = append(pos, "?")
		values = append(values, fn(row))
	}

	sqlText := "INSERT INTO " + row.TableName() + " (" + strings.Join(columns, ",") + ") VALUES (" + strings.Join(pos, ",") + ")"
	res, err := r.clt.DB.ExecContext(ctx, sqlText, values...)
	if err != nil {
		return err
	}
	row.ID, err = res.LastInsertId()
	return err
}

func (r *I18nRepo) Update(ctx context.Context, row *model.I18n) error {
	mapping := row.Mapping(false)
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

func (r *I18nRepo) Delete(ctx context.Context, ids string) error {
	sqlText := "DELETE FROM " + new(model.User).TableName() + " WHERE id IN (?)"
	_, err := r.clt.DB.ExecContext(ctx, sqlText, ids)
	return err
}
