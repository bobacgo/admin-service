package apps

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/bobacgo/admin-service/apps/common/ecode"
	"github.com/bobacgo/admin-service/apps/mgr/dto"
	"github.com/bobacgo/admin-service/apps/mgr/service"
	"github.com/bobacgo/admin-service/pkg/kit/hs"
	"github.com/bobacgo/admin-service/pkg/kit/hs/response"
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

func makeLoginHandler(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = new(dto.LoginReq)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.JSON(w, response.Resp{
				Code: ecode.ErrCodeParam,
				Msg:  err.Error(),
			})
			return
		}

		data, err := svc.Login(r.Context(), req)
		if err != nil {
			response.JSON(w, response.Resp{
				Code: ecode.ErrCodeUsernameOrPassword,
				Msg:  err.Error(),
			})
			return
		}

		response.JSON(w, response.Resp{
			Code: ecode.OK,
			Msg:  "ok",
			Data: data,
		})
	}
}

func RegisterRoutes(container *Container) http.Handler {
	mux := http.NewServeMux()

	// 创建一个处理链，先应用Cors中间件设置CORS头，再应用OptionsMiddleware处理OPTIONS请求
	public := hs.NewGroup("/", mux, hs.Logger, hs.Cors)
	hs.RegisterService(public, "/", container.svc.Basic)

	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)
	api.HandleFunc("POST /login", makeLoginHandler(container.svc.User))

	hs.RegisterService(api, "/user", container.svc.User)
	hs.RegisterService(api, "/menu", container.svc.Menu)
	hs.RegisterService(api, "/role", container.svc.Role)
	hs.RegisterService(api, "/i18n", container.svc.I18n)

	return hs.Cors(OptionsMiddleware(mux))
}