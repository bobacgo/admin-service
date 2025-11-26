package api

import (
	"net/http"

	"github.com/bobacgo/admin-service/apps/service"
	"github.com/bobacgo/admin-service/pkg/kit/hs/response"
)

type Handler struct {
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, response.Resp{
		Code: OK,
		Msg:  "ok",
	})
}
