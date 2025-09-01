package api

import (
	"net/http"

	"github.com/bobacgo/admin-service/pkg/kit-web/response"
	"github.com/bobacgo/admin-service/service"
)

type Handler struct {
	User *UserHandler
	Menu *MenuHandler
	I18n *I18nHandler
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		User: NewUserHandler(svc),
		Menu: NewMenuHandler(svc),
		I18n: NewI18nHandler(svc),
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, response.Resp{
		Code: OK,
		Msg:  "ok",
	})
}
