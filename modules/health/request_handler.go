package health

import (
	"github.com/bowoBp/csmart/modules/bookingbook"
	"github.com/bowoBp/csmart/utils/httpserver"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestHandler struct {
	// TODO: field
	repo     VersionGetterRepository
	ctrBooks bookingbook.BooksController
}

func (h RequestHandler) check(ctx httpserver.Context) {
	// TODO: check context and version
	version := h.repo.GetVersion()

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "System is healthy",
		"mysqlVersion": version,
	})
}
