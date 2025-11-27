package i18n

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

func setupI18nService(t *testing.T) (*I18nService, *http.ServeMux) {
	clt := data.NewData()
	repo := NewI18nRepo(clt)
	svc := NewI18nService(repo, validator.New())

	mux := http.NewServeMux()
	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)

	hs.RegisterService(api, "/i18n", svc)

	return svc, mux
}

func TestI18nGetOne(t *testing.T) {
	svc, mux := setupI18nService(t)
	ctx := context.Background()

	i18n := &I18n{Lang: "en", Key: "hello", Value: "Hello"}
	svc.repo.Create(ctx, i18n)

	req := httptest.NewRequest("GET", "/api/i18n/one?id=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestI18nGetOne: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestI18nGetOne passed, response: %s", w.Body.String())
}

func TestI18nGetList(t *testing.T) {
	svc, mux := setupI18nService(t)
	ctx := context.Background()

	svc.repo.Create(ctx, &I18n{Lang: "en", Key: "hello", Value: "Hello"})
	svc.repo.Create(ctx, &I18n{Lang: "zh", Key: "hello", Value: "你好"})

	req := httptest.NewRequest("GET", "/api/i18n/list", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestI18nGetList: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestI18nGetList passed, response: %s", w.Body.String())
}

func TestI18nPost(t *testing.T) {
	_, mux := setupI18nService(t)

	newI18n := &I18nCreateReq{Lang: "en", Key: "goodbye", Value: "Goodbye"}
	body, _ := json.Marshal(newI18n)

	req := httptest.NewRequest("POST", "/api/i18n", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestI18nPost: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestI18nPost passed, response: %s", w.Body.String())
}

func TestI18nPut(t *testing.T) {
	svc, mux := setupI18nService(t)
	ctx := context.Background()

	i18n := &I18n{Lang: "en", Key: "hello", Value: "Hello"}
	svc.repo.Create(ctx, i18n)

	updateI18n := &I18nUpdateReq{ID: 1, Lang: "en", Value: "Hello World"}
	body, _ := json.Marshal(updateI18n)

	req := httptest.NewRequest("PUT", "/api/i18n", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestI18nPut: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestI18nPut passed, response: %s", w.Body.String())
}

func TestI18nDelete(t *testing.T) {
	_, mux := setupI18nService(t)

	req := httptest.NewRequest("DELETE", "/api/i18n?ids=1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestI18nDelete: expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	t.Logf("✅ TestI18nDelete passed, response: %s", w.Body.String())
}
