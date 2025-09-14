package apps

import (
	"github.com/bobacgo/admin-service/apps/api"
	"github.com/bobacgo/admin-service/apps/repo"
	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/service"
)

type Container struct {
	clt  *data.Client
	repo *repo.Repo
	svc  *service.Service
	api  *api.Handler
}

func NewContainer() *Container {
	clt := data.NewData()
	r := repo.NewRepo(clt)
	s := service.NewService(r)
	a := api.NewHandler(s)
	return &Container{
		clt:  clt,
		repo: r,
		svc:  s,
		api:  a,
	}
}
