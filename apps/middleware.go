package apps

import (
	"log/slog"
	"net/http"
	"slices"
	"strings"

	"github.com/bobacgo/admin-service/pkg/util"
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

		claims, err := util.ParseJWT(parts[1])
		if err != nil {
			slog.Error("AuthMiddleware", slog.String("error", err.Error()))
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		// Optionally attach user info to context for downstream handlers
		ctx := r.Context()
		ctx = WithUserContext(ctx, claims.UserID, claims.Account)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
