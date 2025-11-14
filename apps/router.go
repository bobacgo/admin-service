package apps

import (
	"log/slog"
	"net/http"

	"github.com/bobacgo/admin-service/pkg/kit/hs"
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
	public := hs.NewGroup("/", mux, hs.Logger, hs.Cors)
	public.HandleFunc("GET /health", container.api.Health)
	public.HandleFunc("POST /api/login", container.api.User.Login)

	// api := kitweb.NewGroup("/api", mux, kitweb.Logger, kitweb.Cors, kitweb.AuthMiddleware)
	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)
	api.HandleFunc("POST /logout", container.api.User.Logout)

	// User
	api.HandleFunc("GET /user/info", container.api.User.GetInfo)
	api.HandleFunc("GET /user/list", container.api.User.List)
	api.HandleFunc("POST /user/add", container.api.User.Create)
	api.HandleFunc("PUT /user/update", container.api.User.Update)
	api.HandleFunc("DELETE /user/delete", container.api.User.Delete)

	// Menu
	api.HandleFunc("GET /menu/get", container.api.Menu.Get)
	api.HandleFunc("GET /menu/list", container.api.Menu.GetList)
	api.HandleFunc("POST /menu/add", container.api.Menu.Create)
	api.HandleFunc("PUT /menu/update", container.api.Menu.Update)
	api.HandleFunc("DELETE /menu/delete", container.api.Menu.Delete)
	api.HandleFunc("GET /get-menu-list-i18n", container.api.Menu.Tree)

	// Role
	api.HandleFunc("GET /role/get", container.api.Role.Get)
	api.HandleFunc("GET /role/list", container.api.Role.List)
	api.HandleFunc("POST /role/add", container.api.Role.Create)
	api.HandleFunc("PUT /role/update", container.api.Role.Update)
	api.HandleFunc("DELETE /role/delete", container.api.Role.Delete)

	// I18n
	api.HandleFunc("GET /i18n/get", container.api.I18n.Get)
	api.HandleFunc("GET /i18n/list", container.api.I18n.List)
	api.HandleFunc("POST /i18n/add", container.api.I18n.Create)
	api.HandleFunc("PUT /i18n/update", container.api.I18n.Update)
	api.HandleFunc("DELETE /i18n/delete", container.api.I18n.Delete)

	handlerChain := hs.Cors(OptionsMiddleware(mux))
	return handlerChain
}
