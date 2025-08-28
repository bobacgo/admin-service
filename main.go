package main

import (
	"log/slog"

	kitweb "github.com/bobacgo/admin-service/pkg/kit-web"
)

func main() {
	// 注册路由
	container := NewContainer()
	handler := RegisterRoutes(container)
	server := kitweb.New(":8080")
	// 使用处理链而不是直接使用mux
	server.SetHandler(handler)
	if err := server.Run(); err != nil {
		slog.Error("Server failed", "error", err)
	}
}