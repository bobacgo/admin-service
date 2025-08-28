package main

import (
	"log/slog"
	"net/http"

	"github.com/bobacgo/admin-service/pkg/kit-web"
)

// OptionsMiddleware 处理所有OPTIONS请求，确保跨域预检请求能够正确响应
func OptionsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 如果是OPTIONS请求，直接处理并返回200 OK
		if r.Method == http.MethodOptions {
			slog.Info("Handling OPTIONS request", "path", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			return
		}
		// 否则，继续处理请求
		next.ServeHTTP(w, r)
	})
}

func RegisterRoutes(container *Container) http.Handler {
	mux := http.NewServeMux()

	// 创建一个处理链，先应用Cors中间件设置CORS头，再应用OptionsMiddleware处理OPTIONS请求
	handlerChain := kitweb.Cors(OptionsMiddleware(mux))
	public := kitweb.NewGroup("/", mux, kitweb.Logger, kitweb.Cors)
	public.HandleFunc("GET /health", container.api.Health)
	public.HandleFunc("POST /api/login", container.api.User.Login)

	// api := kitweb.NewGroup("/api", mux, kitweb.Logger, kitweb.Cors, kitweb.AuthMiddleware)
	api := kitweb.NewGroup("/api", mux, kitweb.Logger, kitweb.Cors)
	api.HandleFunc("POST /logout", container.api.User.Logout)
	api.HandleFunc("GET /user/info", container.api.User.GetInfo)
	api.HandleFunc("GET /user/list", container.api.User.List)
	api.HandleFunc("POST /user/add", container.api.User.Create)
	api.HandleFunc("PUT /user/update", container.api.User.Update)
	api.HandleFunc("DELETE /user/delete", container.api.User.Delete)
	api.HandleFunc("GET /get-menu-list-i18n", container.api.Menu.GetList)

	return handlerChain
}