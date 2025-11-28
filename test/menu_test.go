package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobacgo/admin-service/apps/menu"
	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/pkg/kit/hs"
	"github.com/go-playground/validator/v10"
)

func setupMenuService(t *testing.T) *http.ServeMux {
	clt := data.NewData()
	repo := menu.NewMenuRepo(clt)
	svc := menu.NewMenuService(repo, validator.New())

	mux := http.NewServeMux()
	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)

	hs.RegisterService(api, "/menu", svc)

	return mux
}

func TestMenuGetOne(t *testing.T) {
	mux := setupMenuService(t)

	req := httptest.NewRequest("GET", "/api/menu/one?id=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestMenuGetList(t *testing.T) {
	mux := setupMenuService(t)

	req := httptest.NewRequest("GET", "/api/menu/list", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestMenuPost(t *testing.T) {
	mux := setupMenuService(t)

	newMenu := &menu.MenuCreateReq{Path: "/newmenu", Name: "New Menu", Component: "NewMenu"}
	body, _ := json.Marshal(newMenu)

	req := httptest.NewRequest("POST", "/api/menu", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestMenuPut(t *testing.T) {
	mux := setupMenuService(t)

	updateMenu := &menu.MenuUpdateReq{Path: "/menu1_updated", Name: "Menu1 Updated"}
	body, _ := json.Marshal(updateMenu)

	req := httptest.NewRequest("PUT", "/api/menu", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestMenuDelete(t *testing.T) {
	mux := setupMenuService(t)

	req := httptest.NewRequest("DELETE", "/api/menu?ids=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestMenuGetTree(t *testing.T) {
	mux := setupMenuService(t)

	req := httptest.NewRequest("GET", "/api/menu/tree", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}
