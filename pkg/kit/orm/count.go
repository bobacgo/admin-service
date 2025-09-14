package orm

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func Count(ctx context.Context, db *sql.DB, tableName string, where map[string]any) (int64, error) {
	cond := []string{
		"1 = 1", // Always true condition to simplify appending
	}
	condValues := make([]any, 0)
	for k, v := range where {
		cond = append(cond, k)
		condValues = append(condValues, v)
	}

	var count sql.NullInt64

	sqlText := "SELECT COUNT(*) FROM " + tableName + " WHERE " + strings.Join(cond, " ")
	res := db.QueryRowContext(ctx, sqlText, condValues...)
	if err := res.Scan(&count); err != nil {
		return 0, fmt.Errorf("sql count error: %w", err)
	}
	return count.Int64, nil
}
