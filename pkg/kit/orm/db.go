package orm

import (
	"database/sql"
)

type dbCommon struct {
	db    *sql.DB
	debug bool
	err   error

	sql   string // 需要执行的 sql 语句
	table string // 表名
	args  []any  // 占位符对应参数
}

type DB struct {
	dbCommon

	slect  *Select
	insert *Insert
	update *Update
	delete *Delete
}

func NewDB(db *sql.DB) *DB {
	return &DB{
		dbCommon: dbCommon{
			db: db,
		},
	}
}

func (d *DB) Debug() *DB {
	d.debug = true
	return d
}

func (d *DB) SELECT(rows ...Model) *Select {
	d.slect = &Select{
		dbCommon: d.dbCommon,
	}

	d.slect.slect(rows)
	return d.slect
}

func (d *DB) INSERT(rows ...Model) *Insert {
	d.insert = &Insert{
		dbCommon: d.dbCommon,
	}
	d.insert.insert(rows)
	return d.insert
}

func (d *DB) UPDATE(tableName string) *Update {
	d.update = &Update{
		dbCommon: d.dbCommon,
	}
	d.dbCommon.table = tableName
	return d.update
}

func (d *DB) DELETE() *Delete {
	d.delete = &Delete{
		dbCommon: d.dbCommon,
	}
	return d.delete
}
