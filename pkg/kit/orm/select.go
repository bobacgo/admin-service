package orm

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"testing"
)

type Select struct {
	dbCommon
	cols []string // 查询字段
	res  []any    // 查询映射结果字段

	row  Model
	rows []Model

	where   []string // 查询语法条件 例如：["AND id = ?", "OR account = ?"]
	groupBy []string
	having  string
	orderBy []string
	limit   int
	offset  int
}

// select a, b from ab where a = 1 group by a order by a limit 1, 2
func (d *Select) slect(rows []Model) {
	switch len(rows) {
	case 0:
		d.err = errors.New("model is nil")
	case 1:
		d.row = rows[0]
		mapping := d.row.Mapping(true)
		for k, v := range mapping {
			d.cols = append(d.cols, k)
			d.res = append(d.res, v)
		}
	default:
		d.rows = rows
	}
}

func (d *Select) FROM(table string) *Select {
	if table == "" {
		d.err = errors.New("table name is empty")
	}
	d.table = table
	return d
}

func (d *Select) WHERE(where map[string]any) *Select {
	for k, v := range where {
		d.where = append(d.where, k)
		d.args = append(d.args, v)
	}
	return d
}

func (d *Select) GROUP_BY(cols ...string) *Select {
	d.groupBy = cols
	return d
}

func (d *Select) HAVING(text string) *Select {
	d.having = text
	return d
}

func (d *Select) ORDER_BY(orders ...string) *Select {
	d.orderBy = orders
	return d
}

func (d *Select) LIMIT(limit int) *Select {
	d.limit = limit
	return d
}

func (d *Select) OFFSET(offset int) *Select {
	d.offset = offset
	return d
}

func (d *Select) COUNT(col string, v any) Model {
	if col == "" {
		col = "*"
	}
	text := fmt.Sprintf("COUNT(%s)", col)
	d.cols = append(d.cols, text)
	d.res = append(d.res, v)
	return &model{}
}

func (d *Select) SUM(col string, v any) Model {
	if col == "" {
		col = "*"
	}
	text := fmt.Sprintf("SUM(%s)", col)
	d.cols = append(d.cols, text)
	d.res = append(d.res, v)
	return &model{}
}

func (d *Select) Builder() string {
	return "" // TODO xxx
}

// 查到单个
func (d *Select) QueryRow(ctx context.Context) error {
	if d.err != nil {
		return d.err
	}
	sqlText := d.queryBuilder(false)
	if d.debug {
		slog.InfoContext(ctx, "sql text", "sql", sqlText, "args", d.args)
	}
	return d.db.QueryRowContext(ctx, sqlText, d.args...).Scan(d.res...)
}

// 查询多个
func (d *Select) Query(ctx context.Context) error {
	if d.err != nil {
		return d.err
	}

	sqlText := d.queryBuilder(true)
	if d.debug {
		slog.InfoContext(ctx, "sql text", "sql", sqlText, "args", d.args)
	}
	rows, err := d.db.QueryContext(ctx, sqlText, d.args...)
	if err != nil {
		return fmt.Errorf("stmt.Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(d.res...); err != nil {
			return fmt.Errorf("rows.Scan: %w", err)
		}
		cp := d.row // TODO 这里拷贝
		d.rows = append(d.rows, cp)
	}
	return rows.Err()
}

func (d *Select) queryBuilder(isMulti bool) string {
	sqlText := "SELECT " + strings.Join(d.cols, ", ") + " FROM " + d.table
	if len(d.where) > 0 {
		sqlText += " WHERE " + strings.TrimLeft(strings.Join(d.where, " "), "AND ")
	}
	if len(d.groupBy) > 0 {
		sqlText += " GROUP BY " + strings.Join(d.groupBy, ", ")
	}
	if d.having != "" {
		sqlText += " HAVING " + d.having
	}
	if len(d.orderBy) > 0 {
		sqlText += " ORDER BY " + strings.Join(d.orderBy, ", ")
	}

	if !isMulti {
		sqlText += " LIMIT 1"
	} else {
		if d.limit > 0 {
			sqlText += " LIMIT " + strconv.Itoa(d.limit)
		}
		if d.offset > 0 {
			sqlText += " OFFSET " + strconv.Itoa(d.offset)
		}
	}
	return sqlText
}

func TestSelect(t *testing.T) {
	db := NewDB(nil)
	// SELECT a, b FROM xx WHERE id = 1 GROUP BY a HAVING a > 0 ORDER BY a desc, b LIMIT 1, 1
	err := db.SELECT(&TestModel{}).
		FROM("xx").
		WHERE(map[string]any{"AND id = ?": 1}).
		GROUP_BY("a").
		HAVING("a > 0").
		ORDER_BY("a desc", "b").
		OFFSET(1).LIMIT(1).
		QueryRow(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
