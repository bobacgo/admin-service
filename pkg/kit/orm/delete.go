package orm

import (
	"context"
	"database/sql"
)

func Delete(ctx context.Context, db *sql.DB, tableName, ids string) (int64, error) {
	sqlText := "DELETE FROM " + tableName + " WHERE id IN (?)"
	stmt, err := db.PrepareContext(ctx, sqlText)
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, ids)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
