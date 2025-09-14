package orm

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func FindOne[T any, PT interface {
	*T
	Model
}](ctx context.Context, db *sql.DB, where map[string]any) (PT, error) {
	row := PT(new(T)) // 显式转换成 PT
	mapping := row.Mapping(true)
	var (
		columns = make([]string, 0, len(mapping))
		values  = make([]any, 0, len(mapping))
	)
	for k, v := range mapping {
		columns = append(columns, k)
		values = append(values, v)
	}

	condValues := make([]any, 0)
	cond := []string{
		"1 = 1", // Always true condition to simplify appending
	}

	for k, v := range where {
		cond = append(cond, k)
		condValues = append(condValues, v)
	}

	sqlText := "SELECT " + strings.Join(columns, ",") + " FROM " + row.TableName() + " WHERE " + strings.Join(cond, " ")
	res := db.QueryRowContext(ctx, sqlText, condValues...)
	err := res.Scan(values...)
	return row, err
}

func Find[T any, PT interface {
	*T
	Model
}](ctx context.Context, db *sql.DB, where map[string]any) ([]PT, error) {
	row := PT(new(T)) // 显式转换成 PT
	mapping := row.Mapping(true)

	columns := make([]string, 0, len(mapping))
	values := make([]any, 0, len(mapping))
	for k, v := range mapping {
		columns = append(columns, k)
		values = append(values, v)
	}

	cond := []string{
		"1 = 1", // Always true condition to simplify appending
	}
	condValues := make([]interface{}, 0)

	for k, v := range where {
		cond = append(cond, k)
		condValues = append(condValues, v)
	}

	sqlText := "SELECT " + strings.Join(columns, ",") + " FROM " + row.TableName() + " WHERE " + strings.Join(cond, " ") + " ORDER BY id DESC"
	rows, err := db.QueryContext(ctx, sqlText, condValues...)
	if err != nil {
		return nil, fmt.Errorf("stmt.Query: %w", err)
	}
	defer rows.Close()

	list := make([]PT, 0)
	for rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		cp := *row
		list = append(list, &cp)
	}
	return list, nil
}

func FindPage[T any, PT interface {
	*T
	Model
}](ctx context.Context, db *sql.DB, where map[string]any, page, pageSize int) ([]PT, int64, error) {
	list := make([]PT, 0)

	if page <= 0 || pageSize <= 0 {
		return list, 0, nil
	}

	row := PT(new(T)) // 显式转换成 PT

	// 过滤条件
	cond := []string{"1 = 1"}
	condValues := make([]interface{}, 0)
	for k, v := range where {
		cond = append(cond, k)
		condValues = append(condValues, v)
	}

	var total sql.NullInt64

	sqlText := "SELECT COUNT(*) FROM " + row.TableName() + " WHERE " + strings.Join(cond, " ")
	if err := db.QueryRowContext(ctx, sqlText, condValues...).
		Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("COUNT: %w", err)
	}
	if total.Int64 == 0 {
		return list, 0, nil
	}

	mapping := row.Mapping(true)

	// 查询字段和对应的字段映射
	columns := make([]string, 0, len(mapping))
	values := make([]any, 0, len(mapping))
	for k, v := range mapping {
		columns = append(columns, k)
		values = append(values, v)
	}

	// 分页
	condValues = append(condValues, (page-1)*pageSize, pageSize)

	sqlText = "SELECT " + strings.Join(columns, ",") + " FROM " + row.TableName() + " WHERE " + strings.Join(cond, " ") + " ORDER BY id DESC LIMIT ?, ?"
	rows, err := db.QueryContext(ctx, sqlText, condValues...)
	if err != nil {
		return nil, 0, fmt.Errorf("stmt.Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, 0, fmt.Errorf("rows.Scan: %w", err)
		}
		cp := *row
		list = append(list, &cp)
	}
	return list, total.Int64, nil
}
