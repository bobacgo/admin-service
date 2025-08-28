package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/bobacgo/admin-service/pkg/kit-web/response"
	"github.com/bobacgo/admin-service/service"
)

type MenuHandler struct {
	svc *service.Service
}

func NewMenuHandler(svc *service.Service) *MenuHandler {
	return &MenuHandler{svc: svc}
}

func (h *MenuHandler) GetList(w http.ResponseWriter, r *http.Request) {
	// 从文件中读取菜单列表
	bytes, err := os.ReadFile("./api/menu.json")
	if err != nil {
		response.JSON(w, response.Resp{
			Code: ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}

	var menus []any
	if err := json.Unmarshal(bytes, &menus); err != nil {
		response.JSON(w, response.Resp{
			Code: ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}

	response.JSON(w, response.Resp{
		Code: OK,
		Data: map[string]any{"list": menus},
	})
}