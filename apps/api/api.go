package api

import (
	"net/http"

	"github.com/bobacgo/admin-service/apps/i18n"
	"github.com/bobacgo/admin-service/apps/menu"
	"github.com/bobacgo/admin-service/apps/role"
	"github.com/bobacgo/admin-service/apps/service"
	"github.com/bobacgo/admin-service/apps/user"
	"github.com/bobacgo/admin-service/pkg/kit/hs/response"
)

type Handler struct {
	User *user.UserHandler
	Menu *menu.MenuHandler
	I18n *i18n.I18nHandler
	Role *role.RoleHandler
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		User: user.NewUserHandler(svc.User),
		Menu: menu.NewMenuHandler(svc.Menu, svc.Validator),
		I18n: i18n.NewI18nHandler(svc.I18n, svc.Validator),
		Role: role.NewRoleHandler(svc.Role, svc.Validator),
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, response.Resp{
		Code: OK,
		Msg:  "ok",
	})
}
