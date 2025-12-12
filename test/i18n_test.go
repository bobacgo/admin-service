package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobacgo/admin-service/apps/mgr/dto"
	"github.com/bobacgo/admin-service/apps/mgr/repo"
	"github.com/bobacgo/admin-service/apps/mgr/service"
	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/pkg/kit/hs"
	"github.com/go-playground/validator/v10"
)

func setupI18nService() *http.ServeMux {
	clt := data.NewData()
	repo := repo.NewI18nRepo(clt)
	svc := service.NewI18nService(repo, validator.New())

	mux := http.NewServeMux()
	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)

	hs.RegisterService(api, "/i18n", svc)

	return mux
}

func TestI18nGetOne(t *testing.T) {
	mux := setupI18nService()

	req := httptest.NewRequest("GET", "/api/i18n/one?id=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestI18nGetList(t *testing.T) {
	mux := setupI18nService()

	req := httptest.NewRequest("GET", "/api/i18n/list", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestI18nPost(t *testing.T) {
	mux := setupI18nService()

	newI18n := &dto.I18nCreateReq{Lang: "en", Key: "goodbye", Value: "Goodbye"}
	body, _ := json.Marshal(newI18n)

	req := httptest.NewRequest("POST", "/api/i18n", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestI18nPut(t *testing.T) {
	mux := setupI18nService()

	updateI18n := &dto.I18nUpdateReq{ID: 1, Lang: "en", Value: "Hello World"}
	body, _ := json.Marshal(updateI18n)

	req := httptest.NewRequest("PUT", "/api/i18n", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}

func TestI18nDelete(t *testing.T) {
	mux := setupI18nService()
	req := httptest.NewRequest("DELETE", "/api/i18n?ids=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}