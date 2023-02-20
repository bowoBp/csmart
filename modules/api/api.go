package api

import (
	"github.com/bowoBp/csmart/utils/httpserver"
)

type Api struct {
	httpServer httpserver.HttpServer
	routers    []Router
}

type Router interface {
	Route(rgh httpserver.RouteHandler)
}

func (a Api) Start() error {
	root := a.httpServer.Group("/")

	for _, router := range a.routers {
		router.Route(root)
	}
	var err error
	if err = a.httpServer.Run(); err != nil {
		return err
	}
	return nil
}

func New(
	httpServer httpserver.HttpServer,
	routers []Router,
) *Api {
	return &Api{
		httpServer: httpServer,
		routers:    routers,
	}
}
