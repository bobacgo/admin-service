package main

import (
	"github.com/bobacgo/admin-service/api"
	"github.com/bobacgo/admin-service/repo"
	"github.com/bobacgo/admin-service/repo/data"
	"github.com/bobacgo/admin-service/service"
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