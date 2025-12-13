package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobacgo/admin-service/apps/common/repo/data"
	"github.com/bobacgo/admin-service/apps/mgr/dto"
	"github.com/bobacgo/admin-service/apps/mgr/repo"
	"github.com/bobacgo/admin-service/apps/mgr/service"
	"github.com/bobacgo/admin-service/pkg/kit/hs"
	"github.com/go-playground/validator/v10"
)

func setupRoleService(t *testing.T) *http.ServeMux {
	clt := data.NewData()
	roleRepo := repo.NewRoleRepo(clt)
	// create user repo for tests
	userRepo := repo.NewUserRepo(clt)
	svc := service.NewRoleService(roleRepo, nil, userRepo, validator.New())

	mux := http.NewServeMux()
	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)

	hs.RegisterService(api, "/role", svc)

	return mux
}

func TestRoleGetOne(t *testing.T) {
	mux := setupRoleService(t)

	req := httptest.NewRequest("GET", "/api/role/one?id=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestRoleGetList(t *testing.T) {
	mux := setupRoleService(t)

	req := httptest.NewRequest("GET", "/api/role/list", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestRolePost(t *testing.T) {
	mux := setupRoleService(t)

	newRole := &dto.RoleCreateReq{RoleName: "guest", Description: "Guest Role", Status: 1}
	body, _ := json.Marshal(newRole)

	req := httptest.NewRequest("POST", "/api/role", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestRolePut(t *testing.T) {
	mux := setupRoleService(t)

	updateRole := &dto.RoleUpdateReq{ID: 1, RoleName: "admin_updated", Description: "Updated Administrator", Status: 1}
	body, _ := json.Marshal(updateRole)

	req := httptest.NewRequest("PUT", "/api/role", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestRoleDelete(t *testing.T) {
	mux := setupRoleService(t)

	req := httptest.NewRequest("DELETE", "/api/role?ids=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}
