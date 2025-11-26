package apps

import (
	"github.com/bobacgo/admin-service/apps/api"
	"github.com/bobacgo/admin-service/apps/i18n"
	"github.com/bobacgo/admin-service/apps/menu"
	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/role"
	"github.com/bobacgo/admin-service/apps/service"
	"github.com/bobacgo/admin-service/apps/user"
)

type Container struct {
	clt *data.Client
	svc *service.Service
	api *api.Handler
}

func NewContainer() *Container {
	clt := data.NewData()

	// Initialize repos
	userRepo := user.NewUserRepo(clt)
	menuRepo := menu.NewMenuRepo(clt)
	roleRepo := role.NewRoleRepo(clt)
	i18nRepo := i18n.NewI18nRepo(clt)

	// Initialize services
	userSvc := user.NewUserService(userRepo)
	menuSvc := menu.NewMenuService(menuRepo)
	roleSvc := role.NewRoleService(roleRepo)
	i18nSvc := i18n.NewI18nService(i18nRepo)

	// Initialize service container
	svc := &service.Service{
		Validator: service.GetValidator(),
		User:      userSvc,
		Menu:      menuSvc,
		Role:      roleSvc,
		I18n:      i18nSvc,
	}

	// Initialize API handlers
	apiHandler := api.NewHandler(svc)

	return &Container{
		clt: clt,
		svc: svc,
		api: apiHandler,
	}
}
