package main

import (
	"log/slog"

	"github.com/bobacgo/admin-service/apps"
	"github.com/bobacgo/admin-service/pkg/kit/hs"
)

func main() {
	// 注册路由
	container := apps.NewContainer()
	handler := apps.RegisterRoutes(container)
	server := hs.New(":8080")
	// 使用处理链而不是直接使用mux
	server.SetHandler(handler)
	if err := server.Run(); err != nil {
		slog.Error("Server failed", "error", err)
	}
}
