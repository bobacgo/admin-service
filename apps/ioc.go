package apps

import (
	"github.com/bobacgo/admin-service/apps/basic"
	"github.com/bobacgo/admin-service/apps/common/repo/data"
	"github.com/bobacgo/admin-service/apps/mgr/repo"
	"github.com/bobacgo/admin-service/apps/mgr/service"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	Validator *validator.Validate
	Basic     *basic.Service
	User      *service.UserService
	I18n      *service.I18nService
	Menu      *service.MenuService
	Role      *service.RoleService
}

func GetValidator() *validator.Validate {
	return validator.New()
}

type Container struct {
	clt *data.Client
	svc *Service
}

func NewContainer() *Container {
	clt := data.NewData()

	// Initialize repos
	userRepo := repo.NewUserRepo(clt)
	menuRepo := repo.NewMenuRepo(clt)
	roleRepo := repo.NewRoleRepo(clt)
	i18nRepo := repo.NewI18nRepo(clt)

	// Initialize services
	basicSvc := basic.NewService()
	userSvc := service.NewUserService(userRepo, GetValidator())
	menuSvc := service.NewMenuService(menuRepo, GetValidator())
	roleSvc := service.NewRoleService(roleRepo, menuRepo, userRepo, GetValidator())
	i18nSvc := service.NewI18nService(i18nRepo, GetValidator())

	// Initialize service container
	svc := &Service{
		Validator: GetValidator(),
		Basic:     basicSvc,
		User:      userSvc,
		Menu:      menuSvc,
		Role:      roleSvc,
		I18n:      i18nSvc,
	}

	return &Container{
		clt: clt,
		svc: svc,
	}
}