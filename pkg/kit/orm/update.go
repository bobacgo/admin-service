package orm

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/bobacgo/admin-service/pkg/util"
)

func Update(ctx context.Context, db *sql.DB, id int64, row Model) (int64, error) {
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

	if len(columns) == 0 { // 没有更新字段
		return 0, nil
	}

	values = append(values, id)

	sqlText := "UPDATE " + row.TableName() + " SET " + strings.Join(columns, ",") + " WHERE id = ?"
	stmt, err := db.PrepareContext(ctx, sqlText)
	if err != nil {
		return 0, fmt.Errorf("sql prepare error: %w", err)
	}
	res, err := stmt.ExecContext(ctx, values...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
