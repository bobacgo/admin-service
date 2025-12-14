package apps

import (
	"log/slog"
	"net/http"
	"slices"
	"strings"

	"github.com/bobacgo/admin-service/apps/common/contextx"
	"github.com/bobacgo/admin-service/pkg/kit/hs"
)

// 白名单路径，不需要认证
var apiWhiteList = []string{
	"/api/user/login",
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 白名单路径，不需要认证
		if slices.Contains(apiWhiteList, r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		auth := r.Header.Get("Authorization")
		if auth == "" {
			slog.Error("AuthMiddleware", slog.String("error", "missing Authorization header"))
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Expect header format: "Bearer <token>"
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			slog.Error("AuthMiddleware", slog.String("error", "invalid auth scheme"))
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := hs.ParseJWT[contextx.User](parts[1])
		if err != nil {
			slog.Error("AuthMiddleware", slog.String("error", err.Error()))
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := contextx.WithUserContext(r.Context(), contextx.User{
			Account: claims.User.Account,
			RoleIds: claims.User.RoleIds,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SetCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// 这里可以设置一些公共的上下文值
		ctx = contextx.WithIPContext(ctx, r.RemoteAddr)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
