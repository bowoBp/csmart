package bookingbook

import (
	"github.com/bowoBp/csmart/utils/httpserver"
	"github.com/bowoBp/csmart/utils/validator"
	"net/http"
	"strconv"
)

type BooksRequestHandler struct {
	ctrl BooksController
}

func (h BooksRequestHandler) GetBook(c httpserver.Context) {
	res := h.ctrl.GetBooks(getQueryParam(c))
	c.JSON(http.StatusOK, res)
}

func (h BooksRequestHandler) SubmitBook(c httpserver.Context) {
	var request BodySubmitBook
	if !validator.BindAndValidateWithAbort(c, &request) {
		return

	}
	res := h.ctrl.SubmitBooks(request)
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
