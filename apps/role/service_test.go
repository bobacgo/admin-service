package role

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

func setupRoleService(t *testing.T) (*RoleService, *http.ServeMux) {
	clt := data.NewData()
	repo := NewRoleRepo(clt)
	svc := NewRoleService(repo, validator.New())

	mux := http.NewServeMux()
	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)

	hs.RegisterService(api, "/role", svc)

	return svc, mux
}

func TestRoleGetOne(t *testing.T) {
	svc, mux := setupRoleService(t)
	ctx := context.Background()

	role := &Role{Code: "admin", Description: "Administrator", Status: 1}
	svc.repo.Create(ctx, role)

	req := httptest.NewRequest("GET", "/api/role/one?id=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestRoleGetOne: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestRoleGetOne passed, response: %s", w.Body.String())
}

func TestRoleGetList(t *testing.T) {
	svc, mux := setupRoleService(t)
	ctx := context.Background()

	svc.repo.Create(ctx, &Role{Code: "admin", Description: "Administrator", Status: 1})
	svc.repo.Create(ctx, &Role{Code: "user", Description: "User", Status: 1})

	req := httptest.NewRequest("GET", "/api/role/list", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestRoleGetList: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestRoleGetList passed, response: %s", w.Body.String())
}

func TestRolePost(t *testing.T) {
	_, mux := setupRoleService(t)

	newRole := &RoleCreateReq{Code: "guest", Description: "Guest Role", Status: 1}
	body, _ := json.Marshal(newRole)

	req := httptest.NewRequest("POST", "/api/role", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestRolePost: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestRolePost passed, response: %s", w.Body.String())
}

func TestRolePut(t *testing.T) {
	svc, mux := setupRoleService(t)
	ctx := context.Background()

	role := &Role{Code: "admin", Description: "Administrator", Status: 1}
	svc.repo.Create(ctx, role)

	updateRole := &RoleUpdateReq{ID: 1, Code: "admin_updated", Description: "Updated Administrator", Status: 1}
	body, _ := json.Marshal(updateRole)

	req := httptest.NewRequest("PUT", "/api/role", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestRolePut: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestRolePut passed, response: %s", w.Body.String())
}

func TestRoleDelete(t *testing.T) {
	_, mux := setupRoleService(t)

	req := httptest.NewRequest("DELETE", "/api/role?ids=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestRoleDelete: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestRoleDelete passed, response: %s", w.Body.String())
}
