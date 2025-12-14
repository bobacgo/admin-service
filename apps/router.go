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
	public := hs.NewGroup("/", mux, hs.Logger, hs.Cors, SetCtx)
	hs.RegisterService(public, "/", container.svc.Basic)

	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors, SetCtx, AuthMiddleware)

	hs.RegisterService(api, "/user", container.svc.User)
	hs.RegisterService(api, "/menu", container.svc.Menu)
	hs.RegisterService(api, "/role", container.svc.Role)
	hs.RegisterService(api, "/i18n", container.svc.I18n)

	return hs.Cors(OptionsMiddleware(mux))
}
