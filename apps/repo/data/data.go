package data

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // 引入 pure Go sqlite
)

const (
	DBName = "admin.db"
)

type Client struct {
	DB *sql.DB
}

func NewData() *Client {
	return &Client{
		DB: NewDB(DBName),
	}
}

func (d *Client) Close() error {
	if d.DB != nil {
		return d.DB.Close()
	}
	return nil
}

func NewDB(dataSourceName string) *sql.DB {
	// 打开数据库，不存在会自动创建
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
