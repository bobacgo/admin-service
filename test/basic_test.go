package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobacgo/admin-service/apps/basic"
	"github.com/bobacgo/admin-service/pkg/kit/hs"
	"github.com/bobacgo/admin-service/pkg/kit/hs/response"
)

func assertResponse(t *testing.T, w *httptest.ResponseRecorder) {
	if w.Code != http.StatusOK {
		t.Fatalf("%s: expected status 200, got %d, body: %s", t.Name(), w.Code, w.Body.String())
	}

	var resp response.Resp
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("%s: failed to unmarshal response: %v", t.Name(), err)
	}
	// 检查返回的数据格式
	if resp.Code != 0 {
		t.Fatalf("%s: expected ret_code 0, got %d, msg: %s", t.Name(), resp.Code, resp.Msg)
	}

	t.Logf("✅ %s passed, response: %v", t.Name(), w.Body.String())
}

func setupBasicService() *http.ServeMux {
	mux := http.NewServeMux()
	api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)
	hs.RegisterService(api, "/", basic.NewService())
	return mux
}

func TestHealth(t *testing.T) {
	mux := setupBasicService()

	req := httptest.NewRequest("GET", "/api/health", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	assertResponse(t, w)
}
