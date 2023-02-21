package bookingbook

import (
	"github.com/bowoBp/csmart/utils/httpserver"
)

type RouterBooks struct {
	requestHandler *BooksRequestHandler
}

func NewRouteBooks() *RouterBooks {
	return &RouterBooks{
		requestHandler: &BooksRequestHandler{
			ctrl: booksController{
				booksUseCase: bookUseCase{},
			},
		},
	}
}

func (h RouterBooks) Route(router httpserver.RouteHandler) {
	//TODO implement me
	basePath := "/online-library"
	onlineLibrary := router.Group(basePath)

	onlineLibrary.GET(
		"/books",
		h.requestHandler.GetBook,
	)

	onlineLibrary.POST(
		"/books",
		h.requestHandler.SubmitBook,
	)

}
