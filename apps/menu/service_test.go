package menu

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

func setupMenuService(t *testing.T) (*MenuService, *http.ServeMux) {
	clt := data.NewData()
	repo := NewMenuRepo(clt)
	svc := NewMenuService(repo, validator.New())

	mux := http.NewServeMux()
	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)

	hs.RegisterService(api, "/menu", svc)

	return svc, mux
}

func TestMenuGetOne(t *testing.T) {
	svc, mux := setupMenuService(t)
	ctx := context.Background()

	menu := &Menu{Path: "/users", Name: "Users", Component: "Users"}
	svc.repo.Create(ctx, menu)

	req := httptest.NewRequest("GET", "/api/menu/one?id=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestMenuGetOne: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestMenuGetOne passed, response: %s", w.Body.String())
}

func TestMenuGetList(t *testing.T) {
	svc, mux := setupMenuService(t)
	ctx := context.Background()

	svc.repo.Create(ctx, &Menu{Path: "/menu1", Name: "Menu1"})
	svc.repo.Create(ctx, &Menu{Path: "/menu2", Name: "Menu2"})

	req := httptest.NewRequest("GET", "/api/menu/list", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestMenuGetList: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestMenuGetList passed, response: %s", w.Body.String())
}

func TestMenuPost(t *testing.T) {
	_, mux := setupMenuService(t)

	newMenu := &MenuCreateReq{Path: "/newmenu", Name: "New Menu", Component: "NewMenu"}
	body, _ := json.Marshal(newMenu)

	req := httptest.NewRequest("POST", "/api/menu", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestMenuPost: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestMenuPost passed, response: %s", w.Body.String())
}

func TestMenuPut(t *testing.T) {
	svc, mux := setupMenuService(t)
	ctx := context.Background()

	menu := &Menu{Path: "/menu1", Name: "Menu1"}
	svc.repo.Create(ctx, menu)

	updateMenu := &MenuUpdateReq{Path: "/menu1_updated", Name: "Menu1 Updated"}
	body, _ := json.Marshal(updateMenu)

	req := httptest.NewRequest("PUT", "/api/menu", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestMenuPut: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestMenuPut passed, response: %s", w.Body.String())
}

func TestMenuDelete(t *testing.T) {
	_, mux := setupMenuService(t)

	req := httptest.NewRequest("DELETE", "/api/menu?ids=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestMenuDelete: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestMenuDelete passed, response: %s", w.Body.String())
}

func TestMenuGetTree(t *testing.T) {
	svc, mux := setupMenuService(t)
	ctx := context.Background()

	svc.repo.Create(ctx, &Menu{Path: "/users", Name: "Users"})
	svc.repo.Create(ctx, &Menu{Path: "/roles", Name: "Roles"})

	req := httptest.NewRequest("GET", "/api/menu/tree", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestMenuGetTree: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestMenuGetTree passed, response: %s", w.Body.String())
}
