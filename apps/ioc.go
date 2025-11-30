package apps

import (
	"github.com/bobacgo/admin-service/apps/basic"
	"github.com/bobacgo/admin-service/apps/i18n"
	"github.com/bobacgo/admin-service/apps/menu"
	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/role"
	"github.com/bobacgo/admin-service/apps/user"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	Validator *validator.Validate
	Basic     *basic.Service
	User      *user.UserService
	I18n      *i18n.I18nService
	Menu      *menu.MenuService
	Role      *role.RoleService
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
	userRepo := user.NewUserRepo(clt)
	menuRepo := menu.NewMenuRepo(clt)
	roleRepo := role.NewRoleRepo(clt)
	i18nRepo := i18n.NewI18nRepo(clt)

	// Initialize services
	basicSvc := basic.NewService()
	userSvc := user.NewUserService(userRepo, GetValidator())
	menuSvc := menu.NewMenuService(menuRepo, GetValidator())
	roleSvc := role.NewRoleService(roleRepo, menuRepo, userRepo, GetValidator())
	i18nSvc := i18n.NewI18nService(i18nRepo, GetValidator())

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
