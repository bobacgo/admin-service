package kitorm

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func Count(ctx context.Context, db *sql.DB, tableName string, cond []string, condValues []any) (int64, error) {
	var count sql.NullInt64
	sqlText := "SELECT COUNT(*) FROM " + tableName + " WHERE " + strings.Join(cond, " ")
	res := db.QueryRowContext(ctx, sqlText, condValues...)
	if err := res.Scan(&count); err != nil {
		return 0, fmt.Errorf("sql count error: %w", err)
	}
	return count.Int64, nil
}
