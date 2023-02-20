package bookingbook

import (
	"fmt"
)

type BooksController interface {
	GetBooks(query map[string]string) BooksResponse
}

type booksController struct {
	booksUseCase BooksUseCase
}

func (c booksController) GetBooks(query map[string]string) BooksResponse {
	res := BooksResponse{}
	result, err := c.booksUseCase.GetBooks(query)
	if err != nil {
		fmt.Printf("error")
		return BooksResponse{}
	}
	res.Success = true
	res.MessageTitle = "Get list book by subject"
	res.Message = "Success get list book"
	res.Data = result

	return res
}
