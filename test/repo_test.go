package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/user"
)

func TestDB(t *testing.T) {
	Init()
}

func Init() {
	clt := data.NewData()
	defer clt.Close()

	// 创建表
	_, err := clt.DB.Exec(
		`CREATE TABLE IF NOT EXISTS users (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					account VARCHAR(255) NOT NULL,
					password VARCHAR(255) NOT NULL,
					phone VARCHAR(255),
					email VARCHAR(255),
					status INT,
					register_at INT,
					register_ip VARCHAR(255),
					login_at INT,
					login_ip VARCHAR(255),
					created_at INT DEFAULT CURRENT_TIMESTAMP,
					updated_at INT DEFAULT CURRENT_TIMESTAMP
			)`,
	)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}

	// 创建唯一索引
	_, err = clt.DB.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_account ON users(account)`)
	if err != nil {
		log.Fatal("创建索引失败:", err)
	}

	userRepo := user.NewUserRepo(clt)
	err = userRepo.Create(context.Background(), &user.User{
		Account:    "admin",
		Password:   "admin",
		Phone:      "12345678901",
		Email:      "admin@example.com",
		Status:     1,
		RegisterAt: time.Now().Unix(),
		RegisterIp: "127.0.0.1",
		LoginAt:    time.Now().Unix(),
		LoginIp:    "127.0.0.1",
	})
	if err != nil {
		log.Fatal("创建用户失败:", err)
	}
}
