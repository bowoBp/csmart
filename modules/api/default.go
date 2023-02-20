package api

import (
	"github.com/bowoBp/csmart/modules/bookingbook"
	"github.com/bowoBp/csmart/modules/health"
	"github.com/bowoBp/csmart/utils/httpserver"
)

func Default() *Api {
	routers := []Router{
		health.NewRouter(),
		bookingbook.NewRouteBooks(),
	}
	return New(
		httpserver.Default(),
		routers,
	)
}
