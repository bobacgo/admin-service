package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/user"
	"github.com/bobacgo/admin-service/pkg/kit/hs"
	"github.com/go-playground/validator/v10"
)

// setupUserService 设置用户服务和路由
func setupUserService(t *testing.T) *http.ServeMux {
	// 初始化数据源
	clt := data.NewData()
	repo := user.NewUserRepo(clt)
	svc := user.NewUserService(repo, validator.New())

	// 创建路由
	mux := http.NewServeMux()
	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)

	// 注册路由
	hs.RegisterService(api, "/user", svc)
	return mux
}

func TestUserGetOne(t *testing.T) {
	mux := setupUserService(t)

	req := httptest.NewRequest("GET", "/api/user/one?id=2", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestUserGetList(t *testing.T) {
	mux := setupUserService(t)

	req := httptest.NewRequest("GET", "/api/user/list", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestUserPost(t *testing.T) {
	mux := setupUserService(t)

	// 测试 POST /api/user
	newUser := &user.User{Account: "admin", Password: "admin", Status: 1}
	body, _ := json.Marshal(newUser)

	req := httptest.NewRequest("POST", "/api/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestUserPut(t *testing.T) {
	mux := setupUserService(t)

	// 测试 PUT /api/user
	updateUser := &user.User{Account: "user1_updated", Status: 1}
	body, _ := json.Marshal(updateUser)

	req := httptest.NewRequest("PUT", "/api/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestUserDelete(t *testing.T) {
	mux := setupUserService(t)

	// 测试 DELETE /api/user?ids=1
	req := httptest.NewRequest("DELETE", "/api/user?ids=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}
