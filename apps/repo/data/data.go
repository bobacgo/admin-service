package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultDBName = "admin_db"
)

type Client struct {
	DB *sql.DB
}

func NewData() *Client {
	dsn := buildDSNFromEnv()
	return &Client{
		DB: NewDB(dsn),
	}
}

func (d *Client) Close() error {
	if d.DB != nil {
		return d.DB.Close()
	}
	return nil
}

func NewDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

// buildDSNFromEnv builds a MySQL DSN from environment variables.
// Supported env vars:
//
//	DB_USER (default: root)
//	DB_PASS (default: empty)
//	DB_HOST (default: 127.0.0.1)
//	DB_PORT (default: 3306)
//	DB_NAME (default: admin_db)
//	DB_PARAMS (default: parseTime=true&charset=utf8mb4&loc=Local)
func buildDSNFromEnv() string {
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}
	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = "admin123"
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = defaultDBName
	}
	params := os.Getenv("DB_PARAMS")
	if params == "" {
		params = "parseTime=true&charset=utf8mb4&loc=Local"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, pass, host, port, name, params)
}
