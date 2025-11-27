package user

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/pkg/kit/hs"
	"github.com/go-playground/validator/v10"
)

// setupUserService 设置用户服务和路由
func setupUserService(t *testing.T) (*UserService, *http.ServeMux) {
	// 初始化数据源
	clt := data.NewData()
	repo := NewUserRepo(clt)
	svc := NewUserService(repo, validator.New())

	// 创建路由
	mux := http.NewServeMux()
	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)

	// 注册路由
	hs.RegisterService(api, "/user", svc)

	return svc, mux
}

func TestUserGetOne(t *testing.T) {
	svc, mux := setupUserService(t)
	ctx := context.Background()

	// 创建测试用户
	user := &User{Account: "testuser", Password: "123456", Status: 1}
	svc.repo.Create(ctx, user)

	// 测试 GET /api/user/one?id=1
	req := httptest.NewRequest("GET", "/api/user/one?id=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestUserGetOne: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestUserGetOne passed, response: %s", w.Body.String())
}

func TestUserGetList(t *testing.T) {
	svc, mux := setupUserService(t)
	ctx := context.Background()

	// 创建测试用户
	svc.repo.Create(ctx, &User{Account: "user1", Status: 1})
	svc.repo.Create(ctx, &User{Account: "user2", Status: 1})

	// 测试 GET /api/user/list
	req := httptest.NewRequest("GET", "/api/user/list", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestUserGetList: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestUserGetList passed, response: %s", w.Body.String())
}

func TestUserPost(t *testing.T) {
	_, mux := setupUserService(t)

	// 测试 POST /api/user
	newUser := &User{Account: "newuser", Password: "123456", Status: 1}
	body, _ := json.Marshal(newUser)

	req := httptest.NewRequest("POST", "/api/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestUserPost: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestUserPost passed, response: %s", w.Body.String())
}

func TestUserPut(t *testing.T) {
	svc, mux := setupUserService(t)
	ctx := context.Background()

	// 创建初始用户
	user := &User{Account: "user1", Status: 1}
	svc.repo.Create(ctx, user)

	// 测试 PUT /api/user
	updateUser := &User{Account: "user1_updated", Status: 1}
	body, _ := json.Marshal(updateUser)

	req := httptest.NewRequest("PUT", "/api/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestUserPut: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestUserPut passed, response: %s", w.Body.String())
}

func TestUserDelete(t *testing.T) {
	_, mux := setupUserService(t)

	// 测试 DELETE /api/user?ids=1
	req := httptest.NewRequest("DELETE", "/api/user?ids=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestUserDelete: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestUserDelete passed, response: %s", w.Body.String())
}
