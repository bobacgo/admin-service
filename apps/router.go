package apps

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/bobacgo/admin-service/apps/ecode"
	"github.com/bobacgo/admin-service/apps/menu"
	"github.com/bobacgo/admin-service/apps/user"
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

func makeLoginHandler(svc *user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = new(user.LoginReq)
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

func makeLogoutHandler(svc *user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := svc.Logout(r.Context())
		if err != nil {
			response.JSON(w, response.Resp{
				Code: ecode.ErrCodeServer,
				Msg:  err.Error(),
			})
			return
		}

		response.JSON(w, response.Resp{
			Code: ecode.OK,
			Msg:  "ok",
		})
	}
}

func makeMenuTreeHandler(svc *menu.MenuService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		menuTree, err := svc.GetTree(r.Context(), nil)
		if err != nil {
			slog.Error("GetTree error", "err", err)
			response.JSON(w, response.Resp{
				Code: ecode.ErrCodeServer,
				Msg:  err.Error(),
			})
			return
		}

		response.JSON(w, response.Resp{
			Code: ecode.OK,
			Msg:  "success",
			Data: map[string]any{
				"list": menuTree,
			},
		})
	}
}

func RegisterRoutes(container *Container) http.Handler {
	mux := http.NewServeMux()

	// 创建一个处理链，先应用Cors中间件设置CORS头，再应用OptionsMiddleware处理OPTIONS请求
	public := hs.NewGroup("/", mux, hs.Logger, hs.Cors)
	public.HandleFunc("GET /health", container.api.Health)

	// api := kitweb.NewGroup("/api", mux, kitweb.Logger, kitweb.Cors, kitweb.AuthMiddleware)
	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)

	// Special handlers for login/logout
	api.HandleFunc("POST /login", makeLoginHandler(container.svc.User))
	api.HandleFunc("POST /logout", makeLogoutHandler(container.svc.User))

	// User - Auto-register routes from service methods
	userConfig := &hs.HandlerConfig{
		Validator: container.svc.Validator,
	}
	hs.RegisterServiceRoutes(api, container.svc.User, "/user", userConfig)

	// Menu - Auto-register routes from service methods
	menuConfig := &hs.HandlerConfig{
		Validator: container.svc.Validator,
	}
	hs.RegisterServiceRoutes(api, container.svc.Menu, "/menu", menuConfig)

	// Menu tree endpoint - custom route
	api.HandleFunc("GET /get-menu-list-i18n", makeMenuTreeHandler(container.svc.Menu))

	// Role - Auto-register routes from service methods
	roleConfig := &hs.HandlerConfig{
		Validator: container.svc.Validator,
	}
	hs.RegisterServiceRoutes(api, container.svc.Role, "/role", roleConfig)

	// I18n - Auto-register routes from service methods
	i18nConfig := &hs.HandlerConfig{
		Validator: container.svc.Validator,
	}
	hs.RegisterServiceRoutes(api, container.svc.I18n, "/i18n", i18nConfig)

	handlerChain := hs.Cors(OptionsMiddleware(mux))
	return handlerChain
}
