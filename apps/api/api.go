package api

import (
	"net/http"

	"github.com/bobacgo/admin-service/apps/service"
	"github.com/bobacgo/admin-service/pkg/kit/hs/response"
)

type Handler struct {
	User *UserHandler
	Menu *MenuHandler
	I18n *I18nHandler
	Role *RoleHandler
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		User: NewUserHandler(svc),
		Menu: NewMenuHandler(svc),
		I18n: NewI18nHandler(svc),
		Role: NewRoleHandler(svc),
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, response.Resp{
		Code: OK,
		Msg:  "ok",
	})
}
