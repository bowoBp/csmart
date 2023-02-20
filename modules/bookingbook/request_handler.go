package bookingbook

import (
	"github.com/bowoBp/csmart/utils/httpserver"
	"net/http"
	"strconv"
)

type BooksRequestHandler struct {
	ctrl BooksController
}

func (h BooksRequestHandler) getBook(c httpserver.Context) {
	res := h.ctrl.GetBooks(getQueryParam(c))
	c.JSON(http.StatusOK, res)
}

func getQueryParam(c httpserver.Context) map[string]string {
	query := make(map[string]string)
	if c.Query("limit") == "" {
		query["limit"] = strconv.Itoa(1)
	} else {
		query["limit"] = c.Query("limit")
	}

	if c.Query("offset") == "" {
		query["offset"] = strconv.Itoa(10)
	} else {
		query["offset"] = c.Query("offset")
	}
	query["subject"] = c.Query("subject")
	return query
}
