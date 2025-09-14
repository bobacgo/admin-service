package orm

import (
	"context"
	"database/sql"
	"strings"
)

func Create(ctx context.Context, db *sql.DB, row Model) (int64, error) {
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
	stmt, err := db.PrepareContext(ctx, sqlText)
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, values...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
