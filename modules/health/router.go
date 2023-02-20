package health

import (
	"github.com/bowoBp/csmart/utils/httpserver"
)

type Router struct {
	requestHandler *RequestHandler
}

func NewRouter() *Router {
	return &Router{
		requestHandler: &RequestHandler{
			repo: versionGetterRepository{},
		},
	}
}

func (r Router) Route(route httpserver.RouteHandler) {
	route.GET("/", r.requestHandler.check)
}
