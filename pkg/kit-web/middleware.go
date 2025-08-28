package kitweb

import (
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/bobacgo/admin-service/pkg/kit-web/response"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.ErrorContext(r.Context(), "Recovered from panic", slog.Any("error", err))
				debug.PrintStack()
				response.JSON(w, map[string]any{
					"code": 1,
					"msg":  "Internal Server Error",
				})
			}
		}()

		now := time.Now()
		next.ServeHTTP(w, r)
		slog.InfoContext(r.Context(), "Request received",
			slog.String("time", time.Since(now).String()),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("remote_addr", r.RemoteAddr),
		)
	})
}

// Cors 设置跨域请求所需的响应头
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置CORS响应头
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// 继续处理请求
		next.ServeHTTP(w, r)
	})
}

// curl -H "Authorization: 123" localhost:8080/user/123
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		// check auth
		if auth == "" {
			slog.Error("AuthMiddleware", slog.String("error", "no auth"))
			http.Error(w, "no auth", http.StatusUnauthorized)
			return
		}
		slog.Debug("AuthMiddleware", slog.String("auth", auth))
		next.ServeHTTP(w, r)
	})
}